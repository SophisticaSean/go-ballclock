package main

import "fmt"
import "os"
import "strconv"
import "github.com/pkg/profile"
// import "runtime/pprof"

func main() {
  defer profile.Start().Stop()
  temp := os.Args[1]
  i, _ := strconv.Atoi(temp)
  var queue = make([]int,0)
  var minute = make([]int,0)
  var five = make([]int,0)
  var hours = make([]int,0)

  j := 1
  for j <= i {
    // fmt.Println(j)
    queue = append(queue, j)
    j = j + 1
  }

  pristine := []int{}
  pristine = append(pristine, queue...)
  loop_count := 0
  for {
    //if loop_count > 4 {
      //fmt.Println(minute)
      //fmt.Println(five)
      //fmt.Println(hours)
      //fmt.Println(queue)
      //break
    //}
    var ball = queue[0]
    loop_count = loop_count + 1
    //var rollover = false
    queue = append(queue[:0], queue[1:]...)
    if len(minute) == 4 {
      for j := len(minute) - 1; j >= 0; j-- {
        //fmt.Println(j)
        ball_to_move := minute[j]
        minute = minute[:len(minute)-1]
        queue = append(queue, ball_to_move)
      }
      if len(five) == 11 {
        for j := len(five) - 1; j >= 0; j-- {
          //fmt.Println(j)
          ball_to_move := five[j]
          five = five[:len(five)-1]
          queue = append(queue, ball_to_move)
        }
        if len(hours) == 11 {
          for j := len(hours) - 1; j >= 0; j-- {
            //fmt.Println(j)
            ball_to_move := hours[j]
            hours = hours[:len(hours)-1]
            queue = append(queue, ball_to_move)
          }
          queue = append(queue, ball)
        } else {
          hours = append(hours, ball)
        }
      } else {
        five = append(five, ball)
      }
    } else {
      minute = append(minute, ball)
    }
    //fmt.Println(minute)
    //fmt.Println(five)
    //fmt.Println(hours)
    //fmt.Println(queue)
    if loop_count % 720 == 0 && len(queue) == len(pristine){
      should_break := false
      for j := len(pristine) - 1; j >= 0; j-- {
        equal := pristine[j] == queue[j]
        if j == 0 && equal == true {
          should_break = true
        }
        if equal != true {
          j = 0
        }
      }
      if should_break == true {
        fmt.Println(minute)
        fmt.Println(five)
        fmt.Println(hours)
        fmt.Println(queue)
        fmt.Println(pristine)
        fmt.Println((((loop_count)/60)/12)/2)
        break
      }
    }
  }
}
