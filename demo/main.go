package main

import (
	"github.com/gin-gonic/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	_ "github.com/chenhg5/go-admin/adapter/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "0000",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "demo.go-admin.cn",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "/data/www/go-admin/uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
		INDEX:    "/",
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "/data/www/go-admin/uploads")

	// you can custom your pages like:

	r.GET("/"+cfg.PREFIX+"/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	r.Run(":9033")
}
