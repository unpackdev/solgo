package storage

import (
	"sort"
)

type StorageLayout struct {
	Slots []*SlotDescriptor
}

func (s *StorageLayout) GetSlots() []*SlotDescriptor {
	return s.Slots
}

func (s *StorageLayout) GetSlot(slot int64) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Slot == slot {
			return slotInfo
		}
	}
	return nil
}

func (s *StorageLayout) GetSlotByName(name string) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Name == name {
			return slotInfo
		}
	}
	return nil
}

func (s *StorageLayout) GetSlotByOffset(offset int64) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Offset == offset {
			return slotInfo
		}
	}
	return nil
}

func (s *StorageLayout) GetSlotByType(t string) *SlotDescriptor {
	for _, slotInfo := range s.Slots {
		if slotInfo.Type == t {
			return slotInfo
		}
	}
	return nil
}

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

func (s *StorageLayout) SlotExists(slot int64) bool {
	for _, slotInfo := range s.Slots {
		if slotInfo.Slot == slot {
			return true
		}
	}
	return false
}
