package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/flrnd/goching"
)

type lines []string

func getLines(l []string) lines {
	return lines(l)
}

func (l lines) contains(e string) bool {
	for _, element := range l {
		if element == e {
			return true
		}
	}
	return false
}

func printHexagram(l lines) {
	white := color.New(color.FgWhite).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yin := "=== ==="
	yang := "======="

	oYinLine := fmt.Sprintf("%s   %s", red(yin), white(yang))
	oYangLine := fmt.Sprintf("%s   %s", red(yang), white(yin))

	var yinLine string
	var yangLine string

	if l.contains("OYin") || l.contains("OYang") {
		yinLine = fmt.Sprintf("%s   %s", white(yin), white(yin))
		yangLine = fmt.Sprintf("%s   %s", white(yang), white(yang))
	} else {
		yinLine = fmt.Sprintf("%v", white(yin))
		yangLine = fmt.Sprintf("%v", white(yang))
	}

	for i := len(l) - 1; i >= 0; i-- {
		switch l[i] {
		case "OYin":
			fmt.Println(oYinLine)
		case "OYang":
			fmt.Println(oYangLine)
		case "Yin":
			fmt.Println(yinLine)
		case "Yang":
			fmt.Println(yangLine)
		}
	}
}

func getMovingLinesString(m []int) string {
	var b strings.Builder
	for _, l := range m {
		fmt.Fprintf(&b, "%d.", l+1)
	}
	return b.String()
}

func printReadingString(r goching.Reading) {
	var b strings.Builder

	fmt.Fprintf(&b, "Hexagram %d", r.Hexagram.Number)
	if len(r.MovingLines) > 0 {
		s := getMovingLinesString(r.MovingLines)
		s = s[:len(s)-1]
		fmt.Fprintf(&b, " with line(s) %v", s)
		fmt.Fprintf(&b, " to Hexagram %d\n", r.RelatingHex.Number)
	}

	fmt.Println(b.String())
}

func main() {
	myReading := goching.CastReading(goching.NewCast)
	fmt.Println()
	printHexagram(getLines(myReading.Lines))
	fmt.Println()
	printReadingString(myReading)
}
