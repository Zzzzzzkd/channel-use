package trylock

import (
	"testing"
	"time"
)

func TestLock(T *testing.T) {
	m := NewMutex()
	var a = 0
	m.Lock()
	go func() {
		m.Lock()
		a = 1
	}()
	if a != 0 {
		T.Fatalf("Mutex can be repeat locked")
	}
}

func TestUnlock(T *testing.T) {
	m := NewMutex()
	var a = 0
	m.Lock()
	go func() {
		m.Lock()
		a = 1
	}()
	m.Unlock()
	time.Sleep(1 * time.Second)
	if a != 1 {
		T.Fatalf("Unlock failed")
	}
	m.Unlock()
	defer func() {
		perr := recover()
		if perr == nil {
			T.Fatalf("unlock of unlocked mutex should panic")
		}
	}()
	m.Unlock()
}

func TestTrylock(T *testing.T) {
	m := NewMutex()
	m.Lock()
	go func() {
		time.Sleep(time.Second * 1)
		m.Unlock()
	}()
	if !m.Trylock(3 * time.Second) {
		T.Fatalf("trylock failed")
	}
}

func TestIslocked(T *testing.T) {
	m := NewMutex()
	if m.Islocked() {
		T.Fatalf("Islocked should be false")
	}
	m.Lock()

	if !m.Islocked() {
		T.Fatalf("Islocked should be true")
	}
}
