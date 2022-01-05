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
			t.Fatalf("预期王李静的排名是第一，实际排名是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期王强的体脂率是0.28，实际是：%f", fatRateOfLJ)
		}
	}

}

func TestCase2(t *testing.T) {

}
