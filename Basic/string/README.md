## String in GO
### Declaring Go string queries
    str := "hello"
    var str = "hello"
    var str string = "hello"
If you declare a variable without specifying a value,
it will be initialized to zero.
The null value of type string is empty quotes (""):

    var str string

### Is it possible to change strings in Golang?
Strings in Golang are not meant to be manipulated.
You can assign another string to a variable, but
the strings themselves cannot be edited:

    str := "hello"
    str = "home" // NO
    str[0] = 'q' // NO

### Decoding strings into runes

- Strings use an encoding called UTF-8, where each character requires
  from 1 to 4 bytes;
- byte is another name for type uint8, and rune is another name
  type int32;
- The range keyword can decode a UTF-8 encoded string into runs. 

### An efficient way to concatenate strings


    func concat2builder(x, y string) string {
      var builder strings.Builder
      builder.Grow(len(x) + len(y)) // this line allocates memory
      builder.WriteString(x)
      builder.WriteString(y)
      return builder.String()
    }