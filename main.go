package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func Scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer func() {
		if err := writer.Flush(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	//var a,b int
	//Scanf("%d %d",&a,&b)
	//Printf("%d %d",a,b)10
	var s string
	StrScanf(&s)
	data := Split(s)
	var num, tmp float64
	num = float64(data[0])
	tmp = num
	for i := 0; i < data[1]-1; i++ {
		sqrt := math.Sqrt(num)
		tmp = tmp + sqrt
		num = sqrt
	}
	Printf("%.2f", tmp)
}

// Use " " split input numbers
// Return []int
func Split(s string) []int {
	l := strings.Split(s, " ")
	count, _ := strconv.Atoi(s)
	repo := make([]int, count)
	for _, s := range l {
		v, _ := strconv.Atoi(s)
		repo = append(repo, v)
	}
	return repo
}

// Use bufio get input strings
// input *string
func StrScanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
