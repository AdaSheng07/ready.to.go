package main

import "testing"

/*
案例1：
- 王强第一次录入的时候，他的体脂是38
- 王强第二次录入的时候，他的体脂是32
此时，王强的最佳体脂是32
- 李静录入她的体脂是28
此时，李静的最佳体脂是28

李静排名第一，体脂28；王强排名第二，体脂32。
*/
func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		if rankOfWQ != 1 {
			t.Fatalf("预期王强的排名是第一，实际排名是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率是0.32，实际是：%f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强的排名是第二，实际排名是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率是0.32，实际是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfLJ, fatRateOfLJ := getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静的排名是第一，实际排名是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是0.28，实际是：%f", fatRateOfLJ)
		}
	}

}

/*
案例2：
- 王强录入体脂是38
- 张伟录入体脂是38
- 李静录入体脂是28

李静排名第一，体脂28，王强、张伟排名第二，体脂38。
*/

func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)
	{
		rankOfLJ, fatRateOfLJ := getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静的排名是第一，实际排名是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是0.38，实际是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强的排名是第二，实际排名是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率是0.38，实际是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, fatRateOfZW := getRank("张伟")
		if rankOfZW != 2 {
			t.Fatalf("预期张伟的排名是第二，实际排名是：%d", rankOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期张伟的体脂率是0.38，实际是：%f", fatRateOfZW)
		}
	}
}

/*
案例3：
- 王强录入体脂是38
- 李静录入体脂是28
- 张伟注册成功，未录入体脂

李静排名第一，体脂28；王强排名第二，体脂28；张伟排名第三，没有体脂记录。
*/

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("李静", 0.28)
	inputRecord("张伟")
	{
		rankOfLJ, fatRateOfLJ := getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静的排名是第一，实际排名是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是0.38，实际是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强的排名是第二，实际排名是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率是0.38，实际是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, _ := getRank("张伟")
		if rankOfZW != 3 {
			t.Fatalf("预期张伟的排名是第三，实际排名是：%d", rankOfZW)
		}
	}
}
