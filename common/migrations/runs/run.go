package migrate

// import (
// 	"log"
// 	"rice-wine-shop/core/configs"

// 	"github.com/golang-migrate/migrate"
// 	"gorm.io/gorm"
// )

// func RunMigrations() {

// 	// Kết nối tới cơ sở dữ liệu bằng GORM
// 	db, err := gorm.Open(postgres.Open(configs.Get().DataSource), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Không thể kết nối tới cơ sở dữ liệu: %v", err)
// 	}

// 	// Lấy đối tượng sql.DB từ GORM
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("Không thể lấy đối tượng sql.DB từ GORM: %v", err)
// 	}

// 	// Tạo driver cho golang-migrate
// 	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
// 	if err != nil {
// 		log.Fatalf("Không thể tạo driver cho migration: %v", err)
// 	}

// 	// Tạo đối tượng migrate, trỏ đến thư mục migrations
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file://common/migrations",
// 		"postgres",
// 		driver,
// 	)
// 	if err != nil {
// 		log.Fatalf("Không thể tạo đối tượng migrate: %v", err)
// 	}

// 	// Thực hiện migration
// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("Lỗi khi chạy migration: %v", err)
// 	}

// 	log.Println("Migration đã hoàn tất!")
// }
