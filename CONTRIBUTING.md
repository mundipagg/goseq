
### Pull Request Labels

This section show all the labels that should be used on pull request messages.
Each label will represent a section in the Changelog file.

| Label |  Description |
| --- | --- |
| `enhancement` |  New features or improvements |
| `bug` | Bugfix |
|`documentation`| Documentation update |
| `fire` |  Code removal |

## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally
* Consider starting the commit message with an applicable emoji:

| Código                | Emoji               | Descrição                                       |
|-----------------------|---------------------|-------------------------------------------------|
| `:art:`               | :art:               | when improving the format/structure of the code |
| `:racehorse:`         | :racehorse:         | when improving performance                      |
| `:non-potable_water:` | :non-potable_water: | when plugging memory leaks                      |
| `:memo:`              | :memo:              | when writing docs                               |
| `:checkered_flag:`    | :checkered_flag:    | when fixing something on windows                |
| `:bug:`               | :bug:               | when fixing a bug                               |
| `:fire:`              | :fire:              | when removing code or files                     |
| `:green_heart:`       | :green_heart:       | when fixing CI build                            |
| `:white_check_mark:`  | :white_check_mark:  | when adding tests                               |
| `:lock:`              | :lock:              | when dealing with security                      |
| `:arrow_up:`          | :arrow_up:          | when upgrading dependencies                     |
| `:arrow_down:`        | :arrow_down:        | when downgrading dependencies                   |
| `:shirt:`             | :shirt:             | when removing linter warnings                   |
| `:bulb:`              | :bulb:              | new idea                                        |
| `:construction:`      | :construction:      | work in progress                                |
| `:heavy_plus_sign:`   | :heavy_plus_sign:   | when adding feature                             |
| `:heavy_minus_sign:`  | :heavy_minus_sign:  | when removing feature                           |
| `:speaker:`           | :speaker:           | when adding logging                             |
| `:mute:`              | :mute:              | when removing logging                           |
| `:facepunch:`         | :facepunch:         | when resolving conflicts                        |
| `:wrench:`            | :wrench:            | when changing configuration files               |


### Tests Styleguide

* Our tests doesn't have a defined standard, we just keep methods with similar signatures


#### Example
```go
  func Test_TestCenario_SuccessOrFail(t *testing.T) {
      //...
  }
```


### References

[Atom CONTRIB](https://github.com/atom/atom/blob/master/CONTRIBUTING.md)