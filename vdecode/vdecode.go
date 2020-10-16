package main

import (
  "fmt"
  "github.com/atotto/clipboard"
  "encoding/base64"
  "strings"
)

func main(){
//  clipboard.WriteAll("日本語")
text, _ := clipboard.ReadAll()
data := []byte(text)
data2 := "<start>some data</start><end>some more data</end><andsomemore>there is another data set</andsomemor>"
fmt.Println(strings.Contains(data2, "some"))

// use replace to replace input
fmt.Println(strings.Replace(data2, "some", "SOME", 2))

fmt.Println(strings.SplitN(data2, "/", -1))

decoded := dec(data)
fmt.Println(decoded)
}

func dec(inStr []byte) string{
//	data := []byte("")
	str := base64.StdEncoding.EncodeToString(inStr)
  //decoded := fmt.Println("test")
  return str
}
