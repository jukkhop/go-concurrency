package main

import "fmt"

func main() {
    thread_count := 10
    channel := make(chan result)

    for i := 0; i < thread_count; i++ {
        go fibonacci(channel, i)
    }

    for i := 0; i < thread_count; i++ {
        r := <- channel
        s := fmt.Sprintf("Thread %d returned result %d", r.thread_no, r.sum)
        fmt.Println(s)
    }
}

func fibonacci(channel chan result, thread_no int) {
    var a uint64 = 0
    var b uint64 = 1
    var sum uint64 = 0

    for i := 0; i < 92; i++ {
        sum = a + b
        a = b
        b = sum
    }

    channel <- result { sum: sum, thread_no: thread_no }
}

type result struct {
    sum uint64
    thread_no int
}