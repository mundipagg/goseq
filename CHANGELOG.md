# Changelog

## 0.1.2 (2017-04-30)

Added
- Adds support to n (n > 0) http Seq consumer at Logger to prevent the application to consume a lot of memory when we have a thousands of requests
- Change SeqClient.Send return type, from bool to error

## 0.1.1 (2017-03-17)

Added
- Adds error return when `Logger` is created 

## 0.1.0 (2017-03-17)

Added
- Initial release
