//https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem?isFullScreen=true

package main

import(
  "fmt"
  //"os"
  //"math"
)

func series_sum(a uint32,d uint32, l uint32)uint32{
  var n uint32
  if l%d == 0{
    n = (l-d)/d
  }else{
    n = l/d
  }
  return uint32(0.5*float64(n) * (2*float64(a) + (float64(n)-1)*float64(d)))
}


func main(){
	var N uint32 = 100
  fmt.Println(series_sum(3, 3,N) + series_sum(5,5,N)-series_sum(15,15,N))
}
