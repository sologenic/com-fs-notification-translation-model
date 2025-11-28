# Notification Translation Documentation

## Table of Contents

- [Overview](#overview)
- [translation.proto](#translation)
  - [Messages](#messages)
    - [NotificationTranslationRequest](#notificationtranslationrequest)
    - [NotificationTranslationResponse](#notificationtranslationresponse)
    - [EmailTranslationRequest](#emailtranslationrequest)
    - [EmailTranslationResponse](#emailtranslationresponse)
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

### Messages

#### NotificationTranslationRequest {#notificationtranslationrequest}

The `NotificationTranslationRequest` message provides notificationtranslationrequest data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| Language | `language.Language` | Required | Language information |
| NotificationType | `notification.types.NotificationType` | Required | Type classification for this item (see related enum) |
| Parameter | `notification.parameters.Parameter` | Optional | Parameter information |
| Format | `notification.formats.Format` | Required | Format information |

**Use Cases:**
- Creating new notificationtranslationrequest records
- Retrieving notificationtranslationrequest information
- Updating notificationtranslationrequest data

**Important Notes:**
- This message provides the notificationtranslationrequest representation

#### NotificationTranslationResponse {#notificationtranslationresponse}

The `NotificationTranslationResponse` message provides notificationtranslationresponse data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| Text | `string` | Required | Text value |

**Use Cases:**
- Creating new notificationtranslationresponse records
- Retrieving notificationtranslationresponse information
- Updating notificationtranslationresponse data

**Important Notes:**
- This message provides the notificationtranslationresponse representation

#### EmailTranslationRequest {#emailtranslationrequest}

The `EmailTranslationRequest` message provides emailtranslationrequest data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| LanguageCode | `language.Lang` | Required | Internal language enum, refer to `language.proto` in `com-fs-utils-lib/models/language` |
| EmailTemplateType | `emailtemplate.EmailTemplateType` | Required | Type classification for this item (see related enum) |

**Use Cases:**
- Creating new emailtranslationrequest records
- Retrieving emailtranslationrequest information
- Updating emailtranslationrequest data

**Important Notes:**
- This message provides the emailtranslationrequest representation

#### EmailTranslationResponse {#emailtranslationresponse}

The `EmailTranslationResponse` message provides emailtranslationresponse data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| Subject | `string` | Required | Subject value |
| Body | `string` | Required | Body value |

**Use Cases:**
- Creating new emailtranslationresponse records
- Retrieving emailtranslationresponse information
- Updating emailtranslationresponse data

**Important Notes:**
- This message provides the emailtranslationresponse representation

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
