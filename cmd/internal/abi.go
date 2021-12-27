package internal

type ABIProps struct {
	Version          string                     `json:"version"`
	Types            []ABITypesProps            `json:"types"`
	Structs          []ABIStructProps           `json:"structs"`
	Actions          []ABIActionsProps          `json:"actions"`
	Tables           []ABITablesProps           `json:"tables"`
	RicardianClauses []ABIRicardianClausesProps `json:"ricardian_clauses"`
	ErrorMessages    []ABIErrorMessagesProps    `json:"error_messages"`
	AbiExtensions    []ABIAbiExtensionsProps    `json:"abi_extensions"`
	Variants         []ABIVariantsProps         `json:"variants"`
	ActionResults    []ABIActionResultsProps    `json:"action_results"`
	KVTables         interface{}                `json:"kv_tables"`
	// []map[string]ABIKVTablesProps this field is not yet sure
}

type ABITypesProps struct {
	NewTypeName string `json:"new_type_name"`
	Type        string `json:"type"`
}

type ABIStructProps struct {
	Name   string                 `json:"name"`
	Base   string                 `json:"base"`
	Fields []ABIStructFieldsProps `json:"fields"`
}

type ABIStructFieldsProps struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ABITablesProps struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	IndexType string   `json:"index_type"`
	KeyNames  []string `json:"key_names"`
	KeyTypes  []string `json:"key_types"`
}

type ABIActionsProps struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	RicardianContract string `json:"ricardian_contract"`
}

type ABIRicardianClausesProps struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

type ABIErrorMessagesProps struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type ABIAbiExtensionsProps struct {
	Tag   int    `json:"tag"`
	Value string `json:"value"`
}

type ABIVariantsProps struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type ABIActionResultsProps struct {
	Name       string `json:"name"`
	ResultType string `json:"result_type"`
}

type ABIKVTablesProps struct {
	Type             string                                        `json:"type"`
	PrimaryIndex     ABIKVTablesPrimaryIndexProps                  `json:"primary_index"`
	SecondaryIndices []map[string]ABIKVTablesSecondaryIndicesProps `json:"secondary_indices"`
}

type ABIKVTablesPrimaryIndexProps struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ABIKVTablesSecondaryIndicesProps struct {
	Type string `json:"type"`
}
