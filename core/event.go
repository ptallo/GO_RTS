package core

// EventData is a wrapper around a generic data type
type EventData struct {
	data interface{}
}

// Event is a class for handling firing events to handlers
type Event struct {
	handlers []chan EventData
}

// Subscribe allow a handler to listen for events from this source
func (e *Event) Subscribe(handler chan EventData) {
	if e.handlers == nil {
		e.handlers = make([]chan EventData, 0)
	}

	e.handlers = append(e.handlers, handler)
}

// Unsubscribe allow a handler to no longer listen for events from this source
func (e *Event) Unsubscribe(handler chan EventData) {
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

// Fire sends the data to all subscribed handlers
func (e *Event) Fire(data EventData) {
	for _, handler := range e.handlers {
		go func(h chan EventData) {
			h <- data
		}(handler)
	}
}
