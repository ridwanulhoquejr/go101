package order

type UserProvider interface {
	GetEmail() string
}

type Order struct {
	ID    int
	Owner UserProvider
}

func (o Order) CallUser() {
	o.Owner.GetEmail()
}
