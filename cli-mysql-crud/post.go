package main

type Post struct {
	Id             int
	Title, Content string
}

func GetAllPosts() []Post {
	db := dbConnection()
	query, err := db.Query("SELECT * FROM posts")

	if err != nil {
		panic(err.Error())
	}

	p := Post{}
	posts := []Post{}

	for query.Next() {
		var id int
		var title, content string

		query.Scan(&id, &title, &content)
		p = Post{Id: id, Title: title, Content: content}
		posts = append(posts, p)
	}

	defer db.Close()

	return posts
}

func CreateNewPost(post *Post) {
	db := dbConnection()
	query, err := db.Prepare("INSERT INTO posts(title, content) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(post.Title, post.Content)

	defer db.Close()
}

func FindPostById(id int) Post {
	db := dbConnection()
	query, err := db.Query("SELECT * FROM posts WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	var p Post

	for query.Next() {
		var id int
		var title, content string

		query.Scan(&id, &title, &content)

		p = Post{Id: id, Title: title, Content: content}
	}

	defer db.Close()

	return p
}

func UpdatePost(post *Post) {
	db := dbConnection()
	query, err := db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(post.Title, post.Content, post.Id)

	defer db.Close()
}

func DeletePost(id int) {
	db := dbConnection()
	_, err := db.Query("DELETE FROM posts WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
