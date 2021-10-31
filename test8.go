package main
import "fmt"
import "net/http"

func middleware(originalhandler http.Handler)http.Handler{
   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
     fmt.Println("Executing middleware before request")
     originalhandler.ServeHTTP(w,r)
     fmt.Println("Executing middleware after request")  
   })
}


func handle(w http.ResponseWriter,r *http.Request){
  fmt.Println("Executing main handler")
  w.Write([]byte("ok"))
}

func main(){
  originalhandler:=http.HandlerFunc(handle)
  http.Handle("/",middleware(originalhandler))
  http.ListenAndServe(":8000",nil)
}
