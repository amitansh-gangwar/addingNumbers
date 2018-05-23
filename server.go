package main

import (
    "fmt"
    "log"
	"net/http"
	"strconv"
	"encoding/json"
	"strings"
)

type solution struct{
	Result int
	Success string
}

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST"{
			
			requestMessage := r.URL.Path[5:]
			
			var numbers []string
			numbers = strings.Split(requestMessage,"/")

			sum := 0
			errortype:=""
			for _,number := range numbers{
				integer,error := strconv.Atoi(number)
				if error!=nil {
					errortype="type error"
					break
				}
				sum = sum + integer
			}

			var multiplyBy int
			multiplyBy,_=strconv.Atoi(r.Header.Get("multiply"))
			
			multiplication := sum * multiplyBy

			if len(requestMessage)==0 {
				fmt.Fprintf(w,"No integers given as input")
			}else if errortype=="type error" {
				fmt.Fprintf(w,"Entered numbers must be of type int")	
			}else{
				answer := solution{multiplication,"ok"}
				resultJson,_ := json.Marshal(answer)
				ans := string(resultJson)
				fmt.Fprintf(w,ans)
			}
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}