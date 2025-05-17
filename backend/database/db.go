package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./shifts.db")
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS shifts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
		role TEXT NOT NULL,
		location TEXT
	);
	
	CREATE TABLE IF NOT EXISTS shift_requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		shift_id INTEGER NOT NULL,
		worker_id INTEGER NOT NULL,
		status TEXT CHECK(status IN ('pending', 'approved', 'rejected')) DEFAULT 'pending',
		requested_at TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (shift_id) REFERENCES shifts(id),
		FOREIGN KEY (worker_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        role TEXT NOT NULL
    );
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	row := db.QueryRow("SELECT COUNT(*) FROM shifts")
	var count int
	_ = row.Scan(&count)

	if count == 0 {
		log.Println("Seeding initial data...")

		db.Exec(`
    INSERT INTO shifts (date, start_time, end_time, role, location) VALUES
		('2025-05-16', '08:00', '12:00', 'Cashier', 'Store A'),
        ('2025-05-17', '08:00', '12:00', 'Cashier', 'Store A'),
        ('2025-05-18', '13:00', '17:00', 'Stocker', 'Store B'),
        ('2025-05-19', '09:00', '13:00', 'Helper', 'Store A'),
        ('2025-05-20', '10:00', '14:00', 'Cashier', 'Store C'),
        ('2025-05-21', '15:00', '19:00', 'Stocker', 'Store C')
    `)

		db.Exec(`
    INSERT INTO shift_requests (shift_id, worker_id, status) VALUES
        (1, 2, 'pending'),
        (2, 3, 'pending'),
        (3, 4, 'approved'),
        (1, 4, 'rejected'),
        (6, 2, 'approved')
    `)

		db.Exec(`
	INSERT INTO users (username, password, role) VALUES 
		('admin', '$2a$14$qUwJDm6ohqqGZA2FCMcTRuVmvOnPZyfEVWNofRmN2vSmONDzlRdqu', 'admin'),
		('justin', '$2a$14$c94N4QLzt5t4mqFtWKoxpOlpHw3hzFubU7tKDA34nwBC.Eyaca4Ue', 'worker'),
		('payd', '$2a$14$c94N4QLzt5t4mqFtWKoxpOlpHw3hzFubU7tKDA34nwBC.Eyaca4Ue', 'worker'),
		('darma', '$2a$14$c94N4QLzt5t4mqFtWKoxpOlpHw3hzFubU7tKDA34nwBC.Eyaca4Ue', 'worker')
	`)
	}

	fmt.Println("Database connected and migrated.")
}

func GetDB() *sql.DB {
	return db
}
