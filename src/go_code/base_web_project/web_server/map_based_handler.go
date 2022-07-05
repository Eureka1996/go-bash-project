package web_server

/**
 * @author wufuqiang
 * @date 2022/7/3 21:46
 */

type Routable interface {
	Route(method string, pattern string, handlerFunc handlerFunc)
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}
