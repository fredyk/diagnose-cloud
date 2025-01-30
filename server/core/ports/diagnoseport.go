package ports

type DiagnosePort interface {
	GetPatientId() string
	GetDiagnose() string
	GetPrescription() string
	GetDate() string
}
