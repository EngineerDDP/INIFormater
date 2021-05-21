package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/hello/iniutil"
)


func main() {
	doc, err := iniutil.LoadINIDocument("rulesmd.ini")
	if err != "OK" {
		fmt.Println("Load file error.")
		return
	}

	// Types
	var types = [...]string{"InfantryTypes", "VehicleTypes", "AircraftTypes", "BuildingTypes"}
	var buildin_objs = [...]string{"PROC", "BARRACKS", "POWER", "FACTORY", "CONYARD", "TECH", "RADAR"}

	// revt idx
	var objs = make(map[string]string)

	// load buildin
	for _, v := range buildin_objs{
		objs[v] = "buildin"
	}

	// load all object
	for _, type_name := range types {
		if sec, ok := doc.TryGet(type_name); ok{
			for _, att := range sec.GetAttributes(){
				objs[att.Get()] = sec.GetName()
			}
		}
	}

	var check_requisite = func (keyname string, sec *iniutil.Section){
		if att, ok := sec.TryGet(keyname); ok{
			for _, key := range att.GetValues(){
				// if keyname cannot found
				if _, ok := objs[strings.TrimSpace(key)]; !ok{
					att.Remove(key)
					fmt.Println("Invalid " + keyname + " , in [" + sec.GetName() + "], KEY=" + key)
				}
			}
			// fmt.Println("Pass Section [" + sec.GetName() + "]" + "\tKEY :" + keyname)
		}
	}

	var check_list = [...]string{"Prerequisite", "NegativePrerequisite", "AlternatePrerequisite"}

	var check_section_link = func (keyname string, sec *iniutil.Section){
		if att, ok := sec.TryGet(keyname); ok{
			for _, key := range att.GetValues(){
				// if keyname cannot found
				if _, ok := doc.TryGet(strings.TrimSpace(key)); !ok{
					att.Remove(key)
					// remove key if empty
					if att.Get() == ""{
						sec.Remove(att.GetName())
						fmt.Println("Invalid " + keyname + "  \tin [" + sec.GetName() + "]  \tKEY Removed")
					} else {
						fmt.Println("Invalid " + keyname + "  \tin [" + sec.GetName() + "]  \tKEY=" + key)
					}
				}
			}
			// fmt.Println("Pass Section [" + sec.GetName() + "]" + "\tKEY :" + keyname)
		}
	}

	var link_check_list = [...]string{"Primary", "ElitePrimary", "Secondary", "EliteSecondary", "Projectile", "Warhead"}
	
	// get all country
	var country_list map[string]bool = make(map[string]bool)
	for _, att := range doc.Get("Sides").GetAttributes(){
		for _, country := range att.GetValues(){
			country_list[country] = true
		}
	}

	// check country
	var country_check = func (keyname string, sec *iniutil.Section){
		if att, ok := sec.TryGet(keyname); ok{
			for _, key := range att.GetValues(){
				// if keyname cannot found
				if _, ok := country_list[strings.TrimSpace(key)]; !ok{
					att.Remove(key)
					fmt.Println("Invalid " + keyname + " , in [" + sec.GetName() + "], Country=" + key)
				}
			}
			fmt.Println("Pass Section [" + sec.GetName() + "]" + "\tKEY :" + keyname)
		}
	}

	// check obj exists
	for obj, sec_name := range objs{
		if sec, ok := doc.TryGet(obj); !ok{
			// remove invalid
			if sec, ok := doc.TryGet(sec_name); ok{
				sec.Remove(obj)
				fmt.Println("Remove invalid registration: [" + obj + "]")
			}
		} else {
			// check link exists
			for _, link := range link_check_list{
				if false {
					check_section_link(link, sec)
				}
			}
		}
		// fmt.Println("Valid Section [" + obj + "]")
	}

	// check all Section
	for _, sec := range doc.GetSections(){
		for _, req := range check_list{
			check_requisite(req, sec)
			country_check("Owner", sec)
		}
	}

	os.WriteFile("auto_" + doc.GetName(), []byte(doc.SerializeToString()), os.ModePerm)
}
