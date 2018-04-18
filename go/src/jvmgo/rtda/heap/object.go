package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

func (obj *Object) Extra() interface{} {
	return obj.extra
}

func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(obj.class)
}

func (obj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (obj *Object) GetRefVar(name, descriptor string) *Object {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(field.slotId)
}
