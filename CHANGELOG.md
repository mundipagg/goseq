# Changelog

## 0.1.2 (2017-04-30)

Added
- Adds support to n (n > 0) http Seq consumer at Logger to prevent the application to consume a lot of memory when we have a thousands of requests
- Change SeqClient.Send return type, from bool to error
- Assert how many consumers was passed by user to the Logger
- Added support to complex objects on log, now the properties are of type `string`/`interface{}` instead of `string`/`string`

Changed
- Changes method name `DefaultProperies` to `SetDefaultProperties`
- Changes not necessary public variables to private ones

## 0.1.1 (2017-03-17)

Added
- Adds error return when `Logger` is created 

## 0.1.0 (2017-03-17)

Added
- Initial release
