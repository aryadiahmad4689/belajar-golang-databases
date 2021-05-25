package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryImplement struct {
	DB *sql.DB
}

func newCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImplement{DB: db}
}

func (repo *commentRepositoryImplement) Delete(ctx context.Context, id int32) error {
	script := "delete from comments where id=?"
	_, err := repo.DB.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}

	return errors.New("Berhasil Hapus Data")
}

func (repo *commentRepositoryImplement) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

	script := "Insert into comments(comment) values(?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Comment)

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

func (repo *commentRepositoryImplement) FindId(ctx context.Context, id int32) (entity.Comment, error) {

	script := "Select id, comment from comments where id=? limit 1"

	result, err := repo.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer result.Close()
	if result.Next() {

		result.Scan(&comment.Id, &comment.Comment)
		return comment, nil

	} else {
		return comment, errors.New("Id: " + strconv.Itoa(int(id)) + "tidak ditemukan")
	}

}

func (repo *commentRepositoryImplement) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "Select id ,comment from comments"

	result, err := repo.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer result.Close()
	var comments []entity.Comment
	for result.Next() {
		comment := entity.Comment{}

		result.Scan(&comment.Id, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil

}
