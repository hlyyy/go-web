package main

import (
	"fmt"
	"github.com/hly/ginny1/model"
	"github.com/hly/ginny1/router"
)

func main() {
	model.Init()
	defer model.DB.Close()

	router.InitRouter()
	router.Router.Run()
	fmt.Println("Running...Successful!")
}
