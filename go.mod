module github.com/grangerp/test-pkg-repo

go 1.23.4

require github.com/grangerp/test-pkg-repo/tenant v0.0.0-00010101000000-000000000000

require (
	github.com/docker/distribution v2.8.3+incompatible
	github.com/oklog/ulid/v2 v2.1.0 // indirect
)

replace github.com/grangerp/test-pkg-repo/tenant v0.0.0-00010101000000-000000000000 => ./tenant
