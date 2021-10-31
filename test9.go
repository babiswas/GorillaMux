package main
import "encoding/json"
import "github.com/justinas/alice"
import "net/http"
import "fmt"
import "log"

type City struct{
  Name string `json:"name"`
  Area int64 `json:"area"`
}


func setservertimecookie(handler http.Handler)http.Handler{
   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
     handler.ServeHTTP(w,r)
     cookie:=http.Cookie{}
     http.SetCookie(w,&cookie)
     log.Println("Currently server time middleware")
   })
}


func filtercontenttype(handler http.Handler)http.Handler{
   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
     log.Println("Verify content middleware")
     if r.Header.Get("Content-type")!="application/json"{
       w.WriteHeader(http.StatusUnsupportedMediaType)
       w.Write([]byte("415-Unsupported media type"))
       return
     }
     handler.ServeHTTP(w,r)
   })
}

func posthandler(w http.ResponseWriter,r *http.Request){
  fmt.Println("Executing post handler")
  if r.Method=="POST"{
    fmt.Println("Executing post method")
    var tempcity City
    decoder:=json.NewDecoder(r.Body)
    err:=decoder.Decode(&tempcity)
    if err!=nil{
      panic(err)
    }
    defer r.Body.Close()
    fmt.Printf("Got %s city with area %d",tempcity.Name,tempcity.Area)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("201-Status OK"))
  }else{
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("405-Method not allowed"))
  }
}

func main(){
   fmt.Println("Server Started")
   originalhandler:=http.HandlerFunc(posthandler)
   chain:=alice.New(filtercontenttype,setservertimecookie).Then(originalhandler)
   http.Handle("/city",chain)
   http.ListenAndServe(":8000",nil)
}