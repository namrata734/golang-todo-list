package controller

import "golang-session/models"

func InitializingAndAddingToArray(Students *[]models.Student) {
	student1 := models.Student{
		Id:      1,
		Email:   "vaibhav@gamil.com",
		Name:    "Vaibhav",
		Age:     21,
		PhoneNo: 9900000650,
	}
	student2 := models.Student{
		Id:      2,
		Email:   "ankit@gmail.com",
		Name:    "Ankit",
		Age:     21,
		PhoneNo: 9900000372,
	}
	student3 := models.Student{
		Id:      3,
		Email:   "tushar@gamil.com",
		Name:    "Tushar",
		Age:     21,
		PhoneNo: 9900006372,
	}

	*Students = append(*Students, student1, student2, student3)
}
