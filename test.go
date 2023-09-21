package main

import "fmt"

func main(){
	// var warnaLampu = "merah";

	// if warnaLampu == "merah"{
	// 	fmt.Println("Halo Merah")
	// }else{
	// 	fmt.Println("Halo lampu")
	// }

	// Switch case
	// switch warnaLampu{
	// case "merah":
	// 	fmt.Println("berhenti")
	// case "kuning":
	// 	fmt.Println("hati-hati")
	// default:
	// 	fmt.Println("maju")
	// }

	// looping for
	// for i := 1; i < 5; i++ {
	// 	fmt.Println("Angka",i)
	// }

	// looping while
	// var i = 0;
	// for i < 5{
	// 	fmt.Println(i)
	// 	i++
	// }

	// loop do while
	// var i = 0
	// for{
	// 	fmt.Println("angka",i)

	// 	i++ 
	// 	if i == 5{
	// 		break
	// 	}
	// }

	// array
	// var names[3]string
	// names[0] = "Bagaskara"
	// names[1] = "Bagaskoro"
	// names[2] = "Bagas Aja"
	var names = [...]string{
		"Bagaskara",
		"Bagaskoro",
		"Bagas Aja",
	}

	var nilai = [...]int{
		90,80,80,
	}
	
	fmt.Println(names[0], names[1], names[2])
	fmt.Println(nilai[0], nilai[1], nilai[2])
	fmt.Println("Hello")
}
