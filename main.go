package main

import (
  "math"
  "sync"
)

func isPrime(n int) bool {
  if n <= 2 {
    return false
  }
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ { //Algorithm to check for prime numbers
    if n%i == 0 {
      return false
    }
  }
  return true
}
func primeGenerator(start, end int, ch chan<- int, wg *sync.WaitGroup) {
  defer wg.Done() //Makes sure the goroutine signals completion

  for i := start; i <= end; i++ {
    if isPrime(i) {
      ch <- i //Checks if i is prime and sends it to the channel
    }
  }
}

func conPrimes(limit int, numWorkers int) []int {
  
}
func main() {

}
