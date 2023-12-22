package application

import "strings"

// Receiver is the output port
type Receiver interface {
	Receive() []string
}

// Ponger is the input port
type Ponger interface {
	Pong() []string
}

type PongService struct {
	repository Receiver
}

func NewPongService(receiverRepository Receiver) *PongService {
	return &PongService{repository: receiverRepository}
}

func (p *PongService) Pong() (pongs []string) {
	for _, ping := range p.repository.Receive() {
		pong := strings.Replace(ping, "i", "o", 1)
		pongs = append(pongs, pong)
	}

	return pongs
}
