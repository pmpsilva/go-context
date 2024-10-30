# Go context examples

This repository contains several examples demonstrating the usage of the Go package context.

The project will feature multiple branches, each showcasing a specific example.

##  Context WithTimeout

In this example, we are utilizing `context.WithTimeout(ctx, 200*time.Millisecond)` to create a context with a timeout. During the call, we deliberately make the `fetchExternalApi(identifier)` function sleep longer than the specified timeout. This allows us to simulate a slow API call and observe how the context timeout operates in practice.

### Usage
```bash
go run main.go 
```

**Expected Result**
```bash
2024/10/30 16:22:05 request imeout
exit status 1
```

*To ensure a successful call, use the `fetchDate(ctx, id)` function with the id parameter instead of slowId in the main section.*