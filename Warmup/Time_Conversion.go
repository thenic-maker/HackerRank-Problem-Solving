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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {

	parts := strings.Split(s, ":")

	hour, _ := strconv.Atoi(parts[0])
	min, _ := strconv.Atoi(parts[1])

	sec, _ := strconv.Atoi(parts[2][:len(parts[2])-2])
	period := strings.ToUpper(parts[2][len(parts[2])-2:])

	if period == "PM" {
		if hour != 12 {
			hour += 12
			if hour == 24 {
				hour = 0
			}
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}
	time1 := format(hour) + ":" + format(min) + ":" + format(sec)
	return time1
}

func format(n int) string {
	var result string = strconv.Itoa(n)
	if n < 10 {
		result = "0" + result
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
