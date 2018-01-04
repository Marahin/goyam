package goyam

/*
Summary is a structure holding all warnings and errors
for comparison purposes
*/
type Summary struct {
	Warnings []string
	Errors   []string
}

// WarningsCount returns amount of warnings
func (s *Summary) WarningsCount() int {
	return len(s.Warnings)
}

// ErrorsCount returns amount of errors
func (s *Summary) ErrorsCount() int {
	return len(s.Errors)
}

// AddError appends passed error value to Errors array
// and returns it
func (s *Summary) AddError(err string) string {
	s.Errors = append(s.Errors, err)

	return err
}

// AddWarning appends passed error value to Warnings array and returns it
func (s *Summary) AddWarning(warn string) string {
	s.Warnings = append(s.Warnings, warn)

	return warn
}

// Merge merges passed summary into receiver summary
func (s *Summary) Merge(s2 Summary) {
	s.Warnings = append(s.Warnings, s2.Warnings...)
	s.Errors = append(s.Errors, s2.Errors...)
}

// NewSummary is a constructor function for Summary structure
func NewSummary(warnings []string, errors []string) *Summary {
	return &Summary{warnings, errors}
}
