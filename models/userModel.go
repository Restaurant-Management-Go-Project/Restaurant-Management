package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	First_name      *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name       *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password        *string            `json:"Password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	User_type       *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token   *string            `json:"refresh_token"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	User_id         string             `json:"user_id"`
	UserCart        []FoodUser         `json:"usercart"`
	Address_Details []Address          `json:"address"`
	Order_Status    []Order            `json:"orders"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []Food             `json:"order_list"  bson:"order_list"`
	Orderered_At   time.Time          `json:"ordered_on"  bson:"ordered_on"`
	Price          int64              `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount"    bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod"     bson:"cod"`
}

type FoodUser struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      int64              `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      int64              `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"menu_id"`
}
