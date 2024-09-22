package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type DBConfig struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
}

func LoadConfig() (*DBConfig, error) {
	file, err := os.Open("/home/jxlu/project/PhishDetect/PhishGraph/chromedp-crawler/config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var config DBConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func ConnectDB() (*pgx.Conn, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func URLExists(url string) (bool, error) {
	conn, err := ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM webpage_links WHERE url = $1)`
	err = conn.QueryRow(context.Background(), query, url).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
