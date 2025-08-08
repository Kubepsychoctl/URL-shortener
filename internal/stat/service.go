package stat

import (
	"app/url-shorter/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus *event.EventBus
	StatRepo *StatRepository
}

type StatService struct {
	EventBus *event.EventBus
	StatRepo *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus: deps.EventBus,
		StatRepo: deps.StatRepo,
	}
}
func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.LinkVisitedEvent {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventlinkVisited Data: ", msg.Data)
				continue
			}
			s.StatRepo.AddClick(id)
		}
	}
}
