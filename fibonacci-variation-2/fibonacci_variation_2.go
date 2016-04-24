package main

import "fmt"

func main() {
    thread_count := 10
    part_result_ch := make(chan result)
    continue_ch := make(chan bool)
    result_ch := make(chan result)

    for i := 0; i < thread_count; i++ {
        go fibonacci(part_result_ch, continue_ch, result_ch, i)
    }

    for i := 0; i < thread_count; i++ {
        r := <- part_result_ch
        s := fmt.Sprintf("Thread %d returned part result %d", r.thread_no, r.sum)
        fmt.Println(s)
    }

    fmt.Println("Received part result from every thread, signaling threads to continue")

    for i := 0; i < thread_count; i++ {
        continue_ch <- true;
    }

    for i := 0; i < thread_count; i++ {
        r := <- result_ch
        s := fmt.Sprintf("Thread %d returned result %d", r.thread_no, r.sum)
        fmt.Println(s)
    }
}

func fibonacci(part_result_ch chan result,
               continue_ch chan bool,
               result_ch chan result,
               thread_no int) {

    var a uint64 = 0
    var b uint64 = 1
    var sum uint64 = 0

    for i := 0; i < 92; i++ {
        sum = a + b
        a = b
        b = sum

        if i == 46 {
            part_result_ch <- result { sum: sum, thread_no: thread_no }
            <- continue_ch
        }
    }

    result_ch <- result { sum: sum, thread_no: thread_no }
}

type result struct {
    sum uint64
    thread_no int
}