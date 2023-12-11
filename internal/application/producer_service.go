package application

type Messenger interface {
	Post(message string) bool
}

type MessagingService struct {
	repository Messenger
}

func NewMessagingService(messageRepository Messenger) *MessagingService {
	return &MessagingService{repository: messageRepository}
}

func (p *MessagingService) Post(message string) {
	p.repository.Post(message)
}
