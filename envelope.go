package processor

type (
	Envelope struct {
		EOF      bool
		Input    AddressInput
		Output   AddressOutput
		Sequence int
	}

	AddressInput struct {
		Street1 string
		City    string
		State   string
		ZIPCode string
	}

	AddressOutput struct {
		Status        string
		DeliveryLine1 string
		LastLine      string
		City          string
		State         string
		ZIPCode       string
	}
)
