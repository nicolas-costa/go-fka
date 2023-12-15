package application

// Receiver is the output port
type Receiver interface {
	Receive() []string
}

type Ponger interface {
	Pong() []string
}

type PongService struct {
	repository Receiver
}

func NewPongService(receiverRepository Receiver) *PongService {
	return &PongService{repository: receiverRepository}
}

func (p *PongService) Pong() []string {
	return p.repository.Receive()
}
