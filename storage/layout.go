package storage

import (
	"sort"
)

// StorageLayout represents the layout of storage with multiple slots.
// It holds a slice of SlotDescriptor pointers, each representing a storage slot.
type StorageLayout struct {
	Slots []*SlotDescriptor `json:"slots"` // Slots is a slice of SlotDescriptor pointers.
}

// GetSlots returns a slice of pointers to SlotDescriptor representing all slots.
func (s *StorageLayout) GetSlots() []*SlotDescriptor {
	return s.Slots
}

// GetSlot retrieves a SlotDescriptor based on a given slot ID.
// Returns nil if no slot with the given ID exists.
func (s *StorageLayout) GetSlot(slot int64) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Slot == slot {
			return slotInfo
		}
	}
	return nil
}

// GetSlotByName searches for and returns a SlotDescriptor based on a slot name.
// Returns nil if no slot with the given name is found.
func (s *StorageLayout) GetSlotByName(name string) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Name == name {
			return slotInfo
		}
	}
	return nil
}

// GetSlotByOffset finds and returns a SlotDescriptor based on a slot offset.
// Returns nil if no slot with the given offset is found.
func (s *StorageLayout) GetSlotByOffset(offset int64) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Offset == offset {
			return slotInfo
		}
	}
	return nil
}

// GetSlotByType finds and returns a SlotDescriptor based on a slot type.
// Returns nil if no slot of the given type is found.
func (s *StorageLayout) GetSlotByType(t string) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Type == t {
			return slotInfo
		}
	}
	return nil
}

// AppendSlot adds a new SlotDescriptor to the Slots slice at a specified index.
// The function returns true if the slot is successfully added. It returns false
// if the slot already exists. After adding, the slots are sorted based on their
// declaration order using the DeclarationId field.
func (s *StorageLayout) AppendSlot(index int64, slot *SlotDescriptor) bool {
	if !s.SlotExists(slot.Slot) {
		s.Slots[index] = slot

		// Ensure that the slot is in the correct position
		// It's basically a trick to order them by their declaration order
		sort.Slice(s.Slots, func(i, j int) bool {
			return s.Slots[i].DeclarationId < s.Slots[j].DeclarationId
		})

		return true
	}
	return false
}

// SlotExists checks if a slot with the given ID already exists in the layout.
// Returns true if the slot exists, and false otherwise.
func (s *StorageLayout) SlotExists(slot int64) bool {
	for _, slotInfo := range s.Slots {
		if slotInfo.Slot == slot {
			return true
		}
	}
	return false
}
