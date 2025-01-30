package data

type QueueInfo struct {
	Name     string
	Key      string
	Exchange string
	NoWait   bool
}

type AdditionInput struct {
	FirstNumber  float64 `json:"firstNumber"`
	SecondNumber float64 `json:"secondNumber"`
	RequestId    string  `json:"requestId"`
}

type AdditionResult struct {
	Result    float64 `json:"result"`
	RequestId string  `json:"requestId"`
}

type PersonalInformationInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	City      string `json:"city"`
	Zipcode   string `json:"zipcode"`
	Street    string `json:"street"`
	RequestId string `json:"requestId"`
}

type PersonalInformationOutput struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Address   Address  `json:"address"`
	RequestId string   `json:"requestId"`
	MetaData  MetaData `json:"metaData"`
}

type Address struct {
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Street  string `json:"street"`
}

type MetaData struct {
	ServiceData ServiceData `json:"serviceData"`
}

type ServiceData struct {
	ServiceName    string `json:"serviceName"`
	ServiceVersion string `json:"serviceVersion"`
}
