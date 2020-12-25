//
// Created by 嘟嘟 on 2020/10/29.
//

//strand: 线；串；
package strand

type Strand struct {
	Str []byte //存放串的byte字节切片,0索引处不存放,从1索引处开始存放,
	Len int    //串的长度
}

//接受一个字符串,返回一个串实例
func NewStrand(str string) *Strand {
	var strand Strand
	strand.Str = make([]byte, 1)
	for i := 0; i < len(str); i++ {
		strand.Str = append(strand.Str, str[i])
	}
	strand.Len = len(str)
	return &strand
}

//实现将 Str 赋值给地址为:strAddress 的字符串变量
func (Str Strand) StrAssign(strAddress *Strand) *Strand {
	//具体实现...

	strAddress.Str = Str.Str
	strAddress.Len = Str.Len

	return strAddress
}

//实现将 Str 复制到地址为: strAddress 的字符串变量
func (Str Strand) StrCopy(strAddress *Strand) *Strand {
	//具体实现...
	if Str.Len == 0 {
		return strAddress
	}

	strAddress.Str = make([]byte, 1)

	for i := 1; i < Str.Len; i++ {
		strAddress.Str = append(strAddress.Str, Str.Str[i])
		strAddress.Len++
	}

	return strAddress
}

//判断字符串是否为空
func (Str Strand) IsEmpty() bool {
	//具体实现...
	return Str.StrLen() == 0
}

//计算字符串的长度
func (Str Strand) StrLen() int {
	return Str.Len
}

//清空字符串
func (Str Strand) ClearStr() {
	Str.Len = 0
}

//销毁字符串
func (Str Strand) DestroyStr() {
	//具体实现...
	Str.Str = nil
	Str.Len = 0

}

//将两个字符串 Str str 连接成新的字符串:Strstr
func (Str Strand) ConcatStr(str *Strand) *Strand {
	//具体实现...
	var merge = Strand{
		Str: nil,
		Len: 0,
	}
	merge.Str = make([]byte, 1)

	for i := 1; i <= Str.StrLen(); i++ {
		merge.Str = append(merge.Str, Str.Str[i])
	}

	for i := 1; i <= str.StrLen(); i++ {
		merge.Str = append(merge.Str, str.Str[i])
	}

	merge.Len = Str.StrLen() + str.StrLen()

	return &merge
}

//求子串,即返回串Str 中的第 start 个字符起长度为 len 的子串,存在子串返回子串地址与true,否则返回地址(地址中没有数据)与false
func (Str Strand) Substring(start, len int) (*Strand, bool) {
	//具体实现...
	var substring = Strand{
		Str: nil,
		Len: 0,
	}
	substring.Str = make([]byte, 1)

	if start+len-1 > Str.Len {
		return &substring, false
	}

	for i := start; i < start+len; i++ {
		substring.Str = append(substring.Str, Str.Str[i])
	}

	substring.Len = len

	return &substring, true
}

//定位子串, 若主串:Str 中含有子串: substring,则返回第一次出现的位置,否则返回0
func (Str Strand) Index(substring *Strand) int {
	//具体实现...

	//i, n, m, 分别为主串开始索引、主串长度、子串长度
	i, n, m := 1, Str.StrLen(), substring.StrLen()
	//fmt.Println("+++++++++++",n,m)
	//n - m + 1 为子串在主串中寻找的次数
	for i <= n-m+1 {
		//在主串:Str 中寻找一个长度等于目标子串长度的子串: sub
		sub, _ := Str.Substring(i, substring.Len)

		//判断子串:sub 是否等于:subString
		if sub.StrCompare(substring) != 0 {
			i++
		} else {
			return i
		}

	}

	//经过以上for循环说明主串Str中没有子串subString
	return 0
}

//比较字符串的大小
//以utf8编码规则作比较
//若 Str > t 则返值大于0, 若 Str == t 相等返值等于0, 若 Str < t 则返值小于0,
func (Str Strand) StrCompare(t *Strand) int {
	for i := 1; i <= Str.Len && i <= t.Len; i++ {
		if Str.Str[i] != t.Str[i] {
			return int(Str.Str[i] - t.Str[i])
		}
	}
	return (Str.Len - t.Len)
}







//实现求KMP算法中next数组
func GetNextList(strand *Strand) []int {
	//串的前缀: 包含第一个字符,  且不包含最后一个字符的子串
	//串的后缀: 包含最后一个字符,且不包含第一个字符的子串

	//next数组的由来:
		//当第 N 个字符匹配失败时,由前 1 -- N-1 个字符组成的串记为S,则
		//next[N] = (S的最长相等前后缀长度+1)
	next := make([]int, 1)

	next = append(next, 0)
	next = append(next, 1)

	for i := 3; i <= strand.Len; i++ {
		max := 0
		for j := 1; j <= i -2; j++ {
			//前缀
			prefix, _ := strand.Substring(1,j)

			//后缀
			postfix, _ := strand.Substring(i - j, j)

			//判断前后缀是否相等
			if prefix.StrCompare(postfix) == 0 {
				max = j
			}

		}

		next = append(next, max + 1)
	}

	return next
}

//KMP算法中next数组的优化
func GetNextVal(strand *Strand) []int {
	nextVal := make([]int, 1)
	nextVal = append(nextVal,0)

	next := GetNextList(strand)

	for i := 2; i <= strand.Len; i++ {
		if strand.Str[next[i]] == strand.Str[i] {
			nextVal = append(nextVal,nextVal[next[i]])
		} else {
			nextVal = append(nextVal,next[i])
		}
	}

	return nextVal
}

//KMP算法:朴素模式匹配算法的优化版
func (Str Strand) Index_KMP(substring *Strand) int {

	//next := GetNextList(substring)
	nextVal := GetNextVal(substring)
	i, j := 1, 1

	for i <= Str.StrLen() && j <= substring.StrLen() {
		if j == 0 || Str.Str[i] == substring.Str[j] {
			i++
			j++
		}else {
			j = nextVal[j]
		}
	}

	//进过以上循环,若j>子串的长度,说明匹配成功
	if j > substring.Len {
		return i - substring.Len
	} else {
		return 0
	}

}
