package main

import (
	"net/http"
	"strconv"

	"github.com/jordanmatinwebdev/rss_blog_agregator_go/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
