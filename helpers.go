package pipl

// Summarize returns a string summary of the attributes of a person object
/*func (p Person) Summarize() (response string, err error) {
	builder := strings.Builder{}
	_, err = builder.WriteString(fmt.Sprintf("Match Confidence: %.f%%\n", p.Match*100))
	if err != nil {
		return
	}
	if len(p.Names) > 0 {
		_, err = builder.WriteString("Names:\n")
		if err != nil {
			return
		}
		for _, name := range p.Names {
			_, err = builder.WriteString(fmt.Sprintf("\t%s %s %s\n", name.First, name.Middle, name.Last))
			if err != nil {
				return
			}
		}
	}

	if len(p.Emails) > 0 {
		_, err = builder.WriteString("Email Addresses:\n")
		if err != nil {
			return
		}
		for _, email := range p.Emails {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", email.Address))
			if err != nil {
				return
			}
		}
	}

	if len(p.Usernames) > 0 {
		_, err = builder.WriteString("Usernames:\n")
		if err != nil {
			return
		}
		for _, username := range p.Usernames {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", username.Content))
			if err != nil {
				return
			}
		}
	}

	if len(p.Phones) > 0 {
		_, err = builder.WriteString("Phone Numbers:\n")
		if err != nil {
			return
		}
		for _, phone := range p.Phones {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", phone.Display))
			if err != nil {
				return
			}
		}
	}

	if p.Gender != nil {
		_, err = builder.WriteString(fmt.Sprintf("Gender:\n\t%s\n", p.Gender.Content))
		if err != nil {
			return
		}
	}

	if p.DateOfBirth != nil {
		_, err = builder.WriteString(fmt.Sprintf("Date of Birth:\n\t%s\n", p.DateOfBirth.Display))
		if err != nil {
			return
		}
	}

	if len(p.Languages) > 0 {
		_, err = builder.WriteString("Languages:\n")
		if err != nil {
			return
		}
		for _, language := range p.Languages {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", language.Display))
			if err != nil {
				return
			}
		}
	}

	if len(p.Ethnicities) > 0 {
		_, err = builder.WriteString("Ethnicities:\n")
		if err != nil {
			return
		}
		for _, ethnicity := range p.Ethnicities {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", ethnicity.Content))
			if err != nil {
				return
			}
		}
	}

	if len(p.OriginCountries) > 0 {
		_, err = builder.WriteString("Origin Countries:\n")
		if err != nil {
			return
		}
		for _, originCountry := range p.OriginCountries {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", originCountry.Country))
			if err != nil {
				return
			}
		}
	}

	if len(p.Addresses) > 0 {
		_, err = builder.WriteString("Addresses:\n")
		if err != nil {
			return
		}
		for _, address := range p.Addresses {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", address.Display))
			if err != nil {
				return
			}
		}
	}

	if len(p.Jobs) > 0 {
		_, err = builder.WriteString("Jobs:\n")
		if err != nil {
			return
		}
		for _, job := range p.Jobs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", job.Display))
			if err != nil {
				return
			}
		}
	}

	if len(p.Educations) > 0 {
		_, err = builder.WriteString("Education:\n")
		if err != nil {
			return
		}
		for _, education := range p.Educations {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", education.Display))
			if err != nil {
				return
			}
		}
	}

	if len(p.Relationships) > 0 {
		_, err = builder.WriteString("Relationships:\n")
		if err != nil {
			return
		}
		for _, relation := range p.Relationships {
			_, err = builder.WriteString(fmt.Sprintf("\t%s (%s, %s)\n", relation.Names[0].Display, relation.Type, relation.Subtype))
			if err != nil {
				return
			}
		}
	}

	if len(p.UserIDs) > 0 {
		_, err = builder.WriteString("User IDs:\n")
		if err != nil {
			return
		}
		for _, id := range p.UserIDs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", id.Content))
			if err != nil {
				return
			}
		}
	}

	if len(p.URLs) > 0 {
		_, err = builder.WriteString("Related URLs:\n")
		if err != nil {
			return
		}
		for _, url := range p.URLs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", url.URL))
			if err != nil {
				return
			}
		}
	}

	// Set the response
	response = builder.String()
	return
}*/

