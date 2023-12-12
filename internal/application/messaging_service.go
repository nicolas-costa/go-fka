package application

type Messenger interface {
	Send(message string) bool
}

type Pinger interface {
	Ping()
}

type MessagingService struct {
	repository Messenger
}

func NewMessagingService(messageRepository Messenger) *MessagingService {
	return &MessagingService{repository: messageRepository}
}

func (p *MessagingService) Ping() {
	p.repository.Send("Ping")
}
