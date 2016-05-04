package main

import "fmt"
import "strconv"

func main() {
    b := 999

    for a:=999 ; a>1 ; a-- {
        productstr := strconv.Itoa(a*b)
        if len(productstr)%2 == 0 {
            midPoint := len(productstr)/2
            firstHalf := productstr[0:midPoint]
            lastHalf := productstr[midPoint:len(productstr)]
            
//            fmt.Println(productstr + "\n")
            fmt.Println(firstHalf + " and " + lastHalf )
        }
        b--
    }

}
 
