package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/tonly18/xgin/iface"
	"github.com/tonly18/xgin/wrapper"
)

type XRouteGroup struct {
	ginRouteGroup *gin.RouterGroup
}

func (g *XRouteGroup) GET(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.GET(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) POST(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.POST(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) PUT(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.PUT(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) DELETE(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.DELETE(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) HEAD(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.HEAD(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) OPTIONS(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.OPTIONS(relativePath, wrapper.HandlerFuncWrapper(handler))
}

func (g *XRouteGroup) Match(method []string, relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.ginRouteGroup.Match(method, relativePath, wrapper.HandlerFuncWrapper(handler))
}
