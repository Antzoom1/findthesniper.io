package domains

// Domain is an interface for all domain data structures
type Domain interface {
	Validate() error
}
