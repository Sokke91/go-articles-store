package main

import (
	"fmt"

	"github.com/Sokke91/go-articles-store/dialog"
	"github.com/Sokke91/go-articles-store/store"
)



func main(){
store_name := dialog.CreateStoreDialog()  
store, _ := store.NewStore(store_name)
fmt.Println(store)
for {
menu:= dialog.MainMenu()
fmt.Printf("Auswahl %s", menu)
createArticleAndAddToStore(store)
fmt.Println(store)

  }
}


func createArticleAndAddToStore(s *store.Article_store)  {
  name, costs, count := dialog.CreateArticleDialog()
  article, _ := store.NewArticle(name, costs , count )
  s.AddArticle(article)
}



