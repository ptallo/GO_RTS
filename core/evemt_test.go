package core_test

import (
	"fmt"
	"go_rts/core"
	"testing"
	"time"
)

func Test_GivenEvent_CanSubscribeToEvent(t *testing.T) {
	event := core.Event{}
	eventListener := make(chan bool)

	event.Subscribe(eventListener)
}

func Test_GivenEvent_CanFire(t *testing.T) {
	event := core.Event{}

	event.Fire(true)
}

func Test_GivenEventWithSubscribers_WhenEventFires_ThenSubcribersRecieveEvent(t *testing.T) {
	event := core.Event{}
	eventListener := make(chan bool)
	event.Subscribe(eventListener)

	event.Fire(true)

	data := <-eventListener

	if data == false {
		t.Errorf("data not recieved from main thread")
	}
}

func Test_GivenEventWithManySubscribers_WhenEventFires_ThenAllSubscribersGetEvent(t *testing.T) {
	event := core.Event{}
	el1 := make(chan bool)
	event.Subscribe(el1)
	el2 := make(chan bool)
	event.Subscribe(el2)

	event.Fire(true)

	select {
	case _ = <-el1:
		fmt.Println("Success")
	case <-time.After(1 * time.Second):
		t.Error("Data was not recieved from channel 1")
	}

	select {
	case _ = <-el2:
		fmt.Println("Success")
	case <-time.After(1 * time.Second):
		t.Error("Data was not recieved from channel 2")
	}
}

func Test_GivenEventWithSubscribers_WhenSubscribersAreRemoved_ThenSubscribersDoNotGetEvent(t *testing.T) {
	event := core.Event{}
	eventListener := make(chan bool)
	event.Subscribe(eventListener)

	event.Unsubscribe(eventListener)
	event.Fire(true)

	select {
	case _ = <-eventListener:
		t.Error("Data was recieved from the channel")
	case <-time.After(1 * time.Second):
		fmt.Println("Success")
	}
}
