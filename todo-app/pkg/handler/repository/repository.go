package repository

type Authorization interface{
	
}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository() *Repository {
	return &Repository{}
}