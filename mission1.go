package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func testCases(size int) []int {
    // Return list of test cases in integer
    //   from user space-separated input.
    ynScanner := bufio.NewScanner(os.Stdin)
    ynScanner.Scan()
    list_yn := strings.Fields(ynScanner.Text())
    new_yn := make([]int, size, len(list_yn))
    count := 0

str_to_int:
    if (size > 0 && size <= 100) && count < len(list_yn) {
        num, _ := strconv.Atoi(list_yn[count])
        if num >= -100 && num <= 100 {
            new_yn[count] = int(num)
        }
        count += 1
        goto str_to_int
    }
    return new_yn
}

func sumTestCases(yn []int) int {
    // Sum all positive integers in yn list.
    index := 0
    sum := 0

add_pos:
    if index < len(yn) {
        if yn[index] > 0 {
            sum += yn[index] * yn[index]
        }
        index += 1
        goto add_pos
    }
    return sum
}

func output(list []int) {
    // prints output in a new line.
    index := 0
print_item:
    if index < len(list) {
        fmt.Println(list[index])
        index += 1
        goto print_item
    }
}

func main() {
    // Main function.

    // Total number of test cases to be follow.
    nScanner := bufio.NewScanner(os.Stdin)
    nScanner.Scan()
    n, _ := strconv.Atoi(nScanner.Text()) // convert to int.

    // List to store all the sum of each test case.
    list_of_sums := make([]int, n, n+1)

test_case:
    if n > 0 && n <= 100 {
        xScanner := bufio.NewScanner(os.Stdin)
        xScanner.Scan()
        x, _ := strconv.Atoi(xScanner.Text())

        // Getting testcases with testCases() function.
        // And then...
        // Add all positive integers with sumTestCases() function.
        list_of_sums[len(list_of_sums)-n] = sumTestCases(testCases(x))
        n -= 1
        goto test_case
    }

    // print output in console.
    output(list_of_sums)
}
