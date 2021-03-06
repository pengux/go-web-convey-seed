[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

# go-web-convey-seed
A Go seed project for building RESTful JSON API using [gocraft/web](https://github.com/gocraft/web) as router and [Goconvey](http://smartystreets.github.io/goconvey/) as test framework

## Features
* Modular structure with pluggable endpoints
* Base controller with common actions
* Base data service with common methods
* Testable using [Goconvey](http://smartystreets.github.io/goconvey/)

## Usage
1. Clone/download this repo into your project folder
2. Edit the "foo" endpoint (foo.go, foo_controller.go, foo_service.go, foo_test.go) with your own endpoint. You can create new endpoints based on these files.
3. Edit the dao.go and remove all test code (with comment "ONLY FOR TEST") and start implementing your own DB integration
4. Run your tests with Goconvey

### Changes
* 2013-12-27 The current dao.go mockup does not work with tests yet.
