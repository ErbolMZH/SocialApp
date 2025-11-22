package db

import (
	"Tiago/internal/store"
	"context"
	"fmt"
	"log"
	"math/rand"
)

var usernames = []string{
	"Erbol", "Airo", "Arni", "Mukha",
}
var titles = []string{
	"Understanding RESTful APIs with Go",
	"Building a Blog Backend Using PostgreSQL",
	"Optimizing SQL Queries for Better Performance",
	"Implementing JWT Authentication in Go",
}
var contents = []string{
	"This article explains how to design and build RESTful APIs using Go.",
	"Learn how to create a complete blog backend connected to PostgreSQL.",
	"A detailed guide on optimizing SQL queries for faster performance.",
	"Step-by-step tutorial on implementing JWT authentication in Go applications.",
}
var tags = []string{
	"go backend development",
	"postgresql database",
	"sql optimization",
	"jwt authentication",
}
var comments = []string{
	"Great explanation, helped me a lot!",
	"Thanks for sharing this tutorial.",
	"Very clear and easy to understand.",
	"Looking forward to more content like this!",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	users := generateUsers(4)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}
	posts := generatePosts(4, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}

	}
	comments := generateComments(4, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}

	}
	return
	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}
	return users

}
func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				titles[rand.Intn(len(titles))],
				titles[rand.Intn(len(titles))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
