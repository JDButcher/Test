package pkg

import "errors"

type (
	Computer struct {
		ID          int64  `db:"id"`
		MACAddress  string `db:"mac_address"`
		Name        string `db:"name"`
		IP          string `db:"ip"`
		Employee    int64  `db:"employee"`
		Description string `db:"description"`
	}

	Employee struct {
		ID        int64  `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}
)

func (c *Computer) IsValid() error {

	if c.IP == "" {
		return errors.New("wrong IP")
	}

	return nil
}
