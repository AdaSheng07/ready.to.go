package main

import (
	"fmt"
)

// print a poem
/* fmt.Print的几种变种
Print:   输出到控制台,不接受任何格式化操作
Println: 输出到控制台并换行
Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出别的类型）
Sprintf：格式化并返回一个字符串而不带任何输出
Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout
*/

func main() {
	fmt.Println("Do not go gentle into that good night," +
		"\nOld age should burn and rave at close of day;" +
		"\nRage, rage against the dying of the light.")
	fmt.Println("\nThough wise men at their end know dark is right," +
		"\nBecause their words had forked no lightning they" +
		"\nDo not go gentle into that good night.")
}
