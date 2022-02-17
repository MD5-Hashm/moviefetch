package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func Get(inputtext string) string {
	fmt.Printf(inputtext)
	scanner.Scan()
	var name string = scanner.Text()
	return name
}

func GetInt(inputtext string, limit int) (int, int) {
	if limit <= 0 {
		for {
			fmt.Printf(inputtext)
			scanner.Scan()
			var str string = scanner.Text()
			if valid.IsInt(str) {
				out, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				return out, 0
			} else {
				fmt.Println("Not a int :(")
			}
		}
	} else {
		for {
			fmt.Printf(inputtext)
			scanner.Scan()
			var str string = scanner.Text()
			if valid.IsInt(str) {
				out, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				if out > limit || out <= 0 {
					fmt.Println("Out of Range :(")
				} else {
					return out, 0
				}
			} else {
				fmt.Println("Not an int :(")
			}
		}
	}
}
