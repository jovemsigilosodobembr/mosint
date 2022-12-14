// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type EmailRepStruct struct {
	Email      string `json:"email"`
	Reputation string `json:"reputation"`
	Suspicious bool   `json:"suspicious"`
	References int    `json:"references"`
	Details    struct {
		Blacklisted             bool     `json:"blacklisted"`
		MaliciousActivity       bool     `json:"malicious_activity"`
		MaliciousActivityRecent bool     `json:"malicious_activity_recent"`
		CredentialsLeaked       bool     `json:"credentials_leaked"`
		CredentialsLeakedRecent bool     `json:"credentials_leaked_recent"`
		DataBreach              bool     `json:"data_breach"`
		FirstSeen               string   `json:"first_seen"`
		LastSeen                string   `json:"last_seen"`
		DomainExists            bool     `json:"domain_exists"`
		DomainReputation        string   `json:"domain_reputation"`
		NewDomain               bool     `json:"new_domain"`
		DaysSinceDomainCreation int      `json:"days_since_domain_creation"`
		SuspiciousTld           bool     `json:"suspicious_tld"`
		Spam                    bool     `json:"spam"`
		FreeProvider            bool     `json:"free_provider"`
		Disposable              bool     `json:"disposable"`
		Deliverable             bool     `json:"deliverable"`
		AcceptAll               bool     `json:"accept_all"`
		ValidMx                 bool     `json:"valid_mx"`
		PrimaryMx               string   `json:"primary_mx"`
		Spoofable               bool     `json:"spoofable"`
		SpfStrict               bool     `json:"spf_strict"`
		DmarcEnforced           bool     `json:"dmarc_enforced"`
		Profiles                []string `json:"profiles"`
	} `json:"details"`
}

func EmailRep(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var key string = GetAPIKey("Emailrep")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://emailrep.io/"+email, nil)
	req.Header.Set("Key", key)
	req.Header.Set("User-Agent", "mosint")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &emailrep_result)
}
