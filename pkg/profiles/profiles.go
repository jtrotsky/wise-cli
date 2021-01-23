package profiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	ID      int64          `json:"id"`
	Type    string         `json:"type"`
	Details ProfileDetails `json:"details"`
}

// ProfileDetails contains the information about the profile entity.
type ProfileDetails struct {
	// Personal profile deatils
	FirstName      string
	LastName       string
	DateOfBirth    string // yy-mm-dd
	PhoneNumber    string
	Avatar         string
	Occupation     string
	PrimaryAddress int // address object ID
	// Business profile deatils
	Name                  string
	RegistrationNumber    string
	ACN                   string
	ABN                   string
	ARBN                  string
	CompanyType           string
	CompanyRole           string
	DescriptionOfBusiness string
	// same as personal: primaryAddress
	Webpage string
}

// [
//   {
//     "id": 217896,
//     "type": "personal",
//     "details": {
//       "firstName": "Oliver",
//       "lastName": "Wilson",
//       "dateOfBirth": "1977-07-01",
//       "phoneNumber": "+3725064992",
//       "avatar": "https://lh6.googleusercontent.com/-XzeFZ2PJE1A/AAAAAAAI/AAAAAAAAAAA/RvuvhXFsqs0/photo.jpg",
//       "occupation": null,
//       "primaryAddress": 236532
//     }
//   },
//   {
//     "id": 220192,
//     "type": "business",
//     "details": {
//       "name": "ABC Logistics Ltd",
//       "registrationNumber": "12144939",
//       "acn": null,
//       "abn": null,
//       "arbn": null,
//       "companyType": "LIMITED",
//       "companyRole": "OWNER",
//       "descriptionOfBusiness": "Information and communication",
//       "primaryAddress": 240402,
//       "webpage": "https://abc-logistics.com"
//     }
//   }
// ]

// Get all the profiles on a given Wise account
func Get(client *client.Client) ([]Profile, error) {
	profiles := []Profile{}

	response, err := client.DoRequest(http.MethodGet, "/v1/profiles", "")
	if err != nil {
		log.Fatal(err)
	}

	// Make sure response body is closed at end.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: test
	fmt.Printf("%s", body)

	err = json.Unmarshal(body, &profiles)
	if err != nil {
		log.Fatal(err)
	}

	return profiles, nil
}

// GetProfileByType returns the profile entity specified
func GetProfileByType(profiles []Profile, profileType string) (Profile, error) {
	for _, profile := range profiles {
		if strings.ToUpper(profile.Type) == EntityPersonal && strings.ToUpper(profileType) == EntityPersonal {
			return profile, nil
		} else if strings.ToUpper(profile.Type) == EntityBusiness && strings.ToUpper(profileType) == EntityBusiness {
			return profile, nil
		}
	}
	return Profile{}, errors.New("no profile match found")
}
