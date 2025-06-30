// // package main1

// import(
// 	"fmt"
	
// )

// func mkain() {
// 	var numberOfOccupiedSeats int = 0
// 	var numberOfAvailableSeats int = 5
// 	var word string = ""
// 	var whoPerson string = ""
// 	var ferstPerson string = "-"
// 	var secondPerson string = "-"
// 	var threePerson string = "-"
// 	var fhourPerson string = "-"
// 	var fivePerson string = "-"

// 	for{
// 		fmt.Scan(&word)
// 		switch word{
// 		case "количество":
// 			fmt.Println("Осталось свободных мест:", numberOfAvailableSeats, "\nВсего человек в очереди:",  numberOfOccupiedSeats)
// 		case "очередь":
// 			fmt.Printf("1. %s \n2. %s \n3. %s \n4. %s \n5. %s \n", ferstPerson, secondPerson, threePerson, fhourPerson, fivePerson)
// 		case "конец":
// 			fmt.Printf("1. %s \n2. %s \n3. %s \n4. %s \n5. %s \n", ferstPerson, secondPerson, threePerson, fhourPerson, fivePerson)
// 			return
// 		default:
// 			whoPerson = word
// 			var num int
// 			fmt.Scan(&num)

// 			if int(num) < 1 || int(num) > 5{
// 				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
// 				break
// 			}
			
// 			if numberOfOccupiedSeats == 5{
// 				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
// 				break
// 			}
// 			switch num{
// 			case 1:
// 				if ferstPerson == "-"{
// 				ferstPerson = whoPerson
// 				numberOfOccupiedSeats++
// 				numberOfAvailableSeats--
// 				} else{
// 					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
// 				}
// 			case 2:
// 				if secondPerson == "-"{
// 					secondPerson = whoPerson
// 					numberOfOccupiedSeats++
// 					numberOfAvailableSeats--
// 				}else{
// 					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
// 				}
// 			case 3:
// 				if threePerson == "-"{
// 					threePerson = whoPerson
// 					numberOfOccupiedSeats++
// 					numberOfAvailableSeats--
// 				}else{
// 					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
// 				}
// 			case 4:
// 				if fhourPerson == "-"{
// 					fhourPerson = whoPerson
// 					numberOfOccupiedSeats++
// 					numberOfAvailableSeats--
// 				}else{
// 					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
// 				}
// 			case 5:
// 				if fivePerson == "-"{
// 					fivePerson = whoPerson
// 					numberOfOccupiedSeats++
// 					numberOfAvailableSeats--
// 				}else{
// 					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
// 				}
				
// 			}
// 		}
// 	}
// }