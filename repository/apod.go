package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"tesk-task-betera/models"
)

type Apod struct {
	pool *pgxpool.Pool
}

func NewApodRepository(pool *pgxpool.Pool) *Apod {
	return &Apod{pool: pool}
}

func (r *Apod) CreateApod(ctx context.Context, astronomyData *models.ApodDto) error {
	query := `INSERT INTO apodData ("copyright", "date", "explanation", "media_type", "service_version", "title", "url", "hdurl", "data")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.pool.Exec(ctx, query,
		astronomyData.Copyright, astronomyData.Date, astronomyData.Explanation, astronomyData.MediaType,
		astronomyData.ServiceVersion, astronomyData.Title, astronomyData.URL, astronomyData.HDURL, astronomyData.Data)

	if err != nil {
		return fmt.Errorf("failed to create apod query: %v", err)
	}

	return nil
}

func (r *Apod) GetApods(ctx context.Context, getApodsRequest *models.GetApodsRequest) (*models.GetApodsResponse, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT copyright, date, explanation, media_type, service_version, title, url, hdurl, data
		FROM apodData LIMIT $1 OFFSET $2`,
		getApodsRequest.Limit, getApodsRequest.Offset)

	if err != nil {
		return nil, fmt.Errorf("failed to get apods query: %v", err)
	}

	var GetApodsResponse models.GetApodsResponse

	for rows.Next() {
		var apod models.ApodDto

		err := rows.Scan(&apod.Copyright, &apod.Date, &apod.Explanation, &apod.MediaType,
			&apod.ServiceVersion, &apod.Title, &apod.URL, &apod.HDURL, &apod.Data)

		if err != nil {
			return nil, fmt.Errorf("failed to scan apod %w", err)
		}

		GetApodsResponse.Apods = append(GetApodsResponse.Apods, apod)
	}

	return &GetApodsResponse, nil
}

func (r *Apod) GetApodsByDate(ctx context.Context, getApodByDateRequest models.GetApodByDateRequest) (*models.GetApodByDateResponse, error) {
	var GetApodByDateResponse models.GetApodByDateResponse
	var apod models.ApodDto

	err := r.pool.QueryRow(ctx,
		`SELECT copyright, date, explanation, media_type, service_version, title, url, hdurl, data
		FROM apodData where date=$1`, getApodByDateRequest.Date).Scan(
		&apod.Copyright, &apod.Date, &apod.Explanation, &apod.MediaType,
		&apod.ServiceVersion, &apod.Title, &apod.URL, &apod.HDURL, &apod.Data)

	if err != nil {
		return nil, fmt.Errorf("failed to get apod by date query: %v", err)
	}

	GetApodByDateResponse.Apod = apod

	return &GetApodByDateResponse, nil
}
