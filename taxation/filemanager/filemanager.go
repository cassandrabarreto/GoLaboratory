package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}
 
func (fm FileManager) ReadLines() ([]string, error){
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, fmt.Errorf("opening file %q: %w", fm.InputFilePath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {

	file, err := os.Create(fm.OutputFilePath)
	defer file.Close()
	
	if err != nil {
		return errors.New("Failed to create file.")
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		//file.Close()
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return nil
	//return file.Close()
}

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}