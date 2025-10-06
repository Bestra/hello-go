# Top Level Directory Structure

- `cmd/` - Main application commands
- `pkg/` - Library code that's okay to use by applications (e.g., `import "github.com/org/repo/pkg/foo`)
- `internal/` - Private application and library code
- `api/` - OpenAPI/Swagger specs, JSON schema files, protocol definition files
- `web/` - Web application specific components (HTML, CSS, JS)
- `scripts/` - Scripts for building, packaging, etc.
- `test/` - Additional external test applications and test data
- `docs/` - Documentation files
