// 二维数组  1是陆地 0是海   1,1  一个陆地
// 0,0,0,0,0,0
// 0,0,0,0,0,0
// 0,0,1,0,0,0
// 0,0,0,1,1,0
// 0,0,0,0,0,0
package algo

func Process(arrs [5][6]int) int {
	var num int
	for _, arr := range arrs {
		for index, value := range arr {
			if value == 0 {
				continue
			}
			if value == 1 {
				if index == 0 {
					num = num + 1
				}
				if index > 0 {
					if arr[index-1] == 1 || arr[index+1] == 1 {
						num = num + 1
					}
				}
			}
		}
	}
	return num
}
