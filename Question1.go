
// Online IDE - Code Editor, Compiler, Interpreter

package main

import (
	"fmt"
	"strings"
	"time"
)

const inputtext = 
`7 6 2022
19 8 2022`
/*
const inputtext = 
`15 8 2022
22 8 2022`
*/
func main() {
    var date1 = strings.Split(inputtext, "\n")[0];
    var day1 = twochar(strings.Split(date1, " ")[0]);
    var month1 = twochar(strings.Split(date1, " ")[1]);
    var year1 = twochar(strings.Split(date1, " ")[2]);
    var date2 = strings.Split(inputtext, "\n")[1];
    var day2 = twochar(strings.Split(date2, " ")[0]);
    var month2 = twochar(strings.Split(date2, " ")[1]);
    var year2 = twochar(strings.Split(date2, " ")[2]);
    
    t1, err := time.Parse("2006-01-02 15:04", year1+"-"+month1+"-"+day1+" 04:35")
    if err != nil { fmt.Println(err) }
    t2, err := time.Parse("2006-01-02 15:04", year2+"-"+month2+"-"+day2+" 04:35")
    if err != nil { fmt.Println(err) }
    
    year, month, day := diff(t1, t2)
	
	final := 0
	if day>0 {final = day*15}
	if month>0 {final = month*500}
	if year>0 {final = year*12000}
	
    fmt.Printf("%d.", final)
}

func twochar(isi string) string {
   
    if len(isi) == 1 {
        isi = "0"+isi 
    }
   
   return isi 
}

func diff(a, b time.Time) (year, month, day int) {
	//if a.Location() != b.Location() {
	//	b = b.In(a.Location())
	//}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	// Normalize negative values
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}