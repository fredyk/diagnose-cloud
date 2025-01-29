package ports

type ExternalPatientPort interface {
	GetDiagnoses() error
}
