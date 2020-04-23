package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/algobox_cf_services/structs"
)

func getUserURL(username string) string {
	return fmt.Sprintf(
		"https://codeforces.com/api/user.status?handle=%s&from=1&count=%d", username, MaxSubs)
}

// GetAllSubs gets all user submissions
func getAllSubs(handle string) []structs.Submission {
	resp, err := http.Get(getUserURL(handle))
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Printf("%s", e.Error())
		return nil
	}
	response := structs.UserStatus{}
	json.Unmarshal(body, &response)
	return response.Result
}

func getDashboardData(handle string) structs.DashboardData {
	subs := getAllSubs(handle)
	byTags := map[string]int{}
	byIndex := map[string]int{}
	byRating := map[int]int{}
	for _, v := range subs {
		if v.Verdict != "OK" {
			continue
		}
		for _, tag := range v.Problem.Tags {
			byTags[tag]++
		}
		byIndex[string(v.Problem.Index[0])]++
		byRating[v.Problem.Rating]++
	}
	return structs.DashboardData{
		ByIndex:  byIndex,
		ByRating: byRating,
		ByTags:   byTags,
	}
}
