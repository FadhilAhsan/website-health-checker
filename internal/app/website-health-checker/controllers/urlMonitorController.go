package controllers

import (
  "log"
  "net/http"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/services"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/models"
  "encoding/json"
)

func URLMonitorPage(w http.ResponseWriter, r *http.Request) {
	pageDate := models.Page{"Web Health Checker",""}
  utils.RenderTemplate(w,"urlmonitor",pageDate)
  return
}

func URLMonitorGetAll(w http.ResponseWriter, r *http.Request){

  meta := models.JSONResponseMeta{500,"Failed","Failed get all URL"}
  data, err := services.GetAllURL()
  if err != nil {
    log.Println("[Error] URLMonitorPage AddURL err : ", err)
  }else{
    meta = models.JSONResponseMeta{200,"Success","Success get all URL"}
  }
  
  response := models.JSONResponse{meta,data}
  responseJson, _ := json.Marshal(response)
  w.Header().Set("Content-Type","application/json")
  w.Write(responseJson)

  return
}

func URLMonitorPost(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST" {
    meta := models.JSONResponseMeta{500,"Failed","Invalid URL"}
    log.Println("Request Body: ",r.Body)
    decoder := json.NewDecoder(r.Body)
    log.Println("Request Body decode: ",decoder)
    var url models.URLMonitor
    err := decoder.Decode(&url)
    if err != nil {
        log.Println("[Error] URLMonitorPost Decode() err : ", err)
    }else{
      err = services.AddURL(&url)
      if err != nil {
        log.Println("[Error] URLMonitorPage AddURL err : ", err)
      }else{
        meta = models.JSONResponseMeta{200,"Success","Success Check URL"}
      }
    }
    response := models.JSONResponse{meta,url}
    responseJson, _ := json.Marshal(response)
    w.Header().Set("Content-Type","application/json")
    w.Write(responseJson)
  }
  return
}

