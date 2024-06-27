package repositories

// StringPtr transforms a string to a pointer to string.
func StringPtr(s string) *string { return &s }

// IntPtr transforms an int to a pointer to int.
func IntPtr(i int) *int { return &i }
