package main

import "fmt"
import "os"
import "strconv"
//import "github.com/pkg/profile"

func main() {
  //defer profile.Start().Stop()
  debug := false
  temp := os.Args[1]
  i, _ := strconv.Atoi(temp)
  zero_val := 0
  // zero pointer to use for emptying our pointer slices
  zero := &zero_val
  zero_arr := [11]*int{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero}
  var queue = make([]*int,200)
  minute := zero_arr[0:4]
  var five = make([]*int,11)
  var hours = make([]*int,11)

  pristine := []int{}
  queue = queue[:i]
  j := 1
  for j <= i {
    // fmt.Println(j)
    pristine = append(pristine, j)
    queue[j-1] = &pristine[j-1]
    j = j + 1
  }
  for j := len(minute) - 1; j >= 0; j-- {
    //minute[j] = zero
  }

  for j := len(five) - 1; j >= 0; j-- {
    five[j] = zero
    hours[j] = zero
  }

  p_minute := make([]*int,len(minute))
  copy(p_minute, minute)
  p_five := make([]*int,0)
  copy(five, p_five)
  p_hours := make([]*int,0)
  copy(hours, p_hours)

  fmt.Println(p_minute)
  //minute[0] = &queue[0]
  loop_count := 0
  //cur_ball_index := 0
  cur_ball := queue[0]
  for {
    //if cur_ball_index == i {
      //cur_ball_index= 0
      //break
    //}

    if loop_count > 5 && debug == true{
      fmt.Println(len(queue))

      var temp = make([]int,0)
      for j := 0; j < len(queue); j++ {
        temp = append(temp, *queue[j])
      }
      fmt.Println(temp)

      temp = make([]int,0)
      for j := len(minute) - 1; j >= 0; j-- {
        temp = append(temp, *minute[j])
      }
      fmt.Println(temp)

      temp = make([]int,0)
      for j := len(five) - 1; j >= 0; j-- {
        temp = append(temp, *five[j])
      }
      fmt.Println(temp)

      temp = make([]int,0)
      for j := len(hours) - 1; j >= 0; j-- {
        temp = append(temp, *hours[j])
      }
      fmt.Println(temp)
      break
    }

    loop_count = loop_count + 1
    cur_ball = queue[0]
    queue = append(queue[:0], queue[1:]...)
    //queue[cur_ball_index] = zero

    if minute[0] != zero {
      queue = append(queue, minute...)
      copy(minute, p_minute)

      if five[0] != zero {
        for j := 0; j < len(five); j++ {
          queue = append(queue, five[j])
          five[j] = zero
        }
        if hours[0] != zero {
          for j := 0; j < len(hours); j++ {
            queue = append(queue, hours[j])
            hours[j] = zero
          }
          queue = append(queue, cur_ball)
        } else {
          for j := len(hours) - 1; j >= 0; j-- {
            if hours[j] == zero {
              hours[j] = cur_ball
              j = 0
            }
          }
        }
      } else {
        for j := len(five) - 1; j >= 0; j-- {
          if five[j] == zero {
            five[j] = cur_ball
            j = 0
          }
        }
      }
    } else {
      for j := len(minute) - 1; j >= 0; j-- {
        if minute[j] == zero {
          minute[j] = cur_ball
          j = 0
        }
      }
    }

    //cur_ball_index = cur_ball_index+ 1
    //fmt.Println(cur_ball)
        //fmt.Println(len(minute))
        //fmt.Println(five)
        //fmt.Println(hours)
        //fmt.Println("kek")
        //fmt.Println(len(queue))
    if loop_count % 720 == 0 && len(queue) == len(pristine){
      should_break := false
      for j := len(pristine) - 1; j >= 0; j-- {
        equal := pristine[j] == *queue[j]
        if j == 0 && equal == true {
          should_break = true
        }
        if equal != true {
          j = 0
        }
      }
      if should_break == true {
        //fmt.Println(minute)
        //fmt.Println(five)
        //fmt.Println(hours)
        //fmt.Println(queue)
        //fmt.Println(pristine)
        fmt.Println((((loop_count)/60)/12)/2)
        //fmt.Println(loop_count)
        //var temp = make([]int,0)
        //for j := len(queue) - 1; j >= 0; j-- {
          //temp = append(temp, *queue[j])
        //}
        //fmt.Println(temp)
        //fmt.Println(pristine)
        break
      }
    }
  }
}
