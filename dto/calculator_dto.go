package dto

type Numbers struct{
	Num1 int `binding:"required" json:"num_1"`
	Num2 int `binding:"required" json:"num_2"`
}

type Result struct{
	Value int `binding:"required" json:"value"`
}