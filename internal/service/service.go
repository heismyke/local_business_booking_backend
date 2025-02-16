package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
  _"github.com/lib/pq"
	"github.com/heismyke/local_business_booking_app/internal/db/sqlc"
	"github.com/heismyke/local_business_booking_app/types"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Close() error
  CreateUserService(ctx context.Context, req *types.RegisterUser) (*sqlc.User, error)
}

type service struct {
	db      *sql.DB
	queries *sqlc.Queries
}



func New() Service {
	// Establishing connection to the database
	connStr :=  buildConnectionString() 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database", err)
	}

  db.SetMaxOpenConns(25)
  db.SetMaxIdleConns(25)
  db.SetConnMaxLifetime(2 * time.Hour)
  
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  if err := db.PingContext(ctx); err != nil {
    log.Fatalf("there was an error connecting to the database %v", err)
  }

  return &service{
    db : db,
    queries : sqlc.New(db),
  }
}

func buildConnectionString()string{
  return fmt.Sprintf(
    "postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
    os.Getenv("DB_USERNAME"),
		os.Getenv("DB_ROOT_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SCHEMA"),
    )

}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "Its healthy"

	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_close"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	if dbStats.OpenConnections > 40 {
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func (s *service) Close() error {
	log.Printf("Disconnected from database %s", os.Getenv("DB_DATABASE"))
	return s.db.Close()
}



func (s *service) CreateUserService (ctx context.Context, req *types.RegisterUser)(*sqlc.User, error){
    arg := sqlc.CreateUserParams{
      Name: req.Name, 
      Email: req.Email,
      Phone:req.Phone,
      Role: req.Role,
    }

    user, err := s.queries.CreateUser(ctx, arg)
     if err != nil {
       return nil, err
    }

    return &user,nil

}
