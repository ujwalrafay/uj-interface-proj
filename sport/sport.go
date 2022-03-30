package sport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type candidates struct {
	Name   string
	Age    int
	Gender string
	Sports []string
	Height float32
	Weight int
}
type DetailWriter interface {
	WriteToFile([]candidates, string) error
}

type DetailWriterJson struct{}
type DetailWriterYaml struct{}

func (dy DetailWriterYaml) WriteToFile(cl []candidates, file_path string) error {
	yaml_data, err := yaml.Marshal(&cl)
	if err != nil {
		fmt.Println("Error while marshaling yaml", err)
	}
	write_file(file_path, yaml_data)
	return nil
}

func (dj DetailWriterJson) WriteToFile(cl []candidates, file_path string) error {

	json_data, err := json.MarshalIndent(&cl, "", "\t")
	if err != nil {
		fmt.Println("Error while marshaling yaml", err)
	}
	write_file(file_path, json_data)
	return nil
}

func GetCandidates(filepath string) {

	dataf, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("some error occured", err)
	}

	lines := make([]candidates, 0)
	reader := bufio.NewScanner(dataf)

	for reader.Scan() {
		ln := string(reader.Text())
		temp, sports := process_file(ln)

		age, err := strconv.Atoi(temp[1])
		height, err := strconv.ParseFloat(temp[4], 32)
		weight, err := strconv.Atoi(temp[5])
		if err != nil {
			fmt.Println("some error occured", err)
		}

		var temp_candidate = candidates{temp[0], age, string(temp[2][0]), sports, float32(height), weight}

		lines = append(lines, temp_candidate)
	}

	var djw DetailWriter
	var dyw DetailWriter

	djw = DetailWriterJson{}
	djw.WriteToFile(lines, "json.json")

	dyw = DetailWriterYaml{}
	dyw.WriteToFile(lines, "yaml.yaml")

}

func process_file(ip string) (ips []string, sports []string) {

	start := strings.Index(ip, "[")
	end := strings.Index(ip, "]")

	sport := ip[start+1 : end]

	sports = strings.Split(sport, ",")

	ip = strings.ReplaceAll(ip, ip[start:end+1], "")

	ips = strings.Split(ip, ",")

	return ips, sports

}

func write_file(write_file_path string, data_to_write []byte) {

	f, err := os.Create(write_file_path)
	if err != nil {
		fmt.Println("Error occured while creating file to write", err)
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(data_to_write))

	if err != nil {
		fmt.Println("Error while writing ", err)
	}
	w.Flush()

}
