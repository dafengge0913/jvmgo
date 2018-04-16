package heap

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) IsImplements(inter *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == inter || i.IsSubInterfaceOf(inter) {
				return true
			}
		}
	}
	return false
}

func (class *Class) IsSubInterfaceOf(inter *Class) bool {
	for _, superInter := range class.interfaces {
		if superInter == inter || superInter.IsSubInterfaceOf(inter) {
			return true
		}
	}
	return false
}

func (class *Class) IsAssignableFrom(other *Class) bool {
	if other == class {
		return true
	}
	if !class.IsInterface() {
		return other.IsSubClassOf(class)
	} else {
		return other.IsImplements(class)
	}
}

func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(class)
}
