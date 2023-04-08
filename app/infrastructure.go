package app

import "context"

var (
	// globalCounter reflects dbReservations
	globalCounter = 2
)

type infra struct {
	queueCounter infraQueueCounter
}

func NewInfra() *infra {
	return &infra{}
}

func (infr *infra) SetQueueCounter(qc infraQueueCounter) {
	infr.queueCounter = qc
}

type infraQueueCounter interface {
	Count(ctx context.Context) (int, error)
}

type queueCounterImpl struct{}

func NewQueueCounterImpl() *queueCounterImpl {
	return &queueCounterImpl{}
}

func (qc *queueCounterImpl) Count(ctx context.Context) (int, error) {
	globalCounter++
	return globalCounter, nil
}
