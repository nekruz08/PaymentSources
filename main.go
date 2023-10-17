// package main

// import "fmt"

// func main() {
// 	var operations[3]int64
// 	operations[0]=9
// 	operations[1]=-10
// 	operations[2]=10
// 	fmt.Println(operations)

// }

// package main

// import "fmt"

// func main() {
// 	operations:=[3]int64{}
// 	fmt.Println(operations)
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[3]int64{-10,9,10}
// 	fmt.Println(operations)
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{9,-10,10}
// 	fmt.Println(operations)
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{10,-10,9}
// 	for i := 0; i < len(operations); i++ {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{10,-10,9}
// 	for i := 0; i < len(operations); i++ {
// 		fmt.Println(operations[i])
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{10,-9,10}
// 	for _, v := range operations {
// 		fmt.Println(v)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{10,9,-10}
// 	for _,index := range operations {
// 		fmt.Println(index)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{9,-10,10}
// 	sum:=int64(0)
// 	for _, v := range operations {
// 		sum+=v
// 	}

// 	fmt.Println(sum)
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{-100000,1,100}

// 	max:=operations[0]

// 	for _, v := range operations {
// 		if max<=v{
// 			max=v
// 		}
// 	}
// 	fmt.Println(max)
// }

// package main

// import "fmt"

// func main() {
// 	operations:=[...]int64{9,-10,10}
// 	sum:=sum(operations)
// 	max:=max(operations)

// 	fmt.Println(sum)
// 	fmt.Println(max)
// }

// func sum(sl [3]int64) int64 {
// 	sum:=int64(0)
// 	for _, v := range sl {
// 		sum+=v
// 	}
// 	return int64(sum)
// }

// func max(operations [3]int64) int64 {
// 	max:=operations[0]
// 	for _, operation := range operations {
// 		if max<=operation{
// 			max=operation
// 		}
// 	}
// 	return max
// }

// package main

// import "fmt"

// func main() {
// 	var operations []int64
// 	fmt.Println(operations[0])
// }

// package main

// import "fmt"

// func main() {
// operations:=make([]int64,4)
// fmt.Println(operations[0])

// operations:=[]int64{9,-10,10}
// fmt.Println(operations[0])

// 	var operations []int64
// 	operations=append(operations,1)
// 	operations=append(operations,2)
// 	fmt.Println(operations[1])

// }

// package main

// import "fmt"

// func main() {
// 	var operations []int64
// 	operations = append(operations, 9)
// 	operations=append(operations, -10)
// 	operations=append(operations, 10)

// 	sum:=sum(operations)
// 	max:=max(operations)

// 	fmt.Println(sum)
// 	fmt.Println(max)
// }

// func sum(s []int64) int64 {
// 	sum:=int64(0)
// 	for _, v := range s {
// 		sum+=v
// 	}
// 	return sum
// }

// func max(s []int64) int64 {
// 	max:=int64(0)
// 	for _, v := range s {
// 		if max<=v{
// 			max=v
// 		}
// 	}
// 	return max
// }

package main

import "fmt"

func main() {
	type person struct{
		name string
		age int
	}

	p:=[]person{
		{
			name: "A",
			age: 1,
		},
		{
			name: "B",
			age: 2,
		},
		{
			name: "c",
			age: 3,
		},
	}
	fmt.Printf("%+v",p)

}