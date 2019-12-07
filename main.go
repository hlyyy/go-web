package main

import(
	"github.com/hly/ginny1/router"
	"github.com/hly/ginny1/model"
	"fmt"
)

func main() {
	model.DB.Init()
	defer model.DB.Close()
	router.InitRouter()
	router.Router.Run()
	fmt.Println("Running...Successful!")
}
