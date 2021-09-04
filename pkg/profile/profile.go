package profile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jtrotsky/wise-cli/pkg/client"
)

// Business and personal profiles assigned to the account
const (
	EntityBusiness = "BUSINESS"
	EntityPersonal = "PERSONAL"
)

// Profile is a personal and / or business entity associated with a Wise account.
type Profile struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"`
	UserID       int64  `json:"userId"`
	CurrentState string

	// Address Address
	// Email                string
	// CreatedAt            time.Time
	// UpdatedAt            time.Time
	// Obfuscated           bool
	// Avatar               string
	// LocalizedInformation []string
	// FirstName            string
	// LastName             string
	// DateOfBirth          string // yyyy-mm-dd
	// PlaceOfBirth PlaceOfBirth
	// PhoneNumber        string
	// SecondaryAddresses []string
	// FullName           string
	// BusinessName                string
	// RegistrationNumber          string
	// DescriptionOfBusiness       string
	// CompanyType                 string
	// CompanyRole                 string
	// BusinessFreeFormDescription string
	// FirstLevelCategory          string
	// SecondLevelCategory         string
	// OperationalAddresses        []string
}

// Get all the profiles on a given Wise account
func Get(client *client.Client) ([]Profile, error) {
	profiles := []Profile{}

	response, err := client.DoRequest(http.MethodGet, "/v2/profiles", url.Values{})
	if err != nil {
		log.Fatal(err)
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &profiles)
	if err != nil {
		log.Fatal(err)
	}

	return profiles, nil
}

// GetProfileByType returns the profile entity specified
func GetProfileByType(profiles []Profile, profileType string) (Profile, error) {
	// TODO: user input to choose profile instead of dump all
	for _, profile := range profiles {
		if strings.ToUpper(profile.Type) == EntityPersonal && strings.ToUpper(profileType) == EntityPersonal {
			return profile, nil
		} else if strings.ToUpper(profile.Type) == EntityBusiness && strings.ToUpper(profileType) == EntityBusiness {
			return profile, nil
		}
	}
	return Profile{}, errors.New("no profile match found")
}
