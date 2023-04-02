package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
)

type CityUseCase struct {
	repo CityRepo
}

type CityRepo interface {
	ReadAll(ctx context.Context) (*proto.Cities, error)
	Insert(ctx context.Context, t *proto.City) error
	Update(ctx context.Context, t *proto.City) error
}

func NewCityUseCase(r CityRepo) *CityUseCase {
	return &CityUseCase{r}
}

func (r *CityUseCase) Save(ctx context.Context, t *proto.City) (*proto.City, error) {
	len := len(t.GetId())
	if len == 0 {
		t.Id = getHex()
		if err := r.repo.Insert(ctx, t); err != nil {
			return nil, err
		}
		return t, nil
	} else {
		if err := r.repo.Update(ctx, t); err != nil {
			return nil, err
		}
		return t, nil
	}
}

func (r *CityUseCase) ReadAll(ctx context.Context) (*proto.Cities, error) {
	return r.repo.ReadAll(ctx)
}
func getHex() string {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		fmt.Println("error")
	}
	var buf [8]byte
	hex.Encode(buf[:], b[:])
	return string(buf[:])
}
