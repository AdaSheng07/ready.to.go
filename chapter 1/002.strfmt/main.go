package main

import (
	"fmt"
)

func main() {
	// 字符串的运算：只支持加法拼接字符串
	fmt.Println(
		"And we count these moments. \n" +
			"我们珍惜这些时刻 💕️️\n" +
			"These moments when we dare to aim higher, \n" +
			"这些我们敢于追求卓越 🚀\n" +
			"to break barriers, \n" +
			"突破障碍 🔥\n" +
			"to reach for the stars, \n" +
			"探索星空 🌕\n" +
			"to make the unknown known. \n" +
			"揭开未知面纱的时刻 🧬\n" +
			"We count these moments as our proudest achievements.\n" +
			"我们将这些时刻视为我们最值得骄傲的成就 👑")

	// 格式化输出：参考文档 http://doc.golang.ltd/ 中的 package fmt
	// %s: 直接输出字符串或者[]byte
	// %q: 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	// %x: 每个字节用两字符十六进制数表示（使用a-f）
	// %X: 每个字节用两字符十六进制数表示（使用A-F）
	fmt.Printf("This is %s, %s\n", "December", "2021")
	fmt.Printf("This is %q, %q\n", "December", "2021")
	fmt.Printf("This is %x, %x\n", "December", "2021")
	fmt.Printf("This is %X, %X\n", "December", "2021")

}
