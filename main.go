package main

import "log"

func main(){
	f()
}

func f() (result int) {
	defer func(){
		// good, it prints
		// 2021/05/06 07:05:34 result: 1
		log.Println("result:", result)
	}()
	return 1
}
