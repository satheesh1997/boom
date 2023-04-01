package games

type (
    Controller struct {}
    Service struct {}
)

func NewController() *Controller {
    return &Controller{}
}

func NewService() *Service {
    return &Service{}
}
