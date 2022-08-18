# golang_trapping_rain_water

Given `n` non-negative integers representing an elevation map where the width of each bar is `1`
, compute how much water it can trap after raining.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

```

**Example 2:**

```
Input: height = [4,2,0,3,2,5]
Output: 9

```

**Constraints:**

- `n == height.length`
- `1 <= n <= 2 * $10^4$`
- `0 <= height[i] <= $10^5$`

## 解析

給定一個非負整數陣列 height 

每個height[i] 代表在 i 位置的高度

假設每個 height[i] 寬度是 1 單位長度

如果要在位置 i 累積雨水

需要有存在  0 ≤ k < i , i < j ≤ len(height) - 1 使得 min (height[j], height[k]) > height[i]

且累積的雨水面積 =  min (height[j], height[k]) - height[i]

![](https://i.imgur.com/QrMBGDN.png)

要求寫一個演算法算出在給定的 height 中能累積的雨水面積

由雨水累積的條件可以看出

對於每一個位置 i 所能累積的雨水 

必須從位置 0 到 i-1 找出一個  k , 使得 height[k]> height[i] 且 是 height[k]從 0 到 i-1 的最大值 leftMax

必須從位置 i+1 到 len(height)-1 找出一個  j , 使得 height[j]> height[i] 且 是 height[j]從 i+1

 到 len(height) -1 的 最大值 rightMax

該位置累積的雨水= min(leftMax, rightMax) - height[i]


![](https://i.imgur.com/CE7pcVN.png)

對於這樣的特性可以使用左右兩個 pointer 來做找尋每個當下的 leftMax, rightMax

並且在遇到有符合 min(leftMax,rightMax) >  height[i]累加

具體作法如下

初始化 lp =0, rp = len(height) - 1, leftMax =0 , rightMax = 0 , result = 0

當 lp < rp 時做以下運算

  當  height[rp] > height[lp] 代表 左界比較少

     檢查 height[lp] > leftMax  決定是否需要累積面積 

     當 height[lp] > leftMax 時 代表不需要累積水位 更新 leftMax =  height[lp]

     否則 更新 result += leftMax - height[lp] (累積 lp 水位)

     更新 lp += 1

當  height[rp] ≤ height[lp] 代表 右界比較少

     檢查 height[rp] > rightMax  決定是否需要累積面積 

     當 height[rp] > rightMax 時 代表不需要累積水位 更新 rightMax =  height[rp]

     否則 更新 result += rightMax - height[rp] (累積 rp 水位)

     更新 rp -= 1

回傳 result

![](https://i.imgur.com/a5Slape.png)

## 程式碼
```go
package sol

func trap(height []int) int {
	leftMax, rightMax, result := 0, 0, 0
	lp, rp := 0, len(height)-1
	for lp < rp {
		if height[lp] > height[rp] {
			if height[rp] > rightMax {
				rightMax = height[rp]
			} else {
				result += rightMax - height[rp]
			}
			rp--
		} else {
			if height[lp] > leftMax {
				leftMax = height[lp]
			} else {
				result += leftMax - height[lp]
			}
			lp++
		}
	}
	return result
}

```
## 困難點

1. 要思考出能夠累積雨水的條件
2. 要思考出利用 two pointer 來找出每個位置所能累積雨水的關係
3. 要思考出如何做最佳的移動方式，因為只有較短的 height 會影響累積雨水的狀況所以先移動較短的部份

## Solve Point

- [x]  初始化 lp =0, rp = len(height) - 1, leftMax = 0, rightMax = 0, result = 0
- [x]  當 lp < rp 時 做以下計算
- [x]  當 height[lp] > height[rp] 做以下運算 (rp 是比較低的部份 有機會形成 雨水累積)
- [x]  當 height[rp] > rightMax 時， 更新 rightMax = height[rp] (無法累積雨水) 否則更新 result += rightMax - height[rp]
- [x]  更新 rp -= 1
- [x]  當 height[lp] ≤ height[rp] 做以下運算 (lp 是比較低的部份 有機會形成 雨水累積)
- [x]  當 height[lp] > leftMax 時， 更新 leftMax = height[lp] (無法累積雨水) , 否則更新 result += leftMax - height[lp]
- [x]  更新 lp += 1
- [x]  回傳 result