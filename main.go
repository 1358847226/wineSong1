package main

import (
	. "wineSong/connect"
	"wineSong/router"
)

func main() {
	defer Db.Close()
	router.InitRouter()
}
