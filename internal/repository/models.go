// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Paste struct {
	ID          int64              `json:"id"`
	FileID      string             `json:"file_id"`
	FileContent []byte             `json:"file_content"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}
