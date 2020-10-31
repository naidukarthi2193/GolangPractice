package main

import (
	"fmt"
)

//Employee ..
type Employee interface {
	GiveAllDetails() []string
	GiveID() []string
}

//AppDev ...
type AppDev struct {
	Name        string
	Address     string
	ID          string
	Environment string
}

//WebDev ...
type WebDev struct {
	Name     string
	Address  string
	ID       string
	FrontEnd string
	BackEnd  string
}

//GiveAllDetails for Appdev ...
func (a *AppDev) GiveAllDetails() []string {
	var alldetails []string
	alldetails = append(alldetails, a.Name, a.Address, a.ID, a.Environment)
	return alldetails
}

//GiveAllDetails for WebDev ...
func (w *WebDev) GiveAllDetails() []string {
	var alldetails []string
	alldetails = append(alldetails, w.Name, w.Address, w.ID, w.FrontEnd, w.BackEnd)
	return alldetails
}

//GiveID for AppDev...
func (a *AppDev) GiveID() string {
	return a.ID
}

//GiveID for WebDev...
func (w *WebDev) GiveID() string {
	return w.ID
}

func main() {

	var appguy AppDev
	appguy.Name = "Karthikraj"
	appguy.Address = "THane"
	appguy.ID = "12"
	appguy.Environment = "Kotlin"
	fmt.Println(appguy.GiveAllDetails())
	fmt.Println(appguy.GiveID())

	var webguy WebDev
	webguy.ID = "69"
	webguy.Name = "Thiru"
	webguy.BackEnd = "Nodejs"
	webguy.Address = "Mulund"
	webguy.FrontEnd = "HTML"
	fmt.Println(webguy.GiveAllDetails())
	fmt.Println(webguy.GiveID())

}
