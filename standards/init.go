package standards

func init() {
	// Initialize the storage map so it can be accessed globally.
	storage = make(map[Standard]EIP)
}
