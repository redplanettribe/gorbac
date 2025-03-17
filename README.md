# gorbac - Go Role-Based Access Control

## Overview
The `gorbac` library provides a flexible and easy-to-use role-based access control (RBAC) system for managing permissions and roles in your Go applications. It allows you to define permissions through a fluent configuration API and integrates seamlessly with your application's authorization logic.

## Installation

To install the library, run the following command:
```
go get github.com/redplanettribe/gorbac
```


## Usage

### Define Permissions Configuration

Create a permissions configuration as shown below:

```go
package main

import "github.com/redplanettribe/gorbac"

func main () {
    p := gorbac.NewPermissions().
        AddRole("user").
        /* */Read("users").
        /* */Write("users").
        /* */Write("projects").
        /* */Read("projects").
        /* */Delete("projects").
        /* */Read("publishers").
        /* */Read("posts").
        AddRole("admin").Inherit("user").
        /* */Read("roles").
        /* */Write("roles").
        /* */Delete("roles")
    authorizator := gorbac.NewAuthorizer(p)
    
}
```

## Versioning

This library follows semantic versioning. For production use, we recommend pinning to a specific version:

```go
import (
    gorbac "github.com/redplanettribe/gorbac/<version>"
)
```

