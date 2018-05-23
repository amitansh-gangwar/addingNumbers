package main

import (
	"os"
    "fmt"
	"strconv"
	"encoding/json"
	"errors"
)

func main() {
    
	sum:=0

	numbers:=os.Args[1:]
	if(len(numbers)==0){
		fmt.Println(errors.New("No integers given as input"))
		os.Exit(0)
	}
	for _,number := range numbers{
		integer,error:= strconv.Atoi(number)
		if error==nil{
			sum=sum+integer
		}else{
			fmt.Println(errors.New("Entered numbers must be of type int"))
			os.Exit(0)
		}
	}

	var resultMap map[string]int
	resultMap = make(map[string]int)
	resultMap["result"]=sum
	resultJson,_ := json.Marshal(resultMap)
	ans:=string(resultJson)
	fmt.Println(ans)

}