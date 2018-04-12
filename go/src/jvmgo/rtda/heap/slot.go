package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (slot Slots) SetInt(index uint, val int32) {
	slot[index].num = val
}

func (slot Slots) GetInt(index uint) int32 {
	return slot[index].num
}

func (slot Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	slot[index].num = int32(bits)
}

func (slot Slots) GetFloat(index uint) float32 {
	bits := uint32(slot[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (slot Slots) SetLong(index uint, val int64) {
	slot[index].num = int32(val)
	slot[index+1].num = int32(val >> 32)
}

func (slot Slots) GetLong(index uint) int64 {
	low := uint32(slot[index].num)
	high := uint32(slot[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (slot Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	slot.SetLong(index, int64(bits))
}

func (slot Slots) GetDouble(index uint) float64 {
	bits := uint64(slot.GetLong(index))
	return math.Float64frombits(bits)
}

func (slot Slots) SetRef(index uint, ref *Object) {
	slot[index].ref = ref
}

func (slot Slots) GetRef(index uint) *Object {
	return slot[index].ref
}
