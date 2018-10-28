package main

import "fmt"

func main() {


	//var dia Dialog
	//dia := new(Dialog)
	//dia := Dialog{"Name", 1, 2}
	dia := Dialog{name: "N", positionX: 3, positionY: 5}

	//dia.name = "Name"
	//dia.positionX = 1
	//dia.positionY = 2
	//dia.updateDialogPosition()

	//size := dia.getSquare()

	var w Window

	str := getNames(&w, &dia)
	fmt.Println(str)

	//var m Multi

	//append(&m.items, w)
	//append(&m.items, dia)
	//
	//fmt.Println(m)



}

type Browser interface {
	render() string
}

type Dialog struct {

	Browser
	name string
	positionX int
	positionY int
}

func (d *Dialog) getSquare() int {

	return d.positionX * d.positionY
}
func (Dialog) render() string {

	return "Dialog"
}


type Window struct {
	Browser
	body string
}

func (Window) render() string {
	return "Window"
}



func getNames(items ... Browser) string {

	var str string

	for _, v := range items {

		str += " " + v.render() + "\n"
	}
	return str
}

type Multi struct {
	items []Browser
}

func (m *Multi) render() string {

	str := ""

	for _,i := range m.items {

		str += i.render() + "\n"
	}

	return str
}
