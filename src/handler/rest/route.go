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
	var result []RoutePrefix

	userRoutes := RoutePrefix{
		"/user",
		[]Route{
			{
				"GetUserByID",
				"GET",
				"/{userId}",
				e.GetUserByID,
				false,
			},
		},
	}

	result = append(result, userRoutes)

	return result
}
