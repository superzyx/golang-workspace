package web


type Filter func(c *Context)
type FilterBuilder func(f Filter)Filter
