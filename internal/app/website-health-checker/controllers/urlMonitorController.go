package controllers

import (
  "log"
  "net/http"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/services"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/models"
)

func URLMonitorPage(w http.ResponseWriter, r *http.Request) {
	log.Println("URLMonitor Page")
    
	pageDate := models.URLMonitorPage{"Web Health Checker","",models.URLMonitors{}}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
            log.Println("[Error] URLMonitorPage ParseForm() err : ", err)
            pageDate.ErrorMessage = "Input Invalid."
        }
        urlString := r.FormValue("url")
        err = services.AddURL(urlString)
        if err != nil {
        	log.Println("[Error] URLMonitorPage AddURL err : ", err)
          pageDate.ErrorMessage = "Input Invalid."
        }
	}

  urls,err := services.GetAllURL()
	if err != nil {
		pageDate.ErrorMessage = "Can't load data." 
	}else {
    pageDate.URLs = urls
  }

  utils.RenderTemplate(w,"urlmonitor",pageDate)
}

