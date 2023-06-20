package main
import (
    "fmt"
    "os"
)
func main() {
    fmt.Printf("Hello %s!\n", os.Getenv("USR_FROM_ENV_FILE"))
    fmt.Printf("Hello %s!\n", os.Getenv("USR_WITH_DEFAULT"))
    fmt.Printf("Hello %s!\n", os.Getenv("USR"))
}