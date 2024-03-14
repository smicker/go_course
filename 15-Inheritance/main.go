package main

import "fmt"

type author struct {
	firstName string
	lastName  string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
}

func main() {
	myBlogPost := blogPost{
		"My first blogpost",
		"This is something I want to blog about",
		author{
			"Mikael",
			"Berglund",
		},
	}

	myBlogPost.details()
}
