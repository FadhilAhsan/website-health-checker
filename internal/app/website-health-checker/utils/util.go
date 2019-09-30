package utils

import (
  "time"
  "github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/models"
  "github.com/FadhilAhsan/website-health-checker/configs"
  "log"
  "os"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "io"
  "net/http"
  "path/filepath"
  "sync"
  "bytes"
  "html/template"
)

var client = http.Client{
  Timeout: time.Millisecond * configs.TIMEOUT_CLIENT_HTTP,
}

var lock sync.Mutex
var URLList models.URLMonitors

func isDoNow(t time.Time)(isOk bool){
  return int8(t.Minute()) % configs.TIME_SCHEDULE_CHECK_HEALTH_URL == 0
}

func ProcessTask() {  
  for {
    now := time.Now()
      if isDoNow(now) {
        log.Println("Start Health URL Monitor")
        go checkAllHealthURL()
      }
    time.Sleep(time.Second * 45)
  }
}

func checkAllHealthURL(){
  log.Println("Start check health URL")
  err := LoadFromFile(configs.PATH_FILE_CACHE_URL_MONITOR, &URLList)
  if err != nil {
    log.Println("[Error] CheckAllHealthURL LoadFromFile : ",err)
    return
  }else{
    for i := 0; i < len(URLList.URLMonitors); i++ {
      fmt.Println("URL: " + URLList.URLMonitors[i].URL)
      URLList.URLMonitors[i].StatusOK = false
      resp, err := client.Get(URLList.URLMonitors[i].URL)
      if err != nil {
        log.Println("[Error] CheckAllHealthURL : ", err)
      }else if resp.StatusCode == 200 {
        log.Println("[Success] CheckAllHealthURL : ", 200)
        URLList.URLMonitors[i].StatusOK = true
      }else{
        log.Println("[Failed] CheckAllHealthURL : ", URLList.URLMonitors[i].URL)
      }
    }

    if err := SaveToFile(configs.PATH_FILE_CACHE_URL_MONITOR, &URLList); err != nil {
      log.Fatalln(err)
    }
    return
  }
}

func SaveToFile(path string, obj interface{}) error {
  lock.Lock()
  defer lock.Unlock()
  absPath, err := filepath.Abs(path)
  file, err := os.Create(absPath)
  if err != nil {
    log.Println("[Error] Create File : ",err)
    return err
  }
  defer file.Close()
  r, err := marshalObjectJson(obj)
  if err != nil {
    log.Println("[Error] Marshal Object Json : ",err)
    return err
  }
  _, err = io.Copy(file, r)
  if err != nil {
    log.Println("[Error] Copy() : ",err)
  }
  return err
}

func marshalObjectJson(obj interface{}) (io.Reader, error) {
  b, err := json.MarshalIndent(obj, "", "\t")
  if err != nil {
    return nil, err
  }
  return bytes.NewReader(b), nil
}


func LoadFromFile(path string,obj interface{})(error){
  absPath, err := filepath.Abs(path)
  if err != nil {
    log.Println("[Error] LoadFromFile absolute path : ",err)
    return err
  }

  jsonFile, err := os.Open(absPath)
  defer jsonFile.Close()
  if err != nil {
    log.Println("[Error] LoadFromFile open file : ",err)
    return err
  }

  log.Println("Successfully Opened URLMonitor.json")
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, obj)
  return nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, in interface{}) {
  absPath, err := filepath.Abs(configs.PATH_TEMPLATES_PAGE+tmpl + ".html")
    t, err := template.ParseFiles(absPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, in)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}