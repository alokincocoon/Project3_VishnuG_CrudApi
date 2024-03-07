package models

import (
	"context"
	"encoding/json"

	"github.com/edgedb/edgedb-go"
)

type Phone struct {
	Type   string `edgedb:"type" json:"type"`
	Number string `edgedb:"number" json:"number"`
}

type Student struct {
	ID     edgedb.UUID              `edgedb:"id" json:id`
	Name   string                   `edgedb:"name" json:"name"`
	Email  string                   `edgedb:"email" json:"email"`
	Phones []map[string]interface{} `edgedb:"phones" json:"phones"`
}

func GetStudents(db *edgedb.Client) ([]Student, error) {
	ctx := context.Background()
	var students []Student
	query := "SELECT Student{id, name, email, phones}"
	err := db.Query(ctx, query, &students)
	if err != nil {
		return nil, err
	}
	return students, err
}

func GetStudentDetails(db *edgedb.Client, uuid edgedb.UUID) (Student, error) {
	var student Student
	ctx := context.Background()
	query := "SELECT Student{id, name, email, phones} filter .id = <uuid>$0"
	err := db.QuerySingle(ctx, query, &student, uuid)
	if err != nil {
		return student, err
	}
	return student, nil
}

func UpdateStudentDetails(db *edgedb.Client, student *Student) error {
	ctx := context.Background()
	query := `
		UPDATE Student
		FILTER .id = <uuid>$0
		SET {
			name := <str>$1,
			email := <str>$2,
			phones := <array<json>>$3
		}
	`
	var result Student
	var phones []interface{}

	for _, phone := range student.Phones {
		phoneEncoded, _ := json.Marshal(phone)
		phones = append(phones, phoneEncoded)
	}
	if err := db.QuerySingle(ctx, query, &result, student.ID, student.Name, student.Email, phones); err != nil {
		return err
	}

	return nil
}

func CreateStudent(db *edgedb.Client, student *Student) (edgedb.UUID, error) {
	ctx := context.Background()
	query := `
		INSERT Student {
			name := <str>$0,
			email := <str>$1,
			phones := <array<json>>$2
		}
	`
	var out struct{ id edgedb.UUID }
	var phones []interface{}

	for _, phone := range student.Phones {
		phoneEncoded, _ := json.Marshal(phone)
		phones = append(phones, phoneEncoded)
	}
	err := db.QuerySingle(ctx, query, &out, student.Name, student.Email, phones)
	if err != nil {
		return edgedb.UUID{}, err
	}
	return out.id, nil
}

func DeleteStudent(db *edgedb.Client, id edgedb.UUID) (interface{}, error) {
	ctx := context.Background()

	var student Student
	err := db.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
		query := `
			DELETE Student
			filter .id = <uuid>$0
		`
		if e := tx.QuerySingle(ctx, query, &student, id); e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return student, nil
}
