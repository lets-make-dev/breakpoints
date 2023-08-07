package main

import "fmt"

func printHelp() {
	logo := `
 ___                  _               _       _       
| _ ) _ _  ___  __ _ | |__ _ __  ___ (_) _ _ | |_  ___
| _ \| '_|/ -_)/ _` + "`" + ` || / /| '_ \/ _ \| || ' \|  _|(_-<
|___/|_|  \___|\__,_||_\_\| .__/\___/|_||_||_|\__|/__/
                          |_|
`
	// Splash screen
	fmt.Println(logo)

	darkGreen := "\033[38;5;2m" // ANSI code for dark green text

	//green := "\033[38;5;82m"            // ANSI code for green text
	black := "\033[38;5;0m"             // ANSI code for black text
	bold := "\033[1m"                   // ANSI code for bold text
	whiteBackground := "\033[48;5;15m"  // ANSI code for white background
	purpleBackground := "\033[48;5;93m" // ANSI code for purple background
	reset := "\033[0m"                  // ANSI code to reset to default color

	fmt.Printf("  Breakpoints is an %s%s%s X%sDebug %s CLI tool for %s PHPStorm %s\n", whiteBackground, bold, darkGreen, black, reset, purpleBackground, reset)
	fmt.Println("")
	fmt.Println("------------------------------------------------------")
	fmt.Println("                  Make Dev <m@ke.dev>                 ")
	fmt.Println("------------------------------------------------------")

	fmt.Println()
	fmt.Println(`Usage: bp [OPTIONS] <input> | COMMAND`)
	fmt.Println()
	fmt.Println(`Options:`)
	fmt.Println(`  -h, --help        Show this help message and exit.`)
	fmt.Println()
	fmt.Println(`Commands:`)
	fmt.Println(`  <file1:line> <file2:line> ...  `)
	fmt.Println(`                    Set Xdebug breakpoints. This can be one or more filenames`)
	fmt.Println(`                    with a line number.`)
	fmt.Println()
	fmt.Println(`  COMMAND           Any command that outputs relevant data can be piped to 'bp'`)
	fmt.Println(`                    to set breakpoints. For example: grep 'functionName' -rn . | bp.`)
	fmt.Println()
	fmt.Println(`  clear             Clear all set Xdebug breakpoints.`)
	fmt.Println()
	fmt.Println(`Examples:`)
	fmt.Println(`  bp app/Commands/Kernel.php:15 app/Commands/Kernel.php:23`)
	fmt.Println(`                    Set two breakpoints in the Kernel.php file on lines 15 and 25.`)
	fmt.Println(`  grep 'functionName' -rn . | bp`)
	fmt.Println(`                    Search the current dir for 'functionName' and set breakpoints.`)
	fmt.Println(`  bp clear          Clear all set breakpoints.`)
	//fmt.Println()
	//fmt.Println(`For more details, visit https://docs.example.com/bp-cli`)

	fmt.Println("\nNote:")
	fmt.Println("  Filenames should include the relative or absolute path, and line numbers should be valid integers.")
}
