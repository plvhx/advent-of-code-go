package aoc2020

func ReportRepair(arr []uint32) uint32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == 2020 {
				return arr[i] * arr[j]
			}
		}
	}

	return 0
}
