package controllers

import (
	"pendulum/app/chatroom"

	"github.com/revel/revel"
)

type Chat struct {
	*revel.Controller
}

type ChatEvents struct {
	Events []chatroom.Event `json:"events"`
}

func (c Chat) Room(user string) revel.Result {
	chatroom.Join(user)
	return c.Render(user)
}

func (c Chat) Say(user, message string) revel.Result {
	chatroom.Say(user, message)
	return nil
}

func (c Chat) WaitMessages(lastReceived int) revel.Result {
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	// See if anything is new in the archive.
	var events []chatroom.Event
	for _, event := range subscription.Archive {
		if event.Timestamp > lastReceived {
			events = append(events, event)
		}
	}

	var chatEvents ChatEvents
	chatEvents.Events = events

	// If we found one, grand.
	if len(events) > 0 {

		return c.RenderJSON(chatEvents)
	}

	// Else, wait for something new.
	event := <-subscription.New
	chatEvents.Events = []chatroom.Event{event}
	return c.RenderJSON(chatEvents)
}

func (c Chat) Leave(user string) revel.Result {
	chatroom.Leave(user)
	return c.RenderText("leave")
}
