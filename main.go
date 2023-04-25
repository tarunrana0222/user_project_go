package main

import (
	configs "github.com/tarunrana0222/user_project_go/config"
	"github.com/tarunrana0222/user_project_go/routes"
)

func init() {
	if !configs.EnvLoaded {
		configs.LoadEnv()
	}
}

func main() {
	r := configs.GetRouter()
	routes.UserRoutes(r)
	routes.ClientRoutes(r)
	r.Run(configs.Port)
}
