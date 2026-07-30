# genome-nexus-go-api-client

Go API clients for the [Genome Nexus](https://github.com/genome-nexus/genome-nexus) public and internal APIs, generated with [OpenAPI Generator](https://openapi-generator.tech).

- `genome-nexus-public-api/` — client for the public API (spec: `genome-nexus-public-api/api/openapi.yaml`)
- `genome-nexus-internal-api/` — client for the internal API (spec: `genome-nexus-internal-api/api/openapi.yaml`)
- `templates/go/` — custom Mustache templates used to override the default `go` generator output

## Regenerating the clients

The clients are generated with the `openapi-generator-cli` `go` generator, using the custom templates in `templates/go/` (via `-t`) so that generated code differs from stock output — e.g. `model_simple.mustache` omits the `decoder.DisallowUnknownFields()` call so clients tolerate unknown/unrecognized response fields instead of failing to unmarshal.

1. Install `openapi-generator-cli` and pin it to the generator version recorded in that client's `.openapi-generator/VERSION` (`7.13.0` for `genome-nexus-public-api`):

   ```sh
   npm install @openapitools/openapi-generator-cli -g
   openapi-generator-cli version-manager set 7.13.0
   ```

2. Edit the OpenAPI spec (`api/openapi.yaml`) for whichever client you're updating, and/or edit the templates in `templates/go/`.

3. Regenerate, pointing at the custom templates:

   ```sh
   openapi-generator-cli generate \
     -i genome-nexus-public-api/api/openapi.yaml \
     -g go \
     -o genome-nexus-public-api \
     -t templates/go \
     --package-name genome_nexus_public_api
   ```

   Swap the `-i`, `-o`, and `--package-name` values for `genome-nexus-internal-api` to regenerate that client instead (check its own `.openapi-generator/VERSION` first, since it may differ).

4. Review the diff — the generator will rewrite all files not listed in that directory's `.openapi-generator-ignore`. Discard any unwanted changes (e.g. accidental `go.mod` regeneration) before committing.

5. Build to confirm the regenerated client still compiles:

   ```sh
   cd genome-nexus-public-api && go build ./...
   ```
