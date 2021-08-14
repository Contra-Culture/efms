package emailConfirmation

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var emailConfirmationRecordMarking = []int{
	EMAIL_FLEN,
	TOKEN_FLEN,
	CREATED_AT_FLEN,
}

type emailConfirmation struct {
	email     string
	token     string
	createdAt time.Time
}

func emailConfirmation2Rec(ec *emailConfirmation, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(ec.email), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	bs, err = Normalize([]byte(ec.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(ec.createdAt.Format(DATETIME_FORMAT)), CREATED_AT_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	return
}
func rec2EmailConfirmation(rec []byte, handleErr func(error)) *emailConfirmation {
	splitted := Split(rec, emailConfirmationRecordMarking)
	dt, err := time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	e := emailConfirmation{
		email:     string(splitted[0]),
		token:     string(splitted[1]),
		createdAt: dt,
	}
	return &e
}
