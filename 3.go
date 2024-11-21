package main
import (
    "fmt"
)
var count int
func main() {
    fmt.Println("Введите количество строк: ")
    fmt.Scan(&count)
    array := make([][]int, count)
    for i := range array {
        array[i] = make([]int, count)
    }
    for i := 0; i < count; i++ {
        for j := 0; j < count; j++ {
            if j == 0 {
                array[i][j] = 1
            } else {
                array[i][j] = 0
            }
        }
    }
    for i := 1; i < count; i++ {
        for j := 1; j < count; j++ {
            array[i][j] = array[i-1][j-1] + array[i-1][j]
        }
    }
    fmt.Println("Результат: ")
    for _, row := range array {
        for _, value := range row {
            if value == 0 {
                fmt.Print("") 
            } else {
                fmt.Print(value, " ") 
            }
        }
        fmt.Println()
    }
    
}