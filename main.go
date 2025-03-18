package main

import (
  "math"
)

func isPrime(n int) bool {
  if n <= 2 {
    return false
  }
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n%2 == 0 {
      return false
    }
  }
  return true
}
func primeGenerator {
  
}
func main(){

}
