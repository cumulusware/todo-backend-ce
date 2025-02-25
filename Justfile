set dotenv-filename := "config/.env.dev"

app_name := "todo-backend"
app_port := "8080"
coverage_file := "coverage.out"

# List the available justfile recipes.
[group('general')]
@default:
  just --list --unsorted

# List the lines of code in the project.
[group('general')]
loc:
  scc --remap-unknown "-*- Justfile -*-":"justfile"

# List the outdated direct dependencies (can be slow).
[group('dependencies')]
outdated:
  # (requires https://github.com/psampaz/go-mod-outdated).
  go list -u -m -json all | go-mod-outdated -update -direct

# Run go mod tidy and verify.
[group('dependencies')]
tidy:
  go mod tidy
  go mod verify

# Format and vet Go code. Runs before tests.
[group('test')]
check:
	go fmt ./...
	go vet ./...

# Lint using staticcheck. Format and vet code.
[group('test')]
lint: check
	staticcheck -f stylish ./...

# Run the unit tests.
[group('test')]
unit *FLAGS: check
  go test ./... -cover -vet=off -race {{FLAGS}} -short

# Run the integration tests.
[group('test')]
int *FLAGS: check
  go test ./... -cover -vet=off -race {{FLAGS}} -run Integration

# HTML report for unit (default), int, e2e, or all tests.
[group('test')]
cover test='unit': check
  go test ./... -vet=off -coverprofile={{coverage_file}} \
  {{ if test == 'all' { '' } \
    else if test == 'int' { '-run Integration' } \
    else if test == 'e2e' { '-run E2E' } \
    else { '-short' } }}
  go tool cover -html={{coverage_file}}

# Use curl to interact with todos endpoints.
[group('test')]
mod todos

# Build for local operating system.
[group('build')]
local:
  env go build -o dist/{{app_name}} ./cmd/api/

# Build and run for local operating system.
[group('build')]
dev: local
	dist/{{app_name}}

# Build in Docker container.
[group('build')]
docker:
	docker build -t {{app_name}}:latest .

# Build and run in Docker container.
[group('build')]
dock: docker
	docker run --rm -p {{app_port}}:{{app_port}} {{app_name}}:latest

# Login to IBM Cloud, target resource group, and project.
[group('deploy')]
login:
  ibmcloud login --sso
  ibmcloud target -g todo-rg
  ibmcloud ce project select --name todo-ce
  ibmcloud ce project list

# Create the Code Engine app from source code.
[group('deploy')]
create: check tidy
  ibmcloud ce app create --name  {{app_name}} --build-source .

# Deploy to Code Engine from source code.
[group('deploy')]
deploy: check tidy
  ibmcloud ce app update --name {{app_name}} --build-source .

# Determine URL for Code Engine app.
[group('deploy')]
url:
  ibmcloud ce app get --name {{app_name}} --output=url

# Check the application status.
[group('deploy')]
status:
  ibmcloud ce app get -n {{app_name}}

# Delete application on Code Engine.
[group('deploy')]
delete:
  ibmcloud ce app delete -n {{app_name}}
