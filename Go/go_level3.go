// // package main

// import (
// 	"fmt"
// 	"strings"
// 	// "unicode/utf8"
// 	"sort"
// )

// func ThisTvar(str string)bool{
// 	switch str{
// 	case " ":
// 		return true
// 	case "!":
// 		return true
// 	case ".":
// 		return true
// 	case "?":
// 		return true
// 	case ",":
// 		return true
// 	default:
// 		return false
// 	}
// }

// func Preobraz(str []string) map[string]int{
// 	mapa := make(map[string]int) 
// 	for _, val := range str{
// 		mapa[strings.ToLower(val)]++
// 	}
// 	return mapa
// }

// func commonWord(mapa map[string]int) (int, map[int][]string) {
// 	MaxValue := 0
// 	mapa2 := make(map[int][]string)
// 	for Key, Value := range mapa {
// 		if Value > MaxValue {
// 			MaxValue = Value
// 		}
// 		mapa2[Value] = append(mapa2[Value], Key)
// 	}
// 	return MaxValue, mapa2
// }



// func getTopWords(wordMap map[string]int, n int) []string{
// 	var outstr []string
// 	mapa2 := make(map[int]string)
// 	var array []int
// 	for	Key, Value := range wordMap{
// 		if _, ok:= mapa2[Value]; ok{
// 			continue
// 		}
// 		mapa2[Value] = Key
// 		array = append(array, Value)
// 	}
// 	sort.Ints(array)

// 	for i := 0; i < n; i++{
// 		outstr = append(outstr, mapa2[array[len(mapa2) - 1 - i]])
// 	}
// 	return outstr
// }

// func getTopWords2(wordMap map[string]int, n int) ([]string, []int) {
// 	countToWords := make(map[int][]string)
// 	var counts []int

// 	// Собираем мапу: частота -> слова
// 	for word, count := range wordMap {
// 		countToWords[count] = append(countToWords[count], word)
// 	}

// 	// Собираем все уникальные частоты
// 	for count := range countToWords {
// 		counts = append(counts, count)
// 	}

// 	// Сортируем по убыванию
// 	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

// 	var resultWords []string
// 	var resultCounts []int

// 	for _, count := range counts {
// 		words := countToWords[count]
// 		sort.Strings(words) // для стабильности

// 		for _, word := range words {
// 			if len(resultWords) < n {
// 				resultWords = append(resultWords, word)
// 				resultCounts = append(resultCounts, count)
// 			}
// 		}
// 	}

// 	return resultWords, resultCounts
// }


// func AnalyzeText(text string){
// text1 := []rune(text)
// var word string = ""
// var newText []string
// for i := 0; i < utf8.RuneCountInString(text); i++{
// 	if ThisTvar(string(text1[i])){
// 		if word != ""{
// 			newText = append(newText, word)
// 			word = ""
// 		}else{
// 			continue
// 		}
// 	} else {
// 		word += string(text1[i])
// 	}
// }
// if word != ""{
// 	newText = append(newText, word)
// }

// mapa := Preobraz(newText)
// MaxValue, mapaInv := commonWord(mapa)

// fmt.Println("Количество слов:", len(newText))
// fmt.Println("Количество уникальных слов:", len(mapa))
// fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", mapaInv[MaxValue], MaxValue)

// outstr, array := getTopWords2(mapa, 5)
// fmt.Println("Топ-5 самых часто встречающихся слов:")
// for i := 0; i < 5; i++ {
// 	fmt.Printf("\"%s\": %d раз\n", outstr[i], array[i])
// }

// fmt.Println(getTopWords(mapa, 5))

// }
