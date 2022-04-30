package main

import "testing"

func TestPubsub(t *testing.T) {
	ps := NewPubsub()
	ch := ps.Subscribe("topic")
	ps.Publish("topic", "hello")
	msg := <-ch
	if msg != "hello" {
		t.Errorf("expected hello, got %q", msg)
	}
}

func TestPubsubClose(t *testing.T) {
	ps := NewPubsub()
	ch := ps.Subscribe("topic")
	ps.Publish("topic", "hello")
	msg := <-ch
	if msg != "hello" {
		t.Errorf("expected hello, got %q", msg)
	}
	ps.Close()
	_, ok := <-ch
	if ok {
		t.Errorf("expected closed channel")
	}
}

func TestConcurrentPubSub(t *testing.T) {
	ps := NewPubsub()
	ch := ps.Subscribe("topic")
	// Publish from a goroutine
	go func() {
		for i := 0; i < 100; i++ {
			ps.Publish("topic", "hello")
		}
	}()
	for i := 0; i < 100; i++ {
		msg := <-ch
		if msg != "hello" {
			t.Errorf("expected hello, got %q", msg)
		}
	}
}
