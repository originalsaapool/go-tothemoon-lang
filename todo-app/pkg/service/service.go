package service

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

func NewService(repos *repository.NewRepository) *Service {
	return &Service{}
}