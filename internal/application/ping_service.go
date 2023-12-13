package application

// Sender is the output port
type Sender interface {
	Send(message string) bool
}

// Pinger is the input port
type Pinger interface {
	Ping()
}

type PingService struct {
	repository Sender
}

func NewPingService(senderRepository Sender) *PingService {
	return &PingService{repository: senderRepository}
}

func (p *PingService) Ping() {
	p.repository.Send("Ping")
}
