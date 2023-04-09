package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
const formatString = "%.6f\n"

func plusMinus(arr []int32) {
	// Write your code here
	var p, n, z int32 = 0, 0, 0
	l := len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			p++
		} else if arr[i] < 0 {
			n++
		} else {
			z++
		}
	}
	pf := float64(p) / float64(l)
	nf := float64(n) / float64(l)
	zf := float64(z) / float64(l)

	fmt.Printf(formatString, pf)
	fmt.Printf(formatString, nf)
	fmt.Printf(formatString, zf)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
