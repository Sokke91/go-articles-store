package main

import (
	"fmt"
	"github.com/Sokke91/go-articles-store/dialog"
	"github.com/Sokke91/go-articles-store/store"
)

type Store struct {
  name string
}


func main()  {
  storename:= dialog.CreateStoreDialog()
  store, _:= store.NewStore(storename)
  fmt.Println(store)
}