// NewPerson makes a new blank person object to be filled with terms
func NewPerson() *Person {
	return new(Person)
}

// AddName adds a name to the search object. For well defined names. Omit unused fields.
func (p *Person) AddName(firstName, middleName, lastName, prefix, suffix string) {
	// todo: add min/max validation
	newName := new(Name)
	newName.First = firstName
	newName.Middle = middleName
	newName.Last = lastName
	newName.Prefix = prefix
	newName.Suffix = suffix
	p.Names = append(p.Names, *newName)
}

// AddNameRaw can be used when you're unsure how to handle breaking down the name in
// question into its constituent parts. Basically, let Pipl handle parsing it.
func (p *Person) AddNameRaw(fullName string) {
	// todo: add min/max validation
	newName := new(Name)
	newName.Raw = fullName
	p.Names = append(p.Names, *newName)
}

// AddEmail appends an email address to the specified search object
func (p *Person) AddEmail(emailAddress string) {
	// todo: add min/max validation
	// todo: add email validation
	newEmail := new(Email)
	newEmail.Address = emailAddress
	p.Emails = append(p.Emails, *newEmail)
}

// AddUsername appends a username to the specified search object
func (p *Person) AddUsername(username string) {
	// todo: add min/max validation
	newUsername := new(Username)
	newUsername.Content = username
	p.Usernames = append(p.Usernames, *newUsername)
}

// AddUserID appends a user ID to the specified search object
func (p *Person) AddUserID(userID string) {
	// todo: add min/max validation
	newUserID := new(UserID)
	newUserID.Content = userID
	p.UserIDs = append(p.UserIDs, *newUserID)
}

// AddURL appends a URL to the specified search object (website, social, etc)
func (p *Person) AddURL(url string) {
	// todo: add min/max validation
	// todo: validate basic url requirements
	newURL := new(URL)
	newURL.URL = url
	p.URLs = append(p.URLs, *newURL)
}

// AddPhone appends a phone to the specified search object
func (p *Person) AddPhone(phoneNumber, countryCode int) {
	// todo: add min/max validation
	newPhone := new(Phone)
	newPhone.Number = phoneNumber
	if countryCode > 0 {
		newPhone.CountryCode = countryCode
	}
	p.Phones = append(p.Phones, *newPhone)
}

// AddPhoneRaw appends a phone to the specified search object
func (p *Person) AddPhoneRaw(phoneNumber string) {
	// todo: add min/max validation
	newPhone := new(Phone)
	newPhone.Raw = phoneNumber
	p.Phones = append(p.Phones, *newPhone)
}

// AddAddress appends an address to the specified search object
func (p *Person) AddAddress(house, street, apartment, city, state, country, poBox string) {
	// todo: add min/max validation
	// todo: validation on city/state/country?
	newAddress := new(Address)
	newAddress.House = house
	newAddress.Street = street
	newAddress.Apartment = apartment
	newAddress.City = city
	newAddress.State = state
	newAddress.Country = country
	newAddress.POBox = poBox
	p.Addresses = append(p.Addresses, *newAddress)
}

// AddAddressRaw can be used when many of the address parts are missing, or
// you're unsure how to split it up. Let Pipl handle parsing.
func (p *Person) AddAddressRaw(fullAddress string) {
	// todo: add min/max validation
	newAddress := new(Address)
	newAddress.Raw = fullAddress
	p.Addresses = append(p.Addresses, *newAddress)
}

// AddJob appends a job entry to the specified search object
func (p *Person) AddJob(title, organization, industry, dateRangeStart, dateRangeEnd string) {
	// todo: add min/max validation
	// todo: same test for dates (like DOB)
	newJob := new(Job)
	newJob.Title = title
	newJob.Organization = organization
	newJob.Industry = industry
	newJob.DateRange.Start = dateRangeStart
	newJob.DateRange.End = dateRangeEnd
	p.Jobs = append(p.Jobs, *newJob)
}

// AddEducation appends an education entry to the specified search object
func (p *Person) AddEducation(degree, school, dateRangeStart, dateRangeEnd string) {
	// todo: add min/max validation
	// todo: same test for dates (like DOB)
	newEducation := new(Education)
	newEducation.Degree = degree
	newEducation.School = school
	newEducation.DateRange.Start = dateRangeStart
	newEducation.DateRange.End = dateRangeEnd
	p.Educations = append(p.Educations, *newEducation)
}

