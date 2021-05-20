/*
 * This file described how to load and serialize ini comments.
 * Definitions of Comment concludes anything that cannot be described 
 * as Attribute, and Comment class will automatically delete any 
 * unrecognized string doesn't start with ';'.
 * 
 * Comment is an implementation of INIDocumentComponent interface
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

import "strings"

type Comment struct{
	// 注释
	content string
}

/*
 * Build comment only document part based on input string
 * return *Comment and "OK" if success
 * return *Comment and 
 */
 func newComment(line string) (object *Comment, err string){
	strings.TrimSpace(line)
	// find comment
	var comment string = ""
	_pos_split := strings.Index(line, ";")
	if _pos_split != -1 {
		comment = line[_pos_split+1:]
	} // else pass

	if _pos_split == 0{
		err = "OK"
	} else {
		err = "Unused prefix before comments."
	}
	object = &Comment{
		content: comment,
	}
	return
}

func (self *Comment) GetName() string{
	return "Comment"
}

func (self *Comment) getComment() string{
	return self.content
}

func (self *Comment) SerializeToString() string{
	if self.content == ""{
		return "\n"
	} else {
		return comment_prefix + self.content + "\n"
	}
}