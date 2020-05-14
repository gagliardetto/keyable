package keyable

import (
	"sync"

	"github.com/eiannone/keyboard"
)

type Keyable struct {
	mu      *sync.RWMutex
	runeMap map[rune]func()
	keyMap  map[keyboard.Key]func()
}

func New() *Keyable {
	return &Keyable{
		mu:      &sync.RWMutex{},
		runeMap: make(map[rune]func()),
		keyMap:  make(map[keyboard.Key]func()),
	}
}

func (kb *Keyable) OnChar(callback func(), rs ...rune) *Keyable {
	kb.mu.Lock()
	defer kb.mu.Unlock()

	for _, r := range rs {
		kb.runeMap[r] = callback
	}

	return kb
}
func (kb *Keyable) OnKey(callback func(), ks ...keyboard.Key) *Keyable {
	kb.mu.Lock()
	defer kb.mu.Unlock()

	for _, k := range ks {
		kb.keyMap[k] = callback
	}

	return kb
}

// Start starts the keyboard listener and listens to the pressed
// keys and executes the associated callback functions.
func (kb *Keyable) Start() error {
	err := keyboard.Open()
	if err != nil {
		return err
	}
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			func(r rune, k keyboard.Key) {
				kb.mu.RLock()
				defer kb.mu.RUnlock()

				{
					callback, ok := kb.runeMap[r]
					if ok && callback != nil {
						callback()
					}
				}
				{
					callback, ok := kb.keyMap[k]
					if ok && callback != nil {
						callback()
					}
				}

			}(char, key)
		}
	}()
	return nil
}

// Stop stops the keyboard
func (kb *Keyable) Stop() {
	keyboard.Close()
}
