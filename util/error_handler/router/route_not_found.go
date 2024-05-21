package router

type RouteNotFound struct {
	Code int
	Path string
}

func (e RouteNotFound) GetError() (int, string) {
	return e.Code, "Route not found"
}

func RouteNotFoundError(path string) RouteNotFound {
	return RouteNotFound{Code: 404, Path: path}
}
