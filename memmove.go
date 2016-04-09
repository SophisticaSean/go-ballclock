package main

//import "fmt"
//import "os"
//import "strconv"
import "github.com/pkg/profile"

func main() {
  defer profile.Start().Stop()

  loop := 156751200
  //to_copy := make([]int,11)
  //to_copy[0] = 1
  //fmt.Println(to_copy)
  //new_copy := make([]*int,0)

  for loop > 0 {
    //item := &to_copy[0]
    //to_copy = append(to_copy[:0], to_copy[1:]...)
    //new_copy = append(new_copy, item)
    loop = loop - 1
  }
}
