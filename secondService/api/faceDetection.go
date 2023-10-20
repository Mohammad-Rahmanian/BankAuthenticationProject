package api

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	apiKey            = "acc_c6dc3b742095bd3"
	secretKey         = "eb001f6e29796a76c991e6036fdb7430"
	faceDetectionURL  = "https://api.imagga.com/v2/faces/detections?return_face_id=1"
	faceSimilarityURL = "https://api.imagga.com/v2/faces/similarity"
)

func FaceDetection(file *os.File) (string, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	file, err := os.Open(file.Name())
	if err != nil {
		logrus.Println("Can not open file", err)
		return "", err
	}

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		logrus.Println("Can not create form file:", err)
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		logrus.Println("Can not copy image to request:", err)
		return "", err
	}

	err = writer.Close()
	if err != nil {
		logrus.Println("Can not close writer:", err)
		return "", err
	}

	request, err := http.NewRequest("POST", faceDetectionURL, &requestBody)
	if err != nil {
		logrus.Println("Can not create request:", err)
		return "", err
	}
	request.SetBasicAuth(apiKey, secretKey)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logrus.Println("Can not make request:", err)
		return "", err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		logrus.Println("Can not read response:", err)
		return "", err
	}
	logrus.Println("Response:", buf.String())
	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		logrus.Println("can not unmarshaling JSON:", err)
		return "", err
	}
	faces, ok := result["result"].(map[string]interface{})["faces"].([]interface{})
	if !ok || len(faces) == 0 {
		logrus.Println(err)
		return "", err
	}
	face := faces[0].(map[string]interface{})
	faceID, _ := face["face_id"].(string)
	if !ok {
		return "", err
	}
	return faceID, nil
}

func FaceSimilarity(firstFace, secondFace string) (int, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", faceSimilarityURL+"?face_id="+firstFace+"&second_face_id="+secondFace, nil)
	if err != nil {
		logrus.Println("Can not create request:", err)
		return 0, err
	}
	request.SetBasicAuth(apiKey, secretKey)

	response, err := client.Do(request)
	if err != nil {
		logrus.Println("Can not send request to the server:", err)
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		logrus.Println("Non-OK status code received:", response.Status)
		return 0, err
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Println("Can not read response body:", err)
		return 0, err
	}
	var scoreResponse struct {
		Result struct {
			Score float64 `json:"score"`
		} `json:"result"`
	}
	err = json.Unmarshal(respBody, &scoreResponse)
	if err != nil {
		logrus.Println("Error unmarshaling JSON:", err)
		return 0, err
	}

	return int(scoreResponse.Result.Score), err
}
