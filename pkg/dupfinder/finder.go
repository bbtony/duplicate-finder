package dupfinder

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	ByValue ScanType = "value"
	ByKey   ScanType = "key"
)

var (
	ErrTypeScanNotAllowed = errors.New("type of scan is not allowed")
)

type ScanType string

type Param struct {
	File  string
	Key   string
	Value string
}

type Match struct {
	Matches    interface{} `json:"matches"`
	Duplicates []Param     `json:"duplicates"`
}

type ReportOfScanning struct {
	TypeScan ScanType `json:"type_scan"`
	Result   []Match  `json:"result"`
}

type fileInfo struct {
	FilePath string
	ENV      map[string]string
}

type DupFinder struct {
	TypeScan  ScanType
	FilesPath []string
}

func NewScanDupFinder(params []string, typeScan ScanType) *DupFinder {
	return &DupFinder{
		TypeScan:  typeScan,
		FilesPath: params,
	}
}

func (f *DupFinder) FindDuplicates() (map[string][]Param, error) {
	fileENVs, err := readEnvFile(f.FilesPath)
	if err != nil {
		return nil, err
	}

	var resultMap map[string][]Param

	switch f.TypeScan {
	case ByValue:
		resultMap = duplicatesByValue(fileENVs)
	case ByKey:
		resultMap = duplicatesByKey(fileENVs)
	default:
		return nil, ErrTypeScanNotAllowed
	}

	for key, el := range resultMap {
		if len(el) <= 1 {
			delete(resultMap, key)
		}
	}

	return resultMap, nil
}

func readEnvFile(params []string) ([]fileInfo, error) {
	var result []fileInfo
	for _, envFile := range params {
		envMap, err := godotenv.Read(envFile)
		if err != nil {
			return nil, err
		}
		result = append(result, fileInfo{FilePath: envFile, ENV: envMap})
	}

	return result, nil
}

func duplicatesByValue(fileENVs []fileInfo) map[string][]Param {
	resultMap := map[string][]Param{}
	for _, envFile := range fileENVs {
		for key, val := range envFile.ENV {
			if _, ok := resultMap[val]; ok { // add variable if it exists
				resultMap[val] = append(resultMap[val], Param{envFile.FilePath, key, val})
				continue
			}
			// if variable does not exist then create
			resultMap[val] = []Param{
				{envFile.FilePath, key, val},
			}
		}
	}

	return resultMap
}

func duplicatesByKey(fileENVs []fileInfo) map[string][]Param {
	resultMap := map[string][]Param{}
	for _, envFile := range fileENVs {
		for key, val := range envFile.ENV {
			if _, ok := resultMap[key]; ok { // add variable if it exists
				resultMap[key] = append(resultMap[key], Param{envFile.FilePath, key, val})
				continue
			}
			// if variable does not exist than create
			resultMap[key] = []Param{
				{envFile.FilePath, key, val},
			}
		}
	}

	return resultMap
}

func (f *DupFinder) Report(resOfScan map[string][]Param) *ReportOfScanning {
	report := &ReportOfScanning{
		TypeScan: f.TypeScan,
	}

	for key, val := range resOfScan {
		report.Result = append(report.Result, Match{
			Matches:    key,
			Duplicates: val,
		})
	}

	return report
}

func (r ReportOfScanning) ReportToJSONFile(filePath string) error {
	output, err := json.Marshal(r)
	if err != nil {
		return errors.New("error of prepare output")
	}
	f, err := os.Create(filePath)
	if err != nil {
		return errors.New("error of create report file")
	}

	_, err = f.Write(output)
	if err != nil {
		return errors.New("error of write to report file")
	}

	return nil
}

func Verbose(mapOfVariables map[string][]Param, t ScanType) {
	switch t {
	case ByValue:
		for val, listOfKeys := range mapOfVariables {
			if len(listOfKeys) >= 2 {
				fmt.Printf("Duplicate value: %s\n", val)
				for _, el := range listOfKeys {
					fmt.Printf("File: %s Variable: %s \n", el.File, el.Value)
				}
				fmt.Println()
			}
		}
	case ByKey:
		for _, listOfKeys := range mapOfVariables {
			if len(listOfKeys) >= 2 {
				fmt.Printf("Duplicate key: %s\n", listOfKeys[0].Key)
				for _, el := range listOfKeys {
					fmt.Printf("File: %s Variable: %s \n", el.File, el.Key)
				}
				fmt.Println()
			}
		}
	}
	for val, listOfKeys := range mapOfVariables {
		if len(listOfKeys) >= 2 {
			fmt.Printf("Duplicate value: %s\n", val)
			for _, el := range listOfKeys {
				fmt.Printf("File: %s Variable: %s \n", el.File, el.Key)
			}
			fmt.Println()
		}
	}
}
