package main

import "fmt"

func main() {

	lst := make(map[string]int)

	lst["One"] = 1
	lst["2"] = 2

	delete(lst, "2")

	two, fail := lst["One"]

	fmt.Println(two, fail)


	var arr [5]int
	arr[2] = 15

	var total int

	for _, value := range arr {

		total += value
	}

	fmt.Println(total)
	fmt.Println(arr)
	fmt.Println(float64(15))

	/*

	elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
	}
	elements := map[string]map[string]string{
            "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
	x := make([]float64, 5, 10)

	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]

	x := [5]float64{
    98,
    93,
    77,
    82,
    83,
}
	 */
	//for i:=1; i < 10; i++ {

	/*
	const X = "X_CONSTANT"

	var (
		a = 1
		b = 2
		c = 3
	)

	var ss string = "Hya"

	var xx int
	xx = 44

	xx += a + b + c

	yy := 55.0


	var input float64

	fmt.Scanf("%f", &input)

	fmt.Println("Output")
	fmt.Println(input)
	fmt.Println(ss)
	fmt.Println(xx)
	fmt.Println(yy)
	*/

	//http.HandleFunc("/", index)

	//http.Handle("/assets", http.StripPrefix("assets", http.FileServer(http.Dir("./assets"))))

	//http.ListenAndServe(":3000", nil)
}

/*
func index(writer http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("templates/index.html", "templates/header.html")

	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	tpl.ExecuteTemplate(writer, "index", nil)

	//writer.Header().Set("Content-type", "text/plain")
	//writer.Write([]byte("Hey igor"))
}
*/

func calsS(arr []float64) float64 {

	return 0.00
}