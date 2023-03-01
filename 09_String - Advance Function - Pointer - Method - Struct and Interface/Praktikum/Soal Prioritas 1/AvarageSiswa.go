package main

import "fmt"

type Student struct {
	Name  []string
	Score []int
}

func (s Student) average() int {
	var total float64
	for _, v := range s.Score {
		total += float64(v)
	}
	return int(total) / len(s.Score)
}

func (s Student) min() (string, int) {
	var min int = s.Score[0]
	var name string = s.Name[0]

	for i, v := range s.Score {
		if v < min {
			min = v
			name = s.Name[i]
		}
	}
	return name, min
}

func (s Student) max() (string, int) {
	var max int = s.Score[0]
	var name string = s.Name[0]

	for i, v := range s.Score {
		if v > max {
			max = v
			name = s.Name[i]
		}
	}
	return name, max
}

func main() {
	result := Student{
		Name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopam"},
		Score: []int{80, 75, 70, 75, 60},
	}

	minName, minScore := result.min()
	maxName, maxScore := result.max()
	fmt.Println("Nilai rata-rata :", result.average())
	fmt.Printf("Nilai terendah : %s (%d)\n", minName, minScore)
	fmt.Printf("Nilai tertinggi : %s (%d)\n", maxName, maxScore)
}
