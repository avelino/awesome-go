Awesome Go [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/avelino/awesome-go](https://badges.gitter.im/avelino/awesome-go.svg)](https://gitter.im/avelino/awesome-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
==========================================================================================================================================================================================================================================================================================================================================================================================================================================================================================================================

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python) .

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors) ; you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

-   [Awesome Go](#awesome-go)

    -   [Audio & Music](#audiomusic)
    -   [Authentication & OAuth](#authentication--oauth)
    -   [Command Line](#command-line)
    -   [Configuration](#configuration)
    -   [Continuous Integration](#continuous-integration)
    -   [CSS Preprocessors](#css-preprocessors)
    -   [Data Structures](#data-structures)
    -   [Database](#database)
    -   [Database Drivers](#database-drivers)
    -   [Date & Time](#date--time)
    -   [Distributed Systems](#distributed-systems)
    -   [Email](#email)
    -   [Embeddable Scripting Languages](#embeddable-scripting-languages)
    -   [Financial](#financial)
    -   [Forms](#forms)
    -   [Game Development](#game-development)
    -   [Generation & Generics](#generation--generics)
    -   [Go Compilers](#go-compilers)
    -   [Goroutines](#goroutines)
    -   [GUI](#gui)
    -   [Hardware](#hardware)
    -   [Images](#images)
    -   [Logging](#logging)
    -   [Machine Learning](#machine-learning)
    -   [Messaging](#messaging)
    -   [Miscellaneous](#miscellaneous)
    -   [Natural Language Processing](#natural-language-processing)
    -   [Networking](#networking)
    -   [OpenGL](#opengl)
    -   [ORM](#orm)
    -   [Package Management](#package-management)
    -   [Query Language](#query-language)
    -   [Resource Embedding](#resource-embedding)
    -   [Science and Data Analysis](#science-and-data-analysis)
    -   [Security](#security)
    -   [Serialization](#serialization)
    -   [Template Engines](#template-engines)
    -   [Testing](#testing)
    -   [Text Processing](#text-processing)
    -   [Third-party APIs](#third-party-apis)
    -   [Utilities](#utilities)
    -   [Validation](#validation)
    -   [Version Control](#version-control)
    -   [Video](#video)
    -   [Web Frameworks](#web-frameworks)
        -   [Middlewares](#middlewares)
            -   [Actual middlewares](#actual-middlewares)
            -   [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
    -   [Windows](#windows)
-   [Tools](#tools)

    -   [Code Analysis](#code-analysis)
    -   [Editor Plugins](#editor-plugins)
    -   [Go Tools](#go-tools)
    -   [Software Packages](#software-packages)
        -   [DevOps Tools](#devops-tools)
        -   [Other Software](#other-software)
-   [Server Applications](#server-applications)

-   [Resources](#resources)

    -   [Benchmarks](#benchmarks)
    -   [Conferences](#conferences)
    -   [E-Books](#e-books)
    -   [Twitter](#twitter)
    -   [Websites](#websites)
        -   [Tutorials](#tutorials)

Audio/Music
-----------

*Libraries for manipulating audio.*

-   [flac](https://github.com/eaburns/flac) <span> ★ 48, pushed 44 days ago </span> - A native Go FLAC decoder.
-   [flac](https://github.com/mewkiz/flac) <span> ★ 23, pushed 3 days ago </span> - A native Go FLAC decoder.
-   [go-sox](https://github.com/krig/go-sox) <span> ★ 36, pushed 403 days ago </span> - libsox bindings for go.
-   [go\_mediainfo](https://github.com/zhulik/go_mediainfo) <span> ★ 3, pushed 123 days ago </span> - libmediainfo bindings for go.
-   [mix](https://github.com/go-mix/mix) <span> ★ 18, pushed 22 days ago </span> - Sequence-based Go-native audio mixer for music apps.
-   [mp3](https://github.com/tcolgate/mp3) <span> ★ 14, pushed 138 days ago </span> - A native Go MP\# decoder.
-   [music-theory](https://github.com/go-music-theory/music-theory) <span> ★ 94, pushed 22 days ago </span> - Music theory models in Go.
-   [PortAudio](https://github.com/gordonklaus/portaudio) <span> ★ 37, pushed 30 days ago </span> - Go bindings for the PortAudio audio I/O library.
-   [portmidi](https://github.com/rakyll/portmidi) <span> ★ 86, pushed 24 days ago </span> - Go bindings for PortMidi.
-   [taglib](https://github.com/wtolson/go-taglib) <span> ★ 40, pushed 163 days ago </span> - Go bindings for taglib.
-   [vorbis](https://github.com/mccoyst/vorbis) <span> ★ 9, pushed 192 days ago </span> - A "native" Go Vorbis decoder (uses CGO, but has no dependencies).
-   [waveform](https://github.com/mdlayher/waveform) <span> ★ 120, pushed 499 days ago </span> - Go package capable of generating waveform images from audio streams.

Authentication & OAuth
----------------------

*Libraries for implementing authentications schemes.*

-   [authboss](https://github.com/go-authboss/authboss) <span> ★ 598, pushed 61 days ago </span> - A modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.
-   [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth) <span> ★ 123, pushed 34 days ago </span> - AWS (Amazon Web Services) request signing library.
-   [go-jose](https://github.com/square/go-jose) <span> ★ 501, pushed 2 days ago </span> - A fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
-   [go.auth](https://github.com/bradrydzewski/go.auth) <span> ★ 306, pushed 372 days ago </span> - Authentication API for Go web applications.
-   [gologin](https://github.com/dghubble/gologin) <span> ★ 630, pushed 15 days ago </span> - chainable handlers for login with OAuth1 and OAuth2 authentication providers.
-   [gorbac](https://github.com/mikespook/gorbac) <span> ★ 328, pushed 53 days ago </span> - provides a lightweight role-based access control (RBAC) implementation in Golang.
-   [goth](https://github.com/markbates/goth) <span> ★ 680, pushed 4 days ago </span> - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple provides out of the box.
-   [httpauth](https://github.com/goji/httpauth) <span> ★ 91, pushed 0 days ago </span> - HTTP Authentication middleware.
-   [jwt-go](https://github.com/dgrijalva/jwt-go) <span> ★ 1106, pushed 0 days ago </span> - Golang implementation of JSON Web Tokens (JWT).
-   [oauth2](https://github.com/golang/oauth2) <span> ★ 701, pushed 0 days ago </span> - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.
-   [osin](https://github.com/RangelReale/osin) <span> ★ 861, pushed 5 days ago </span> - Golang OAuth2 server library.
-   [permissions2](https://github.com/xyproto/permissions2) <span> ★ 157, pushed 20 days ago </span> - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.
-   [yubigo](https://github.com/GeertJohan/yubigo) <span> ★ 55, pushed 705 days ago </span> - a Yubikey client package that provides a simple API to integrate the Yubico Yubikey into a go application.

Command Line
------------

### Standard CLI

*Libraries for building standard or basic Command Line applications*

-   [cli](https://github.com/mkideal/cli) <span> ★ 81, pushed 1 days ago </span> - A feature-rich and easy to use command-line package based on golang tag
-   [cli-init](https://github.com/tcnksm/gcli) <span> ★ 480, pushed 130 days ago </span> - The easy way to start building Golang command line application.
-   [climax](http://github.com/tucnak/climax) - An alternative CLI with "human face", in spirit of Go command
-   [cobra](https://github.com/spf13/cobra) <span> ★ 1977, pushed 12 days ago </span> - A Commander for modern Go CLI interactions
-   [codegangsta/cli](https://github.com/codegangsta/cli) <span> ★ 4087, pushed 0 days ago </span> - A small package for building command line apps in Go.
-   [docopt.go](https://github.com/docopt/docopt.go) <span> ★ 663, pushed 69 days ago </span> - A command-line arguments parser that will make you smile.
-   [go-flags](https://github.com/jessevdk/go-flags) <span> ★ 568, pushed 32 days ago </span> - go command line option parser
-   [kingpin](https://github.com/alecthomas/kingpin) <span> ★ 602, pushed 0 days ago </span> - A command line and flag parser supporting sub commands.
-   [liner](https://github.com/peterh/liner) <span> ★ 301, pushed 30 days ago </span> - A Go readline-like library for command-line interfaces.
-   [mitchellh/cli](https://github.com/mitchellh/cli) <span> ★ 403, pushed 31 days ago </span> - A Go library for implementing command-line interfaces.
-   [mow.cli](https://github.com/jawher/mow.cli) <span> ★ 342, pushed 64 days ago </span> - A Go library for building CLI applications with sophisticated flag and argument parsing and validation.
-   [readline](https://github.com/chzyer/readline) <span> ★ 554, pushed 4 days ago </span> - A pure golang implementation that provide most of features in GNU-Readline under MIT license.
-   [ukautz/clif](https://github.com/ukautz/clif) <span> ★ 63, pushed 107 days ago </span> - A small command line interface framework.
-   [wlog](https://github.com/dixonwille/wlog) <span> ★ 0, pushed 3 days ago </span> - A simple logging interface that supports cross-platform color and concurrency.
-   [wmenu](https://github.com/dixonwille/wmenu) <span> ★ 0, pushed 0 days ago </span> - An easy to use menu structure for cli applications that prompts users to make choices.

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces*

-   [chalk](https://github.com/ttacon/chalk) <span> ★ 154, pushed 312 days ago </span> - Intuitive package for prettifying terminal/console output.
-   [color](https://github.com/fatih/color) <span> ★ 852, pushed 39 days ago </span> - Versatile package for colored terminal output.
-   [colourize](https://github.com/TreyBastian/colourize) <span> ★ 4, pushed 69 days ago </span> - Go library for ANSI colour text in terminals.
-   [go-colortext](https://github.com/daviddengcn/go-colortext) <span> ★ 142, pushed 9 days ago </span> - Go library for color output in terminals.
-   [gocui](https://github.com/jroimartin/gocui) <span> ★ 647, pushed 5 days ago </span> - Minimalist Go library aimed at creating Console User Interfaces.
-   [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
-   [termbox-go](https://github.com/nsf/termbox-go) <span> ★ 1406, pushed 94 days ago </span> - Termbox is a library for creating cross-platform text-based interfaces.
-   [termtables](https://github.com/apcera/termtables) <span> ★ 51, pushed 139 days ago </span> - A Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output
-   [termui](https://github.com/gizak/termui) <span> ★ 4637, pushed 18 days ago </span> - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib) .
-   [uilive](https://github.com/gosuri/uilive) <span> ★ 428, pushed 83 days ago </span> - A library for updating terminal output in realtime.
-   [uiprogress](https://github.com/gosuri/uiprogress) <span> ★ 792, pushed 83 days ago </span> - A flexible library to render progress bars in terminal applications.
-   [uitable](https://github.com/gosuri/uitable) <span> ★ 328, pushed 21 days ago </span> - A library to improve readability in terminal apps using tabular data.

Configuration
-------------

*Libraries for configuration parsing*

-   [config](https://github.com/olebedev/config) <span> ★ 71, pushed 104 days ago </span> - JSON or YAML configuration wrapper with environment variables and flags parsing.
-   [configure](https://github.com/paked/configure) <span> ★ 9, pushed 17 days ago </span> - Provides configuration through multiple sources, including JSON, flags and environment variables.
-   [env](https://github.com/caarlos0/env) <span> ★ 92, pushed 10 days ago </span> - Parse environment variables to Go structs (with defaults).
-   [envcfg](https://github.com/tomazk/envcfg) <span> ★ 76, pushed 55 days ago </span> - Un-marshaling environment variables to Go structs.
-   [envconf](https://github.com/ian-kent/envconf) <span> ★ 3, pushed 547 days ago </span> - Configuration from environment
-   [envconfig](https://github.com/vrischmann/envconfig) <span> ★ 89, pushed 81 days ago </span> - Read your configuration from environment variables.
-   [gcfg](https://github.com/go-gcfg/gcfg) <span> ★ 33, pushed 113 days ago </span> - read INI-style configuration files into Go structs; supports user-defined types and subsections
-   [gofigure](https://github.com/ian-kent/gofigure) <span> ★ 33, pushed 71 days ago </span> - Go application configuration made easy
-   [ingo](https://github.com/schachmat/ingo) <span> ★ 6, pushed 10 days ago </span> - Flags persisted in an ini-like config file
-   [ini](https://github.com/go-ini/ini) <span> ★ 268, pushed 12 days ago </span> - Go package for read and write INI files
-   [mini](https://github.com/FogCreek/mini) <span> ★ 82, pushed 272 days ago </span> - A golang package for parsing ini-style configuration files
-   [store](https://github.com/tucnak/store) <span> ★ 38, pushed 107 days ago </span> - A lightweight configuration manager for Go
-   [viper](https://github.com/spf13/viper) <span> ★ 1492, pushed 5 days ago </span> - Go configuration with fangs

Continuous Integration
----------------------

*Tools for help with continuous integration*

-   [drone](https://github.com/drone/drone) <span> ★ 6709, pushed 0 days ago </span> - Drone is a Continuous Integration platform built on Docker, written in Go
-   [goveralls](https://github.com/mattn/goveralls) <span> ★ 250, pushed 36 days ago </span> - Go integration for Coveralls.io continuous code coverage tracking system.
-   [overalls](https://github.com/go-playground/overalls) <span> ★ 18, pushed 214 days ago </span> - Multi-Package go project coverprofile for tools like goveralls

CSS Preprocessors
-----------------

*Libraries for preprocessing CSS files*

-   [c6](https://github.com/c9s/c6) <span> ★ 339, pushed 198 days ago </span> - High performance SASS compatible-implementation compiler written in Go
-   [gcss](https://github.com/yosssi/gcss) <span> ★ 340, pushed 561 days ago </span> - Pure Go CSS Preprocessor.
-   [go-libsass](https://github.com/wellington/go-libsass) <span> ★ 26, pushed 27 days ago </span> - Go wrapper to the 100% Sass compatible libsass project.

Data Structures
---------------

*Generic datastructures and algorithms in Go.*

-   [binpacker](https://github.com/zhuangsirui/binpacker) <span> ★ 12, pushed 70 days ago </span> - Binary packer and unpacker helps user build custom binary stream.
-   [bitset](https://github.com/willf/bitset) <span> ★ 200, pushed 60 days ago </span> - Go package implementing bitsets.
-   [bloom](https://github.com/zhenjl/bloom) <span> ★ 82, pushed 182 days ago </span> - Bloom filters implemented in Go.
-   [boomfilters](https://github.com/tylertreat/BoomFilters) <span> ★ 692, pushed 97 days ago </span> - probabilistic data structures for processing continuous, unbounded streams
-   [count-min-log](https://github.com/seiflotfy/count-min-log) <span> ★ 19, pushed 35 days ago </span> - A Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).
-   [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) <span> ★ 143, pushed 228 days ago </span> - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
-   [encoding](https://github.com/zhenjl/encoding) <span> ★ 40, pushed 598 days ago </span> - Integer Compression Libraries for Go.
-   [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) <span> ★ 0, pushed 12 days ago </span> - A Go implementation of Adaptive Radix Tree.
-   [go-datastructures](https://github.com/Workiva/go-datastructures) <span> ★ 2438, pushed 0 days ago </span> - a collection of useful, performant, and thread-safe data structures
-   [go-geoindex](https://github.com/hailocab/go-geoindex) <span> ★ 167, pushed 89 days ago </span> - In-memory geo index.
-   [golang-set](https://github.com/deckarep/golang-set) <span> ★ 365, pushed 65 days ago </span> - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
-   [goskiplist](https://github.com/ryszard/goskiplist) <span> ★ 104, pushed 410 days ago </span> - A skip list implementation in Go.
-   [mafsa](https://github.com/smartystreets/mafsa) <span> ★ 205, pushed 34 days ago </span> - MA-FSA implementation with Minimal Perfect Hashing
-   [roaring](https://github.com/RoaringBitmap/roaring) <span> ★ 119, pushed 1 days ago </span> - Go package implementing compressed bitsets.
-   [skiplist](https://github.com/gansidui/skiplist) <span> ★ 27, pushed 521 days ago </span> - Skiplist implementation in Go
-   [trie](https://github.com/derekparker/trie) <span> ★ 120, pushed 53 days ago </span> - Trie implementation in Go
-   [ttlcache](https://github.com/diegobernardes/ttlcache) <span> ★ 22, pushed 18 days ago </span> - An in-memory LRU string-interface{} map with expiration for golang
-   [willf/bloom](https://github.com/willf/bloom) <span> ★ 232, pushed 60 days ago </span> - Go package implementing Bloom filters.

Database
--------

*Databases implemented in Go.*

-   [bolt](https://github.com/boltdb/bolt) <span> ★ 4062, pushed 1 days ago </span> - A low-level key/value database for Go.
-   [cache2go](https://github.com/muesli/cache2go) <span> ★ 71, pushed 188 days ago </span> - An in-memory key:value cache which supports automatic invalidation based on timeouts.
-   [cockroach](https://github.com/cockroachdb/cockroach) <span> ★ 6764, pushed 0 days ago </span> - A Scalable, Geo-Replicated, Transactional Datastore
-   [couchcache](https://github.com/codingsince1985/couchcache) <span> ★ 16, pushed 219 days ago </span> - A RESTful caching micro-service backed by Couchbase server.
-   [dgraph](https://github.com/dgraph-io/dgraph) <span> ★ 1126, pushed 5 days ago </span> - Scalable, Distributed, Low Latency, High Throughput Graph Database.
-   [diskv](https://github.com/peterbourgon/diskv) <span> ★ 343, pushed 21 days ago </span> - A home-grown disk-backed key-value store.
-   [forestdb](https://github.com/couchbase/goforestdb) <span> ★ 18, pushed 27 days ago </span> - Go bindings for ForestDB.
-   [GCache](https://github.com/bluele/gcache) <span> ★ 93, pushed 50 days ago </span> - Cache library with support for expirable Cache, LFU, LRU and ARC.
-   [go-cache](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
-   [goleveldb](https://github.com/syndtr/goleveldb) <span> ★ 1021, pushed 0 days ago </span> - An implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in the Go.
-   [groupcache](https://github.com/golang/groupcache) <span> ★ 4414, pushed 74 days ago </span> - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
-   [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics
-   [ledisdb](https://github.com/siddontang/ledisdb) <span> ★ 1583, pushed 14 days ago </span> - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
-   [levigo](https://github.com/jmhodges/levigo) <span> ★ 269, pushed 265 days ago </span> - Levigo is a Go wrapper for LevelDB.
-   [prometheus](https://github.com/prometheus/prometheus) <span> ★ 4436, pushed 0 days ago </span> - Monitoring system and time series database.
-   [rqlite](https://github.com/otoolep/rqlite) <span> ★ 1317, pushed 0 days ago </span> - Replicated SQLite, using Raft consensus.
-   [tidb](https://github.com/pingcap/tidb) <span> ★ 3846, pushed 0 days ago </span> - TiDB is a distributed SQL database. Inspired by the design of Google F1.
-   [tiedot](https://github.com/HouzuoGuo/tiedot) <span> ★ 1508, pushed 96 days ago </span> - Your NoSQL database powered by Golang.
-   [Tile38](https://github.com/tidwall/tile38) <span> ★ 1092, pushed 6 days ago </span> - A geolocation DB with spatial index and realtime geofencing.

*Database tools.*

-   [go-mysql](https://github.com/siddontang/go-mysql) <span> ★ 209, pushed 41 days ago </span> - A go toolset to handle MySQL protocol and replication.
-   [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) <span> ★ 263, pushed 110 days ago </span> - Sync your MySQL data into Elasticsearch automatically.
-   [goose](https://github.com/steinbacher/goose) <span> ★ 7, pushed 35 days ago </span> - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
-   [kingshard](https://github.com/flike/kingshard) <span> ★ 1666, pushed 0 days ago </span> - kingshard is a high performance proxy for MySQL powered by Golang.
-   [migrate](https://github.com/mattes/migrate) <span> ★ 586, pushed 17 days ago </span> - Database migration handling in Golang support MySQL,PostgreSQL,Cassandra and SQLite.
-   [myreplication](https://github.com/2tvenom/myreplication) <span> ★ 48, pushed 446 days ago </span> - MySql binary log replication listener. Support statement and row based replication.
-   [orchestrator](https://github.com/outbrain/orchestrator) <span> ★ 411, pushed 4 days ago </span> - MySQL replication topology manager & visualizer
-   [pgweb](https://github.com/sosedoff/pgweb) <span> ★ 3378, pushed 10 days ago </span> - A web-based PostgreSQL database browser
-   [pravasan](https://github.com/pravasan/pravasan) <span> ★ 11, pushed 402 days ago </span> - Simple Migration tool - currently for MySQL but planning to support soon for Postgres, SQLite, MongoDB, etc.,
-   [sql-migrate](https://github.com/rubenv/sql-migrate) <span> ★ 444, pushed 55 days ago </span> - Database migration tool. Allows embedding migrations into the application using go-bindata.
-   [vitess](https://github.com/youtube/vitess) <span> ★ 3423, pushed 0 days ago </span> - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.

*SQL query builder, libraries for building and using SQL.*

-   [dat](https://github.com/mgutz/dat) <span> ★ 285, pushed 44 days ago </span> - Go Postgres Data Access Toolkit
-   [Dotsql](https://github.com/gchaincl/dotsql) <span> ★ 222, pushed 127 days ago </span> - Go library that helps you keep sql files in one place and use it with ease.
-   [goqu](https://github.com/doug-martin/goqu) <span> ★ 253, pushed 111 days ago </span> - An idiomatic SQL builder and query library.
-   [igor](https://github.com/galeone/igor) <span> ★ 48, pushed 6 days ago </span> - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.
-   [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) <span> ★ 86, pushed 6 days ago </span> - Powerful data retrieval methods as well as DB-agnostic query building capabilities.
-   [scaneo](https://github.com/variadico/scaneo) <span> ★ 67, pushed 284 days ago </span> - Generate Go code to convert database rows into arbitrary structs.
-   [sqrl](https://github.com/elgris/sqrl) <span> ★ 38, pushed 404 days ago </span> - SQL query builder, fork of Squirrel with improved performance.
-   [Squirrel](https://github.com/Masterminds/squirrel) <span> ★ 796, pushed 17 days ago </span> - Go library that helps you build SQL queries.
-   [xo](https://github.com/knq/xo) <span> ★ 353, pushed 14 days ago </span> - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.

Database Drivers
----------------

*Libraries for connecting and operating databases.*

-   Relational Databases

    -   [firebirdsql](https://github.com/nakagami/firebirdsql) <span> ★ 42, pushed 5 days ago </span> - Firebird RDBMS SQL driver for Go
    -   [go-adodb](https://github.com/mattn/go-adodb) <span> ★ 33, pushed 43 days ago </span> - Microsoft ActiveX Object DataBase driver for go that using database/sql.
    -   [go-bqstreamer](https://github.com/rounds/go-bqstreamer) <span> ★ 77, pushed 0 days ago </span> - BigQuery fast and concurrent stream insert.
    -   [go-mssqldb](https://github.com/denisenkom/go-mssqldb) <span> ★ 312, pushed 11 days ago </span> - Microsoft MSSQL driver prototype in go language.
    -   [go-oci8](https://github.com/mattn/go-oci8) <span> ★ 127, pushed 2 days ago </span> - Oracle driver for go that using database/sql.
    -   [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) <span> ★ 2304, pushed 9 days ago </span> - MySQL driver for Go.
    -   [go-sqlite3](https://github.com/mattn/go-sqlite3) <span> ★ 1192, pushed 3 days ago </span> - SQLite3 driver for go that using database/sql.
    -   [gofreetds](https://github.com/minus5/gofreetds) <span> ★ 35, pushed 21 days ago </span> Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org) .
    -   [pgx](https://github.com/jackc/pgx) <span> ★ 556, pushed 3 days ago </span> - PostgreSQL driver supporting features beyond those exposed by database/sql.
    -   [pq](https://github.com/lib/pq) <span> ★ 2037, pushed 10 days ago </span> - Pure Go Postgres driver for database/sql.
-   NoSQL Databases

    -   [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) <span> ★ 168, pushed 23 days ago </span> - Aerospike client in Go language.
    -   [arangolite](https://github.com/solher/arangolite) <span> ★ 25, pushed 3 days ago </span> - Lightweight golang driver for ArangoDB.
    -   [cayley](https://github.com/google/cayley) <span> ★ 7388, pushed 0 days ago </span> - A graph database with support for multiple backends.
    -   [dynago](https://github.com/underarmour/dynago) <span> ★ 14, pushed 19 days ago </span> - Dynago is a principle of least surprise client for DynamoDB
    -   [go-couchbase](https://github.com/couchbase/go-couchbase) <span> ★ 207, pushed 4 days ago </span> - Couchbase client in Go
    -   [go-couchdb](https://github.com/fjl/go-couchdb) <span> ★ 30, pushed 182 days ago </span> - Yet another CouchDB HTTP API wrapper for Go
    -   [gocb](https://github.com/couchbase/gocb) <span> ★ 169, pushed 2 days ago </span> - Official Couchbase Go SDK
    -   [gocql](http://gocql.github.io) - A Go language driver for Apache Cassandra.
    -   [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
    -   [gorethink](https://github.com/dancannon/gorethink) <span> ★ 871, pushed 7 days ago </span> - Go language driver for RethinkDB
    -   [mgo](https://godoc.org/labix.org/v2/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    -   [neo4j](https://github.com/cihangir/neo4j) <span> ★ 14, pushed 389 days ago </span> - Neo4j Rest API Bindings for Golang
    -   [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) <span> ★ 62, pushed 1248 days ago </span> - Neo4j REST Client in golang.
    -   [neoism](https://github.com/jmcvetta/neoism) <span> ★ 244, pushed 73 days ago </span> - Neo4j client for Golang
    -   [redigo](https://github.com/garyburd/redigo) <span> ★ 2134, pushed 11 days ago </span> - Redigo is a Go client for the Redis database.
    -   [redis](https://github.com/go-redis/redis) <span> ★ 788, pushed 0 days ago </span> - Redis client for Golang
    -   [redis](https://github.com/hoisie/redis) <span> ★ 503, pushed 220 days ago </span> - A simple, powerful Redis client for Go.
    -   [redis](https://github.com/bsm/redeo) <span> ★ 81, pushed 436 days ago </span> - Redis-protocol compatible TCP servers/services.
-   Search and Analytic Databases

    -   [bleve](https://github.com/blevesearch/bleve) <span> ★ 2479, pushed 2 days ago </span> - A modern text indexing library for go.
    -   [elastic](https://github.com/olivere/elastic) <span> ★ 610, pushed 3 days ago </span> - Elasticsearch client for Google Go.
    -   [elastigo](https://github.com/mattbaird/elastigo) <span> ★ 729, pushed 13 days ago </span> - A Elasticsearch client library.
    -   [goes](https://github.com/belogik/goes) <span> ★ 70, pushed 118 days ago </span> - A library to interact with Elasticsearch.
    -   [skizze](https://github.com/seiflotfy/skizze) <span> ★ 17, pushed 14 days ago </span> - A probabilistic data-structures service and storage.

Date & Time
-----------

*Libraries for working with dates and times.*

-   [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) <span> ★ 5, pushed 53 days ago </span> - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
-   [goweek](https://github.com/grsmv/goweek) <span> ★ 4, pushed 87 days ago </span> - Library for working with week entity in golang.
-   [now](https://github.com/jinzhu/now) <span> ★ 589, pushed 207 days ago </span> - Now is a time toolkit for golang.
-   [NullTime](https://github.com/kirillDanshin/nulltime) <span> ★ 1, pushed 21 days ago </span> - Nullable time.Time
-   [timeutil](https://github.com/leekchan/timeutil) <span> ★ 103, pushed 267 days ago </span> - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.

Distributed Systems
-------------------

*Packages that help with building Distributed Systems.*

-   [celeriac](https://github.com/svcavallar/celeriac.v1) <span> ★ 11, pushed 95 days ago </span> - A library for adding support for interacting and monitoring Celery workers, tasks and events in Go
-   [flowgraph](https://github.com/vectaport/flowgraph) <span> ★ 10, pushed 167 days ago </span> - MPI-style ready-send coordination layer.
-   [glow](https://github.com/chrislusf/glow) <span> ★ 896, pushed 1 days ago </span> - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.
-   [go-jump](https://github.com/dgryski/go-jump) <span> ★ 65, pushed 628 days ago </span> - A port of Google's "Jump" Consistent Hash function.
-   [gorpc](https://github.com/valyala/gorpc) <span> ★ 223, pushed 31 days ago </span> - Simple, fast and scalable RPC library for high load.
-   [grpc-go](https://github.com/grpc/grpc-go) <span> ★ 1693, pushed 0 days ago </span> - The Go language implementation of gRPC. HTTP/2 based RPC.
-   [micro](https://github.com/micro/micro) <span> ★ 1796, pushed 0 days ago </span> - A pluggable microservice toolkit and distributed systems platform.
-   [raft](https://github.com/hashicorp/raft) <span> ★ 581, pushed 5 days ago </span> - Golang implementation of the Raft consensus protocol, by HashiCorp.
-   [torrent](https://github.com/anacrolix/torrent) <span> ★ 1116, pushed 5 days ago </span> - BitTorrent client package.
    -   [dht](https://godoc.org/github.com/anacrolix/torrent/dht) - BitTorrent Kademlia DHT implementation.
    -   [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) <span> ★ 145, pushed 21 days ago </span> - Video streaming torrent client.

Email
-----

*Libraries that implement email creation and sending*

-   [douceur](https://github.com/aymerick/douceur) <span> ★ 43, pushed 27 days ago </span> - CSS inliner for your HTML emails.
-   [email](https://github.com/jordan-wright/email) <span> ★ 670, pushed 22 days ago </span> - A robust and flexible email library for Go.
-   [go-dkim](https://github.com/toorop/go-dkim) <span> ★ 15, pushed 136 days ago </span> - A DKIM library, to sign & verify email.
-   [Gomail](https://github.com/go-gomail/gomail/) - Gomail is a very simple and powerful package to send emails.
-   [Hectane](https://github.com/hectane/hectane) <span> ★ 50, pushed 4 days ago </span> - Lightweight SMTP client providing an HTTP API
-   [MailHog](https://github.com/mailhog/MailHog) <span> ★ 842, pushed 34 days ago </span> - Email and SMTP testing with web and API interface
-   [SendGrid](https://github.com/sendgrid/sendgrid-go) <span> ★ 176, pushed 34 days ago </span> - SendGrid's Go library for sending email
-   [smtp](https://github.com/mailhog/smtp) <span> ★ 20, pushed 36 days ago </span> - SMTP server protocol state machine

Embeddable Scripting Languages
------------------------------

*Embedding other languages inside your go code*

-   [agora](https://github.com/PuerkitoBio/agora) <span> ★ 217, pushed 456 days ago </span> - Dynamically typed, embeddable programming language in Go
-   [anko](https://github.com/mattn/anko) <span> ★ 257, pushed 30 days ago </span> - Scriptable interpreter written in Go
-   [gisp](https://github.com/jcla1/gisp) <span> ★ 321, pushed 666 days ago </span> - Simple LISP in Go
-   [go-duktape](https://github.com/olebedev/go-duktape) <span> ★ 320, pushed 35 days ago </span> - Duktape JavaScript engine bindings for Go
-   [go-lua](https://github.com/Shopify/go-lua) <span> ★ 649, pushed 26 days ago </span> - A port of the Lua 5.2 VM to pure Go
-   [go-php](https://github.com/deuill/go-php) <span> ★ 116, pushed 12 days ago </span> - PHP bindings for Go
-   [go-python](https://github.com/sbinet/go-python) <span> ★ 363, pushed 24 days ago </span> - naive go bindings to the CPython C-API
-   [golua](https://github.com/aarzilli/golua) <span> ★ 268, pushed 18 days ago </span> - Go bindings for Lua C API
-   [gopher-lua](https://github.com/yuin/gopher-lua) <span> ★ 1166, pushed 0 days ago </span> - a Lua 5.1 VM and compiler written in Go
-   [otto](https://github.com/robertkrimen/otto) <span> ★ 2092, pushed 0 days ago </span> - A JavaScript interpreter written in Go
-   [purl](https://github.com/ian-kent/purl) <span> ★ 11, pushed 505 days ago </span> - Perl 5.18.2 embedded in Go

Financial
---------

*Packages for accounting and finance*

-   [accounting](https://github.com/leekchan/accounting) <span> ★ 197, pushed 247 days ago </span> - money and currency formatting for golang
-   [decimal](https://github.com/shopspring/decimal) <span> ★ 268, pushed 0 days ago </span> - Arbitrary-precision fixed-point decimal numbers

Forms
-----

*Libraries for working with forms.*

-   [bind](https://github.com/robfig/bind) <span> ★ 12, pushed 618 days ago </span> - Bind form data to any Go values
-   [binding](https://github.com/mholt/binding) <span> ★ 464, pushed 133 days ago </span> - Binds form and JSON data from net/http Request to struct.
-   [conform](https://github.com/leebenson/conform) <span> ★ 11, pushed 103 days ago </span> - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
-   [formam](https://github.com/monoculum/formam) <span> ★ 46, pushed 69 days ago </span> - decode form's values into a struct.
-   [forms](https://github.com/albrow/forms) <span> ★ 58, pushed 173 days ago </span> - A framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
-   [gorilla/csrf](https://github.com/gorilla/csrf) <span> ★ 108, pushed 11 days ago </span> - CSRF protection for Go web applications & services.
-   [nosurf](https://github.com/justinas/nosurf) <span> ★ 640, pushed 130 days ago </span> - A CSRF protection middleware for Go.

Game Development
----------------

*Awesome game development libraries.*

-   [engo](https://github.com/EngoEngine/engo) <span> ★ 126, pushed 0 days ago </span> - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.
-   [GarageEngine](https://github.com/vova616/GarageEngine) <span> ★ 223, pushed 965 days ago </span> - 2d game engine written in Go working on OpenGL.
-   [glop](https://github.com/runningwild/glop) <span> ★ 70, pushed 214 days ago </span> - Glop (Game Library Of Power) is a fairly simple cross-platform game library.
-   [go-astar](https://github.com/beefsack/go-astar) <span> ★ 146, pushed 264 days ago </span> - Go implementation of the A\* path finding algorithm
-   [go-collada](https://github.com/GlenKelley/go-collada) <span> ★ 8, pushed 942 days ago </span> - Go package for working with the Collada file format.
-   [go-sdl2](https://github.com/veandco/go-sdl2) <span> ★ 364, pushed 0 days ago </span> - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/) .
-   [go3d](https://github.com/ungerik/go3d) <span> ★ 102, pushed 386 days ago </span> - A performance oriented 2D/3D math package for Go
-   [gonet](https://github.com/xtaci/gonet) <span> ★ 644, pushed 153 days ago </span> - A game server skeleton implemented with golang
-   [Leaf](https://github.com/name5566/leaf) <span> ★ 626, pushed 61 days ago </span> - A lightweight game server framework
-   [termloop](https://github.com/JoelOtter/termloop) <span> ★ 607, pushed 79 days ago </span> - Terminal-based game engine for Go, built on top of Termbox

Generation & Generics
---------------------

*Tools to enhance the language with features like generics via code generation*

-   [gen](https://github.com/clipperhouse/gen) <span> ★ 675, pushed 40 days ago </span> - Code generation tool for ‘generics’-like functionality.
-   [go-linq](https://github.com/ahmetalpbalkan/go-linq) <span> ★ 761, pushed 227 days ago </span> - .NET LINQ-like query methods for Go.
-   [interfaces](https://github.com/rjeczalik/interfaces) <span> ★ 68, pushed 93 days ago </span> - Command line tool for generating interface definitions.
-   [pkgreflect](https://github.com/ungerik/pkgreflect) <span> ★ 31, pushed 496 days ago </span> - A Go preprocessor for package scoped reflection.

Go Compilers
------------

*Tools for compiling Go to other languages*

-   [gopherjs](https://github.com/gopherjs/gopherjs) <span> ★ 3780, pushed 10 days ago </span> - A compiler from Go to JavaScript.
-   [llgo](https://github.com/go-llvm/llgo) <span> ★ 733, pushed 476 days ago </span> - LLVM-based compiler for Go.
-   [tardisgo](https://github.com/tardisgo/tardisgo) <span> ★ 308, pushed 17 days ago </span> - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.

Goroutines
----------

*Tools for managing and working with Goroutines*

-   [grpool](https://github.com/ivpusic/grpool) <span> ★ 99, pushed 159 days ago </span> - Lightweight Goroutine pool.
-   [pool](https://github.com/go-playground/pool) <span> ★ 42, pushed 103 days ago </span> - Go consumer goroutine pool for easy goroutine handling + time saving.
-   [tunny](https://github.com/Jeffail/tunny) <span> ★ 304, pushed 287 days ago </span> - A goroutine pool for golang.

GUI
---

*Libraries for building GUI Applications*

-   [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK
-   [go-qml](https://github.com/go-qml/qml) <span> ★ 1652, pushed 19 days ago </span> - QML support for the Go language
-   [goqt](https://github.com/visualfc/goqt) <span> ★ 984, pushed 60 days ago </span> - Golang bindings to the Qt cross-platform application framework.
-   [gosx-notifier](https://github.com/deckarep/gosx-notifier) <span> ★ 283, pushed 80 days ago </span> - OSX Desktop Notifications library for Go.
-   [gotk3](https://github.com/gotk3/gotk3) <span> ★ 67, pushed 7 days ago </span> - Go bindings for GTK3.
-   [sciter](https://github.com/oskca/sciter) <span> ★ 238, pushed 54 days ago </span> - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development.
-   [systray](https://github.com/getlantern/systray) <span> ★ 70, pushed 21 days ago </span> - Cross platform Go library to place an icon and menu in the notification area
-   [trayhost](https://github.com/shurcooL/trayhost) <span> ★ 36, pushed 140 days ago </span> - Cross-platform Go library to place an icon in the host operating system's taskbar.
-   [ui](https://github.com/andlabs/ui) <span> ★ 2792, pushed 5 days ago </span> - Platform-native GUI library for Go.
-   [walk](https://github.com/lxn/walk) <span> ★ 1279, pushed 26 days ago </span> - Windows application library kit for Go.

Hardware
--------

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

Images
------

*Libraries for manipulating images.*

-   [bimg](https://github.com/h2non/bimg) <span> ★ 156, pushed 4 days ago </span> - Small package for fast and efficient image processing using libvips
-   [geopattern](https://github.com/pravj/geopattern) <span> ★ 847, pushed 96 days ago </span> - Create beautiful generative image patterns from a string.
-   [gift](https://github.com/disintegration/gift) <span> ★ 701, pushed 79 days ago </span> - Package of image processing filters.
-   [go-cairo](https://github.com/ungerik/go-cairo) <span> ★ 51, pushed 185 days ago </span> - Go binding for the cairo graphics library.
-   [go-gd](https://github.com/bolknote/go-gd) <span> ★ 39, pushed 237 days ago </span> - Go binding for GD library
-   [go-nude](https://github.com/koyachi/go-nude) <span> ★ 173, pushed 381 days ago </span> - Nudity detection with Go.
-   [go-opencv](https://github.com/lazywei/go-opencv) <span> ★ 381, pushed 90 days ago </span> - Go bindings for OpenCV.
-   [go-webcolors](https://github.com/jyotiska/go-webcolors) <span> ★ 16, pushed 248 days ago </span> - Port of webcolors library from Python to Go.
-   [imagick](https://github.com/gographics/imagick) <span> ★ 407, pushed 1 days ago </span> - Go binding to ImageMagick's MagickWand C API.
-   [imaginary](https://github.com/h2non/imaginary) <span> ★ 734, pushed 4 days ago </span> - Fast and simple HTTP microservice for image resizing
-   [imaging](https://github.com/disintegration/imaging) <span> ★ 791, pushed 57 days ago </span> - Simple Go image processing package.
-   [img](https://github.com/hawx/img) <span> ★ 76, pushed 360 days ago </span> - A selection of image manipulation tools.
-   [mpo](https://github.com/donatj/mpo) <span> ★ 4, pushed 133 days ago </span> - A decoder and conversion tool for MPO 3D Photos.
-   [picfit](https://github.com/thoas/picfit) <span> ★ 469, pushed 17 days ago </span> - An image resizing server written in Go
-   [resize](https://github.com/nfnt/resize) <span> ★ 1050, pushed 87 days ago </span> - Image resizing for the Go with common interpolation methods.
-   [rez](https://github.com/bamiaux/rez) <span> ★ 115, pushed 55 days ago </span> - Image resizing in pure Go and SIMD.
-   [smartcrop](https://github.com/muesli/smartcrop) <span> ★ 200, pushed 161 days ago </span> - Finds good crops for arbitrary images and crop sizes
-   [svgo](https://github.com/ajstarks/svgo) <span> ★ 719, pushed 28 days ago </span> - Go Language Library for SVG generation.
-   [tga](https://github.com/ftrvxmtrx/tga) <span> ★ 12, pushed 337 days ago </span> - Package tga is a TARGA image format decoder/encoder.

Logging
-------

*Libraries for generating and working with log files.*

-   [glog](https://github.com/golang/glog) <span> ★ 972, pushed 2 days ago </span> - Leveled execution logs for Go.
-   [go-log](https://github.com/siddontang/go-log) <span> ★ 13, pushed 616 days ago </span> - Log lib supports level and multi handlers.
-   [go-log](https://github.com/ian-kent/go-log) <span> ★ 17, pushed 103 days ago </span> - A log4j implementation in Go.
-   [go-logger](https://github.com/apsdehal/go-logger) <span> ★ 154, pushed 209 days ago </span> - Simple logger of Go Programs, with level handlers.
-   [gologger](https://github.com/sadlil/gologger) <span> ★ 17, pushed 78 days ago </span> - Simple easy to use log lib for go, logs in Colored Cosole, Simple Console, File or Elasticsearch.
-   [log](https://github.com/apex/log) <span> ★ 221, pushed 7 days ago </span> - Structured logging package for Go.
-   [log](https://github.com/go-playground/log) <span> ★ 13, pushed 0 days ago </span> - Simple, configurable and scalable Structured Logging for Go.
-   [log-voyage](https://github.com/firstrow/logvoyage) <span> ★ 53, pushed 65 days ago </span> - Full-featured logging saas written in golang.
-   [log15](https://github.com/inconshreveable/log15) <span> ★ 422, pushed 24 days ago </span> - Simple, powerful logging for Go
-   [logex](https://github.com/chzyer/logex) <span> ★ 22, pushed 24 days ago </span> - An golang log lib, supports tracking and level, wrap by standard log lib
-   [logger](https://github.com/azer/logger) <span> ★ 63, pushed 53 days ago </span> - Minimalistic logging library for Go.
-   [logrus](https://github.com/Sirupsen/logrus) <span> ★ 2618, pushed 0 days ago </span> - a structured logger for Go.
-   [logrusly](https://github.com/sebest/logrusly) <span> ★ 6, pushed 78 days ago </span> - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/) .
-   [logutils](https://github.com/hashicorp/logutils) <span> ★ 132, pushed 321 days ago </span> - Utilities for slightly better logging in Go (Golang) extending the standard logger.
-   [logxi](https://github.com/mgutz/logxi) <span> ★ 245, pushed 9 days ago </span> - A 12-factor app logger that is fast and makes you happy.
-   [lumberjack](https://github.com/natefinch/lumberjack) <span> ★ 296, pushed 6 days ago </span> - Simple rolling logger, implements io.WriteCloser.
-   [mlog](https://github.com/jbrodriguez/mlog) <span> ★ 3, pushed 58 days ago </span> - A simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
-   [ozzo-log](https://github.com/go-ozzo/ozzo-log) <span> ★ 40, pushed 79 days ago </span> - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
-   [seelog](https://github.com/cihub/seelog) <span> ★ 647, pushed 5 days ago </span> - logging functionality with flexible dispatching, filtering, and formatting.
-   [slf](https://github.com/ventu-io/slf) <span> ★ 26, pushed 30 days ago </span> - The Structured Logging Facade (SLF) for Go (like SLF4J but structured and for Go)
-   [slog](https://github.com/ventu-io/slog) <span> ★ 16, pushed 30 days ago </span> - The reference implementation of the Structured Logging Facade (SLF) for Go
-   [stdlog](https://github.com/alexcesaro/log) <span> ★ 28, pushed 223 days ago </span> - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
-   [tail](https://github.com/hpcloud/tail) <span> ★ 487, pushed 11 days ago </span> - A Go package striving to emulate the features of the BSD tail program.
-   [xlog](https://github.com/rs/xlog) <span> ★ 65, pushed 6 days ago </span> - A structured logger for `    net/context   ` aware HTTP handlers with flexible dispatching.

Machine Learning
----------------

*Libraries for Machine Learning.*

-   [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang.
-   [CloudForest](https://github.com/ryanbressler/CloudForest) <span> ★ 348, pushed 56 days ago </span> - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
-   [gago](https://github.com/MaxHalford/gago) <span> ★ 126, pushed 4 days ago </span> - Multi-population, flexible, parallel genetic algorithm.
-   [go-fann](https://github.com/white-pony/go-fann) <span> ★ 73, pushed 447 days ago </span> - Go bindings for Fast Artificial Neural Networks(FANN) library.
-   [go-galib](https://github.com/thoj/go-galib) <span> ★ 120, pushed 119 days ago </span> - Genetic Algorithms library written in Go / golang
-   [go-pr](https://github.com/daviddengcn/go-pr) <span> ★ 36, pushed 1052 days ago </span> - Pattern recognition package in Go lang.
-   [gobrain](https://github.com/goml/gobrain) <span> ★ 82, pushed 253 days ago </span> - Neural Networks written in go
-   [godist](https://github.com/e-dard/godist) <span> ★ 6, pushed 350 days ago </span> - Various probability distributions, and associated methods.
-   [goga](https://github.com/tomcraven/goga) <span> ★ 34, pushed 8 days ago </span> - Genetic algorithm library for Go.
-   [GoLearn](https://github.com/sjwhitworth/golearn) <span> ★ 2725, pushed 10 days ago </span> - General Machine Learning library for Go.
-   [golinear](https://github.com/danieldk/golinear) <span> ★ 25, pushed 95 days ago </span> - liblinear bindings for Go
-   [goml](https://github.com/cdipaolo/goml) <span> ★ 460, pushed 57 days ago </span> - On-line Machine Learning in Go
-   [goRecommend](https://github.com/timkaye11/goRecommend) <span> ★ 36, pushed 636 days ago </span> - Recommendation Algorithms library written in Go.
-   [libsvm](https://github.com/datastream/libsvm) <span> ★ 37, pushed 920 days ago </span> - libsvm golang version derived work based on LIBSVM 3.14.
-   [mlgo](https://github.com/NullHypothesis/mlgo) <span> ★ 0, pushed 140 days ago </span> - This project aims to provide minimalistic machine learning algorithms in Go.
-   [neural-go](https://github.com/schuyler/neural-go) <span> ★ 42, pushed 754 days ago </span> - A multilayer perceptron network implemented in Go, with training via backpropagation.
-   [probab](https://github.com/ThePaw/probab) <span> ★ 1, pushed 224 days ago </span> - Probability distribution functions. Bayesian inference. Written in pure Go.
-   [regommend](https://github.com/muesli/regommend) <span> ★ 99, pushed 759 days ago </span> - Recommendation & collaborative filtering engine
-   [shield](https://github.com/eaigner/shield) <span> ★ 77, pushed 40 days ago </span> - Bayesian text classifier with flexible tokenizers and storage backends for Go

Messaging
---------

*Libraries that implement messaging systems*

-   [Centrifugo](https://github.com/centrifugal/centrifugo) <span> ★ 1008, pushed 0 days ago </span> - Real-time messaging (Websockets or SockJS) server in Go.
-   [dbus](https://github.com/godbus/dbus) <span> ★ 81, pushed 5 days ago </span> - Native Go bindings for D-Bus.
-   [emitter](https://github.com/olebedev/emitter) <span> ★ 50, pushed 32 days ago </span> - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
-   [EventBus](https://github.com/asaskevich/EventBus) <span> ★ 138, pushed 125 days ago </span> - The lightweight event bus with async compatibility.
-   [go-longpoll](https://github.com/ventu-io/go-longpoll) <span> ★ 7, pushed 10 days ago </span> - PubSub with long polling.
-   [go-notify](https://github.com/TheCreeper/go-notify) <span> ★ 12, pushed 83 days ago </span> - Native implementation of the freedesktop notification spec.
-   [go-nsq](https://github.com/nsqio/go-nsq) <span> ★ 577, pushed 3 days ago </span> - the official Go package for NSQ
-   [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) <span> ★ 1194, pushed 38 days ago </span> - gopush-cluster is a go push server cluster.
-   [gorush](https://github.com/appleboy/gorush) <span> ★ 41, pushed 1 days ago </span> - A push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm) .
-   [machinery](https://github.com/RichardKnop/machinery) <span> ★ 795, pushed 1 days ago </span> - An asynchronous task queue/job queue based on distributed message passing.
-   [mangos](https://github.com/go-mangos/mangos) <span> ★ 821, pushed 4 days ago </span> - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.
-   [NATS](https://github.com/nats-io/nats) <span> ★ 668, pushed 0 days ago </span> - A lightweight and highly performant publish-subscribe and distributed queueing messaging system.
-   [oplog](https://github.com/dailymotion/oplog) <span> ★ 65, pushed 170 days ago </span> - A generic oplog/replication system for REST APIs
-   [pubsub](https://github.com/tuxychandru/pubsub) - A simple pubsub package for go.
-   [sarama](https://github.com/Shopify/sarama) <span> ★ 954, pushed 0 days ago </span> - A Go library for Apache Kafka.
-   [Uniqush-Push](https://github.com/uniqush/uniqush-push) <span> ★ 812, pushed 14 days ago </span> - A redis backed unified push service for server-side notifications to mobile devices.
-   [zmq4](https://github.com/pebbe/zmq4) <span> ★ 373, pushed 59 days ago </span> - A Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2) .

Miscellaneous
-------------

*These libraries were placed here because none of the other categories seemed to fit*

-   [afero](https://github.com/spf13/afero) <span> ★ 569, pushed 2 days ago </span> - A FileSystem Abstraction System for Go.
-   [archiver](https://github.com/mholt/archiver) <span> ★ 90, pushed 1 days ago </span> - Library and command for making and extracting .zip and .tar.gz archives
-   [autoflags](https://github.com/artyom/autoflags) <span> ★ 12, pushed 313 days ago </span> - Go package to automatically define command line flags from struct fields.
-   [banner](https://github.com/dimiro1/banner) <span> ★ 78, pushed 27 days ago </span> - Add beautiful banners into your Go applications.
-   [browscap\_go](https://github.com/digitalcrab/browscap_go) <span> ★ 15, pushed 102 days ago </span> - GoLang Library for [Browser Capabilities Project](http://browscap.org/) .
-   [datacounter](https://github.com/miolini/datacounter) <span> ★ 4, pushed 126 days ago </span> - Go counters for readers/writer/http.ResponseWriter.
-   [go-chat-bot](https://github.com/go-chat-bot/bot) <span> ★ 59, pushed 26 days ago </span> - IRC, Slack & Telegram bot written in Go.
-   [go-commons-pool](https://github.com/jolestar/go-commons-pool) <span> ★ 237, pushed 93 days ago </span> - A generic object pool for Golang.
-   [go-multierror](https://github.com/hashicorp/go-multierror) <span> ★ 189, pushed 56 days ago </span> - A Go (golang) package for representing a list of errors as a single error.
-   [go-shortid](https://github.com/ventu-io/go-shortid) <span> ★ 51, pushed 67 days ago </span> - Distributed generation of super short, unique, non-sequential, URL friendly IDs.
-   [gopsutil](https://github.com/shirou/gopsutil) <span> ★ 999, pushed 1 days ago </span> - A cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).
-   [gosms](https://github.com/haxpax/gosms) <span> ★ 916, pushed 244 days ago </span> - Your own local SMS gateway in Go that can be used to send SMS
-   [gountries](https://github.com/pariz/gountries) <span> ★ 83, pushed 67 days ago </span> - A package that exposes country and subdivision data.
-   [health](https://github.com/dimiro1/health) <span> ★ 171, pushed 21 days ago </span> - A Easy to use, extensible health check library.
-   [jobs](https://github.com/albrow/jobs) <span> ★ 324, pushed 116 days ago </span> - A persistent and flexible background jobs library.
-   [margelet](https://github.com/zhulik/margelet) <span> ★ 31, pushed 4 days ago </span> - A framework for building Telegram bots.
-   [notify](https://github.com/rjeczalik/notify) <span> ★ 151, pushed 49 days ago </span> - File system event notification library with simple API, similar to os/signal.
-   [stats](https://github.com/go-playground/stats) <span> ★ 17, pushed 201 days ago </span> - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...
-   [werr](https://github.com/txgruppi/werr) <span> ★ 4, pushed 46 days ago </span> - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.
-   [xkg](https://github.com/go-xkg/xkg) <span> ★ 10, pushed 473 days ago </span> - X Keyboard Grabber
-   [xstrings](https://github.com/huandu/xstrings) <span> ★ 315, pushed 147 days ago </span> - A collection of useful string functions ported from other languages.

Natural Language Processing
---------------------------

*Libraries for working with human languages.*

-   [dpar](https://github.com/danieldk/dpar/) - Transition-based statistical dependency parser.
-   [go-eco](https://github.com/ThePaw/go-eco) <span> ★ 1, pushed 224 days ago </span> - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.
-   [go-i18n](https://github.com/nicksnyder/go-i18n/) - A package and an accompanying tool to work with localized text.
-   [go-nlp](https://github.com/nuance/go-nlp) <span> ★ 45, pushed 1623 days ago </span> - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.
-   [go-stem](https://github.com/agonopol/go-stem) <span> ★ 37, pushed 300 days ago </span> - Implementation of the porter stemming algorithm.
-   [go2vec](https://github.com/danieldk/go2vec) <span> ★ 8, pushed 23 days ago </span> - Reader and utility functions for word2vec embeddings.
-   [golibstemmer](https://github.com/rjohnsondev/golibstemmer) <span> ★ 10, pushed 678 days ago </span> - Go bindings for the snowball libstemmer library including porter 2
-   [gounidecode](https://github.com/fiam/gounidecode) <span> ★ 37, pushed 215 days ago </span> - Unicode transliterator (also known as unidecode) for Go
-   [icu](https://github.com/goodsign/icu) <span> ★ 15, pushed 1215 days ago </span> - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
-   [libtextcat](https://github.com/goodsign/libtextcat) <span> ★ 8, pushed 1215 days ago </span> - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
-   [MMSEGO](https://github.com/awsong/MMSEGO) <span> ★ 49, pushed 1468 days ago </span> - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
-   [paicehusk](https://github.com/rookii/paicehusk) <span> ★ 13, pushed 861 days ago </span> - Golang implementation of the Paice/Husk Stemming Algorithm
-   [porter](https://github.com/a2800276/porter) <span> ★ 3, pushed 935 days ago </span> - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.
-   [porter2](https://github.com/zhenjl/porter2) <span> ★ 21, pushed 240 days ago </span> - Really fast Porter 2 stemmer.
-   [segment](https://github.com/blevesearch/segment) <span> ★ 16, pushed 111 days ago </span> - A Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex \#29](http://www.unicode.org/reports/tr29/)
-   [sentences](https://github.com/neurosnap/sentences) <span> ★ 125, pushed 10 days ago </span> - A sentence tokenizer: converts text into a list of sentences.
-   [snowball](https://github.com/goodsign/snowball) <span> ★ 13, pushed 1231 days ago </span> - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/) .
-   [stemmer](https://github.com/dchest/stemmer) <span> ★ 29, pushed 1302 days ago </span> - Stemmer packages for Go programming language. Includes English and German stemmers.
-   [textcat](https://github.com/pebbe/textcat) <span> ★ 41, pushed 315 days ago </span> - A Go package for n-gram based text categorization, with support for utf-8 and raw text

Networking
----------

*Libraries for working with various layers of the network*

-   [arp](https://github.com/mdlayher/arp) <span> ★ 34, pushed 242 days ago </span> - Package arp implements the ARP protocol, as described in RFC 826.
-   [buffstreams](https://github.com/stabbycutyou/buffstreams) <span> ★ 132, pushed 9 days ago </span> - Streaming protocolbuffer data over TCP made easy
-   [canopus](https://github.com/zubairhamed/canopus) <span> ★ 28, pushed 0 days ago </span> - CoAP Client/Server implementation (RFC 7252)
-   [dhcp6](https://github.com/mdlayher/dhcp6) <span> ★ 8, pushed 35 days ago </span> - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
-   [dns](https://github.com/miekg/dns) <span> ★ 1475, pushed 6 days ago </span> - Go library for working with DNS
-   [ether](https://github.com/songgao/ether) <span> ★ 31, pushed 20 days ago </span> - A cross-platform Go package for sending and receiving ethernet frames.
-   [ethernet](https://github.com/mdlayher/ethernet) <span> ★ 9, pushed 242 days ago </span> - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.
-   [fasthttp](https://github.com/valyala/fasthttp) <span> ★ 2240, pushed 0 days ago </span> - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http
-   [ftp](https://github.com/jlaffaye/ftp) <span> ★ 123, pushed 40 days ago </span> - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959) .
-   [go-getter](https://github.com/hashicorp/go-getter) <span> ★ 230, pushed 4 days ago </span> - A Go library for downloading files or directories from various sources using a URL.
-   [go-stun](https://github.com/ccding/go-stun) <span> ★ 72, pushed 0 days ago </span> - A go implementation of the STUN client (RFC 3489 and RFC 5389).
-   [golibwireshark](https://github.com/sunwxg/golibwireshark) <span> ★ 4, pushed 14 days ago </span> - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.
-   [gopacket](https://github.com/google/gopacket) <span> ★ 569, pushed 0 days ago </span> - A Go library for packet processing with libpcap bindings
-   [gopcap](https://github.com/akrennmair/gopcap) <span> ★ 208, pushed 51 days ago </span> - A Go wrapper for libpcap
-   [goshark](https://github.com/sunwxg/goshark) <span> ★ 2, pushed 159 days ago </span> - Package goshark use tshark to decode IP packet and create data struct to analyse packet.
-   [gosnmp](https://github.com/soniah/gosnmp) <span> ★ 104, pushed 26 days ago </span> - Native Go library for performing SNMP actions
-   [gotcp](https://github.com/gansidui/gotcp) <span> ★ 153, pushed 176 days ago </span> - A Go package for quickly writing tcp applications
-   [grab](https://github.com/cavaliercoder/grab) <span> ★ 21, pushed 64 days ago </span> - Go package for managing file downloads
-   [graval](https://github.com/koofr/graval) <span> ★ 8, pushed 101 days ago </span> - An experimental FTP server framework.
-   [linkio](https://github.com/ian-kent/linkio) <span> ★ 17, pushed 483 days ago </span> - Network link speed simulation for Reader/Writer interfaces
-   [llb](https://github.com/kirillDanshin/llb) <span> ★ 1, pushed 21 days ago </span> - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.
-   [mdns](https://github.com/hashicorp/mdns) <span> ★ 230, pushed 26 days ago </span> - Simple mDNS (Multicast DNS) client/server library in Golang
-   [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
-   [portproxy](https://github.com/aybabtme/portproxy) <span> ★ 21, pushed 499 days ago </span> - Simple TCP proxy which adds CORS support to API's which don't support it.
-   [raw](https://github.com/mdlayher/raw) <span> ★ 15, pushed 242 days ago </span> - Package raw enables reading and writing data at the device driver level for a network interface.
-   [sftp](https://github.com/pkg/sftp) <span> ★ 247, pushed 12 days ago </span> - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.
-   [sslb](https://github.com/eduardonunesp/sslb) <span> ★ 60, pushed 17 days ago </span> - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.
-   [tcp\_server](https://github.com/firstrow/tcp_server) <span> ★ 43, pushed 26 days ago </span> - A Go library for building tcp servers faster.
-   [utp](https://github.com/anacrolix/utp) <span> ★ 47, pushed 0 days ago </span> - Go uTP micro transport protocol implementation.

OpenGL
------

*Libraries for using OpenGL in Go.*

-   [gl](https://github.com/go-gl/gl) <span> ★ 196, pushed 50 days ago </span> - Go bindings for OpenGL (generated via glow).
-   [glfw](https://github.com/go-gl/glfw) <span> ★ 230, pushed 5 days ago </span> - Go bindings for GLFW 3.
-   [goxjs/gl](https://github.com/goxjs/gl) <span> ★ 65, pushed 24 days ago </span> - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
-   [goxjs/glfw](https://github.com/goxjs/glfw) <span> ★ 25, pushed 148 days ago </span> - Go cross-platform glfw library for creating an OpenGL context and receiving events.
-   [mathgl](https://github.com/go-gl/mathgl) <span> ★ 116, pushed 2 days ago </span> - Pure Go math package specialized for 3D math, with inspiration from GLM.

ORM
---

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

-   [beego orm](https://github.com/astaxie/beego/tree/master/orm) - A powerful orm framework for go. Support: pq/mysql/sqlite3.
-   [go-store](https://github.com/gosuri/go-store) <span> ★ 70, pushed 138 days ago </span> - A simple and fast Redis backed key-value store library for Go.
-   [gomodel](https://github.com/cosiner/gomodel) <span> ★ 46, pushed 16 days ago </span> - A lightweight, fast, orm-like library helps interactive with database.
-   [GORM](https://github.com/jinzhu/gorm) <span> ★ 3669, pushed 14 days ago </span> - The fantastic ORM library for Golang, aims to be developer friendly.
-   [gorp](https://github.com/go-gorp/gorp) <span> ★ 2051, pushed 33 days ago </span> - Go Relational Persistence, ORM-ish library for Go.
-   [QBS](https://github.com/coocood/qbs) <span> ★ 386, pushed 144 days ago </span> - Stands for Query By Struct. A Go ORM.
-   [Storm](https://github.com/asdine/storm) <span> ★ 134, pushed 1 days ago </span> - Simple and powerful ORM for BoltDB.
-   [upper.io/db](https://github.com/upper/db) <span> ★ 331, pushed 3 days ago </span> - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
-   [Xorm](https://github.com/go-xorm/xorm) <span> ★ 1174, pushed 3 days ago </span> - Simple and powerful ORM for Go.
-   [Zoom](https://github.com/albrow/zoom) <span> ★ 128, pushed 27 days ago </span> - A blazing-fast datastore and querying engine built on Redis.

Package Management
------------------

*Libraries for package and dependency management.*

-   [gigo](https://github.com/LyricalSecurity/gigo) <span> ★ 189, pushed 263 days ago </span> - PIP-like dependency tool for golang, with support for private repositories and hashes.
-   [glide](https://github.com/Masterminds/glide) <span> ★ 1716, pushed 0 days ago </span> - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.
-   [godep](https://github.com/tools/godep) <span> ★ 3389, pushed 10 days ago </span> - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.
-   [gom](https://github.com/mattn/gom) <span> ★ 1078, pushed 37 days ago </span> - Go Manager - bundle for go.
-   [goop](https://github.com/nitrous-io/goop) <span> ★ 759, pushed 145 days ago </span> - A simple dependency manager for Go (golang), inspired by Bundler.
-   [gopm](https://github.com/gpmgo/gopm) <span> ★ 963, pushed 71 days ago </span> - Go Package Manager
-   [gpm](https://github.com/pote/gpm) <span> ★ 944, pushed 96 days ago </span> - Barebones dependency manager for Go.
-   [johnny-deps](https://github.com/VividCortex/johnny-deps) <span> ★ 215, pushed 587 days ago </span> - Minimal dependency version using Git
-   [nut](https://github.com/jingweno/nut) <span> ★ 246, pushed 305 days ago </span> - Vendor Go dependencies
-   [VenGO](https://github.com/DamnWidget/VenGO) <span> ★ 85, pushed 48 days ago </span> - create and manage exportable isolated go virtual environments

Query Language
--------------

-   [graphql](https://github.com/tmc/graphql) <span> ★ 35, pushed 173 days ago </span> - graphql parser + utilities.
-   [graphql](https://github.com/sevki/graphql) <span> ★ 28, pushed 111 days ago </span> - GraphQL implementation in go.
-   [graphql-go](https://github.com/chris-ramon/graphql-go) - An implementation of GraphQL for Go.
-   [jsonql](https://github.com/elgs/jsonql) <span> ★ 50, pushed 25 days ago </span> - JSON query expression library in Golang.

Resource Embedding
------------------

-   [fileb0x](https://github.com/UnnoTed/fileb0x) <span> ★ 43, pushed 84 days ago </span> - Simple tool to embed files in go with focus on "customization" and ease to use.
-   [go-bindata](https://github.com/jteeuwen/go-bindata) <span> ★ 1785, pushed 13 days ago </span> - Package that converts any file into managable Go source code.
-   [go-embed](https://github.com/pyros2097/go-embed) <span> ★ 2, pushed 13 days ago </span> - Generates go code to embed resource files into your library or executable
-   [go-resources](https://github.com/omeid/go-resources) <span> ★ 103, pushed 42 days ago </span> - Unfancy resources embedding with Go.
-   [go.rice](https://github.com/GeertJohan/go.rice) <span> ★ 707, pushed 15 days ago </span> - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.
-   [statics](https://github.com/go-playground/statics) <span> ★ 31, pushed 157 days ago </span> - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.
-   [vfsgen](https://github.com/shurcooL/vfsgen) <span> ★ 78, pushed 14 days ago </span> - Generates a vfsdata.go file that statically implements the given virtual filesystem.

Science and Data Analysis
-------------------------

*Libraries for scientific computing and data analyzing.*

-   [blas](https://github.com/ziutek/blas) <span> ★ 90, pushed 895 days ago </span> - Implementation of BLAS (Basic Linear Algebra Subprograms)
-   [chart](https://github.com/vdobler/chart) <span> ★ 307, pushed 224 days ago </span> - Simple Chart Plotting library for Go. Supports many graphs types.
-   [evaler](https://github.com/soniah/evaler) <span> ★ 21, pushed 62 days ago </span> - A simple floating point arithmetic expression evaluator
-   [ewma](https://github.com/VividCortex/ewma) <span> ★ 153, pushed 349 days ago </span> - Exponentially-weighted moving averages
-   [geom](https://github.com/skelterjohn/geom) <span> ★ 33, pushed 1105 days ago </span> - 2D geometry for golang
-   [go-dsp](https://github.com/mjibson/go-dsp) <span> ★ 382, pushed 222 days ago </span> - Digital Signal Processing for Go
-   [go-fn](https://github.com/ematvey/go-fn) <span> ★ 1, pushed 363 days ago </span> - Mathematical functions written in Go language, that are not covered by math pkg
-   [go-gt](https://github.com/ThePaw/go-gt) <span> ★ 1, pushed 224 days ago </span> - Graph theory algorithms written in "Go" language
-   [go.matrix](https://github.com/skelterjohn/go.matrix) <span> ★ 240, pushed 88 days ago </span> - linear algebra for go (has been stalled)
-   [gocomplex](https://github.com/varver/gocomplex) <span> ★ 1, pushed 305 days ago </span> - A complex number library for the Go programming language.
-   [gofrac](https://github.com/anschelsc/gofrac) <span> ★ 6, pushed 252 days ago </span> - A (goinstallable) fractions library for go with support for basic arithmetic.
-   [gohistogram](https://github.com/VividCortex/gohistogram) <span> ★ 70, pushed 109 days ago </span> - Approximate histograms for data streams
-   [gonum/mat64](https://github.com/gonum/matrix) <span> ★ 246, pushed 0 days ago </span> - The general purpose package for matrix computation. Package mat64 provides basic linear algebra operations for float64 matrices.
-   [gonum/plot](https://github.com/gonum/plot) <span> ★ 241, pushed 3 days ago </span> - gonum/plot provides an API for building and drawing plots in Go.
-   [goraph](https://github.com/gyuho/goraph) <span> ★ 210, pushed 28 days ago </span> - A pure Go graph theory library(data structure, algorith visualization)
-   [gostat](https://github.com/ematvey/gostat) <span> ★ 8, pushed 294 days ago </span> - A statistics library for the go language
-   [mudlark-go](https://github.com/pwil3058/mudlark-go-pkgs) <span> ★ 0, pushed 409 days ago </span> - A collection of packages providing (hopefully) useful code for use in software using Google's Go programming language.
-   [pagerank](https://github.com/alixaxel/pagerank) <span> ★ 10, pushed 50 days ago </span> - Weighted PageRank algorithm implemented in Go
-   [stats](https://github.com/montanaflynn/stats) <span> ★ 586, pushed 9 days ago </span> - A statistics package with common functions missing from the Golang standard library.
-   [streamtools](https://github.com/nytlabs/streamtools) <span> ★ 1271, pushed 283 days ago </span> - general purpose, graphical tool for dealing with streams of data.
-   [vectormath](https://github.com/spate/vectormath) <span> ★ 52, pushed 1129 days ago </span> - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code. (currently inactive)

Security
--------

*Libraries that are used to help make your application more secure.*

-   [acmetool](https://github.com/hlandau/acme) <span> ★ 757, pushed 4 days ago </span> — ACME (Let's Encrypt) client tool with automatic renewal.
-   [BadActor](https://github.com/jaredfolkins/badactor) <span> ★ 169, pushed 329 days ago </span> - An in-memory, application-driven jailer built in the spirit of fail2ban
-   [go-yara](https://github.com/hillu/go-yara) <span> ★ 28, pushed 45 days ago </span> - Go Bindings for [YARA](https://github.com/plusvic/yara) , the "pattern matching swiss knife for malware researchers (and everyone else)"
-   [lego](https://github.com/xenolf/lego) <span> ★ 1101, pushed 0 days ago </span> - Pure Go ACME client library and CLI tool (for use with Let's Encrypt)
-   [passlib](https://github.com/hlandau/passlib) <span> ★ 94, pushed 206 days ago </span> - Futureproof password hashing library.
-   [simple-scrypt](https://github.com/elithrar/simple-scrypt) <span> ★ 60, pushed 30 days ago </span> - an scrypt package with a simple, obvious API and automatic cost calibration built-in.

Serialization
-------------

*Libraries and tools for binary serialization*

-   [asn1](https://github.com/PromonLogicalis/asn1) <span> ★ 2, pushed 49 days ago </span> - Asn.1 BER and DER encoding library for golang
-   [go-capnproto](https://github.com/glycerine/go-capnproto) <span> ★ 241, pushed 85 days ago </span> - Cap'n Proto library and parser for go
    -   [bambam](https://github.com/glycerine/bambam) <span> ★ 53, pushed 120 days ago </span> - generator for Cap'n Proto schemas from go.
-   [go-codec](https://github.com/ugorji/go) <span> ★ 577, pushed 28 days ago </span> - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support
-   [gogoprotobuf](https://github.com/gogo/protobuf) <span> ★ 395, pushed 0 days ago </span> - Protocol Buffers for Go with Gadgets
-   [goprotobuf](https://github.com/golang/protobuf) <span> ★ 879, pushed 0 days ago </span> - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.
-   [mapstructure](https://github.com/mitchellh/mapstructure) <span> ★ 579, pushed 24 days ago </span> - Go library for decoding generic map values into native Go structures.
-   [php *session* decoder](https://github.com/yvasiyarov/php_session_decoder) <span> ★ 50, pushed 139 days ago </span> - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions
-   [structomap](https://github.com/tuvistavie/structomap) <span> ★ 31, pushed 326 days ago </span> - Library to easily and dynamically generate maps from static structures.

Server Applications
-------------------

-   [algernon](https://github.com/xyproto/algernon) <span> ★ 230, pushed 0 days ago </span> - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
-   [Caddy](https://github.com/mholt/caddy) <span> ★ 5613, pushed 1 days ago </span> - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
-   [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
-   [devd](https://github.com/cortesi/devd) <span> ★ 2072, pushed 11 days ago </span> - A local webserver for developers
-   [etcd](https://github.com/coreos/etcd) <span> ★ 9385, pushed 0 days ago </span> - A highly-available key value store for shared configuration and service discovery.
-   [minio](https://github.com/minio/minio) <span> ★ 1103, pushed 0 days ago </span> - Minio is a distributed object storage server.
-   [nsq](http://nsq.io/) - A realtime distributed messaging platform
-   [yakvs](https://github.com/sci4me/yakvs) <span> ★ 6, pushed 258 days ago </span> - A small, networked, in-memory key-value store.

Template Engines
----------------

*Libraries and tools for templating and lexing.*

-   [ace](https://github.com/yosssi/ace) <span> ★ 456, pushed 114 days ago </span> - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.
-   [amber](https://github.com/eknkc/amber) <span> ★ 586, pushed 4 days ago </span> - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.
-   [damsel](https://github.com/dskinner/damsel) <span> ★ 16, pushed 18 days ago </span> - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.
-   [ego](https://github.com/benbjohnson/ego) <span> ★ 250, pushed 74 days ago </span> - A lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
-   [fasttemplate](https://github.com/valyala/fasttemplate) <span> ★ 59, pushed 41 days ago </span> - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/) .
-   [gofpdf](https://github.com/jung-kurt/gofpdf) <span> ★ 215, pushed 47 days ago </span> - A PDF document generator with high level support for text, drawing and images.
-   [kasia.go](https://github.com/ziutek/kasia.go) <span> ★ 68, pushed 238 days ago </span> - Templating system for HTML and other text documents - go implementation.
-   [mustache](https://github.com/hoisie/mustache) <span> ★ 787, pushed 35 days ago </span> - A Go implementation of the Mustache template language.
-   [pongo2](https://github.com/flosch/pongo2) <span> ★ 705, pushed 2 days ago </span> - A Django-like template-engine for Go.
-   [quicktemplate](https://github.com/valyala/quicktemplate) <span> ★ 379, pushed 1 days ago </span> - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.
-   [raymond](https://github.com/aymerick/raymond) <span> ★ 83, pushed 315 days ago </span> - A complete handlebars implementation in Go.
-   [Razor](https://github.com/sipin/gorazor) <span> ★ 505, pushed 301 days ago </span> - Razor view engine for Golang.
-   [Soy](https://github.com/robfig/soy) <span> ★ 112, pushed 35 days ago </span> - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/)

Testing
-------

*Libraries for testing codebases and generating test data.*

-   Testing Frameworks

    -   [assert](https://github.com/go-playground/assert) <span> ★ 5, pushed 80 days ago </span> - Basic Assertion Library used along side native go testing, with building blocks for custom assertions
    -   [assert](https://github.com/bmizerany/assert) <span> ★ 125, pushed 245 days ago </span> - Asserts to Go testing
    -   [badio](https://github.com/cavaliercoder/badio) <span> ★ 2, pushed 72 days ago </span> - Extensions to Go's `      testing/iotest     ` package
    -   [bro](https://github.com/marioidival/bro) <span> ★ 11, pushed 258 days ago </span> - Watch files in directory and run tests for them
    -   [frisby](https://github.com/verdverm/frisby) <span> ★ 132, pushed 40 days ago </span> - a REST API testing framework
    -   [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go
    -   [go-carpet](https://github.com/msoap/go-carpet) <span> ★ 139, pushed 37 days ago </span> - Tool for viewing test coverage in terminal
    -   [go-mutesting](https://github.com/zimmski/go-mutesting) <span> ★ 51, pushed 448 days ago </span> - Mutation testing for Go source code
    -   [go-vcr](https://github.com/dnaeon/go-vcr) <span> ★ 87, pushed 50 days ago </span> - Record and replay your HTTP interactions for fast, deterministic and accurate tests
    -   [goblin](https://github.com/franela/goblin) <span> ★ 267, pushed 93 days ago </span> - Mocha like testing framework fo Go
    -   [gocheck](http://labix.org/gocheck) - A more advanced testing framework alternative to gotest.
    -   [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload
    -   [godog](https://github.com/DATA-DOG/godog) <span> ★ 56, pushed 25 days ago </span> - Cucumber or Behat like BDD framework for Go.
    -   [gofight](https://github.com/appleboy/gofight) <span> ★ 29, pushed 1 days ago </span> - API Handler Testing for Golang Router framework.
    -   [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
    -   [GoSpec](https://github.com/orfjackal/gospec) <span> ★ 108, pushed 634 days ago </span> - BDD-style testing framework for the Go programming language.
    -   [gospecify](https://github.com/stesla/gospecify) <span> ★ 51, pushed 1651 days ago </span> - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
    -   [Hamcrest](https://github.com/rdrdr/hamcrest) <span> ★ 22, pushed 1897 days ago </span> - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
    -   [restit](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test.
    -   [testfixtures](https://github.com/go-testfixtures/testfixtures) <span> ★ 40, pushed 4 days ago </span> - A helper for Rails' like test fixtures to test database applications.
    -   [Testify](https://github.com/stretchr/testify) <span> ★ 1841, pushed 2 days ago </span> - A sacred extension to the standard go testing package.
-   Mock

    -   [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) <span> ★ 63, pushed 5 days ago </span> - Tool for generating self-contained mock objects
    -   [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) <span> ★ 215, pushed 11 days ago </span> - Mock SQL driver for testing database interactions
    -   [go-txdb](https://github.com/DATA-DOG/go-txdb) <span> ★ 5, pushed 292 days ago </span> - Single transaction based database driver mainly for testing purposes.
    -   [gock](https://github.com/h2non/gock) <span> ★ 207, pushed 1 days ago </span> - Versatile HTTP mocking made easy.
    -   [gomock](https://github.com/golang/mock) <span> ★ 250, pushed 42 days ago </span> - Mocking framework for the Go programming language.
    -   [mockhttp](https://github.com/tv42/mockhttp) <span> ★ 19, pushed 544 days ago </span> - Mock object for Go http.ResponseWriter
-   Fuzzing and delta-debugging/reducing/shrinking

    -   [go-fuzz](https://github.com/dvyukov/go-fuzz) <span> ★ 1365, pushed 5 days ago </span> - A randomized testing system
    -   [gofuzz](https://github.com/google/gofuzz) <span> ★ 228, pushed 84 days ago </span> - A library for populating go objects with random values
    -   [gogenerate](https://github.com/arschles/gogenerate) - A Scalacheck-like library for Go
    -   [Tavor](https://github.com/zimmski/tavor) <span> ★ 142, pushed 35 days ago </span> - A generic fuzzing and delta-debugging framework

Text Processing
---------------

*Libraries for parsing and manipulating texts.*

-   Specific Formats
    -   [blackfriday](https://github.com/russross/blackfriday) <span> ★ 1803, pushed 4 days ago </span> - Markdown processor in Go
        -   [github *flavored* markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links.
    -   [bluemonday](https://github.com/microcosm-cc/bluemonday) <span> ★ 369, pushed 93 days ago </span> - HTML Sanitizer
    -   [enca](https://github.com/endeveit/enca) <span> ★ 2, pushed 41 days ago </span> - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/) .
    -   [genex](https://github.com/alixaxel/genex) <span> ★ 36, pushed 43 days ago </span> - Count and expand Regular Expressions into all matching Strings
    -   [go-humanize](https://github.com/dustin/go-humanize) <span> ★ 679, pushed 4 days ago </span> - Formatters for time, numbers, and memory size to human readable format.
    -   [go-nmea](https://github.com/adrianmo/go-nmea) <span> ★ 14, pushed 245 days ago </span> - NMEA parser library for the Go language.
    -   [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) <span> ★ 266, pushed 19 days ago </span> - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    -   [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) <span> ★ 108, pushed 146 days ago </span> - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
    -   [go-runewidth](https://github.com/mattn/go-runewidth) <span> ★ 52, pushed 33 days ago </span> - Functions to get fixed width of the character or string.
    -   [gofeed](https://github.com/mmcdole/gofeed) <span> ★ 423, pushed 0 days ago </span> - Parse RSS and Atom feeds in Go
    -   [gographviz](https://github.com/awalterschulze/gographviz) <span> ★ 39, pushed 249 days ago </span> - Parses the Graphviz DOT language.
    -   [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - Format bytes to string.
    -   [gonameparts](https://github.com/polera/gonameparts) <span> ★ 16, pushed 38 days ago </span> - Parses human names into individual name parts
    -   [GoQuery](https://github.com/PuerkitoBio/goquery) <span> ★ 2578, pushed 17 days ago </span> - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.
    -   [goregen](https://github.com/zach-klippenstein/goregen) <span> ★ 16, pushed 53 days ago </span> - A library for generating random strings from regular expressions.
    -   [guesslanguage](https://github.com/endeveit/guesslanguage) <span> ★ 18, pushed 496 days ago </span> - Functions to determine the natural language of a unicode text.
    -   [mxj](https://github.com/clbanning/mxj) <span> ★ 113, pushed 11 days ago </span> - Encode / decode XML as JSON or map\[string\]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.
    -   [slug](https://github.com/gosimple/slug) <span> ★ 76, pushed 35 days ago </span> - URL-friendly slugify with multiple languages support.
    -   [Slugify](https://github.com/avelino/slugify) <span> ★ 11, pushed 378 days ago </span> - A Go slugify application that handles string.
    -   [toml](https://github.com/BurntSushi/toml) <span> ★ 850, pushed 28 days ago </span> - TOML configuration format (encoder/decoder with reflection).
-   Utility
    -   [gotabulate](https://github.com/bndr/gotabulate) <span> ★ 118, pushed 56 days ago </span> - Easily pretty-print your tabular data with Go.
    -   [kace](https://github.com/codemodus/kace) <span> ★ 2, pushed 285 days ago </span> - Common case conversions covering common initialisms.
    -   [parseargs-go](https://github.com/nproc/parseargs-go) <span> ★ 3, pushed 46 days ago </span> - A string argument parser that understands quotes and backslashes
    -   [parth](https://github.com/codemodus/parth) <span> ★ 6, pushed 18 days ago </span> - URL path segmentation parsing.
    -   [xurls](https://github.com/mvdan/xurls) <span> ★ 185, pushed 2 days ago </span> - Extract urls from text

Third-party APIs
----------------

*Libraries for accessing third party APIs.*

-   [anaconda](https://github.com/ChimeraCoder/anaconda) <span> ★ 523, pushed 20 days ago </span> - A Go client library for the Twitter 1.1 API
-   [aws-sdk-go](https://github.com/aws/aws-sdk-go) <span> ★ 2461, pushed 0 days ago </span> - The official AWS SDK for the Go programming language.
-   [brewerydb](https://github.com/naegelejd/brewerydb) <span> ★ 10, pushed 312 days ago </span> - Go library for accessing the BreweryDB API.
-   [clarifai](https://github.com/samuelcouch/clarifai) - A Go client library for interfacing with the Clarifai API.
-   [discordgo](https://github.com/bwmarrin/discordgo) <span> ★ 76, pushed 0 days ago </span> - Go bindings for the Discord Chat API
-   [facebook](https://github.com/huandu/facebook) <span> ★ 303, pushed 191 days ago </span> - Go Library that supports the Facebook Graph API
-   [gads](https://github.com/emiddleton/gads) <span> ★ 15, pushed 196 days ago </span> - Google Adwords Unofficial API
-   [gami](https://github.com/bit4bit/gami) <span> ★ 14, pushed 213 days ago </span> - Go library for Asterisk Manager Interface.
-   [gcm](https://github.com/Aorioli/gcm) <span> ★ 27, pushed 143 days ago </span> - Go library for Google Cloud Messaging
-   [geo-golang](https://github.com/codingsince1985/geo-golang) <span> ★ 118, pushed 0 days ago </span> - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro) , [MapQuest](http://open.mapquestapi.com/geocoding/) , [Nominatim](http://open.mapquestapi.com/nominatim/) , [OpenCage](http://geocoder.opencagedata.com/api.html) , [HERE](https://developer.here.com/rest-apis/documentation/geocoder) , [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx) , and [Mapbox](https://www.mapbox.com/developers/api/geocoding/) geocoding / reverse geocoding APIs.
-   [ghost](https://github.com/neuegram/ghost) <span> ★ 16, pushed 123 days ago </span> - Go Library for accessing the Snapchat API.
-   [github](https://github.com/google/go-github) <span> ★ 1650, pushed 0 days ago </span> - Go library for accessing the GitHub API.
-   [go-jira](https://github.com/andygrunwald/go-jira) <span> ★ 18, pushed 2 days ago </span> - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)
-   [go-marathon](https://github.com/gambol99/go-marathon) <span> ★ 70, pushed 7 days ago </span> - A Go library for interacting with Mesosphere's Marathon PAAS.
-   [go-trending](https://github.com/andygrunwald/go-trending) <span> ★ 70, pushed 228 days ago </span> - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
-   [go-twitter](https://github.com/dghubble/go-twitter) <span> ★ 55, pushed 15 days ago </span> - Go client library for the Twitter v1.1 APIs.
-   [go-xkcd](https://github.com/nishanths/go-xkcd) <span> ★ 11, pushed 38 days ago </span> - Go client for the xkcd API.
-   [goamz](https://github.com/mitchellh/goamz) <span> ★ 622, pushed 19 days ago </span> - Popular fork of [goamz](https://launchpad.net/goamz) which adds some missing API calls to certain packages.
-   [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) <span> ★ 17, pushed 195 days ago </span> - a Go MusicBrainz WS2 client library.
-   [google](https://github.com/google/google-api-go-client) <span> ★ 625, pushed 17 days ago </span> - Auto-generated Google APIs for Go.
-   [google-analytics](https://github.com/chonthu/go-google-analytics) <span> ★ 4, pushed 321 days ago </span> - A simple wrapper for easy google analytics reporting.
-   [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) <span> ★ 297, pushed 3 days ago </span> - Google Cloud APIs Go Client Library.
-   [gostorm](https://github.com/jsgilmore/gostorm) <span> ★ 72, pushed 19 days ago </span> - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.
-   [hipchat](https://github.com/andybons/hipchat) <span> ★ 98, pushed 32 days ago </span> - This project implements a golang client library for the Hipchat API.
-   [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) <span> ★ 93, pushed 68 days ago </span> - A golang package to communicate with HipChat over XMPP.
-   [Medium](https://github.com/Medium/medium-sdk-go) <span> ★ 43, pushed 12 days ago </span> - A Golang SDK for Medium's OAuth2 API.
-   [megos](https://github.com/andygrunwald/megos) <span> ★ 33, pushed 29 days ago </span> - A client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster
-   [minio-go](https://github.com/minio/minio-go) <span> ★ 91, pushed 0 days ago </span> - Minio Go Library for Amazon S3 compatible cloud storage.
-   [mixpanel](https://github.com/dukex/mixpanel) <span> ★ 12, pushed 23 days ago </span> - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
-   [paypal](https://github.com/logpacker/paypalsdk) - Wrapper for PayPal payment API
-   [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) <span> ★ 0, pushed 50 days ago </span> - The Playlyfe Rest API Go SDK
-   [pushover](https://github.com/gregdel/pushover) <span> ★ 10, pushed 241 days ago </span> - Go wrapper for the Pushover API.
-   [rrdaclient](https://github.com/Omie/rrdaclient) <span> ★ 3, pushed 584 days ago </span> - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.
-   [shopify](https://github.com/rapito/go-shopify) <span> ★ 9, pushed 84 days ago </span> - Go Library to make CRUD request to the Shopify API.
-   [slack](https://github.com/nlopes/slack) <span> ★ 419, pushed 0 days ago </span> - Slack API in Go.
-   [smite](https://github.com/sergiotapia/smitego) <span> ★ 6, pushed 647 days ago </span> - Go package to wraps access to the Smite game API.
-   [spotify](https://github.com/rapito/go-spotify) <span> ★ 6, pushed 378 days ago </span> - Go Library to access Spotify WEB API.
-   [steam](https://github.com/sostronk/go-steam) <span> ★ 7, pushed 348 days ago </span> - Go Library to interact with Steam game servers.
-   [stripe](https://github.com/stripe/stripe-go) <span> ★ 399, pushed 2 days ago </span> - Go client for the Stripe API
-   [telebot](https://github.com/tucnak/telebot) <span> ★ 161, pushed 10 days ago </span> - Telegram bot framework written in Go.
-   [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client.
-   [textbelt](https://github.com/dietsche/textbelt) <span> ★ 4, pushed 234 days ago </span> - Go client for the textbelt.com txt messaging API.
-   [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) <span> ★ 4, pushed 174 days ago </span> - A simple golang package to communicate with [themoviedb.org](https://themoviedb.org)
-   [translate](https://github.com/poorny/translate) - Go online translation package.
-   [tumblr](https://github.com/mattcunningham/gumblr) <span> ★ 0, pushed 287 days ago </span> - Go wrapper for the Tumblr v2 API.
-   [webhooks](https://github.com/go-playground/webhooks) <span> ★ 14, pushed 67 days ago </span> - Webhook reciever for GitHub and Bitbucket.

Utilities
---------

*General utilities and tools to make your life easier.*

-   [abutil](https://github.com/bahlo/abutil) <span> ★ 5, pushed 236 days ago </span> - A collection of often-used Golang helpers.
-   [apm](https://github.com/topfreegames/apm) <span> ★ 33, pushed 46 days ago </span> - A process manager for Golang applications with an HTTP API.
-   [boilr](https://github.com/tmrts/boilr) <span> ★ 47, pushed 50 days ago </span> - A blazingly fast CLI tool for creating projects from boilerplate templates.
-   [command](https://github.com/txgruppi/command) <span> ★ 1, pushed 5 days ago </span> - Command pattern for Go with thread safe serial and parallel dispatcher
-   [coop](https://github.com/rakyll/coop) <span> ★ 1142, pushed 202 days ago </span> - Cheat sheet for some of the common concurrent flows in Go.
-   [Death](https://github.com/vrecan/death) <span> ★ 24, pushed 51 days ago </span> - Managing go application shutdown with signals.
-   [Deepcopier](https://github.com/ulule/deepcopier) <span> ★ 81, pushed 7 days ago </span> - Simple struct copying for Go.
-   [delve](https://github.com/derekparker/delve) <span> ★ 3556, pushed 0 days ago </span> - Go debugger.
-   [fastlz](https://github.com/digitalcrab/fastlz) <span> ★ 4, pushed 440 days ago </span> - Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang.
-   [filetype](https://github.com/h2non/filetype) <span> ★ 55, pushed 125 days ago </span> - Small package to infer the file type checking the magic numbers signature.
-   [fzf](https://github.com/junegunn/fzf) <span> ★ 4715, pushed 0 days ago </span> - A command-line fuzzy finder written in Go
-   [generate](https://github.com/go-playground/generate) <span> ★ 1, pushed 103 days ago </span> - runs go generate recursively on a specified path or environment variable and can filter by regex.
-   [gentleman](https://github.com/h2non/gentleman) <span> ★ 263, pushed 4 days ago </span> - Full-featured plugin-driven HTTP client library.
-   [go-cron](https://github.com/rk/go-cron) <span> ★ 114, pushed 346 days ago </span> - A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
-   [go-debug](https://github.com/tj/go-debug) <span> ★ 313, pushed 116 days ago </span> - Conditional debug logging for Golang libraries & applications.
-   [go-dry](https://github.com/ungerik/go-dry) <span> ★ 146, pushed 14 days ago </span> - DRY (don't repeat yourself) package for Go.
-   [go-rate](https://github.com/beefsack/go-rate) <span> ★ 187, pushed 608 days ago </span> - A timed rate limiter for Go.
-   [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) <span> ★ 11, pushed 11 days ago </span> - XML Sitemap generator written in Go.
-   [go-trigger](https://github.com/sadlil/go-trigger) <span> ★ 43, pushed 54 days ago </span> - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.
-   [go-underscore](https://github.com/tobyhede/go-underscore) <span> ★ 840, pushed 301 days ago </span> - A useful collection of helpfully functional Go collection utilities.
-   [goback](https://github.com/carlescere/goback) <span> ★ 24, pushed 408 days ago </span> - Go simple exponential backoff package.
-   [godaemon](https://github.com/VividCortex/godaemon) <span> ★ 251, pushed 228 days ago </span> - Utility to write daemons.
-   [godotenv](https://github.com/joho/godotenv) <span> ★ 282, pushed 64 days ago </span> - A Go port of Ruby's dotenv library (Loads environment variables from `    .env   ` .)
-   [godropbox](https://github.com/dropbox/godropbox) <span> ★ 3031, pushed 60 days ago </span> - Common libraries for writing Go services/applications from Dropbox.
-   [gohper](https://github.com/cosiner/gohper) <span> ★ 211, pushed 4 days ago </span> - Various tools/modules help for development.
-   [gojq](https://github.com/elgs/gojq) <span> ★ 33, pushed 4 days ago </span> - JSON query in Golang.
-   [golarm](https://github.com/msempere/golarm) <span> ★ 16, pushed 245 days ago </span> - Fire alarms with system events.
-   [golog](https://github.com/mlimaloureiro/golog) <span> ★ 15, pushed 83 days ago </span> - Easy and lightweight CLI tool to time track your tasks.
-   [gopencils](https://github.com/bndr/gopencils) <span> ★ 342, pushed 95 days ago </span> - Small and simple package to easily consume REST APIs.
-   [goplaceholder](https://github.com/michiwend/goplaceholder) <span> ★ 10, pushed 99 days ago </span> - a small golang lib to generate placeholder images.
-   [goreq](https://github.com/franela/goreq) <span> ★ 452, pushed 2 days ago </span> - Minimal and simple request library for Go language.
-   [goreq](https://github.com/smallnest/goreq) <span> ★ 16, pushed 68 days ago </span> - An enhanced simplified HTTP client based on gorequest.
-   [gorequest](https://github.com/parnurzeal/gorequest) <span> ★ 669, pushed 44 days ago </span> - Simplified HTTP client with rich features for Go.
-   [gotenv](https://github.com/subosito/gotenv) <span> ★ 49, pushed 568 days ago </span> - Load environment variables from `    .env   ` or any `    io.Reader   ` in Go
-   [grequests](https://github.com/levigross/grequests) <span> ★ 709, pushed 15 days ago </span> - An elegant and simple `    net/http   ` wrapper that follows Python's requests library
-   [htcat](https://github.com/htcat/htcat) <span> ★ 421, pushed 322 days ago </span> - Parallel and Pipelined HTTP GET Utility
-   [httpcontrol](https://github.com/facebookgo/httpcontrol) <span> ★ 404, pushed 282 days ago </span> - Package httpcontrol allows for HTTP transport level control around timeouts and retries.
-   [hystrix-go](https://github.com/afex/hystrix-go) <span> ★ 387, pushed 57 days ago </span> - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.
-   [JobRunner](https://github.com/bamzi/jobrunner) <span> ★ 265, pushed 178 days ago </span> - Smart and featureful cron job scheduler with job queuing and live monitoring built in.
-   [jsonf](https://github.com/miolini/jsonf) <span> ★ 25, pushed 336 days ago </span> - Console tool for highlighted formatting and struct query fetching JSON.
-   [jsongo](https://github.com/ricardolonga/jsongo) <span> ★ 48, pushed 226 days ago </span> - Fluent API to make it easier to create Json objects.
-   [lrserver](https://github.com/jaschaephraim/lrserver) <span> ★ 67, pushed 206 days ago </span> - LiveReload server for Go
-   [mc](https://github.com/minio/mc) <span> ★ 221, pushed 0 days ago </span> - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.
-   [mergo](https://github.com/imdario/mergo) <span> ★ 124, pushed 47 days ago </span> - A helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.
-   [moldova](https://github.com/StabbyCutyou/moldova) <span> ★ 116, pushed 9 days ago </span> - A utility for generating random data based on an input template.
-   [mp](https://github.com/sanbornm/mp) <span> ★ 11, pushed 679 days ago </span> - A simple cli email parser. It currently takes stdin and outputs JSON.
-   [multitick](https://github.com/VividCortex/multitick) <span> ★ 44, pushed 707 days ago </span> - Multiplexor for aligned tickers.
-   [netbug](https://github.com/e-dard/netbug) <span> ★ 30, pushed 179 days ago </span> - Easy remote profiling of your services.
-   [ngrok](https://github.com/inconshreveable/ngrok) <span> ★ 6940, pushed 16 days ago </span> - Introspected tunnels to localhost.
-   [okrun](https://github.com/xta/okrun) <span> ★ 11, pushed 567 days ago </span> - go run error steamroller.
-   [panicparse](https://github.com/maruel/panicparse) <span> ★ 1259, pushed 39 days ago </span> - Groups similar goroutines and colorizes stack dump.
-   [peco](https://github.com/peco/peco) <span> ★ 2942, pushed 33 days ago </span> - Simplistic interactive filtering tool.
-   [pester](https://github.com/sethgrid/pester) <span> ★ 65, pushed 3 days ago </span> - Go HTTP client calls with retries, backoff, and concurrency.
-   [pm](https://github.com/VividCortex/pm) <span> ★ 39, pushed 591 days ago </span> - Process (i.e. goroutine) manager with an HTTP API.
-   [profile](https://github.com/davecheney/profile) <span> ★ 420, pushed 97 days ago </span> - Simple profiling support package for Go.
-   [request](https://github.com/mozillazg/request) <span> ★ 121, pushed 61 days ago </span> - Go HTTP Requests for Humans™.
-   [rerun](https://github.com/ivpusic/rerun) <span> ★ 102, pushed 82 days ago </span> - Recompiling and rerunning go apps when source changes.
-   [resty](https://github.com/go-resty/resty) <span> ★ 281, pushed 81 days ago </span> - Simple HTTP and REST client for Go inspired by Ruby rest-client.
-   [robustly](https://github.com/VividCortex/robustly) <span> ★ 103, pushed 347 days ago </span> - Runs functions resiliently, catching and restarting panics.
-   [scheduler](https://github.com/carlescere/scheduler) <span> ★ 87, pushed 41 days ago </span> - Cronjobs scheduling made easy.
-   [sling](https://github.com/dghubble/sling) <span> ★ 294, pushed 0 days ago </span> - Go HTTP requests builder for API clients.
-   [spinner](https://github.com/briandowns/spinner) <span> ★ 227, pushed 57 days ago </span> - Go package to easily provide a terminal spinner with options.
-   [sqlx](https://github.com/jmoiron/sqlx) <span> ★ 1723, pushed 37 days ago </span> - provides a set of extensions on top of the excellent built-in database/sql package.
-   [ugo](https://github.com/alxrm/ugo) <span> ★ 7, pushed 34 days ago </span> - ugo is slice toolbox with concise syntax for Go.
-   [xlsx](https://github.com/tealeg/xlsx) <span> ★ 962, pushed 24 days ago </span> - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.

Validation
----------

*Libraries for validation.*

-   [govalidator](https://github.com/asaskevich/govalidator) <span> ★ 1031, pushed 2 days ago </span> - Validators and sanitizers for strings, numerics, slices and structs.
-   [validator](https://github.com/go-playground/validator) <span> ★ 461, pushed 62 days ago </span> - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.

Version Control
---------------

*Libraries for version control.*

-   [gh](https://github.com/rjeczalik/gh) <span> ★ 37, pushed 221 days ago </span> - Scriptable server and net/http middleware for GitHub Webhooks.
-   [git2go](https://github.com/libgit2/git2go) <span> ★ 786, pushed 2 days ago </span> - Go bindings for libgit2.
-   [go-vcs](https://github.com/sourcegraph/go-vcs) <span> ★ 41, pushed 70 days ago </span> - manipulate and inspect VCS repositories in Go.
-   [hgo](https://github.com/beyang/hgo) <span> ★ 7, pushed 244 days ago </span> - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.

Video
-----

*Libraries for manipulating video.*

-   [aac/h264](https://github.com/nareix/codec) <span> ★ 161, pushed 4 days ago </span> - Golang aac/h264 encoder and decoder.
-   [gmf](https://github.com/3d0c/gmf) <span> ★ 185, pushed 41 days ago </span> - Go bindings for FFmpeg av\* libraries.
-   [goav](https://github.com/giorgisio/goav) <span> ★ 131, pushed 56 days ago </span> - Comphrensive Go bindings for FFmpeg.
-   [gst](https://github.com/ziutek/gst) <span> ★ 109, pushed 170 days ago </span> - Go bindings for GStreamer.

Web Frameworks
--------------

*Full stack web frameworks.*

-   [Beego](https://github.com/astaxie/beego) <span> ★ 6885, pushed 0 days ago </span> - beego is an open-source, high-performance web framework for the Go programming language.
-   [Bone](https://github.com/go-zoo/bone) <span> ★ 894, pushed 6 days ago </span> - Lightning Fast HTTP Multiplexer.
-   [chi](https://github.com/pressly/chi) <span> ★ 296, pushed 5 days ago </span> - Small, fast and expressive HTTP router built on net/context.
-   [Echo](https://github.com/labstack/echo) <span> ★ 4178, pushed 0 days ago </span> - A fast and unfancy micro web framework for Go.
-   [Gin](https://github.com/gin-gonic/gin) <span> ★ 6370, pushed 1 days ago </span> - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
-   [Gizmo](https://github.com/NYTimes/gizmo) <span> ★ 1354, pushed 3 days ago </span> - Microservice toolkit used by the New York Times.
-   [Glue](https://github.com/desertbit/glue) <span> ★ 146, pushed 36 days ago </span> - Robust Go and Javascript Socket Library (Alternative to Socket.io).
-   [go-json-rest](https://github.com/ant0ine/go-json-rest) <span> ★ 2288, pushed 15 days ago </span> - A quick and easy way to setup a RESTful JSON API.
-   [go-kit](https://github.com/go-kit/kit) <span> ★ 3885, pushed 0 days ago </span> - A Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.
-   [go-relax](https://github.com/codehack/go-relax) <span> ★ 120, pushed 30 days ago </span> - A framework of pluggable components to build RESTful API's.
-   [go-rest](https://github.com/ungerik/go-rest) <span> ★ 88, pushed 213 days ago </span> - A small and evil REST framework for Go.
-   [go-socket.io](https://github.com/googollee/go-socket.io) <span> ★ 1115, pushed 8 days ago </span> - socket.io library for golang, a realtime application framework.
-   [goa](https://github.com/raphael/goa) - Framework for developing microservices based on the design of Ruby's Praxis.
-   [Goat](https://github.com/bahlo/goat) <span> ★ 83, pushed 19 days ago </span> - A minimalistic REST API server in Go.
-   [gocraft/web](https://github.com/gocraft/web) <span> ★ 1040, pushed 116 days ago </span> - A mux and middleware package in Go.
-   [Goji](https://github.com/goji/goji) <span> ★ 144, pushed 22 days ago </span> - Goji is a minimalistic and flexible HTTP request multiplexer with support for `    net/context   ` .
-   [Golf](https://github.com/dinever/golf) <span> ★ 121, pushed 0 days ago </span> - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.
-   [golongpoll](https://github.com/jcuga/golongpoll) <span> ★ 278, pushed 74 days ago </span> - HTTP longpoll server library that makes web pub-sub simple.
-   [Gondola](https://github.com/rainycape/gondola) <span> ★ 297, pushed 142 days ago </span> - The web framework for writing faster sites, faster
-   [goose](https://github.com/ian-kent/goose) <span> ★ 20, pushed 491 days ago </span> - Server Sent Events in Go
-   [Gorilla](https://github.com/gorilla/) - Gorilla is a web toolkit for the Go programming language.
-   [httprouter](https://github.com/julienschmidt/httprouter) <span> ★ 3072, pushed 4 days ago </span> - A high performance router. Use this and the standard http handlers to form a very high performance web framework.
-   [httptreemux](https://github.com/dimfeld/httptreemux) <span> ★ 145, pushed 20 days ago </span> - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.
-   [Iris](https://kataras.github.io/iris) - A very minimal but flexible and high-performance golang web application framework, providing a robust set of features for building web applications.
-   [lars](https://github.com/go-playground/lars) <span> ★ 276, pushed 12 days ago </span> - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.
-   [Macaron](https://github.com/go-macaron/macaron) <span> ★ 1079, pushed 39 days ago </span> - Macaron is a high productive and modular design web framework in Go.
-   [mango](https://github.com/paulbellamy/mango) <span> ★ 295, pushed 1078 days ago </span> - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
-   [medeina](https://github.com/imdario/medeina) <span> ★ 15, pushed 464 days ago </span> - Medeina is a HTTP routing tree based on HttpRouter, inspired by Roda and Cuba.
-   [mux](https://github.com/gorilla/mux) <span> ★ 2281, pushed 24 days ago </span> - A powerful URL router and dispatcher for golang.
-   [neo](https://github.com/ivpusic/neo) <span> ★ 251, pushed 37 days ago </span> - Neo is minimal and fast Go Web Framework with extremely simple API.
-   [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) <span> ★ 64, pushed 48 days ago </span> - A high-performance HTTP router and Web framework supporting routes with regular expressions. Comes with full support for quickly building a RESTful API application.
-   [pat](https://github.com/bmizerany/pat) <span> ★ 940, pushed 68 days ago </span> - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.
-   [Resoursea](https://github.com/resoursea/api) <span> ★ 24, pushed 449 days ago </span> - A REST framework for quickly writing resource based services.
-   [Revel](https://github.com/revel/revel) <span> ★ 6692, pushed 0 days ago </span> - A high-productivity web framework for the Go language.
-   [rex](https://github.com/goanywhere/rex) <span> ★ 10, pushed 160 days ago </span> - Rex is a library for modular development built upon gorilla/mux, fully compatible with `    net/http   ` .
-   [sawsij](http://sawsij.com/) - lightweight, open-source web framework for building high-performance, data-driven web applications.
-   [Siesta](https://github.com/VividCortex/siesta) <span> ★ 337, pushed 14 days ago </span> - Composable framework to write middleware and handlers
-   [tango](https://github.com/lunny/tango) <span> ★ 393, pushed 10 days ago </span> - Micro & pluggable web framework for Go.
-   [tigertonic](https://github.com/rcrowley/go-tigertonic) <span> ★ 929, pushed 90 days ago </span> - A Go framework for building JSON web services inspired by Dropwizard
-   [traffic](https://github.com/pilu/traffic) <span> ★ 472, pushed 151 days ago </span> - Sinatra inspired regexp/pattern mux and web framework for Go.
-   [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and ouput handling.
-   [vestigo](https://github.com/husobee/vestigo) <span> ★ 33, pushed 34 days ago </span> - A performant, stand-alone, HTTP compliant URL Router for go web applications.
-   [Volatile](https://github.com/volatile/core) <span> ★ 79, pushed 15 days ago </span> - Minimalist middleware stack promoting flexibility, good practices and clean code.
-   [web.go](https://github.com/hoisie/web) <span> ★ 2641, pushed 5 days ago </span> - A simple framework to write webapps in Go.
-   [xmux](https://github.com/rs/xmux) <span> ★ 49, pushed 11 days ago </span> - A high performance muxer based on `    httprouter   ` with `    net/context   ` support.
-   [Zerver](https://github.com/cosiner/zerver) <span> ★ 125, pushed 10 days ago </span> - Zerver is an expressive, modular, feature completed RESTful framework.
-   [zeus](https://github.com/daryl/zeus) <span> ★ 92, pushed 316 days ago </span> - A very simple and fast HTTP router for Go.

### Middlewares

#### Actual middlewares

-   [CORS](https://github.com/rs/cors) <span> ★ 325, pushed 24 days ago </span> - Easily add CORS capabilities to your API.
-   [formjson](https://github.com/rs/formjson) <span> ★ 18, pushed 130 days ago </span> - Transparently handle JSON input as a standard form POST.
-   [Limiter](https://github.com/ulule/limiter) <span> ★ 214, pushed 75 days ago </span> - Dead simple rate limit middleware for Go.
-   [Tollbooth](https://github.com/didip/tollbooth) <span> ★ 368, pushed 25 days ago </span> - Rate limit HTTP request handler.
-   [XFF](https://github.com/sebest/xff) <span> ★ 45, pushed 183 days ago </span> - Handle `    X-Forwarded-For   ` header and friends.

#### Libraries for creating HTTP middlewares

-   [alice](https://github.com/justinas/alice) <span> ★ 887, pushed 22 days ago </span> - Painless middleware chaining for Go.
-   [catena](https://github.com/codemodus/catena) <span> ★ 5, pushed 190 days ago </span> - http.Handler wrapper catenation (same API as "chain").
-   [chain](https://github.com/codemodus/chain) <span> ★ 47, pushed 165 days ago </span> - Handler wrapper chaining with scoped data (net/context-based "middleware").
-   [go-wrap](https://github.com/go-on/wrap) <span> ★ 50, pushed 497 days ago </span> - Small middlewares package for net/http.
-   [gores](https://github.com/alioygur/gores) <span> ★ 39, pushed 48 days ago </span> - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
-   [httpware](https://github.com/nstogner/httpware) <span> ★ 13, pushed 6 days ago </span> - Stackable middleware (using net/context) with easy chaining.
-   [interpose](https://github.com/carbocation/interpose) <span> ★ 243, pushed 434 days ago </span> - Minimalist net/http middleware for golang.
-   [muxchain](https://github.com/stephens2424/muxchain) <span> ★ 198, pushed 547 days ago </span> - Lightweight middleware for net/http.
-   [negroni](https://github.com/codegangsta/negroni) <span> ★ 3581, pushed 6 days ago </span> - Idiomatic HTTP middleware for Golang.
-   [render](https://github.com/unrolled/render) <span> ★ 663, pushed 54 days ago </span> - Go package for easily rendering JSON, XML, and HTML template responses.
-   [stats](https://github.com/thoas/stats) <span> ★ 335, pushed 11 days ago </span> - A Go middleware that stores various information about your web application.

Tools
=====

Go software and plugins.

Code Analysis
-------------

-   [dupl](https://github.com/mibk/dupl) <span> ★ 50, pushed 17 days ago </span> - A tool for code clone detection.
-   [errcheck](https://github.com/kisielk/errcheck) <span> ★ 555, pushed 9 days ago </span> - Errcheck is a program for checking for unchecked errors in Go programs.
-   [gcvis](https://github.com/davecheney/gcvis) <span> ★ 511, pushed 0 days ago </span> - Visualise Go program GC trace data in real time.
-   [Go Metalinter](https://github.com/alecthomas/gometalinter) <span> ★ 862, pushed 16 days ago </span> - Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form.
-   [go-checkstyle](https://github.com/qiniu/checkstyle) <span> ★ 28, pushed 639 days ago </span> checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments.
-   [go-outdated](https://github.com/firstrow/go-outdated) <span> ★ 21, pushed 31 days ago </span> - Console application that displays outdated packages.
-   [goast-viewer](https://github.com/yuroyoro/goast-viewer) <span> ★ 139, pushed 657 days ago </span> - Web based Golang AST visualizer.
-   [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
-   [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
-   [GoLint](https://github.com/golang/lint) <span> ★ 1320, pushed 9 days ago </span> - Golint is a linter for Go source code.
-   [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
-   [goreturns](https://sourcegraph.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
-   [gostatus](https://github.com/shurcooL/gostatus) <span> ★ 132, pushed 37 days ago </span> - A command line tool, shows the status of repositories that contain Go packages.
-   [interfacer](https://github.com/mvdan/interfacer) <span> ★ 471, pushed 2 days ago </span> - A linter that suggests interface types.
-   [validate](https://github.com/mccoyst/validate) <span> ★ 58, pushed 28 days ago </span> - Automatically validates struct fields with tags.

Editor Plugins
--------------

-   [go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) <span> ★ 3279, pushed 0 days ago </span> Go plugin for IntelliJ IDEA.
-   [go-plus](https://github.com/joefitzgerald/go-plus) <span> ★ 849, pushed 0 days ago </span> - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting
-   [Goclipse](https://github.com/GoClipse/goclipse) <span> ★ 521, pushed 5 days ago </span> - An Eclipse plugin for Go.
-   [gocode](https://github.com/nsf/gocode) <span> ★ 3054, pushed 40 days ago </span> - An autocompletion daemon for the Go programming language.
-   [GoSublime](https://github.com/DisposaBoy/GoSublime) <span> ★ 2203, pushed 3 days ago </span> - A Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features.
-   [velour](https://github.com/velour/velour) <span> ★ 13, pushed 53 days ago </span> - An IRC client for the acme editor.
-   [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) <span> ★ 65, pushed 657 days ago </span> - A Vim plugin to highlight syntax errors on save.
-   [vim-go](https://github.com/fatih/vim-go) <span> ★ 4657, pushed 0 days ago </span> - Go development plugin for Vim.
-   [Watch](https://github.com/eaburns/Watch) <span> ★ 105, pushed 170 days ago </span> - Runs a command in an acme win on file changes.

Go Tools
--------

-   [colorgo](https://github.com/songgao/colorgo) <span> ★ 63, pushed 301 days ago </span> - A wrapper around `    go   ` command for colorized `    go build   ` output.
-   [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.
-   [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) <span> ★ 29, pushed 144 days ago </span> - Bash completion for go and wgo.
-   [rts](https://github.com/galeone/rts) <span> ★ 97, pushed 6 days ago </span> - RTS: response to struct. Generates Go structs from server responses.

Software Packages
-----------------

Software written in Go.

### DevOps Tools

-   [aptly](https://github.com/smira/aptly) <span> ★ 1041, pushed 0 days ago </span> - aptly is a Debian repository management tool.
-   [awsenv](https://github.com/soniah/awsenv) <span> ★ 7, pushed 150 days ago </span> - a small binary that loads Amazon (AWS) environment variables for a profile.
-   [Banshee](https://github.com/eleme/banshee) <span> ★ 220, pushed 0 days ago </span> - Anomalies detection system for periodic metrics.
-   [Boom](https://github.com/rakyll/boom) <span> ★ 4182, pushed 3 days ago </span> - Boom is a tiny program that sends some load to a web application.
-   [bosun](https://github.com/bosun-monitor/bosun) <span> ★ 1770, pushed 3 days ago </span> - Time Series Alerting Framework.
-   [dogo](https://github.com/liudng/dogo) <span> ★ 90, pushed 214 days ago </span> - Monitoring changes in the source file and automatically compile and run (restart).
-   [Dropship](https://github.com/chrismckenzie/dropship) <span> ★ 22, pushed 7 days ago </span> - A tool for deploying code via cdn.
-   [EasySSH](https://github.com/hypersleep/easyssh) <span> ★ 93, pushed 184 days ago </span> - Golang package for easy remote execution through SSH and SCP downloading.
-   [Go Metrics](https://github.com/rcrowley/go-metrics) <span> ★ 1167, pushed 3 days ago </span> - Go port of Coda Hale's Metrics library: https://github.com/codahale/metrics.
-   [go-selfupdate](https://github.com/sanbornm/go-selfupdate) <span> ★ 362, pushed 191 days ago </span> - Enable your Go applications to self update.
-   [gobrew](https://github.com/cryptojuice/gobrew) <span> ★ 151, pushed 173 days ago </span> - gobrew lets you easily switch between multiple versions of go.
-   [godbg](https://github.com/sirnewton01/godbg) <span> ★ 185, pushed 86 days ago </span> - Web-based gdb front-end application.
-   [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
-   [gonative](https://github.com/inconshreveable/gonative) <span> ★ 232, pushed 112 days ago </span> - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
-   [gox](https://github.com/mitchellh/gox) <span> ★ 1546, pushed 26 days ago </span> - A dead simple, no frills Go cross compile tool.
-   [goxc](https://github.com/laher/goxc) <span> ★ 1268, pushed 26 days ago </span> - build tool for Go, with a focus on cross-compiling and packaging.
-   [GVM](https://github.com/moovweb/gvm) <span> ★ 1964, pushed 19 days ago </span> - GVM provides an interface to manage Go versions.
-   [hk](https://github.com/heroku/hk) <span> ★ 766, pushed 59 days ago </span> - Heroku command-line interface in Go.
-   [kala](https://github.com/ajvb/kala) <span> ★ 793, pushed 18 days ago </span> - Simplistic, modern, and performant job scheduler.
-   [kubernetes](https://github.com/kubernetes/kubernetes) <span> ★ 14058, pushed 0 days ago </span> - Container Cluster Manager from Google.
-   [Mora](https://github.com/emicklei/mora) <span> ★ 148, pushed 89 days ago </span> - REST server for accessing MongoDB documents and meta data.
-   [ostent](https://github.com/ostrost/ostent) <span> ★ 61, pushed 0 days ago </span> - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.
-   [Packer](https://github.com/mitchellh/packer) <span> ★ 5251, pushed 0 days ago </span> - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
-   [Rodent](https://github.com/alouche/rodent) <span> ★ 30, pushed 357 days ago </span> - Rodent helps you manage Go versions, projects and track dependencies.
-   [s3gof3r](https://github.com/rlmcpherson/s3gof3r) <span> ★ 662, pushed 75 days ago </span> - A small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.
-   [Scaleway-cli](https://github.com/scaleway/scaleway-cli) <span> ★ 124, pushed 0 days ago </span> - Manage BareMetal Servers from Command Line (as easily as with Docker).
-   [Vegeta](https://github.com/tsenart/vegeta) <span> ★ 3738, pushed 5 days ago </span> - HTTP load testing tool and library. It's over 9000!
-   [webhook](https://github.com/adnanh/webhook) <span> ★ 382, pushed 11 days ago </span> - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.
-   [Wide](https://wide.b3log.org/login) - A Web-based IDE for Teams using Golang.

### Other Software

-   [boxed](https://github.com/tejo/boxed) <span> ★ 46, pushed 2 days ago </span> - Dropbox based blog engine
-   [Cherry](https://github.com/rafael-santiago/cherry) <span> ★ 86, pushed 16 days ago </span> - A tiny webchat server in Go.
-   [Circuit](https://github.com/gocircuit/circuit) <span> ★ 1274, pushed 216 days ago </span> - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
-   [Comcast](https://github.com/tylertreat/Comcast) <span> ★ 4208, pushed 40 days ago </span> - Simulate bad network connections.
-   [confd](https://github.com/kelseyhightower/confd) <span> ★ 2592, pushed 1 days ago </span> - Manage local application configuration files using templates and data from etcd or consul.
-   [Docker](http://www.docker.com/) - An open platform for distributed applications for developers and sysadmins.
-   [fleet](https://github.com/coreos/fleet) <span> ★ 2162, pushed 0 days ago </span> - A Distributed init System.
-   [Go Package Store](https://github.com/shurcooL/Go-Package-Store#go-package-store-) - An app that displays updates for the Go packages in your GOPATH.
-   [gocc](https://github.com/goccmack/gocc) <span> ★ 35, pushed 11 days ago </span> - Gocc is a compiler kit for Go written in Go.
-   [GoDocTooltip](https://github.com/diankong/GoDocTooltip) <span> ★ 7, pushed 87 days ago </span> - A chrome extension for Go Doc sites, which shows function description as tooltip at funciton list.
-   [Gor](https://github.com/buger/gor) <span> ★ 5022, pushed 0 days ago </span> - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
-   [heka](https://github.com/mozilla-services/heka) <span> ★ 3029, pushed 0 days ago </span> - universal tool for data processing from Mozilla. Large collection of built-in plugins. Extendable via Go and Lua plugin API.
-   [hsync](http://ambrevar.bitbucket.org/hsync/) - A filesystem hierarchy synchronizer.
-   [hugo](http://gohugo.io/) - A Fast and Modern Static Website Engine.
-   [ipe](https://github.com/dimiro1/ipe) <span> ★ 149, pushed 34 days ago </span> - An open source Pusher server implementation compatible with Pusher client libraries written in GO.
-   [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
-   [limetext](http://limetext.org/) Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
-   [LiteIDE](https://github.com/visualfc/liteide) <span> ★ 2775, pushed 0 days ago </span> LiteIDE is a simple, open source, cross-platform Go IDE.
-   [mockingjay](https://github.com/quii/mockingjay-server) <span> ★ 234, pushed 1 days ago </span> Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.
-   [naclpipe](https://github.com/unix4fun/naclpipe) <span> ★ 5, pushed 132 days ago </span> - A simple NaCL EC25519 based crypto pipe tool written in Go.
-   [nes](https://github.com/fogleman/nes) <span> ★ 2255, pushed 164 days ago </span> - A Nintendo Entertainment System (NES) emulator written in Go.
-   [orange-cat](https://github.com/noraesae/orange-cat) <span> ★ 140, pushed 226 days ago </span> - A Markdown previewer written in Go.
-   [peg](https://github.com/pointlander/peg) <span> ★ 285, pushed 39 days ago </span> - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
-   [Postman](https://github.com/zachlatta/postman) <span> ★ 634, pushed 181 days ago </span> - Command-line utility for batch-sending email.
-   [restic](https://github.com/restic/restic) <span> ★ 595, pushed 0 days ago </span> - De-duplicating backup program.
-   [rkt](https://github.com/coreos/rkt) <span> ★ 5431, pushed 0 days ago </span> - An App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM.
-   [Seaweed File System](https://github.com/chrislusf/seaweedfs) <span> ★ 2397, pushed 8 days ago </span> - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
-   [shell2http](https://github.com/msoap/shell2http) <span> ★ 43, pushed 1 days ago </span> - Executing shell commands via http server (for prototyping or remote control).
-   [snap](https://github.com/intelsdi-x/snap) <span> ★ 674, pushed 0 days ago </span> - A powerful telemetry framework.
-   [Stack Up](https://github.com/pressly/sup) <span> ★ 917, pushed 11 days ago </span> - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
-   [syncthing](https://syncthing.net/) - An open, decentralized file synchronization tool and protocol.
-   [Tenyks](https://github.com/kyleterry/tenyks) <span> ★ 153, pushed 22 days ago </span> - Service oriented IRC bot using Redis and JSON for messaging.
-   [toto](https://github.com/blogcin/ToTo) <span> ★ 5, pushed 46 days ago </span> - A simple proxy server written in Go language, can be used together with browser.
-   [toxiproxy](https://github.com/shopify/toxiproxy) <span> ★ 1227, pushed 0 days ago </span> - Proxy to simulate network and system conditions for automated tests.
-   [tsuru](https://tsuru.io/) - An extensible and open source Platform as a Service software.
-   [websysd](https://github.com/ian-kent/websysd) - Web based process manager (like Marathon or Upstart).
-   [wellington](https://github.com/wellington/wellington) <span> ★ 180, pushed 2 days ago </span> - Sass project management tool, extends the language with sprite functions (like Compass).

Resources
=========

Where to discover new Go libraries.

Benchmarks
----------

-   [autobench](https://github.com/davecheney/autobench) <span> ★ 84, pushed 637 days ago </span> - Framework to compare the performance between different Go versions.
-   [go-benchmarks](https://github.com/tylertreat/go-benchmarks) <span> ★ 24, pushed 60 days ago </span> - A few miscellaneous Go microbenchmarks. Compare some language features to alternative aproaches.
-   [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) <span> ★ 696, pushed 12 days ago </span> - Go HTTP request router benchmark and comparison.
-   [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) <span> ★ 3, pushed 545 days ago </span> - Naive performance test of two ways to do type assertion in Go.
-   [go *serialization* benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) <span> ★ 295, pushed 1 days ago </span> - Benchmarks of Go serialization methods.
-   [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) <span> ★ 47, pushed 477 days ago </span> - Benchmarks of common basic operations for the Go language.
-   [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) <span> ★ 5, pushed 129 days ago </span> - Tiny collection of Go micro benchmarks. The intent is to compare some language features to others.
-   [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) <span> ★ 21, pushed 409 days ago </span> - A collection of benchmarks for popular Go database/SQL utilities.
-   [gospeed](https://github.com/feyeleanor/GoSpeed) <span> ★ 49, pushed 89 days ago </span> - Go micro-benchmarks for calculating the speed of language constructs.
-   [kvbench](https://github.com/jimrobinson/kvbench) <span> ★ 11, pushed 736 days ago </span> - Key/Value database benchmark.
-   [skynet](https://github.com/atemerev/skynet) <span> ★ 576, pushed 1 days ago </span> - Skynet 1M threads microbenchmark.
-   [speedtest-resize](https://github.com/fawick/speedtest-resize) <span> ★ 69, pushed 226 days ago </span> - Compare various Image resize algorithms for the Go language.

Conferences
-----------

-   [dotGo](http://www.dotgo.eu) - Paris, France
-   [GoCon](http://gocon.connpass.com/) - Tokyo, Japan
-   [GolangUK](http://golanguk.com/) - London, UK
-   [GopherChina](http://gopherchina.org) - Shanghai, China
-   [GopherCon](http://www.gophercon.com/) - Denver, USA
-   [GopherCon Dubai](http://www.gophercon.ae/) - Dubai, UAE
-   [GopherCon India](http://www.gophercon.in/) - Bengaluru, India
-   [GothamGo](http://gothamgo.com/) - New York City, USA

E-Books
-------

-   [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
-   [An Introduction to Programming in Go](http://www.golang-book.com/)
-   [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
-   [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
-   [Go Bootcamp](http://golangbootcamp.com)
-   [GoBooks](https://github.com/dariubs/GoBooks) <span> ★ 2090, pushed 31 days ago </span> - A curated list of Go books
-   [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
-   [Network Programming With Go](https://jan.newmarch.name/go/)
-   [The Go Programming Language](http://www.gopl.io/)

Twitter
-------

-   [@golang](https://twitter.com/golang)
-   [@golang\_news](https://twitter.com/golang_news)
-   [@golangweekly](https://twitter.com/golangweekly)

Websites
--------

-   [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) <span> ★ 5387, pushed 0 days ago </span> - A curated list of awesome remote jobs. A lot of them is looking for Go hackers.
-   [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) <span> ★ 16051, pushed 1 days ago </span> - List of other amazingly awesome lists.
-   [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - A collection of Go articles and tutorials.
-   [Go Blog](http://blog.golang.org) - The official Go blog.
-   [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
-   [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
-   [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
-   [golang-graphics](https://github.com/mholt/golang-graphics) <span> ★ 76, pushed 245 days ago </span> - A collection of Go images, graphics, and art.
-   [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
-   [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for \#golang enthusiasts.
-   [gowalker.org](https://gowalker.org) - Go Project API documentation.
-   [r/Golang](https://www.reddit.com/r/golang) - News about Go.
-   [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.

### Tutorials

-   [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
-   [Go By Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated example programs.
-   [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
-   [Working with Go](https://github.com/mkaz/working-with-go) <span> ★ 502, pushed 200 days ago </span> - An intro to go for experienced programmers.

Windows
-------

-   [go-ole](https://github.com/go-ole/go-ole) <span> ★ 191, pushed 45 days ago </span> - Win32 OLE implementation for golang.

