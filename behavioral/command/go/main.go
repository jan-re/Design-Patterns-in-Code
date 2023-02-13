package main

func main() {
	tv := tv{
		isRunning: false,
	}

	// Concrete command
	switchOn := switchOn{
		device: &tv,
	}
	// Concrete command
	switchOff := switchOff{
		device: &tv,
	}

	// Understand that the button type only has the execution method.
	// Depending on the type of button we need for our UI, we can feed it any command that fulfills the interface.
	// The button type by design carries no intention to be a button for switching on a TV.
	onButton := button{
		command: &switchOn,
	}

	offButton := button{
		command: &switchOff,
	}

	onButton.press()
	onButton.press()

	offButton.press()
	offButton.press()

	onButton.press()
}
