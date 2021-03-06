/*
 * memo-go
 *
 * メモのCRUD用API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject struct {

	// メモのタイトル
	Title string `json:"title"`

	// メモの中身
	Content string `json:"content"`
}

// AssertInlineObjectRequired checks if the required fields are not zero-ed
func AssertInlineObjectRequired(obj InlineObject) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"content": obj.Content,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseInlineObjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineObject (e.g. [][]InlineObject), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineObjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineObject, ok := obj.(InlineObject)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineObjectRequired(aInlineObject)
	})
}
