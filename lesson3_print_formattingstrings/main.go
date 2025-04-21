//%T değişkenin tipini gösterir
package main 

import "fmt"

func main(){
	age := 35
	name := "ahmet"	
	// Print

   fmt.Print("Hello, ") //print stringin sonuna yeni line eklemez.
   fmt.Print("World!\n",)
    fmt.Println("------------------")
   // Println
   fmt.Println("Hello World")
   fmt.Println("Goodbye World")
   fmt.Println("my age is", age, "and my name is", name)
   fmt.Println("--------------------------")
   //Printf (formatted strings) %v and %_ = format specifier
   fmt.Printf("my age is %v and my name is %v\n", age, name)	
   fmt.Printf("my age is %q and my name is %q\n", age, name)	
   fmt.Printf("age is of type %T\n", age)
   fmt.Printf("you scored %f points! \n",225.5)
    fmt.Println("--------------------------------")

   // Sprintf (save formatted strings)

   var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
    fmt.Println("the saved string is ", str)
  
}