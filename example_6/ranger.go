package ranger

import "fmt"

func Runner() {

	nums := []int{2, 3, 4}
	sum := 0
    numsMap := make(map[int]int)
	for order , num := range nums {
		sum += num
        numsMap[order]= num 
	}
	fmt.Println("sum:", sum)
	fmt.Println("nums map:", numsMap)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i,".th variable is " ,c)
	}
}