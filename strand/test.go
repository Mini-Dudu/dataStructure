//
// Created by 嘟嘟 on 2020/11/3.
//


package strand

import (
	"fmt"
)

//测试串的各个功能
func TestStrand() {


	strand1 := NewStrand("KeAiDeXiaoDudu")
	strand2 := NewStrand("Dudu")
	fmt.Printf("新建两个串:%s  %s\n", strand1.Str, strand2.Str)

	Index := strand1.Index(strand2)

	fmt.Printf("定位子串:%s 在主串:%s 中的位置为:%d\n",strand2.Str, strand1.Str, Index)

	str := strand1.ConcatStr(strand2)
	fmt.Printf("拼接两个串:%s\n", str.Str)

	start := 7
	len := 8
	strand3, b := str.Substring(start,len)
	if !b {
		fmt.Printf("串%s 中不存在位序为%d开始长度为:%d的子串哦!\n", str.Str, start, len)
	}else {
		fmt.Printf("串%s 中位序为%d开始的长度为:%d的子串为:%s\n", str.Str, start, len, strand3.Str)
	}

	n := strand1.StrCompare(strand2)
	fmt.Printf("比较串:%s %s 的大小:", strand1.Str, strand2.Str)
	if n > 0 {
		fmt.Printf("串:%s >%s\n", strand1.Str, strand2.Str)
	}else if n < 0 {
		fmt.Printf("串:%s <%s\n", strand1.Str, strand2.Str)
	} else {
		fmt.Printf("串:%s =%s\n", strand1.Str, strand2.Str)
	}

}
