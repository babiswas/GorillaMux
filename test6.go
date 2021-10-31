package main
import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "time"
import "log"


func ArticleHandler(w http.ResponseWriter,r *http.Request){
  fmt.Println("Article handler executed")
  queryparam:=r.URL.Query()
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w,"Got parameter id:%s!\n",queryparam["id"][0])
  fmt.Fprintf(w,"Got parameter category:%s!",queryparam["category"][0])
}

func main(){
   fmt.Println("Started a server")
   r:=mux.NewRouter()
   r.HandleFunc("/articles",ArticleHandler)
   srv:=&http.Server{Handler:r,Addr:"127.0.0.1:8003",WriteTimeout:15*time.Second,ReadTimeout:15*time.Second,}
   log.Fatal(srv.ListenAndServe())
}