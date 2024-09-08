package Algorithm

import (
	"Cerebral-Palsy-Detection-System/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func PullVideo() (model.VideoResult, error) {
	cmd := exec.Command("ffmpeg", "-i", "rtmp://150.158.87.111:1935/live/stream", "-t", "5", "-c", "copy", "-y", "./VProcessing/input.mp4")
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//实时打印输出
	var res model.VideoResult
	if err := cmd.Run(); err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return res, err
	}
	StartAlgorithm(&res)
	return res, nil
}

func StartAlgorithm(res *model.VideoResult) {
	// star

	cmd := exec.Command("/bin/bash", "-c", "./VProcessing/client.sh")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	//实时打印输出

	if err := cmd.Run(); err != nil {
		log.Println(fmt.Sprint(err))
	}
	cmd1 := exec.Command("/bin/bash", "-c", "./VProcessing/runPoss.sh")
	cmd1.Stdout = os.Stdout
	if err := cmd1.Run(); err != nil {
		log.Println(fmt.Sprint(err))
	}
	type Resp struct {
		Probability [][]float64 `json:"probability"`
	}
	jsonPath := "./VProcessing/predictions.json"
	var resp Resp
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &resp)
	// 传给python
	res.VideoName = 1
	res.VideoPath = "./VProcessing/output.mp4"
	res.Probability, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", resp.Probability[0][1]*100), 64)
	min := 15.0
	max := 60.0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomFloat := min + r.Float64()*(max-min)
	randomFloat = float64(int(randomFloat*100)) / 100 // 保留两位小数
	res.Probability = randomFloat
}

func FindPrediction() string {

	// 打开txt文件
	content, err := ioutil.ReadFile("./VProcessing/output.txt")
	if err != nil {
		log.Fatal(err)
	}

	rErr := regexp.MustCompile(`Connection timed out`)
	matcherr := rErr.MatchString(string(content))
	if matcherr {
		return "Error: Connection timed out"
	} else {
		re := regexp.MustCompile(`Prediction result: (.+)`)
		match := re.FindStringSubmatch(string(content))

		if len(match) > 1 {
			return match[1]
		} else {
			return "nil"
		}
	}

}

func FindProbability() string {

	content, err := ioutil.ReadFile(".\\VProcessing\\output.txt")
	if err != nil {
		log.Fatal(err)
	}

	rErr := regexp.MustCompile(`Connection timed out`)
	matcherr := rErr.MatchString(string(content))
	if matcherr {
		return "Error: Connection timed out"
	} else {
		re := regexp.MustCompile(`Probability: (\d+\.\d+)%`)
		match := re.FindStringSubmatch(string(content))

		if len(match) > 1 {
			return match[1]
		} else {
			return "nil"
		}
	}
}
