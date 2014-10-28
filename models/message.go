package models

import (
  "time"
)

type Message struct {
  Id          int32     `json:"id"`
  MessageBody string    `json:"messageBody"`
  CreatedAt   time.Time `json:"createdAt"`
  // SenderId    int32     `json:"senderId"`
  ConversationId   int32 `json:"conversationId"`
}