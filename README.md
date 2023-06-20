<p align="center">
<img src="./assets/photos/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/erfanmomeniii/lfu?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# lfu

<i>lfu</i> is a lightweight package that implements [lfu](https://en.wikipedia.org/wiki/Least_frequently_used) cache algorithm in Go.

# Documentation

## Install

```bash
go get github.com/erfanmomeniii/lfu
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/lfu"
``` 
## Quick Start
The following example illustrates how to use this package
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/lru"
)

func main() {
	l := lru.New(2)
	// 2 is the size of lru (default 2000)
	
	l.Set("a", 12)
	l.Set("b", 13)
	
	fmt.Println(l.Get("a"))
	// 12
	
	l.Set("c", "hi")
	
	fmt.Println(l.Has("c"))
	// false
}

```
## Usage

#### type Lfu
Lfu is an instantiation of the lfu.

#### func New
```go
func New(size ...int64) *Lfu
```
New creates a new instance of a lfu.

#### func Set
```go
func (l *Lfu) Set(key any, value any)
```
Set adds or updates with inputted key and value.

#### func Get
```go
func (l *Lfu) Get(key any) any
```
Get returns the value associated with the inputted key.

#### func Has
```go
func (l *Lfu) Has(key any) bool 
```
Has checks whether the key exists or not.

## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.



