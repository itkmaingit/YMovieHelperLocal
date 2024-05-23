package utils

// 二次元スライスの文字列の縦横を入れ替える関数
func Transpose(input [][]string) [][]string {
	rowNum := len(input)
	if rowNum == 0 {
		return input
	}
	colNum := len(input[0])

	output := make([][]string, colNum)
	for i := range output {
		output[i] = make([]string, rowNum)
	}

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			output[j][i] = input[i][j]
		}
	}

	return output
}
