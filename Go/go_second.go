// // package main2

// import(
// 	"time"
// 	"strings"
// 	"unicode/utf8"
// 	"errors"
// )

// var err error = errors.New("исправь свой ответ, а лучше ложись поспать")

// func currentDayOfTheWeek() string{
// 	var time time.Time = TimeNow()
	
// 	switch time.Weekday().String(){
// 	case "Monday":
// 		return "Понедельник"
// 	case "Tuesday":
// 		return "Вторник"
// 	case "Wednesday":
// 		return "Среда"
// 	case "Thursday":
// 		return "Четверг"
// 	case "Friday":
// 		return "Пятница"
// 	case "Saturday":
// 		return "Суббота"
// 	case "Sunday":
// 		return "Воскресенье"
// 	default:
// 		return "error in currentDayOfTheWeek"
// 	}
// }

// func dayOrNight() string{
// 	var time time.Time = TimeNow()
// 	if time.Hour() < 10 || time.Hour() > 22{
// 		return "Ночь"
// 	} else {
// 		return "День"
// 	}
// }

// func nextFriday() int{
// 	var time time.Time = TimeNow()
// 	switch time.Weekday().String(){
// 	case "Monday":
// 		return 4
// 	case "Tuesday":
// 		return 3
// 	case "Wednesday":
// 		return 2
// 	case "Thursday":
// 		return 1
// 	case "Friday":
// 		return 0
// 	case "Saturday":
// 		return 6
// 	case "Sunday":
// 		return 5
// 	default:
// 		return 7
// 	}
// }

// func CheckCurrentDayOfTheWeek(answer string) bool{
// 	var lowWeek string = strings.ToLower(currentDayOfTheWeek())
// 	answer = strings.ToLower(answer)
// 	if answer == lowWeek {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func CheckNowDayOrNight(answer string) (bool, error){

// if utf8.RuneCountInString(answer) != 4{
// return false, err
// }
// answer = strings.ToLower(answer)
// var timeniw string = strings.ToLower(dayOrNight())
// if answer == timeniw {
// 	return true, nil
// } else {
// 	return false, nil
// }
// }
