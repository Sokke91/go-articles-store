package store

import "fmt"


type Article struct{
  name string
  costs float64
  count int
}

type Article_store struct {
  name string
  articles []Article
}

func NewStore(name string) (*Article_store, error) {
   if len(name) == 0 {
    return &Article_store{} , fmt.Errorf("invalid Storename %s", name) 
  } 
  return &Article_store{name: name}, nil
}

func NewArticle(name string, costs float64, count int) (Article, error) {
  if name == "" || costs < 0 || count < 0 {
    return Article{}, fmt.Errorf("invalid data")
  }
  return Article{name: name, costs: costs, count: count}, nil
}


func (s *Article_store) AddArticle(article Article){
  s.articles = append(s.articles, article)
}

func (s Article_store) String() string {
  result:= fmt.Sprintf("ArticleStore: %s", s.name)
  result += "---------------------\n"
  result += "Artikels: \n"
  for _ ,v := range s.articles {
    result += fmt.Sprintf("Name: %s \n", v.name)
    result += fmt.Sprintf("Costs: %f \n", v.costs)
    result += fmt.Sprintf("Stock: %d \n", v.count)
    result += "-------------------"
  }
  return result
}

