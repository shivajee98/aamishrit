package routes

import "github.com/shivajee98/aamishrit/internal/handlers"

type AdminDeps struct {
	ProductHandler  *handlers.ProductHandler
	UserHandler     *handlers.UserHandler
	OrderHandler    *handlers.OrderHandler
	CategoryHandler *handlers.CategoryHandler
}
