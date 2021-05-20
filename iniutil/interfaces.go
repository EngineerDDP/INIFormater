/*
 * This file defines some basic constrains and interfaces for ini serilization.
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

const (
	splitor = "*****************************************************"
	space =   "                                                    "
	comment_prefix = "; *"
	section_splitor = "; *" + splitor + "\n"
	section_name_prefix = comment_prefix + " Section: "
	section_end = "\r\n\r\n"
	doc_header = section_splitor + comment_prefix + space + "*\n" + comment_prefix + "      This document is generated automatically.     *\n" + comment_prefix + space + "*\n" + section_splitor + section_end
)

type INIDocumentComponent interface{
	GetName() string
    SerializeToString() string
}

