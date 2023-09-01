package kuehnenagel

type BarcodeQualifier string

const (
	MessageVersion = "01.40"

	BarcodeIDS BarcodeQualifier = "IDS"
	BarcodeKDN                  = "KDN"
	BarcodeNVE                  = "NVE"
)

type TransportOrderExtFO struct {
	Envelope Envelope
	Message  []Message
}

type Envelope struct {
	SenderIdentification   string               `validate:"required,max=13"`
	ReceiverIdentification string               `validate:"required,max=13"`
	MessageType            string               `validate:"required,max=10"`
	MessageVersion         string               `validate:"required"`
	EnvelopeIdentification string               `validate:"required,max=13"`
	TransmissionDateTime   TransmissionDateTime `validate:"required"`
}

type TransmissionDateTime struct {
	Date string `validate:"required,len=10"`
	Time string `xml:",omitempty" json:",omitempty"`
}

type Message struct {
	TransportOrderHeader   TransportOrderHeader
	TransportOrderShipment []TransportOrderShipment
	TransportOrderTotal    TransportOrderTotal
}

type TransportOrderHeader struct {
	TransportOrderNumber string              `xml:",omitempty" json:",omitempty" validate:"max=35"`
	TransportOrderDate   *TransportOrderDate `validate:"dive"`
	CustomerNumber       string              `validate:"required,max=35"`
	CustomerName         string              `xml:",omitempty" json:",omitempty" validate:"max=35"`
	EDIRecipient         string              `xml:"EdiRecipient,omitempty" json:"EdiRecipient,omitempty" validate:"max=35"`
	CodeList             string              `xml:"Codelist,omitempty" json:"Codelist,omitempty" validate:"max=4"`
}

type TransportOrderDate struct {
	Date string `validate:"required,len=9"`
	Time string `xml:",omitempty" json:",omitempty"`
}

type TransportOrderShipment struct {
	TransportOrderPosition int    `validate:"required"`
	ShipmentNumber         string `validate:"required,max=35"`
	ShipmentDate           ShipmentDate
	TermsOfDelivery        TermsOfDelivery
	ShipmentReference      []string `xml:",omitempty" validate:"max=99,dive,max=35"`
	DeliveryNoteNumber     []string `xml:",omitempty" validate:"max=99,dive,max=35"`
	CustomOrderNumber      []string `xml:",omitempty" validate:"max=99,dive,max=35"`
	PickupOrderNumber      []string `xml:",omitempty" validate:"max=99,dive,max=35"`
	ShipperAddress         ShipperAddress
	ConsigneeAddress       ConsigneeAddress
	AdditionalInformation  *AdditionalInformation
	OtherAddress           []OtherAddress
	ShipmentItem           []ShipmentItem
	ShipmentTotal          ShipmentTotal
}

type ShipmentDate struct {
	Date string `validate:"required,len=10"`
	Time string `xml:",omitempty" json:",omitempty"`
}

type TermsOfDelivery struct {
	TermsOfDeliveryCode string `validate:"required,max=4"`
	TermsOfDeliveryText string `validate:"required,max=70"`
	TermsOfDeliveryCity string `validate:"required,max=70"`
}

type ShipperAddress struct {
	AddressID         string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Name1             string `validate:"required,max=35"`
	Name2             string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Street1           string `validate:"required,max=35"`
	Street2           string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Country           string `validate:"required,max=3"`
	Zip               string `validate:"required,max=9"`
	City              string `validate:"required,max=35"`
	CommunicationType string `xml:",omitempty" json:",omitempty" validate:"max=4"`
	Communication     string `xml:",omitempty" json:",omitempty" validate:"max=70"`
}

type ConsigneeAddress struct {
	AddressID         string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Name1             string `validate:"required,max=35"`
	Name2             string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Street1           string `validate:"required,max=35"`
	Street2           string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Country           string `validate:"required,max=3"`
	Zip               string `validate:"required,max=9"`
	City              string `validate:"required,max=35"`
	CommunicationType string `xml:",omitempty" json:",omitempty" validate:"max=4"`
	Communication     string `xml:",omitempty" json:",omitempty" validate:"max=70"`
}

