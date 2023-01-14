package dialog

import (
	"bufio"
	"encoding/base64"
	"os"
)

func CreateStoreDialog() string {
  println("name of the store")
  value:= readUserInput()
  return value
}


func readUserInput() string {
  reader:= bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  return input
}

func CreateArticleDialog() (string, float64, int){
 println("name of the article")
 name:= readUserInput()
 println("costs of the product")
 costs:= readUserInput()
}

func StartOptions ()  {
}
