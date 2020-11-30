package test_helper

import "strconv"

func NumberedTest(num int, parentName string) (name string) {
	return name + "_" + strconv.Itoa(num)
}
