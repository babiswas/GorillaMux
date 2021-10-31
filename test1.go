package main
import "io"
import "log"
import "net/http"
import "time"

func HealthCheck(w http.ResponseWriter,req *http.Request){
  currentTime:=time.Now()
  io.WriteString(w,currentTime.String())
}

func main(){
  http.HandleFunc("/health",HealthCheck)
  log.Fatal(http.ListenAndServe(":8000",nil))
}