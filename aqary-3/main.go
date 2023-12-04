package main

import "fmt"

func main() {

	input := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	output := swap(input)
	fmt.Println(input)
	fmt.Println(output)

	fmt.Println("_____________________________")

	input = map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
		6: "f",
	}

	output = swap(input)
	fmt.Println(input)
	fmt.Println(output)

}

func swap(input map[int]string) map[int]string {

	length := len(input)
	maxLength := length
	output := make(map[int]string, length)

	if length%2 == 0 {
		maxLength = length
	} else {
		maxLength = length - 1
		output[length] = input[length]
	}

	for counter := 1; counter < maxLength; counter += 2 {
		output[counter] = input[counter+1]
		output[counter+1] = input[counter]
	}

	return output

}
