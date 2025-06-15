package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf(headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//允许所有来源 (*)
			c.Header("Access-Control-Allow-Origin", "*")
			//允许多种请求方式
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// 请求头字段列表
			// Origin - 指示请求的来源（协议、域名和端口）。
			// X-Requested-With - 通常用于标识 AJAX 请求，常见的值是 XMLHttpRequest。
			// Content-Type - 指定请求体的 MIME 类型，例如 application/json 或 application/x-www-form-urlencoded。
			// Accept - 指定客户端能够处理的内容类型。
			// Authorization - 通常用于传递认证信息，如 Bearer Token。
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			// 用于指定服务器允许客户端通过 JavaScript 访问的响应头字段列表。
			// Content-Length - 指示响应体的大小（以字节为单位）。
			// Access-Control-Allow-Origin - 指示哪些域可以访问资源。
			// Access-Control-Allow-Headers - 指示服务器允许的请求头。
			// Cache-Control - 指示缓存机制，如 no-cache 或 max-age。
			// Content-Language - 指示响应内容的语言。
			// Content-Type - 指示响应体的 MIME 类型。
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			//缓存请求信息 单位秒
			c.Header("Access-Control-Max-Age", "172800")
			// 跨域请求是否携带token
			c.Header("Access-Control-Allow-Credentials", "false")
			//返回格式为json
			c.Set("content-type", "application/json")
		}
		// 方向所有Options请求
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request")
		}
		// 处理请求
		c.Next()
	}
}
