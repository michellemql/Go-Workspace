package main

import "fmt"

type Student struct {
	name string
	grade []int
	age int
}
func (s Student) getAge() int {
	return s.age
} 

// Need to add * whenever you need to modify the object itself
func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) setAge_2(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float32(sum) / float32(len(s.grade))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grade {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 85}, 19}
	fmt.Println(s1)  //{Tim [70 90 85] 19} 
	s1.getAge()
	s1.setAge(7)
	fmt.Println(s1)  //{Tim [70 90 85] 7} 
	s1.setAge_2(10)
	fmt.Println(s1)   // {Tim [70 90 85] 7} because without * it won't change

	s1 := Student{"Tim", []int{70, 90, 85}, 19}
	fmt.Println(s1)  //{Tim [70 90 85] 19} 
	s1.getAge()
	s1.setAge(7)
	fmt.Println(s1)  //{Tim [70 90 85] 7} 

	average := s1.getAverageGrade()
	fmt.Println(average) // 81.25
	s2 := Student{"Joe", []int{70, 90, 85, 95, 90, 90}, 19}
	average2 := s2.getAverageGrad3()
	fmt.Println(average, average2) // 81.25 85.71

	max := s1.getMaxGrade()
	fmt.Println(max)  // 90
}