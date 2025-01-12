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
