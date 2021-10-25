package main

import "fmt"

func makeOddGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i+1
        i += 2
        return
    }
}
func main() {
    nextEven := makeOddGenerator()
    fmt.Println(nextEven()) // 1
    fmt.Println(nextEven()) // 3
    fmt.Println(nextEven()) // 5
}