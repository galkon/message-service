package main

import (
  "log"
  "io/ioutil"
  "fmt"
  "os"
  "encoding/json"

  "github.com/galkon/messages/resources"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

type Config struct {
  DatabaseUser string `json:"databaseUser"`
  DatabaseName string `json:"databaseName"`
}

var config Config

func main() {
  r := gin.Default()

  LoadConfig()

  conversationResource := &resources.ConversationResource{Db: GetDb()}

  r.GET("/conversations/:id", conversationResource.GetConversation)
  r.GET("/conversations", conversationResource.GetAllConversations)

  r.Run(":8080")
}

func GetDb() (gorm.DB) {

  db, err := gorm.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.DatabaseUser, config.DatabaseName))
  
  if err != nil {
    log.Fatal(err)
  }

  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

  db.LogMode(true)

  return db
}

func LoadConfig() {
  config_file, e := ioutil.ReadFile("./config/config.json")
  
  if e != nil {
    fmt.Printf("File error: %v\n", e)
    os.Exit(1)
  }

  json.Unmarshal(config_file, &config)
}