// AddLanguage appends a language to the specified search object.
// Language is a 2 character language code (e.g. "en")
// Region  is a country code (e.g "US")
func (p *Person) AddLanguage(languageCode, regionCode string) {
	// todo: add min/max validation
	// todo: validate accepted languages and regions
	// todo: for the case of the characters? (uppercase)
	newLanguage := new(Language)
	newLanguage.Language = languageCode
	newLanguage.Region = regionCode
	p.Languages = append(p.Languages, *newLanguage)
}

// AddEthnicity appends an ethnicity to the specified search object
func (p *Person) AddEthnicity(ethnicity string) {
	// todo: add min/max validation
	// todo: accepted list of ethnicities?
	newEthnicity := new(Ethnicity)
	newEthnicity.Content = ethnicity
	p.Ethnicities = append(p.Ethnicities, *newEthnicity)
}

// AddOriginCountry appends an origin country to the specified search object
func (p *Person) AddOriginCountry(countryCode string) {
	// todo: add min/max validation
	// todo: accepted list of countries?
	newCountry := new(OriginCountry)
	newCountry.Country = countryCode
	p.OriginCountries = append(p.OriginCountries, *newCountry)
}

// AddRelationship appends a relationship entry to the specified search object
func (p *Person) AddRelationship(relationship Relationship) {
	p.Relationships = append(p.Relationships, relationship)
}

// SetGender sets the gender of the specified search object
func (p *Person) SetGender(gender string) {
	if gender != "male" && gender != "female" {
		gender = "male"
		// todo: return an error?
	}
	newGender := new(Gender)
	newGender.Content = gender
	p.Gender = newGender
}

// SetDateOfBirth sets the DOB of the specified search object
// DOB string format: "YYYY-MM-DD"
func (p *Person) SetDateOfBirth(dob string) {
	// todo: add validation to the DOB (yyyy-mm-dd)
	// todo: test date compared to today (cannot be future date)
	newDOB := new(DateOfBirth)
	newDOB.DateRange.Start = dob
	newDOB.DateRange.End = dob
	p.DateOfBirth = newDOB
}

// HasEmail returns true if the person has an email address
func (p *Person) HasEmail() bool {
	if len(p.Emails) > 0 {
		for _, email := range p.Emails {
			if email.Address != "" {
				return true
			}
		}
	}
	return false
}

// HasPhone returns true if the person has a phone number
func (p *Person) HasPhone() bool {
	if len(p.Phones) > 0 {
		for _, phone := range p.Phones {
			if (phone.CountryCode != 0 && phone.Number != 0) || (phone.Raw != "") {
				return true
			}
		}
	}
	return false
}

// HasUserID returns true if the person has a user id
func (p *Person) HasUserID() bool {
	if len(p.UserIDs) > 0 {
		for _, userID := range p.UserIDs {
			if userID.Content != "" {
				return true
			}
		}
	}
	return false
}

// HasUsername returns true if the person has a username
func (p *Person) HasUsername() bool {
	if len(p.Usernames) > 0 {
		for _, username := range p.Usernames {
			if username.Content != "" {
				return true
			}
		}
	}
	return false
}

// HasURL returns true if the person has a url
func (p *Person) HasURL() bool {
	if len(p.URLs) > 0 {
		for _, u := range p.URLs {
			if u.URL != "" {
				return true
			}
		}
	}
	return false
}

// HasName returns true if the person has a name (minimum required)
func (p *Person) HasName() bool {
	if len(p.Names) > 0 {
		for _, name := range p.Names {
			if (name.First != "" && name.Last != "") || (name.Raw != "") {
				return true
			}
		}
	}
	return false
}

// HasAddress returns true if the person has an address (minimum required)
func (p *Person) HasAddress() bool {
	if len(p.Addresses) > 0 {
		for _, address := range p.Addresses {
			if address.House != "" && address.Street != "" && address.City != "" && address.State != "" {
				return true
			}
		}
	}
	return false
}
