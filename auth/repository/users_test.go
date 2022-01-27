package repository

import (
	"log"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rs401/myauth/auth/models"
	"github.com/rs401/myauth/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	// Delete all previous test's users
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error Connecting to database: %v\n", err)
	}
	r := &usersRepository{
		db: conn.DB(),
	}
	err = r.DeleteAll()
	if err != nil {
		log.Fatalf("Error deleting all users before tests: %v", err)
	}
}

func TestGetAllEmpty(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	testempty := struct {
		name      string
		db        *gorm.DB
		wantUsers []*models.User
		wantErr   bool
	}{
		name:      "Get all.",
		db:        conn.DB(),
		wantUsers: []*models.User{},
		wantErr:   false,
	}

	t.Run(testempty.name, func(t *testing.T) {
		r := &usersRepository{
			db: testempty.db,
		}
		gotUsers, err := r.GetAll()
		if (err != nil) != testempty.wantErr {
			t.Errorf(
				"usersRepository.GetAll() error = %v, wantErr %v",
				err,
				testempty.wantErr,
			)
			return
		}
		if !reflect.DeepEqual(gotUsers, testempty.wantUsers) {
			t.Errorf(
				"usersRepository.GetAll() = %v, want %v",
				gotUsers,
				testempty.wantUsers,
			)
		}
	})

}

func TestSave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	testSave := struct {
		name    string
		db      *gorm.DB
		user    *models.User
		wantErr bool
	}{
		name: "Save User",
		db:   conn.DB(),
		user: &models.User{
			Name:     "TestUser1",
			Email:    "testuser1@test.com",
			Password: []byte("testpassword"),
		},
		wantErr: false,
	}

	t.Run(testSave.name, func(t *testing.T) {
		r := &usersRepository{
			db: testSave.db,
		}
		if err := r.Save(testSave.user); (err != nil) != testSave.wantErr {
			t.Errorf(
				"usersRepository.Save() error = %v, wantErr %v",
				err,
				testSave.wantErr,
			)
		}
	})

}

func TestSaveAndGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	// Guaranty a user exists
	testSave := struct {
		name    string
		db      *gorm.DB
		user    *models.User
		wantErr bool
	}{
		name: "Save User",
		db:   conn.DB(),
		user: &models.User{
			Name:     "TestUser2",
			Email:    "testuser2@test.com",
			Password: []byte("testpassword"),
		},
		wantErr: false,
	}

	t.Run(testSave.name, func(t *testing.T) {
		r := &usersRepository{
			db: testSave.db,
		}
		if err := r.Save(testSave.user); (err != nil) != testSave.wantErr {
			t.Errorf("usersRepository.Save() error = %v, wantErr %v", err, testSave.wantErr)
		}
	})

	// Grab first user for ID
	var first models.User
	result := conn.DB().Find(&first)
	assert.NoError(t, result.Error)

	testGetId := struct {
		name    string
		db      *gorm.DB
		id      uint
		wantErr bool
	}{
		name:    "Get User By ID",
		db:      conn.DB(),
		id:      first.ID,
		wantErr: false,
	}

	t.Run(testGetId.name, func(t *testing.T) {
		r := &usersRepository{
			db: testGetId.db,
		}
		if _, err := r.GetById(testGetId.id); (err != nil) != testGetId.wantErr {
			t.Errorf("usersRepository.GetById() error = %v, wantErr %v", err, testGetId.wantErr)
		}
	})
}

func TestDeleteAll(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error Connecting to database: %v\n", err)
	}

	testDelAll := struct {
		name    string
		db      *gorm.DB
		wantErr bool
	}{

		name:    "DeleteAll",
		db:      conn.DB(),
		wantErr: false,
	}

	t.Run(testDelAll.name, func(t *testing.T) {
		r := &usersRepository{
			db: testDelAll.db,
		}
		if err := r.DeleteAll(); (err != nil) != testDelAll.wantErr {
			t.Errorf("usersRepository.DeleteAll() error = %v, wantErr %v", err, testDelAll.wantErr)
		}
	})

}
