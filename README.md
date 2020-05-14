## Install

```bash
go get -u github.com/gagliardetto/keyable
```

## Why I wrote it

I needed to make things happen in a program **while** it was running.

Plus, I also didn't want to write a lot of `if` statements.

So I wrote `pausable` to have a clean interface for having callbacks
associated with keys or key combinations.

## Usage

```golang
// Create a new Keyable object and add callbacks
// for the various characters and key combinations you send from
// the keyboard:
kb := keyable.New()

// onStat will be the callback for information for your program
// (progress, etc.)
onStat := func() {
	fmt.Println("some stats here")
	// TODO: add stats
}
// onQuit is a callback for when you quit the program:
onQuit := func() {
	fmt.Println("Exiting...")
	// TODO: clean up, send exit signals, etc.
	os.Exit(0)
}

// Add the callbacks:
kb.
	OnChar(
		onStat,
		'i',
	).
	OnKey(
		onQuit,
		keyboard.KeyEsc, keyboard.KeyCtrlC,
	).
	OnKey(
		func() {
			fmt.Println("FORCE-QUITTING...")
			os.Exit(1)
		},
		keyboard.KeyCtrlF,
	)
// Start the keyboard listener:
err := kb.Start()
if err != nil {
	panic(err)
}
// Don't forget to stop it!
defer kb.Stop()

```