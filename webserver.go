package webserver

import "net/http"
import "fmt"


func Start() {
  var port string 
  fmt.Println("Port>> ")
  fmt.Scanln(&port)
  http.ListenAndServe("0.0.0.0:"+port, nil)
 
  

}
