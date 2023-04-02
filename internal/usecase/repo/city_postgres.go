package repo

import (
	"context"
	"fmt"

	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
	"github.com/audryus/boleto-palm-tree/pkg/postgres"
)

const _defaultEntityCap = 64

type CityRepo struct {
	*postgres.Postgres
}

func NewCityRepo(pg *postgres.Postgres) *CityRepo {
	return &CityRepo{pg}
}

func (r *CityRepo) Insert(ctx context.Context, t *proto.City) error {
	sql, args, err := r.Builder.
		Insert("city").
		Columns("id, name").
		Values(t.GetId(), t.GetName()).
		ToSql()

	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}
	return nil
}

func (r *CityRepo) Update(ctx context.Context, t *proto.City) error {
	sql, args, err := r.Builder.
		Update("city").Set("name", t.GetName()).
		Where("id = ?", t.GetId()).
		ToSql()
	fmt.Println(sql)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}
	return nil
}

func (r *CityRepo) ReadAll(ctx context.Context) (*proto.Cities, error) {
	var result = &proto.Cities{Cities: make([]*proto.City, 0)}

	sql, _, err := r.Builder.
		Select("id, name").
		From("city").
		ToSql()
	if err != nil {
		return result, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return result, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]*proto.City, 0, _defaultEntityCap)

	for rows.Next() {
		e := &proto.City{}

		err = rows.Scan(&e.Id, &e.Name)
		if err != nil {
			return result, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}
		entities = append(entities, e)
	}

	result.Cities = entities

	return result, nil
}
