package main

import "fmt"

func main() {
	api := getDepReqApi()
	api.put("test", "test")
	val, err := api.get("test")	
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(val)
}
