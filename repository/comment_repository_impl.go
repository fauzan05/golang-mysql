package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments(email,comment) VALUES(?,?)"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	defer repository.DB.Close()
	defer stmt.Close()
	if err != nil {
		return comment, err
	}
	result, err := stmt.ExecContext(ctx, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comment, error) {
	query := "SELECT * FROM comments WHERE id = ? LIMIT 1"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	comment := entity.Comment{}
	defer repository.DB.Close()
	defer stmt.Close()
	if err != nil {
		return comment, err
	}
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		comment_id := int(id)
		return entity.Comment{}, errors.New("Komentar dengan id " + strconv.Itoa(comment_id) + " tidak ditemukan!")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT * FROM comments"
	defer repository.DB.Close()
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		// memasukkan datanya satu per satu ke variabel comments
		comments = append(comments, comment)

	}
	return comments, nil
}
