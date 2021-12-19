package repository

/*
import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"github.com/AleksK1NG/api-mc/internal/models"
)

func TestNewsRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	priceGroupRepo := NewPriceGroupRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		authorUID := uuid.New()
		title := "title"
		content := "content"

		rows := sqlmock.NewRows([]string{"author_id", "title", "content"}).AddRow(authorUID, title, content)

		priceGroup := &models.News{
			AuthorID: authorUID,
			Title:    title,
			Content:  content,
		}

		mock.ExpectQuery(createNews).WithArgs(priceGroup.AuthorID, priceGroup.Title, priceGroup.Content, priceGroup.Category).WillReturnRows(rows)

		createdNews, err := priceGroupRepo.Create(context.Background(), priceGroup)

		require.NoError(t, err)
		require.NotNil(t, createdNews)
		require.Equal(t, priceGroup.Title, createdNews.Title)
	})
}

func TestNewsRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	priceGroupRepo := NewPriceGroupRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		priceGroupUID := uuid.New()
		title := "title"
		content := "content"

		rows := sqlmock.NewRows([]string{"priceGroup_id", "title", "content"}).AddRow(priceGroupUID, title, content)

		priceGroup := &models.News{
			NewsID:  priceGroupUID,
			Title:   title,
			Content: content,
		}

		mock.ExpectQuery(updateNews).WithArgs(priceGroup.Title,
			priceGroup.Content,
			priceGroup.ImageURL,
			priceGroup.Category,
			priceGroup.NewsID,
		).WillReturnRows(rows)

		updatedNews, err := priceGroupRepo.Update(context.Background(), priceGroup)

		require.NoError(t, err)
		require.NotNil(t, updateNews)
		require.Equal(t, updatedNews, priceGroup)
	})
}

func TestNewsRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	priceGroupRepo := NewPriceGroupRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		priceGroupUID := uuid.New()
		mock.ExpectExec(deleteNews).WithArgs(priceGroupUID).WillReturnResult(sqlmock.NewResult(1, 1))

		err := priceGroupRepo.Delete(context.Background(), priceGroupUID)

		require.NoError(t, err)
	})
}
*/
