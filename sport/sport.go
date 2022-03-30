package sport

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
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
	WriteToFile(candidates) error
}

type DetailWriterJson struct{}
type DetailWriterYaml struct{}

func (dj DetailWriterJson) WriteToFile(c candidates) error {
	fmt.Println("yo")
	fmt.Println(c)
	return nil
}

func (dy DetailWriterYaml) WriteToFile(c candidates) error {
	fmt.Println("yo")
	fmt.Println(c)
	return nil
}

func GetCandidates(filepath string) {

	fmt.Println(viper.Get("SHELL"))

	dataf, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("some error occured", err)
	}
	fmt.Println(os.Getenv("SHELL"))

	lines := make([]candidates, 0)
	reader := bufio.NewScanner(dataf)
	list := list.New()
	for reader.Scan() {
		ln := string(reader.Text())
		// fmt.Println(ln)
		temp, sports := create_obj(ln)
		fmt.Println("----------------", temp, sports)
		age, err := strconv.Atoi(temp[1])
		height, err := strconv.ParseFloat(temp[4], 32)
		weight, err := strconv.Atoi(temp[5])
		if err != nil {
			fmt.Println("some error occured", err)
		}

		var temp_candidate = candidates{temp[0], age, string(temp[2][0]), sports, float32(height), weight}
		// temp_candidate.Sports = sports

		list.PushBack(temp_candidate)

		lines = append(lines, temp_candidate)
	}

	fmt.Println(lines[0])
	fmt.Println(lines)
	yaml_data, err := yaml.Marshal(&lines)
	yaml_data1, err := yaml.Marshal(&lines[0])
	json_data, err := json.Marshal(&lines)
	fmt.Println(string(yaml_data1))
	fmt.Println(string(yaml_data))

	write_yaml(yaml_data)
	// write_yaml(yaml_data1)
	write_json(json_data)

	var djw DetailWriter
	var dyw DetailWriter

	djw = DetailWriterJson{}
	djw.WriteToFile(lines[0])

	dyw = DetailWriterYaml{}
	dyw.WriteToFile(lines[0])

}

func create_obj(ip string) (ips []string, sports []string) {

	start := strings.Index(ip, "[")
	end := strings.Index(ip, "]")

	sport := ip[start+1 : end]

	sports = strings.Split(sport, ",")

	ip = strings.ReplaceAll(ip, ip[start:end+1], "")

	ips = strings.Split(ip, ",")
	// fmt.Println(ips, sports)

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
