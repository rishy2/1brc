package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var measurementsMap = make(map[string][4]float64)

func main() {
	file, err := os.Open("../../../test/resources/samples/measurements-3.txt")

	if err != nil {
		log.Fatal("Error opening file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	for _, record := range records {
		var station string = record[0]
		measurement, _ := strconv.ParseFloat(record[1], 64)

		val, ok := measurementsMap[station]
		if ok {
			val[3]++
			val[2] += measurement
			val[1] = max(val[1], measurement)
			val[0] = min(val[0], measurement)
			measurementsMap[station] = val
		} else {
			// [min, max, sum, count]
			measurementsMap[station] = [4]float64{measurement, measurement, measurement, 1}
		}
	}

	fmt.Print("{")
	for k, v := range measurementsMap {
		fmt.Printf("%s=%0.1f/%0.1f/%0.1f, ", k, v[0], v[2]/v[3], v[1])
	}
	fmt.Print("}")
}
