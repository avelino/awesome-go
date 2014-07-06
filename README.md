# Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).


- [Awesome Go](#awesome-go)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares) 
    - [Template Engine](#template-engine)
    - [Forms](#forms)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Email](#email)
    - [ORM](#orm)
    - [Imagery](#imagery)
    - [Text Processing](#text-processing)
    - [Machine Learning](#machine-learning)
    - [Testing](#testing)
    - [Audio](#audio)
    - [Video](#video)
    - [DevOps Tools](#devops-tools)
- [Resources](#resources)
    - [Websites](#websites)
    - [(e)Books](#ebooks)

## Web Frameworks

*Full stack web frameworks.*

* [Martini](http://martini.codegangsta.io/) - Martini is a powerful package for quickly writing modular web applications/services in Golang.
* [Gorilla](http://www.gorillatoolkit.org/) - Gorilla is a web toolkit for the Go programming language.
* [Gin](http://gin-gonic.github.io/gin/) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
* [Goji](https://goji.io) - Goji is a minimalistic web framework for Golang that's high in antioxidants.
* [web.go](http://webgo.io/) - A simple framework to write webapps in Go.
* [pat](https://github.com/bmizerany/pat) - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.
* [Revel](http://revel.github.io/) - A high-productivity web framework for the Go language.
* [Beego](http://beego.me/) - beego is an open-source, high-performance web framework for the Go programming language.
* [traffic](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go.
* [httprouter](https://github.com/julienschmidt/httprouter) - A high performance router. Use this and the standard http handlers to form a very high performance web framework.
* [gocraft/web](https://github.com/gocraft/web) - A mux and middleware package in Go. 
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.


### Middlewares

* [negroni](https://github.com/codegangsta/negroni) - Idiomatic HTTP Middleware for Golang
* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go

## Template Engine

*Libraries and tools for templating and lexing.*

* [mustache](https://github.com/hoisie/mustache) - A Go implementation of the Mustache template language.
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation.
* [gold](https://github.com/yosssi/gold) - Gold is a template engine for Go. This simplifies HTML coding in Go web application development. This is influenced by Slim and Jade.
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang.
* [Pongo](https://github.com/flosch/pongo) - A Django-like template engine for Go.
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/)


## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) - A CSRF protection middleware for Go.


## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [goauth](http://alloy-d.net/goauth/) - A Go library for doing header-based OAuth over HTTP or HTTPS. Mostly created for working with Twitter.
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middlewares.


## Database

*Databases implemented in Go.*

* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang.
* [diskv](https://github.com/peterbourgon/diskv) - A home-grown disk-backed key-value store.
* [bolt](https://github.com/boltdb/bolt) - A low-level key/value database for Go.
* [go-cache](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.

## Database Drivers

*Libraties for connecting and operating databases.*

* Relational Databases
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql.
    * [go-pgsql](https://github.com/lxn/go-pgsql) - A PostgreSQL client package for the Go Programming Language.
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go.
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that using database/sql.
    * [go-db](https://github.com/phf/go-db) - Generic database API for Go.
* NoSQL Databases
    * [mgo](http://labix.org/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    * [redis](https://github.com/hoisie/redis) - A simple, powerful Redis client for Go.
    * [redigo](https://github.com/garyburd/redigo) - Redigo is a Go client for the Redis database.
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang.
    * [gocouch](https://github.com/hoisie/gocouch) - Couchdb client for Go.
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.

## Email

*Libraries that implement email creation and sending*

* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go.

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [BeeDB](https://github.com/astaxie/beedb) - go ORM,support database/sql interface，pq/mysql/sqlite.
* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
* [gorp](https://github.com/coopernurse/gorp) - Go Relational Persistence, ORM-ish library for Go.
* [hood](https://github.com/eaigner/hood) - Database agnostic ORM for Go.
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM.
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go.
* [upper.io/db](https://upper.io/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.


## Imagery

*Libraries for manipulating images.*

* [img](https://github.com/hawx/img) - A selection of image manipulation tools.
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation.
* [resize](https://github.com/nfnt/resize) - Image resizing for the Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez) - Image resizing, functionality similar to resize

## Text Processing

* Specific Formats
    * [yaml](https://bitbucket.org/zombiezen/yaml) - Implements a YAML 1.2 parser in Go.
    * [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
    * [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer


## Machine Learning

*Libraries for Machine Learning.*

* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
* [mlgo](https://code.google.com/p/mlgo/) - This project aims to provide minimalistic machine learning algorithms in Go.
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library.
* [neural-go](https://github.com/schuyler/neural-go) - A multilayer perceptron network implemented in Go, with training via backpropagation.
* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang.
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go
* [probab](https://code.google.com/p/probab/) - Probability distribution functions. Bayesian inference. Written in pure Go.
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14.
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang.
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang


## Testing

*Libraries for testing codebases and generating test data.*

* [gocheck](http://labix.org/gocheck) - A more advanced testing framework alternative to gotest.
* [GoConvey](http://goconvey.co/) - BDD-ish, rspec inspirated testing framework, automatic testing, coverage report and web UI
* [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language.
* [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
* [gomock](https://code.google.com/p/gomock/) - Mocking framework for the Go programming language.
* [mockhttp.go](https://github.com/tv42/mockhttp.go) - Mock object for Go http.ResponseWriter
* [assert](https://github.com/bmizerany/assert) - Asserts to Go testing
* [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
* [restit](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test.
* [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go


## Audio

*Libraries for manipulating audio.*

* [PortAudio](https://code.google.com/p/portaudio-go/) - Go bindings for the PortAudio audio I/O library.
* [gosndfile](https://github.com/mkb218/gosndfile) - Go bindings for libsndfile.
* [go-sox](https://github.com/krig/go-sox) - libsox bindings for go.


## Video
*Libraries for manipulating video.*
* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av* libraries.
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer.
* [aac/h264](https://github.com/go-av/codec) - Golang aac/h264 encoder and decoder.


## DevOps Tools

*Software and libraries for DevOps.*

* [Docker](http://www.docker.com/) - An open platform for distributed applications for developers and sysadmins.
* [juju](https://juju.ubuntu.com/) - Automate your cloud infrastructure
* [tsuru](http://www.tsuru.io/) - An extensible and open source Platform as a Service software.


## Logging

*Libraries for generating and working with log files.*

* [glog](https://github.com/golang/glog) - Leveled execution logs for Go.
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers.


# Resources

Where to discover new Go libraries.

## Websites

* [r/Golang](http://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - A collection of Go articles and tutorials.

### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go
* [An Introduction to Programming in Go](http://www.golang-book.com/) - A
  beginner's introduction to programming which uses the Go language
* [Working with Go](https://github.com/mkaz/working-with-go) - An intro to go for experienced programmers
* [Go By Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated example programs


## Twitter

* [@golang_news](https://twitter.com/golang_news)
* [@golangweekly](https://twitter.com/golangweekly)

## (e)Books

* [golang-book](http://www.golang-book.com/)
* [golangbootcamp](http://golangbootcamp.com)
* [network-programming](http://jan.newmarch.name/go/)
* [learning-go](http://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [build-applications-web](https://docs.google.com/file/d/0B2GBHFyTK2N8TzM4dEtIWjBJdEk/edit?pli=1)

# Contributing

Your contributions are always welcome!
