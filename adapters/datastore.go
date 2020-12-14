package adapters

import (
	"database/sql"

	"github.com/VividCortex/mysqlerr"
	"github.com/Zzocker/bl-otp/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
	"github.com/go-sql-driver/mysql"
)

type OTPDatastore struct {
	DB *sql.DB
}

const (
	create = "INSERT INTO otp(email,otp,expiry_time) values(?,?,?)"
	read   = "SELECT email,otp,expiry_time FROM otp WHERE email=?"
	update = "UPDATE otp SET otp=?,expiry_time=? WHERE email=?"
	delete = "DELETE FROM otp WHERE email=?"
)

// Create :
func (ds *OTPDatastore) Create(in model.EmailOTP) *errors.Er {
	stmt, err := ds.DB.Prepare(create)
	if err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	if _, err := stmt.Exec(in.Email, in.OTP, in.ExpiryTime); err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	return nil
}

// Read :
func (ds *OTPDatastore) Read(email string) (*model.EmailOTP, *errors.Er) {
	stmt, err := ds.DB.Prepare(read)
	if err != nil {
		return nil, errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	var out model.EmailOTP
	if err := stmt.QueryRow(email).Scan(&out.Email, &out.OTP, &out.ExpiryTime); err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_KEY_NOT_FOUND {
				return nil, errors.NewMsgln(errors.NOT_FOUND, "otp not found")
			}
		}
		return nil, errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	return &out, nil
}

// Update :
func (ds *OTPDatastore) Update(in model.EmailOTP) *errors.Er {
	stmt, err := ds.DB.Prepare(update)
	if err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	res, err := stmt.Exec(in.OTP, in.ExpiryTime, in.Email)
	if err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	count, _ := res.RowsAffected()
	if count != 1 {
		return errors.NewMsgln(errors.NOT_FOUND, "otp not found")
	}
	return nil
}

// Delete :
func (ds *OTPDatastore) Delete(email string) *errors.Er {
	stmt, err := ds.DB.Prepare(delete)
	if err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	res, err := stmt.Exec(email)
	if err != nil {
		return errors.NewMsgln(errors.INTERNAL, err.Error())
	}
	count, _ := res.RowsAffected()
	if count != 1 {
		return errors.NewMsgln(errors.NOT_FOUND, "otp not found")
	}
	return nil
}

// CREATE TABLE otp(
// 	email VARCHAR(200) PRIMARY KEY,
// 	otp VARCHAR(4) NOT NULL,
// 	expiry_time BIGINT NOT NULL
// );

// docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root -d mysql