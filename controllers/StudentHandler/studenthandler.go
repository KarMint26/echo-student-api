package studenthandler

import (
	"net/http"

	"github.com/KarMint26/echo-student-api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Define Struct With Gorm Database
type Database struct {
	DB *gorm.DB
}

// Define and Make Function that get all students data in the postgresql database
func (d *Database) GetStudents(c echo.Context) error {
	students := &[]models.Students{}

	if err := d.DB.Find(students).Error; err != nil {
        c.JSON(http.StatusNotFound, echo.Map{"message":"Can't find students"})
        return err
    }

	c.JSON(http.StatusOK, echo.Map{"data": students, "message": "Successfully Get All Students"})
	return nil
}

// Define and Make Function that get spesific student data in the postgresql database
func (d *Database) GetStudentById(c echo.Context) error {
	studentModel := &models.Students{}
	id := c.Param("id")
	
	if id == ""{
		c.JSON(http.StatusNotFound, echo.Map{"message":"Can't find id on the database"})
		return nil
	}

	if err := d.DB.First(studentModel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, echo.Map{"message":"Can't find student data on the database"})
        return err
	}

	c.JSON(http.StatusOK, echo.Map{"data":studentModel, "message": "Successfully Get Student by Id"})
	return nil
}

// Define and Make Function that input or create new student data on the postgresql database
func (d *Database) CreateStudent(c echo.Context) error {
	studentModel := &models.Students{}

	if err := c.Bind(studentModel); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message":"Bad request from the client"})
		return err
	}
	
	if err := d.DB.Create(studentModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"message":"Can't create student data"})
		return err
	}

	c.JSON(http.StatusOK, echo.Map{"data":studentModel, "message":"Successfully added data"})
	
	return nil
}

// Define and Make Function that updated student data on the postgresql database
func (d *Database) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	studentModel := &models.Students{}
	responseStudent := &models.Students{}

	if err := c.Bind(studentModel); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message":"Can't fetched data from the client request"})
		return err
	}

	if studentModel.ID != 0 {
		responseStudent.ID = studentModel.ID
	}

	if studentModel.Name != "" {
		responseStudent.Name = studentModel.Name
	}

	if studentModel.Age != 0 {
		responseStudent.Age = studentModel.Age
	}

	if studentModel.Grade != "" {
		responseStudent.Grade = studentModel.Grade
	}
	
	if d.DB.Model(studentModel).Where("id = ?", id).Updates(studentModel).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, echo.Map{"message":"Can't find student id on the database"})
		return nil
	}

	c.JSON(http.StatusOK, echo.Map{"new_data":responseStudent, "message":"Successfully updated data"})
	
	return nil
}

// Define and Make Function that deleted spesific student data by id on the postgresql database
func (d *Database) DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	studentModel := &models.Students{}

	if id == ""{
		c.JSON(http.StatusNotFound, echo.Map{"message":"Can't find id on the database"})
		return nil
	}

	if err := d.DB.Delete(studentModel, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message":"Failed to deteled data"})
		return err
	}

	c.JSON(http.StatusOK, echo.Map{"message":"Successfully deleted data"})
	
	return nil
}