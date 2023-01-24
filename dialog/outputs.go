package dialog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateStoreDialog() string {
  println("name of the store")
  value:= readUserInput()
  return value
}

func MainMenu() string {
  result:= "What next ? \n"
  result += "1: Add Artiecle to Store\n"
  result += "2: Print all Articles\n"
  fmt.Println(result)
  mod:= readUserInput()
  return mod
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
 costs_input:= readUserInput()
 costs, _ := strconv.ParseFloat(strings.TrimSpace(costs_input), 64)
 println("count of the article")
 count_input:= readUserInput()
 count, _ := strconv.Atoi(strings.TrimSpace(count_input))
 return name, costs, count 
}

func StartOptions ()  {
}