type AdditionalInformation struct {
	DeliveryDateIndicator       string `xml:",omitempty" validate:"max=4"`
	DeliveryDateTimeFrom        *DeliveryDateTimeFrom
	DeliveryDateTimeTo          *DeliveryDateTimeTo
	DeliveryInstructionCoded    []DeliveryInstructionCoded `xml:",omitempty" validate:"dive"`
	DeliveryInstructionFreeText []string                   `xml:",omitempty" validate:"dive,max=70"`
	PickupInstructionFreeText   []string                   `xml:",omitempty" validate:"dive,max=70"`
}

type DeliveryDateTimeFrom struct {
	Date string `validate:"required,len=10"`
	Time string `xml:",omitempty" json:",omitempty"`
}

type DeliveryDateTimeTo struct {
	Date string `validate:"required,len=10"`
	Time string `xml:",omitempty" json:",omitempty"`
}

type DeliveryInstructionCoded struct {
	DeliveryInstructionsCode string
	DeliveryInstructionsText string `xml:",omitempty" validate:"max=70"`
}

type OtherAddress struct {
	AddressType       string `validate:"required"`
	AddressID         string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Name1             string `validate:"required,max=35"`
	Name2             string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Street1           string `validate:"required,max=35"`
	Street2           string `xml:",omitempty" json:",omitempty" validate:"max=35"`
	Country           string `validate:"required,max=3"`
	Zip               string `validate:"required,max=9"`
	City              string `validate:"required,max=35"`
	CommunicationType string `xml:",omitempty" json:",omitempty" validate:"max=4"`
	Communication     string `xml:",omitempty" json:",omitempty" validate:"max=70"`
}

type ShipmentItem struct {
	ShipmentItemCounter      int `validate:"required"`
	NumberOfPackages         NumberOfPackages
	NumberOfPackagesOnPallet *NumberOfPackagesOnPallet
	Dimensions               *Dimensions
	GrossWeight              GrossWeight
	Content                  string           `validate:"max=70"`
	MarksAndNumbers          string           `xml:",omitempty" validate:"max=70"`
	PackageBarcodes          []PackageBarcode `validate:"dive"`
}

type NumberOfPackages struct {
	Value       int    `validate:"required"`
	PackageType string `validate:"required,max=4"`
}

type NumberOfPackagesOnPallet struct {
	Value       int    `validate:"required"`
	PackageType string `validate:"required,max=4"`
}

type Dimensions struct {
	Length      float64 `xml:",omitempty" json:",omitempty"`
	Width       float64 `xml:",omitempty" json:",omitempty"`
	Height      float64 `xml:",omitempty" json:",omitempty"`
	MeasureUnit string  `xml:",omitempty" json:",omitempty"`
}

type GrossWeight struct {
	Value       float64 `validate:"required"`
	MeasureUnit string  `validate:"required,max=3"`
}

type PackageBarcode struct {
	BarcodeQualifier BarcodeQualifier `validate:"required,max=4"`
	Barcode          string           `validate:"max=35"`
}

type ShipmentTotal struct {
	TotalShipmentItem        int `validate:"required"`
	TotalShipmentPackages    int `validate:"required"`
	TotalShipmentGrossWeight TotalShipmentGrossWeight
}

type TotalShipmentGrossWeight struct {
	Value       float64 `validate:"required"`
	MeasureUnit string  `validate:"required,max=3"`
}

type TransportOrderTotal struct {
	TotalTransportOrderShipments     int `validate:"required"`
	TotalTransportOrderShipmentItems int `validate:"required"`
	TotalTransportOrderPackages      int `validate:"required"`
	TotalTransportOrderGrossWeight   TotalTransportOrderGrossWeight
}

type TotalTransportOrderGrossWeight struct {
	Value       float64 `validate:"required"`
	MeasureUnit string  `validate:"required,max=3"`
}
