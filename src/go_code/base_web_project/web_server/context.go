package web_server

import "net/http"

/**
 * @author wufuqiang
 * @date 2022/7/3 16:51
 */

type Context struct {
	W http.ResponseWriter
	R *http.Request
}
