package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/tuken/poyang/graph/generated"
	"github.com/tuken/poyang/graph/model"
)

func (r *mutationResolver) CreateManager(ctx context.Context, input model.NewManager) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Managers(ctx context.Context) ([]*model.Manager, error) {
	log.Printf("queryResolver.Managers: %v", ctx)
	managers := []Manager{}
	if err := r.DB.Find(&managers).Error; err != nil {
		return nil, err
	}

	results := []*model.Manager{}
	for _, m := range managers {
		results = append(results, &model.Manager{
			ID:            strconv.Itoa(int(m.ID)),
			Email:         m.Email,
			FirstName:     m.FirstName,
			LastName:      m.LastName,
			ManagemnetsID: int(m.ManagemnetsID),
		})
	}

	return results, nil
}

func (r *queryResolver) Manager(ctx context.Context, id string) (*model.Manager, error) {
	log.Printf("queryResolver.Manager: %v, %s", ctx, id)

	manager := Manager{}
	if err := r.DB.Where("id = ?", id).Find(&manager).Error; err != nil {
		return nil, err
	}

	results := model.Manager{
		ID:            strconv.Itoa(int(manager.ID)),
		Email:         manager.Email,
		FirstName:     manager.FirstName,
		LastName:      manager.LastName,
		ManagemnetsID: int(manager.ManagemnetsID),
	}

	return &results, nil
}

func (r *queryResolver) Managements(ctx context.Context) ([]*model.Management, error) {
	log.Printf("queryResolver.Managements: %v", ctx)

	managements := []Management{}
	if err := r.DB.Find(&managements).Error; err != nil {
		return nil, err
	}

	results := []*model.Management{}
	for _, m := range managements {
		results = append(results, &model.Management{
			ID:      strconv.Itoa(int(m.ID)),
			Name:    m.Name,
			Zip:     m.Zip,
			Address: m.Address,
		})
	}

	return results, nil
}

func (r *queryResolver) Management(ctx context.Context, id string) (*model.Management, error) {
	log.Printf("queryResolver.Management: %v, %s", ctx, id)

	management := Management{}
	if err := r.DB.Where("id = ?", id).Find(&management).Error; err != nil {
		return nil, err
	}

	results := model.Management{
		ID:      strconv.Itoa(int(management.ID)),
		Name:    management.Name,
		Zip:     management.Zip,
		Address: management.Address,
	}

	return &results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
