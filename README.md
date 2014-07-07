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
    - [Natural Language Processing](#natural-language-processing)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Machine Learning](#machine-learning)
    - [Testing](#testing)
    - [Audio](#audio)
    - [Video](#video)
    - [Game Development](#game-development)
    - [DevOps Tools](#devops-tools)
    - [Utilities](#utilities)
    - [Logging](#logging)
    - [Code Analysis and Linter](#code-analysis-and-linter)
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
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http

## Template Engine

*Libraries and tools for templating and lexing.*

* [mustache](https://github.com/hoisie/mustache) - A Go implementation of the Mustache template language.
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation.
* [gold](https://github.com/yosssi/gold) - Gold is a template engine for Go. This simplifies HTML coding in Go web application development. This is influenced by Slim and Jade.
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang.
* [pongo2](https://github.com/flosch/pongo2) - A Django-like template-engine for Go.
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/)


## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) - A CSRF protection middleware for Go.


## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [goauth](http://alloy-d.net/goauth/) - A Go library for doing header-based OAuth over HTTP or HTTPS. Mostly created for working with Twitter.
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middlewares.
* [osin](https://github.com/RangelReale/osin) - Golang OAuth2 server library.
* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT).

## Database

*Databases implemented in Go.*

* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang.
* [diskv](https://github.com/peterbourgon/diskv) - A home-grown disk-backed key-value store.
* [bolt](https://github.com/boltdb/bolt) - A low-level key/value database for Go.
* [go-cache](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
* [goleveldb](https://github.com/syndtr/goleveldb) - An implementation of the [LevelDB](https://code.google.com/p/leveldb/) key/value database in the Go.
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
* [skydb.io](http://skydb.io/) - Sky is an open source database used for flexible, high performance analysis of behavioral data.

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
* [Go-MailHog](https://github.com/ian-kent/Go-MailHog) - Catches mail and serves it through a dream. Inspired by MailCatcher, easier to install.

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
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection).
    * [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
    * [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go
        * [github_flavored_markdown](http://godoc.org/github.com/shurcooL/go/github_flavored_markdown) - GitHub Flavored Markdown renderer in Go.
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer


## Natural Language Processing

*Libraries for working with human languages.*

* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm.
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality.
* [paicehusk](https://github.com/Rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm
* [go-porterstemmer](https://github.com/reiver/go-porterstemmer) - A native Go clean room implementation of the Porter Stemming algorithm.
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers.
* [snowball](https://github.com/kljensen/snowball) - Go implementation of the Snowball stemmers
* [porter](https://github.com/a2800276/porter) - This is a fairly straighforward port of Martin Porter's C implementation of the Porter stemming algorithm. 
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2
* [snowball](https://github.com/tebeka/snowball) - Snowball Stemmer for Go [Snowball native](http://snowball.tartarus.org/)
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
* [textcat](https://github.com/pebbe/textcat) - A Go package for n-gram based text categorization, with support for utf-8 and raw text
* [go-eco](https://code.google.com/p/go-eco/) - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.


## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [go.matrix](https://github.com/skelterjohn/go.matrix) - linear algebra for go
* [gocomplex](https://code.google.com/p/gocomplex/) - A complex number library for the Go programming language.
* [mudlark-go](https://code.google.com/p/mudlark-go-pkgs/) - A collection of packages providing (hopefully) useful code for use in software using Google's Go programming language.
* [gostat](https://code.google.com/p/gostat/) - A statistics library for the go language
* [gofrac](https://github.com/anschelsc/gofrac) - A (goinstallable) fractions library for go with support for basic arithmetic.
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang
* [blas](https://github.com/ziutek/blas) - Implementation of BLAS (Basic Linear Algebra Subprograms)
* [go-fn](https://code.google.com/p/go-fn/) - Mathematical functions written in Go language, that are not covered by math pkg
* [go-gt](https://code.google.com/p/go-gt/) - Graph theory algorithms written in "Go" language
* [vectormath](https://github.com/spate/vectormath) - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code.


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

* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries.
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer.
* [aac/h264](https://github.com/go-av/codec) - Golang aac/h264 encoder and decoder.


## Game Development

*Awesome game development libraries.*

* [GarageEngine](https://github.com/vova616/GarageEngine) - 2d game engine written in Go working on OpenGL.
* [fungo](https://github.com/beoran/fungo) - Fun Unified Game library for te gO Programming language.
* [go-rpg](https://github.com/viking/go-rpg) - Go package for creating role playing games
* [terrago](https://github.com/sarenji/terrago) - Fractal terrain generator in Go.
* [rog](https://github.com/ajhager/rog/) - A roguelike game library written in go
* [glop](https://github.com/runningwild/glop) - Glop (Game Library Of Power) is a fairly simple cross-platform game library.


## DevOps Tools

*Software and libraries for DevOps.*

* [Docker](http://www.docker.com/) - An open platform for distributed applications for developers and sysadmins.
* [juju](https://juju.ubuntu.com/) - Automate your cloud infrastructure
* [tsuru](http://www.tsuru.io/) - An extensible and open source Platform as a Service software.
* [Gogs](http://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
* [gaudi](http://gaudi.io/) - Gaudi automates the setup of isolated and decoupled dev environments.
* [fleet](https://github.com/coreos/fleet) - A Distributed init System.
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul.
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery.

## Utilities

*General utilities and tools to make you're life easier.*

* [Postman](https://github.com/zachlatta/postman) - Command-line utility for batch-sending email.
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions.

## Logging

*Libraries for generating and working with log files.*

* [glog](https://github.com/golang/glog) - Leveled execution logs for Go.
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers.
* [logrus](https://github.com/sirupsen/logrus) - Structured, pluggable logging
  for Go.


## Code Analysis and Linter

*Libraries and tools for analysing, parsing and manipulation codebases.*

* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code.


# Resources

Where to discover new Go libraries.

## Websites

* [Go Projects](https://code.google.com/p/go-wiki/wiki/Projects) - List of projects on the Go community wiki
* [godoc.org](http://godoc.org/) - Documentation for open source Go packages.
* [r/Golang](http://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - A collection of Go articles and tutorials.

### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go
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
