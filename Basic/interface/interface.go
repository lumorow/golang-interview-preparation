package main

import (
	"fmt"
	"reflect"
)

type MobilePhone struct {
	modelName string
	price     int
}

type phone interface {
	call()
}

func NewMobilePhone(modelName string, price int) *MobilePhone {
	return &MobilePhone{
		modelName: modelName,
		price:     price,
	}
}

func (m *MobilePhone) call() {
	fmt.Printf("%s ringtone tuz tuz...\n", m.modelName)
}

func main() {
	mobPh1 := NewMobilePhone("samsung", 800)
	var mobPh2 *MobilePhone

	ph1 := phone(mobPh1)
	ph2 := phone(mobPh2)

	fmt.Printf("%t\n", ph1 == nil) // false
	fmt.Printf("%t\n", ph2 == nil) // false

	ph1.call()

	fmt.Println("Type ph1 is: " + reflect.TypeOf(ph1).String()) // Type ph1 is: *main.MobilePhone
	obj1, ok := ph1.(*MobilePhone)
	if ok == true {
		fmt.Printf("Success converting to MobilePhone: %#v\n", obj1)
	} else {
		fmt.Println("Fail converting to MobilePhone")
	}
	// Success converting to MobilePhone: &main.MobilePhone{modelName:"samsung", price:800}
}
