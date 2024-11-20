package main
import "fmt"
var x int
func main() {
  fmt.Println("Введите год:")
  fmt.Scan(&x)
  if (x%4 == 0 && x%100 != 0) || (x%400 == 0) {
       fmt.Println("Год весокосный")
  } else {
      fmt.Println("Год не весокосный")
  }
  
}