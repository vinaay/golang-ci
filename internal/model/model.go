// Package model provides data models and validation functions.
package model

// Data represents a university/school data entry.
type Data struct {
	GUID       string `json:"guid"`
	School     string `json:"school"`
	Mascot     string `json:"mascot"`
	Nickname   string `json:"nickname"`
	Location   string `json:"location"`
	LatLong    string `json:"latlong"`
	NCAA       string `json:"ncaa,omitempty"`
	Conference string `json:"conference,omitempty"`
}

// ValidateGUID validates if a string is a valid GUID format.
// A GUID should be in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func ValidateGUID(guid string) bool {
	if len(guid) != 36 {
		return false
	}

	// Check format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	expectedDashes := []int{8, 13, 18, 23}
	for _, pos := range expectedDashes {
		if guid[pos] != '-' {
			return false
		}
	}

	// Check that all other characters are hexadecimal
	for i, r := range guid {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
			return false
		}
	}

	return true
}
