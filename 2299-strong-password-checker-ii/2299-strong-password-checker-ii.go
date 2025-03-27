import "regexp"

func strongPasswordCheckerII(password string) bool {
    if len(password) < 8 {
        return false
    }

    if isMatch, _ := regexp.MatchString("[a-z]", password); !isMatch {
        return false
    }

    if isMatch, _ := regexp.MatchString("[A-Z]", password); !isMatch {
        return false
    }

    if isMatch, _ := regexp.MatchString("[0-9]", password); !isMatch {
        return false
    }

    if isMatch, _ := regexp.MatchString("[!@#$%^&*()\\-+]", password); !isMatch {
        return false
    }

    for i := 1; i < len(password); i++ {
        if password[i] == password[i-1] {
            return false
        }
    }

    return true
}