package main
import "fmt"

// In one package, Go allows only ONE main() function.
// You currently have two main() functions in the same package (usually package main).
func main () {
	i:=1
	for i <=5 {
        fmt.Print(i, " ") // Print without newline
        i++
    }
    fmt.Println() // Move to the next line after loop ends
}


// Use "go run ." to run all files in a folder .