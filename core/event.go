package core

type Event struct {
	handlers []chan bool
}

func (e *Event) Subscribe(handler chan bool) {
	if e.handlers == nil {
		e.handlers = make([]chan bool, 0)
	}

	e.handlers = append(e.handlers, handler)
}

func (e *Event) Unsubscribe(handler chan bool) {
	idxToRemove := -1
	for i, h := range e.handlers {
		if h == handler {
			idxToRemove = i
			break
		}
	}

	if idxToRemove < 0 {
		return
	}
	e.handlers[idxToRemove] = e.handlers[len(e.handlers)-1]
	e.handlers = e.handlers[:len(e.handlers)-1]
}

func (e *Event) Fire(data bool) {
	for _, handler := range e.handlers {
		go func(h chan bool) {
			h <- data
		}(handler)
	}
}
