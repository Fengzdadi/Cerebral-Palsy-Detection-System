package Algorithm

import (
	"Cerebral-Palsy-Detection-System/model"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
)

func StartAlgorithm(res *model.VideoResult) {
	// star

	cmd := exec.Command("cmd.exe", "/C", ".\\VProcessing\\runV2.bat")
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//实时打印输出

	if err := cmd.Run(); err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	fmt.Println("Result: " + out.String())

	res.VideoName = 1
	res.VideoPath = ".\\VProcessing\\output.mp4"
	res.VideoRes = FindPrediction()
	res.Probability = FindProbability()
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
