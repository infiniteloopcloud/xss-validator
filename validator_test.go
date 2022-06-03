package xssvalidator

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

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
		err := Validate(input, DefaultRules...)
		if err == nil {
			t.Errorf("Line: %d, input: %s", line, input)
		}
	}
}

func TestValidateForRightData(t *testing.T) {
	valueTester(t, gofakeit.Name)
	valueTester(t, gofakeit.Email)
	valueTester(t, gofakeit.Phone)
	valueTester(t, gofakeit.PhoneFormatted)
	valueTester(t, gofakeit.BS)
	valueTester(t, gofakeit.BeerName)
	valueTester(t, gofakeit.Company)
	valueTester(t, gofakeit.HackerPhrase)
	valueTester(t, func() string {
		return gofakeit.Password(true, true, true, true, true, 40)
	})
	valueTester(t, gofakeit.State)
	valueTester(t, gofakeit.Street)
	valueTester(t, gofakeit.StreetNumber)
	valueTester(t, gofakeit.Gamertag)
	valueTester(t, gofakeit.CarMaker)
	valueTester(t, gofakeit.UUID)
	valueTester(t, gofakeit.HexColor)
	valueTester(t, gofakeit.URL)
	valueTester(t, gofakeit.MacAddress)
	valueTester(t, gofakeit.ChromeUserAgent)
	valueTester(t, gofakeit.SafariUserAgent)
	valueTester(t, gofakeit.BitcoinAddress)
}

func valueTester(t *testing.T, fn func() string) {
	for i := 0; i < 20; i++ {
		v := fn()
		if err := Validate(v, DefaultRules...); err != nil {
			t.Errorf("Error should be nill for value: %s, err: %s", v, err.Error())
		}
	}
}

func BenchmarkValidate(b *testing.B) {
	input := `<IMG SRC="javascript:alert('XSS')"`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := Validate(input, DefaultRules...)
		if err == nil {
			b.Error(err)
		}
	}
}
