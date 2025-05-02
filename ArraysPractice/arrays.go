package main

import (
	"fmt"
)

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.




func main(){

	hobbies := [3]string{"Drawing", "Music", "Videogames"}

	// 1
	fmt.Println("Step 1")
	fmt.Println(hobbies)
	
	// 2
	
	fmt.Println(hobbies[0])

	twoHobbies := hobbies[1:3]
	fmt.Println(twoHobbies)

	//3 
	fmt.Println("Step 3")
	firstTwoHobbies := hobbies[0:2]
	firstTwoHobbiesVariation := hobbies[:2]
	
	fmt.Println(firstTwoHobbies)
	fmt.Println(firstTwoHobbiesVariation)

	//4
	fmt.Println("Step 4")
	firstTwoHobbies = hobbies[1:3]
	fmt.Println(firstTwoHobbies)

	//5 Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learning Go", "Making Software"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	fmt.Println("Step 5")
	secondGoal := []string{}
	secondGoal = append(secondGoal, goals[1])
	fmt.Println(secondGoal)

	goals = append(goals, "Having fun")
	fmt.Println(goals)


// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

	type Product struct{
		Title string 
		Id string 
		Price float64
	}

	products := []Product{
		{Title: "Macintosh Laptop", Id : "123", Price: 1000.1},
		{Title: "Windows Laptop", Id : "1234", Price: 1500.2},
	}

	print(products)


	



	
	

}