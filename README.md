# Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

### Contents

- [Awesome Go](#awesome-go)
    - [Audio & Music](#audiomusic)
    - [Authentication & OAuth](#authentication--oauth)
    - [CUI](#cui)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date & Time](#date--time)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Forms](#forms)
    - [Game Development](#game-development)
    - [Generation & Generics](#generation--generics)
    - [GUI](#gui)
    - [Images](#images)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Natural Language Processing](#natural-language-processing)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [E-Books](#e-books)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)



## Audio/Music

*Libraries for manipulating audio.*

* [flac](https://github.com/eaburns/flac) - A native Go FLAC decoder.
* [go-sox](https://github.com/krig/go-sox) - libsox bindings for go.
* [PortAudio](https://code.google.com/p/portaudio-go/) - Go bindings for the PortAudio audio I/O library.
* [portmidi](https://github.com/rakyll/portmidi) - Go bindings for PortMidi.
* [vorbis](https://github.com/mccoyst/vorbis) - A "native" Go Vorbis decoder (uses CGO, but has no dependencies).



## Authentication & OAuth

*Libraries for implementing authentications schemes.*

* [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth) - AWS (Amazon Web Services) request signing library.
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middleware.
* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT).
* [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.
* [osin](https://github.com/RangelReale/osin) - Golang OAuth2 server library.


## CUI

*Libraries for building CUI Applications*

* [color](https://github.com/fatih/color) Versatile package for colored terminal output.
* [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals.
* [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces.


## Database

*Databases implemented in Go.*

* [bolt](https://github.com/boltdb/bolt) - A low-level key/value database for Go.
* [diskv](https://github.com/peterbourgon/diskv) - A home-grown disk-backed key-value store.
* [go-cache](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
* [goleveldb](https://github.com/syndtr/goleveldb) - An implementation of the [LevelDB](https://code.google.com/p/leveldb/) key/value database in the Go.
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
* [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics
* [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
* [skydb.io](https://github.com/skydb/sky) - Sky is an open source database used for flexible, high performance analysis of behavioral data.
* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang.

## Database Drivers

*Libraries for connecting and operating databases. (see [Testing](#testing) for mock drivers)*

* Relational Databases
    * [go-pgsql](https://github.com/lxn/go-pgsql) - A PostgreSQL client package for the Go Programming Language.
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go.
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that using database/sql.
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql.

* NoSQL Databases
    * [cayley](https://github.com/google/cayley) - A graph database with support for multiple backends.
    * [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
    * [mgo](http://godoc.org/labix.org/v2/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    * [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang.
    * [redigo](https://github.com/garyburd/redigo) - Redigo is a Go client for the Redis database.
    * [redis](https://github.com/hoisie/redis) - A simple, powerful Redis client for Go.


## Date & Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang.


## Email

*Libraries that implement email creation and sending*

* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go.
* [Go-MailHog](https://github.com/ian-kent/Go-MailHog) - Catches mail and serves it through a dream. Inspired by MailCatcher, easier to install.
* [gomail](https://github.com/alexcesaro/mail) - Gomail provides a very simple API to send emails. It supports attachments, multipart emails and encoding of non-ASCII characters.


## Embeddable Scripting Languages

*Embedding other languages inside your go code*

* [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go
* [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API
* [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API
* [otto](https://github.com/robertkrimen/otto) - A JavaScript interpreter written in Go
* [v8-go](https://github.com/idada/v8.go/) - V8 JavaScript engine bindings for Go


## Forms

*Libraries for working with forms.*

* [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct.
* [nosurf](https://github.com/justinas/nosurf) - A CSRF protection middleware for Go.



## Game Development

*Awesome game development libraries.*

* [GarageEngine](https://github.com/vova616/GarageEngine) - 2d game engine written in Go working on OpenGL.
* [glop](https://github.com/runningwild/glop) - Glop (Game Library Of Power) is a fairly simple cross-platform game library.
* [go3d](https://github.com/ungerik/go3d) - A performance oriented 2D/3D math package for Go
* [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A* path finding algorithm


## Generation & Generics

*Tools to enhance the language with features like generics via code generation*

* [gen](https://github.com/clipperhouse/gen) - Code generation tool for ‘generics’-like functionality.
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go.
* [pkgreflect](https://github.com/ungerik/pkgreflect) - A Go preprocessor for package scoped reflection.


## GUI

*Libraries for building GUI Applications*

* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK
* [gotk3](https://github.com/conformal/gotk3) - Go bindings for GTK3.
* [go-qml](https://github.com/go-qml/qml) - QML support for the Go language
* [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go.


## Images

*Libraries for manipulating images.*

* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library.
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go.
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go.
* [img](https://github.com/hawx/img) - A selection of image manipulation tools.
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API.
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package.
* [resize](https://github.com/nfnt/resize) - Image resizing for the Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez) - Image resizing, functionality similar to resize
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation.


## Logging

*Libraries for generating and working with log files.*

* [glog](https://github.com/golang/glog) - Leveled execution logs for Go.
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers.
* [logrus](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
* [seelog](https://github.com/cihub/seelog) -   logging functionality with flexible dispatching, filtering, and formatting.
* [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.


## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang.
* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library.
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go
* [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go.
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang.
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14.
* [mlgo](https://code.google.com/p/mlgo/) - This project aims to provide minimalistic machine learning algorithms in Go.
* [neural-go](https://github.com/schuyler/neural-go) - A multilayer perceptron network implemented in Go, with training via backpropagation.
* [probab](https://code.google.com/p/probab/) - Probability distribution functions. Bayesian inference. Written in pure Go.
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go


## Messaging

*Libraries that implement messaging systems*

* [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus.
* [go-nsq](https://github.com/bitly/go-nsq) - the official Go package for NSQ
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster.
* [NATS](https://github.com/apcera/nats) - A lightweight and highly performant publish-subscribe and distributed queueing messaging system.
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - A redis backed unified push service for server-side notifications to mobile devices.



## Natural Language Processing

*Libraries for working with human languages.*

* [go-eco](https://code.google.com/p/go-eco/) - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.
* [go-porterstemmer](https://github.com/reiver/go-porterstemmer) - A native Go clean room implementation of the Porter Stemming algorithm.
* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm.
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
* [paicehusk](https://github.com/Rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm
* [porter](https://github.com/a2800276/porter) - This is a fairly straighforward port of Martin Porter's C implementation of the Porter stemming algorithm.
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers.
* [textcat](https://github.com/pebbe/textcat) - A Go package for n-gram based text categorization, with support for utf-8 and raw text


## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL. Requires an external dependency GLEW.
* [glfw3](https://github.com/go-gl/glfw3) - Go bindings for GLFW 3.
* [glow](https://github.com/errcw/glow) - Go binding generator and bindings for OpenGL.
* [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM.



## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [BeeDB](https://github.com/astaxie/beedb) - go ORM,support database/sql interface，pq/mysql/sqlite.
* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
* [gorp](https://github.com/coopernurse/gorp) - Go Relational Persistence, ORM-ish library for Go.
* [hood](https://github.com/eaigner/hood) - Database agnostic ORM for Go.
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM.
* [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go.


## Package Management

*Libraries for package and dependency management.*

* [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.
* [gom](https://github.com/mattn/gom) - Go Manager - bundle for go.
* [goop](https://github.com/nitrous-io/goop) - A simple dependency manager for Go (golang), inspired by Bundler.
* [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go.


## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [blas](https://github.com/ziutek/blas) - Implementation of BLAS (Basic Linear Algebra Subprograms)
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang
* [gocomplex](https://code.google.com/p/gocomplex/) - A complex number library for the Go programming language.
* [go-fn](https://code.google.com/p/go-fn/) - Mathematical functions written in Go language, that are not covered by math pkg
* [gofrac](https://github.com/anschelsc/gofrac) - A (goinstallable) fractions library for go with support for basic arithmetic.
* [go-gt](https://code.google.com/p/go-gt/) - Graph theory algorithms written in "Go" language
* [go.matrix](https://github.com/skelterjohn/go.matrix) - linear algebra for go
* [gostat](https://code.google.com/p/gostat/) - A statistics library for the go language
* [mudlark-go](https://code.google.com/p/mudlark-go-pkgs/) - A collection of packages providing (hopefully) useful code for use in software using Google's Go programming language.
* [plotinum](https://code.google.com/p/plotinum/) - Plotinum provides an API for building and drawing plots in Go.
* [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data.
* [vectormath](https://github.com/spate/vectormath) - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code.

## Template Engines

*Libraries and tools for templating and lexing.*

* [amber](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.
* [gold](https://github.com/yosssi/gold) - Gold is a template engine for Go. This simplifies HTML coding in Go web application development. This is influenced by Slim and Jade.
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation.
* [mustache](https://github.com/hoisie/mustache) - A Go implementation of the Mustache template language.
* [pongo2](https://github.com/flosch/pongo2) - A Django-like template-engine for Go.
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang.
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/)


## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/bmizerany/assert) - Asserts to Go testing
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go
    * [gocheck](http://labix.org/gocheck) - A more advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload
    * [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language.
    * [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
    * [restit](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test.
    * [Testify](https://github.com/stretchr/testify) - A sacred extension to the standard go testing package.

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects
    * [gomock](https://code.google.com/p/gomock/) - Mocking framework for the Go programming language.
    * [mockhttp.go](https://github.com/tv42/mockhttp.go) - Mock object for Go http.ResponseWriter
    * [go-sqlmock](github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database transactions


## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go
        * [github_flavored_markdown](http://godoc.org/github.com/shurcooL/go/github_flavored_markdown) - GitHub Flavored Markdown renderer in Go.
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer
    * [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format.
    * [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    * [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
    * [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support.
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection).
    * [yaml](https://bitbucket.org/zombiezen/yaml) - Implements a YAML 1.2 parser in Go.
* Utility


## Third-party APIs

*Libraries for accessing third party APIs.*

* [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API
* [github](https://github.com/google/go-github) - Go library for accessing the GitHub API.
* [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API.
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP.


## Utilities

*General utilities and tools to make your life easier.*

* [cli](https://github.com/codegangsta/cli) - A small package for building command line apps in Go.
* [coop](https://github.com/rakyll/coop) - Cheat sheet for some of the common concurrent flows in Go.
* [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go.
* [mp](https://github.com/sanbornm/mp) - A simple cli email parser. It currently takes stdin and outputs JSON.


## Video

*Libraries for manipulating video.*

* [aac/h264](https://github.com/go-av/codec) - Golang aac/h264 encoder and decoder.
* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries.
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer.


## Web Frameworks

*Full stack web frameworks.*

* [Beego](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language.
* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
* [gocraft/web](https://github.com/gocraft/web) - A mux and middleware package in Go.
* [Goji](https://github.com/zenazn/goji) - Goji is a minimalistic web framework for Golang that's high in antioxidants.
* [Gorilla](https://github.com/gorilla/) - Gorilla is a web toolkit for the Go programming language.
* [go-rest](https://github.com/ungerik/go-rest) - A small and evil REST framework for Go
* [httprouter](https://github.com/julienschmidt/httprouter) - A high performance router. Use this and the standard http handlers to form a very high performance web framework.
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
* [Martini](https://github.com/go-martini/martini) - Martini is a powerful package for quickly writing modular web applications/services in Golang.
* [pat](https://github.com/bmizerany/pat) - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.
* [Revel](https://github.com/revel/revel) - A high-productivity web framework for the Go language.
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - A Go framework for building JSON web services inspired by Dropwizard
* [traffic](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go.
* [web.go](https://github.com/hoisie/web) - A simple framework to write webapps in Go.


### Middlewares

*Libraries for creating HTTP middlewares.*

* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go.
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http.
* [negroni](https://github.com/codegangsta/negroni) - Idiomatic HTTP Middleware for Golang.
* [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses.









# Tools

Go software and plugins.

## Code Analysis

* [doc](http://godoc.org/code.google.com/p/rspace.cmd/doc) - Go documentation tool that produces an alternative doc format.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer.
* [goimports](http://godoc.org/code.google.com/p/go.tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code.
* [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags.


## Editor Plugins

* [go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) Go plugin for IntelliJ IDEA.
* [gocode](https://github.com/nsf/gocode) - An autocompletion daemon for the Go programming language
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - A Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features.
* [velour](https://github.com/velour/velour) - An IRC client for the acme editor.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - A Vim plugin to highlight syntax errors on save.
* [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim.
* [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes.

## Software Packages

Software written in Go.

### DevOps Tools

* [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool
* [Boom](https://github.com/rakyll/boom) - Boom is a tiny program that sends some load to a web application.
* [gaudi](http://gaudi.io/) - Gaudi automates the setup of isolated and decoupled dev environments.
* [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go.
* [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application.
* [Gogs](http://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update.
* [gox](https://github.com/mitchellh/gox) - A dead simple, no frills Go cross compile tool.
* [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging.
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions.
* [hk](https://github.com/heroku/hk) - Heroku command-line interface in Go.
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data.

### Other Software

* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul.
* [Docker](http://www.docker.com/) - An open platform for distributed applications for developers and sysadmins.
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery.
* [fleet](https://github.com/coreos/fleet) - A Distributed init System.
* [juju](https://juju.ubuntu.com/) - Configure, deploy, manage and monitor Openstack infrastructure.
* [nsq](http://nsq.io/) - A realtime distributed messaging platform
* [Postman](https://github.com/zachlatta/postman) - Command-line utility for batch-sending email.
* [tsuru](http://www.tsuru.io/) - An extensible and open source Platform as a Service software.











# Resources

Where to discover new Go libraries.

## Benchmarks

* [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions.
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods.
* [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) - Naive performance test of two ways to do type assertion in Go.
* [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language.

## E-Books

* [build-applications-web](https://docs.google.com/file/d/0B2GBHFyTK2N8TzM4dEtIWjBJdEk/edit?pli=1)
* [golang-book](http://www.golang-book.com/)
* [golangbootcamp](http://golangbootcamp.com)
* [learning-go](http://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [network-programming](http://jan.newmarch.name/go/)

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangweekly](https://twitter.com/golangweekly)


## Websites

* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists.
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - A collection of Go articles and tutorials.
* [godoc.org](http://godoc.org/) - Documentation for open source Go packages.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list
* [Go Projects](https://code.google.com/p/go-wiki/wiki/Projects) - List of projects on the Go community wiki
* [r/Golang](http://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [go generate](https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit) - A proposed code generation syntax from Rob Pike.

### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go
* [Go By Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated example programs
* [Working with Go](https://github.com/mkaz/working-with-go) - An intro to go for experienced programmers

