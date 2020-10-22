package model

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	Music struct {
		ID            uuid.UUID `json:"id"`
		CategoryID    uuid.UUID `json:"category_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		ImageCoverURL string    `json:"image_cover_url"`
		FilePath      string    `json:"file_path"`
	}

	MusicResponse struct {
		ID            uuid.UUID `json:"id"`
		CategoryID    uuid.UUID `json:"category_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		ImageCoverURL string    `json:"image_cover_url"`
		FilePath      string    `json:"-"`
	}

	AllMusic struct {
		FilterBy    string `json:"filter_by"`
		FilterValue string `json:"filter_value"`
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (m *Music) Response() MusicResponse {
	return MusicResponse{
		ID:            m.ID,
		CategoryID:    m.CategoryID,
		Title:         m.Title,
		Description:   m.Description,
		ImageCoverURL: m.ImageCoverURL,
		FilePath:      m.FilePath,
	}
}

func (m *Music) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "music" (category_id,title,description,image_cover_url,file_path) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), m.CategoryID, m.Title, m.Description, m.ImageCoverURL, m.FilePath).Scan(&m.ID)
	if err != nil {
		return m.ID, err
	}

	return m.ID, nil
}

func (m *Music) All(ctx context.Context, db *sql.DB, param AllMusic) ([]Music, error) {
	all := []Music{}

	query := `SELECT id,category_id,title,description,image_cover_url,file_path FROM "music" WHERE %s = $1 AND %s LIKE $2 ORDER BY %s %s OFFSET $3 LIMIT $4`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.FilterBy, param.SearchBy, param.OrderBy, param.OrderDir), param.FilterValue, "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Music{}
		err = rows.Scan(
			&one.ID, &one.CategoryID, &one.Title, &one.Description, &one.ImageCoverURL, &one.FilePath,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (m *Music) One(ctx context.Context, db *sql.DB) (Music, error) {
	one := Music{}

	query := `SELECT id,category_id,title,description,image_cover_url,file_path FROM "music" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), m.ID).Scan(
		&one.ID, &one.CategoryID, &one.Title, &one.Description, &one.ImageCoverURL, &one.FilePath,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (m *Music) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "music" SET category_id = $1,title = $2,description = $3,image_cover_url = $4 ,file_path = $5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), m.CategoryID, m.Title, m.Description, m.ImageCoverURL, m.FilePath, m.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *Music) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "music" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), m.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
