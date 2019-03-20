package levenshtein

import (
	"math"
)

//DamereauLevenshtein : Compute the distance beetween two words
func DamereauLevenshtein(frstString string, scndString string) float64 {
	lenFrst := len(frstString)
	lenScnd := len(scndString)
	matrix := make([][]float64, lenFrst+1)
	for k := 0; k < lenFrst; k++ {
		matrix[k] = make([]float64, lenScnd+1)
		if k == 0 {
			for j := 0; j < lenScnd; j++ {
				matrix[k][j] = float64(j)
			}
		} else {
			matrix[k][0] = float64(k)
		}
	}
	cost := 0.0
	for i := 1; i < lenFrst; i++ {
		for j := 1; j < lenScnd; j++ {
			if frstString[i] == scndString[j] {
				cost = 0
			} else {
				cost = 1
			}
			newValue := math.Min(matrix[i-1][j]+1, matrix[i][j-1]+1)
			matrix[i][j] = math.Min(newValue, matrix[i-1][j-1]+cost)
			if i > 1 && j > 1 && frstString[i] == scndString[j-1] && frstString[i-1] == scndString[j] {
				matrix[i][j] = math.Min(matrix[i][j], matrix[i-2][j-2])
			}
		}
	}
	// for i := 0; i < lenFrst; i++ {
	// 	for j := 0; j < lenScnd; j++ {
	// 		fmt.Printf(" %f ", matrix[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }
	// fmt.Print(matrix[lenFrst-1][lenScnd-1])
	return matrix[lenFrst-1][lenScnd-1]
}
