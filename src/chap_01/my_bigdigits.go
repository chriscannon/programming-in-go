package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var helpFlag bool
var barFlag bool

const (
	usage = "Usage: %s [-b|--bar] <whole-number>\n" +
		"-b --bar draw an overbar and underbar\n"
)

func init() {
	flag.BoolVar(&helpFlag, "help", false, usage)
	flag.BoolVar(&helpFlag, "h", false, usage)
	flag.BoolVar(&barFlag, "bar", false, usage)
	flag.BoolVar(&barFlag, "b", false, usage)
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 || helpFlag || (barFlag && len(os.Args) == 1) {
		fmt.Printf(usage, filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var stringOfDigits string
	if barFlag {
		stringOfDigits = os.Args[2]
	} else {
		stringOfDigits = os.Args[1]
	}

	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		if barFlag && row == 0 {
			drawBar(len(line))
		}
		fmt.Println(line)
		if barFlag && row+1 == len(bigDigits[0]) {
			drawBar(len(line))
		}
	}
}

func drawBar(length int) {
	fmt.Println(strings.Repeat("*", length))
}

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
