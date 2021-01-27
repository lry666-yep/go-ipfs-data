package chaintypes

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Link struct{ x ipld.Link }
type Link = *_Link

func (n Link) Link() ipld.Link {
	return n.x
}
func (_Link__Prototype) FromLink(v ipld.Link) (Link, error) {
	n := _Link{v}
	return &n, nil
}

type _Link__Maybe struct {
	m schema.Maybe
	v Link
}
type MaybeLink = *_Link__Maybe

func (m MaybeLink) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeLink) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeLink) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeLink) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Absent
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeLink) Must() Link {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Link)(&_Link{})
var _ schema.TypedNode = (Link)(&_Link{})

func (Link) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link) LookupByString(string) (ipld.Node, error) {
	return mixins.Link{"chaintypes.Link"}.LookupByString("")
}
func (Link) LookupByNode(ipld.Node) (ipld.Node, error) {
	return mixins.Link{"chaintypes.Link"}.LookupByNode(nil)
}
func (Link) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Link{"chaintypes.Link"}.LookupByIndex(0)
}
func (Link) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return mixins.Link{"chaintypes.Link"}.LookupBySegment(seg)
}
func (Link) MapIterator() ipld.MapIterator {
	return nil
}
func (Link) ListIterator() ipld.ListIterator {
	return nil
}
func (Link) Length() int {
	return -1
}
func (Link) IsAbsent() bool {
	return false
}
func (Link) IsNull() bool {
	return false
}
func (Link) AsBool() (bool, error) {
	return mixins.Link{"chaintypes.Link"}.AsBool()
}
func (Link) AsInt() (int, error) {
	return mixins.Link{"chaintypes.Link"}.AsInt()
}
func (Link) AsFloat() (float64, error) {
	return mixins.Link{"chaintypes.Link"}.AsFloat()
}
func (Link) AsString() (string, error) {
	return mixins.Link{"chaintypes.Link"}.AsString()
}
func (Link) AsBytes() ([]byte, error) {
	return mixins.Link{"chaintypes.Link"}.AsBytes()
}
func (n Link) AsLink() (ipld.Link, error) {
	return n.x, nil
}
func (Link) Prototype() ipld.NodePrototype {
	return _Link__Prototype{}
}

type _Link__Prototype struct{}

func (_Link__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Link__Builder
	nb.Reset()
	return &nb
}

type _Link__Builder struct {
	_Link__Assembler
}

func (nb *_Link__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Link__Builder) Reset() {
	var w _Link
	var m schema.Maybe
	*nb = _Link__Builder{_Link__Assembler{w: &w, m: &m}}
}

type _Link__Assembler struct {
	w *_Link
	m *schema.Maybe
}

func (na *_Link__Assembler) reset() {}
func (_Link__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.LinkAssembler{"chaintypes.Link"}.BeginMap(0)
}
func (_Link__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.LinkAssembler{"chaintypes.Link"}.BeginList(0)
}
func (na *_Link__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.LinkAssembler{"chaintypes.Link"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	panic("unreachable")
}
func (_Link__Assembler) AssignBool(bool) error {
	return mixins.LinkAssembler{"chaintypes.Link"}.AssignBool(false)
}
func (_Link__Assembler) AssignInt(int) error {
	return mixins.LinkAssembler{"chaintypes.Link"}.AssignInt(0)
}
func (_Link__Assembler) AssignFloat(float64) error {
	return mixins.LinkAssembler{"chaintypes.Link"}.AssignFloat(0)
}
func (_Link__Assembler) AssignString(string) error {
	return mixins.LinkAssembler{"chaintypes.Link"}.AssignString("")
}
func (_Link__Assembler) AssignBytes([]byte) error {
	return mixins.LinkAssembler{"chaintypes.Link"}.AssignBytes(nil)
}
func (na *_Link__Assembler) AssignLink(v ipld.Link) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	if na.w == nil {
		na.w = &_Link{}
	}
	na.w.x = v
	*na.m = schema.Maybe_Value
	return nil
}
func (na *_Link__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Link); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v2, err := v.AsLink(); err != nil {
		return err
	} else {
		return na.AssignLink(v2)
	}
}
func (_Link__Assembler) Prototype() ipld.NodePrototype {
	return _Link__Prototype{}
}
func (Link) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Link) Representation() ipld.Node {
	return (*_Link__Repr)(n)
}

type _Link__Repr = _Link

var _ ipld.Node = &_Link__Repr{}

type _Link__ReprPrototype = _Link__Prototype
type _Link__ReprAssembler = _Link__Assembler
