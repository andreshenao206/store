package utils

// FMap The higher-order-function takes an array and a function as arguments in order
// to apply the transformation to each element
func FMap(arr []interface{}, fn func(it interface{}) (r interface{}, err error)) ([]interface{}, error) {
	var nArray []interface{}
	for _, it := range arr {
		if mappedItem, err := fn(it); err != nil {
			return nil, err
		} else {
			// executing the passed method
			nArray = append(nArray, mappedItem)
		}
	}
	return nArray, nil
}
