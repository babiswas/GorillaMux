package main
import "net/http"


func main(){
  newmux:=http.NewServeMux()
  newmux.HandleFunc("/static/*filepath",http.Dir("./"))
  http.ListenAndServe(":8002",newmux)
}