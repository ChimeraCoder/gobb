package controllers

import (
    "net/http"
    "sirjtaa/utils"
    "sirjtaa/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")
        confirm  := r.FormValue("password2")

        var error string
        if(password != confirm) {
            error = "Passwords don't match"
        }

        if error != "" {
            utils.RenderTemplate(w, r, "register.mustache", map[string]interface{} {
                "error": error,
            })
            return
        }

        // We're good, let's make it
        db_map := models.GetDbSession()
        user := models.NewUser(username, password)
        db_map.Insert(user)
    }

    utils.RenderTemplate(w, r, "register.mustache", nil)
}