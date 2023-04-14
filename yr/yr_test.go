package yr

import (
	"bufio"
	"encoding/csv"
	"math"
	"os"
	"reflect"
	"testing"
)

func lineCount(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}

func TestLineCount(t *testing.T) {
	want := 16756
	got, err := lineCount("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("error counting lines in file: %v", err)
	}
	if got != want {
		t.Errorf("incorrect number of lines in file: got %d, want %d", got, want)
	}
}

func TestConvertTemperatures(t *testing.T) {

	ConvertTemperatures()

	// Åpner output fil
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Leser data fra output fil
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Finn linjen med data vi vil teste
	var lineToTest []string
	for _, line := range lines {
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "18.03.2022 01:50" {
			lineToTest = line
			break
		}
	}

	// Sjekk om linjen ble funnet
	if lineToTest == nil {
		t.Fatal("Line to test not found")
	}

	// Expected output data
	want := []string{"Kjevik", "SN39040", "18.03.2022 01:50", "42.80"}

	// Sjekk om linjen er lik det vi forventer
	if !reflect.DeepEqual(lineToTest, want) {
		t.Errorf("got %v, want %v", lineToTest, want)
	}
}

func TestConvertTemperatures2(t *testing.T) {

	ConvertTemperatures()

	// Åpner output fil
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Leser data fra output fil
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Finn linjen med data vi vil teste
	var lineToTest []string
	for _, line := range lines {
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "07.03.2023 18:20" {
			lineToTest = line
			break
		}
	}

	// Sjekk om linjen ble funnet
	if lineToTest == nil {
		t.Fatal("Line to test not found")
	}

	// Forventet output data
	want := []string{"Kjevik", "SN39040", "07.03.2023 18:20", "32.00"}

	// Sjekk om linjen er lik det vi forventer
	if !reflect.DeepEqual(lineToTest, want) {
		t.Errorf("got %v, want %v", lineToTest, want)
	}
}

func TestConvertTemperatures3(t *testing.T) {

	ConvertTemperatures()

	// Åpner output fil
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Leser data fra output fil
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Finn linjen med data vi vil teste
	var lineToTest []string
	for _, line := range lines {
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "08.03.2023 02:20" {
			lineToTest = line
			break
		}
	}

	// Sjekk om linjen ble funnet
	if lineToTest == nil {
		t.Fatal("Line to test not found")
	}

	// Forventet output data
	want := []string{"Kjevik", "SN39040", "08.03.2023 02:20", "12.20"}

	// Compare the result with the expected output
	if !reflect.DeepEqual(lineToTest, want) {
		t.Errorf("got %v, want %v", lineToTest, want)
	}
}

func TestConvertSisteLinje(t *testing.T) {

	ConvertTemperatures()

	// Åpner output fil
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Leser data fra output fil
	reader := csv.NewReader(file)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Finn linjen med data vi vil teste
	var lineToTest []string
	for _, line := range lines {
		if line[0] == "Data er gyldig per 18.03.2023 (CC BY 4.0)" && line[1] == "Meteorologisk institutt (MET)" && line[2] == "endringen er gjort av" {
			lineToTest = line
			break
		}
	}

	// Sjekk om linjen ble funnet
	if lineToTest == nil {
		t.Fatal("Line to test not found")
	}

	// Forventet output data
	want := []string{"Data er gyldig per 18.03.2023 (CC BY 4.0)", "Meteorologisk institutt (MET)", "endringen er gjort av", "Adrian Viken"}

	// Sjekk om linjen er lik det vi forventer
	if !reflect.DeepEqual(lineToTest, want) {
		t.Errorf("got %v, want %v", lineToTest, want)
	}
}

// Funksjon for å sjekke om to tall er like innenfor en viss feilmargin
func withinTolerance(a, b, error float64) bool {
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	if b == 0 {
		return difference < error
	}
	return (difference / math.Abs(b)) < error
}
func TestAverageTemperature(t *testing.T) {

	type test struct {
		want float64
	}
	tests := []test{
		{want: 8.56},
	}
	for _, tc := range tests {
		got := AverageTemperature("C")

		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("AverageTemperature() = %v, want %v", got, tc.want)
		}

	}
}
