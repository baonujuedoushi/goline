package main

import "fmt"

func bag() {
	//bag weight  = 15
	//item number = 5
	//item values = [0=>4, 1=>5, 2=>10, 3=>11, 4=>13]
	//itemWeight  = [0=>3, 1=>4, 2=>7, 3=>8, 4=>9]
	//w 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
	//0 0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0
	//1 0  0  0  4  4  4  4  4  4  4  4  4  4  4  4  4
	//2 0  0  0  0  5  5  5  9  9  9  9  9  9  9  9  9
	//3 0  0  0  0  0  0  0 10 10 10 10 15 15 15 19 19
	//4 0  0  0  0  0  0  0  0 11 11 11 15 15 15 19 21
	//5 0  0  0  0  0  0  0  0  0 13 13 15 15 15 19 21
	bagWeight := 15
	itemValues := []int{4, 5, 10, 11, 13}
	itemWeight := []int{3, 4, 7, 8, 9}
	itemNum := len(itemWeight)
	dp := make([][]int, itemNum+1)
	for k := range dp {
		dp[k] = make([]int, bagWeight+1)
	}
	//不放入任何物品时价值总是为0
	for v := 0; v <= bagWeight; v++ {
		dp[0][v] = 0
	}
	for i := 1; i <= itemNum; i++ {
		//任何物品在背包载重为0时，值都为0（无法放入）
		dp[i][0] = 0
		for currentWeight := 1; currentWeight <= bagWeight; currentWeight++ {
			if itemWeight[i-1] <= currentWeight {
				dropValue := dp[i-1][currentWeight]
				//观察当前背包重量 - 物品重量时的背包容量能放下的价值，加上当前物品的价值（选取新的物品时），能否超过丢弃它的情况
				pickValue := dp[i-1][currentWeight-itemWeight[i-1]] + itemValues[i-1]
				dp[i][currentWeight] = max(dropValue, pickValue) //选取考虑0-i物品时最佳方案
			}
		}
	}
	j := bagWeight
	selectedItems := []int{}
	for i := itemNum; i > 0; i-- {
		// 如果值不等于上一行，说明拿了
		if dp[i][j] != dp[i-1][j] {
			selectedItems = append(selectedItems, i-1) // 记录下标
			j -= itemWeight[i-1]                       // 减去重量，继续找剩下的
		}
	}
	fmt.Printf("选中的物品下标为: %v\n", selectedItems)
}
