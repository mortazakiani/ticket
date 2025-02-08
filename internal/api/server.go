package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/mortazakiani/ticket/internal/db/sqlc"
	"github.com/mortazakiani/ticket/internal/utiles"
)
type Server struct{
 Config utiles.Config
 Store *db.Store
 App  *fiber.App
}
func Newserver(config utiles.Config , store *db.Store)(*Server,error){
	app := fiber.New()
server := &Server{
	App: app,
	Config: config,
	Store: store,
}
return server,nil
}


func (server *Server)  Start(add string) error {
	return server.App.Listen(fmt.Sprintf(":%s",add))
}