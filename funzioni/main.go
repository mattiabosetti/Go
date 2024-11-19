package main

import (
  "errors"
  "fmt"
)

func main(){
  num := 4
  div := 0
  var ris, res, err = divisione(num, div)

  switch  {
  case err !=nil:
    fmt.Printf(err.Error())
  case res==0:
    fmt.Printf("Il risultato della divisione è: %v", ris)
  default:
    fmt.Printf("\nIl risultato della divisione è %v con %v di resto\n", ris, res)
  }
}

func divisione(numeratore int, denominatore int) (int, int, error) {
  var err error
  if(denominatore == 0){
    err = errors.New("Non puoi dividere per 0\n")
    return 0,0, err
  }
  var risultato int = numeratore / denominatore 
  var resto int = numeratore % denominatore

  return risultato, resto, err
}

