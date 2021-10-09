package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//	"io/ioutil"
)

func main() {
	// data, err := ioutil.ReadFile("log-00.txt")
	file, err := os.Open("log-00.txt")
	if err != nil {
		log.Fatalf("Failed opening the file: %s", err)
		return
	}
	fmt.Println("File opened.")

	fileOut, err := os.Create("output.xlsx")
	if err != nil {
		fmt.Println("Cannot create the output file")
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	i := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i++
	}
	file.Close()
	fmt.Println("Done reading: ", i, " lines")
	var gpu0Lines []string

	// Parsing lines
	for _, line := range lines {
		if strings.HasPrefix(line, "GPU #0:") {
			gpu0Lines = append(gpu0Lines, line)
		}
	}
	fmt.Println("GPU 0 statistics:")
	for _, l := range gpu0Lines {

		entry := parseEntry(l) + "\n"
		fileOut.Write([]byte(entry))
	}

	fileOut.Close()

	// HR := strings.TrimLeft(gpu0Lines[0], "GPU #0: EVGA RTX 3090    - ")
	// fmt.Println("1/2 trimmed: ", HR)
	// idxEndHashRate := strings.Index(HR, " MH/s,")
	// HRF := HR[0:idxEndHashRate]
	// fmt.Println("HR: ", HRF)

	// type gpudata map[int]map[string]string
	// var re = regexp.MustCompile(`GPU #0: EVGA RTX 3090    - `)
	// //v := make(gpudata)
	// fmt.Println("regex result: ", r.FindStringSubmatch(gpu0Lines[0]))
	// r, _ := regexp.Compile(`GPU #0: EVGA RTX 3090    - \d+.\d{2} MH/s,*`)
	// fmt.Println("regex index:", r.FindStringIndex(gpu0Lines[0]))
	// matches := re.FindStringSubmatch(gpu0Lines[0])
	// fmt.Println("Matches:", matches)
	// //	v[0] = make

	//fmt.Println("Power: ", power)
	// for _, f := range testArray {
	// 	fmt.Println(f)
	// }

	fmt.Println("Exiting...")
}

func parseEntry(entry string) string {
	// fmt.Println("entry: ", entry)
	entryCleaned := strings.ReplaceAll(entry, "F: ", "F:")
	// if strings.Compare(entry, entryCleaned) != 0 {
	// 	fmt.Println("String had to be cleaned: ", entry, " --> ", entryCleaned)
	// }
	entryFields := strings.Fields(entryCleaned)
	var data []string
	data = append(data, entryFields[1][1:len(entryFields[1])-1]) // GPU
	data = append(data, entryFields[6])                          // hash rate
	temp := entryFields[8][3 : len(entryFields[8])-2]
	temps := strings.Split(temp, "/")
	data = append(data, temps[0])                                  // core temp
	data = append(data, temps[1])                                  // cbridge temp
	data = append(data, entryFields[9][2:len(entryFields[9])-2])   // power
	data = append(data, entryFields[10][2:len(entryFields[10])-2]) // fan speed
	data = append(data, entryFields[11][2:len(entryFields[11])-6]) // hash/power ratio

	return strings.Join(data, ",")
}
