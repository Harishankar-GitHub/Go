package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	// var fruitList = []string{}
	//	var sliceName = [lengthNeedNotBeSpecified]datatype{}
	//	{} -> Initialization

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// We can add as many values to the Slice.
	// No need to specify the length at the begnining.
	fmt.Printf("Type of Fruit List is %T\n", fruitList)

	// Adding values to Slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit List after adding Mango and Banana:\n", fruitList)

	fruitList = append(fruitList[1:])
	// : is used to slice up the slice.
	// Now the fruitList will have elements from 1st index till the end.
	fmt.Println("Fruit List from [1:]\n", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Fruit List [1:3]\n", fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println("Fruit List [:3]\n", fruitList)

	// Defining a Slice using make() keyword.
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777
	// This won't work because the default allocation of the memory is 4.

	fmt.Println("High Scores", highScores)

	highScores = append(highScores, 555, 666, 321)
	fmt.Println("High Scores", highScores)
	// This works because it is reallcoate the memory.

	fmt.Println("Is scores sorted ?\n", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// There are other methods and functionalities available in sort.
	fmt.Println("Sorted scores", highScores)
	fmt.Println("Is scores sorted ?\n", sort.IntsAreSorted(highScores))

	// Removing a value from Slices based on Index.

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
