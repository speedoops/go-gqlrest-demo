package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/speedoops/go-gqlrest-demo/graph/errorsx"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/speedoops/go-gqlrest-demo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	if r.todos == nil {
		todo.ID = "T9527"
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	list := r.todos
	for _, l := range list {
		if l.ID == input.ID {
			l.Text = input.Text
			l.UserID = input.UserID
			return l, nil
		}
	}
	return nil, errorsx.NewNotFoundError(fmt.Errorf("not found(%s)", input.ID))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	n := 0
	list := r.todos
	for _, l := range list {
		if l.ID != id {
			list[n] = l
			n++
		}
	}
	if len(r.todos) == n {
		return false, errorsx.NewNotFoundError(fmt.Errorf("not found(%s)", id))
	}
	r.todos = list[:n]
	return true, nil
}

func (r *mutationResolver) DeleteTodoByUser(ctx context.Context, userID string) (bool, error) {
	n := 0
	list := r.todos
	for _, l := range list {
		if l.UserID != userID {
			list[n] = l
			n++
		}
	}
	if len(r.todos) == n {
		return false, errorsx.NewNotFoundError(fmt.Errorf("not found(%s)", userID))
	}
	r.todos = list[:n]
	return true, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string, name *string, tmp *int) (*model.Todo, error) {
	for _, l := range r.todos {
		if l.ID == id {
			return l, nil
		}
	}
	return nil, errorsx.NewNotFoundError(errors.New("not found"))
}

func (r *queryResolver) Todos(ctx context.Context, ids []string, userID *string, userID2 string, text *string, text2 string, done *bool, done2 bool, pageOffset *int, pageSize *int) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID, Role: "test"}, nil
}

func (r *todoResolver) Category(ctx context.Context, obj *model.Todo) ([]*model.Category, error) {
	category := model.Category{ID: "1", Name: "Category"}
	return []*model.Category{&category}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
