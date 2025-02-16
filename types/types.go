package types

import (
	"errors"
	"regexp"
)


type RegisterUser struct{
  Name string `json:"name" validate:"required,min=3"`
  Email string `json:"email"`
  Phone string `json:"phone"`
  Role string `json:"role"`
}

type User struct{
  ID string `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  Phone string `json:"phone"`
  Role string `json:"role"`
}



func (r *RegisterUser) Validate() error {
  if len(r.Name) < 3 {
    return errors.New("name must be at least 3 characters long")
  }  
  if !isValidEmail(r.Email) {
    return errors.New("invalid email format!")
  }
  if !isValidPhone(r.Phone){
    return errors.New("invalid phone number format")
  }
  return nil
}




func isValidEmail(email string) bool {
  emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
  return emailRegex.MatchString(email)
}

func isValidPhone(phone string) bool {
  phoneRegex := regexp.MustCompile(`^\d{10}$`)
  return phoneRegex.MatchString(phone)
}

