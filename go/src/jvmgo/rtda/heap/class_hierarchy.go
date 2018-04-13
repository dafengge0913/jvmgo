package heap

func (class *Class) isSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) isImplements(inter *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == inter || i.isSubInterfaceOf(inter) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(inter *Class) bool {
	for _, superInter := range class.interfaces {
		if superInter == inter || superInter.isSubInterfaceOf(inter) {
			return true
		}
	}
	return false
}

func (class *Class) isAssignableFrom(other *Class) bool {
	if other == class {
		return true
	}
	if !class.IsInterface() {
		return other.isSubClassOf(class)
	} else {
		return other.isImplements(class)
	}
}
