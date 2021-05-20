/*
 * This file described how to load and serialize ini Document.
 * Document is an implementation of INIDocumentComponent interface
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

type Document struct{
	// 文档名称
	name string
	// Section 列表
	sections map[string]*Section
	// 文档部件
	doc_content []INIDocumentComponent
}

func (self *Document) GetName() string{
	return self.name
}

func newDocument(name string) *Document{
	return &Document{
		name: name,
		sections: make(map[string]*Section),
		doc_content: make([]INIDocumentComponent, 0),
	}
}

func (self *Document) Get(key string) *Section{
	result, ok := self.sections[key]
	if ok {
		return result
	} else {
		panic(ok)
	}
}

func (self *Document) Add(key string, val *Section) {
	if _, ok := self.sections[key]; !ok {
		self.sections[key] = val
		self.doc_content = append(self.doc_content, val)
	}
}

func (self *Document) AddContent(content INIDocumentComponent){
	self.doc_content = append(self.doc_content, content)
}

func (self *Document) SerializeToString() string {
	var res string = doc_header
	
	for _, doc := range self.doc_content {
		res += doc.SerializeToString()
	}
	return res
}
