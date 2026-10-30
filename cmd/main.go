package main

import (
	"fmt"
	"log"

	"github.com/LuanSilva9/hub-estudos-backend/conf"
	"github.com/LuanSilva9/hub-estudos-backend/conf/db"
	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/LuanSilva9/hub-estudos-backend/pkg/shared/modules/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := conf.LoadGlobalConfig()

	if err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente: ", err.Error())
	}

	db, errs := db.PostgresConnection(cfg.DatabaseDsn, cfg.Env)

	if errs != nil {
		log.Fatal("Erro ao conectar no banco")
	}

	r := gin.Default()

	auth.InitAuthModule(r.Group("/auth"), db)
	users.InitUserModule(r.Group("/users"), db, auth.RequireAuth())

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
