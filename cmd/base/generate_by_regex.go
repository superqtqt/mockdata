package base

import (
	"github.com/spf13/cast"
	"math/rand"
	"regexp"
	"strings"
)

var BigBracketsMatch = regexp.MustCompile("\\{([0-9]+),([0-9]+)?\\}")
var NumberMatch = regexp.MustCompile("\\d+")
var AnyMatch = regexp.MustCompile("\\*")
var MustMatch = regexp.MustCompile("\\+")
var OneZeroMatch = regexp.MustCompile("\\?")

type Generate struct {
	MaxLength int
}

func (g *Generate) Generate(regex string, maxLength int) string {
	return g.generateRegex(regex, 1, 1)
}

func (g *Generate) generateRegex(regex string, min, max int) string {
	str := strings.Builder{}
	match := BigBracketsMatch.FindAllString(regex, -1)
	if len(match) > 0 {
		nodes := BigBracketsMatch.Split(regex, -1)
		for i := range match {
			innerMin, innerMax := getBigBracketsMinMax(match[i])
			valueNode := g.generateRegex(nodes[i], innerMin, innerMax)
			str.WriteString(valueNode)
		}
		return str.String()
	}
	match = AnyMatch.FindAllString(regex, -1)
	if len(match) > 0 {
		for i := range match {
			valueNode := g.generateRegex(match[i], 0, g.MaxLength)
			str.WriteString(valueNode)
		}
		return str.String()
	}

	match = MustMatch.FindAllString(regex, -1)
	if len(match) > 0 {
		for i := range match {
			valueNode := g.generateRegex(match[i], 1, g.MaxLength)
			str.WriteString(valueNode)
		}
		return str.String()
	}
	match = OneZeroMatch.FindAllString(regex, -1)
	if len(match) > 0 {
		for i := range match {
			valueNode := g.generateRegex(match[i], 0, 1)
			str.WriteString(valueNode)
		}
		return str.String()
	}
	v := rand.Intn(max)
	for i := 0; i < 4 && v < min; i++ {
		v = rand.Intn(max)
	}
	if v < min {
		v = min
	}
	wordArray := generateWordArray(regex)
	for i := 0; i < v; i++ {
		index := rand.Intn(len(wordArray))
		str.WriteString(wordArray[index])
	}
	return str.String()
}

var BigWord = []string{"A", "B", "C", "D", "E", "F", "G", "H", "K"}
var littleWord = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
var numberWord = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

//解法参照上面:按 [],() 转义符、.匹配符来

//var middleBracketsMatch = regexp.MustCompile("\\[((?!\\[).)*\\]")
//var littleBracketsMatch = regexp.MustCompile("\\(((?!\\().)*\\)")
//var encodeMatch = regexp.MustCompile("\\\\\\w")

func generateWordArray(regex string) []string {
	rs := make([]string, 0)
	regexArray := strings.Split(regex, "")
	for i := 0; i < len(regexArray); i++ {
		word := regexArray[i]
		switch word {
		case ".":
			rs = append(rs, BigWord...)
			rs = append(rs, littleWord...)
			rs = append(rs, numberWord...)
		case "[":
			var newItems []string
			i,newItems=scanMiddleContent(regexArray,i+1)
			rs=append(rs, newItems...)
		case "(":
			var newItems []string
			i,newItems=scanLittleContent(regexArray,i+1)
			rs=append(rs, newItems...)
		case "\\":
			rs=append(rs,getEncodeMean(regexArray[i+1])...)
			i++
		default:
			rs=append(rs,word)
		}
	}
	return rs
}

func scanMiddleContent(src []string, i int) (int, []string) {
	rs := make([]string, 0)
	currentWord := src[i]
	for ; currentWord != "]"; {
		switch {
		case currentWord == "\\":
			i=i+2
			rs = append(rs, getEncodeMean(rs[i])...)
		case i+2 < len(src) && src[i+1] == "-":
			//生成a-z这种情况
			begin, end := currentWord, src[i+2]
			rs=append(rs, appendArray(numberWord,begin,end)...)
			rs=append(rs, appendArray(littleWord,begin,end)...)
			rs=append(rs, appendArray(BigWord,begin,end)...)
			i = i+3
		default:
			rs = append(rs, currentWord)
			i++
		}
		currentWord = src[i]
	}
	return i,rs
}

func scanLittleContent(src []string,i int)(int, []string)  {
	rs := make([]string, 0)
	currentWord := src[i]
	str:=strings.Builder{}
	for ; currentWord != ")"; i++ {
		currentWord = src[i]
		switch {
		case currentWord == "\\":
			i++
		case currentWord=="|":
			rs=append(rs,str.String())
			str.Reset()
		default:
			str.WriteString(currentWord)
		}
	}
	rs=append(rs,str.String())
	return i,rs
}

func appendArray(target []string, begin, end string) []string {
	beginAppend, rs := false, make([]string, 0)
	for i := range target {
		if target[i] == begin {
			beginAppend = true
		}
		if target[i] == end {
			beginAppend = false
			rs = append(rs, target[i])
			break
		}
		if beginAppend {
			rs = append(rs, target[i])
		}
	}
	return rs
}




func getEncodeMean(word string) []string {
	switch word {
	//TODO 加入更多的转义符
	case "d":
		return numberWord

	}
	return []string{word}
}

//var GenerateByRegex= func(regex string,min,max int)string {
//
//
//}

func getBigBracketsMinMax(str string) (int, int) {
	v := NumberMatch.FindAllString(str, -1)
	min := cast.ToInt(v[0])
	max:=10
	if len(v) >= 0 {
		max = cast.ToInt(v[1])
	}
	return min,max
}

type FixedValueGenerate struct {
}

func (f *FixedValueGenerate) Generate(regex []string, min, max int) string {
	v := rand.Intn(max)
	for i := 0; i < 4 && v < min; i++ {
		v = rand.Intn(max)
	}
	if v < min {
		v = min
	}
	str := strings.Builder{}
	for i := 0; i < v; i++ {
		index := rand.Intn(len(regex))
		str.WriteString(regex[index])
	}
	return str.String()
}

type RegexValueGenerate struct {
}

func (r *RegexValueGenerate) Generate(regex []string, min, max int) string {
	panic("implement me")
}
