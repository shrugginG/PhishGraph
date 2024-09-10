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
	log.Print(connStr)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
