package main

import "fmt"

func main(){
  age := 45

  fmt.Println(age<=50)
  fmt.Println(age>=50)
  fmt.Println(age==45)
  // conditionals
  if age < 30 {
	fmt.Println("age is less than 30")
  } else if age < 40 {
	fmt.Println("age is less than 40")
  }else {
	fmt.Println("age is not than 45")
  }
   
  names := []string{"mario","luigi","yoshi","peach","bowser"}

  for index, value := range names {
	if index == 1 {
		fmt.Println("continuing at pos ", index)
		continue 
	}
	if index > 2 {
		fmt.Println("breaking at position", index)
		break
	}

	fmt.Printf("the value at pos %v is %v\n", index,value)
  }
 
}