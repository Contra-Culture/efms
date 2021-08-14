package email

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var emailRecordMarking = []int{
	LOGIN_FLEN,
	EMAIL_FLEN,
	CREATED_AT_FLEN,
}

type email struct {
	login     string
	email     string
	createdAt time.Time
}

func email2Rec(e *email, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(e.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	bs, err = Normalize([]byte(e.email), EMAIL_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(e.createdAt.Format(DATETIME_FORMAT)), EMAIL_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	return
}
func rec2Email(rec []byte, handleErr func(error)) *email {
	splitted := Split(rec, emailRecordMarking)
	dt, err := time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	e := email{
		login:     string(splitted[0]),
		email:     string(splitted[1]),
		createdAt: dt,
	}
	return &e
}
