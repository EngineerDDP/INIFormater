/*
 * This file described how to load and serialize ini Section.
 * When serializing, Section always start with comment, this is for
 * easier idenfication for human reading.
 *
 * Section is an implementation of INIDocumentComponent interface
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

import "strings"

type Section struct{
	// 名称
	name string
	// 属性列表
	attributes map[string]*Attribute
	// 文档部件
	doc_content []INIDocumentComponent
}

func newSection(line string) (object *Section, err string){
	start, end := strings.Index(line, "["), strings.Index(line, "]")
	if start != -1 && end != -1 {
		err = "OK"
		object = &Section{
			name: line[start + 1 : end],
			attributes: make(map[string]*Attribute),
			doc_content: make([]INIDocumentComponent, 0),
	   }
	} else {
		err = "Parse line failed."
	}
    return
}

func (self *Section) GetName() string{
	return self.name
}

func (self *Section) Get(key string) *Attribute{
    val, ok := self.attributes[key]
	if ok {
		return val
	} else {
		panic(ok)
	}
}

func (self *Section) TryGet(key string) (obj *Attribute, ok bool){
	obj, ok = self.attributes[key]
	return
}

func (self *Section) Add(att *Attribute) bool{
	if _, ok := self.attributes[att.GetName()]; !ok {
		self.attributes[att.GetName()] = att
		self.doc_content = append(self.doc_content, att)
		return true
	}
	return false
}

func (self *Section) AddComment(comment *Comment){
	self.doc_content = append(self.doc_content, comment)
}

func (self *Section) Remove(key string) bool{
	if _, ok := self.attributes[key]; ok{
		// 从索引中删除
		delete(self.attributes, key)

		// 从内容中删除
		mark := false
		for i := 0; i < len(self.doc_content) - 1; i++{
			if self.doc_content[i].GetName() == key{
				mark = true
			}
			if mark {
				self.doc_content[i] = self.doc_content[i+1]
			}
		}
		return true
	}
	return false
}

func (self *Section) GetAttributes() []*Attribute{
	res := make([]*Attribute, 0, len(self.attributes))
	for _, v := range self.attributes{
		res = append(res, v)
	}
	return res
}

func (self *Section) SerializeToString() string{
    var res string
	res += section_splitor
	res += section_name_prefix + self.name + "\n"
	res += "\n"
	res += "[" + self.name + "]\n"
	for _, doc := range self.doc_content {
		res += doc.SerializeToString()
	}
	res += section_end
	return res
}
