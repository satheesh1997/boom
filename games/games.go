package games


type (
    // Controller is the controller for the games package
    Controller struct {}

    // Service is the service for the games package
    Service struct {}
)

// NewController returns a new instance of the games controller
func NewController() *Controller {
    return &Controller{}
}

// NewService returns a new instance of the games service
func NewService() *Service {
    return &Service{}
}
