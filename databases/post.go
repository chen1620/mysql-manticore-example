package databases

import (
	"context"
	"fmt"
	"github.com/manticoresoftware/go-sdk/manticore"
	"gorm.io/gorm"

	"mysql-manticore-example/models"
)

// PostRepository database structure.
type PostRepository interface {
	AddPost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, id int) (*models.Post, error)
}

// postRepo post repository implement struct.
type postRepo struct {
	mysqlDB         *gorm.DB
	manticoreClient *manticore.Client
}

// NewPostRepository ...
func NewPostRepository(mysqlDB *gorm.DB, manticoreClient *manticore.Client) PostRepository {
	return &postRepo{
		mysqlDB:         mysqlDB,
		manticoreClient: manticoreClient,
	}
}

// AddPost add new post.
// firstly, insert post to mysql database in a transaction.
// secondly, insert post to manticore-search database, if it succeeds, commit mysql db transaction else rollback.
func (r *postRepo) AddPost(ctx context.Context, p *models.Post) error {
	tx := r.mysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert data into Manticore search
	cmd := fmt.Sprintf(
		"replace into posts (id, title, content) values (%d, '%s', '%s')",
		p.ID, p.Title, p.Content,
	)
	res, err := r.manticoreClient.Sphinxql(cmd)
	fmt.Println(res, err)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

// GetPostByID ...
func (r *postRepo) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	var p models.Post
	err := r.mysqlDB.Preload("User").First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
