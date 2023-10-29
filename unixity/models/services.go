package models

import "fmt"

// Exmple on interface for services
type ServiceInfo interface {
	GetServiceID() string
	GetServiceName() string
	GetTeamMembers() []string
}

type Service struct {
	ServiceId   string
	ServiceName string
	TeamMembers []string
}

func NewService(serviceId string, service string, team []string) Service {
	return Service{ServiceId: serviceId, ServiceName: service, TeamMembers: team}
}

// methods for the ServiceInfo interface
func (s Service) GetServiceID() string {
	return s.ServiceId
}

func (s Service) GetServiceName() string {
	return s.ServiceName
}

func (s Service) GetTeamMembers() []string {
	return s.TeamMembers
}

func (s Service) DisplayServiceInfo() {
	fmt.Println("Service ID:", s.GetServiceID())
	fmt.Println("Service Name:", s.GetServiceName())
	fmt.Println("Team Members:", s.GetTeamMembers())
}

func (s Service) ListServices() string {
	return fmt.Sprintf("%v,%v,%v", s.ServiceId, s.ServiceName, s.TeamMembers)
}
