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
			requestMessage:=r.URL.Path[5:]
			index:=0
			firstNumber:=""
			secondNumber:=""
			for ;requestMessage[index]<='9' && requestMessage[index]>='0';{
				firstNumber=firstNumber+string(requestMessage[index])
				index++
			}
			index++
			for ;index<len(requestMessage) && requestMessage[index]<='9' && requestMessage[index]>='0';{
				secondNumber=secondNumber+string(requestMessage[index])
				index++
			}
			
			firstInteger,_:=strconv.Atoi(firstNumber)
			
			secondInteger,_:=strconv.Atoi(secondNumber)
			sum:=firstInteger+secondInteger
			sumString:=strconv.Itoa(sum)
			fmt.Fprintf(w, sumString)
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}