package ports

// StringPtr generates a pointer to a string.
func StringPtr(s string) *string { return &s }

// UintPtr generates a pointer to a uint.
func UintPtr(u uint) *uint { return &u }

// IntPtr generates a pointer to an int.
func IntPtr(i int) *int { return &i }
