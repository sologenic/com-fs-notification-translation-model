# Notification Translation Documentation

## Table of Contents

- [Overview](#overview)
- [translation.proto](#translation)
- [Version Information](#version-information)
- [Support](#support)

## Overview

The Notification Translation provides a comprehensive data structure for managing notification translation within the system.
## translation.proto

### Package Information

- **Package Name**: `translation`
- **Go Package Path**: `./translation`

### Overview

The `translation.proto` file defines the core translation model for notification translation management. It provides message types for representing translation data and operations. The file integrates with external utility libraries: `types.proto`, `parameters.proto`, `formats.proto`.

## Version Information

This documentation corresponds to the Protocol Buffer definitions in `translation.proto`. The proto file(s) use `proto3` syntax. When referencing this documentation, ensure that the version of the proto files matches the version of the generated code and API implementations you are using.

## Support

For additional information and support:
- See `README.md` for project setup, installation, and usage instructions
- Refer to the Protocol Buffer definitions in `translation.proto` for the authoritative source of truth
- Check the imported utility libraries for details on related types:
  - `sologenic/com-fs-notification-model/types/types.proto`
  - `sologenic/com-fs-notification-model/parameters/parameters.proto`
  - `sologenic/com-fs-notification-model/formats/formats.proto`
  - `sologenic/com-fs-email-template-model/emailtemplate.proto`
  - `sologenic/com-fs-utils-lib/models/language/language.proto`
