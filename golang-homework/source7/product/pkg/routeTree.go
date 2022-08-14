package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var httpMethod = [4]string{
	http.MethodGet,
	http.MethodPut,
	http.MethodPost,
	http.MethodHead,
}

type MethodOnTree interface {
	ServeHTTP(c *Context)
}

func (t *Tree) ServeHTTP(c *Context) {
	path := c.R.URL.Path
	method := c.R.Method

	if root, ok := t.root[method]; !ok {
		panic("no method")
	}else {
		funcName := root.findNode(path)
		if funcName == nil{
			func(c *Context){
				c.W.WriteHeader(http.StatusNotFound)
			}(c)
			return
		}
		time.Sleep(time.Second)
		funcName(c)
	}
}

type SetOnTree interface {
	isNullOnTree()bool
	addNode(path string, method HandleMethod)
	deleteNode()
	findNode(path string) HandleMethod
}

func (n *Node) isNullOnTree() bool{
	if n == nil {
		panic("error")
	}
	if n.path == "" {
		return false
	}
	return true
}

func (n *Node) addNode(path string, method HandleMethod) {
	tmp := n
	pathSlice := strings.Split(path, "/")
	for _, i := range pathSlice[1:] {
		if val, ok := tmp.child[i]; ok {
			tmp = val
		}else {
			node := NewNode(i)
			if tmp.child == nil {
				tmp.child = make(map[string]*Node, 1)
			}
			tmp.child[i] = node
			tmp = node
		}
	}
	tmp.method = method

}

func (n *Node) deleteNode() {
	panic("implement me")
}

func (n *Node) findNode(path string) HandleMethod{
	tmp := n
	pathSlice := strings.Split(path, "/")
	fmt.Println(pathSlice)
	for _, i := range pathSlice[1:] {
		if val, ok := tmp.child[i]; ok {
			tmp = val
		}else {
			fmt.Printf("No find route like %v \n", i)
			return nil
		}
	}
	return tmp.method
}

type Node struct {
	path string
	child map[string]*Node
	method HandleMethod
}

type Tree struct {
	root map[string]*Node
}

func NewNode(path string) *Node{
	return &Node{path: path}
}

func NewTree() *Tree{
	tree := &Tree{make(map[string]*Node, 4)}
	for _, method := range httpMethod {
		tree.root[method] = NewNode("/")
	}
	return tree
}
