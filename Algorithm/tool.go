package Algorithm

import (
	"Cerebral-Palsy-Detection-System/model"
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

func StartAlgorithm(res *model.Result) {
	// star
	cmd := exec.Command("VProcessing/runv2.bat")

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//data, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Print(err, "print")
	//	log.Fatal(err)
	//}
	// out 转 string
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	res.VideoName = 18
	res.VideoRes = findPrediction(out.String())
}

func findPrediction(data string) string {

	rErr := regexp.MustCompile(`Connection timed out`)
	matcherr := rErr.MatchString(data)
	if matcherr {
		return "Error: Connection timed out"
	} else {
		re := regexp.MustCompile(`Prediction result: (label_\d+)`)
		match := re.FindStringSubmatch(data)
		label := match[1]
		if label == "" {
			log.Fatal()
		}
		return label
	}

}
