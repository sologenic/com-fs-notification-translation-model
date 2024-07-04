# Notification translation model

The notification translation model re-uses the parameter, format and notification type from the notification model to construct a message for parsing the translations.

## Building the proto

See the [README](https://github.com/sologenic/fs-notification-model/README.md) for the used imports.

Important to note is:

* The proto linter does not seem to recognize imports so states the proto is invalid. This is not the case.
* The com-fs-notification-model is required to be (git) checked out in the location indicated in the proto_path. Do not change this path: Change your local structure to go idiomatic repo style instead (else we require complex scripts or personal scripts): The proto generator can not use git to find imports: They are all required to be present in the path. The generation scripts and the proto files take the repo structure into account in go idiomatic style only.
