package me


type (
    // Controller is the controller for the me package
    Controller struct {}

    // Service is the service for the me package
    Service struct {}
)

// NewController returns a new instance of the me controller
func NewController() *Controller {
    return &Controller{}
}

// NewService returns a new instance of the me service
func NewService() *Service {
    return &Service{}
}
