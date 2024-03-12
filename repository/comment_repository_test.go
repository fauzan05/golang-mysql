package repository

import (
	"context"
	"fmt"
	golangmysql "golang-mysql"
	"golang-mysql/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangmysql.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "testrepo@mail.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err) 
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangmysql.GetConnection())
	ctx := context.Background()
	
	comment, err := commentRepository.FindById(ctx, 17)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangmysql.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("Id : ", comment.Id, " , Email : ", comment.Email, " , Comment : ", comment.Comment)
	}
}