package base

import "fmt"
import "time"

func Base() {

	fmt.Println("======1.base=======")
	//1.define variable
	var name string
	name = "ni hao"

	address := "beijing"

	fmt.Println(name, address)

	//2.slice:An array of unrestricted lengths

	fruits := []string{"pear", "orange"}
	fruits[0] = "apple"

	fmt.Println(fruits)

	//3.map

	fruitsMap := map[string]int{"pear": 15, "apple": 12}
	fmt.Println(fruitsMap)

	//4. if switch

	var date string = "20170228"

	if date == "20170228" {
		switch address {
		case "beijing":
			fmt.Println("study go")
		default:
			fmt.Println("no record")
		}
	} else {
		fmt.Println("no record")
	}

	//5. for

	for index := 0; index < len(fruits); index++ {
		fmt.Println(fruits[index])
	}

	for k, v := range fruits {
		fmt.Println(k, "+++++", v)
	}

	for k, v := range fruitsMap {
		fmt.Println(k, "~~~~~~", v)
	}

	//6.time
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("20060102"))
	fmt.Println(now.Add(time.Hour * 24).Format("2006-01-02"))
}
