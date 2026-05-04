// Package service provides business logic for data loading and caching.
package service

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"api/internal/model"
	"api/pkg/logger"
)

// DataService defines the interface for data operations.
type DataService interface {
	GetAllData() []model.Data
	GetDataByGUID(guid string) *model.Data
}

// Service handles data loading and caching.
type Service struct {
	data     []model.Data
	mu       sync.RWMutex
	filePath string
}

// NewService creates a new service instance and loads data from the specified file.
func NewService(filePath string) (*Service, error) {
	s := &Service{
		filePath: filePath,
	}

	if err := s.LoadData(); err != nil {
		return nil, fmt.Errorf("failed to load data: %w", err)
	}

	return s, nil
}

// LoadData loads data from the JSON file into memory.
func (s *Service) LoadData() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("could not read data file: %w", err)
	}

	var data []model.Data
	if err := json.Unmarshal(file, &data); err != nil {
		return fmt.Errorf("could not unmarshal data: %w", err)
	}

	s.data = data
	logger.Info("Data loaded successfully", "count", len(data), "file", s.filePath)

	return nil
}

// GetAllData returns all data entries.
func (s *Service) GetAllData() []model.Data {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]model.Data, len(s.data))
	copy(result, s.data)
	return result
}

// GetDataByGUID returns a data entry by GUID, or nil if not found.
func (s *Service) GetDataByGUID(guid string) *model.Data {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i := range s.data {
		if s.data[i].GUID == guid {
			// Return a copy
			result := s.data[i]
			return &result
		}
	}

	return nil
}

// Reload reloads data from the file.
func (s *Service) Reload() error {
	return s.LoadData()
}
