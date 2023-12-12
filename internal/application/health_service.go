package application

type Healther interface {
	Ping() bool
}

type Checker interface {
	Check() bool
}

type HealthService struct {
	healther Healther
}

func NewHealthService(healther Healther) *HealthService {
	return &HealthService{
		healther: healther,
	}
}

func (h *HealthService) Check() bool {
	return h.healther.Ping()
}
