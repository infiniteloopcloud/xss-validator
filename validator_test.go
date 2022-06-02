package xssvalidator

import (
	"bufio"
	"os"
	"testing"
)

func TestValidatorVulnerableList(t *testing.T) {
	file, err := os.Open("test/payloads.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	line := 0
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	for scanner.Scan() {
		line++
		input := scanner.Text()
		err := Validate(input)
		if err == nil {
			t.Fatalf("Line: %d, input: %s", line, input)
		}
	}

}
