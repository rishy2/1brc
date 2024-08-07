package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

var measurementsMap = make(map[string][4]float64)

// var fileToRead = "../../../test/resources/samples/measurements-3.txt"

var fileToRead2 = "../../../test/resources/samples/measurements-10000-unique-keys.txt"

// var measurementsFile = "../../../../data/measurements.txt"

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create("./profiles/" + *cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	file, err := os.Open(fileToRead2)

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
			measurementsMap[station] = val // can we modify in place?
		} else {
			// [min, max, sum, count]
			measurementsMap[station] = [4]float64{measurement, measurement, measurement, 1}
		}
	}

	out := "{"
	for k, v := range measurementsMap {
		out += fmt.Sprintf("%s=%0.1f/%0.1f/%0.1f, ", k, v[0], RoundUp(v[2]/v[3]), v[1])
	}
	out = out[:len(out)-2] + "}"
	fmt.Println(out)
	fmt.Println("Time taken:", time.Since(start))

	if *memprofile != "" {
		f, err := os.Create("./profiles/" + *memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func RoundUp(num float64) float64 {
	return math.Round(num*10) / 10
}
