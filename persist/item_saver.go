package persist

import "log"

// ItemSaver to save items
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itmeCounter := 0
		for {
			item := <-out
			log.Printf("Got item #%d: %v", itmeCounter, item)
			itmeCounter++
		}
	}()
	return out
}
