//pass_fail program tells whether one is passed or failed
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	   main lines for comments
	*/
	fmt.Print("Enter a grade: ") //get the grade from a user
	/*
	   make "buffer reader" to read the texts from keyboard
	   Stdin means standard input
	*/
	reader := bufio.NewReader(os.Stdin)
	/*
	   all texts before pushing the Enter keyboard
	   ReadString argument type is rune
	*/
	input, err := reader.ReadString('\n') //사용하지 않는 var err 추가
	fmt.Println(input)

	//get several returns
	//bool, err := strconv.ParseBool("true")         //string을 bool로 변환할 수 없는 경우 err return
	//file, err := os.Open("myfile.txt")             //file을 열 수 없는ㄴ 경우 err return
	//response, err := http.Get("http://golang.org") //page를 못 가져오는 경우 err return

}
