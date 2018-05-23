package main

import (
    "fmt"
    "log"
	"net/http"
	"strconv"
	"encoding/json"
	"strings"
)

func main1() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST"{
			errortype:=""
			requestMessage := r.URL.Path[5:]
			
			var numbers []string
			numbers = strings.Split(requestMessage,"/")
			
			sum := 0
			for _,number := range numbers{
				integer,error := strconv.Atoi(number)
				if error!=nil {
					errortype="type error"
					break
				}
				sum = sum + integer
			}

			if len(requestMessage)==0 {
				fmt.Fprintf(w,"No integers given as input")
			}else if errortype=="type error" {
				fmt.Fprintf(w,"Entered numbers must be of type int")	
			}else{
				var resultMap map[string]int
				resultMap = make(map[string]int)
				resultMap["result"] = sum
				//resultMap["success"]="ok"
				resultJson,_ := json.Marshal(resultMap)
				ans := string(resultJson)
				fmt.Fprintf(w,ans)
			}
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}