# Scripting language

Lightweight go-like scripting language.

## Package

```go
var source, _ = os.ReadFile(path)
script.Compile(string(source)).Run(storage)
```

## Command

```sh
gs [args...]
```

## Syntax

```go
func () {}
new {}
if true {}
for true {}
{}

value ()
value . key
key = value

true
false
key
0.0
""
```

## Examples

- Root Directory
  - async.cs
  - closure.cs
  - main.cs
  - object.cs
  - sum.cs

```go
main = func () {
    s = slice(1.2 2.3)
    sum = this.sum(s)
    fmt.Println(sum)
    // Output:
    // 3.5
}
sum = func (s) {
    i = 0
    sum = 0
    for lt(i len(s)) {
        sum = add(sum get(s i))
        i = add(i 1)
    }
    return(sum)
}
async = func () {
    msgs = chan(0)
    go(func () {
        for true {
            msg = recv(this.msgs)
            if eq(msg "exit") {
                fmt.Println("server exited")
                break()
            }
            fmt.Println(msg)
        }
    })
    for true {
        msg = fmt.Scanln()
        if eq(msg "bye") {
            fmt.Println("client exited")
            break()
        }
        send(msgs msg)
    }
    // Output:
    // 123
    // 123
    // exit
    // server exited
    // bye
    // client exited
}
object = func () {
    User = func (id name) {
        user = new {
            id = this.id
            name = this.name
            work = func () {
                fmt.Println("I'm working, and my name is" this.name)
            }
        }
        return(user)
    }
    User(1 "Tom").work()
    t = time.Now()
    fmt.Println(t.Year())
    fmt.Println(t.Format("2006-01-02 15:04:05"))
    // Output:
    // I'm working, and my name is Tom
    // 2024
    // 2024-03-27 08:23:27
}
closure = func () {
    inc = func () {
        val = add(this.val 10)
        this.val = 30
        return(val)
    }
    val = 10
    fmt.Println(inc())
    fmt.Println(val)
    // Output:
    // 20
    // 30
}
```
