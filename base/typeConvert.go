package base

import (
	"fmt"
	"strconv"
)

func TypeConvert() {
	/**type:
			  bool          1
			  byte          1
			  int,uint      4
		      int8,uint8    1 //-128 ~ 127, 0 ~255
		      int16,uint16  2 //0 ~65535
	          int32,uint32  4 //-21 Billion
	          float32       4
	          float64       4


	          array
	          struct
	          string

	          slice
	          map
	          interface

	*/

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
