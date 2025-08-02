package parser

import (
	"os"
	"strings"

	"github.com/yigithakverdi/concord/internal/utils"
	"github.com/yigithakverdi/concord/internal/utils/globals"
)

// Very simple parser, just reads the properties file from defualt location, then
// slurps the contents into memory, splits the content into lines, also ignoring
// specific tokens, then trims, and splits it according to split token values.
// Appends the results into a global map of properties.
func ParseApplicationProperties() {

	// Tokens to ignore in the properties file, such as commented lines, 
	// new lines, etc.
	var tokensToIgnore = []string{"#", "\n"}
	var tokensToSplit = []string{"=", ":", " ", "\t"}

	// Slurping all the contents to the memory, since all the time, properties file contaings
	// small amount of data.
	data, err := os.ReadFile(globals.DefaultPropertiesFileLocation)
	utils.Check(err)

	// Conver the byte slice to a `strings.Split` for processing the properties file,
	// line by line.
	//
	// TODO O(n^2) complexity, since we are splitting the whole file into lines, then checking each line
	// for tokens to ignore, this could be with a very small chance that the file is large, then 
	// below method is highly inefficient, could explore more efficient ways
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		
		// Check if the current line contains a token that needs to be ignored, if so
		// skip it., trims leading and trailing spaces, sets ignore flag to false by 
		// default, if the trimmed line contains token that needs to be ignored, 
		// flag set to true. 
		trimmed := strings.TrimSpace(line)
		ignore := false
		for _, token := range tokensToIgnore {
			if strings.HasPrefix(trimmed, token) || strings.Contains(trimmed, token) {
				ignore = true
				break
			}
		}

		// Depending on the ignore flag, and the trimmed line is empty string or not (
		// indicating empty line), either continue with further processing or skip the line
		if ignore || trimmed == "" {
			continue
		}

		// If no ignored tokens are found, then the line is processed further, by splitting it
		// into key-value pairs based on the tokens to split.
		//
		// TODO later on this could be extended to support more complex parsing, such as 
		// dynamic token splitting, depending on what type of split token used.
		parts := strings.Split(line, tokensToSplit[0])
		
		// Guard if the propert contain extra token to split thus making it more then 2 parts,
		// directly fail the parsing.
		if(len(parts) != 2) {
			panic("Invalid property format, expected key=value format, got: " + line)
		}

		// Append it to the global map of application properties.
		globals.GlobalProperties[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
}
