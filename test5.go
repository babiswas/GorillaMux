package main
import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "time"
import "log"

func ArticleHandler(w http.ResponseWriter,r *http.Request){
  fmt.Println("Article handler executed")
  vars:=mux.Vars(r)
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w,"Category is:%v\n",vars["category"])
  fmt.Fprintf(w,"Id is:%v\n",vars["id"])
}

func main(){
   fmt.Println("Started Server")
   r:=mux.NewRouter()
   r.HandleFunc("/articles/{category}/{id:[0-9]+}",ArticleHandler)
   srv:=&http.Server{Handler:r,Addr:"127.0.0.1:8002",WriteTimeout:15*time.Second,ReadTimeout:15*time.Second,}
   log.Fatal(srv.ListenAndServe())
}