package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	fmt.Println("Welcome to the CLI CRUD with MYSQL")

	for {
		fmt.Println("========== OPTIONS ==========")

		fmt.Println("1 _ CREATE POST")
		fmt.Println("2 _ READ POSTS")
		fmt.Println("3 _ UPDATE POST")
		fmt.Println("4 _ DELETE POST")
		fmt.Println("9 _ EXIT")

		fmt.Println("========== COMAND ==========")

		var input int
		fmt.Scanln(&input)

		fmt.Println("========== RESULT ==========")

		switch input {
		case 1:
			create()
		case 2:
			read()
		case 3:
			update()
		case 4:
			delete()
		case 9:
			fmt.Println("bye bye...")
			os.Exit(0)
		default:
			fmt.Println("an error has occurred...")
			os.Exit(-1)
		}
	}

}

func read() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "ID\tTITLE\tCONTENT")

	posts := GetAllPosts()

	for _, post := range posts {
		fmt.Fprintln(w, strconv.Itoa(post.Id)+"\t"+post.Title+"\t"+post.Content)
		w.Flush()
	}
}

func create() {
	fmt.Print("Enter the title of the new post: ")
	title := ""
	fmt.Scanln(&title)

	fmt.Print("Enter the content of the new post: ")
	content := ""
	fmt.Scanln(&content)

	post := Post{Title: title, Content: content}

	CreateNewPost(&post)

	fmt.Println("Post created")
}

func update() {
	fmt.Print("Enter the id of the post to be edited: ")
	id := 0
	fmt.Scan(&id)

	fmt.Print("Enter the new title of the post: ")
	title := ""
	fmt.Scanln(&title)

	fmt.Print("Enter the new content of the post: ")
	content := ""
	fmt.Scanln(&content)

	post := Post{Id: id, Title: title, Content: content}

	UpdatePost(&post)

	fmt.Println("Post updated")
}

func delete() {
	fmt.Print("Enter the id of the post to be deleted: ")
	id := 0
	fmt.Scan(&id)

	DeletePost(id)

	fmt.Println("Post deleted")
}
