package main

import (
	"fmt"
	"log"
	"github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/controllers"
	"github.com/FadhilAhsan/website-health-checker/configs"
	"github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
	"net/http"
)


func handleRequests() {
    http.HandleFunc("/", controllers.URLMonitorPage)
    // http.HandleFunc("/add", controllers.HomePage)
    err := http.ListenAndServe(configs.PORT_SERVER, nil)
    log.Fatal(err)
}

func main(){
	fmt.Println("Application is running on port : ", configs.PORT_SERVER)
	go utils.ProcessTask()
	handleRequests()
}