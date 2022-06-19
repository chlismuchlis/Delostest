// Online IDE - Code Editor, Compiler, Interpreter

package main

import (
	"fmt"
    "strconv"
	"strings"
    "os"
)

const inputtext = 
`5
1 5 7 2 4` 
/*
`4
1 3 4 9`
`4
1 3 5 4`*/ 

func main() {
    var arrayInput = strings.Split(inputtext, "\n");
    lengths, err := strconv.Atoi(arrayInput[0])
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    kolomtengah := 0
    if lengths%2 != 0 {
        kolomtengah = ((lengths - 1)/2) + 1
    } else {
        kolomtengah = (lengths/2) + 1
    }
    //fmt.Println(kolomtengah)
    var derets = strings.Split(arrayInput[1], " ");
    kiri := 0;kanan := 0 ;
    status := "kiri"
    for i := 0; i < lengths; i++ {
        //fmt.Println(derets[i]+ganjilgenap+" ")
        angka, err := strconv.Atoi(derets[i])
        if err != nil {
            // handle error
            fmt.Println(err)
            os.Exit(2)
        }
        if(i != (kolomtengah-1)){
            if(status == "kiri"){
                kiri = kiri + angka
            } else {
                kanan = kanan + angka
            }
        }
        if(i == (kolomtengah-1)){
            status = "kanan"
        }
    }
    /*
        fmt.Println(" ")
        fmt.Println(kiri)
        fmt.Println(kanan)*/
    
    if kiri == kanan {
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}