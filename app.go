package main

// import (
// 	"gin-test/server"
// )


// func main() {
// 	r := server.SetupRouter()
// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }

import (
	"gin-test/config"
	"gin-test/models"
	"gin-test/server"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
	 fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	r := server.SetupRouter()
	//running
	r.Run()
   }