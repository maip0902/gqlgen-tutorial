package gqlgen_todos

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []Todo
	users []User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		Text:   input.Text,
		Title:  input.Title,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, *todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	user := &User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	r.users = append(r.users, *user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return r.todos, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	for _, user := range r.users {
		if (user.ID == id) {
			return &user, nil
		}
	}
	return &User{}, errors.New("Sorry, Not Found.")
}
func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	return r.users, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
