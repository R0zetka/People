package service

import (
	"TestZD/internal/Person"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	client *pgxpool.Pool
}

func NewRepository(client *pgxpool.Pool) *repository {
	return &repository{
		client: client,
	}
}
func (r *repository) Create(ctx context.Context, person *Person.Person) error {
	// Добавление нового человека
	q := "INSERT INTO person (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	if err := r.client.QueryRow(ctx, q, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality).Scan(&person.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			return nil
		}
		return err
	}
	return nil
}

func (r *repository) People(ctx context.Context) (m []Person.Person, err error) {
	// Показать всех людей
	q := "SELECT id, name, surname, patronymic, age, gender, nationality FROM person"
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err

	}
	people := make([]Person.Person, 0)
	for rows.Next() {
		var p Person.Person

		err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
		if err != nil {
			return nil, err

		}
		people = append(people, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err

	}
	fmt.Println(people)
	return people, nil

}

func (r *repository) OnePerson(ctx context.Context, id string) (Person.Person, error) {
	// Показать одного человека по id
	q := "SELECT id, name, surname, patronymic, age, gender, nationality FROM person WHERE id = $1"
	var p Person.Person
	err := r.client.QueryRow(ctx, q, id).Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
		}
		return Person.Person{}, err
	}
	return p, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := "DELETE FROM person WHERE id =$1"
	var p Person.Person
	err := r.client.QueryRow(ctx, q, id).Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
		}
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, person Person.Person) error {
	fmt.Println("Update")
	return nil
}
