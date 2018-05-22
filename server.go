package main

import (
    "fmt"
    "log"
	"net/http"
	"strconv"
	"encoding/json"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method=="POST"{
			requestMessage:=r.URL.Path[5:]
			index:=0
			numberString:=""
			var s []int
			sum:=0
			for ;index<len(requestMessage);{
				for ;index<len(requestMessage) && requestMessage[index]<='9' && requestMessage[index]>='0';{
					numberString=numberString+string(requestMessage[index])
					index++
				}
				index++
				integer,_:=strconv.Atoi(numberString)
				s=append(s,integer)
				sum=sum+integer
				numberString=""
			}
			
			var resultMap map[string]int
			resultMap = make(map[string]int)
			resultMap["result"]=sum
			resultJson,_ := json.Marshal(resultMap)
			ans:=string(resultJson)
			fmt.Fprintf(w,ans)
			header:=r.Header["User-Agent"]
			fmt.Println(header)
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}