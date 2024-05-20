package router

type RouteNotFound struct {
	Code int
}

func (e RouteNotFound) GetError() (int, string) {
	return e.Code, "Route not found"
}

func RouteNotFoundError() RouteNotFound {
	return RouteNotFound{Code: 404}
}
