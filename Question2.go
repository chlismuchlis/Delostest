
// Online IDE - Code Editor, Compiler, Interpreter

package main

import (
	"fmt"
    "strconv"
	"strings"
    "os"
)

const inputtext = 
`5 3 1`
//`3 5 2`

func main() {
    var arrayInput = strings.Split(inputtext, " ");
    student, err := strconv.Atoi(arrayInput[0])
    candy, err := strconv.Atoi(arrayInput[1])
    first, err := strconv.Atoi(arrayInput[2])
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    urutan := first
    //flg := 1
    for i := 0; i < candy; i++ {
        fmt.Println(urutan)
        
        if urutan==student{
            urutan = 1
        } else {
            urutan++
        }
    }
}
