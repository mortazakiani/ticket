package api

import (
	db "github.com/mortazakiani/ticket/internal/db/sqlc"
	"github.com/mortazakiani/ticket/internal/utiles"
)
type server struct{

}
func Newserver(config utiles.Config , store *db.Store)(*server,error){
server := &server{}
return server,nil
}