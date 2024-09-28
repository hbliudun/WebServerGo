package web

import (
	"fmt"
	"strings"
)

type node struct {
	path     string
	children map[string]*node
	// 加一个通配符匹配
	starChild *node
	handler   HandleFunc
}

type router struct {
	trees map[string]*node
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

func (r *router) addRouter(method string, path string, handleFunc HandleFunc) {
	// 规则
	// /开头
	// 不能以/结尾
	if path == "" {
		panic("web: 路径不能为空")
	}
	if path[0] != '/' {
		panic("web: 路径要以/开始")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路径不能以/结束")
	}

	// 找到树
	root, ok := r.trees[method]
	if !ok {
		// 没有根节点则创建
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}

	if path == "/" {
		if root.handler != nil {
			panic("根节点重复注册")
		}
		root.handler = handleFunc
		return
	}

	sigs := strings.Split(path[1:], "/")
	for _, seg := range sigs {
		child := root.childOrCreate(seg)
		root = child
	}
	//r.trees[method].PrintNode()
	if root.handler != nil {
		panic("该节点已注册")
	}
	root.handler = handleFunc
}

// /user
func (n *node) childOrCreate(sig string) *node {
	if sig == "*" {
		n.starChild = &node{
			path: sig,
		}
		return n.starChild
	}
	if n.children == nil {
		n.children = map[string]*node{}
	}
	nd, ok := n.children[sig]
	if !ok {
		nd = &node{
			path: sig,
		}
		n.children[sig] = nd
	}
	return nd
}

func (r *router) findRouter(method string, path string) (*node, bool) {

	if path == "" {
		return nil, false
	}
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}

	if path == "/" && root.path == "/" {
		return root, true
	}

	sigs := strings.Split(path[1:], "/")
	for _, seg := range sigs {
		if root == nil {
			return nil, false
		}
		if root.children == nil {
			return root.starChild, root.starChild != nil
		}
		root = root.children[seg]
	}
	return root, root != nil
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	child, ok := n.children[path]
	return child, ok
}

func (n *node) PrintNode() {
	if n == nil {
		return
	}

	fmt.Printf("%v", n.path)

	if n.children == nil {
		return
	}
	child, ok := n.children[n.path]
	if ok {
		child.PrintNode()
	}
}
