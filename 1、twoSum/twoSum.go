
func twoSum(nums []int, target int) []int {
	var l = len(nums) - 1
	var m = int(l / 2)

	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int)

	for i := 0; i <= m; i++ {

		b := nums[i]
		a := nums[l-i]

		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}
		index[b] = i

		if i != l-i {
			if a+b == target {
				return []int{i, l - i}
			}

			// 通过查询map，获取a = target - b的序列号
			if j, ok := index[target-a]; ok {
				// ok 为 true
				// 说明在i之前，存在 nums[j] == a
				return []int{j, l - i}
				// 注意，顺序是j，i
			}

			index[a] = l - i

		}

		// 把b和i的值，存入map
	}

	return nil
}

