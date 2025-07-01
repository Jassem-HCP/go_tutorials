package main

import (
	"fmt"
	"unicode/utf8"
)

type Animal struct {
	Name      string
	Age       int
	Height    float64
	ownerInfo owner
}

type owner struct {
	Name string
	Age  uint16
}

func main() {

	// fmt.Println("Hello, World!")

	var str string = "lololemon"
	//  number of bytes in the string
	// fmt.Println("String length:", len(str))
	// Count the number of runes (aka characters) in the string
	fmt.Println("Rune count:", utf8.RuneCountInString(str))
	// fmt.Println("First character:", string(str[0]))
	// fmt.Println("Last character:", string(str[len(str)-1]))

	// var char rune = 'a'
	// fmt.Println("Character as rune:", char)

	strArr := []string{"mo", "ka", "rii"}
	strArr = append(strArr, "lemo")
	for i, v := range strArr {
		// fmt.Printf("Index: %d, value : %v\n", i, v)
		fmt.Printf("memory adress of [%v] is: %p \n", v, &strArr[i])
	}

	ageMap := map[string]uint16{"Nick": 20, "Leyla": 29}
	age, isValid := ageMap["tito"]

	fmt.Printf("Nick's age is: %d\nIs valid?: %v\n", age, isValid)

	// make(int[], 0, n) this creates an empty slice of integers with a capacity of n and 0

	cat := Animal{Name: "Moka", Age: 2, Height: 0.300, ownerInfo: owner{"Leyla", 29}}
	fmt.Printf("Cat's name is: %s, age: %d, height: %.3f, Owner_Name: %s, Owner_Age:%d\n", cat.Name, cat.Age, cat.Height, cat.ownerInfo.Name, cat.ownerInfo.Age)

	// // change the cats age every year
	// // loop with range of 10 using range
	// for i := 0; i < 10; i++ {
	// 	cat.Age += 1
	// 	fmt.Printf("Cat's name is: %s, age: %d, height: %.3f\n", cat.Name, cat.Age, cat.Height)
	// }

}
