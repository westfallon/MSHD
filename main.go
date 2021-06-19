package main

import (
	"github.com/westfallon/naseum/db"
	"github.com/westfallon/naseum/gin_router"
)

func main() {
	db.Init()
	gin_router.InitRouter()
}
