package grpc

import vCop "gopkg.in/go-playground/validator.v9"

var validator *vCop.Validate

func init() {
	validator = vCop.New()

	//Users
	validator.RegisterStructValidation(func(sl vCop.StructLevel) {
		r := sl.Current().Interface().(CreateUserRequest) //we cast it user

		if r.GetUser() == nil {
			sl.ReportError("NewUser", "newuser", "NewUser", "valid new user", "")
		} else {
			if len(r.GetUser().GetEmail()) == 0 {
				sl.ReportError("Email", "email", "Email", "valid-email", "")
			}
			if len(r.GetUser().GetFirstName()) == 0 {
				sl.ReportError("FirstName", "firstname", "FirstName", "valid-firstname", "")
			}
			if len(r.GetUser().GetLastName()) == 0 {
				sl.ReportError("LastName", "lastname", "LastName", "valid-lastname", "")
			}
			if len(r.GetUser().GetPassword()) == 0 {
				sl.ReportError("Password", "password", "Password", "valid-password", "")
			}
			if len(r.GetUser().GetConfirmPassword()) == 0 {
				sl.ReportError("ConfirmPassword", "confirmpassword", "ConfirmPassword", "valid-confirmpassword", "")
			}
		}

	})
}
