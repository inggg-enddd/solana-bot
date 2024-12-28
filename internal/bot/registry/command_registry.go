package registry

import (
	"sync"

	"github.com/your/project/internal/bot/domain"
	"github.com/your/project/internal/bot/ports"
)

type CommandRegistry struct {
	handlers map[domain.Command]ports.MessageHandler
	mu       sync.RWMutex
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		handlers: make(map[domain.Command]ports.MessageHandler),
	}
}

func (r *CommandRegistry) Register(command domain.Command, handler ports.MessageHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[command] = handler
}

func (r *CommandRegistry) Get(command domain.Command) (ports.MessageHandler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	handler, exists := r.handlers[command]
	return handler, exists
}
