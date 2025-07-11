# Concord 🤝

**A CLI tool for managing `application.properties` files across multiple environments.**

Concord is a Go-based command-line tool for generating, validating, and managing environment-specific `application.properties` files. It uses a hierarchical source structure to produce consistent configuration artifacts for use in CI/CD pipelines.

-----

## Purpose

Managing `application.properties` across multiple environments (e.g., dev, preprod, beta, prod) can lead to configuration drift, inconsistencies, and accidental secret exposure. Concord addresses this by automating the generation and validation of these files from a defined source structure, reducing manual intervention and potential errors.

-----

## Architecture & Workflow

Concord merges a base configuration with an environment-specific override file, validates the result against a defined ruleset, and generates a final properties file. This ensures that every configuration is consistent and compliant with your defined policies before deployment.

-----

## Features

### Core Features

  * **Hierarchical Configuration**: Define common values in a `base.properties` file and specify only the differences in environment-specific files (e.g., `dev.properties`, `preprod.properties`).
  * **Rule-Based Validation**: Define rules in a central `concord.rules.yml` file to:
      * Enforce **required keys** for every environment.
      * Validate key naming conventions with **regex**.
      * Check for valid value formats (e.g., URLs, ports, enums).
      * **Disallow specific keys or values** in certain environments.
  * **Environment Diffing**: Compare the final generated properties of two environments using the `concord diff <env1> <env2>` command to quickly spot differences.
  * **CI/CD Integration**: The tool compiles to a **single, static Go binary** with no external runtime dependencies, making it easy to drop into any pipeline.

### Additional & Planned Features

  * **Secret Placeholder Resolution**: Support for placeholders (e.g., `db.password={{VAULT:secret/app/db-password}}`) for secrets stored in external systems.
  * **Multi-Format Support**: Use `.properties` or `.yml` files as configuration sources.
  * **Interactive Mode**: A `concord new` command to guide a user through creating a new environment file.
  * **Dry-Run Capability**: A `--dry-run` flag to preview the generated output without writing to a file.
  * **Comment Preservation**: Preserve comments from source files in the generated output.
  * **Pre-Commit Hook**: A script to run `concord validate` automatically before a commit.

-----

## Usage

### Generate a Configuration File

Generates a final `application-dev.properties` file by merging `base.properties` and `dev.properties`.

```sh
concord generate dev --output-file ./build/application-dev.properties
```

### Validate an Environment's Configuration

Validates the `preprod` environment against the rules defined in `concord.rules.yml`. Exits with a non-zero status code if validation fails, making it ideal for CI checks.

```sh
concord validate preprod
```

### See the Difference Between Environments

Shows a diff of the final generated properties for the `dev` and `preprod` environments.

```sh
concord diff dev preprod
```

-----

## Configuration Example

Rules are defined in a `concord.rules.yml` file in your repository's root.

**`concord.rules.yml.example`**

```yaml
# Concord Validation Rules
validation:
  # Rules apply to all environments unless specified otherwise
  global:
    - type: required_keys
      keys:
        - server.port
        - spring.application.name
        - graylog.host
        - graylog.port

    - type: naming_convention
      pattern: "^server\\.tomcat\\..+"
      description: "All Tomcat keys must start with 'server.tomcat.'"

  # Environment-specific rules
  environments:
    preprod:
      - type: disallowed_value
        key: "logging.level.root"
        value: "DEBUG"
        description: "Root logging level cannot be DEBUG in preprod."

    prod:
      - type: required_keys
        keys:
          - server.ssl.key-store # Example: an extra key required only for prod

```
