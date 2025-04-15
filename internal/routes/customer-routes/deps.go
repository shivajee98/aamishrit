package routes

import "github.com/shivajee98/aamishrit/internal/handlers"

type Deps struct {
	UserHandler     *handlers.UserHandler
	ProductHandler  *handlers.ProductHandler
	CartHandler     *handlers.CartHandler
	ReviewHandler   *handlers.ReviewHandler
	AddressHandler  *handlers.AddressHandler
	OrderHandler    *handlers.OrderHandler
	CategoryHandler *handlers.CategoryHandler
}
