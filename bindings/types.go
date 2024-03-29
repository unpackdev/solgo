package bindings

// BindingType defines a custom type for representing different types of contract bindings. It is a string type that
// enables easy identification and differentiation of various contract standards and interfaces, such as ERC20,
// ERC20Ownable, etc. This type aids in the classification and management of contract interactions within the
// bindings package.
type BindingType string

// String is a receiver method of the BindingType type that converts the BindingType value back into a string.
// This method provides a straightforward way to obtain the string representation of a BindingType, facilitating
// logging, debugging, and any other scenarios where a string representation is more convenient or necessary.
func (b BindingType) String() string {
	return string(b)
}
