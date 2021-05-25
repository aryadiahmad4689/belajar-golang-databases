package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	commentRepo := newCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Comment: "Test Insert Repo",
	}
	result, err := commentRepo.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindId(t *testing.T) {
	commentRepo := newCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	result, err := commentRepo.FindId(ctx, 79)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindAll(t *testing.T) {
	commentRepo := newCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	result, err := commentRepo.FindAll(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestDelete(t *testing.T) {
	commentRepo := newCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	err := commentRepo.Delete(ctx, 10)

	fmt.Println(err)

}
