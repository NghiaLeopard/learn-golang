package slice

import "fmt"

type Product struct {
    title string
    id int
    price int
}

type floatMap map[string]float64

func main() {
    // exercise 1
	hobbies := [3]string{"learn","study","gym"}
    fmt.Println(hobbies)


    // exercise 2
 
    fmt.Println(hobbies[0])

    newArr := hobbies[1:]

    newArr[0] = "studying"

    fmt.Println(newArr,hobbies)

    // exercise 3
    // cap have three, capacity will contain 

//     hobbies:     [ learn | study | gym ]
//                      ^       ^       ^
//                      |       |       |
//     createSlice: [ learn | study ] 
    createSlice := hobbies[:2]

    fmt.Println(createSlice)

    // // exercise 4
    fmt.Println(cap(createSlice))

    createSlice = createSlice[1:3]
    fmt.Println(createSlice)

    

    // exercise 7
    dynamicArr := []Product{Product{
        title: "Nghia",
        id: 1,
        price: 1000000000,
    },Product{
        title: "Beo",
        id: 2,
        price: 1000000000,
    }}

    newDynamicArr := append(dynamicArr,Product{title: "Định",
    id: 3,
    price: 1000000000,})

    fmt.Println(newDynamicArr)

    
}

