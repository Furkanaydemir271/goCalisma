package main

import "fmt"

func main(){
	// ARRAYS = THE LENGTH OF ARRAYS CANNOT CHANGE
   // var ages [3]int = [3]int{20, 25 ,30}
    var ages = [3]int{20, 25 ,30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
     names[1]="furkan"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))       
       

	//slices (use arrays under the hood ) LENGTH OF SLICES CAN CHANGE
	var scores 	= []int{100, 50, 60}
	scores[2] = 25 
	scores = append(scores, 85) 
	fmt.Println(scores, len(scores))


	// slice ranges 
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
    fmt.Println(rangeOne,rangeTwo,rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
     

}