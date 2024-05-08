package main
import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func names() {
	file, err := os.Open("names.yaml")
	if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	content := bufio.NewScanner(file)
	for content.Scan() {
		fmt.Printf("Your name is: %s\n", content.Text())
	}
}
func main(){
	names()
}