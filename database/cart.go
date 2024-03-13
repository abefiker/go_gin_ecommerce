package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart(context.Context,*mongo.Collection,*mongo.Collection,primitive.ObjectID,string) {

}

func RemoveCartItem(context.Context,*mongo.Collection,*mongo.Collection,primitive.ObjectID,string) {

}

func GetCartItem(){

}
func BuyItemFromCart() {

}

func InstantBuyer(context.Context,*mongo.Collection,*mongo.Collection,primitive.ObjectID,string) {

}
