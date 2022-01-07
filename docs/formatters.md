# Formatters

Logger allows you to turn raw string messages into a variety of formats
using its modular Formatters. Logger currently support the following 
formats:
* [Logger Standard Log Format](#logger-standard-log-format)
* [Kubernetes Structured Log Format](#kubernetes-structured-log-format)
* [Graylog Extended Log Format](#graylog-extended-log-format)

## Logger Standard Log Format
This is the default Logger formatter. It provides a standard, consistent format for your logs while still being both file and CLI friendly.

LSLF provides 3 required fields, and one optional field.

```log
time="2022-01-06T13:10:53-0600" level="fatal" msg="example standard out message with identifier" ident="aaaa-bbbb-cccc-dddd"
```

The `time` and `level` fields are generated programmatically. The `time` field the current time, and is formatted based on the `DateTimeFormat` value given when calling the `DefaultHandler{}` struct. If unset, the `time` field defaults to *ISO-8601* format.

The `level` field is determined by the logging function called, which also takes in the user message and sets the log `msg`. `.Error("foobar")` sets the level to `error` and the `msg` to `foobar`, for example.

The `ident` field is an optional field that can be used to specify a group identifier
for your logs. In the event that you may want to see a group of logs together (such)
as related to a user login flow or a Kubernetes controller reconciliation loop, you
can provide all of the logs the same `ident` using 
`l.WithIdentifier('aaa-bbb').Error("foobar)` to attach an identifier to that log entry.

### Output Reference

|  Field  | User-Provided |                     Details                        |
|:-------:|:-------------:|:-------------------------------------------------------:|
|  `time` |      No       | The current time. Defaults to **ISO 8601** formatting   |
| `level` |      No       | The log level                                           |
|  `msg`  |      Yes      | The message to be logged                                |
| `ident` |      Yes      | A grouping identifier, used to identify related entries |

## Kubernetes Structured Log Format
This is the default Logger formatter. It provides a standard, consistent format for your logs while still being both file and CLI friendly.

LSLF provides 3 required fields, and one optional field.

```log

```

The `ts` and `level` fields are generated programmatically. The `time` field the current time, and is formatted based on the `DateTimeFormat` value given when calling the `DefaultHandler{}` struct. If unset, the `time` field defaults to *ISO-8601* format.

The `level` field is determined by the logging function called, which also takes in the user message and sets the log `msg`. `.Error("foobar")` sets the level to `error` and the `msg` to `foobar`, for example.

The `ident` field is an optional field that can be used to specify a group identifier
for your logs. In the event that you may want to see a group of logs together (such)
as related to a user login flow or a Kubernetes controller reconciliation loop, you
can provide all of the logs the same `ident` using 
`l.WithIdent('aaa-bbb').Error("foobar)` to attach an identifier to that log entry.

### Output Reference

|  Field  | User-Provided |                     Details                        |
|:-------:|:-------------:|:-------------------------------------------------------:|
|  `time` |      No       | The current time. Defaults to **ISO 8601** formatting   |
| `level` |      No       | The log level                                           |
|  `msg`  |      Yes      | The message to be logged                                |
| `ident` |      Yes      | A grouping identifier, used to identify related entries |


