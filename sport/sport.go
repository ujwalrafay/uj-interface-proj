package sport

import (
	"bufio"
	"encoding/json"
	"log"
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
	WriteToFile([]candidates, string) error
}

type DetailWriterJson struct{}
type DetailWriterYaml struct{}

func (dy DetailWriterYaml) WriteToFile(cl []candidates, file_path string) error {
	yaml_data, err := yaml.Marshal(&cl)
	if err != nil {
		log.Fatal("Error while marshaling yaml \t", err)
	}
	_, err = Write_file(file_path, yaml_data)
	if err != nil {
		log.Fatal("error writing yaml file")
	}
	return nil
}

func (dj DetailWriterJson) WriteToFile(cl []candidates, file_path string) error {

	json_data, err := json.MarshalIndent(&cl, "", "\t")
	if err != nil {
		log.Fatal("Error while marshaling json \t", err)
	}
	_, err = Write_file(file_path, json_data)
	if err != nil {
		log.Fatal("error writing json file")
	}
	return nil
}

func GetCandidates(filepath string) {
	//geting env varible
	viper.BindEnv("FORMAT")
	env_variable := viper.Get("FORMAT")
	format := "empty"
	//handling if env variable format not set
	if env_variable != nil {

		format = env_variable.(string)
	}

	dataf, err := os.Open(filepath)
	if err != nil {
		log.Fatal("some error occured \t", err)
	}

	lines := make([]candidates, 0)
	reader := bufio.NewScanner(dataf)

	for reader.Scan() {
		ln := string(reader.Text())
		temp, sports := Process_file(ln)

		age, err := strconv.Atoi(temp[1])
		height, err := strconv.ParseFloat(temp[4], 32)
		weight, err := strconv.Atoi(temp[5])
		if err != nil {
			log.Fatal("some error occured \t", err)
		}

		var temp_candidate = candidates{temp[0], age, string(temp[2][0]), sports, float32(height), weight}

		lines = append(lines, temp_candidate)
	}

	var djw DetailWriter
	var dyw DetailWriter
	//discarding case sensitivity by using EqualFold
	if strings.EqualFold(format, "json") {
		djw = DetailWriterJson{}
		djw.WriteToFile(lines, "json.json")
	} else if strings.EqualFold(format, "yaml") {
		dyw = DetailWriterYaml{}
		dyw.WriteToFile(lines, "yaml.yaml")
	} else {
		log.Fatal("Environment variable not set for format")
	}
}

func Process_file(ip string) (ips []string, sports []string) {

	start := strings.Index(ip, "[")
	end := strings.Index(ip, "]")

	sport := ip[start+1 : end]

	sports = strings.Split(sport, ",")

	ip = strings.ReplaceAll(ip, ip[start:end+1], "")

	ips = strings.Split(ip, ",")

	return ips, sports

}

func Write_file(write_file_path string, data_to_write []byte) (int, error) {

	f, err := os.Create(write_file_path)
	if err != nil {
		log.Fatal("Error occured while creating file to write \t", err)
	}
	w := bufio.NewWriter(f)
	n, err := w.WriteString(string(data_to_write))

	if err != nil {
		log.Fatal("Error while writing \t", err)
	}
	w.Flush()
	return n, err

}
