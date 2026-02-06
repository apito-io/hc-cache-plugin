package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-cache-plugin] Starting Cache plugin...")
	plugin := sdk.Init("hc-cache-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("CacheStatus", "Cache plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getCacheStatus",
		sdk.ComplexObjectField("Get cache plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-cache-plugin] Plugin ready")
	plugin.Serve()
}
