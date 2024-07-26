package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	columns := []string{
		"helloah",
		"hellobh",
		"helloch",
		"hellodh",
		"helloeh",
		"hellofh",
		"hellogh",
		"hellohh",
		"helloih",
		"hellojh",
		"hellokh",
		"hellolh",
		"hellomh",
		"hellonh",
		"hellooh",
		"helloph",
		"helloqh",
		"hellorh",
		"hellosh",
		"helloth",
		"hellouh",
		"hellovh",
		"hellowh",
		"helloxh",
		"helloa",
		"hellob",
		"helloc",
		"hellod",
		"helloe",
		"hellof",
		"hellog",
		"helloh",
		"helloi",
		"helloj",
		"hellok",
		"hellol",
		"hellom",
		"hellon",
		"helloo",
		"hellop",
		"helloq",
		"hellor",
		"hellos",
		"hellot",
		"hellou",
		"hellov",
		"hellow",
		"hellox",
		"ah",
		"bh",
		"ch",
		"dh",
		"eh",
		"fh",
		"gh",
		"hh",
		"ih",
		"jh",
		"kh",
		"lh",
		"mh",
		"nh",
		"oh",
		"ph",
		"qh",
		"rh",
		"sh",
		"th",
		"uh",
		"vh",
		"wh",
		"xh",
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
	}
	ssi := make([]map[string]any, 0, 1000000)
	for i := 0; i < cap(ssi); i++ {
		m := make(map[string]any)
		for _, c := range columns {
			m[c] = c
		}
		m["acustom1"] = "custom1"
		m["acustom2"] = "custom2"
		m["acustom3"] = "custom3"
		m["acustom4"] = "custom4"
		m["acustom5"] = "custom5"
		m["acustom6"] = "custom6"
		m["acustom7"] = "custom7"
		m["acustom8"] = "custom8"
		m["acustom9"] = "custom9"
		ssi = append(ssi, m)
	}

	start := time.Now()
	calcAtOnce(ssi, columns)
	fmt.Println("calcAtOnce:", time.Since(start))
	start = time.Now()
	calcEachTime(ssi, columns)
	fmt.Println("calcEachTime:", time.Since(start))
}

func calcAtOnce(ssi []map[string]any, columns []string) {
	baseColumnsCache := make(map[string]struct{})
	customColumns := make([]string, 0)

	for _, item := range ssi {
		for k := range item {
			if _, ok := baseColumnsCache[k]; ok {
				continue
			}
			// if the column is not found in the base columns, then it is a custom column
			if idx := slices.Index(columns, k); idx == -1 {
				customColumns = append(customColumns, k)
			} else {
				baseColumnsCache[k] = struct{}{}
			}
		}
	}
}

func calcEachTime(ssi []map[string]any, columns []string) {
	customColumns := make(map[string]struct{})

	for _, item := range ssi {
		for k := range item {
			// if the column is not found in the base columns, then it is a custom column
			if idx := slices.Index(columns, k); idx == -1 {
				customColumns[k] = struct{}{}
			}
		}
	}
}
