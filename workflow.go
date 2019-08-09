package redmine

import (
	"encoding/json"
	"errors"
//	"net/http"
	"strconv"
	"strings"
)


type transitionResult struct {
	 Transition Transition `json:"transition"`
}


type Transition struct {
	OldStatusID int `json:"old_status_id"`
    NewStatusID int `json:"new_status_id"`
}




// Endpoint to get possible transitions for given role and tracker

func (c *Client) Transition (role_id int, tracker_id int) (*Transition, error) {
	res, err := c.Get(c.endpoint + "/workflow/transitions.json?" + "role_id=" + strconv.Itoa(role_id) + "&" + "tracker_id=" + strconv.Itoa(tracker_id) + c.apikey)



	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r transitionResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Transition, nil
}


