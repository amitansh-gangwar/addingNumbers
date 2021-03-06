package main

import (
	"os"
    "fmt"
	"strconv"
	"encoding/json"
	"errors"
)

type result struct{
	Addition int
	Success string
}

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

	answer := result{sum,"ok"}
	resultJson,_ := json.Marshal(answer)
	ans:=string(resultJson)
	fmt.Println(ans)

}