/*
 * memo-go
 *
 * メモのCRUD用API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse200 struct {

	// メモのid
	Id int32 `json:"id"`

	// メモのタイトル
	Title string `json:"title"`
}

// AssertInlineResponse200Required checks if the required fields are not zero-ed
func AssertInlineResponse200Required(obj InlineResponse200) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"title": obj.Title,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseInlineResponse200Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineResponse200 (e.g. [][]InlineResponse200), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineResponse200Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineResponse200, ok := obj.(InlineResponse200)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineResponse200Required(aInlineResponse200)
	})
}
