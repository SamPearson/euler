package main

// What is the largest prime factor of the number 600851475143 ?
import "fmt"

func isPrime(product int) bool {
    for possFactor := 2 ; possFactor < product ; possFactor ++ {
        if (product%possFactor==0){
            return false 
        }
    }
    return true
}    

func main() {
    n := 600851475143
    for denom := 2; denom<n/2 ; denom ++{
        if n%denom == 0{
            if isPrime(n/denom) {
                fmt.Println(n/denom)
            }
        }
    }
}
 
