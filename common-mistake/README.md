
| Mistake Title                     | One-line explanation                                                                                              | Example            |
| --------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ------------------ |
| Variable shadowing in inner scope | Using `:=` inside an inner block can redeclare and hide an outer variable, leading to unexpected unchanged state. | [example](./001.go) |
| Unnecessary nested code | Deep `if-else` chains can be simplified using ordered returns, making logic easier to read and maintain. | [example](./002.go) |


