package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"pavel-vacha.cz/ctc/internal/paingas/types"
)

// GetConfig načte vstupní soubor na základě cesty.
// Očekává vstupní formát ve tvaru Configuration.
func GetConfig(filePath string) (*types.Configuration, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := &types.Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filePath, err)
	}

	return c, nil
}
