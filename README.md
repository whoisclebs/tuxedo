# Tuxedo  
[![GoDoc](https://pkg.go.dev/badge/github.com/whoisclebs/tuxedo.svg)](https://pkg.go.dev/github.com/whoisclebs/tuxedo) ![GitHub Created At](https://img.shields.io/github/created-at/whoisclebs/tuxedo?style=flat) [![GitHub release](https://img.shields.io/github/v/tag/whoisclebs/tuxedo?include_prereleases&style=flat&color=blue)](https://github.com/whoisclebs/tuxedo/releases) [![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10092/badge)](https://www.bestpractices.dev/projects/10092) [![FOSSA Status](https://app.fossa.com/api/projects/custom%2B52040%2Fgithub.com%2Fwhoisclebs%2Ftuxedo.svg?type=shield&issueType=license)](https://app.fossa.com/projects/custom%2B52040%2Fgithub.com%2Fwhoisclebs%2Ftuxedo?ref=badge_shield&issueType=license) [![FOSSA Status](https://app.fossa.com/api/projects/custom%2B52040%2Fgithub.com%2Fwhoisclebs%2Ftuxedo.svg?type=shield&issueType=security)](https://app.fossa.com/projects/custom%2B52040%2Fgithub.com%2Fwhoisclebs%2Ftuxedo?ref=badge_shield&issueType=security)
<img align="right" height="320" src="./.github/assets/tuxedo.png">

Tuxedo is a lightweight, chainable HTTP client library for Go that simplifies working with the standard net/http package. With a clean, intuitive API inspired by popular libraries like Resty, Tuxedo allows you to build and execute HTTP requests with minimal boilerplate code. It supports setting headers, bodies, and even tracing, and comes with built-in JSON unmarshaling for easy response handling. Ideal for developers who need a simple yet powerful tool to streamline HTTP interactions in Go projects.

## Instalalation

```
go get -u github.com/whoisclebs/tuxedo
```

## Basic Usage
Below is a quick example demonstrating how to make a GET request with Tuxedo. Feel free to adjust as needed:

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/whoisclebs/tuxedo"
)

func main() {
    // Create a client with a 30-second timeout
    client := tuxedo.NewClient(30 * time.Second)
    // Build a request, adding a custom header
    response, err := client.R().
        AddHeader("Accept", "application/json").
        Get("https://api.github.com")
    if err != nil {
        log.Fatalf("Error while making request: %v", err)
    }
    // Print out the status code and body
    fmt.Printf("Status: %d\n", response.StatusCode)
    fmt.Println("Body:", string(response.Body))
}
```
You can also easily unmarshal the response into a struct using Json() or Xml():
```go
type GithubAPIResponse struct {
    CurrentUserURL string `json:"current_user_url"`
}

func main() {
    client := tuxedo.NewClient(30 * time.Second)
    response, err := client.R().Get("https://api.github.com")
    if err != nil {
        log.Fatalf("Request error: %v", err)
    }
    var data GithubAPIResponse
    if err := response.Json(&data); err != nil {
        log.Fatalf("JSON unmarshal error: %v", err)
    }
    fmt.Println("Current User URL:", data.CurrentUserURL)
}

```

## Project Structure

```
📦 tuxedo  
┣ 📜 client.go                 # HTTP client implementation  
┣ 📜 client_test.go            # Unit tests for the HTTP client  
┣ 📜 CODE_OF_CONDUCT.md        # Contributor guidelines and behavior expectations  
┣ 📜 CONTRIBUTING.md           # Instructions on how to contribute to the project  
┣ 📜 go.mod                    # Go module file with dependencies  
┣ 📜 LICENSE                   # License for the project  
┣ 📜 README.md                 # Main documentation file with usage instructions  
┣ 📜 request.go                # Request builder and customization  
┣ 📜 response.go               # HTTP response handling  
┣ 📜 SECURITY.md               # Security policies and vulnerability reporting guidelines  
┣ 📜 tuxedo_test.go            # General tests for the library  
┣ 📜 utils.go                  # Utility functions for internal use  
```
## Commit Message Standardization
We use Conventional Commits enforced by Git Commitizen to maintain a clear and structured commit history.

### Commit Types
<details> <summary>Click here to expand the full list of commit types</summary>

- **feat:** A new feature
- **fix:** A bug fix
- **docs:** Documentation changes only
- **style:** Changes that do not affect code behavior (e.g., formatting, missing semicolons)
- **refactor:** Code refactoring that does not change functionality
- **perf:** Performance improvements
- **test:** Adding or modifying tests
- **chore:** Maintenance or build process changes (e.g., CI/CD, dependencies updates)
- **build:** Changes that affect the build system or external dependencies
- **ci:** Changes to the CI configuration files or scripts
- **revert:** Reverting a previous commit
- **hotfix:** Urgent bug fix deployed immediately
</details>

#### Commit Format
```
<type>(<scope>): <short description>
[optional blank line]
<long description (if necessary)>
```
#### Example Commit Messages:
```
git commit -m "feat(client): add custom timeout setting for Tuxedo client"
git commit -m "fix(request): ensure 'Content-Type' header is set to 'application/json' if none provided"
git commit -m "refactor(utils): reorganize helper functions and clean up code"
git commit -m "docs: add example usage in README for JSON and XML parsing"
```

Using these standardized commit messages ensures better traceability, clearer change logs, and a more streamlined workflow for everyone contributing to the project.

---
## License
Distributed under the MIT License.