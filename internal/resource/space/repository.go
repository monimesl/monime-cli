package space

import (
	"context"
	"errors"
	"fmt"
	"github.com/monimesl/monime-cli/cli-utils/monimeapis"
	"github.com/monimesl/monime-cli/internal/store"
	"net/http"
)

type Repository interface {
	ListSpaces(ctx context.Context) (List, error)
	GetActiveSpace(_ context.Context) (Space, bool, error)
	ActivateSpace(Space) error
}

var (
	_ Repository = &defaultRepository{}
)

const (
	activeSpaceField = "active_space"
)

type defaultRepository struct{}

func (r *defaultRepository) ActivateSpace(space Space) error {
	return store.Get().SetConfig(activeSpaceField, space)
}

func (r *defaultRepository) GetActiveSpace(_ context.Context) (Space, bool, error) {
	sp := Space{}
	err := store.Get().GetConfig(activeSpaceField, &sp)
	if errors.Is(err, store.ErrKeyNotFound) {
		return Space{}, false, nil
	} else if err != nil {
		return Space{}, false, err
	}
	return sp, true, nil
}

func (r *defaultRepository) ListSpaces(ctx context.Context) (List, error) {
	type SpaceDto struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Alias string `json:"alias"`
	}
	result, err := monimeapis.ApiRequest[any, []SpaceDto](ctx, nil,
		http.MethodGet, "/spaces/", nil)
	if err != nil {
		return List{}, err
	}
	list := List{}
	for _, spc := range result.Result {
		list.Add(Space{
			Id:    spc.Id,
			Name:  spc.Name,
			Alias: spc.Alias,
			URL:   fmt.Sprintf("https://%s.monime.space", spc.Alias),
		})
	}
	return list, nil
}
