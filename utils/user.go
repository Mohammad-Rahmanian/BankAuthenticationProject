package utils

type User struct {
	Email       string `json:"email" bson:"email"`
	Lastname    string `json:"lastname" bson:"lastname"`
	NationalID  string `json:"nationalID" bson:"_id"`
	IP          string `json:"ip" bson:"ip"`
	FirstImage  string `json:"firstImage" bson:"firstImage"`
	SecondImage string `json:"secondImage" bson:"secondImage"`
	State       string `json:"state" bson:"state"`
}

func NewUSer(email string, lastname string, nationalId string, ip string, firstImage string, secondImage string, state string) *User {
	return &User{Email: email,
		Lastname:    lastname,
		NationalID:  nationalId,
		IP:          ip,
		FirstImage:  firstImage,
		SecondImage: secondImage,
		State:       state}

}
