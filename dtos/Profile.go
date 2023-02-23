package dtos;

type Auth0User struct {
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	Nickname string `json:"nickname"`
	FamilyName string `json:"family_name"`
	GivenName string `json:"given_name"`
	Name string `json:"name"`
	Picture string `json:"picture"`
	UserID	string `json:"user_id"`
	Identities []struct{
		Provider	string `json:"provider"`
		UserID		string `json:"user_id"`
		Connection	string `json:"connection"`
		IsSocial	bool `json:"isSocial"`
	} 
}