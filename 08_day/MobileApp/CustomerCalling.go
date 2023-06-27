package main

import "fmt"

//create the interface
type Calling interface {
	SetRingtune()
	ChangeLanguage()
	TalkWithCustomerCare()
}

//create the struct
type User struct {
	Name string
}

// Implementing interface
func (u User) SetRingtune() {
	fmt.Println(u.Name, "Press 1 for Chnage the ring tune")
}
func (u User) ChangeLanguage() {
	fmt.Println("Press 2 for Chnage the Language")
}
func (u User) TalkWithCustomerCare() {
	fmt.Println("Press 5 for Chnage the ring tune")
}
func main() {
	fmt.Println("Welcome to Jio Care Center")
	user := User{Name: "LovePreet"}
	user.SetRingtune()
	user.TalkWithCustomerCare()
	user.ChangeLanguage()

}
