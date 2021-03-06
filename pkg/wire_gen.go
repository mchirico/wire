// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package pkg

// Injectors from init.go:

func InitializeEvent(phrase string) (Event, error) {
	message := NewMessage(phrase)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
