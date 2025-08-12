package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/SurekshyaS/distributed/internal/db"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func GetUsersHandler(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, name FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, u)
    }
    c.JSON(http.StatusOK, users)
}