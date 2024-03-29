package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test_db"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Employee struct {
	ID int `sql:"id"`
	Name string `sql:"name"`
	age int `sql:"age"`
	address string `sql:"address"`
	salary float32 `sql:"salary"`
}

type EmployeeDepartment struct {
	ID int `sql:"ID"`
	Name string `sql:"Name"`
	Department string `sql:"Department"`
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	// create Employees table
	// _, err = db.Exec(`
	// 	CREATE TABLE Employees (
	// 		ID INT PRIMARY KEY,
	// 		NAME VARCHAR(255),
	// 		AGE INT,
	// 		ADDRESS VARCHAR(255),
	// 		SALARY INT
	// 	)
	// `)

	// update salary data type
	// _, err = db.Exec(`
	// 	ALTER TABLE Employees
	// 	ALTER COLUMN SALARY TYPE FLOAT
	// `)

	// insert employees data
	// _, err = db.Exec(`
	// 	INSERT INTO Employees
	// 	VALUES
	// 	(1, 'Rizki', 25, 'Jl. Kebon Jeruk', 2000000),
	// 	(2, 'Andi', 27, 'Jl. Kebon Sirih', 3000000),
	// 	(3, 'Budi', 30, 'Jl. Kebon Melati', 4000000),
	// 	(4, 'Caca', 32, 'Jl. Kebon Anggrek', 5000000),
	// 	(5, 'Deni', 35, 'Jl. Kebon Mawar', 6000000)
	// `)

	// scan employees data
	// rows, err := db.Query("SELECT * FROM Employees")
	// if err != nil {
	// 	panic(err)
	// }

	// var employees []Employee

	// for rows.Next() {
	// 	var employee Employee 

	// 	err := rows.Scan(&employee.ID, &employee.Name, &employee.age, &employee.address, &employee.salary)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	employees = append(employees, employee)
	// }

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(employees)

	// scan only one employee data
	// row := db.QueryRow("SELECT * FROM Employees WHERE ID = 1")
	// if err != nil {
	// 	panic(err)
	// }

	// var employee Employee 

	// err = row.Scan(&employee.ID,&employee.Name, &employee.age, &employee.address, &employee.salary)
	// if errors.Is(err, sql.ErrNoRows) {
	// 	fmt.Println("Data tidak diterima")
	// 	return
	// } else if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(employee)

	// inner join
	rows, err := db.Query(`
		SELECT Employees.ID AS "ID", Employees.NAME AS "Name", Departments.NAME AS "Department"
		FROM Employees INNER JOIN Departments
		ON Employees.id_department = Departments.ID
		ORDER BY Employees.ID ASC
	`)
	if err != nil {
		panic(err)
	}

	var employeeDepartments []EmployeeDepartment

	for rows.Next() {
		var employeeDepartment EmployeeDepartment

		err := rows.Scan(&employeeDepartment.ID, &employeeDepartment.Name, &employeeDepartment.Department)
		if err != nil {
			panic(err)
		}

		employeeDepartments = append(employeeDepartments, employeeDepartment)
	}

	fmt.Println(employeeDepartments)
}