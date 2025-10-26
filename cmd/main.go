package main


import (
	"log"
    userent "github.com/deeptesh-rout/project-root/internal/user/ent"
    adminent "github.com/deeptesh-rout/project-root/internal/admin/ent"
    "github.com/deeptesh-rout/project-root/internal/orm"
)




func main() {
	db, err := orm.NewDB()
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	log.Println("Database initialized successfully:", db)

	// Example: creating Ent clients
	_ = userent.NewClient()
	_ = adminent.NewClient()
}
