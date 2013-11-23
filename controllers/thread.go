package controllers

import (
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func Thread(w http.ResponseWriter, r *http.Request) {
    board_id_str := mux.Vars(r)["board_id"]
    board_id, _ := strconv.Atoi(board_id_str)
    err, board := models.GetBoard(board_id)

    post_id_str := mux.Vars(r)["post_id"]
    post_id, _ := strconv.Atoi(post_id_str)
    err, op, posts := models.GetThread(post_id)

    if err != nil {
        http.NotFound(w, r)
        return
    }

    utils.RenderTemplate(w, r, "thread.html", map[string]interface{} {
        "board": board,
        "op": op,
        "posts": posts,
    })
}