package repository

//
//import (
//	"PennieAI/config"
//	"PennieAI/models"
//	"fmt"
//)
//
//func CreateAnalyzedDocument(doc *models.AnalyzedDocument) error {
//	db := config.GetDB()
//	_, err := db.Exec("INSERT INTO analyzed_documents (user_id, title, content, summary, created_at) VALUES ($1, $2, $3, $4, $5)", doc.UserID, doc.Title, doc.Content, doc.Summary, doc.CreatedAt)
//	if err != nil {
//		fmt.Println("Error inserting analyzed document into database:", err)
//		return err
//	}
//	return nil
//}
