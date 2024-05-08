package entity

type Order struct {
	id        int
	userId    int
	productId int
}

func NewOrder(id int, userId int, productId int) Order {
	return Order{
		id:        id,
		userId:    userId,
		productId: productId,
	}
}

func (o *Order) GetId() int {
	return o.id
}

func (o *Order) GetUserId() int {
	return o.userId
}

func (o *Order) GetProductId() int {
	return o.productId
}
