package rest

import "net/http"

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

func (e *rest) InitRoute() []RoutePrefix {
	var routes []RoutePrefix

	accountRoutes := RoutePrefix{
		"/login",
		[]Route{
			{
				"AccountLogin",
				"POST",
				"",
				e.AccountLogin,
				false,
			},
		},
	}

	userRoutes := RoutePrefix{
		"/user",
		[]Route{
			{
				"GetUserByID",
				"GET",
				"/{userID}",
				e.GetUserByID,
				false,
			},
		},
	}

	productRoutes := RoutePrefix{
		"/product",
		[]Route{
			{
				"GetProductByID",
				"GET",
				"/{productID}",
				e.GetProductByID,
				false,
			},
		},
	}

	routes = append(routes, accountRoutes, userRoutes, productRoutes)

	return routes
}
