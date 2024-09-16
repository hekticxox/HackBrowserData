package main

import (
    "fmt"
    "myproject/pkg/utils"  // Ensure this path is correct
)

func main() {
    result := utils.SomeFunction()  // This function should be defined in pkg/utils
    fmt.Println(result)
}

