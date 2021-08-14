package registration

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var registrationRecordMarking = []int{
	LOGIN_FLEN,
	PASSWORD_FLEN,
	CREATED_AT_FLEN,
}

type registration struct {
	login     string
	password  string
	createdAt time.Time
}

func registration2Rec(r *registration, handleErr func(error)) (rec []byte) {
	var bs []byte
	var err error
	rec, err = Normalize([]byte(r.login), LOGIN_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	bs, err = Normalize([]byte(r.password), PASSWORD_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(r.createdAt.Format(DATETIME_FORMAT)), CREATED_AT_FLEN)
	if err != nil {
		handleErr(err)
		return []byte{}
	}
	rec = append(rec, bs...)
	return
}
func rec2Registration(rec []byte, handleErr func(error)) *registration {
	splitted := Split(rec, registrationRecordMarking)
	dt, err := time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	r := registration{
		login:     string(splitted[0]),
		password:  string(splitted[1]),
		createdAt: dt,
	}
	return &r
}
