package ports

type PatientPort interface {
	GetDiagnoses() error
}
