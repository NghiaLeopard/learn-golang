package main

import "fmt"

func main() {
	website := map[string]string{"Google": "https://google.com", "aws": "https://aws.com"}

	// fmt.Println(website)

    // fmt.Println(website["Google"])

    // website["linkedIn"] = "https://linkedIn.com"

	// fmt.Println(website)

    // delete(website,"Google")

	// fmt.Println(website)


    websites := make([]string,2,5)

    fmt.Print(websites,len(websites),cap(websites))

    // mapUrl := make(map[string]string,5)

    for key,value := range website {
        fmt.Println("Key: ",key)
        fmt.Println("Value: ",value)
    }
}
