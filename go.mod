module github.com/deeptesh-rout/project-root

go 1.25.3

require (
    entgo.io/ent v0.14.5
    github.com/google/uuid v1.6.0
    github.com/lib/pq v1.10.9
)

replace github.com/deeptesh-rout/project-root/internal/user/ent => ./internal/user/ent
replace github.com/deeptesh-rout/project-root/internal/admin/ent => ./internal/admin/ent
replace github.com/deeptesh-rout/project-root/internal/orm => ./internal/orm
