package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2{
		fmt.Print("ERROR!!! : WRONG ARGUMENTS")
	}
	arr := ReadArray("E:\\unsorted.txt")
	arr = quicksort(arr)
	WriteToFile(arr,"E:\\Sorted.txt")
}

func ReadArray(fileName string) [] int {
	_, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ERROR !!! : WRONG FILE")
		os.Exit(1)
		return nil
	}
	file, err := os.Open(fileName)
	if err != nil {//failed to open the file
		log.Fatal(err)
	}
	defer file.Close()//after the end of scope
	scanner := bufio.NewScanner(file)//scanner that read the file
	var A [] int // Array of integers
	var S []string // The read text from the file in string form
	var Value int // The converted string
	for scanner.Scan(){
		S = strings.Fields(scanner.Text())//divide the string
		Value, _ = strconv.Atoi(S[0])//convert string to integer
		A = append(A,Value )//store the number in the array
	}
	return A //the read array is returned
}


func quicksort(A [] int) [] int{
	if len(A) < 2 {//if array size is 1 or array is empty
		return A
	}

	left, right := 0, len(A)-1//left and right assginment

	pivot := rand.Int() % len(A) // random pivot generator to try to speed up

	A[pivot], A[right] = A[right], A[pivot] //swap the pivot with the most right element

	for i, _ := range A { // partitioning the array
		if A[i] < A[right] {
			A[left], A[i] = A[i], A[left]
			left++
		}
	}

	A[left], A[right] = A[right], A[left]

	quicksort(A[:left])//sort the left array
	quicksort(A[left+1:])//sort the right array

	return A
}

func WriteToFile(A [] int,FName string){
	f, err := os.Create(FName)//Create the output file
	if err != nil {//if failed to create
		panic(err)
	}
	defer f.Close()//end of the scope
	length := len(A)//store the size of the array

	str := ""
	for i:= 0 ;i < length ;i ++ {//loop on the array and store it in the string
		str += strconv.Itoa(A[i])//Convert the integer value to a string
		str += " "
	}


	_,_ = f.WriteString(str) // write the file

}