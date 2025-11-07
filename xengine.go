package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/tonly18/xgin/iface"
	"github.com/tonly18/xgin/xglobal"
)

type XGinEngine struct {
	xRouteGroup *XRouteGroup
	ginEngine   *gin.Engine
}

func NewXGinEngine(opts ...gin.OptionFunc) *XGinEngine {
	//设置gin的运行模式: debug、release、test
	switch xglobal.GinMode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	ginEngine := gin.New(opts...)

	//默认中间件
	ginEngine.Use(defaultMiddleware())

	//return
	return &XGinEngine{
		xRouteGroup: &XRouteGroup{
			ginRouteGroup: &ginEngine.RouterGroup,
		},
		ginEngine: ginEngine,
	}
}

func (e *XGinEngine) GetGinEngine() *gin.Engine {
	return e.ginEngine
}

func (e *XGinEngine) Use(middleware ...gin.HandlerFunc) *XGinEngine {
	e.ginEngine.Use(middleware...)

	return e
}

func (e *XGinEngine) SetTrustedProxies(trustedProxies []string) error {
	return e.ginEngine.SetTrustedProxies(trustedProxies)
}

func (e *XGinEngine) Run(addr ...string) error {
	return e.ginEngine.Run(addr...)
}

func (e *XGinEngine) Group(path string, handlers ...gin.HandlerFunc) *XRouteGroup {
	e.xRouteGroup.ginRouteGroup = e.ginEngine.Group(path, handlers...)

	return e.xRouteGroup
}

func (e *XGinEngine) GET(relativePath string, handler iface.IHandler) gin.IRoutes {
	return e.xRouteGroup.GET(relativePath, handler)
}

func (g *XGinEngine) POST(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.POST(relativePath, handler)
}

func (g *XGinEngine) PUT(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.PUT(relativePath, handler)
}

func (g *XGinEngine) DELETE(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.DELETE(relativePath, handler)
}

func (g *XGinEngine) HEAD(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.HEAD(relativePath, handler)
}

func (g *XGinEngine) OPTIONS(relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.OPTIONS(relativePath, handler)
}

func (g *XGinEngine) Match(method []string, relativePath string, handler iface.IHandler) gin.IRoutes {
	return g.xRouteGroup.Match(method, relativePath, handler)
}
