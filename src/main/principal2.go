package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"bufio"
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
	fmt.Println(len(files))
	//w := bufio.NewWriter(file)
	for _, f := range files {
		byteSlice = append(byteSlice, (f.Name()+"\n")...)
		//fmt.Println(f.Name())
		//w.WriteString(f.Name())
		//w.WriteString("\n")
	} 
	//err = w.Flush()
	
	if err != nil{
		log.Fatal(err)
	}
}

