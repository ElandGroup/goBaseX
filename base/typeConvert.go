package base

import (
	"encoding/json"
	"fmt"
	. "gobasex/dto"
	"strconv"
)

func ConvertBase() {
	/**type:
			  bool          1
			  byte          1
			  int,uint      4
		      int8,uint8    1 //-128 ~ 127, 0 ~255
		      int16,uint16  2 //0 ~65535
	          int32,uint32  4 //-21 Billion
	          float32       4
	          float64       4




	*/

	fmt.Println("======2.Convert Base=======")
	//1. convert to string
	fmt.Println(strconv.FormatBool(true))
	var dollar int64 = 20
	fmt.Println(strconv.FormatInt(dollar, 10))
	var bonus int32 = 11
	fmt.Println(strconv.FormatInt(int64(bonus), 10))
	var coat float32 = 100.2
	fmt.Println(strconv.FormatFloat(float64(coat), 'f', -1, 32))

	//2. convert from string
	if skirt, err := strconv.ParseInt("12", 10, 32); err != nil {

	} else {
		//var skirt64 int64 = skirt
		var skirt32 int32 = int32(skirt)
		fmt.Printf("%T,%v", skirt32, skirt32)
	}

	skirt, _ := strconv.ParseFloat("12.3", 32)
	fmt.Printf("\n%T,%v\n", skirt, skirt)

	//3. convert int,float
	var computer float32 = 500
	fmt.Println(int32(computer))

}

func ConvertCollection() {

	/*
	   array
	   struct
	   string

	   slice
	   map
	   interface
	*/

	fmt.Println("======3.Convert Collection=======")
	//1.slice,map to json

	//slice
	fruits := []string{"apple", "orange", "pear"}
	fruitByte, _ := json.Marshal(fruits)
	fmt.Println(string(fruitByte))

	//map
	fruitsMap := map[string]string{"apple": "red", "orange": "yellow", "pear": "green"}
	fruitByteMap, _ := json.Marshal(fruitsMap)
	fmt.Println(string(fruitByteMap))

	//2.json to slice,map
	jsonSlice := `["apple","orange","pear"]`
	jsonMap := `{"apple":"red","orange":"yellow","pear":"green"}`
	slice := []string{}
	json.Unmarshal([]byte(jsonSlice), &slice)

	mapString := make(map[string]string, 0)
	json.Unmarshal([]byte(jsonMap), &mapString)
	fmt.Println(mapString)

}

func ConvertFromObject() {
	/*
	   struct
	*/

	//1.object
	fmt.Println("======4.Convert Object to Collection=======")
	fruit := Fruit{Name: "apple", Color: "red"}
	fruitByte, _ := json.Marshal(&fruit)
	fmt.Println(string(fruitByte))

	//2.slice object
	fruits := []Fruit{Fruit{Name: "apple", Color: "red"}, Fruit{Name: "pear", Color: "white"}}
	fruitsByte, _ := json.Marshal(&fruits)
	fmt.Println(string(fruitsByte))

	//3.map object

	fruitsMap := map[string]Fruit{"fruit1": Fruit{Name: "apple", Color: "red"}, "fruit2": Fruit{Name: "pear", Color: "white"}}
	fruitsByteMap, _ := json.Marshal(&fruitsMap)
	fmt.Println(string(fruitsByteMap))

}

func ConvertToObject() {
	fmt.Println("======5.Convert  Collection to Object=======")
	//1.json to object
	fruitSlice := `{"name":"apple","color":"red","price":0}`
	fruit := Fruit{}
	if err := json.Unmarshal([]byte(fruitSlice), &fruit); err != nil {

	} else {
		fmt.Printf("\n%#v\n", fruit)
	}
	//2.json to slice object
	fruitsSlice := `[{"name":"apple","color":"red","price":0},{"name":"pear","color":"white","price":0}]`
	fruits := []Fruit{}
	if err := json.Unmarshal([]byte(fruitsSlice), &fruits); err != nil {

	} else {
		fmt.Printf("\n%#v\n", fruits)
	}
	//3.json to map
	fruitStringMap := `{"fruit1":{"name":"apple","color":"red","price":0},"fruit2":{"name":"pear","color":"white","price":0}}`
	fruitMap := make(map[string]Fruit, 0)
	if err := json.Unmarshal([]byte(fruitStringMap), &fruitMap); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fruitMap)
	}

	//4. json(struct ) to map
	fruitMap2 := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(fruitSlice), &fruitMap2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("struct -> map", fruitMap2)
	}

	//5.json(struct slice) to map
	fruitMap3 := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(fruitsSlice), &fruitMap3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("[]struct -> map", fruitMap3)
	}

}
