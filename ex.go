package main

func main(){
	// a := 15

	// if a % 3 == 0{
	// 	print("Fizz")
	// }
	// if a % 5 == 0{
	// 	println("Buzz")
	// }

	// month := "May"

	// switch month{
	// case "December", "January", "February":
	// 	println("Winter")
	// case "March", "April", "May":
	// 	println("Spring")
	// case "June", "July", "August":
	// 	println("Summer")
	// case "September", "October", "November":
	// 	println("Fall")
	// default:
	// 	println("5chi fasl qachon kashf etildi")
	// }

	// sumOfOdds := 0
	// sumOfEvens := 0

	// for i := 1; i <= 100; i++{
	// 	if i % 2 == 0{
	// 		sumOfEvens += i
	// 	}else{
	// 		sumOfOdds += i
	// 	}
	// }
	// println(sumOfEvens - sumOfOdds)

	isPrime := false
	n := 1
	for i := 2; i < n/2; i++{
		if n % i == 0{
			isPrime = true
			break
		}
	}
	if isPrime || n == 1{
		println("Tub emas")
	}else{
		println("Tub son")
	}

}