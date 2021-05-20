package iniutil_test

import (
	"fmt"
	"testing"

	"example.com/hello/iniutil"
)

const (
	Normal = `
[A]
1=a
2=b ; aaa
; abc

[B]
A=A
`
	NoEqual = `
adgsad
[A]
afdsf
A=13415dfsf
`
	NoSection = `
afdsaf
afsdf=afdsafas
gfdgd;
;gdfgd
;;;;;;;
`
	CommentOnly = `
;adfsf
;afasfsd


;afasfs
;;;;;
`
	SectionOnly = `
[A]
[B]
[
	A
]
[SFSDF]
[AFDFSF
AFASFDSF]
`
)

func TestAttribute(t *testing.T){
	var doc *iniutil.Document
	doc,_ = iniutil.LoadINIString(Normal, "1")
	fmt.Println(doc.SerializeToString())

	doc,_ = iniutil.LoadINIString(NoEqual, "2")
	fmt.Println(doc.SerializeToString())

	doc,_ = iniutil.LoadINIString(NoSection, "3")
	fmt.Println(doc.SerializeToString())

	doc,_ = iniutil.LoadINIString(CommentOnly, "4")
	fmt.Println(doc.SerializeToString())

	doc,_ = iniutil.LoadINIString(SectionOnly, "5")
	fmt.Println(doc.SerializeToString())

	doc.Get("A").Add(iniutil.NewAttribute("ABCDEF"))
	att := doc.Get("A").Get("ABCDEF")
	att.Set("A,B,C")
	att.Remove("B")
	att.GetValues()
	if att.Check("B") {
		t.Error("Delete failed.")
	}
	att.Append("B")
	if !att.Check("B") {
		t.Error("Append failed.")
	}
}