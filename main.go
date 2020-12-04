package main 

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().Unix())
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d回目のおみくじ結果: ", i)
    }
}