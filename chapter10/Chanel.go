package main

import (
    "fmt"
    "time"
)

func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong" // c <- Otpravka
    }
}

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping" // c <- Otpravka
    }
}
func printer(c chan string) {
    for {
        msg := <- c // <-c - Poluchenie
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
func main() {
    var c chan string = make(chan string) // chan - kanal

    go pinger(c)
    go printer(c)
    go ponger(c)
    
    var input string
    fmt.Scanln(&input)
}