package insertdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Result :How we recive the data
type Result struct {
	Data []struct {
		Type       string `json:"type"`
		Attributes struct {
			TeamOne      string `json:"TeamOne"`
			TeamTwo      string `json:"TeamTwo"`
			TeamOneGoals int    `json:"TeamOneGoals"`
			TeamTwoGoals int    `json:"TeamTwoGoals"`
		} `json:"attributes"`
	} `json:"data"`
}

// Error :How we send the Error info
type Error struct {
	Status string `json:"status"`
	Source struct {
		Pointer string `json:"pointer"`
	} `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// SuccessMessage :How we send the Success info
type SuccessMessage struct {
	Meta struct {
		Success struct {
			Title   string `json:"title"`
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"success"`
	} `json:"meta"`
}

//InsertData recive the data and process the info
func InsertData(r *http.Request) ([]byte, error) {
	var result Result

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &result)

	if err != nil {
		fmt.Println(err)
	}

	//Check if the request body is correct
	isSuccess, message := isJSONCorrect(result)

	if isSuccess {
		return getSuccessJSONBody("Game result was successfully inserted"), nil
	}

	return getErrorFromResult("/insert", "invalid Attribute", message), nil
}

func isJSONCorrect(r Result) (bool, string) {
	if len(r.Data[0].Attributes.TeamOne) == 0 || len(r.Data[0].Attributes.TeamTwo) == 0 {
		fmt.Println("TeamOne and TeamTwo need to have a value")
		return false, "TeamOne and TeamTwo need to have a value"
	}

	if r.Data[0].Attributes.TeamOneGoals == 0 || r.Data[0].Attributes.TeamTwoGoals == 0 {
		fmt.Println("TeamOneGoals and TeamTwoGoals need to have a value")
		return false, "TeamOneGoals and TeamTwoGoals need to have a value"
	}

	return true, "success"
}

func getSuccessJSONBody(message string) []byte {
	var successMessage SuccessMessage

	successMessage.Meta.Success.Title = "success"
	successMessage.Meta.Success.Status = "200"
	successMessage.Meta.Success.Message = message
	jsonBody, _ := json.Marshal(successMessage)

	return jsonBody
}

func getErrorFromResult(pointer string, title string, detail string) []byte {
	var errorMessage Error
	errorMessage.Status = "422"
	errorMessage.Source.Pointer = pointer
	errorMessage.Title = title
	errorMessage.Detail = detail
	jsonBody, _ := json.Marshal(errorMessage)

	return jsonBody
}
