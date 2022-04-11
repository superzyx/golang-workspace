package web

//type Handler interface {
//	ServeHTTP(c *Context)
//}

type HandleMethod func(c *Context)