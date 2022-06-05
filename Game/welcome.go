package game

import "fmt"

func WelcomePlayer() {
	asciiArt := `
	 _   _    _    _   _  ____ __  __    _    _   _
	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|
	`
	fmt.Println(asciiArt)
}
