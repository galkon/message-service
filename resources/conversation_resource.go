package resources

import (
  "github.com/jinzhu/gorm"
  "github.com/galkon/messages/models"
  "github.com/gin-gonic/gin"
  "strconv"
)

type ConversationResource struct {
  Db gorm.DB
}

func (cr *ConversationResource) GetConversation(c *gin.Context) {
    var conversation models.Conversation

    id := parseStrToInt(c.Params.ByName("id"))

    if cr.Db.First(&conversation, id).Related(&conversation.Messages).RecordNotFound() {
        c.JSON(404, gin.H{"error": "not found"})
    } else {
        c.JSON(200, conversation)
    }
}

func (cr *ConversationResource) GetAllConversations(c *gin.Context) {
  var conversations []models.Conversation

  perPageAttr := c.Request.URL.Query()["per_page"]
  pageAttr := c.Request.URL.Query()["page"]

  perPage := 20
  page := 1

  if perPageAttr != nil {
    perPage = parseStrToInt(perPageAttr[0])
  }

  if pageAttr != nil {
    page = parseStrToInt(pageAttr[0])
  }

  offset := strconv.Itoa((page - 1) * perPage)
  limit := strconv.Itoa(perPage)

  cr.Db.Limit(limit).Offset(offset).Find(&conversations)

  for i := range conversations {
    cr.Db.Model(&conversations[i]).Related(&conversations[i].Messages)
  }

  c.JSON(200, conversations)
}

func parseStrToInt(str string) (int) {
  idInt, _ := strconv.Atoi(str)
  id := int(idInt)
  return id
}