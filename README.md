# Clean Architecture in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/gocleanarchitecture)](https://goreportcard.com/report/github.com/josephspurrier/gocleanarchitecture)
[![GoDoc](https://godoc.org/github.com/josephspurrier/gocleanarchitecture?status.svg)](https://godoc.org/github.com/josephspurrier/gocleanarchitecture)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/gocleanarchitecture/badge.svg?branch=master&randid=6)](https://coveralls.io/github/josephspurrier/gocleanarchitecture?branch=master)

A good example of clean architecture for a web application in Go.

The **domain** folder is for **enterprise** business logic without any
dependencies. These can be structs, interfaces, and functions.

The **usecase** folder is for **application** business logic without any
dependencies with the exception of the domain logic. These can be structs,
interfaces, and functions. There is no usecase folder in this example.

The **adapter** folder should contain abstractions for the packages in the
**lib** and **vendor** folders.

The **lib** folder contains internal packages, similar to the **vendor** folder
which contains 3rd party packages.