package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"Restaurant-Management/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIDIsNotValid   = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add review to cart")
	ErrCantRemoveItem     = errors.New("cannot remove item from cart")
	ErrCantGetItem        = errors.New("cannot get item from cart ")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func FoodReviews() gin.HandlerFunc {
	return func(c *gin.Context) {
		food_id := c.Query("food_id")
		if food_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid code"})
			c.Abort()
			return
		}
		food, err := primitive.ObjectIDFromHex(food_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var addReview models.Review
		if err = c.BindJSON(&addReview); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: food}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$reviews"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$food_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		pointcursor, err := foodCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var reviewinfo []bson.M
		if err = pointcursor.All(ctx, &reviewinfo); err != nil {
			panic(err)
		}

		//var size int32
		// for _, address_no := range reviewinfo {
		// 	count := address_no["count"]
		// 	size = count.(int32)
		// }

		// if size < 2{
		// 	filter := bson.D{primitive.E{Key: "_id", Value: food}}
		// 	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "reviews", Value: addReview}}}}
		// 	_, err := foodCollection.UpdateOne(ctx, filter, update)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// } else {
		// 	c.IndentedJSON(400, "Not Allowed")
		// }

		filter := bson.D{primitive.E{Key: "_id", Value: food}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "reviews", Value: addReview}}}}
		_, err = foodCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(400, "Not Allowed")
			return
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(201, "Review Added")
	}
}
