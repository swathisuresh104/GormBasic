package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*The DB is exptected to be created - Company
CREATE SCHEMA IF NOT EXISTS `Company` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci ;
USE `Company` ;
*/

/*type Employee struct {
	empNo     int    `json:"id"`
	FirstName string `json:"FName"`
	DeptId    int    `json:"DeptId"`
}*/

type Employee struct {
	gorm.Model
	empNo     int    `gorm:"primaryKey"`
	FirstName string `json:"FName"`
	DeptId    int    `json:"DeptId"`
}

//gorm.Model adds the updated_at, created_at, deleted_at tables

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3377)/Company?charset=utf8&parseTime=True")
	fmt.Println("db:", db)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Database connection established")
	}

	//Empty table Employee creation.
	db.AutoMigrate(&Employee{})

	//Employee table row insert if row does not exist
	emp := Employee{FirstName: "Jacob", DeptId: 202}
	res := db.FirstOrCreate(&emp)
	if res.Error != nil {
		fmt.Println("Error while creating the employee")
	} else {
		fmt.Println("Successfully created the employee")
	}

	//Searching for Jacob and
	updatedEmpOnName := Employee{FirstName: "George", DeptId: 302}
	res = db.Model(&Employee{}).Where(Employee{FirstName: "Jacob"}).Update(&updatedEmpOnName)
	if res.Error != nil {
		fmt.Println("Error while updating the employee based on FirstName")
	} else {
		fmt.Println("Successfully updated the employee based on FirstName")
	}

	updateEmpOnDeptId := Employee{FirstName: "John", DeptId: 402}
	res = db.Model(&Employee{}).Where(Employee{DeptId: 302}).Update(&updateEmpOnDeptId)
	if res.Error != nil {
		fmt.Println("Error while updating the employee based on DeptId")
	} else {
		fmt.Println("Successfully updated the employee based on DeptId")
	}

	//delete employee based on FirstName - deleted_at gets updated.
	var rowsToDelete int
	res = db.Model(&Employee{}).Where(Employee{FirstName: "John"}).Count(&rowsToDelete)
	if rowsToDelete > 0 {
		res = db.Model(&Employee{}).Where(Employee{FirstName: "John"}).Delete(emp)
		if res.Error != nil {
			fmt.Println("Error while deleting the employee")
		} else {
			fmt.Println("Successfully deleted the employee")
		}
	}

	//HardDelete employee based on FirstName
	res = db.Model(&Employee{}).Where(Employee{FirstName: "George"}).Unscoped().Delete(emp)
	if res.Error != nil {
		fmt.Println("Error while hard deleting the employee")
	} else {
		fmt.Println("Successfully hard deleted the employee")
	}
}
