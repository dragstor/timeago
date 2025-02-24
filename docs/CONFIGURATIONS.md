[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Description](#description)
- [Language](#language)
- [Location](#location)

# ⚙️ Configurations

## Description

We can set package configurations with `SetConfig` function that accepts `Config` structure.

## Language

Optionally you can set the language in your application.

> Default value is `"en"`

```go
import . "github.com/SerhiiCho/timeago"

func main() {
    SetConfig(Config{
        Language: "ru",
    })
}
```

## Location

Optionally you can set the location in your application. The functionality is going to be the
same, but with 1 exception. In the absence of time zone information, package interprets a time as UTC;
With location configuration it interprets the time as in the given location.

> There is no default value for this option.

```go
import . "github.com/SerhiiCho/timeago"

func main() {
    SetConfig(Config{
        Location: "Europe/Kiev",
    })
}
```

> Please make sure that timezone configuration is correct for your location. It is very important for displaying the correct datetime.

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)