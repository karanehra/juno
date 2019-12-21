package models

import (
	"context"
	"fmt"
	"juno/database"
	"time"
)

//GetArticles fetches articles from db
func GetArticles() {
	coll := database.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := coll.Find(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
}
