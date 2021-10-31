package main
import "github.com/gorilla/handlers"
import "github.com/gorilla/mux"
import "log"
import "os"
import "net/http"


func handle(w http.ResponseWriter,r *http.Request){
   log.Println("Processing request!")
   w.Write([]byte("OK"))
   log.Println("Finished Processing")
}

func main(){
  r:=mux.NewRouter()
  r.HandleFunc("/",handle)
  loggedrouter:=handlers.LoggingHandler(os.Stdout,r)
  http.ListenAndServe(":8000",loggedrouter)
}