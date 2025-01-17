package tenant

import "github.com/oklog/ulid/v2"

type Tenant struct{}

func (t Tenant) String() string {
	return ulid.MustNew(ulid.Now(), nil).String()
}
