package platform

import (
	"context"
	"net/http"
)

// 플러그인 핸들러 시그니처
// - ctx: tenant, user 등 컨텍스트 포함
// - core: CoreServices 주입
type PluginHandlerFunc func(ctx context.Context, core *CoreServices, w http.ResponseWriter, r *http.Request)

var routes = map[string]PluginHandlerFunc{}

func RegisterRoute(path string, handler PluginHandlerFunc) {
	routes[path] = handler
}

func GetRoutes() map[string]PluginHandlerFunc {
	return routes
}
