package main

type FooService struct {
	*DBService
}

func (s *FooService) GetTblName() string { return s.TblName }
