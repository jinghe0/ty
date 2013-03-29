package fun

import (
	"reflect"

	"github.com/BurntSushi/ty"
)

// Keys has a parametric type:
//
//	func(m map[A]B) []A
//
// Keys returns a list of the keys of `m` in an unspecified order.
func Keys(m interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]ty.B) []ty.A),
		m)
	vm, tkeys := uni.Args[0], uni.Returns[0]

	vkeys := reflect.MakeSlice(tkeys, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vkeys.Index(i).Set(vkey)
	}
	return vkeys.Interface()
}

// Values has a parametric type:
//
//	func(m map[A]B) []B
//
// Values returns a list of the values of `m` in an unspecified order.
func Values(m interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]ty.B) []ty.B),
		m)
	vm, tvals := uni.Args[0], uni.Returns[0]

	vvals := reflect.MakeSlice(tvals, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vvals.Index(i).Set(vm.MapIndex(vkey))
	}
	return vvals.Interface()
}

// func MapMerge(m1, m2 interface{}) interface{} {
// }
