package web_server

import (
	"net/http"
	"strings"
)

/**
 * @author wufuqiang
 * @date 2022/7/3 17:43
 */

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handlerFunc handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root

	for index, path := range paths {
		// 查找当前节点的子节点
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else { // 没有找到子节点
			// 创建新的节点链路
			h.createSubTree(cur, paths[index:], handlerFunc)
			break
		}
	}
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFunc handlerFunc) {
	cur := root

	for _, path := range paths {
		nn := newNode(path)
		// 在孩子节点中追回新的子节点
		cur.children = append(cur.children, nn)
		cur = nn
	}

	cur.handlerFunc = handlerFunc

}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {

	var wildcardNode *node

	for _, child := range root.children {
		if child.path == path && child.path != "*" {
			return child, true
		}

		// 命中通配符
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")

	cur := h.root

	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前path的节点
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}

	// 到这里，已经是找到了子节点
	if cur.handlerFunc == nil {
		// 到达这里是因为这种场景
		// 注册了/user/profile
		// 然后你访问 /user
		return nil, false
	}
	return cur.handlerFunc, true
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}
