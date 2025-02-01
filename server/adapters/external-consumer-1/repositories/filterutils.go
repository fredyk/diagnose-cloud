package repositories

import (
	"time"

	wst "github.com/fredyk/westack-go/v2/common"
)

func buildWhereAndOrPatientNameDob(patientName string, dob time.Time) wst.Where {
	var where wst.Where
	if patientName != "" && !dob.IsZero() {
		where = wst.Where{
			"patient.name": patientName,
		}
	} else if patientName != "" {
		where = wst.Where{
			"patient.name": patientName,
		}
	} else if !dob.IsZero() {
		where = wst.Where{
			"date": dob,
		}
	}
	return where
}
