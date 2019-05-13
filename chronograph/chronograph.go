package chronograph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptTimestamp() string {
	output := fmt.Sprintf("%s: ", promptTimestamp)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(output)
	input, err := reader.ReadString('\n')
	if ok := err == nil; ok != true {
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
}
