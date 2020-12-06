package main 

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().Unix())

    number := rand.Intn(6)
    switch number {
        case 0:
            fmt.Println("凶です")
        case 1, 2:
            fmt.Println("吉です")
        case 3, 4:
            fmt.Println("中吉です")
        case 5:
            fmt.Println("大吉です")
    }
}