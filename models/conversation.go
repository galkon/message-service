package models

import (
  "time"
)

type Conversation struct {
  Id        int32     `json:"id"`
  Subject   string    `json:"subject"`
  CreatedAt time.Time `json:"createdAt"`
  Messages []Message  `json:"messages"`
}