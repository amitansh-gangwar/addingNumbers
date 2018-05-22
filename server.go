package main

import (
    "fmt"
    "log"
	"net/http"
	"strconv"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method=="POST"{
			request:=r.URL.Path[5:]
			i:=0
			a:=""
			b:=""
			for ;request[i]<='9' && request[i]>='0';{
				a=a+string(request[i])
				i++
			}
			i++
			for ;i<len(request) && request[i]<='9' && request[i]>='0';{
				b=b+string(request[i])
				i++
			}
			fmt.Println(a)
			fmt.Println(b)
			//a:=request[5:6]
			ai,_:=strconv.Atoi(a)
			//b:=request[7:8]
			bi,_:=strconv.Atoi(b)
			ci:=ai+bi
			c:=strconv.Itoa(ci)
			fmt.Fprintf(w, c)
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}