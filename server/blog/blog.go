package blog

import (
	"context"
	"fmt"
	"time"

	"example.com/model"

	"github.com/jackc/pgx/v5"
)

func BlogsByPage(pageSize int, pageNumber int) ([]model.Blog, error) {
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}
	offset := (pageNumber - 1) * pageSize
	var blogs []model.Blog
	rows, err := model.DB.Query(context.Background(), "SELECT * FROM blogs ORDER BY blog_id LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("blogByPage: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(
			&blog.Blog_id,
			&blog.User_id,
			&blog.Title,
			&blog.Content,
			&blog.Create_time,
			&blog.Update_time)
		if err != nil {
			return nil, fmt.Errorf("blogByPage: %v", err)
		}
		name_err := model.DB.QueryRow(context.Background(), "SELECT name FROM users WHERE user_id = $1", blog.User_id).Scan(&blog.User)
		if name_err != nil {
			blog.User = "Not Found"
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("blogByPage: %v", err)
	}
	return blogs, nil
}

func BlogsByUser(name string) ([]model.Blog, error) {
	var blogs []model.Blog
	var user_id int64
	err := model.DB.QueryRow(context.Background(), "SELECT user_id FROM users WHERE name = $1", name).Scan(&user_id)
	if err != nil {
		return nil, fmt.Errorf("searchUser %q: %v", name, err)
	}
	rows, err := model.DB.Query(context.Background(), "SELECT * FROM blogs WHERE user_id = $1", user_id)
	if err != nil {
		return nil, fmt.Errorf("blogsByUser %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(
			&blog.Blog_id,
			&blog.User_id,
			&blog.Title,
			&blog.Content,
			&blog.Create_time,
			&blog.Update_time)
		if err != nil {
			return nil, fmt.Errorf("blogsByUser %q: %v", name, err)
		}
		blog.User = name
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("blogsByUser %q: %v", name, err)
	}
	return blogs, nil
}

func BlogsByID(id int64) (model.Blog, error) {
	var blog model.Blog
	err := model.DB.QueryRow(context.Background(), "SELECT * FROM blogs WHERE blog_id = $1", id).Scan(
		&blog.Blog_id,
		&blog.User_id,
		&blog.Title,
		&blog.Content,
		&blog.Create_time,
		&blog.Update_time)
	name_err := model.DB.QueryRow(context.Background(), "SELECT name FROM users WHERE user_id = $1", blog.User_id).Scan(&blog.User)
	if name_err != nil {
		blog.User = "Not Found"
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return blog, fmt.Errorf("blogByID %d: no such blog", id)
		}
		return blog, fmt.Errorf("blogByID %d: %v", id, err)
	}

	return blog, nil
}

func AddBlog(blog *model.Blog) (int64, error) {
	var newID int64
	err := model.DB.QueryRow(context.Background(), "SELECT blog_id FROM blogs WHERE title = $1 AND user_id = $2", blog.Title, blog.User_id).Scan(&newID)
	if err == nil {
		return 0, fmt.Errorf("addBlog %+v: Blog already exists", blog)
	}
	now_time := time.Now().UTC()
	err = model.DB.QueryRow(context.Background(), `
	INSERT INTO blogs
		(user_id, title, content, create_time, update_time)
	VALUES 
		($1, $2, $3, $4, $5) 
	RETURNING blog_id`, blog.User_id, blog.Title, blog.Content, now_time, now_time).Scan(&newID)
	if err != nil {
		return 0, fmt.Errorf("addAlbum %+v: %v", blog, err)
	}
	return newID, nil
}
