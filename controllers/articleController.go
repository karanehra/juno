package controllers

import (
	"context"
	"juno/database"
	"juno/util"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetArticles fetches articles from db
func GetArticles(res http.ResponseWriter, req *http.Request) {
	size, _ := strconv.Atoi(req.FormValue("size"))
	page, _ := strconv.Atoi(req.FormValue("page"))
	query := req.FormValue("query")
	results, err := database.GetPaginatedArticles(int32(size), int64(page), query)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessReponse(res, results)
}

//PurgeArticles empties the article entries
func PurgeArticles(res http.ResponseWriter, req *http.Request) {
	coll := database.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll.DeleteMany(ctx, bson.D{})
	util.SendSuccessReponse(res, map[string]string{})
}
