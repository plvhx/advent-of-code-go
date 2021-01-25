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

func ReportRepair_Part2(arr []uint32) uint32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					return arr[i] * arr[j] * arr[k]
				}
			}
		}
	}

	return 0
}
