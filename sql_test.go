package golangmysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
	// "time"
)

func TestCreateUser(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	var query string = "INSERT INTO USERS(id,name) VALUES('1', 'Fauzan Nur Hidayat')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Inserting Data User")
}

func TestGetUser(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM users"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id : ", id, ", Name : ", name)
		// fmt.Println("Name : ", name)
	}
}

func TestCreateUserNew(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// var query1 string = "INSERT INTO USERS(id,name,email,balance,rating,birth_date,married) VALUES('1', 'Fauzan Nur Hidayat', 'fauzannurhidayat8@gmail.com', 10000, 4.5, '2001-02-05', false)"
	// var query2 string = "INSERT INTO USERS(id,name,email,balance,rating,birth_date,married) VALUES('1', 'Fauzan Nur Hidayat', 'fauzannurhidayat8@gmail.com', 10000, 4.5, '2001-02-05', false)"
	for i := 1; i <= 5; i++ {
		query := "INSERT INTO USERS(id,name,email,balance,rating,birth_date,married) VALUES('" + strconv.Itoa(i) + "','" + "user " + strconv.Itoa(i) + "', 'user" + strconv.Itoa(i) + "@mail.com', " + "1000" + strconv.Itoa(i) + ", 4.5, '2001-02-05', false)"
		_, err := db.ExecContext(ctx, query)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success Inserting Data User")
}

func TestGetUserNew(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM users"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var email sql.NullString // bisa menerima data null
		var balance int
		var rating float32
		var created_at, birth_date time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		var filter_email string
		if email.Valid {
			filter_email = email.String
		}
		fmt.Println("Id : ", id, ", Name : ", name, ", Email : ", filter_email, ", Balance : ", balance, ", Rating : ", rating, ", Created_At : ", created_at, ", Birth_Date : ", birth_date, ", Is_Married : ", married)
		// fmt.Println("Name : ", name)
	}
}

func TestGetUserWithVariadicArguments(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "SELECT * FROM users WHERE id = ?"
	rows, err := db.QueryContext(ctx, query, "3")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var email sql.NullString // bisa menerima data null
		var balance int
		var rating float32
		var created_at, birth_date time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		var filter_email string
		if email.Valid {
			filter_email = email.String
		}
		fmt.Println("Id : ", id, ", Name : ", name, ", Email : ", filter_email, ", Balance : ", balance, ", Rating : ", rating, ", Created_At : ", created_at, ", Birth_Date : ", birth_date, ", Is_Married : ", married)
		// fmt.Println("Name : ", name)
	}
}

func TestCreateCommentWithLastInsertId(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	email := "fauzannurhidayat8@gmail.com"
	comment := "Hello World"
	query := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	getId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Id yang terakhir diinput adalah : ", getId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 1; i <= 5; i++ {
		email := "email" + strconv.Itoa(i) + "@mail.com"
		comment := "comment " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Komentar id yang berhasil ter-insert : ", id)
	}
}

func TestWithTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := tx.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 5; i++ {
		
		if i == 3 {
			tx.Rollback()
		}
		email := "email" + strconv.Itoa(i) + "@mail.com"
		comment := "comment " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Komentar id yang berhasil ter-insert : ", id)
	}

	tx.Commit()
}
