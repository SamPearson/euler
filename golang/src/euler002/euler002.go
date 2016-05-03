package main

import "fmt"

func main() {
    low := 1
    high := 2
    sum := 0

    for high < 4000000 {
        if (high%2 == 0 ){
            sum += high 
        }
        high = low + high
        low = high - low
    }
    fmt.Println(sum)
}
 

