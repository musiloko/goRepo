package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	files, err := ioutil.ReadDir("D:\\Músicas\\Black Sabbath - 1970 - Paranoid")
	if err != nil{
		log.Fatal(err)
	}
	file, err := os.OpenFile("C:\\Users\\TOTVS\\Documents\\SaidaGO.txt", os.O_WRONLY, 0666)
	if err !=nil{
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := make([]byte,len(files))
	fmt.Println(len(files))
	for _, f := range files {
		fmt.Println(f.Name())
		s := f.Name()+"\n"
		byteSlice = append(byteSlice,s...)
	} 
	bytesWritten , err := file.Write(byteSlice)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote %d ",bytesWritten)
}

