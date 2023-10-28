package models

import "fmt"

type Service struct {
	ServiceId   string
	ServiceName string
	TeamMembers []string
}

func NewService(serviceId string, service string, team []string) Service {
	return Service{ServiceId: serviceId, ServiceName: service, TeamMembers: team}
}

func (s Service) ListServices() string {
	return fmt.Sprintf("%v,%v,%v", s.ServiceId, s.ServiceName, s.TeamMembers)
}
