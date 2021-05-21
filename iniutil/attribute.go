/*
 * This file described how to load and serialize ini attribute
 * Attribute is an implementation of INIDocumentComponent interface
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

import "strings"

type Attribute struct{
	// 键
	key string
	// 值
	val string
    // 附加注释
	comment string
}

/*
 * Build a new KeyAttribute object based on input string
 * return *KeyAttribute and "OK" if success
 * return nil and error message if not
 */
 func newAttribute(line string) (object *Attribute, err string){
	// find comment
	_pos_split := strings.Index(line, ";")
	var comment string
	if _pos_split >= 0 {
		comment = line[_pos_split+1:]
		line = line[:_pos_split]
	} else {
		comment = ""
	}
	
	// find key and val
    _pos_equal := strings.Index(line, "=")
	if _pos_equal == -1{
		object = nil
		err = "\"=\" not found."
		return 
	}
	key := line[:_pos_equal]
	val := line[_pos_equal+1:]

	object = &Attribute{
		key: key,
		val: val,
		comment: comment,
	}
	err = "OK"
	return
}

func NewAttribute(key string) *Attribute{
	return &Attribute{
		key: key,
		val: "",
		comment: "",
	}
}

func (self *Attribute) GetName() string{
	return self.key
}

func (self *Attribute) getComment() string{
	return self.comment
}

func (self *Attribute) Get() string{
	return self.val
}

func (self *Attribute) GetValues() []string{
	return strings.Split(self.val, ",")
}

func (self *Attribute) setValue(vals []string){
	self.val = strings.Join(vals, ",")
}

func (self *Attribute) Check(match string) bool{
	if strings.Index(self.val, match) == -1 {
		return false
	} else {
		return true
	}
}

func (self *Attribute) Append(val string) {
	self.val += val
}

func (self *Attribute) Set(val string) {
	self.val = val
}

func (self *Attribute) Remove(val string) {
	var res string
	for _, str := range self.GetValues(){
		if str != val{
			res += str + ","
		} // else pass
	}
	if len(res) > 0{
		self.val = res[:len(res)-1]
	} else {
		self.val = ""
	}
}

func (self *Attribute) SerializeToString() string{
    result := self.key + " = " + self.val
	if self.comment != "" {
		result += "\t\t\t; " + self.comment
	}
	result += "\n"
	return result
}
