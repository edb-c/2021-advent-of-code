package main

/* Count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.)

Read file,
for i range of lines
if i=0 not greater
check if value is greater
if i+1 > i
*/
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main1() {
	//lastNumber := 0  <- this is incorrect 
	lastNumber := math.MaxInt64
	increasedCount := 0

    // Open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // Close the file at the end of the program
    defer f.Close()

    // Read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		currentNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		fmt.Println("lastNumber is ", lastNumber, "currentNumber is ", currentNumber)

		if currentNumber > lastNumber {
			increasedCount++
			fmt.Println("increasedCount")
		}
		lastNumber = currentNumber
    	
		if err := scanner.Err(); err != nil {
        log.Fatal(err)
    	}
	} //end for
	fmt.Println("increasedCount is ", increasedCount)
}//end main
