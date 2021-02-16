package service

type SortField string

const (
	DateField   SortField = "date"
	ViewsField  SortField = "views"
	ClicksField SortField = "clicks"
	CostField   SortField = "cost"
	CpcField    SortField = "cpc"
	CpmField    SortField = "cpm"
)

type SortOrder string

const (
	AscOrder  SortOrder = "asc"
	DescOrder SortOrder = "desc"
)

const LayoutISO = "2006-01-02"
