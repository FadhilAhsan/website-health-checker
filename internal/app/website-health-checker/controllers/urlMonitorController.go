package controllers

import (
  "log"
  "net/http"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/services"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
)

func URLMonitorPage(w http.ResponseWriter, r *http.Request) {
	log.Println("URLMonitor Page")
    
	
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
            log.Println("[Error] URLMonitorPage ParseForm() err : ", err)
            return
        }
        urlString := r.FormValue("url")
        err = services.AddURL(urlString)
        if err != nil {
        	log.Println("[Error] URLMonitorPage AddURL err : ", err)
            return
        }
	}

  urls,err := services.GetAllURL()
	if err != nil {
		return 
	}

  utils.RenderTemplate(w,"urlmonitor",urls)
}

