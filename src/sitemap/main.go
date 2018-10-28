package main

import (
	"fmt"
	"text/template"
	//"fmt"
	"os"
	"sitemap/post"
	"sitemap/product"
)


func main() {

	posts := post.GetPosts()

	products := product.GetProducts()

	//fmt.Println(posts)

	tpl, err := template.ParseFiles("templates/sitemap.xml")

	if (err != nil) {
		panic(err)
	}

	tplVars := map[string][]map[string]string {
		"Items": posts,
		"Products": products,
	}

	//fmt.Println(tplVars)

	// create file
	w, err := os.Create("sitemap.xml")

	defer w.Close()

	if (err != nil) {
		panic(err)
	}



	err = tpl.ExecuteTemplate(w, "main", tplVars)

	if (err != nil) {
		panic(err)
	}

	fmt.Println("Done")


}


