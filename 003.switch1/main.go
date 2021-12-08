package main

import "fmt"

/* switch
1. 基于不同条件执行不同动作，每个 case 分支唯一，从上至下逐一测试，直到匹配为止
2. 执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break
3. 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果需要执行后面的 case，可使用 fallthrough

switch语句语法：
	switch var1 {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
	* 可以同时测试多个可能符合条件的值，使用逗号分割，例如：case val1, val2, val3
*/

func main() {
	var money int
	switch {
	case money >= 20:
		fmt.Println("点个外卖")
	case money >= 200:
		fmt.Println("下个馆子")
	case money >= 2000:
		fmt.Println("去米其林")
	case money >= 20000:
		fmt.Println("出国玩一圈")
	default: // 非必需分支，如果没有此分支，无输出
		fmt.Println("饿着吧")
	}
}
