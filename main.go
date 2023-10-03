package main

import (
"fmt"
"os"
"log"
)
func main(){
	fmt.Println("Hello");

	portString := os.Getenv("PORT")
	 
	if portString == ""{
	log.Fatal("PORT is not found")
	}

	fmt.Println("PORT: ", portString);
 } 