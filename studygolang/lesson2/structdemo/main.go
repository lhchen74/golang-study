package main

import (
	"fmt"
	"os"
)

type book struct {
	title  string
	author string
	price  int
}

var allBooks = make([]*book, 0, 10)

func showMenu() {
	fmt.Println("1: 添加书籍")
	fmt.Println("2: 修改书籍")
	fmt.Println("3: 显示书籍")
	fmt.Println("4: 退出")
}

func newBook(title, author string, price int) *book {
	return &book{
		title:  title,
		author: author,
		price:  price,
	}
}

func getInputBookInfo() *book {
	var (
		title  string
		author string
		price  int
	)
	fmt.Print("title:")
	fmt.Scanln(&title)
	fmt.Print("author:")
	fmt.Scanln(&author)
	fmt.Print("price:")
	fmt.Scanln(&price)

	book := newBook(title, author, price)

	return book
}

func addBook() {
	book := getInputBookInfo()

	for _, b := range allBooks {
		if b.title == book.title {
			fmt.Printf("《%s》 already exists!", b.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Println("Add book success!")

}

func showBooks() {
	if len(allBooks) == 0 {
		fmt.Println("No books, Please add book first!")
		return
	}
	for _, book := range allBooks {
		fmt.Printf("title:%s, author:%s, price:%d\n", book.title, book.author, book.price)
	}
}

func modifyBook() {
	book := getInputBookInfo()
	for i, b := range allBooks {
		if b.title == book.title {
			allBooks[i] = book
			fmt.Printf("Modify success!\n")
			return
		}
	}
	fmt.Printf("《%s》not exists!\n", book.title)
}

func main() {

	for {
		showMenu()
		var code int
		fmt.Scanln(&code)
		switch code {
		case 1:
			addBook()
		case 2:
			modifyBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Unsupport operation")
		}
	}

}
