package services

import (
  "time"
  "log"
  "net/http"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/models"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
  "github.com/FadhilAhsan/website-health-checker/configs"
  "errors"
  "net/url"
  // "os"
  // "encoding/json"
  // "io/ioutil"
)

var client = http.Client{
		Timeout: time.Millisecond * configs.TIMEOUT_CLIENT_HTTP,
	}

var URLList models.URLMonitors

func AddURL (urlMonitor *models.URLMonitor)(error){
	if !IsURL(urlMonitor.URL){
		return errors.New("URL is not valid")
	}

	um,err := CheckHealthURL(urlMonitor.URL)
	if err != nil {
		log.Println("[Error] AddURL : ",err)
		return err
	}

	err = utils.LoadFromFile(configs.PATH_FILE_CACHE_URL_MONITOR, &URLList)
	if err != nil {
	    log.Println("[Error] AddURL LoadFromFile : ",err)
	    return err
  	}

  	URLList.URLMonitors = append(URLList.URLMonitors,um)
  	err = utils.SaveToFile(configs.PATH_FILE_CACHE_URL_MONITOR,URLList)
  	if err != nil {
  		log.Println("[Error] AddURL SaveToFile : ",err)
	    return err
  	}
	log.Println("[Success] AddURL : ",um.URL)
	urlMonitor.StatusOK = um.StatusOK
	return nil
}

func GetAllURL() (models.URLMonitors,error){
	err := utils.LoadFromFile(configs.PATH_FILE_CACHE_URL_MONITOR, &URLList)
	if err != nil {
	    log.Println("[Error] GetAllURL LoadFromFile : ",err)
	    return models.URLMonitors{},err
  	}
  	return URLList,nil
}

func IsURL(urlString string)(bool){
	u, err := url.Parse(urlString)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func CheckHealthURL(urlString string)(models.URLMonitor,error){
	statusOK := false
	resp, err := client.Get(urlString)
	if err != nil {
		log.Println("[Error] CheckHealthURL : ", err)
	}else if resp.StatusCode == 200 {
		log.Println("[Success] CheckHealthURL : ", 200)
		statusOK =true
	}else{
		log.Println("[Failed] CheckHealthURL : ", urlString)
	}
	um := models.URLMonitor{urlString,statusOK}
	return um,nil
}


