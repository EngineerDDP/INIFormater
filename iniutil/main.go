/*
 * This file described how to read a ini file.
 * author: Engineer_DDP
 * date: 2021-05-20
 */

package iniutil

import (
	"bufio"
	"os"
	"strings"
)


func LoadINIDocument(filename string) (doc *Document, err string){
    _, stat := os.Stat(filename)
	if stat != nil {
		err = "Open file failed."
		return 
	}

	file, stat := os.Open(filename)
	if stat != nil {
		err = "Read file failed."
		return 
	}

	doc = newDocument(filename)
	// 在读的Section
	var sec *Section

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// Check section
		_sec, _err := newSection(text)
		if _err == "OK"{
			sec = _sec
			doc.Add(sec.GetName(), sec)
			continue
		}

		// Check Attribute
		att, _err := newAttribute(text)
		if _err == "OK"{
			if sec != nil{
				sec.Add(att)
			}
			continue
		}

		// Check comment
		com, _ := newComment(text)
		if sec == nil{
			doc.AddContent(com)
		} else {
			sec.AddComment(com)
		}
	}

	file.Close()
	err = "OK"
	return
}

func LoadINIString(content string, doc_name string) (doc *Document, err string){
	lines := strings.Split(content, "\n")

	doc = newDocument(doc_name)
	// 在读的Section
	var sec *Section
	for _, text := range lines{
		_sec, _err := newSection(text)
		if _err == "OK"{
			sec = _sec
			doc.Add(sec.GetName(), sec)
			continue
		}

		// Check Attribute
		att, _err := newAttribute(text)
		if _err == "OK"{
			if sec != nil{
				sec.Add(att)
			}
			continue
		}

		// Check comment
		com, _ := newComment(text)
		if sec == nil{
			doc.AddContent(com)
		} else {
			sec.AddComment(com)
		}
	}

	err = "OK"
	return
}