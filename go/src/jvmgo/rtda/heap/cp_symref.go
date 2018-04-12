package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (ref *SymRef) ResolvedClass() *Class {
	if ref.class == nil {
		ref.ResolvedClassRef()
	}
	return ref.class
}

// 类D通过符号引用N 引用类C 要解析N
// 先用D的类加载器加载C 然后检查D是否有权限访问C
func (ref *SymRef) ResolvedClassRef() {
	d := ref.cp.class
	c := d.loader.LoadClass(ref.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.class = c
}
