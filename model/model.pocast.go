package model

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	Podcast struct {
		ID            uuid.UUID `json:"id"`
		UserID        uuid.UUID `json:"user_id"`
		CategoryID    uuid.UUID `json:"category_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		ImageCoverURL string    `json:"image_cover_url"`
	}

	PodcastResponse struct {
		ID            uuid.UUID `json:"id"`
		UserID        uuid.UUID `json:"user_id"`
		CategoryID    uuid.UUID `json:"category_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		ImageCoverURL string    `json:"image_cover_url"`
	}

	AllPodcast struct {
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

func (p *Podcast) Response() PodcastResponse {
	return PodcastResponse{
		ID:            p.ID,
		UserID:        p.UserID,
		CategoryID:    p.CategoryID,
		Title:         p.Title,
		Description:   p.Description,
		ImageCoverURL: p.ImageCoverURL,
	}
}

func (p *Podcast) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "podcast" (user_id,category_id,title,description,image_cover_url) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.UserID, p.CategoryID, p.Title, p.Description, p.ImageCoverURL).Scan(&p.ID)
	if err != nil {
		return p.ID, err
	}

	return p.ID, nil
}

func (p *Podcast) All(ctx context.Context, db *sql.DB, param AllPodcast) ([]Podcast, error) {
	all := []Podcast{}

	query := `SELECT id,user_id,category_id,title,description,image_cover_url FROM "podcast" WHERE %s = $1 AND %s LIKE $2 ORDER BY %s %s OFFSET $3 LIMIT $4`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.FilterBy, param.SearchBy, param.OrderBy, param.OrderDir), param.FilterValue, "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Podcast{}
		err = rows.Scan(
			&one.ID, &one.UserID, &one.CategoryID, &one.Title, &one.Description, &one.ImageCoverURL,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (p *Podcast) One(ctx context.Context, db *sql.DB) (Podcast, error) {
	one := Podcast{}

	query := `SELECT id,user_id,category_id,title,description,image_cover_url FROM "podcast" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.ID).Scan(
		&one.ID, &one.UserID, &one.CategoryID, &one.Title, &one.Description, &one.ImageCoverURL,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (p *Podcast) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "podcast" SET user_id = $1,category_id = $2,title = $3,description = $4,image_cover_url = $5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.UserID, p.CategoryID, p.Title, p.Description, p.ImageCoverURL, p.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (p *Podcast) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "podcast" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
