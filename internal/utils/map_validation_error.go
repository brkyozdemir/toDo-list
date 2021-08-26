package utils

func MapValidationError(errorsMap *string, paramName string) string {
	*errorsMap = paramName + " is required."

	return *errorsMap
}
