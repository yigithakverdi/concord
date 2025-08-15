package validator

import (
	"fmt"
	"strings"

	"github.com/yigithakverdi/concord/internal/utils"
	"github.com/yigithakverdi/concord/internal/utils/globals"
)

var envKeywords = []string{"alpha", "dev", "preprod"}

func containsRestrictedKeyword(restrictedKeywords []string, property string) bool {
	for _, kw := range restrictedKeywords {
		if strings.Contains(property, kw) {
			fmt.Printf("String: %v contains substring: %v\n", property, kw)
			return true
		}
	}
	return false
}

func SanityCheckBaseEnvironmentPropertiesFile() {
	condition := func(s string) bool { return s != globals.CurrentEnvironment }
	restrictedEnv := utils.Filter(envKeywords, condition)

	for key, val := range globals.GlobalProperties {
		if containsRestrictedKeyword(restrictedEnv, val) || containsRestrictedKeyword(restrictedEnv, key) {
			return
		}
	}
}
