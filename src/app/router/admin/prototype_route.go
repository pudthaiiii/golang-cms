package route

import (
	controller "go-ibooking/src/app/http/admin/controllers/prototype"
	"go-ibooking/src/types"
)

func addPrototypeRoute(c controller.PrototypeController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "Paginate",
			Method:      "GET",
			Pattern:     "/prototype",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Paginate,
		},

		types.Route{
			Name:        "Paginate",
			Method:      "POST",
			Pattern:     "/prototype",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Create,
		},
	}
}
