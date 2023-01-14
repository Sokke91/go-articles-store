package store

import "fmt"


type Article struct{
  name string
  costs float64
  count int
}

type Store struct {
  name string
  articles []Article
}

func NewStore(name string) (Store, error) {
   if len(name) == 0 {
    return Store{} , fmt.Errorf("invalid Storename %s", name) 
  } 
  return Store{name: name}, nil
}

func NewArticle(name string, costs float64, count int) (Article, error) {
  if name == "" || costs < 0 || count < 0 {
    return Article{}, fmt.Errorf("invalid data")
  }
  return Article{name: name, costs: costs, count: count}, nil
}


func (s *Store) addArticle(article Article){
  s.articles = append(s.articles, article)
}

func (s Store) String() string {
  result:= fmt.Sprintf("ArticleStore: %s", s.name)
  result += "test"
  return result
}

