package entity

type Person struct{
	Firstname string `json:"firstname"  binding:"required"`
	Lastname string `json:"lastname"  binding:"required"`
	Age int8 `json:"age"  binding:"gte=1,lte=130"`
	Email string `json:"email"  binding:"required,email"`

}
type Employe struct {
	Name     string `json:"name" binding:"min=2,max=10" validate:"is-cool"`
	Position string `json:"position"  binding:"max=20"`
	Contract string `json:"contract" binding:"required"`
	Author Person `json:"author" binding:"required"`
}
