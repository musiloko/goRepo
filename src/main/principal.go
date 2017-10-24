package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	files, err := ioutil.ReadDir("C:\\Users\\TOTVS\\goWorkspace\\src\\main")
	if err != nil{
		log.Fatal(err)
	}
	
	for _, f := range files {
		fmt.Println(f.Name())
	} 
}

