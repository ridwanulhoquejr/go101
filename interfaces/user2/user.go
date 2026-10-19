package user2

import "example.com/go101/interfaces/order"

var _ order.UserProvider = User{}

type User struct {
	ID    int
	Email string
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) PlaceOrder(orderID int) order.Order {
	return order.Order{
		ID:    orderID,
		Owner: u,
	}
}
