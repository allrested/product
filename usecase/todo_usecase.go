package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/allrested/product/entity"
	"github.com/allrested/product/repository/pgsql"
	"github.com/allrested/product/repository/redis"
	"github.com/allrested/product/transport/request"
	"github.com/allrested/product/utils"
)

// TodoUsecase represent the todos usecase contract
type TodoUsecase interface {
	Create(ctx context.Context, request *request.CreateTodoReq) error
	GetByID(ctx context.Context, id int64) (entity.Todo, error)
	Fetch(ctx context.Context) ([]entity.Todo, error)
	Update(ctx context.Context, id int64, request *request.UpdateTodoReq) error
	Delete(ctx context.Context, id int64) error
}

type todoUsecase struct {
	todoRepo   pgsql.TodoRepository
	redisRepo  redis.RedisRepository
	ctxTimeout time.Duration
}

// NewTodoUsecase will create new an todoUsecase object representation of TodoUsecase interface
func NewTodoUsecase(todoRepo pgsql.TodoRepository, redisRepo redis.RedisRepository, ctxTimeout time.Duration) TodoUsecase {
	return &todoUsecase{
		todoRepo:   todoRepo,
		redisRepo:  redisRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (u *todoUsecase) Create(c context.Context, request *request.CreateTodoReq) (err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.todoRepo.Create(ctx, &entity.Todo{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return
}

func (u *todoUsecase) GetByID(c context.Context, id int64) (todo entity.Todo, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	todo, err = u.todoRepo.GetByID(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		err = utils.NewNotFoundError("todo not found")
		return
	}
	return
}

func (u *todoUsecase) Fetch(c context.Context) (todos []entity.Todo, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	todosCached, _ := u.redisRepo.Get("todos")
	if err = json.Unmarshal([]byte(todosCached), &todos); err == nil {
		return
	}

	todos, err = u.todoRepo.Fetch(ctx)
	if err != nil {
		return
	}

	todosString, _ := json.Marshal(&todos)
	u.redisRepo.Set("todos", todosString, 30*time.Second)
	return
}

func (u *todoUsecase) Update(c context.Context, id int64, request *request.UpdateTodoReq) (err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	todo, err := u.todoRepo.GetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = utils.NewNotFoundError("todo not found")
			return
		}
		return
	}

	todo.Name = request.Name
	todo.UpdatedAt = time.Now()

	err = u.todoRepo.Update(ctx, &todo)
	return
}

func (u *todoUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	_, err = u.todoRepo.GetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = utils.NewNotFoundError("todo not found")
			return
		}
		return
	}

	err = u.todoRepo.Delete(ctx, id)
	return
}
