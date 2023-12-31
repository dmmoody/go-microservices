package database

import (
	"context"
	"fmt"
	"githab.com/dmmoody/go-microservices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, email string) ([]models.Customer, error)
	GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	GetAllServices(ctx context.Context) ([]models.Service, error)
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=US/Eastern",
		"postgres",
		"postgres",
		"postgres",
		"postgres",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return &client, nil
}

func (c *Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
