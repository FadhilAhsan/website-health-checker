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
    http.HandleFunc("/api/url", controllers.URLMonitorGetAll)
    http.HandleFunc("/api/url/add", controllers.URLMonitorPost)
    http.Handle("/assets/",http.StripPrefix("/assets/", http.FileServer(http.Dir("../../web/static/assets/"))))
    err := http.ListenAndServe(configs.PORT_SERVER, nil)
    log.Fatal(err)
}

func main(){
	fmt.Println("Application is running on port : ", configs.PORT_SERVER)
	go utils.ProcessTask()
	handleRequests()
}