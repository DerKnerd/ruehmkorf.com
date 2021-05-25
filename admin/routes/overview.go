package routes

type OverviewModel struct {
	Items      interface{}
	Count      int
	Offset     int
	Page       int
	TotalPages int
	TotalCount int
	NextOffset int
	PrevOffset int
}
