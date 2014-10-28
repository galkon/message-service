package main

import (
  "log"
  "github.com/galkon/messages/resources"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func main() {
  r := gin.Default()

  conversationResource := &resources.ConversationResource{Db: GetDb()}

  r.GET("/conversations/:id", conversationResource.GetConversation)
  r.GET("/conversations", conversationResource.GetAllConversations)

  r.Run(":8080")
}

func GetDb() (gorm.DB) {
  db, err := gorm.Open("postgres", "user=treach3r dbname=message_service sslmode=disable")
  
  if err != nil {
    log.Fatal(err)
  }

  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

  db.LogMode(true)

  return db
}