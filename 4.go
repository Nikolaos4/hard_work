package main
import "fmt"
var count int
func serch(arr []string, str string) bool {
    flag := true
    for i := 0; i < len(arr); i++ {
        if arr[i][:len(str)] != str {
            flag = false
        }
    }
    return flag
}
func mi_n(arr []string) int{
    min := len(arr[0])
    for i := 1; i < len(arr); i++ {
        if len(arr[i]) < min {
            min = len(arr[i])
        }
    }
    return min
}
func main() {
    array := []string{}
    fmt.Print("Введите количество элементов: ")
    fmt.Scan(&count)
    for i := 0; i < count; i++ {
        var alfa string
        fmt.Scan(&alfa)
        array = append(array,alfa)
    }
    first_element := array[0]
    str := ""
    str_alfa := first_element
    min := mi_n(array)
    for _, row := range str_alfa {
        str += string(row)
        if len(str) <= min {
            if serch(array,str) == false {
                break
            }
        }else {
            break
        }
    }
    fmt.Print(str[:len(str)-1])
}