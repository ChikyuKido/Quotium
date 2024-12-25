package untis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

type UntisClient struct {
	School   string
	Username string
	Password string
	BaseUrl  string
	ID       string
	client   *http.Client
}

type Teacher struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Forename    string `json:"forename"`
	LongName    string `json:"longName"`
	Displayname string `json:"displayname"`
	Title       string `json:"title"`
}

type TeacherJson struct {
	Data struct {
		Teachers []Teacher `json:"elements"`
	} `json:"data"`
}

func NewUntisClient(school, username, password string, baseUrl string) *UntisClient {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	return &UntisClient{
		School:   school,
		Username: username,
		Password: password,
		BaseUrl:  baseUrl,
		ID:       username,
		client:   client,
	}
}

func (c *UntisClient) Login() error {
	formData := url.Values{}
	formData.Set("school", c.School)
	formData.Set("j_username", c.Username)
	formData.Set("j_password", c.Password)
	req, err := http.NewRequest("POST", c.BaseUrl+"/WebUntis/j_spring_security_check", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create a request %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make a request %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed. Maybe wrong credentials %o", resp.StatusCode)
	}
	return nil

}
func (c *UntisClient) Logout() error {
	resp, err := c.client.Get(c.BaseUrl + "/WebUntis/saml/logout")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusFound {
		return fmt.Errorf("logout failed %o", resp.StatusCode)
	}
	return nil
}

func (c *UntisClient) GetTeachers() ([]Teacher, error) {
	req, err := http.NewRequest("GET", c.BaseUrl+"/WebUntis/api/public/timetable/weekly/pageconfig?type=2&date=2024-12-22&isMyTimetableSelected=false", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to sucessfully get data with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)

	}
	var teacherJson TeacherJson
	err = json.Unmarshal(body, &teacherJson)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %v", err)
	}

	return teacherJson.Data.Teachers, nil
}

func (c *UntisClient) DownloadTeacherImage(path string, teacherID uint) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/WebUntis/image.do?cat=2&id=%d", c.BaseUrl, teacherID), nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	err = os.WriteFile(path, body, 644)
	if err != nil {
		return fmt.Errorf("error writing teacher image: %v", err)
	}
	return nil
}
