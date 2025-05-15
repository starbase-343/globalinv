package profile

import "fmt"

type Profile string

const (
	Dev  Profile = "dev"
	Prod Profile = "prod"
)

func Parse(profile string) (Profile, error) {
	switch profile {
	case string(Dev):
		return Dev, nil
	case string(Prod):
		return Prod, nil
	default:
		return "unknown", fmt.Errorf("unknown profile: %s", profile)
	}
}
