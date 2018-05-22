package main

import (
	"os"
    "fmt"
	"strconv"
	"encoding/json"
)

func main() {
    
	sum:=0

	numbers:=os.Args[1:]
	for _,number := range numbers{
		integer,_:= strconv.Atoi(number)
		sum=sum+integer
	}

	var resultMap map[string]int
	resultMap = make(map[string]int)
	resultMap["result"]=sum
	resultJson,_ := json.Marshal(resultMap)
	ans:=string(resultJson)
	fmt.Println(ans)

}