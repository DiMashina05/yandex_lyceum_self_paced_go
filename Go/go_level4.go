// package main

// import(
// 	"fmt"
// )

// type Animal interface{
// 	MakeSound() string
// 	GetName() string
// 	GetInfo() string
// }

// type animal struct{
// 	name string
// 	species string
// 	age int
// 	sound string
// }

// func (a animal) MakeSound() string{
// 	return a.sound
// }

// func (a animal) GetName() string{
// 	return a.name
// }

// func (a animal) GetInfo() string{
// 	s := fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
// 	return s
// }

// func NewAnimal(name1, species11 string, age1 int, sound1 string) Animal{
// 	return &animal{name: name1, species: species11, age: age1, sound: sound1}
// }

// func ZooShow(animals []Animal){
// 	for _, v := range animals{
// 	fmt.Println(v.GetInfo())
// 	fmt.Println(v.MakeSound())
// 	}
// }

// type ZooKeeper struct{}

// func (c ZooKeeper) Feed(animal Animal){
// 	fmt.Printf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
// }