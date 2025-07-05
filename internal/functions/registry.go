package functions

var funcCallHandlers = map[string]HandlerStruct{
	"multiplies_two_numbers": {
		Callback: createGenericHandler(multiplyCallback),
		Tool: tool{
			Type:        "function",
			Name:        "multiplies_two_numbers",
			Description: "multiplies two numbers",
			Parameters: parameters{
				Type: "object",
				Properties: map[string]property{
					"number1": {
						Type:        "number",
						Description: "The first number.",
					},
					"number2": {
						Type:        "number",
						Description: "The second number.",
					},
				},
				Required: []string{
					"number1",
					"number2",
				},
			},
		},
	},
}
