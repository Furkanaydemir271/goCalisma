package  main

import "fmt"

func main(){
 /*    x := 0
	for x < 5  {
		fmt.Println("value of x is :", x)
		x++
	}
 
 for i := 0;  i < 5; i++ {
	fmt.Println("value of x is :", i)

 } */

   names := []string{"mario","luigi","yoshi","bowser"} 
/*   for i := 0; i<len(names); i++ {
	fmt.Println(names[i])
  } 
  */
  for index, value := range names {
	fmt.Printf("The value at index %v is %v\n",index,value)

  }

} 