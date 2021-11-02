package store

type Store interface {
	User() UserRepository
	Orders() OrdersRepository
}
