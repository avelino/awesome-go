# Awesome Go [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)


A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).


### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

[Join us on Slack](https://gophers.slack.com/messages/awesome/) to chat with other awesome-go maintainers!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*


### Contents

- [Awesome Go](#awesome-go)
    - [Audio & Music](#audiomusic)
    - [Authentication & OAuth](#authentication--oauth)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date & Time](#date--time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Game Development](#game-development)
    - [Generation & Generics](#generation--generics)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Miscellaneous](#miscellaneous)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [Validation](#validation)
    - [Version Control](#version-control)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
    - [Windows](#windows)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Server Applications](#server-applications)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
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
* [taglib](https://github.com/wtolson/go-taglib) - Go bindings for taglib.
* [vorbis](https://github.com/mccoyst/vorbis) - A "native" Go Vorbis decoder (uses CGO, but has no dependencies).
* [waveform](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams.


## Authentication & OAuth

*Libraries for implementing authentications schemes.*

* [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth) - AWS (Amazon Web Services) request signing library.
* [go-jose](https://github.com/square/go-jose) - A fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
* [go.auth](https://github.com/bradrydzewski/go.auth) - Authentication API for Go web applications.
* [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang.
* [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple provides out of the box.
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middleware.
* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT).
* [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.
* [osin](https://github.com/RangelReale/osin) - Golang OAuth2 server library.
* [permissions2](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.
* [yubigo](https://github.com/GeertJohan/yubigo) - a Yubikey client package that provides a simple API to integrate the Yubico Yubikey into a go application.


## Command Line


### Standard CLI

*Libraries for building standard or basic Command Line applications*

* [cli-init](https://github.com/tcnksm/cli-init) - The easy way to start building Golang command line application.
* [cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [codegangsta/cli](https://github.com/codegangsta/cli) - A small package for building command line apps in Go.
* [kingpin](https://github.com/alecthomas/kingpin) - A command line and flag parser supporting sub commands.
* [liner](https://github.com/peterh/liner) - A Go readline-like library for command-line interfaces.
* [mitchellh/cli](https://github.com/mitchellh/cli) - A Go library for implementing command-line interfaces.
* [ukautz/clif](https://github.com/ukautz/clif) - A small command line interface framework.


### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces*

* [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output.
* [color](https://github.com/fatih/color) - Versatile package for colored terminal output.
* [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals.
* [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces.
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
* [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces.
* [termtables](https://github.com/apcera/termtables) - A Go port of the Ruby library [terminal-tables](https://github.com/visionmedia/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output
* [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).


## Configuration

*Libraries for configuration parsing*

* [config](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing.
* [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults).
* [envcfg](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs.
* [envconf](https://github.com/ian-kent/envconf) - Configuration from environment
* [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables.
* [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy
* [ini](https://github.com/go-ini/ini) - Go package for read and write INI files
* [mini](https://github.com/FogCreek/mini) - A golang package for parsing ini-style configuration files
* [viper](https://github.com/spf13/viper) - Go configuration with fangs

## Continuous Integration

*Tools for help with continuous integration*

* [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system.
* [overalls](https://github.com/bluesuncorp/overalls) - Multi-Package go project coverprofile for tools like goveralls

## CSS Preprocessors

*Libraries for preprocessing CSS files*

* [c6](https://github.com/c9s/c6) - High performance SASS compatible-implementation compiler written in Go
* [gcss](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [bitset](https://github.com/willf/bitset) - Go package implementing bitsets.
* [bloom](https://github.com/surge/bloom) - Bloom filters implemented in Go.
* [boomfilters](https://github.com/tylertreat/BoomFilters) - probabilistic data structures for processing continuous, unbounded streams
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
* [encoding](https://github.com/surge/encoding) - Integer Compression Libraries for Go.
* [go-datastructures](https://github.com/Workiva/go-datastructures) - a collection of useful, performant, and thread-safe data structures
* [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
* [goskiplist](https://github.com/ryszard/goskiplist) - A skip list implementation in Go.
* [mafsa](https://github.com/smartystreets/mafsa) - MA-FSA implementation with Minimal Perfect Hashing
* [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go
* [trie](https://github.com/derekparker/trie) - Trie implementation in Go
* [ttlcache](https://github.com/diegobernardes/ttlcache) - An in-memory LRU string-interface{} map with expiration for golang
* [willf/bloom](https://github.com/willf/bloom) - Go package implementing Bloom filters.

## Database

*Databases implemented in Go.*

* [bolt](https://github.com/boltdb/bolt) - A low-level key/value database for Go.
* [cache2go](https://github.com/muesli/cache2go) - An in-memory key:value cache which supports automatic invalidation based on timeouts.
* [couchcache](https://github.com/codingsince1985/couchcache) - A RESTful caching micro-service backed by Couchbase server.
* [diskv](https://github.com/peterbourgon/diskv) - A home-grown disk-backed key-value store.
* [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB.
* [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC.
* [go-cache](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
* [goleveldb](https://github.com/syndtr/goleveldb) - An implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in the Go.
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
* [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics
* [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
* [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB.
* [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1.
* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang.

*Database tools.*

* [go-mysql](https://github.com/siddontang/go-mysql) - A go toolset to handle MySQL protocol and replication.
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically.
* [goose](https://bitbucket.org/liamstask/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
* [kingshard](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang.
* [myreplication](https://github.com/2tvenom/myreplication) - MySql binary log replication listener. Support statement and row based replication.
* [orchestrator](https://github.com/outbrain/orchestrator) - MySQL replication topology manager & visualizer
* [pgweb](https://github.com/sosedoff/pgweb) - A web-based PostgreSQL database browser
* [pravasan](https://github.com/pravasan/pravasan) - Simple Migration tool - currently for MySQL but planning to support soon for Postgres, SQLite, MongoDB, etc.,
* [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata.
* [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.

*SQL query builder, libraries for building and using SQL.*

* [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use it with ease.
* [goqu](https://github.com/doug-martin/goqu) - An idiomatic SQL builder and query library.
* [scaneo](https://github.com/variadico/scaneo) - Generate Go code to convert database rows into arbitrary structs.
* [Squirrel](https://github.com/lann/squirrel) - Go library that helps you build SQL queries.


## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go
    * [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that using database/sql.
    * [go-bqstreamer](https://github.com/rounds/go-bqstreamer) - BigQuery fast and concurrent stream insert.
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver prototype in go language.
    * [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that using database/sql.
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go.
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that using database/sql.
    * [gofreetds](https://github.com/minus5/gofreetds) Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql.

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language.
    * [cayley](https://github.com/google/cayley) - A graph database with support for multiple backends.
    * [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go
    * [go-couchdb](https://github.com/fjl/go-couchdb) - Yet another CouchDB HTTP API wrapper for Go
    * [gocb](https://github.com/couchbaselabs/gocb) - Official Couchbase Go SDK
    * [gocql](http://gocql.github.io) - A Go language driver for Apache Cassandra.
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
    * [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB
    * [mgo](http://godoc.org/labix.org/v2/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    * [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang.
    * [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang
    * [redigo](https://github.com/garyburd/redigo) - Redigo is a Go client for the Redis database.
    * [redis](https://github.com/go-redis/redis) - Redis client for Golang
    * [redis](https://github.com/hoisie/redis) - A simple, powerful Redis client for Go.

* Search and Analytic Databases
    * [bleve](https://github.com/blevesearch/bleve) - A modern text indexing library for go.
    * [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Google Go.
    * [elastigo](https://github.com/mattbaird/elastigo) - A Elasticsearch client library.
    * [goes](https://github.com/belogik/goes) - A library to interact with Elasticsearch.

## Date & Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang.
* [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.


## Distributed Systems

*Packages that help with building Distributed Systems.*

* [flowgraph](https://github.com/vectaport/flowgraph) - MPI-style ready-send coordination layer.
* [go-jump](https://github.com/dgryski/go-jump) - A port of Google's "Jump" Consistent Hash function.
* [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load.
* [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC.
* [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp.
* [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package.
    * [dht](https://godoc.org/github.com/anacrolix/torrent/dht) - BitTorrent Kademlia DHT implementation.

## Email

*Libraries that implement email creation and sending*

* [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails.
* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go.
* [go-dkim](https://github.com/toorop/go-dkim) - A DKIM library, to sign & verify email.
* [Gomail](https://github.com/go-gomail/gomail/) - Gomail is a very simple and powerful package to send emails.
* [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface
* [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email
* [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine



## Embeddable Scripting Languages

*Embedding other languages inside your go code*

* [agora](https://github.com/PuerkitoBio/agora) - Dynamically typed, embeddable programming language in Go
* [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go
* [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go
* [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go
* [go-lua](https://github.com/Shopify/go-lua) - A port of the Lua 5.2 VM to pure Go
* [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API
* [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API
* [gopher-lua](https://github.com/yuin/gopher-lua) - a Lua 5.1 VM and compiler written in Go
* [otto](https://github.com/robertkrimen/otto) - A JavaScript interpreter written in Go
* [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go


## Financial

*Packages for accounting and finance*

* [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang
* [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers


## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind)  - Bind form data to any Go values
* [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct.
* [formam](https://github.com/monoculum/formam) - decode form's values into a struct.
* [forms](https://github.com/albrow/forms) - A framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
* [nosurf](https://github.com/justinas/nosurf) - A CSRF protection middleware for Go.


## Game Development

*Awesome game development libraries.*

* [GarageEngine](https://github.com/vova616/GarageEngine) - 2d game engine written in Go working on OpenGL.
* [glop](https://github.com/runningwild/glop) - Glop (Game Library Of Power) is a fairly simple cross-platform game library.
* [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A* path finding algorithm
* [go-collada](https://github.com/GlenKelley/go-collada) - Go package for working with the Collada file format.
* [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [go3d](https://github.com/ungerik/go3d) - A performance oriented 2D/3D math package for Go
* [gonet](https://github.com/xtaci/gonet) - A game server skeleton implemented with golang
* [Leaf](https://github.com/name5566/leaf) - A lightweight game server framework


## Generation & Generics

*Tools to enhance the language with features like generics via code generation*

* [gen](https://github.com/clipperhouse/gen) - Code generation tool for ‘generics’-like functionality.
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go.
* [pkgreflect](https://github.com/ungerik/pkgreflect) - A Go preprocessor for package scoped reflection.


## GUI

*Libraries for building GUI Applications*

* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK
* [go-qml](https://github.com/go-qml/qml) - QML support for the Go language
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go.
* [gotk3](https://github.com/conformal/gotk3) - Go bindings for GTK3.
* [gxui](https://github.com/google/gxui) - A Go cross platform UI library.
* [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go.
* [walk](https://github.com/lxn/walk) - Windows application library kit for Go.

## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips
* [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string.
* [gift](https://github.com/disintegration/gift) - Package of image processing filters.
* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library.
* [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go.
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go.
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API.
* [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package.
* [img](https://github.com/hawx/img) - A selection of image manipulation tools.
* [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go
* [resize](https://github.com/nfnt/resize) - Image resizing for the Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD.
* [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation.
* [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder.

## Logging

*Libraries for generating and working with log files.*

* [glog](https://github.com/golang/glog) - Leveled execution logs for Go.
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers.
* [go-log](https://github.com/ian-kent/go-log) - A log4j implementation in Go.
* [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers.
* [log-voyage](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang.
* [logex](https://github.com/go-logex/logex) - An golang log lib, supports tracking and level, wrap by standard log lib
* [logrus](https://github.com/Sirupsen/logrus) - a structured logger for Go.
* [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
* [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger.
* [logxi](https://github.com/mgutz/logxi) - A 12-factor app logger that is fast and makes you happy.
* [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser.
* [mlog](https://github.com/jbrodriguez/mlog) - A simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
* [seelog](https://github.com/cihub/seelog) -   logging functionality with flexible dispatching, filtering, and formatting.
* [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
* [tail](https://github.com/ActiveState/tail) - A Go package striving to emulate the features of the BSD tail program.

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang.
* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library.
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang.
* [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go
* [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods.
* [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go.
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go
* [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go
* [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go.
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14.
* [mlgo](https://code.google.com/p/mlgo/) - This project aims to provide minimalistic machine learning algorithms in Go.
* [neural-go](https://github.com/schuyler/neural-go) - A multilayer perceptron network implemented in Go, with training via backpropagation.
* [probab](https://code.google.com/p/probab/) - Probability distribution functions. Bayesian inference. Written in pure Go.
* [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go


## Messaging

*Libraries that implement messaging systems*

* [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus.
* [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility.
* [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec.
* [go-nsq](https://github.com/bitly/go-nsq) - the official Go package for NSQ
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster.
* [machinery](https://github.com/RichardKnop/machinery) - An asynchronous task queue/job queue based on distributed message passing.
* [NATS](https://github.com/apcera/nats) - A lightweight and highly performant publish-subscribe and distributed queueing messaging system.
* [oplog](https://github.com/dailymotion/oplog) - A generic oplog/replication system for REST APIs
* [pubsub](https://github.com/tuxychandru/pubsub) - A simple pubsub package for go.
* [sarama](https://github.com/Shopify/sarama) - A Go library for Apache Kafka.
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - A redis backed unified push service for server-side notifications to mobile devices.
* [zmq4](https://github.com/pebbe/zmq4) - A Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).


## Miscellaneous

*These libraries were placed here because none of the other categories seemed to fit*

* [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields.
* [browscap_go](https://github.com/fromYukki/browscap_go) - GoLang Library for [Browser Capabilities Project](http://browscap.org/).
* [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser
* [go-multierror](https://github.com/hashicorp/go-multierror) - A Go (golang) package for representing a list of errors as a single error.
* [gopsutil](https://github.com/shirou/gopsutil) - A cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).
* [jobs](https://github.com/albrow/jobs) - A persistent and flexible background jobs library.
* [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal.
* [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber
* [xstrings](https://github.com/huandu/xstrings) - A collection of useful string functions ported from other languages.

## Natural Language Processing

*Libraries for working with human languages.*

* [go-eco](https://code.google.com/p/go-eco/) - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.
* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm.
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
* [paicehusk](https://github.com/Rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm
* [porter](https://github.com/a2800276/porter) - This is a fairly straighforward port of Martin Porter's C implementation of the Porter stemming algorithm.
* [porter2](https://github.com/surge/porter2) - Really fast Porter 2 stemmer.
* [segment](https://github.com/blevesearch/segment) - A Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers.
* [textcat](https://github.com/pebbe/textcat) - A Go package for n-gram based text categorization, with support for utf-8 and raw text

## Networking

*Libraries for working with various layers of the network*

* [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826.
* [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy
* [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252)
* [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
* [dns](https://github.com/miekg/dns) - Go library for working with DNS
* [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.
* [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).
* [go-stun](https://github.com/ccding/go-stun) - A go implementation of the STUN client (RFC 3489 and RFC 5389).
* [gopacket](https://github.com/google/gopacket) - A Go library for packet processing with libpcap bindings
* [gopcap](https://github.com/akrennmair/gopcap) - A Go wrapper for libpcap
* [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions
* [gotcp](https://github.com/gansidui/gotcp) - A Go package for quickly writing tcp applications
* [graval](https://github.com/koofr/graval) - An experimental FTP server framework.
* [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces
* [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang
* [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it.
* [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface.
* [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.
* [tcp_server](https://github.com/firstrow/tcp_server) - A Go library for building tcp servers faster.
* [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation.

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow).
* [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3.
* [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
* [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events.
* [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM.


## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - A powerful orm framework for go.
* [go-store](https://github.com/gosuri/go-store) - A simple and fast Redis backed key-value store library for Go.
* [gomodel](https://github.com/cosiner/gomodel) - A lightweight, fast, orm-like library helps interactive with database.
* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
* [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go.
* [hood](https://github.com/eaigner/hood) - Database agnostic ORM for Go.
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM.
* [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go.
* [Zoom](https://github.com/albrow/zoom) - A blazing-fast datastore and querying engine built on Redis.


## Package Management

*Libraries for package and dependency management.*

* [gigo](https://github.com/LyricalSecurity/gigo) - PIP-like dependency tool for golang, with support for private repositories and hashes.
* [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.
* [gom](https://github.com/mattn/gom) - Go Manager - bundle for go.
* [goop](https://github.com/nitrous-io/goop) - A simple dependency manager for Go (golang), inspired by Bundler.
* [gopm](https://github.com/gpmgo/gopm) - Go Package Manager
* [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go.
* [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git
* [nut](https://github.com/jingweno/nut) - Vendor Go dependencies
* [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments


## Query Language

* [graphql-go](https://github.com/chris-ramon/graphql-go) - An implementation of GraphQL for Go.


## Resource Embedding

* [go-bindata](https://github.com/jteeuwen/go-bindata) - Package that converts any file into managable Go source code.
* [go-resources](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go.
* [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.
* [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem.


## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [blas](https://github.com/ziutek/blas) - Implementation of BLAS (Basic Linear Algebra Subprograms)
* [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types.
* [evaler](https://github.com/soniah/evaler) - A simple floating point arithmetic expression evaluator
* [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang
* [go-fn](https://code.google.com/p/go-fn/) - Mathematical functions written in Go language, that are not covered by math pkg
* [go-gt](https://code.google.com/p/go-gt/) - Graph theory algorithms written in "Go" language
* [go.matrix](https://github.com/skelterjohn/go.matrix) - linear algebra for go (has been stalled)
* [gocomplex](https://code.google.com/p/gocomplex/) - A complex number library for the Go programming language.
* [gofrac](https://github.com/anschelsc/gofrac) - A (goinstallable) fractions library for go with support for basic arithmetic.
* [gohistogram](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams
* [gonum/mat64](https://github.com/gonum/matrix) - The general purpose package for matrix computation. Package mat64 provides basic linear algebra operations for float64 matrices.
* [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go.
* [goraph](https://github.com/gyuho/goraph) - A pure Go graph theory library(data structure, algorith visualization)
* [gostat](https://code.google.com/p/gostat/) - A statistics library for the go language
* [mudlark-go](https://code.google.com/p/mudlark-go-pkgs/) - A collection of packages providing (hopefully) useful code for use in software using Google's Go programming language.
* [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go
* [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data.
* [vectormath](https://github.com/spate/vectormath) - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code. (currently inactive)


## Security

*Libraries that are used to help make your application more secure.*

* [BadActor](https://github.com/jaredfolkins/badactor) - An in-memory, application-driven jailer built in the spirit of fail2ban
* [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)"

## Serialization

*Libraries and tools for binary serialization*

* [cbor](https://github.com/2tvenom/cbor) - Golang library for working with cbor binary format
* [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go
  * [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go.
* [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets
* [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.
* [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures.
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions
* [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures.


## Server Applications

* [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
* [Caddy](https://github.com/mholt/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery.
* [nsq](http://nsq.io/) - A realtime distributed messaging platform
* [yakvs](https://github.com/sci4me/yakvs) - A small, networked, in-memory key-value store.


## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace) - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.
* [amber](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.
* [damsel](https://github.com/dskinner/damsel) - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.
* [ego](https://github.com/benbjohnson/ego) - A lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation.
* [mustache](https://github.com/hoisie/mustache) - A Go implementation of the Mustache template language.
* [pongo2](https://github.com/flosch/pongo2) - A Django-like template-engine for Go.
* [raymond](https://github.com/aymerick/raymond) - A complete handlebars implementation in Go.
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang.
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/)


## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/bluesuncorp/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions
    * [assert](https://github.com/bmizerany/assert) - Asserts to Go testing
    * [bro](https://github.com/marioidival/bro) - Watch files in directory and run tests for them
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go
    * [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code
    * [goblin](https://github.com/franela/goblin) - Mocha like testing framework fo Go
    * [gocheck](http://labix.org/gocheck) - A more advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload
    * [godog](https://github.com/DATA-DOG/godog) - Cucumber or Behat like BDD framework for Go.
    * [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language.
    * [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
    * [restit](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test.
    * [Testify](https://github.com/stretchr/testify) - A sacred extension to the standard go testing package.

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes.
    * [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language.
    * [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter

* Fuzzing and delta-debugging/reducing/shrinking
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) - A randomized testing system
    * [gofuzz](https://github.com/google/gofuzz) - A library for populating go objects with random values
    * [gogenerate](https://github.com/arschles/gogenerate) - A Scalacheck-like library for Go
    * [Tavor](https://github.com/zimmski/tavor) - A generic fuzzing and delta-debugging framework

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go
        * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links.
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer
    * [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).
    * [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings
    * [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format.
    * [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language.
    * [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    * [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
    * [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string.
    * [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language.
    * [gommon/gytes](https://github.com/labstack/gommon/tree/master/gytes) - Format bytes to string.
    * [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts
    * [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.
    * [goregen](https://github.com/zach-klippenstein/goregen) - A library for generating random strings from regular expressions.
    * [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text.
    * [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.
    * [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support.
    * [Slugify](https://github.com/avelino/slugify) - A Go slugify application that handles string.
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection).
* Utility
    * [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go.
    * [xurls](https://github.com/mvdan/xurls) - Extract urls from text


## Third-party APIs

*Libraries for accessing third party APIs.*

* [anaconda](https://github.com/ChimeraCoder/anaconda) - A Go client library for the Twitter 1.1 API
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) - The official AWS SDK for the Go programming language. Caution: The SDK is currently in the process of being developed, and not everything may be working fully yet.
* [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API.
* [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API
* [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API
* [gami](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface.
* [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](http://open.mapquestapi.com/nominatim/), [OpenCage](http://geocoder.opencagedata.com/api.html), [HERE](https://developer.here.com/rest-apis/documentation/geocoder) and [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx) geocoding / reverse geocoding APIs.
* [github](https://github.com/google/go-github) - Go library for accessing the GitHub API.
* [go-marathon](https://github.com/gambol99/go-marathon) - A Go library for interacting with Mesosphere's Marathon PAAS
* [goamz](https://github.com/mitchellh/goamz) - Popular fork of [goamz](https://launchpad.net/goamz) which adds some missing API calls to certain packages.
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - a Go MusicBrainz WS2 client library
* [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go
* [google-analytics](https://github.com/chonthu/go-google-analytics) - A simple wrapper for easy google analytics reporting
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library
* [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.
* [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API.
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP.
* [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
* [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API.
* [rrdaclient](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.
* [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API.
* [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API.
* [snapchat](https://github.com/jamieomatthews/gosnap) - Go wrapper for the snapchat API
* [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API.
* [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers.
* [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) - A simple golang package to communicate with [themoviedb.org](https://themoviedb.org)
* [translate](https://github.com/poorny/translate) - Go online translation package
* [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API.

## Utilities

*General utilities and tools to make your life easier.*

* [coop](https://github.com/rakyll/coop) - Cheat sheet for some of the common concurrent flows in Go.
* [delve](https://github.com/derekparker/delve) - Go debugger.
* [fastlz](https://github.com/fromYukki/fastlz) - Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang.
* [go-cron](https://github.com/rk/go-cron) - A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
* [go-debug](https://github.com/tj/go-debug) - Conditional debug logging for Golang libraries & applications
* [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go.
* [go-underscore](https://github.com/tobyhede/go-underscore) - A useful collection of helpfully functional Go collection utilities.
* [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package.
* [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons
* [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.)
* [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox.
* [gohper](https://github.com/cosiner/gohper) - Various tools/modules help for development.
* [golarm](https://github.com/msempere/golarm) - Fire alarms with system events.
* [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs.
* [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images
* [goreq](https://github.com/franela/goreq) - Minimal and simple request library for Go language.
* [gorequest](https://github.com/parnurzeal/gorequest) - Simplified HTTP client with rich features for Go.
* [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go
* [hystrix-go](https://github.com/afex/hystrix-go) - Imprements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.
* [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON.
* [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects.
* [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go
* [mp](https://github.com/sanbornm/mp) - A simple cli email parser. It currently takes stdin and outputs JSON.
* [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers
* [netbug](https://github.com/e-dard/netbug) - Easy remote profiling of your services.
* [ngrok](https://github.com/inconshreveable/ngrok) - Introspected tunnels to localhost.
* [okrun](https://github.com/xta/okrun) - go run error steamroller
* [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump.
* [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool
* [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency
* [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API
* [profile](https://github.com/davecheney/profile) - Simple profiling support package for Go
* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™.
* [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics
* [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy.
* [sling](https://github.com/dghubble/sling) - Go HTTP requests builder for API clients.
* [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options.
* [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package
* [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.


## Validation

*Libraries for validation.*

* [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs
* [validator](https://github.com/bluesuncorp/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving


## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks
* [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2.
* [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go.
* [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.


## Video

*Libraries for manipulating video.*

* [aac/h264](https://github.com/nareix/codec) - Golang aac/h264 encoder and decoder.
* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries.
* [goav](https://github.com/giorgisio/goav) - Comphrensive Go bindings for FFmpeg.
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer.


## Web Frameworks

*Full stack web frameworks.*

* [Beego](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language.
* [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer.
* [Echo](https://github.com/labstack/echo) - A fast HTTP router (zero memory allocation) and micro web framework in Go.
* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
* [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) - A quick and easy way to setup a RESTful JSON API
* [go-relax](https://github.com/codehack/go-relax) - A framework of pluggable components to build RESTful API's
* [go-rest](https://github.com/ungerik/go-rest) - A small and evil REST framework for Go
* [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework.
* [Goat](https://github.com/bahlo/goat) - A minimalistic REST API server in Go
* [gocraft/web](https://github.com/gocraft/web) - A mux and middleware package in Go.
* [Goji](https://github.com/zenazn/goji) - Goji is a minimalistic web framework for Golang that's high in antioxidants.
* [Gondola](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster
* [goose](https://github.com/ian-kent/goose) - Server Sent Events in Go
* [Gorilla](https://github.com/gorilla/) - Gorilla is a web toolkit for the Go programming language.
* [httprouter](https://github.com/julienschmidt/httprouter) - A high performance router. Use this and the standard http handlers to form a very high performance web framework.
* [Macaron](https://github.com/Unknwon/macaron) - Macaron is a high productive and modular design web framework in Go.
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
* [Martini](https://github.com/go-martini/martini) - Martini is a powerful package for quickly writing modular web applications/services in Golang.
* [medeina](https://github.com/imdario/medeina) - Medeina is a HTTP routing tree based on HttpRouter, inspired by Roda and Cuba.
* [neo](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API.
* [pat](https://github.com/bmizerany/pat) - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.
* [Resoursea](https://github.com/resoursea/api) - A REST framework for quickly writing resource based services.
* [Revel](https://github.com/revel/revel) - A high-productivity web framework for the Go language.
* [rex](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.
* [sawsij](http://sawsij.com/) - lightweight, open-source web framework for building high-performance, data-driven web applications.
* [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers
* [tango](https://github.com/lunny/tango) - Micro & pluggable web framework for Go.
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - A Go framework for building JSON web services inspired by Dropwizard
* [traffic](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go.
* [web.go](https://github.com/hoisie/web) - A simple framework to write webapps in Go.
* [Zerver](https://github.com/cosiner/zerver) - Zerver is a expressive, modular, feature completed RESTful framework.
* [zeus](https://github.com/daryl/zeus) - A very simple and fast HTTP router for Go.


### Middlewares

#### Actual middlewares

* [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API
* [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST
* [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler
* [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go.
* [go-wrap](https://github.com/go-on/wrap) - Small middlewares package for net/http.
* [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http.
* [negroni](https://github.com/codegangsta/negroni) - Idiomatic HTTP middleware for Golang.
* [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses.
* [stats](https://github.com/thoas/stats) - A Go middleware that stores various information about your web application.

# Tools

Go software and plugins.


## Code Analysis

* [doc](http://godoc.org/robpike.io/cmd/doc) - Go documentation tool that produces an alternative doc format.
* [dupl](https://github.com/mibk/dupl) - A tool for code clone detection.
* [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs.
* [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time.
* [Go Metalinter](https://github.com/alecthomas/gometalinter) - Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer.
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code.
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [gostatus](https://github.com/shurcooL/gostatus) - A command line tool, shows the status of repositories that contain Go packages.
* [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags.


## Editor Plugins

* [go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) Go plugin for IntelliJ IDEA.
* [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting
* [gocode](https://github.com/nsf/gocode) - An autocompletion daemon for the Go programming language
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - A Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features.
* [velour](https://github.com/velour/velour) - An IRC client for the acme editor.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - A Vim plugin to highlight syntax errors on save.
* [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim.
* [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes.

## Go Tools

* [colorgo](https://github.com/songgao/colorgo) - A wrapper around `go` command for colorized `go build` output.

## Software Packages

Software written in Go.


### DevOps Tools

* [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool
* [awsenv](https://github.com/soniah/awsenv) - a small binary that loads Amazon (AWS) environment variables for a profile
* [Boom](https://github.com/rakyll/boom) - Boom is a tiny program that sends some load to a web application.
* [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart).
* [EasySSH](https://github.com/hypersleep/easyssh) - Golang package for easy remote execution through SSH and SCP downloading.
* [gaudi](http://gaudi.io/) - Gaudi automates the setup of isolated and decoupled dev environments.
* [Go Metrics](https://github.com/rcrowley/go-metrics) - Go port of Coda Hale's Metrics library: https://github.com/codahale/metrics.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update.
* [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go.
* [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application.
* [Gogs](http://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
* [gox](https://github.com/mitchellh/gox) - A dead simple, no frills Go cross compile tool.
* [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging.
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions.
* [hk](https://github.com/heroku/hk) - Heroku command-line interface in Go.
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data.
* [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB
* [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
* [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies.
* [Vegeta] (https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000!
* [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server
* [Wide](https://wide.b3log.org) - A Web-based IDE for Teams using Golang.


### Other Software
* [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine
* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
* [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul.
* [Docker](http://www.docker.com/) - An open platform for distributed applications for developers and sysadmins.
* [fleet](https://github.com/coreos/fleet) - A Distributed init System.
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store#go-package-store-) - An app that displays updates for the Go packages in your GOPATH.
* [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go.
* [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
* [heka](https://github.com/mozilla-services/heka) - universal tool for data processing from Mozilla. Large collection of built-in plugins. Extendable via Go and Lua plugin API.
* [hugo](http://gohugo.io/) - A Fast and Modern Static Website Engine
* [Juju](https://juju.ubuntu.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [limetext](http://limetext.org/) Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [naclpipe](https://github.com/unix4fun/naclpipe) - A simple NaCL EC25519 based crypto pipe tool written in Go.
* [nes](https://github.com/fogleman/nes) - A Nintendo Entertainment System (NES) emulator written in Go.
* [orange-cat](https://github.com/noraesae/orange-cat) - A Markdown previewer written in Go.
* [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
* [plubi](https://github.com/norwack/plubi) - A Golang Plugin Based IRC Bot.
* [Postman](https://github.com/zachlatta/postman) - Command-line utility for batch-sending email.
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
* [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control)
* [syncthing](http://www.syncthing.net/) - An open, decentralized file synchronization tool and protocol.
* [Tenyks](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging.
* [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests.
* [tsuru](http://www.tsuru.io/) - An extensible and open source Platform as a Service software.
* [websysd](https://github.com/ian-kent/websysd) - Web based process manager (like Marathon or Upstart)







# Resources

Where to discover new Go libraries.


## Benchmarks

* [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions.
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison.
* [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) - Naive performance test of two ways to do type assertion in Go.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods.
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language.
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) - Tiny collection of Go micro benchmarks. The intent is to compare some language features to others.
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - A collection of benchmarks for popular Go database/SQL utilities.
* [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language.


## Conferences

* [dotGo](http://www.dotgo.io) - Paris, France
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan
* [GolangUK](http://www.golanguk.com/) - London, UK
* [GopherChina](http://gopherchina.org) - Shanghai, China
* [GopherCon](http://www.gophercon.com/) - Denver, USA
* [GopherCon India](http://www.gophercon.in/) - Bengaluru, India
* [GothamGo](http://gothamgo.com/) - New York City, USA

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](http://astaxie.gitbooks.io/build-web-application-with-golang/)
* [Building Web Apps With Go](http://codegangsta.gitbooks.io/building-web-apps-with-go/)
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books
* [Learning Go](http://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](http://jan.newmarch.name/go/)


## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangweekly](https://twitter.com/golangweekly)


## Websites

* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - A curated list of awesome remote jobs. A lot of them is looking for Go hackers.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists.
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - A collection of Go articles and tutorials.
* [Go Blog](http://blog.golang.org) - The official Go blog
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki
* [godoc.org](http://godoc.org/) - Documentation for open source Go packages.
* [golang-graphics](https://github.com/mholt/golang-graphics) - A collection of Go images, graphics, and art
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [r/Golang](http://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.


### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go
* [Go By Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated example programs
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql
* [Working with Go](https://github.com/mkaz/working-with-go) - An intro to go for experienced programmers


## Windows

* [go-ole](https://github.com/mattn/go-ole) - Win32 OLE implementation for golang
