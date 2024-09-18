package lib1

type Student struct {
	Person
	Rank    int
	ClassNo int
}

func (s *Student) GetRank() int {
	return s.Rank
}

func (s *Student) SetRank(rank int) {
	s.Rank = rank
}
