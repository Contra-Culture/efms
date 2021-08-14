package emailConfirmationRequirement

import (
	"time"

	. "github.com/Contra-Culture/efms/models/base"
)

var emailConfirmationRequirmentRecordMarking = []int{
	EMAIL_FLEN,
	TOKEN_FLEN,
	CREATED_AT_FLEN,
}

type emailConfirmationRequirement struct {
	email     string
	token     string
	createdAt time.Time
}

func emailConfirmationRequirement2Rec(ecr *emailConfirmationRequirement, handleErr func(error)) (rec []byte) {
	var err error
	var bs []byte
	rec, err = Normalize([]byte(ecr.email), EMAIL_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	bs, err = Normalize([]byte(ecr.token), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	bs, err = Normalize([]byte(ecr.createdAt.Format(DATETIME_FORMAT)), TOKEN_FLEN)
	if err != nil {
		handleErr(err)
		return nil
	}
	rec = append(rec, bs...)
	return
}
func rec2EmailConfirmationRequirement(rec []byte, handleErr func(error)) *emailConfirmationRequirement {
	var err error
	var dt time.Time
	splitted := Split(rec, emailConfirmationRequirmentRecordMarking)
	dt, err = time.Parse(string(splitted[2]), DATETIME_FORMAT)
	if err != nil {
		handleErr(err)
		return nil
	}
	ecr := emailConfirmationRequirement{
		email:     string(splitted[0]),
		token:     string(splitted[1]),
		createdAt: dt,
	}
	return &ecr
}
