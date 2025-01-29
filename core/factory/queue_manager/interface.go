package queue_manager

import (
	"game-go/core/queue_manager"
)

type Factory interface {
	QueueManager() queue_manager.QueueManager
}
