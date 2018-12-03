package helm

// StatusType is an enumeration of the valid status for a release
type StatusType uint8

const (
	// DEPLOYED is the status that denotes a release has been deployed
	DEPLOYED StatusType = iota + 1
)

type status struct {
	Code      StatusType
	Resources string
}

type helmTimestamp struct {
	Seconds int64
	Nanos   int64
}

type info struct {
	Description   string
	Status        *status
	FirstDeployed *helmTimestamp `json:"first_deployed"`
	LastDeployed  *helmTimestamp `json:"last_deployed"`
}

// ReleaseStatus represents the status of a release in Helm
type ReleaseStatus struct {
	Name      string
	Namespace string
	Info      *info
}
