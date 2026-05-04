package service

import (
	"os"
	"testing"
)

func TestService_GetAllData(t *testing.T) {
	// Create a temporary test data file
	testData := `[
		{
			"guid": "05024756-765e-41a9-89d7-1407436d9a58",
			"school": "Test University",
			"mascot": "Test Mascot",
			"nickname": "Testers",
			"location": "Test City, ST, USA",
			"latlong": "0.0,0.0"
		}
	]`

	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	tmpFile.Close()

	svc, err := NewService(tmpFile.Name())
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}

	data := svc.GetAllData()
	if len(data) != 1 {
		t.Errorf("GetAllData() returned %d items, want 1", len(data))
	}
	if data[0].GUID != "05024756-765e-41a9-89d7-1407436d9a58" {
		t.Errorf("GetAllData() GUID = %v, want %v", data[0].GUID, "05024756-765e-41a9-89d7-1407436d9a58")
	}
}

func TestService_GetDataByGUID(t *testing.T) {
	testData := `[
		{
			"guid": "05024756-765e-41a9-89d7-1407436d9a58",
			"school": "Test University",
			"mascot": "Test Mascot",
			"nickname": "Testers",
			"location": "Test City, ST, USA",
			"latlong": "0.0,0.0"
		}
	]`

	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	tmpFile.Close()

	svc, err := NewService(tmpFile.Name())
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}

	// Test existing GUID
	data := svc.GetDataByGUID("05024756-765e-41a9-89d7-1407436d9a58")
	if data == nil {
		t.Fatal("GetDataByGUID() returned nil for existing GUID")
	}
	if data.GUID != "05024756-765e-41a9-89d7-1407436d9a58" {
		t.Errorf("GetDataByGUID() GUID = %v, want %v", data.GUID, "05024756-765e-41a9-89d7-1407436d9a58")
	}

	// Test non-existing GUID
	data = svc.GetDataByGUID("00000000-0000-0000-0000-000000000000")
	if data != nil {
		t.Error("GetDataByGUID() returned data for non-existing GUID")
	}
}

func TestNewService_InvalidFile(t *testing.T) {
	_, err := NewService("/nonexistent/file.json")
	if err == nil {
		t.Error("NewService() expected error for nonexistent file")
	}
}
