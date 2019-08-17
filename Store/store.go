package entirestore

import viewmodels "github.com/femosc2/ia-discord-bot-2/viewmodels"

var Store viewmodels.StoreViewModel

// SetStore Sets the store
func SetStore(store viewmodels.StoreViewModel) {
	Store = store
}

// GetStore Returns the store
func GetStore() viewmodels.StoreViewModel {
	return Store
}
