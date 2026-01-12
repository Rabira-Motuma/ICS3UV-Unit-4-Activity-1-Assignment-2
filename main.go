/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-5
 * Fileoverview: This program prints a pattern of numbers in a right-angled triangle.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	// variables
	var rowsAsString string
	var rowsAsNumber int

	reader := bufio.NewReader(os.Stdin)

	// inputs
	fmt.Printf("How many rows would you like? ")
	rowsAsString, _ = reader.ReadString('\n')
	rowsAsString = strings.TrimSpace(rowsAsString)
	rowsAsNumber, _ = strconv.Atoi(rowsAsString)

	// for statements
	for i := 1; i <= rowsAsNumber; i++ {
		var row = ""
		for j := 1; j <= i; j++ {
			row += strconv.Itoa(j) + " "
		}
		fmt.Printf("%s\n", row)
	}

	fmt.Println("\nDone.")
}
