# Awesome Go 

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [awesome-go chat channel](http://gophers.slack.com/messages/awesome)

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Files](#files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot-internet-of-things)
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
        - [Routers](#routers)
    - [Windows](#windows)
    - [XML](#xml)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
	- [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Server Applications](#server-applications)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [flac](https://github.com/eaburns/flac) (:star: 71) - Native Go FLAC decoder.
* [flac](https://github.com/mewkiz/flac) (:star: 54) - Native Go FLAC decoder.
* [gaad](https://github.com/Comcast/gaad) (:star: 35) - Native Go AAC bitstream parser.
* [go-sox](https://github.com/krig/go-sox) (:star: 56) - libsox bindings for go.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) (:star: 16) - libmediainfo bindings for go.
* [gosamplerate](https://github.com/dh1tw/gosamplerate) (:star: 3) - libsamplerate bindings for go.
* [id3v2](https://github.com/bogem/id3v2) (:star: 51) - Fast and stable ID3 parsing and writing library for Go.
* [malgo](https://github.com/gen2brain/malgo) (:star: 12) - Mini audio library.
* [mix](https://github.com/go-mix/mix) (:star: 67) - Sequence-based Go-native audio mixer for music apps.
* [mp3](https://github.com/tcolgate/mp3) (:star: 64) - Native Go MP3 decoder.
* [music-theory](https://github.com/go-music-theory/music-theory) (:star: 180) - Music theory models in Go.
* [PortAudio](https://github.com/gordonklaus/portaudio) (:star: 165) - Go bindings for the PortAudio audio I/O library.
* [portmidi](https://github.com/rakyll/portmidi) (:star: 150) - Go bindings for PortMidi.
* [taglib](https://github.com/wtolson/go-taglib) (:star: 53) - Go bindings for taglib.
* [vorbis](https://github.com/mccoyst/vorbis) (:star: 19) - "Native" Go Vorbis decoder (uses CGO, but has no dependencies).
* [waveform](https://github.com/mdlayher/waveform) (:star: 198) - Go package capable of generating waveform images from audio streams.

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [authboss](https://github.com/volatiletech/authboss) (:star: 1255) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.
* [casbin](https://github.com/hsluoyz/casbin) (:star: 1554) - Authorization library that supports access control models like ACL, RBAC, ABAC.
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) (:star: 0) - provides parser of cookies.txt file format.
* [Go-AWS-Auth](https://github.com/smartystreets/go-aws-auth) (:star: 179) - AWS (Amazon Web Services) request signing library.
* [go-jose](https://github.com/square/go-jose) (:star: 765) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) (:star: 703) - Standalone, specification-compliant,  OAuth2 server written in Golang.
* [gologin](https://github.com/dghubble/gologin) (:star: 785) - chainable handlers for login with OAuth1 and OAuth2 authentication providers.
* [gorbac](https://github.com/mikespook/gorbac) (:star: 626) - provides a lightweight role-based access control (RBAC) implementation in Golang.
* [goth](https://github.com/markbates/goth) (:star: 1443) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.
* [httpauth](https://github.com/goji/httpauth) (:star: 127) - HTTP Authentication middleware.
* [jwt](https://github.com/robbert229/jwt) (:star: 26) - Clean and easy to use implementation of JSON Web Tokens (JWT).
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) (:star: 87) - JWT middleware for Golang http servers with many configuration options.
* [jwt-go](https://github.com/dgrijalva/jwt-go) (:star: 2798) - Golang implementation of JSON Web Tokens (JWT).
* [loginsrv](https://github.com/tarent/loginsrv) (:star: 511) - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.
* [oauth2](https://github.com/golang/oauth2) (:star: 1342) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.
* [osin](https://github.com/RangelReale/osin) (:star: 1278) - Golang OAuth2 server library.
* [permissions2](https://github.com/xyproto/permissions2) (:star: 245) - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.
* [securecookie](https://github.com/chmike/securecookie) (:star: 9) - Efficient secure cookie encoding/decoding.
* [session](https://github.com/icza/session) (:star: 46) - Go session management for web servers (including support for Google App Engine - GAE).
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) (:star: 3) - Go session management using the SessionGate Redis module.
* [sessions](https://github.com/adam-hanna/sessions) (:star: 25) - Dead simple, highly performant, highly customizable sessions service for go http servers.
* [traefik](https://github.com/containous/traefik) (:star: 11313) - Reverse proxy and load balancer with support for multiple backends.
* [yubigo](https://github.com/GeertJohan/yubigo) (:star: 81) - Yubikey client package that provides a simple API to integrate the Yubico Yubikey into a go application.

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [argparse](https://github.com/akamensky/argparse) (:star: 1) - Command line argument parser inspired by Python's argparse module.
* [argv](https://github.com/cosiner/argv) (:star: 3) - Go library to split command line string as arguments array using the bash syntax.
* [cli](https://github.com/mkideal/cli) (:star: 318) - Feature-rich and easy to use command-line package based on golang struct tags.
* [cli](https://github.com/teris-io/cli) (:star: 18) - Simple and complete API for building command line interfaces in Go.
* [cli-init](https://github.com/tcnksm/gcli) (:star: 764) - The easy way to start building Golang command line applications.
* [climax](http://github.com/tucnak/climax) (:star: 126) - Alternative CLI with "human face", in spirit of Go command.
* [cobra](https://github.com/spf13/cobra) (:star: 5709) - Commander for modern Go CLI interactions.
* [complete](https://github.com/posener/complete) (:star: 388) - Write bash completions in Go + Go command bash completion.
* [docopt.go](https://github.com/docopt/docopt.go) (:star: 872) - Command-line arguments parser that will make you smile.
* [drive](https://github.com/odeke-em/drive) (:star: 3696) - Google Drive client for the commandline.
* [env](https://github.com/codingconcepts/env) (:star: 8) - Tag-based environment configuration for structs.
* [flag](https://github.com/cosiner/flag) (:star: 72) - Simple but powerful command line option parsing library for Go supporting subcommand.
* [go-arg](https://github.com/alexflint/go-arg) (:star: 446) - Struct-based argument parsing in Go.
* [go-flags](https://github.com/jessevdk/go-flags) (:star: 918) - go command line option parser.
* [kingpin](https://github.com/alecthomas/kingpin) (:star: 1586) - Command line and flag parser supporting sub commands.
* [liner](https://github.com/peterh/liner) (:star: 406) - Go readline-like library for command-line interfaces.
* [mitchellh/cli](https://github.com/mitchellh/cli) (:star: 701) - Go library for implementing command-line interfaces.
* [mow.cli](https://github.com/jawher/mow.cli) (:star: 468) - Go library for building CLI applications with sophisticated flag and argument parsing and validation.
* [pflag](https://github.com/spf13/pflag) (:star: 275) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.
* [readline](https://github.com/chzyer/readline) (:star: 1019) - Pure golang implementation that provides most features in GNU-Readline under MIT license.
* [sflags](https://github.com/octago/sflags) (:star: 44) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.
* [ukautz/clif](https://github.com/ukautz/clif) (:star: 79) - Small command line interface framework.
* [urfave/cli](https://github.com/urfave/cli) (:star: 6909) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).
* [wlog](https://github.com/dixonwille/wlog) (:star: 22) - Simple logging interface that supports cross-platform color and concurrency.
* [wmenu](https://github.com/dixonwille/wmenu) (:star: 36) - Easy to use menu structure for cli applications that prompts users to make choices.

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [aurora](https://github.com/logrusorgru/aurora) (:star: 257) - ANSI terminal colors that supports fmt.Printf/Sprintf.
* [chalk](https://github.com/ttacon/chalk) (:star: 220) - Intuitive package for prettifying terminal/console output.
* [color](https://github.com/fatih/color) (:star: 1969) - Versatile package for colored terminal output.
* [colourize](https://github.com/TreyBastian/colourize) (:star: 12) - Go library for ANSI colour text in terminals.
* [go-ataman](https://github.com/workanator/go-ataman) (:star: 5) - Go library for rendering ANSI colored text templates in terminals.
* [go-colorable](https://github.com/mattn/go-colorable) (:star: 215) - Colorable writer for windows.
* [go-colortext](https://github.com/daviddengcn/go-colortext) (:star: 162) - Go library for color output in terminals.
* [go-isatty](https://github.com/mattn/go-isatty) (:star: 189) - isatty for golang.
* [gocui](https://github.com/jroimartin/gocui) (:star: 2053) - Minimalist Go library aimed at creating Console User Interfaces.
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
* [mpb](https://github.com/vbauerster/mpb) (:star: 127) - Multi progress bar for terminal applications.
* [progressbar](https://github.com/schollz/progressbar) (:star: 213) - Basic thread-safe progress bar that works in every OS.
* [termbox-go](https://github.com/nsf/termbox-go) (:star: 2512) - Termbox is a library for creating cross-platform text-based interfaces.
* [termtables](https://github.com/apcera/termtables) (:star: 111) - Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.
* [termui](https://github.com/gizak/termui) (:star: 6651) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).
* [tui-go](https://github.com/marcusolsson/tui-go) (:star: 1073) - Go UI library for building rich terminal applications.
* [uilive](https://github.com/gosuri/uilive) (:star: 597) - Library for updating terminal output in realtime.
* [uiprogress](https://github.com/gosuri/uiprogress) (:star: 1035) - Flexible library to render progress bars in terminal applications.
* [uitable](https://github.com/gosuri/uitable) (:star: 401) - Library to improve readability in terminal apps using tabular data.

## Configuration

*Libraries for configuration parsing.*

* [config](https://github.com/olebedev/config) (:star: 138) - JSON or YAML configuration wrapper with environment variables and flags parsing.
* [configure](https://github.com/paked/configure) (:star: 29) - Provides configuration through multiple sources, including JSON, flags and environment variables.
* [env](https://github.com/caarlos0/env) (:star: 353) - Parse environment variables to Go structs (with defaults).
* [envcfg](https://github.com/tomazk/envcfg) (:star: 83) - Un-marshaling environment variables to Go structs.
* [envconf](https://github.com/ian-kent/envconf) (:star: 5) - Configuration from environment.
* [envconfig](https://github.com/vrischmann/envconfig) (:star: 110) - Read your configuration from environment variables.
* [envh](https://github.com/antham/envh) (:star: 83) - Helpers to manage environment variables.
* [gcfg](https://github.com/go-gcfg/gcfg) (:star: 73) - read INI-style configuration files into Go structs; supports user-defined types and subsections.
* [goConfig](https://github.com/crgimenes/goConfig) (:star: 49) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.
* [godotenv](https://github.com/joho/godotenv) (:star: 845) - Go port of Ruby's dotenv library (Loads environment variables from `.env`).
* [gofigure](https://github.com/ian-kent/gofigure) (:star: 48) - Go application configuration made easy.
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [hjson](https://github.com/hjson/hjson-go) (:star: 104) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.
* [ingo](https://github.com/schachmat/ingo) (:star: 15) - Flags persisted in an ini-like config file.
* [ini](https://github.com/go-ini/ini) (:star: 745) - Go package to read and write INI files.
* [joshbetz/config](https://github.com/joshbetz/config) (:star: 149) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.
* [mini](https://github.com/sasbury/mini) (:star: 15) - Golang package for parsing ini-style configuration files.
* [store](https://github.com/tucnak/store) (:star: 213) - Lightweight configuration manager for Go.
* [viper](https://github.com/spf13/viper) (:star: 4065) - Go configuration with fangs.
* [xdg](https://github.com/OpenPeeDeeP/xdg) (:star: 1) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) (:star: 12142) - Drone is a Continuous Integration platform built on Docker, written in Go.
* [goveralls](https://github.com/mattn/goveralls) (:star: 431) - Go integration for Coveralls.io continuous code coverage tracking system.
* [overalls](https://github.com/go-playground/overalls) (:star: 73) - Multi-Package go project coverprofile for tools like goveralls.
* [roveralls](https://github.com/LawrenceWoodman/roveralls) (:star: 1) - Recursive coverage testing tool.

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [c6](https://github.com/c9s/c6) (:star: 389) - High performance SASS compatible-implementation compiler written in Go.
* [gcss](https://github.com/yosssi/gcss) (:star: 394) - Pure Go CSS Preprocessor.
* [go-libsass](https://github.com/wellington/go-libsass) (:star: 67) - Go wrapper to the 100% Sass compatible libsass project.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [binpacker](https://github.com/zhuangsirui/binpacker) (:star: 61) - Binary packer and unpacker helps user build custom binary stream.
* [bit](https://github.com/yourbasic/bit) (:star: 6) - Golang set data structure with bonus bit-twiddling functions.
* [bitset](https://github.com/willf/bitset) (:star: 316) - Go package implementing bitsets.
* [bloom](https://github.com/zhenjl/bloom) (:star: 106) - Bloom filters implemented in Go.
* [bloom](https://github.com/yourbasic/bloom) (:star: 2) - Golang Bloom filter implementation.
* [boomfilters](https://github.com/tylertreat/BoomFilters) (:star: 870) - Probabilistic data structures for processing continuous, unbounded streams.
* [concurrent-writer](https://github.com/free/concurrent-writer) (:star: 5) - Highly concurrent drop-in replacement for `bufio.Writer`.
* [count-min-log](https://github.com/seiflotfy/count-min-log) (:star: 35) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) (:star: 254) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
* [encoding](https://github.com/zhenjl/encoding) (:star: 79) - Integer Compression Libraries for Go.
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) (:star: 25) - Go implementation of Adaptive Radix Tree.
* [go-datastructures](https://github.com/Workiva/go-datastructures) (:star: 3670) - Collection of useful, performant, and thread-safe data structures.
* [go-ef](https://github.com/amallia/go-ef) (:star: 5) - A Go implementation of the Elias-Fano encoding.
* [go-geoindex](https://github.com/hailocab/go-geoindex) (:star: 263) - In-memory geo index.
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) (:star: 72) - Region quadtrees with efficient point location and neighbour finding.
* [gods](https://github.com/emirpasic/gods) (:star: 3362) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.
* [golang-set](https://github.com/deckarep/golang-set) (:star: 635) - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
* [goset](https://github.com/zoumo/goset) (:star: 6) - A useful Set collection implementation for Go.
* [goskiplist](https://github.com/ryszard/goskiplist) (:star: 140) - Skip list implementation in Go.
* [gota](https://github.com/kniren/gota) (:star: 396) - Implementation of dataframes, series, and data wrangling methods for Go.
* [hilbert](https://github.com/google/hilbert) (:star: 109) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.
* [hyperloglog](https://github.com/axiomhq/hyperloglog) (:star: 555) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.
* [levenshtein](https://github.com/agext/levenshtein) (:star: 6) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [levenshtein](https://github.com/agnivade/levenshtein) (:star: 11) - Implementation to calculate levenshtein distance in Go.
* [mafsa](https://github.com/smartystreets/mafsa) (:star: 255) - MA-FSA implementation with Minimal Perfect Hashing.
* [merkletree](https://github.com/cbergoon/merkletree) (:star: 17) - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.
* [roaring](https://github.com/RoaringBitmap/roaring) (:star: 308) - Go package implementing compressed bitsets.
* [skiplist](https://github.com/gansidui/skiplist) (:star: 41) - Skiplist implementation in Go.
* [trie](https://github.com/derekparker/trie) (:star: 213) - Trie implementation in Go.
* [ttlcache](https://github.com/diegobernardes/ttlcache) (:star: 45) - In-memory LRU string-interface{} map with expiration for golang.
* [willf/bloom](https://github.com/willf/bloom) (:star: 375) - Go package implementing Bloom filters.

## Database

*Databases implemented in Go.*

* [badger](https://github.com/dgraph-io/badger) (:star: 2727) - Fast key-value store in Go.
* [BigCache](https://github.com/allegro/bigcache) (:star: 1273) - Efficient key/value cache for gigabytes of data.
* [bolt](https://github.com/boltdb/bolt) (:star: 7379) - Low-level key/value database for Go.
* [buntdb](https://github.com/tidwall/buntdb) (:star: 1634) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.
* [cache2go](https://github.com/muesli/cache2go) (:star: 353) - In-memory key:value cache which supports automatic invalidation based on timeouts.
* [cockroach](https://github.com/cockroachdb/cockroach) (:star: 11793) - Scalable, Geo-Replicated, Transactional Datastore.
* [couchcache](https://github.com/codingsince1985/couchcache) (:star: 27) - RESTful caching micro-service backed by Couchbase server.
* [dgraph](https://github.com/dgraph-io/dgraph) (:star: 4169) - Scalable, Distributed, Low Latency, High Throughput Graph Database.
* [diskv](https://github.com/peterbourgon/diskv) (:star: 497) - Home-grown disk-backed key-value store.
* [eliasdb](https://github.com/krotik/eliasdb) (:star: 444) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.
* [forestdb](https://github.com/couchbase/goforestdb) (:star: 28) - Go bindings for ForestDB.
* [GCache](https://github.com/bluele/gcache) (:star: 283) - Cache library with support for expirable Cache, LFU, LRU and ARC.
* [geocache](https://github.com/melihmucuk/geocache) (:star: 72) - In-memory cache that is suitable for geolocation based applications.
* [go-cache](https://github.com/pmylund/go-cache) (:star: 1226) - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
* [goleveldb](https://github.com/syndtr/goleveldb) (:star: 1856) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.
* [groupcache](https://github.com/golang/groupcache) (:star: 5945) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
* [influxdb](https://github.com/influxdb/influxdb) (:star: 12005) - Scalable datastore for metrics, events, and real-time analytics.
* [ledisdb](https://github.com/siddontang/ledisdb) (:star: 2316) - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
* [levigo](https://github.com/jmhodges/levigo) (:star: 316) - Levigo is a Go wrapper for LevelDB.
* [moss](https://github.com/couchbase/moss) (:star: 391) - Moss is a simple LSM key-value storage engine written in 100% Go.
* [piladb](https://github.com/fern4lvarez/piladb) (:star: 142) - Lightweight RESTful database engine based on stack data structures.
* [prometheus](https://github.com/prometheus/prometheus) (:star: 13118) - Monitoring system and time series database.
* [rqlite](https://github.com/rqlite/rqlite) (:star: 3355) - The lightweight, distributed, relational database built on SQLite.
* [Scribble](https://github.com/nanobox-io/golang-scribble) (:star: 90) - Tiny flat file JSON store.
* [tempdb](https://github.com/rafaeljesus/tempdb) (:star: 6) - Key-value store for temporary items.
* [tidb](https://github.com/pingcap/tidb) (:star: 10785) - TiDB is a distributed SQL database. Inspired by the design of Google F1.
* [tiedot](https://github.com/HouzuoGuo/tiedot) (:star: 1963) - Your NoSQL database powered by Golang.
* [Tile38](https://github.com/tidwall/tile38) (:star: 3598) - Geolocation DB with spatial index and realtime geofencing.

*Database schema migration.*

* [darwin](https://github.com/GuiaBolso/darwin) (:star: 50) - Database schema evolution library for Go.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) (:star: 4) - Django style fixtures for Golang's excellent built-in database/sql library.
* [gondolier](https://github.com/emvicom/gondolier) (:star: 8) - Gondolier is a library to auto migrate database schemas using structs.
* [goose](https://github.com/steinbacher/goose) (:star: 54) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
* [gormigrate](https://github.com/go-gormigrate/gormigrate) (:star: 71) - Database schema migration helper for Gorm ORM.
* [migrate](https://github.com/mattes/migrate) (:star: 1773) - Database migrations. CLI and Golang library.
* [pravasan](https://github.com/pravasan/pravasan) (:star: 18) - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [soda](https://github.com/markbates/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [sql-migrate](https://github.com/rubenv/sql-migrate) (:star: 798) - Database migration tool. Allows embedding migrations into the application using go-bindata.

*Database tools.*

* [chproxy](https://github.com/Vertamedia/chproxy) (:star: 54) - HTTP proxy for ClickHouse database.
* [go-mysql](https://github.com/siddontang/go-mysql) (:star: 747) - Go toolset to handle MySQL protocol and replication.
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) (:star: 1015) - Sync your MySQL data into Elasticsearch automatically.
* [kingshard](https://github.com/flike/kingshard) (:star: 3088) - kingshard is a high performance proxy for MySQL powered by Golang.
* [myreplication](https://github.com/2tvenom/myreplication) (:star: 85) - MySql binary log replication listener. Supports statement and row based replication.
* [orchestrator](https://github.com/github/orchestrator) (:star: 947) - MySQL replication topology manager & visualizer.
* [pgweb](https://github.com/sosedoff/pgweb) (:star: 4675) - Web-based PostgreSQL database browser.
* [pREST](https://github.com/nuveo/prest) (:star: 1530) - Serve a RESTful API from any PostgreSQL database.
* [rwdb](https://github.com/andizzle/rwdb) (:star: 8) - rwdb provides read replica capability for multiple database servers setup.
* [vitess](https://github.com/youtube/vitess) (:star: 5224) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.

*SQL query builder, libraries for building and using SQL.*

* [dat](https://github.com/mgutz/dat) (:star: 478) - Go Postgres Data Access Toolkit.
* [Dotsql](https://github.com/gchaincl/dotsql) (:star: 318) - Go library that helps you keep sql files in one place and use them with ease.
* [goqu](https://github.com/doug-martin/goqu) (:star: 356) - Idiomatic SQL builder and query library.
* [igor](https://github.com/galeone/igor) (:star: 62) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) (:star: 187) - Powerful data retrieval methods as well as DB-agnostic query building capabilities.
* [scaneo](https://github.com/variadico/scaneo) (:star: 117) - Generate Go code to convert database rows into arbitrary structs.
* [sqrl](https://github.com/elgris/sqrl) (:star: 89) - SQL query builder, fork of Squirrel with improved performance.
* [Squirrel](https://github.com/Masterminds/squirrel) (:star: 1291) - Go library that helps you build SQL queries.
* [xo](https://github.com/knq/xo) (:star: 1385) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [avatica](https://github.com/Boostport/avatica) (:star: 36) - Apache Phoenix/Avatica SQL driver for database/sql.
    * [bgc](https://github.com/viant/bgc) (:star: 5) - Datastore Connectivity for BigQuery for go.
    * [firebirdsql](https://github.com/nakagami/firebirdsql) (:star: 68) - Firebird RDBMS SQL driver for Go.
    * [go-adodb](https://github.com/mattn/go-adodb) (:star: 62) - Microsoft ActiveX Object DataBase driver for go that uses database/sql.
    * [go-bqstreamer](https://github.com/rounds/go-bqstreamer) (:star: 130) - BigQuery fast and concurrent stream insert.
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) (:star: 636) - Microsoft MSSQL driver for Go.
    * [go-oci8](https://github.com/mattn/go-oci8) (:star: 252) - Oracle driver for go that uses database/sql.
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) (:star: 4466) - MySQL driver for Go.
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) (:star: 2128) - SQLite3 driver for go that uses database/sql.
    * [gofreetds](https://github.com/minus5/gofreetds) (:star: 54) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [pgx](https://github.com/jackc/pgx) (:star: 1168) - PostgreSQL driver supporting features beyond those exposed by database/sql.
    * [pq](https://github.com/lib/pq) (:star: 3322) - Pure Go Postgres driver for database/sql.

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) (:star: 235) - Aerospike client in Go language.
    * [arangolite](https://github.com/solher/arangolite) (:star: 53) - Lightweight golang driver for ArangoDB.
    * [asc](https://github.com/viant/asc) (:star: 2) - Datastore Connectivity for Aerospike for go.
    * [cayley](https://github.com/google/cayley) (:star: 10697) - Graph database with support for multiple backends.
    * [dsc](https://github.com/viant/dsc) (:star: 4) - Datastore connectivity for SQL, NoSQL, structured files.
    * [dynago](https://github.com/underarmour/dynago) (:star: 48) - Dynago is a principle of least surprise client for DynamoDB.
    * [go-couchbase](https://github.com/couchbase/go-couchbase) (:star: 253) - Couchbase client in Go.
    * [go-couchdb](https://github.com/fjl/go-couchdb) (:star: 44) - Yet another CouchDB HTTP API wrapper for Go.
    * [gocb](https://github.com/couchbase/gocb) (:star: 233) - Official Couchbase Go SDK.
    * [gocql](http://gocql.github.io) - Go language driver for Apache Cassandra.
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
    * [gorethink](https://github.com/dancannon/gorethink) (:star: 1249) - Go language driver for RethinkDB.
    * [goriak](https://github.com/zegl/goriak) (:star: 19) - Go language driver for Riak KV.
    * [mgo](https://godoc.org/labix.org/v2/mgo) - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    * [neo4j](https://github.com/cihangir/neo4j) (:star: 22) - Neo4j Rest API Bindings for Golang.
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) (:star: 66) - Neo4j REST Client in golang.
    * [neoism](https://github.com/jmcvetta/neoism) (:star: 307) - Neo4j client for Golang.
    * [redigo](https://github.com/garyburd/redigo) (:star: 3744) - Redigo is a Go client for the Redis database.
    * [redis](https://github.com/go-redis/redis) (:star: 2534) - Redis client for Golang.
    * [redis](https://github.com/hoisie/redis) (:star: 548) - Simple, powerful Redis client for Go.
    * [redis](https://github.com/bsm/redeo) (:star: 147) - Redis-protocol compatible TCP servers/services.
    * [xredis](https://github.com/shomali11/xredis) (:star: 7) - Typesafe, customizable, clean & easy to use Redis client.

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) (:star: 3646) - Modern text indexing library for go.
    * [elastic](https://github.com/olivere/elastic) (:star: 2016) - Elasticsearch client for Go.
    * [elasticsql](https://github.com/cch123/elasticsql) (:star: 120) - Convert sql to elasticsearch dsl in Go.
    * [elastigo](https://github.com/mattbaird/elastigo) (:star: 888) - Elasticsearch client library.
    * [goes](https://github.com/belogik/goes) (:star: 77) - Library to interact with Elasticsearch.
    * [riot](https://github.com/go-ego/riot) (:star: 3120) - Go Open Source, Distributed, Simple and efficient Search Engine
    * [skizze](https://github.com/seiflotfy/skizze) (:star: 46) - probabilistic data-structures service and storage.

## Date and Time

*Libraries for working with dates and times.*

* [carbon](https://github.com/uniplaces/carbon) (:star: 145) - Simple Time extension with a lot of util methods, ported from PHP Carbon library.
* [date](https://github.com/rickb777/date) (:star: 5) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
* [dateparse](https://github.com/araddon/dateparse) (:star: 66) - Parse date's without knowing format in advance.
* [durafmt](https://github.com/hako/durafmt) (:star: 167) - Time duration formatting library for Go.
* [feiertage](https://github.com/wlbr/feiertage) (:star: 8) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) (:star: 28) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) (:star: 1) - Calculate the sunrise and sunset times for a given location.
* [goweek](https://github.com/grsmv/goweek) (:star: 12) - Library for working with week entity in golang.
* [now](https://github.com/jinzhu/now) (:star: 1025) - Now is a time toolkit for golang.
* [NullTime](https://github.com/kirillDanshin/nulltime) (:star: 3) - Nullable `time.Time`.
* [timeutil](https://github.com/leekchan/timeutil) (:star: 140) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.
* [tuesday](https://github.com/osteele/tuesday) (:star: 3) - Ruby-compatible Strftime function.

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [celeriac](https://github.com/svcavallar/celeriac.v1) (:star: 27) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
* [digota](https://github.com/digota/digota) (:star: 160) - grpc ecommerce microservice.
* [drmaa](https://github.com/dgruber/drmaa) (:star: 19) - Job submission library for cluster schedulers based on the DRMAA standard.
* [emitter-io](https://github.com/emitter-io/emitter) (:star: 169) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.
* [flowgraph](https://github.com/vectaport/flowgraph) (:star: 33) - MPI-style ready-send coordination layer.
* [gleam](https://github.com/chrislusf/gleam) (:star: 992) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.
* [glow](https://github.com/chrislusf/glow) (:star: 1817) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.
* [go-jump](https://github.com/dgryski/go-jump) (:star: 135) - Port of Google's "Jump" Consistent Hash function.
* [go-kit](https://github.com/go-kit/kit) (:star: 8455) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.
* [gorpc](https://github.com/valyala/gorpc) (:star: 412) - Simple, fast and scalable RPC library for high load.
* [grpc-go](https://github.com/grpc/grpc-go) (:star: 4354) - The Go language implementation of gRPC. HTTP/2 based RPC.
* [hprose](https://github.com/hprose/hprose-golang) (:star: 661) - Very newbility RPC Library, support 25+ languages now.
* [jsonrpc](https://github.com/osamingo/jsonrpc) (:star: 40) - The jsonrpc package helps implement of JSON-RPC 2.0.
* [jsonrpc](https://github.com/ybbus/jsonrpc) (:star: 15) - JSON-RPC 2.0 HTTP client implementation.
* [KrakenD](https://github.com/devopsfaith/krakend) (:star: 174) - Ultra performant API Gateway framework with middlewares.
* [micro](https://github.com/micro/micro) (:star: 3767) - Pluggable microservice toolkit and distributed systems platform.
* [NATS](https://github.com/nats-io/gnatsd) (:star: 3559) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.
* [raft](https://github.com/hashicorp/raft) (:star: 1448) - Golang implementation of the Raft consensus protocol, by HashiCorp.
* [raft](https://github.com/coreos/etcd/tree/master/raft) - Go implementation of the Raft consensus protocol, by CoreOS.
* [ringpop-go](https://github.com/uber/ringpop-go) (:star: 371) - Scalable, fault-tolerant application-layer sharding for Go applications.
* [rpcx](https://github.com/smallnest/rpcx) (:star: 1697) - Distributed pluggable RPC service framework like alibaba Dubbo.
* [sleuth](https://github.com/ursiform/sleuth) (:star: 232) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).
* [tendermint](https://github.com/tendermint/tendermint) (:star: 943) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.
* [torrent](https://github.com/anacrolix/torrent) (:star: 1892) - BitTorrent client package.
    * [dht](https://godoc.org/github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation.
    * [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) (:star: 293) - Video streaming torrent client.

## Email

*Libraries that implement email creation and sending.*

* [douceur](https://github.com/aymerick/douceur) (:star: 106) - CSS inliner for your HTML emails.
* [email](https://github.com/jordan-wright/email) (:star: 820) - A robust and flexible email library for Go.
* [go-dkim](https://github.com/toorop/go-dkim) (:star: 32) - DKIM library, to sign & verify email.
* [go-imap](https://github.com/emersion/go-imap) (:star: 316) - IMAP library for clients and servers.
* [go-message](https://github.com/emersion/go-message) (:star: 30) - Streaming library for the Internet Message Format and mail messages.
* [Gomail](https://github.com/go-gomail/gomail/) - Gomail is a very simple and powerful package to send emails.
* [Hectane](https://github.com/hectane/hectane) (:star: 131) - Lightweight SMTP client providing an HTTP API.
* [hermes](https://github.com/matcornic/hermes) (:star: 1057) - Golang package that generates clean, responsive HTML e-mails.
* [MailHog](https://github.com/mailhog/MailHog) (:star: 2750) - Email and SMTP testing with web and API interface.
* [SendGrid](https://github.com/sendgrid/sendgrid-go) (:star: 301) - SendGrid's Go library for sending email.
* [smtp](https://github.com/mailhog/smtp) (:star: 36) - SMTP server protocol state machine.

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [agora](https://github.com/PuerkitoBio/agora) (:star: 284) - Dynamically typed, embeddable programming language in Go.
* [anko](https://github.com/mattn/anko) (:star: 471) - Scriptable interpreter written in Go.
* [binder](https://github.com/alexeyco/binder) (:star: 6) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).
* [gisp](https://github.com/jcla1/gisp) (:star: 380) - Simple LISP in Go.
* [go-duktape](https://github.com/olebedev/go-duktape) (:star: 524) - Duktape JavaScript engine bindings for Go.
* [go-lua](https://github.com/Shopify/go-lua) (:star: 1112) - Port of the Lua 5.2 VM to pure Go.
* [go-php](https://github.com/deuill/go-php) (:star: 384) - PHP bindings for Go.
* [go-python](https://github.com/sbinet/go-python) (:star: 565) - naive go bindings to the CPython C-API.
* [golua](https://github.com/aarzilli/golua) (:star: 365) - Go bindings for Lua C API.
* [gopher-lua](https://github.com/yuin/gopher-lua) (:star: 1881) - Lua 5.1 VM and compiler written in Go.
* [ngaro](https://github.com/db47h/ngaro) (:star: 8) - Embeddable Ngaro VM implementation enabling scripting in Retro.
* [otto](https://github.com/robertkrimen/otto) (:star: 3328) - JavaScript interpreter written in Go.
* [purl](https://github.com/ian-kent/purl) (:star: 18) - Perl 5.18.2 embedded in Go.

## Files

*Libraries for  handling files and file systems.*

* [afero](https://github.com/spf13/afero) (:star: 1207) - FileSystem Abstraction System for Go.
* [go-csv-tag](https://github.com/artonge/go-csv-tag) (:star: 14) - Load csv file using tag.
* [go-gtfs](https://github.com/artonge/go-gtfs) (:star: 5) - Load gtfs files in go.
* [notify](https://github.com/rjeczalik/notify) (:star: 277) - File system event notification library with simple API, similar to os/signal.
* [skywalker](https://github.com/dixonwille/skywalker) (:star: 14) - Package to allow one to concurrently go through a filesystem with ease.
* [tarfs](https://github.com/posener/tarfs) (:star: 14) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.

## Financial

*Packages for accounting and finance.*

* [accounting](https://github.com/leekchan/accounting) (:star: 325) - money and currency formatting for golang.
* [decimal](https://github.com/shopspring/decimal) (:star: 686) - Arbitrary-precision fixed-point decimal numbers.
* [go-finance](https://github.com/FlashBoys/go-finance) (:star: 499) - Comprehensive financial markets data in Go.
* [go-money](https://github.com/rhymond/go-money) (:star: 362) - Implementation of Fowler's Money pattern.
* [ofxgo](https://github.com/aclindsa/ofxgo) (:star: 22) - Query OFX servers and/or parse the responses (with example command-line client).
* [vat](https://github.com/dannyvankooten/vat) (:star: 38) - VAT number validation & EU VAT rates.

## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind) (:star: 17) - Bind form data to any Go values.
* [binding](https://github.com/mholt/binding) (:star: 628) - Binds form and JSON data from net/http Request to struct.
* [conform](https://github.com/leebenson/conform) (:star: 103) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
* [form](https://github.com/go-playground/form) (:star: 217) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
* [formam](https://github.com/monoculum/formam) (:star: 81) - decode form's values into a struct.
* [forms](https://github.com/albrow/forms) (:star: 75) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
* [gorilla/csrf](https://github.com/gorilla/csrf) (:star: 248) - CSRF protection for Go web applications & services.
* [nosurf](https://github.com/justinas/nosurf) (:star: 797) - CSRF protection middleware for Go.

## Game Development

*Awesome game development libraries.*

* [Azul3D](https://github.com/azul3d/engine) (:star: 320) - 3D game engine written in Go.
* [Ebiten](https://github.com/hajimehoshi/ebiten) (:star: 499) - simple 2D game library in Go.
* [engo](https://github.com/EngoEngine/engo) (:star: 684) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.
* [GarageEngine](https://github.com/vova616/GarageEngine) (:star: 280) - 2d game engine written in Go working on OpenGL.
* [glop](https://github.com/runningwild/glop) (:star: 78) - Glop (Game Library Of Power) is a fairly simple cross-platform game library.
* [go-astar](https://github.com/beefsack/go-astar) (:star: 221) - Go implementation of the A\* path finding algorithm.
* [go-collada](https://github.com/GlenKelley/go-collada) (:star: 11) - Go package for working with the Collada file format.
* [go-sdl2](https://github.com/veandco/go-sdl2) (:star: 708) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [go3d](https://github.com/ungerik/go3d) (:star: 128) - Performance oriented 2D/3D math package for Go.
* [gonet](https://github.com/xtaci/gonet) (:star: 875) - Game server skeleton implemented with golang.
* [goworld](https://github.com/xiaonanln/goworld) (:star: 404) - Scalable game server engine, featuring space-entity framework and hot-swapping
* [Leaf](https://github.com/name5566/leaf) (:star: 1686) - Lightweight game server framework.
* [nano](https://github.com/lonnng/nano) (:star: 317) - Lightweight, facility, high performance golang based game server framework
* [Oak](https://github.com/oakmound/oak) (:star: 412) - Pure Go game engine.
* [Pixel](https://github.com/faiface/pixel) (:star: 799) - Hand-crafted 2D game library in Go.
* [raylib-go](https://github.com/gen2brain/raylib-go) (:star: 179) - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [termloop](https://github.com/JoelOtter/termloop) (:star: 812) - Terminal-based game engine for Go, built on top of Termbox.

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [efaceconv](https://github.com/t0pep0/efaceconv) (:star: 25) - Code generation tool for high performance conversion from interface{} to immutable type without allocations.
* [gen](https://github.com/clipperhouse/gen) (:star: 841) - Code generation tool for ‘generics’-like functionality.
* [go-enum](https://github.com/abice/go-enum) (:star: 23) - Code generation for enums from code comments.
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) (:star: 1319) - .NET LINQ-like query methods for Go.
* [goderive](https://github.com/awalterschulze/goderive) (:star: 478) - Derives functions from input types.
* [interfaces](https://github.com/rjeczalik/interfaces) (:star: 124) - Command line tool for generating interface definitions.
* [jennifer](https://github.com/dave/jennifer) (:star: 552) - Generate arbitrary Go code without templates.
* [pkgreflect](https://github.com/ungerik/pkgreflect) (:star: 59) - Go preprocessor for package scoped reflection.

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) (:star: 5851) - Compiler from Go to JavaScript.
* [llgo](https://github.com/go-llvm/llgo) (:star: 854) - LLVM-based compiler for Go.
* [tardisgo](https://github.com/tardisgo/tardisgo) (:star: 355) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.

## Goroutines

*Tools for managing and working with Goroutines.*

* [go-floc](https://github.com/workanator/go-floc) (:star: 101) - Orchestrate goroutines with ease.
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) (:star: 58) - Control goroutines execution order.
* [GoSlaves](https://github.com/themester/GoSlaves) (:star: 8) - Simple and Asynchronous Goroutine pool library.
* [goworker](https://github.com/benmanns/goworker) (:star: 1844) - goworker is a Go-based background worker.
* [grpool](https://github.com/ivpusic/grpool) (:star: 261) - Lightweight Goroutine pool.
* [pool](https://github.com/go-playground/pool) (:star: 327) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.
* [semaphore](https://github.com/kamilsk/semaphore) (:star: 14) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
* [tunny](https://github.com/Jeffail/tunny) (:star: 632) - Goroutine pool for golang.
* [worker-pool](https://github.com/vardius/worker-pool) (:star: 5) - goworker is a Go simple async worker pool.
* [workerpool](https://github.com/gammazero/workerpool) (:star: 5) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [app](https://github.com/murlokswarm/app) (:star: 1905) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.
* [go-astilectron](https://github.com/asticode/go-astilectron) (:star: 817) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).
* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK.
* [go-qml](https://github.com/go-qml/qml) (:star: 1828) - QML support for the Go language.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) (:star: 783) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.
* [goqt](https://github.com/visualfc/goqt) (:star: 1314) - Golang bindings to the Qt cross-platform application framework.
* [gotk3](https://github.com/gotk3/gotk3) (:star: 316) - Go bindings for GTK3.
* [gowd](https://github.com/dtylman/gowd) (:star: 61) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.
* [qt](https://github.com/therecipe/qt) (:star: 3218) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).
* [ui](https://github.com/andlabs/ui) (:star: 4586) - Platform-native GUI library for Go. Cross platform.
* [walk](https://github.com/lxn/walk) (:star: 2178) - Windows application library kit for Go.
* [webview](https://github.com/zserge/webview) (:star: 458) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).

*Interaction*

* [gosx-notifier](https://github.com/deckarep/gosx-notifier) (:star: 396) - OSX Desktop Notifications library for Go.
* [robotgo](https://github.com/go-vgo/robotgo) (:star: 2270) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.
* [systray](https://github.com/getlantern/systray) (:star: 280) - Cross platform Go library to place an icon and menu in the notification area.
* [trayhost](https://github.com/shurcooL/trayhost) (:star: 111) - Cross-platform Go library to place an icon in the host operating system's taskbar.


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bild](https://github.com/anthonynsimon/bild) (:star: 1593) - Collection of image processing algorithms in pure Go.
* [bimg](https://github.com/h2non/bimg) (:star: 410) - Small package for fast and efficient image processing using libvips.
* [geopattern](https://github.com/pravj/geopattern) (:star: 925) - Create beautiful generative image patterns from a string.
* [gg](https://github.com/fogleman/gg) (:star: 1122) - 2D rendering in pure Go.
* [gift](https://github.com/disintegration/gift) (:star: 972) - Package of image processing filters.
* [go-cairo](https://github.com/ungerik/go-cairo) (:star: 63) - Go binding for the cairo graphics library.
* [go-gd](https://github.com/bolknote/go-gd) (:star: 41) - Go binding for GD library.
* [go-nude](https://github.com/koyachi/go-nude) (:star: 236) - Nudity detection with Go.
* [go-opencv](https://github.com/lazywei/go-opencv) (:star: 817) - Go bindings for OpenCV.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) (:star: 21) - Port of webcolors library from Python to Go.
* [gocv](https://github.com/hybridgroup/gocv) (:star: 449) - Go package for computer vision using OpenCV 3.3+.
* [govatar](https://github.com/o1egl/govatar) (:star: 209) - Library and CMD tool for generating funny avatars.
* [imagick](https://github.com/gographics/imagick) (:star: 648) - Go binding to ImageMagick's MagickWand C API.
* [imaginary](https://github.com/h2non/imaginary) (:star: 1458) - Fast and simple HTTP microservice for image resizing.
* [imaging](https://github.com/disintegration/imaging) (:star: 1335) - Simple Go image processing package.
* [img](https://github.com/hawx/img) (:star: 108) - Selection of image manipulation tools.
* [ln](https://github.com/fogleman/ln) (:star: 2057) - 3D line art rendering in Go.
* [mpo](https://github.com/donatj/mpo) (:star: 3) - Decoder and conversion tool for MPO 3D Photos.
* [picfit](https://github.com/thoas/picfit) (:star: 767) - An image resizing server written in Go.
* [pt](https://github.com/fogleman/pt) (:star: 1488) - Path tracing engine written in Go.
* [resize](https://github.com/nfnt/resize) (:star: 1586) - Image resizing for Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez) (:star: 138) - Image resizing in pure Go and SIMD.
* [smartcrop](https://github.com/muesli/smartcrop) (:star: 981) - Finds good crops for arbitrary images and crop sizes.
* [svgo](https://github.com/ajstarks/svgo) (:star: 945) - Go Language Library for SVG generation.
* [tga](https://github.com/ftrvxmtrx/tga) (:star: 15) - Package tga is a TARGA image format decoder/encoder.

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb) (:star: 81) - Open-Source Platform for Quantified Self & IoT.
* [devices](https://github.com/goiot/devices) (:star: 192) - Suite of libraries for IoT devices, experimental for x/exp/io.
* [eywa](https://github.com/xcodersun/eywa) (:star: 11) - Project Eywa is essentially a connection manager that keeps track of connected devices.
* [flogo](https://github.com/tibcosoftware/flogo) (:star: 313) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt) (:star: 582) - Gatt is a Go package for building Bluetooth Low Energy peripherals.
* [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [mainflux](https://github.com/Mainflux/mainflux) (:star: 293) - Industrial IoT Messaging and Device Management Server.
* [sensorbee](https://github.com/sensorbee/sensorbee) (:star: 122) - Lightweight stream processing engine for IoT.

## Logging

*Libraries for generating and working with log files.*

* [glg](https://github.com/kpango/glg) (:star: 4) - glg is simple and fast leveled logging library for Go.
* [glog](https://github.com/golang/glog) (:star: 1588) - Leveled execution logs for Go.
* [go-cronowriter](https://github.com/utahta/go-cronowriter) (:star: 4) - Simple writer that rotate log files automatically based on current date and time, like cronolog.
* [go-log](https://github.com/siddontang/go-log) (:star: 14) - Log lib supports level and multi handlers.
* [go-log](https://github.com/ian-kent/go-log) (:star: 24) - Log4j implementation in Go.
* [go-logger](https://github.com/apsdehal/go-logger) (:star: 167) - Simple logger of Go Programs, with level handlers.
* [gologger](https://github.com/sadlil/gologger) (:star: 27) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.
* [gomol](https://github.com/aphistic/gomol) (:star: 8) - Multiple-output, structured logging for Go with extensible logging outputs.
* [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.
* [log](https://github.com/apex/log) (:star: 466) - Structured logging package for Go.
* [log](https://github.com/go-playground/log) (:star: 225) - Simple, configurable and scalable Structured Logging for Go.
* [log](https://github.com/teris-io/log) (:star: 13) - Structured log interface for Go cleanly separates logging facade from its implementation.
* [log-voyage](https://github.com/firstrow/logvoyage) (:star: 73) - Full-featured logging saas written in golang.
* [log15](https://github.com/inconshreveable/log15) (:star: 653) - Simple, powerful logging for Go.
* [logdump](https://github.com/ewwwwwqm/logdump) (:star: 4) - Package for multi-level logging.
* [logex](https://github.com/chzyer/logex) (:star: 29) - Golang log lib, supports tracking and level, wrap by standard log lib.
* [logger](https://github.com/azer/logger) (:star: 94) - Minimalistic logging library for Go.
* [logo](https://github.com/mbndr/logo) (:star: 1) - Golang logger to different configurable writers.
* [logrus](https://github.com/Sirupsen/logrus) (:star: 6012) - Structured logger for Go.
* [logrusly](https://github.com/sebest/logrusly) (:star: 17) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
* [logutils](https://github.com/hashicorp/logutils) (:star: 183) - Utilities for slightly better logging in Go (Golang) extending the standard logger.
* [logxi](https://github.com/mgutz/logxi) (:star: 314) - 12-factor app logger that is fast and makes you happy.
* [lumberjack](https://github.com/natefinch/lumberjack) (:star: 648) - Simple rolling logger, implements io.WriteCloser.
* [mlog](https://github.com/jbrodriguez/mlog) (:star: 12) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) (:star: 74) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
* [seelog](https://github.com/cihub/seelog) (:star: 1021) - Logging functionality with flexible dispatching, filtering, and formatting.
* [spew](https://github.com/davecgh/go-spew) (:star: 1687) - Implements a deep pretty printer for Go data structures to aid in debugging.
* [stdlog](https://github.com/alexcesaro/log) (:star: 43) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
* [tail](https://github.com/hpcloud/tail) (:star: 874) - Go package striving to emulate the features of the BSD tail program.
* [xlog](https://github.com/xfxdev/xlog) (:star: 1) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
* [xlog](https://github.com/rs/xlog) (:star: 116) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching.
* [zap](https://github.com/uber-go/zap) (:star: 3046) - Fast, structured, leveled logging in Go.
* [zerolog](https://github.com/rs/zerolog) (:star: 581) - Zero-allocation JSON logger.

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) (:star: 511) - Naive Bayesian Classification for Golang.
* [CloudForest](https://github.com/ryanbressler/CloudForest) (:star: 560) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
* [fonet](https://github.com/Fontinalis/fonet) (:star: 2) - A Deep Neural Network library written in Go.
* [gago](https://github.com/MaxHalford/gago) (:star: 394) - Multi-population, flexible, parallel genetic algorithm.
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) (:star: 7) - Go implementation of the k-modes and k-prototypes clustering algorithms.
* [go-fann](https://github.com/white-pony/go-fann) (:star: 92) - Go bindings for Fast Artificial Neural Networks(FANN) library.
* [go-galib](https://github.com/thoj/go-galib) (:star: 154) - Genetic Algorithms library written in Go / golang.
* [go-pr](https://github.com/daviddengcn/go-pr) (:star: 49) - Pattern recognition package in Go lang.
* [gobrain](https://github.com/goml/gobrain) (:star: 222) - Neural Networks written in go.
* [godist](https://github.com/e-dard/godist) (:star: 17) - Various probability distributions, and associated methods.
* [goga](https://github.com/tomcraven/goga) (:star: 59) - Genetic algorithm library for Go.
* [GoLearn](https://github.com/sjwhitworth/golearn) (:star: 4871) - General Machine Learning library for Go.
* [golinear](https://github.com/danieldk/golinear) (:star: 33) - liblinear bindings for Go.
* [goml](https://github.com/cdipaolo/goml) (:star: 789) - On-line Machine Learning in Go.
* [goRecommend](https://github.com/timkaye11/goRecommend) (:star: 92) - Recommendation Algorithms library written in Go.
* [gorgonia](https://github.com/chewxy/gorgonia) (:star: 1476) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.
* [goscore](https://github.com/asafschers/goscore) (:star: 7) - Go Scoring API for PMML.
* [gosseract](https://github.com/otiai10/gosseract) (:star: 383) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.
* [libsvm](https://github.com/datastream/libsvm) (:star: 53) - libsvm golang version derived work based on LIBSVM 3.14.
* [mlgo](https://github.com/NullHypothesis/mlgo) (:star: 3) - This project aims to provide minimalistic machine learning algorithms in Go.
* [neat](https://github.com/jinyeom/neat) (:star: 35) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).
* [neural-go](https://github.com/schuyler/neural-go) (:star: 55) - Multilayer perceptron network implemented in Go, with training via backpropagation.
* [probab](https://github.com/ThePaw/probab) (:star: 7) - Probability distribution functions. Bayesian inference. Written in pure Go.
* [regommend](https://github.com/muesli/regommend) (:star: 177) - Recommendation & collaborative filtering engine.
* [shield](https://github.com/eaigner/shield) (:star: 109) - Bayesian text classifier with flexible tokenizers and storage backends for Go.
* [tfgo](https://github.com/galeone/tfgo) (:star: 634) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.

## Messaging

*Libraries that implement messaging systems.*

* [Centrifugo](https://github.com/centrifugal/centrifugo) (:star: 2280) - Real-time messaging (Websockets or SockJS) server in Go.
* [dbus](https://github.com/godbus/dbus) (:star: 183) - Native Go bindings for D-Bus.
* [drone-line](https://github.com/appleboy/drone-line) (:star: 43) - Sending [Line](https://business.line.me/en/services/bot) notifications using a binary, docker or Drone CI.
* [emitter](https://github.com/olebedev/emitter) (:star: 162) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
* [event](https://github.com/agoalofalife/event) (:star: 4) - Implementation of the pattern observer.
* [EventBus](https://github.com/asaskevich/EventBus) (:star: 310) - The lightweight event bus with async compatibility.
* [gaurun-client](https://github.com/osamingo/gaurun-client) (:star: 4) - Gaurun Client written in Go.
* [Glue](https://github.com/desertbit/glue) (:star: 248) - Robust Go and Javascript Socket Library (Alternative to Socket.io).
* [go-notify](https://github.com/TheCreeper/go-notify) (:star: 32) - Native implementation of the freedesktop notification spec.
* [go-nsq](https://github.com/nsqio/go-nsq) (:star: 924) - the official Go package for NSQ.
* [go-socket.io](https://github.com/googollee/go-socket.io) (:star: 1935) - socket.io library for golang, a realtime application framework.
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) (:star: 2) - Client library to Viessmann Vitotrol web service.
* [Gollum](https://github.com/trivago/gollum) (:star: 568) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.
* [golongpoll](https://github.com/jcuga/golongpoll) (:star: 352) - HTTP longpoll server library that makes web pub-sub simple.
* [goose](https://github.com/ian-kent/goose) (:star: 30) - Server Sent Events in Go.
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) (:star: 1603) - gopush-cluster is a go push server cluster.
* [gorush](https://github.com/appleboy/gorush) (:star: 1569) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).
* [guble](https://github.com/smancke/guble) (:star: 113) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.
* [machinery](https://github.com/RichardKnop/machinery) (:star: 1840) - Asynchronous task queue/job queue based on distributed message passing.
* [mangos](https://github.com/go-mangos/mangos) (:star: 1249) - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.
* [melody](https://github.com/olahol/melody) (:star: 746) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.
* [messagebus](https://github.com/vardius/message-bus) (:star: 4) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
* [NATS Go Client](https://github.com/nats-io/nats) (:star: 1386) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) (:star: 19) - A tiny wrapper around NSQ topic and channel.
* [oplog](https://github.com/dailymotion/oplog) (:star: 84) - Generic oplog/replication system for REST APIs.
* [pubsub](https://github.com/tuxychandru/pubsub) (:star: 154) - Simple pubsub package for go.
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) (:star: 32) - RapidMQ is a lightweight and reliable library for managing of the local messages queue.
* [sarama](https://github.com/Shopify/sarama) (:star: 2307) - Go library for Apache Kafka.
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) (:star: 933) - Redis backed unified push service for server-side notifications to mobile devices.
* [zmq4](https://github.com/pebbe/zmq4) (:star: 579) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).

## Miscellaneous

*These libraries were placed here because none of the other categories seemed to fit.*

* [alice](https://github.com/magic003/alice) (:star: 7) - Additive dependency injection container for Golang.
* [archiver](https://github.com/mholt/archiver) (:star: 715) - Library and command for making and extracting .zip and .tar.gz archives.
* [autoflags](https://github.com/artyom/autoflags) (:star: 18) - Go package to automatically define command line flags from struct fields.
* [avgRating](https://github.com/kirillDanshin/avgRating) (:star: 3) - Calculate average score and rating based on Wilson Score Equation.
* [banner](https://github.com/dimiro1/banner) (:star: 147) - Add beautiful banners into your Go applications.
* [battery](https://github.com/distatus/battery) (:star: 81) - Cross-platform, normalized battery information library.
* [bitio](https://github.com/icza/bitio) (:star: 43) - Highly optimized bit-level Reader and Writer for Go.
* [browscap_go](https://github.com/digitalcrab/browscap_go) (:star: 25) - GoLang Library for [Browser Capabilities Project](http://browscap.org/).
* [captcha](https://github.com/steambap/captcha) (:star: 4) - Package captcha provides an easy to use, unopinionated API for captcha generation.
* [conv](https://github.com/cstockton/go-conv) (:star: 308) - Package conv provides fast and intuitive conversions across Go types.
* [datacounter](https://github.com/miolini/datacounter) (:star: 20) - Go counters for readers/writer/http.ResponseWriter.
* [errors](https://github.com/pkg/errors) (:star: 2463) - Package that provides simple error handling primitives.
* [go-chat-bot](https://github.com/go-chat-bot/bot) (:star: 273) - IRC, Slack & Telegram bot written in Go.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) (:star: 411) - Generic object pool for Golang.
* [go-multierror](https://github.com/hashicorp/go-multierror) (:star: 362) - Go (golang) package for representing a list of errors as a single error.
* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.
* [go-resiliency](https://github.com/eapache/go-resiliency) (:star: 546) - Resiliency patterns for golang.
* [go-sarah](https://github.com/oklahomer/go-sarah) (:star: 30) - Framework to build bot for desired chat services including LINE, Slack, Gitter and more.
* [go-unarr](https://github.com/gen2brain/go-unarr) (:star: 17) - Decompression library for RAR, TAR, ZIP and 7z archives.
* [go.uuid](https://github.com/satori/go.uuid) (:star: 1524) - Implementation of Universally Unique Identifier (UUID). Supported both creation and parsing of UUIDs.
* [gofakeit](https://github.com/brianvoe/gofakeit) (:star: 53) - Random data generator written in go.
* [goid](https://github.com/jakehl/goid) (:star: 11) - Generate and Parse RFC4122 compliant V4 UUIDs.
* [gopsutil](https://github.com/shirou/gopsutil) (:star: 2135) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).
* [gosms](https://github.com/haxpax/gosms) (:star: 1064) - Your own local SMS gateway in Go that can be used to send SMS.
* [gountries](https://github.com/pariz/gountries) (:star: 143) - Package that exposes country and subdivision data.
* [hanu](https://github.com/sbstjn/hanu) (:star: 38) - Framework for writing Slack bots.
* [health](https://github.com/dimiro1/health) (:star: 256) - Easy to use, extensible health check library.
* [healthcheck](https://github.com/etherlabsio/healthcheck) (:star: 18) - An opinionated and concurrent health-check HTTP handler for RESTful services.
* [hostutils](https://github.com/Wing924/hostutils) (:star: 0) - A golang library for packing and unpacking FQDNs list.
* [indigo](https://github.com/osamingo/indigo) (:star: 24) - Distributed unique ID generator of using Sonyflake and encoded by Base58.
* [jobs](https://github.com/albrow/jobs) (:star: 407) - Persistent and flexible background jobs library.
* [lk](https://github.com/hyperboloide/lk) (:star: 22) - A simple licensing library for golang.
* [margelet](https://github.com/zhulik/margelet) (:star: 49) - Framework for building Telegram bots.
* [persian](https://github.com/mavihq/persian) (:star: 3) - Some utilities for Persian language in go.
* [secdl](https://github.com/xor-gate/secdl) (:star: 5) - Lighttpd ModSecDownload algorithm ported to go to secure download urls.
* [shellwords](https://github.com/Wing924/shellwords) (:star: 0) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.
* [shortid](https://github.com/teris-io/shortid) (:star: 180) - Distributed generation of super short, unique, non-sequential, URL friendly IDs.
* [slacker](https://github.com/shomali11/slacker) (:star: 136) - Easy to use framework to create Slack bots.
* [stats](https://github.com/go-playground/stats) (:star: 81) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...
* [turtle](https://github.com/hackebrot/turtle) (:star: 28) - Emojis for Go.
* [uuid](https://github.com/agext/uuid) (:star: 3) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and ouput handling.
* [werr](https://github.com/txgruppi/werr) (:star: 5) - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.
* [xkg](https://github.com/go-xkg/xkg) (:star: 28) - X Keyboard Grabber.
* [xstrings](https://github.com/huandu/xstrings) (:star: 410) - Collection of useful string functions ported from other languages.

## Natural Language Processing

*Libraries for working with human languages.*

* [dpar](https://github.com/danieldk/dpar/) - Transition-based statistical dependency parser.
* [go-eco](https://github.com/ThePaw/go-eco) (:star: 3) - Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.
* [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text.
* [go-mystem](https://github.com/dveselov/mystem) (:star: 11) - CGo bindings to Yandex.Mystem - russian morphology analyzer.
* [go-nlp](https://github.com/nuance/go-nlp) (:star: 70) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work.
* [go-stem](https://github.com/agonopol/go-stem) (:star: 41) - Implementation of the porter stemming algorithm.
* [go-unidecode](https://github.com/mozillazg/go-unidecode) (:star: 19) - ASCII transliterations of Unicode text.
* [go2vec](https://github.com/danieldk/go2vec) (:star: 22) - Reader and utility functions for word2vec embeddings.
* [gojieba](https://github.com/yanyiwu/gojieba) (:star: 422) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) (:star: 12) - Go bindings for the snowball libstemmer library including porter 2.
* [gounidecode](https://github.com/fiam/gounidecode) (:star: 57) - Unicode transliterator (also known as unidecode) for Go.
* [gse](https://github.com/go-ego/gse) (:star: 139) - Go efficient text segmentation; support english, chinese, japanese and other.
* [icu](https://github.com/goodsign/icu) (:star: 15) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
* [libtextcat](https://github.com/goodsign/libtextcat) (:star: 9) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
* [MMSEGO](https://github.com/awsong/MMSEGO) (:star: 53) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
* [nlp](https://github.com/Shixzie/nlp) (:star: 303) - Extract values from strings and fill your structs with nlp.
* [nlp](https://github.com/james-bowman/nlp) (:star: 69) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).
* [paicehusk](https://github.com/rookii/paicehusk) (:star: 23) - Golang implementation of the Paice/Husk Stemming Algorithm.
* [petrovich](https://github.com/striker2000/petrovich) (:star: 5) - Petrovich is the library which inflects Russian names to given grammatical case.
* [porter](https://github.com/a2800276/porter) (:star: 6) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.
* [porter2](https://github.com/zhenjl/porter2) (:star: 29) - Really fast Porter 2 stemmer.
* [prose](https://github.com/jdkato/prose) (:star: 482) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more.
* [RAKE.go](https://github.com/Obaied/RAKE.go) (:star: 19) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
* [segment](https://github.com/blevesearch/segment) (:star: 34) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)
* [sentences](https://github.com/neurosnap/sentences) (:star: 215) - Sentence tokenizer:  converts text into a list of sentences.
* [shamoji](https://github.com/osamingo/shamoji) (:star: 2) - The shamoji is word filtering package written in Go.
* [snowball](https://github.com/goodsign/snowball) (:star: 15) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
* [stemmer](https://github.com/dchest/stemmer) (:star: 36) - Stemmer packages for Go programming language. Includes English and German stemmers.
* [textcat](https://github.com/pebbe/textcat) (:star: 51) - Go package for n-gram based text categorization, with support for utf-8 and raw text.
* [whatlanggo](https://github.com/abadojack/whatlanggo) (:star: 249) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).
* [when](https://github.com/olebedev/when) (:star: 752) - Natural EN and RU language date/time parser with pluggable rules.

## Networking

*Libraries for working with various layers of the network.*

* [arp](https://github.com/mdlayher/arp) (:star: 109) - Package arp implements the ARP protocol, as described in RFC 826.
* [buffstreams](https://github.com/stabbycutyou/buffstreams) (:star: 181) - Streaming protocolbuffer data over TCP made easy.
* [canopus](https://github.com/zubairhamed/canopus) (:star: 87) - CoAP Client/Server implementation (RFC 7252).
* [cidranger](https://github.com/yl2chen/cidranger) (:star: 220) - Fast IP to CIDR lookup for Go.
* [dhcp6](https://github.com/mdlayher/dhcp6) (:star: 29) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
* [dns](https://github.com/miekg/dns) (:star: 2427) - Go library for working with DNS.
* [ether](https://github.com/songgao/ether) (:star: 46) - Cross-platform Go package for sending and receiving ethernet frames.
* [ethernet](https://github.com/mdlayher/ethernet) (:star: 120) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.
* [fasthttp](https://github.com/valyala/fasthttp) (:star: 5335) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.
* [ftp](https://github.com/jlaffaye/ftp) (:star: 280) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).
* [go-getter](https://github.com/hashicorp/go-getter) (:star: 426) - Go library for downloading files or directories from various sources using a URL.
* [go-stun](https://github.com/ccding/go-stun) (:star: 182) - Go implementation of the STUN client (RFC 3489 and RFC 5389).
* [gobgp](https://github.com/osrg/gobgp) (:star: 1090) - BGP implemented in the Go Programming Language.
* [golibwireshark](https://github.com/sunwxg/golibwireshark) (:star: 9) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.
* [gopacket](https://github.com/google/gopacket) (:star: 1582) - Go library for packet processing with libpcap bindings.
* [gopcap](https://github.com/akrennmair/gopcap) (:star: 286) - Go wrapper for libpcap.
* [goshark](https://github.com/sunwxg/goshark) (:star: 7) - Package goshark use tshark to decode IP packet and create data struct to analyse packet.
* [gosnmp](https://github.com/soniah/gosnmp) (:star: 250) - Native Go library for performing SNMP actions.
* [gotcp](https://github.com/gansidui/gotcp) (:star: 301) - Go package for quickly writing tcp applications.
* [grab](https://github.com/cavaliercoder/grab) (:star: 242) - Go package for managing file downloads.
* [graval](https://github.com/koofr/graval) (:star: 17) - Experimental FTP server framework.
* [jazigo](https://github.com/udhos/jazigo) (:star: 68) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices.
* [kcp-go](https://github.com/xtaci/kcp-go) (:star: 919) - KCP - Fast and Reliable ARQ Protocol.
* [kcptun](https://github.com/xtaci/kcptun) (:star: 6908) - Extremely simple & fast udp tunnel based on KCP protocol.
* [lhttp](https://github.com/fanux/lhttp) (:star: 357) - Powerful websocket framework, build your IM server more easily.
* [linkio](https://github.com/ian-kent/linkio) (:star: 32) - Network link speed simulation for Reader/Writer interfaces.
* [llb](https://github.com/kirillDanshin/llb) (:star: 4) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.
* [mdns](https://github.com/hashicorp/mdns) (:star: 363) - Simple mDNS (Multicast DNS) client/server library in Golang.
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [portproxy](https://github.com/aybabtme/portproxy) (:star: 32) - Simple TCP proxy which adds CORS support to API's which don't support it.
* [publicip](https://github.com/polera/publicip) (:star: 12) - Package publicip returns your public facing IPv4 address (internet egress).
* [raw](https://github.com/mdlayher/raw) (:star: 128) - Package raw enables reading and writing data at the device driver level for a network interface.
* [sftp](https://github.com/pkg/sftp) (:star: 425) - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.
* [ssh](https://github.com/gliderlabs/ssh) (:star: 534) - Higher-level API for building SSH servers (wraps crypto/ssh).
* [sslb](https://github.com/eduardonunesp/sslb) (:star: 96) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.
* [stun](https://github.com/go-rtc/stun) (:star: 46) - Go implementation of RFC 5389 STUN protocol.
* [tcp_server](https://github.com/firstrow/tcp_server) (:star: 152) - Go library for building tcp servers faster.
* [utp](https://github.com/anacrolix/utp) (:star: 118) - Go uTP micro transport protocol implementation.
* [water](https://github.com/songgao/water) (:star: 332) - Simple TUN/TAP library.
* [winrm](https://github.com/masterzen/winrm) (:star: 136) - Go WinRM client to remotely execute commands on Windows machines.
* [xtcp](https://github.com/xfxdev/xtcp) (:star: 26) - TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.
* [YANNFF](https://github.com/intel-go/yanff) (:star: 190) - Framework for rapid development of performant network functions for cloud and bare-metal.

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) (:star: 410) - Go bindings for OpenGL (generated via glow).
* [glfw](https://github.com/go-gl/glfw) (:star: 424) - Go bindings for GLFW 3.
* [goxjs/gl](https://github.com/goxjs/gl) (:star: 103) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
* [goxjs/glfw](https://github.com/goxjs/glfw) (:star: 42) - Go cross-platform glfw library for creating an OpenGL context and receiving events.
* [mathgl](https://github.com/go-gl/mathgl) (:star: 190) - Pure Go math package specialized for 3D math, with inspiration from GLM.

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [go-pg](https://github.com/go-pg/pg) (:star: 1285) - PostgreSQL ORM with focus on PostgreSQL specific features and performance.
* [go-queryset](https://github.com/jirfag/go-queryset) (:star: 231) - 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.
* [go-store](https://github.com/gosuri/go-store) (:star: 85) - Simple and fast Redis backed key-value store library for Go.
* [gomodel](https://github.com/cosiner/gomodel) (:star: 56) - Lightweight, fast, orm-like library helps interactive with database.
* [GORM](https://github.com/jinzhu/gorm) (:star: 7175) - The fantastic ORM library for Golang, aims to be developer friendly.
* [gorp](https://github.com/go-gorp/gorp) (:star: 2613) - Go Relational Persistence, ORM-ish library for Go.
* [lore](https://github.com/abrahambotros/lore) (:star: 3) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.
* [Marlow](https://github.com/dadleyy/marlow) (:star: 14) - Generated ORM from project structs for compile time safety assurances.
* [pop/soda](https://github.com/markbates/pop) (:star: 394) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [QBS](https://github.com/coocood/qbs) (:star: 472) - Stands for Query By Struct. A Go ORM.
* [reform](https://github.com/go-reform/reform) (:star: 568) - Better ORM for Go, based on non-empty interfaces and code generation.
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) (:star: 830) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.
* [upper.io/db](https://github.com/upper/db) (:star: 1061) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
* [Xorm](https://github.com/go-xorm/xorm) (:star: 2542) - Simple and powerful ORM for Go.
* [Zoom](https://github.com/albrow/zoom) (:star: 176) - Blazing-fast datastore and querying engine built on Redis.

## Package Management

*Libraries for package and dependency management.*

* [dep](https://github.com/golang/dep) (:star: 6008) - Go dependency tool.
* [gigo](https://github.com/LyricalSecurity/gigo) (:star: 187) - PIP-like dependency tool for golang, with support for private repositories and hashes.
* [glide](https://github.com/Masterminds/glide) (:star: 5767) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.
* [godep](https://github.com/tools/godep) (:star: 5054) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.
* [gom](https://github.com/mattn/gom) (:star: 1290) - Go Manager - bundle for go.
* [goop](https://github.com/nitrous-io/goop) (:star: 769) - Simple dependency manager for Go (golang), inspired by Bundler.
* [gop](https://github.com/lunny/gop) (:star: 32) - Build and manage your Go applications out of GOPATH
* [gopm](https://github.com/gpmgo/gopm) (:star: 1595) - Go Package Manager.
* [govendor](https://github.com/kardianos/govendor) (:star: 2729) - Go Package Manager. Go vendor tool that works with the standard vendor file.
* [gpm](https://github.com/pote/gpm) (:star: 1180) - Barebones dependency manager for Go.
* [gvt](https://github.com/FiloSottile/gvt) (:star: 703) - `gvt` is a simple vendoring tool made for Go native vendoring (aka GO15VENDOREXPERIMENT), based on gb-vendor.
* [johnny-deps](https://github.com/VividCortex/johnny-deps) (:star: 215) - Minimal dependency version using Git.
* [nut](https://github.com/jingweno/nut) (:star: 247) - Vendor Go dependencies.
* [VenGO](https://github.com/DamnWidget/VenGO) (:star: 105) - create and manage exportable isolated go virtual environments.

## Query Language

* [graphql](https://github.com/tmc/graphql) (:star: 43) - graphql parser + utilities.
* [graphql](https://github.com/sevki/graphql) (:star: 33) - GraphQL implementation in go.
* [graphql](https://github.com/neelance/graphql-go) (:star: 897) - GraphQL server with a focus on ease of use.
* [graphql-go](https://github.com/graphql-go/graphql) (:star: 2227) - Implementation of GraphQL for Go.
* [jsonql](https://github.com/elgs/jsonql) (:star: 114) - JSON query expression library in Golang.

## Resource Embedding

* [esc](https://github.com/mjibson/esc) (:star: 254) - Embeds files into Go programs and provides http.FileSystem interfaces to them.
* [fileb0x](https://github.com/UnnoTed/fileb0x) (:star: 182) - Simple tool to embed files in go with focus on "customization" and ease to use.
* [go-embed](https://github.com/pyros2097/go-embed) (:star: 9) - Generates go code to embed resource files into your library or executable.
* [go-resources](https://github.com/omeid/go-resources) (:star: 130) - Unfancy resources embedding with Go.
* [go.rice](https://github.com/GeertJohan/go.rice) (:star: 1112) - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.
* [statics](https://github.com/go-playground/statics) (:star: 44) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.
* [statik](https://github.com/rakyll/statik) (:star: 824) - Embeds static files into a Go executable.
* [templify](https://github.com/wlbr/templify) (:star: 6) - Embed external template files into Go code to create single file binaries.
* [vfsgen](https://github.com/shurcooL/vfsgen) (:star: 215) - Generates a vfsdata.go file that statically implements the given virtual filesystem.

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [blas](https://github.com/ziutek/blas) (:star: 114) - Implementation of BLAS (Basic Linear Algebra Subprograms).
* [chart](https://github.com/vdobler/chart) (:star: 452) - Simple Chart Plotting library for Go. Supports many graphs types.
* [evaler](https://github.com/soniah/evaler) (:star: 30) - Simple floating point arithmetic expression evaluator.
* [ewma](https://github.com/VividCortex/ewma) (:star: 202) - Exponentially-weighted moving averages.
* [geom](https://github.com/skelterjohn/geom) (:star: 33) - 2D geometry for golang.
* [go-dsp](https://github.com/mjibson/go-dsp) (:star: 498) - Digital Signal Processing for Go.
* [go-fn](https://github.com/ematvey/go-fn) (:star: 6) - Mathematical functions written in Go language, that are not covered by math pkg.
* [go-gt](https://github.com/ThePaw/go-gt) (:star: 3) - Graph theory algorithms written in "Go" language.
* [go.matrix](https://github.com/skelterjohn/go.matrix) (:star: 300) - linear algebra for go (has been stalled).
* [gocomplex](https://github.com/varver/gocomplex) (:star: 3) - Complex number library for the Go programming language.
* [goent](https://github.com/kzahedi/goent) (:star: 3) - GO Implementation of Entropy Measures
* [gofrac](https://github.com/anschelsc/gofrac) (:star: 6) - (goinstallable) fractions library for go with support for basic arithmetic.
* [gohistogram](https://github.com/VividCortex/gohistogram) (:star: 103) - Approximate histograms for data streams.
* [gonum/mat64](https://github.com/gonum/matrix) (:star: 462) - The general purpose package for matrix computation. Package mat64 provides basic linear algebra operations for float64 matrices.
* [gonum/plot](https://github.com/gonum/plot) (:star: 686) - gonum/plot provides an API for building and drawing plots in Go.
* [goraph](https://github.com/gyuho/goraph) (:star: 425) - Pure Go graph theory library(data structure, algorith visualization).
* [gosl](https://github.com/cpmech/gosl) (:star: 958) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.
* [gostat](https://github.com/ematvey/gostat) (:star: 22) - Statistics library for the go language.
* [graph](https://github.com/yourbasic/graph) (:star: 19) - Library of basic graph algorithms.
* [ode](https://github.com/ChristopherRabotin/ode) (:star: 4) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.
* [pagerank](https://github.com/alixaxel/pagerank) (:star: 32) - Weighted PageRank algorithm implemented in Go.
* [PiHex](https://github.com/claygod/PiHex) (:star: 3) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.
* [sparse](https://github.com/james-bowman/sparse) (:star: 18) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.
* [stats](https://github.com/montanaflynn/stats) (:star: 794) - Statistics package with common functions missing from the Golang standard library.
* [streamtools](https://github.com/nytlabs/streamtools) (:star: 1303) - general purpose, graphical tool for dealing with streams of data.
* [vectormath](https://github.com/spate/vectormath) (:star: 59) - Vectormath for Go, an adaptation of the scalar C functions from Sony's Vector Math library, as found in the Bullet-2.79 source code (currently inactive).

## Security

*Libraries that are used to help make your application more secure.*

* [acmetool](https://github.com/hlandau/acme) (:star: 1374) - ACME (Let's Encrypt) client tool with automatic renewal.
* [BadActor](https://github.com/jaredfolkins/badactor) (:star: 211) - In-memory, application-driven jailer built in the spirit of fail2ban.
* [go-yara](https://github.com/hillu/go-yara) (:star: 65) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) (:star: 6) - A probably paranoid package for securely hashing and encrypting passwords.
* [lego](https://github.com/xenolf/lego) (:star: 2180) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt).
* [memguard](https://github.com/awnumar/memguard) (:star: 660) - A pure Go library for handling sensitive values in memory.
* [passlib](https://github.com/hlandau/passlib) (:star: 171) - Futureproof password hashing library.
* [secure](https://github.com/unrolled/secure) (:star: 810) - HTTP middleware for Go that facilitates some quick security wins.
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) (:star: 120) - Scrypt package with a simple, obvious API and automatic cost calibration built-in.
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) (:star: 104) - encrypt/decrypt using ssh keys.

## Serialization

*Libraries and tools for binary serialization.*

* [asn1](https://github.com/PromonLogicalis/asn1) (:star: 18) - Asn.1 BER and DER encoding library for golang.
* [bambam](https://github.com/glycerine/bambam) (:star: 57) - generator for Cap'n Proto schemas from go.
* [colfer](https://github.com/pascaldekloe/colfer) (:star: 299) - Code generation for the Colfer binary format.
* [go-capnproto](https://github.com/glycerine/go-capnproto) (:star: 254) - Cap'n Proto library and parser for go.
* [go-codec](https://github.com/ugorji/go) (:star: 849) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.
* [gogoprotobuf](https://github.com/gogo/protobuf) (:star: 1307) - Protocol Buffers for Go with Gadgets.
* [goprotobuf](https://github.com/golang/protobuf) (:star: 2319) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.
* [jsoniter](https://github.com/json-iterator/go) (:star: 1611) - High-performance 100% compatible drop-in replacement of "encoding/json".
* [mapstructure](https://github.com/mitchellh/mapstructure) (:star: 1156) - Go library for decoding generic map values into native Go structures.
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) (:star: 75) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.
* [structomap](https://github.com/tuvistavie/structomap) (:star: 65) - Library to easily and dynamically generate maps from static structures.

## Server Applications

* [algernon](https://github.com/xyproto/algernon) (:star: 390) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
* [Caddy](https://github.com/mholt/caddy) (:star: 14938) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
* [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
* [devd](https://github.com/cortesi/devd) (:star: 2454) - Local webserver for developers.
* [etcd](https://github.com/coreos/etcd) (:star: 15679) - Highly-available key value store for shared configuration and service discovery.
* [Fider](https://github.com/getfider/fider) (:star: 161) - Fider is an open platform to collect and organize customer feedback.
* [minio](https://github.com/minio/minio) (:star: 9197) - Minio is a distributed object storage server.
* [nsq](http://nsq.io/) - A realtime distributed messaging platform.
* [yakvs](https://github.com/sci4me/yakvs) (:star: 21) - Small, networked, in-memory key-value store.

## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace) (:star: 636) - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.
* [amber](https://github.com/eknkc/amber) (:star: 761) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.
* [damsel](https://github.com/dskinner/damsel) (:star: 19) - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.
* [ego](https://github.com/benbjohnson/ego) (:star: 351) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
* [fasttemplate](https://github.com/valyala/fasttemplate) (:star: 159) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).
* [gofpdf](https://github.com/jung-kurt/gofpdf) (:star: 1939) - PDF document generator with high level support for text, drawing and images.
* [grender](https://github.com/dannyvankooten/grender) (:star: 66) - small wrapper around html/template for file-based templates that support extending other template files.
* [hero](https://github.com/shiyanhui/hero) (:star: 847) - Hero is a handy, fast and powerful go template engine.
* [jet](https://github.com/CloudyKit/jet) (:star: 421) - Jet template engine.
* [kasia.go](https://github.com/ziutek/kasia.go) (:star: 70) - Templating system for HTML and other text documents - go implementation.
* [liquid](https://github.com/osteele/liquid) (:star: 33) - Go implementation of Shopify Liquid templates.
* [mustache](https://github.com/hoisie/mustache) (:star: 864) - Go implementation of the Mustache template language.
* [pongo2](https://github.com/flosch/pongo2) (:star: 1095) - Django-like template-engine for Go.
* [quicktemplate](https://github.com/valyala/quicktemplate) (:star: 858) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.
* [raymond](https://github.com/aymerick/raymond) (:star: 199) - Complete handlebars implementation in Go.
* [Razor](https://github.com/sipin/gorazor) (:star: 603) - Razor view engine for Golang.
* [Soy](https://github.com/robfig/soy) (:star: 120) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).
* [velvet](https://github.com/gobuffalo/velvet) (:star: 52) - Complete handlebars implementation in Go.

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/go-playground/assert) (:star: 8) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions.
    * [badio](https://github.com/cavaliercoder/badio) (:star: 5) - Extensions to Go's `testing/iotest` package.
    * [baloo](https://github.com/h2non/baloo) (:star: 442) - Expressive and versatile end-to-end HTTP API testing made easy.
    * [bro](https://github.com/marioidival/bro) (:star: 16) - Watch files in directory and run tests for them.
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) (:star: 12) - Simple snapshot testing addon for your test framework.
    * [dbcleaner](https://github.com/khaiql/dbcleaner) (:star: 23) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby.
    * [dsunit](https://github.com/viant/dsunit) (:star: 5) - Datastore testing for SQL, NoSQL, structured files.
    * [frisby](https://github.com/verdverm/frisby) (:star: 210) - REST API testing framework.
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) (:star: 164) - Tool for viewing test coverage in terminal.
    * [go-mutesting](https://github.com/zimmski/go-mutesting) (:star: 117) - Mutation testing for Go source code.
    * [go-vcr](https://github.com/dnaeon/go-vcr) (:star: 174) - Record and replay your HTTP interactions for fast, deterministic and accurate tests.
    * [goblin](https://github.com/franela/goblin) (:star: 419) - Mocha like testing framework fo Go.
    * [gocheck](http://labix.org/gocheck) - More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload.
    * [godog](https://github.com/DATA-DOG/godog) (:star: 284) - Cucumber or Behat like BDD framework for Go.
    * [gofight](https://github.com/appleboy/gofight) (:star: 121) - API Handler Testing for Golang Router framework.
    * [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
    * [GoSpec](https://github.com/orfjackal/gospec) (:star: 108) - BDD-style testing framework for the Go programming language.
    * [gospecify](https://github.com/stesla/gospecify) (:star: 50) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
    * [gosuite](https://github.com/pavlo/gosuite) (:star: 4) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) (:star: 24) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
    * [httpexpect](https://github.com/gavv/httpexpect) (:star: 784) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing.
    * [restit](https://github.com/yookoala/restit) (:star: 43) - Go micro framework to help writing RESTful API integration test.
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) (:star: 128) - A helper for Rails' like test fixtures to test database applications.
    * [Testify](https://github.com/stretchr/testify) (:star: 4057) - Sacred extension to the standard go testing package.
    * [wstest](https://github.com/posener/wstest) (:star: 28) - Websocket client for unit-testing a websocket http.Handler.

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) (:star: 177) - Tool for generating self-contained mock objects.
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) (:star: 630) - Mock SQL driver for testing database interactions.
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) (:star: 43) - Single transaction based database driver mainly for testing purposes.
    * [gock](https://github.com/h2non/gock) (:star: 382) - Versatile HTTP mocking made easy.
    * [gomock](https://github.com/golang/mock) (:star: 969) - Mocking framework for the Go programming language.
    * [govcr](https://github.com/seborama/govcr) (:star: 28) - HTTP mock for Golang: record and replay HTTP interactions for offline testing.
    * [minimock](https://github.com/gojuno/minimock) (:star: 92) - Mock generator for Go interfaces.
    * [mockhttp](https://github.com/tv42/mockhttp) (:star: 19) - Mock object for Go http.ResponseWriter.

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) (:star: 2013) - Randomized testing system.
    * [gofuzz](https://github.com/google/gofuzz) (:star: 368) - Library for populating go objects with random values.
    * [Tavor](https://github.com/zimmski/tavor) (:star: 172) - Generic fuzzing and delta-debugging framework.

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp) (:star: 65) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.
    * [chromedp](https://github.com/knq/chromedp) (:star: 1561) - Way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.
    * [ggr](https://github.com/aandryashin/ggr) (:star: 5) - Lightweight server that routes and proxies Selenium Wedriver requests to multiple Selenium hubs.
    * [selenoid](https://github.com/aandryashin/selenoid) (:star: 3) - alternative Selenium hub server that launches browsers within containers.

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align) (:star: 32) - A general purpose application that aligns text.
    * [allot](https://github.com/sbstjn/allot) (:star: 26) - Placeholder and wildcard text parsing for CLI tools and bots.
    * [bbConvert](https://github.com/CalebQ42/bbConvert) (:star: 2) - Converts bbCode to HTML that allows you to add support for custom bbCode tags.
    * [blackfriday](https://github.com/russross/blackfriday) (:star: 2748) - Markdown processor in Go.
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) (:star: 776) - HTML Sanitizer.
    * [colly](https://github.com/asciimoo/colly) (:star: 2416) - Fast and Elegant Scraping Framework for Gophers
    * [doi](https://github.com/hscells/doi) (:star: 0) - Document object identifier (doi) parser in Go.
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) (:star: 18) - Editorconfig file parser and manipulator for Go.
    * [enca](https://github.com/endeveit/enca) (:star: 4) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).
    * [genex](https://github.com/alixaxel/genex) (:star: 41) - Count and expand Regular Expressions into all matching Strings.
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) (:star: 3) - Fixed-width text formatting (encoder/decoder with reflection).
    * [go-humanize](https://github.com/dustin/go-humanize) (:star: 1214) - Formatters for time, numbers, and memory size to human readable format.
    * [go-nmea](https://github.com/adrianmo/go-nmea) (:star: 34) - NMEA parser library for the Go language.
    * [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) (:star: 308) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs.
    * [go-runewidth](https://github.com/mattn/go-runewidth) (:star: 105) - Functions to get fixed width of the character or string.
    * [go-slugify](https://github.com/mozillazg/go-slugify) (:star: 14) - Make pretty slug with multiple languages support.
    * [go-vcard](https://github.com/emersion/go-vcard) (:star: 5) - Parse and format vCard.
    * [gofeed](https://github.com/mmcdole/gofeed) (:star: 708) - Parse RSS and Atom feeds in Go.
    * [gographviz](https://github.com/awalterschulze/gographviz) (:star: 145) - Parses the Graphviz DOT language.
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - Format bytes to string.
    * [gonameparts](https://github.com/polera/gonameparts) (:star: 24) - Parses human names into individual name parts.
    * [goq](https://github.com/andrewstuart/goq) (:star: 62) - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).
    * [GoQuery](https://github.com/PuerkitoBio/goquery) (:star: 4897) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.
    * [goregen](https://github.com/zach-klippenstein/goregen) (:star: 27) - Library for generating random strings from regular expressions.
    * [gotext](https://github.com/leonelquinteros/gotext) (:star: 142) - GNU gettext utilities for Go.
    * [guesslanguage](https://github.com/endeveit/guesslanguage) (:star: 35) - Functions to determine the natural language of a unicode text.
    * [inject](https://github.com/facebookgo/inject) (:star: 744) - Package inject provides a reflect based injector.
    * [mxj](https://github.com/clbanning/mxj) (:star: 213) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.
    * [sh](https://github.com/mvdan/sh) (:star: 778) - Shell parser and formatter.
    * [slug](https://github.com/gosimple/slug) (:star: 148) - URL-friendly slugify with multiple languages support.
    * [Slugify](https://github.com/avelino/slugify) (:star: 14) - Go slugify application that handles string.
    * [toml](https://github.com/BurntSushi/toml) (:star: 1702) - TOML configuration format (encoder/decoder with reflection).
* Utility
    * [gotabulate](https://github.com/bndr/gotabulate) (:star: 163) - Easily pretty-print your tabular data with Go.
    * [kace](https://github.com/codemodus/kace) (:star: 5) - Common case conversions covering common initialisms.
    * [parseargs-go](https://github.com/nproc/parseargs-go) (:star: 3) - string argument parser that understands quotes and backslashes.
    * [parth](https://github.com/codemodus/parth) (:star: 16) - URL path segmentation parsing.
    * [radix](https://github.com/yourbasic/radix) (:star: 22) - fast string sorting algorithm.
    * [xj2go](https://github.com/stackerzzq/xj2go) (:star: 6) - Convert xml or json to go struct.
    * [xurls](https://github.com/mvdan/xurls) (:star: 308) - Extract urls from text.

## Third-party APIs

*Libraries for accessing third party APIs.*

* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) (:star: 21) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).
* [anaconda](https://github.com/ChimeraCoder/anaconda) (:star: 802) - Go client library for the Twitter 1.1 API.
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) (:star: 3571) - The official AWS SDK for the Go programming language.
* [brewerydb](https://github.com/naegelejd/brewerydb) (:star: 12) - Go library for accessing the BreweryDB API.
* [cachet](https://github.com/andygrunwald/cachet) (:star: 39) - Go client library for [Cachet (open source status page system)](https://cachethq.io/).
* [circleci](https://github.com/jszwedko/go-circleci) (:star: 17) - Go client library for interacting with CircleCI's API.
* [clarifai](https://github.com/samuelcouch/clarifai) (:star: 54) - Go client library for interfacing with the Clarifai API.
* [discordgo](https://github.com/bwmarrin/discordgo) (:star: 445) - Go bindings for the Discord Chat API.
* [ethrpc](https://github.com/onrik/ethrpc) (:star: 15) - Go bindings for Ethereum JSON RPC API.
* [facebook](https://github.com/huandu/facebook) (:star: 549) - Go Library that supports the Facebook Graph API.
* [fcm](https://github.com/maddevsio/fcm) (:star: 23) - Go library for Firebase Cloud Messaging.
* [gads](https://github.com/emiddleton/gads) (:star: 28) - Google Adwords Unofficial API.
* [gami](https://github.com/bit4bit/gami) (:star: 22) - Go library for Asterisk Manager Interface.
* [gcm](https://github.com/Aorioli/gcm) (:star: 30) - Go library for Google Cloud Messaging.
* [geo-golang](https://github.com/codingsince1985/geo-golang) (:star: 210) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](http://open.mapquestapi.com/nominatim/), [OpenCage](http://geocoder.opencagedata.com/api.html), [HERE](https://developer.here.com/rest-apis/documentation/geocoder), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.
* [ghost](https://github.com/neuegram/ghost) (:star: 21) - Go Library for accessing the Snapchat API.
* [github](https://github.com/google/go-github) (:star: 3070) - Go library for accessing the GitHub REST API v3.
* [githubql](https://github.com/shurcooL/githubql) (:star: 157) - Go library for accessing the GitHub GraphQL API v4.
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) (:star: 3) - Tiny Go client for HackerNews API.
* [go-imgur](https://github.com/koffeinsource/go-imgur) (:star: 5) - Go client library for [imgur](https://imgur.com)
* [go-jira](https://github.com/andygrunwald/go-jira) (:star: 188) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)
* [go-marathon](https://github.com/gambol99/go-marathon) (:star: 155) - Go library for interacting with Mesosphere's Marathon PAAS.
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) (:star: 9) - Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) (:star: 4) - Go client library for the SPTrans Olho Vivo API.
* [go-telegraph](https://github.com/toby3d/go-telegraph) (:star: 40) - Telegraph publishing platform API client.
* [go-tgbot](https://github.com/olebedev/go-tgbot) (:star: 60) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.
* [go-trending](https://github.com/andygrunwald/go-trending) (:star: 82) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
* [go-twitch](https://github.com/knspriggs/go-twitch) (:star: 8) - Go client for interacting with the Twitch v3 API.
* [go-twitter](https://github.com/dghubble/go-twitter) (:star: 344) - Go client library for the Twitter v1.1 APIs.
* [go-unsplash](https://github.com/hbagdi/go-unsplash) (:star: 2) - Go client library for the [Unsplash.com](https://unsplash.com) API.
* [go-xkcd](https://github.com/nishanths/go-xkcd) (:star: 26) - Go client for the xkcd API.
* [goamz](https://github.com/mitchellh/goamz) (:star: 670) - Popular fork of [goamz](https://launchpad.net/goamz) which adds some missing API calls to certain packages.
* [golyrics](https://github.com/mamal72/golyrics) (:star: 19) - Golyrics is a Go library to fetch music lyrics data from the Wikia website.
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) (:star: 25) - Go MusicBrainz WS2 client library.
* [google](https://github.com/google/google-api-go-client) (:star: 1197) - Auto-generated Google APIs for Go.
* [google-analytics](https://github.com/chonthu/go-google-analytics) (:star: 8) - Simple wrapper for easy google analytics reporting.
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) (:star: 913) - Google Cloud APIs Go Client Library.
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) (:star: 4) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).
* [gostorm](https://github.com/jsgilmore/gostorm) (:star: 100) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.
* [govkbot](https://github.com/nikepan/govkbot) (:star: 12) - Simple Go [VK](https://vk.com) bot library.
* [hipchat](https://github.com/andybons/hipchat) (:star: 105) - This project implements a golang client library for the Hipchat API.
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) (:star: 105) - A golang package to communicate with HipChat over XMPP.
* [Medium](https://github.com/Medium/medium-sdk-go) (:star: 90) - Golang SDK for Medium's OAuth2 API.
* [megos](https://github.com/andygrunwald/megos) (:star: 46) - Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.
* [micha](https://github.com/onrik/micha) (:star: 4) - Go Library for Telegram bot api.
* [minio-go](https://github.com/minio/minio-go) (:star: 389) - Minio Go Library for Amazon S3 compatible cloud storage.
* [mixpanel](https://github.com/dukex/mixpanel) (:star: 22) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
* [patreon-go](https://github.com/mxpv/patreon-go) (:star: 6) - Go library for Patreon API.
* [paypal](https://github.com/logpacker/paypalsdk) (:star: 181) - Wrapper for PayPal payment API.
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) (:star: 0) - The Playlyfe Rest API Go SDK.
* [pushover](https://github.com/gregdel/pushover) (:star: 35) - Go wrapper for the Pushover API.
* [rrdaclient](https://github.com/Omie/rrdaclient) (:star: 4) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.
* [shopify](https://github.com/rapito/go-shopify) (:star: 16) - Go Library to make CRUD request to the Shopify API.
* [slack](https://github.com/nlopes/slack) (:star: 1321) - Slack API in Go.
* [smite](https://github.com/sergiotapia/smitego) (:star: 8) - Go package to wraps access to the Smite game API.
* [spotify](https://github.com/rapito/go-spotify) (:star: 12) - Go Library to access Spotify WEB API.
* [steam](https://github.com/sostronk/go-steam) (:star: 8) - Go Library to interact with Steam game servers.
* [stripe](https://github.com/stripe/stripe-go) (:star: 624) - Go client for the Stripe API.
* [tbot](https://github.com/yanzay/tbot) (:star: 128) - Telegram bot server with API similar to net/http.
* [telebot](https://github.com/tucnak/telebot) (:star: 350) - Telegram bot framework written in Go.
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) (:star: 763) - Simple and clean Telegram bot client.
* [textbelt](https://github.com/dietsche/textbelt) (:star: 11) - Go client for the textbelt.com txt messaging API.
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) (:star: 10) - Simple golang package to communicate with [themoviedb.org](https://themoviedb.org).
* [translate](https://github.com/poorny/translate) (:star: 19) - Go online translation package.
* [Trello](https://github.com/adlio/trello) (:star: 29) - Go wrapper for the Trello API.
* [tumblr](https://github.com/mattcunningham/gumblr) (:star: 5) - Go wrapper for the Tumblr v2 API.
* [webhooks](https://github.com/go-playground/webhooks) (:star: 121) - Webhook receiver for GitHub and Bitbucket.
* [zooz](https://github.com/gojuno/go-zooz) (:star: 1) - Go client for the Zooz API.

## Utilities

*General utilities and tools to make your life easier.*

* [abutil](https://github.com/bahlo/abutil) (:star: 37) - Collection of often-used Golang helpers.
* [apm](https://github.com/topfreegames/apm) (:star: 98) - Process manager for Golang applications with an HTTP API.
* [boilr](https://github.com/tmrts/boilr) (:star: 587) - Blazingly fast CLI tool for creating projects from boilerplate templates.
* [chyle](https://github.com/antham/chyle) (:star: 37) - Changelog generator using a git repository with multiple configuration possibilities.
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) (:star: 520) - Circuit Breakers in Go.
* [clockwerk](http://github.com/onatm/clockwerk) (:star: 32) - Go package to schedule periodic jobs using a simple, fluent syntax.
* [command](https://github.com/txgruppi/command) (:star: 6) - Command pattern for Go with thread safe serial and parallel dispatcher.
* [coop](https://github.com/rakyll/coop) (:star: 1197) - Cheat sheet for some of the common concurrent flows in Go.
* [copy-pasta](https://github.com/jutkko/copy-pasta) (:star: 21) - Universal multi-workstation clipboard that uses S3 like backend for the storage.
* [ctop](https://github.com/bcicen/ctop) (:star: 6252) - [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.
* [Death](https://github.com/vrecan/death) (:star: 70) - Managing go application shutdown with signals.
* [Deepcopier](https://github.com/ulule/deepcopier) (:star: 125) - Simple struct copying for Go.
* [delve](https://github.com/derekparker/delve) (:star: 6783) - Go debugger.
* [dlog](https://github.com/kirillDanshin/dlog) (:star: 8) - Compile-time controlled logger to make your release smaller without removing debug calls.
* [ergo](https://github.com/cristianoliveira/ergo) (:star: 116) - The management of multiple local services running over different ports made easy.
* [evaluator](https://github.com/nullne/evaluator) (:star: 5) - Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.
* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) (:star: 1586) - Golang library for reading and writing Microsoft Excel™ (XLSX) files.
* [fastlz](https://github.com/digitalcrab/fastlz) (:star: 7) - Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang.
* [filetype](https://github.com/h2non/filetype) (:star: 210) - Small package to infer the file type checking the magic numbers signature.
* [filler](https://github.com/yaronsumel/filler) (:star: 9) - small utility to fill structs using "fill" tag.
* [fzf](https://github.com/junegunn/fzf) (:star: 12052) - Command-line fuzzy finder written in Go.
* [generate](https://github.com/go-playground/generate) (:star: 6) - runs go generate recursively on a specified path or environment variable and can filter by regex.
* [gentleman](https://github.com/h2non/gentleman) (:star: 478) - Full-featured plugin-driven HTTP client library.
* [git-time-metric](https://github.com/git-time-metric/gtm) (:star: 403) - Simple, seamless, lightweight time tracking for Git.
* [GJSON](https://github.com/tidwall/gjson) (:star: 1904) - Get a JSON value with one line of code.
* [go-astitodo](https://github.com/asticode/go-astitodo) (:star: 29) - Parse TODOs in your GO code.
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) (:star: 150) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only).
* [go-cron](https://github.com/rk/go-cron) (:star: 152) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
* [go-debug](https://github.com/tj/go-debug) (:star: 377) - Conditional debug logging for Golang libraries & applications.
* [go-dry](https://github.com/ungerik/go-dry) (:star: 377) - DRY (don't repeat yourself) package for Go.
* [go-funk](https://github.com/thoas/go-funk) (:star: 289) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).
* [go-httpheader](https://github.com/mozillazg/go-httpheader) (:star: 10) - Go library for encoding structs into Header fields.
* [go-rate](https://github.com/beefsack/go-rate) (:star: 225) - Timed rate limiter for Go.
* [go-respond](https://github.com/nicklaw5/go-respond) (:star: 7) - Go package for handling common HTTP JSON responses.
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) (:star: 59) - XML Sitemap generator written in Go.
* [go-torch](https://github.com/uber/go-torch) (:star: 2383) - Stochastic flame graph profiler for Go programs.
* [go-trigger](https://github.com/sadlil/go-trigger) (:star: 112) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.
* [go-underscore](https://github.com/tobyhede/go-underscore) (:star: 926) - Useful collection of helpfully functional Go collection utilities.
* [goback](https://github.com/carlescere/goback) (:star: 33) - Go simple exponential backoff package.
* [godaemon](https://github.com/VividCortex/godaemon) (:star: 332) - Utility to write daemons.
* [godropbox](https://github.com/dropbox/godropbox) (:star: 3432) - Common libraries for writing Go services/applications from Dropbox.
* [gohper](https://github.com/cosiner/gohper) (:star: 227) - Various tools/modules help for development.
* [gojq](https://github.com/elgs/gojq) (:star: 85) - JSON query in Golang.
* [gojson](https://github.com/ChimeraCoder/gojson) (:star: 1500) - Automatically generate Go (golang) struct definitions from example JSON.
* [golarm](https://github.com/msempere/golarm) (:star: 24) - Fire alarms with system events.
* [golog](https://github.com/mlimaloureiro/golog) (:star: 33) - Easy and lightweight CLI tool to time track your tasks.
* [gopencils](https://github.com/bndr/gopencils) (:star: 399) - Small and simple package to easily consume REST APIs.
* [goplaceholder](https://github.com/michiwend/goplaceholder) (:star: 13) - a small golang lib to generate placeholder images.
* [goreleaser](https://github.com/goreleaser/goreleaser) (:star: 2088) - Deliver Go binaries as fast and easily as possible.
* [goreporter](https://github.com/wgliang/goreporter) (:star: 1814) - Golang tool that does static analysis, unit testing, code review and generate code quality report.
* [goreq](https://github.com/franela/goreq) (:star: 628) - Minimal and simple request library for Go language.
* [goreq](https://github.com/smallnest/goreq) (:star: 54) - Enhanced simplified HTTP client based on gorequest.
* [gorequest](https://github.com/parnurzeal/gorequest) (:star: 1430) - Simplified HTTP client with rich features for Go.
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) (:star: 4) - SeaweedFS client library with almost full features.
* [gotenv](https://github.com/subosito/gotenv) (:star: 70) - Load environment variables from `.env` or any `io.Reader` in Go.
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) (:star: 7) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.
* [gpath](https://github.com/tenntenn/gpath) (:star: 20) - Library to simplify access struct fields with Go's expression in reflection.
* [grequests](https://github.com/levigross/grequests) (:star: 958) - Elegant and simple `net/http` wrapper that follows Python's requests library.
* [gron](https://github.com/roylee0704/gron) (:star: 477) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.
* [htcat](https://github.com/htcat/htcat) (:star: 442) - Parallel and Pipelined HTTP GET Utility.
* [httpcontrol](https://github.com/facebookgo/httpcontrol) (:star: 464) - Package httpcontrol allows for HTTP transport level control around timeouts and retries.
* [hub](https://github.com/github/hub) (:star: 11744) - wrap git commands with additional functionality to interact with github from the terminal.
* [hystrix-go](https://github.com/afex/hystrix-go) (:star: 790) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.
* [immortal](https://github.com/immortal/immortal) (:star: 440) - *nix cross-platform (OS agnostic) supervisor.
* [intrinsic](https://github.com/mengzhuo/intrinsic) (:star: 10) - Use x86 SIMD without writing any assembly code.
* [JobRunner](https://github.com/bamzi/jobrunner) (:star: 406) - Smart and featureful cron job scheduler with job queuing and live monitoring built in.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) (:star: 3) - Go bindings based on the JSON API errors reference.
* [jsonf](https://github.com/miolini/jsonf) (:star: 46) - Console tool for highlighted formatting and struct query fetching JSON.
* [jsongo](https://github.com/ricardolonga/jsongo) (:star: 73) - Fluent API to make it easier to create Json objects.
* [jsonhal](https://github.com/RichardKnop/jsonhal) (:star: 6) - Simple Go package to make custom structs marshal into HAL compatible JSON responses.
* [kazaam](https://github.com/Qntfy/kazaam) (:star: 55) - API for arbitrary transformation of JSON documents.
* [lrserver](https://github.com/jaschaephraim/lrserver) (:star: 86) - LiveReload server for Go.
* [mc](https://github.com/minio/mc) (:star: 604) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.
* [mergo](https://github.com/imdario/mergo) (:star: 352) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.
* [minify](https://github.com/tdewolff/minify) (:star: 1271) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.
* [mmake](https://github.com/tj/mmake) (:star: 1305) - Modern Make.
* [moldova](https://github.com/StabbyCutyou/moldova) (:star: 139) - Utility for generating random data based on an input template.
* [mp](https://github.com/sanbornm/mp) (:star: 21) - Simple cli email parser. It currently takes stdin and outputs JSON.
* [mssqlx](https://github.com/linxGnu/mssqlx) (:star: 26) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
* [multitick](https://github.com/VividCortex/multitick) (:star: 51) - Multiplexor for aligned tickers.
* [myhttp](https://github.com/inancgumus/myhttp) (:star: 23) - Simple API to make HTTP GET requests with timeout support.
* [netbug](https://github.com/e-dard/netbug) (:star: 56) - Easy remote profiling of your services.
* [ngrok](https://github.com/inconshreveable/ngrok) (:star: 11957) - Introspected tunnels to localhost.
* [okrun](https://github.com/xta/okrun) (:star: 11) - go run error steamroller.
* [onecache](https://github.com/adelowo/onecache) (:star: 58) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
* [panicparse](https://github.com/maruel/panicparse) (:star: 1527) - Groups similar goroutines and colorizes stack dump.
* [peco](https://github.com/peco/peco) (:star: 4206) - Simplistic interactive filtering tool.
* [pester](https://github.com/sethgrid/pester) (:star: 217) - Go HTTP client calls with retries, backoff, and concurrency.
* [pm](https://github.com/VividCortex/pm) (:star: 50) - Process (i.e. goroutine) manager with an HTTP API.
* [profile](https://github.com/pkg/profile) (:star: 675) - Simple profiling support package for Go.
* [rclient](https://github.com/zpatrick/rclient) (:star: 18) - Readable, flexible, simple-to-use client for REST APIs.
* [realize](https://github.com/tockins/realize) (:star: 1844) - Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.
* [request](https://github.com/mozillazg/request) (:star: 239) - Go HTTP Requests for Humans™.
* [rerate](https://github.com/abo/rerate) (:star: 7) - Redis-based rate counter and rate limiter for Go.
* [rerun](https://github.com/ivpusic/rerun) (:star: 133) - Recompiling and rerunning go apps when source changes.
* [resty](https://github.com/go-resty/resty) (:star: 645) - Simple HTTP and REST client for Go inspired by Ruby rest-client.
* [retry](https://github.com/kamilsk/retry) (:star: 27) - Functional mechanism based on context to perform actions repetitively until successful.
* [robustly](https://github.com/VividCortex/robustly) (:star: 123) - Runs functions resiliently, catching and restarting panics.
* [scheduler](https://github.com/carlescere/scheduler) (:star: 177) - Cronjobs scheduling made easy.
* [sling](https://github.com/dghubble/sling) (:star: 588) - Go HTTP requests builder for API clients.
* [spinner](https://github.com/briandowns/spinner) (:star: 391) - Go package to easily provide a terminal spinner with options.
* [sqlx](https://github.com/jmoiron/sqlx) (:star: 3569) - provides a set of extensions on top of the excellent built-in database/sql package.
* [Storm](https://github.com/asdine/storm) (:star: 715) - Simple and powerful toolkit for BoltDB.
* [structs](https://github.com/PumpkinSeed/structs) (:star: 1) - Implement simple functions to manipulate structs.
* [Task](https://github.com/go-task/task) (:star: 428) - simple "Make" alternative.
* [toolbox](https://github.com/viant/toolbox) (:star: 11) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.
* [ugo](https://github.com/alxrm/ugo) (:star: 14) - ugo is slice toolbox with concise syntax for Go.
* [UNIS](https://github.com/esemplastic/unis) (:star: 63) - Common Architecture™ for String Utilities in Go.
* [usql](https://github.com/knq/usql) (:star: 2113) - usql is a universal command-line interface for SQL databases.
* [util](https://github.com/shomali11/util) (:star: 38) - Collection of useful utility functions. (strings, concurrency, manipulations, ...).
* [wuzz](https://github.com/asciimoo/wuzz) (:star: 6997) - Interactive cli tool for HTTP inspection.
* [xferspdy](https://github.com/monmohan/xferspdy) (:star: 46) - Xferspdy provides binary diff and patch library in golang.
* [xlsx](https://github.com/tealeg/xlsx) (:star: 2042) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.

## Validation

*Libraries for validation.*

* [govalidator](https://github.com/asaskevich/govalidator) (:star: 2211) - Validators and sanitizers for strings, numerics, slices and structs.
* [govalidator](https://github.com/thedevsaddam/govalidator) (:star: 162) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) (:star: 489) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.
* [validate](https://github.com/markbates/validate) (:star: 26) - This package provides a framework for writing validations for Go applications.
* [validator](https://github.com/go-playground/validator) (:star: 1334) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh) (:star: 50) - Scriptable server and net/http middleware for GitHub Webhooks.
* [git2go](https://github.com/libgit2/git2go) (:star: 1071) - Go bindings for libgit2.
* [go-vcs](https://github.com/sourcegraph/go-vcs) (:star: 53) - manipulate and inspect VCS repositories in Go.
* [hgo](https://github.com/beyang/hgo) (:star: 10) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf) (:star: 317) - Go bindings for FFmpeg av\* libraries.
* [go-astisub](https://github.com/asticode/go-astisub) (:star: 42) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).
* [go-astits](https://github.com/asticode/go-astits) (:star: 173) - Parse and demux MPEG Transport Streams (.ts) natively in GO.
* [goav](https://github.com/giorgisio/goav) (:star: 383) - Comphrensive Go bindings for FFmpeg.
* [gst](https://github.com/ziutek/gst) (:star: 135) - Go bindings for GStreamer.
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) (:star: 5) - Subtitle format support for go. Supports .srt, .ttml, and .ass.
* [v4l](https://github.com/korandiz/v4l) (:star: 12) - Video capture library for Linux, written in Go.

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
* [Air](https://github.com/sheng/air) (:star: 49) - Ideal RESTful web framework for Go.
* [Beego](https://github.com/astaxie/beego) (:star: 13121) - beego is an open-source, high-performance web framework for the Go programming language.
* [Buffalo](http://gobuffalo.io) - Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) (:star: 8781) - High performance, minimalist Go web framework.
* [Fireball](https://github.com/zpatrick/fireball) (:star: 29) - More "natural" feeling web framework.
* [Florest](https://github.com/jabong/florest-core) (:star: 30) - High-performance workflow based REST API framework.
* [Gem](https://github.com/go-gem/gem) (:star: 147) - Simple and fast web framework, friendly to REST API.
* [Gin](https://github.com/gin-gonic/gin) (:star: 13195) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
* [Gizmo](https://github.com/NYTimes/gizmo) (:star: 1991) - Microservice toolkit used by the New York Times.
* [go-json-rest](https://github.com/ant0ine/go-json-rest) (:star: 2973) - Quick and easy way to setup a RESTful JSON API.
* [go-relax](https://github.com/codehack/go-relax) (:star: 147) - Framework of pluggable components to build RESTful API's.
* [go-rest](https://github.com/ungerik/go-rest) (:star: 105) - Small and evil REST framework for Go.
* [goa](https://github.com/raphael/goa) (:star: 2422) - Framework for developing microservices based on the design of Ruby's Praxis.
* [Goat](https://github.com/bahlo/goat) (:star: 149) - Minimalistic REST API server in Go.
* [Golf](https://github.com/dinever/golf) (:star: 210) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.
* [Gondola](https://github.com/rainycape/gondola) (:star: 311) - The web framework for writing faster sites, faster.
* [gongular](https://github.com/mustafaakin/gongular) (:star: 384) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection.
* [Macaron](https://github.com/go-macaron/macaron) (:star: 2126) - Macaron is a high productive and modular design web framework in Go.
* [mango](https://github.com/paulbellamy/mango) (:star: 313) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
* [Microservice](https://github.com/claygod/microservice) (:star: 40) - The framework for the creation of microservices, written in Golang.
* [neo](https://github.com/ivpusic/neo) (:star: 336) - Neo is minimal and fast Go Web Framework with extremely simple API.
* [Resoursea](https://github.com/resoursea/api) (:star: 29) - REST framework for quickly writing resource based services.
* [REST Layer](http://rest-layer.io) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Revel](https://github.com/revel/revel) (:star: 9123) - High-productivity web framework for the Go language.
* [rex](https://github.com/goanywhere/rex) (:star: 15) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.
* [sawsij](https://github.com/jaybill/sawsij) (:star: 1) - lightweight, open-source web framework for building high-performance, data-driven web applications.
* [tango](https://github.com/lunny/tango) (:star: 674) - Micro & pluggable web framework for Go.
* [tigertonic](https://github.com/rcrowley/go-tigertonic) (:star: 969) - Go framework for building JSON web services inspired by Dropwizard.
* [traffic](https://github.com/pilu/traffic) (:star: 521) - Sinatra inspired regexp/pattern mux and web framework for Go.
* [utron](https://github.com/gernest/utron) (:star: 1965) - Lightweight MVC framework for Go(Golang).
* [violetear](https://github.com/nbari/violetear) (:star: 77) - Go HTTP router.
* [YARF](https://github.com/yarf-framework/yarf) (:star: 39) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way.
* [Zerver](https://github.com/cosiner/zerver) (:star: 139) - Zerver is an expressive, modular, feature completed RESTful framework.

### Middlewares

#### Actual middlewares

* [CORS](https://github.com/rs/cors) (:star: 628) - Easily add CORS capabilities to your API.
* [formjson](https://github.com/rs/formjson) (:star: 28) - Transparently handle JSON input as a standard form POST.
* [Limiter](https://github.com/ulule/limiter) (:star: 277) - Dead simple rate limit middleware for Go.
* [Tollbooth](https://github.com/didip/tollbooth) (:star: 638) - Rate limit HTTP request handler.
* [XFF](https://github.com/sebest/xff) (:star: 64) - Handle `X-Forwarded-For` header and friends.

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice) (:star: 1421) - Painless middleware chaining for Go.
* [catena](https://github.com/codemodus/catena) (:star: 6) - http.Handler wrapper catenation (same API as "chain").
* [chain](https://github.com/codemodus/chain) (:star: 62) - Handler wrapper chaining with scoped data (net/context-based "middleware").
* [go-wrap](https://github.com/go-on/wrap) (:star: 52) - Small middlewares package for net/http.
* [gores](https://github.com/alioygur/gores) (:star: 66) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
* [interpose](https://github.com/carbocation/interpose) (:star: 279) - Minimalist net/http middleware for golang.
* [muxchain](https://github.com/stephens2424/muxchain) (:star: 199) - Lightweight middleware for net/http.
* [negroni](https://github.com/urfave/negroni) (:star: 5155) - Idiomatic HTTP middleware for Golang.
* [render](https://github.com/unrolled/render) (:star: 985) - Go package for easily rendering JSON, XML, and HTML template responses.
* [renderer](https://github.com/thedevsaddam/renderer) (:star: 43) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.
* [rye](https://github.com/InVisionApp/rye) (:star: 62) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.
* [stats](https://github.com/thoas/stats) (:star: 444) - Go middleware that stores various information about your web application.
* [Volatile](https://github.com/volatile/core) (:star: 81) - Minimalist middleware stack promoting flexibility, good practices and clean code.

### Routers

* [alien](https://github.com/gernest/alien) (:star: 87) - Lightweight and fast http router from outer space.
* [Bone](https://github.com/go-zoo/bone) (:star: 1075) - Lightning Fast HTTP Multiplexer.
* [Bxog](https://github.com/claygod/Bxog) (:star: 78) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
* [chi](https://github.com/go-chi/chi) (:star: 2850) - Small, fast and expressive HTTP router built on net/context.
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) (:star: 389) - High performance router forked from `httprouter`. The first router fit for `fasthttp`.
* [FastRouter](https://github.com/razonyang/fastrouter) (:star: 10) - a fast, flexible HTTP router written in Go.
* [gocraft/web](https://github.com/gocraft/web) (:star: 1261) - Mux and middleware package in Go.
* [Goji](https://github.com/goji/goji) (:star: 562) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.
* [GoRouter](https://github.com/vardius/gorouter) (:star: 21) - GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
* [gowww/router](https://github.com/gowww/router) (:star: 132) - Lightning fast HTTP router fully compatible with the net/http.Handler interface.
* [httprouter](https://github.com/julienschmidt/httprouter) (:star: 5997) - High performance router. Use this and the standard http handlers to form a very high performance web framework.
* [httptreemux](https://github.com/dimfeld/httptreemux) (:star: 275) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.
* [lars](https://github.com/go-playground/lars) (:star: 344) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.
* [medeina](https://github.com/imdario/medeina) (:star: 17) - Medeina is a HTTP routing tree based on HttpRouter, inspired by Roda and Cuba.
* [mux](https://github.com/gorilla/mux) (:star: 4879) - Powerful URL router and dispatcher for golang.
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) (:star: 218) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.
* [pat](https://github.com/bmizerany/pat) (:star: 1085) - Sinatra style pattern muxer for Go’s net/http library, by the author of Sinatra.
* [pure](https://github.com/go-playground/pure) (:star: 43) - Is a lightweight HTTP router that sticks to the std "net/http" implementation.
* [Siesta](https://github.com/VividCortex/siesta) (:star: 343) - Composable framework to write middleware and handlers.
* [vestigo](https://github.com/husobee/vestigo) (:star: 190) - Performant, stand-alone, HTTP compliant URL Router for go web applications.
* [xmux](https://github.com/rs/xmux) (:star: 74) - High performance muxer based on `httprouter` with `net/context` support.
* [zeus](https://github.com/daryl/zeus) (:star: 109) - Very simple and fast HTTP router for Go.

## Windows

* [d3d9](https://github.com/gonutz/d3d9) (:star: 56) - Go bindings for Direct3D9.
* [go-ole](https://github.com/go-ole/go-ole) (:star: 342) - Win32 OLE implementation for golang.

## XML

*Libraries and tools for manipulating XML.*

* [go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) (:star: 128) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions.
* [XML-Comp](https://github.com/xml-comp/xml-comp) (:star: 9) - Simple command line XML comparer that generates diffs of folders, files and tags.
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) (:star: 2) - Procedural XML generation API based on libxml2's xmlwriter module.
* [xpath](https://github.com/antchfx/xpath) (:star: 36) - XPath package for Go.
* [xquery](https://github.com/antchfx/xquery) (:star: 89) - XQuery lets you extract data from HTML/XML documents using XPath expression.

# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat) (:star: 124) - Checks recent changes to a Go project for backwards incompatible changes.
* [dupl](https://github.com/mibk/dupl) (:star: 92) - Tool for code clone detection.
* [errcheck](https://github.com/kisielk/errcheck) (:star: 922) - Errcheck is a program for checking for unchecked errors in Go programs.
* [gcvis](https://github.com/davecheney/gcvis) (:star: 727) - Visualise Go program GC trace data in real time.
* [Go Metalinter](https://github.com/alecthomas/gometalinter) (:star: 2083) - Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form.
* [go-checkstyle](https://github.com/qiniu/checkstyle) (:star: 64) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) (:star: 171) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.
* [go-outdated](https://github.com/firstrow/go-outdated) (:star: 35) - Console application that displays outdated packages.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) (:star: 198) - Web based Golang AST visualizer.
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GoLint](https://github.com/golang/lint) (:star: 2105) - Golint is a linter for Go source code.
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - gosimple is a linter for Go source code that specialises on simplifying code.
* [gostatus](https://github.com/shurcooL/gostatus) (:star: 211) - Command line tool, shows the status of repositories that contain Go packages.
* [interfacer](https://github.com/mvdan/interfacer) (:star: 679) - Linter that suggests interface types.
* [lint](https://github.com/surullabs/lint) (:star: 54) - Run linters as part of go test.
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) (:star: 5) - tarp finds functions and methods without direct unit tests in Go source code.
* [unconvert](https://github.com/mdempsky/unconvert) (:star: 181) - Remove unnecessary type conversions from Go source.
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - unused checks Go code for unused constants, variables, functions and types.
* [validate](https://github.com/mccoyst/validate) (:star: 61) - Automatically validates struct fields with tags.

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
* [go-mode](https://github.com/dominikh/go-mode.el) (:star: 678) - Go mode for GNU/Emacs.
* [go-plus](https://github.com/joefitzgerald/go-plus) (:star: 1313) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.
* [Goclipse](https://github.com/GoClipse/goclipse) (:star: 732) - Eclipse plugin for Go.
* [gocode](https://github.com/nsf/gocode) (:star: 4003) - Autocompletion daemon for the Go programming language.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) (:star: 2756) - Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features.
* [velour](https://github.com/velour/velour) (:star: 14) - IRC client for the acme editor.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) (:star: 74) - Vim plugin to highlight syntax errors on save.
* [vim-go](https://github.com/fatih/vim-go) (:star: 7546) - Go development plugin for Vim.
* [vscode-go](https://github.com/Microsoft/vscode-go) (:star: 3007) - Extension for Visual Studio Code (VS Code) which provides support for the Go language.
* [Watch](https://github.com/eaburns/Watch) (:star: 139) - Runs a command in an acme win on file changes.

## Go Generate Tools

* [generic](https://github.com/usk81/generic) (:star: 5) - flexible data type for Go.
* [genny](https://github.com/cheekybits/genny) (:star: 507) - Elegant generics for Go.
* [gonerics](http://github.com/bouk/gonerics) (:star: 92) - Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests) (:star: 888) - Generate Go tests from your source code.
* [re2dfa](https://github.com/opennota/re2dfa) (:star: 123) - Transform regular expressions into finite state machines and output Go source code.

## Go Tools

* [colorgo](https://github.com/songgao/colorgo) (:star: 83) - Wrapper around `go` command for colorized `go build` output.
* [depth](https://github.com/KyleBanks/depth) (:star: 184) - Visualize dependency trees of any package by analyzing imports.
* [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.
* [go-callvis](https://github.com/TrueFurby/go-callvis) (:star: 954) - Visualize call graph of your Go program using dot format.
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) (:star: 31) - Bash completion for go and wgo.
* [go-swagger](https://github.com/go-swagger/go-swagger) (:star: 1685) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.
* [OctoLinker](https://github.com/OctoLinker/browser-extension) (:star: 2394) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
* [richgo](https://github.com/kyoh86/richgo) (:star: 216) - Enrich `go test` outputs with text decorations.
* [rts](https://github.com/galeone/rts) (:star: 151) - RTS: response to struct. Generates Go structs from server responses.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly) (:star: 1438) - aptly is a Debian repository management tool.
* [aurora](https://github.com/xuri/aurora) (:star: 231) - Cross-platform web-based Beanstalkd queue server console.
* [awsenv](https://github.com/soniah/awsenv) (:star: 12) - Small binary that loads Amazon (AWS) environment variables for a profile.
* [Banshee](https://github.com/eleme/banshee) (:star: 858) - Anomalies detection system for periodic metrics.
* [bombardier](https://github.com/codesenberg/bombardier) (:star: 748) - Fast cross-platform HTTP benchmarking tool.
* [bosun](https://github.com/bosun-monitor/bosun) (:star: 2412) - Time Series Alerting Framework.
* [dogo](https://github.com/liudng/dogo) (:star: 149) - Monitoring changes in the source file and automatically compile and run (restart).
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) (:star: 10) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
* [drone-scp](https://github.com/appleboy/drone-scp) (:star: 18) - Copy files and artifacts via SSH using a binary, docker or Drone CI.
* [Dropship](https://github.com/chrismckenzie/dropship) (:star: 37) - Tool for deploying code via cdn.
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) (:star: 33) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
* [Gitea](https://github.com/go-gitea/gitea) (:star: 4294) - Fork of Gogs, entirely community driven.
* [Go Metrics](https://github.com/rcrowley/go-metrics) (:star: 1735) - Go port of Coda Hale's Metrics library: https://github.com/codahale/metrics.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) (:star: 533) - Enable your Go applications to self update.
* [gobrew](https://github.com/cryptojuice/gobrew) (:star: 167) - gobrew lets you easily switch between multiple versions of go.
* [godbg](https://github.com/sirnewton01/godbg) (:star: 209) - Web-based gdb front-end application.
* [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative) (:star: 280) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
* [govvv](https://github.com/ahmetalpbalkan/govvv) (:star: 199) - “go build” wrapper to easily add version information into Go binaries.
* [gox](https://github.com/mitchellh/gox) (:star: 2394) - Dead simple, no frills Go cross compile tool.
* [goxc](https://github.com/laher/goxc) (:star: 1469) - build tool for Go, with a focus on cross-compiling and packaging.
* [grapes](https://github.com/yaronsumel/grapes) (:star: 84) - Lightweight tool designed to distribute commands over ssh with ease.
* [GVM](https://github.com/moovweb/gvm) (:star: 3191) - GVM provides an interface to manage Go versions.
* [Hey](https://github.com/rakyll/hey) (:star: 2435) - Hey is a tiny program that sends some load to a web application.
* [kala](https://github.com/ajvb/kala) (:star: 1117) - Simplistic, modern, and performant job scheduler.
* [kubernetes](https://github.com/kubernetes/kubernetes) (:star: 29362) - Container Cluster Manager from Google.
* [manssh](https://github.com/xwjdsh/manssh) (:star: 40) - manssh is a command line tool for managing your ssh alias config easily.
* [Moby](https://github.com/moby/moby) (:star: 46565) - Collaborative project for the container ecosystem to assemble container-based systems.
* [Mora](https://github.com/emicklei/mora) (:star: 227) - REST server for accessing MongoDB documents and meta data.
* [ostent](https://github.com/ostrost/ostent) (:star: 148) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.
* [Packer](https://github.com/mitchellh/packer) (:star: 7077) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
* [Pewpew](https://github.com/bengadbois/pewpew) (:star: 130) - Flexible HTTP command line stress tester.
* [Rodent](https://github.com/alouche/rodent) (:star: 31) - Rodent helps you manage Go versions, projects and track dependencies.
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) (:star: 869) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) (:star: 373) - Manage BareMetal Servers from Command Line (as easily as with Docker).
* [sg](https://github.com/ChristopherRabotin/sg) (:star: 1) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the reponse code and data between each call for specific server stress based on its previous response.
* [skm](https://github.com/TimothyYe/skm) (:star: 322) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!
* [StatusOK](https://github.com/sanathp/statusok) (:star: 842) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.
* [Vegeta](https://github.com/tsenart/vegeta) (:star: 6690) - HTTP load testing tool and library. It's over 9000!
* [webhook](https://github.com/adnanh/webhook) (:star: 2453) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.
* [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
* [winrm-cli](https://github.com/masterzen/winrm-cli) (:star: 26) - Cli tool to remotely execute commands on Windows machines.

### Other Software
* [borg](https://github.com/crufter/borg) (:star: 1287) - Terminal based search engine for bash snippets.
* [boxed](https://github.com/tejo/boxed) (:star: 64) - Dropbox based blog engine.
* [Cherry](https://github.com/rafael-santiago/cherry) (:star: 137) - Tiny webchat server in Go.
* [Circuit](https://github.com/gocircuit/circuit) (:star: 1593) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
* [Comcast](https://github.com/tylertreat/Comcast) (:star: 5115) - Simulate bad network connections.
* [confd](https://github.com/kelseyhightower/confd) (:star: 4387) - Manage local application configuration files using templates and data from etcd or consul.
* [DDNS](https://github.com/skibish/ddns) (:star: 32) - Personal DDNS client with Digital Ocean Networking DNS as backend.
* [Docker](http://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community) (:star: 314) - Modern wiki software that integrates data from SaaS tools.
* [fleet](https://github.com/coreos/fleet) (:star: 2398) - Distributed init System.
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) (:star: 797) - App that displays updates for the Go packages in your GOPATH.
* [gocc](https://github.com/goccmack/gocc) (:star: 146) - Gocc is a compiler kit for Go written in Go.
* [GoDNS](https://github.com/timothyye/godns) (:star: 166) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) (:star: 11) - Chrome extension for Go Doc sites, which shows function description as tooltip at funciton list.
* [Gogland](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor) (:star: 8119) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
* [hugo](http://gohugo.io/) - Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide) (:star: 152) - Browser accessible IDE. Designed for Go with Go.
* [ipe](https://github.com/dimiro1/ipe) (:star: 213) - Open source Pusher server implementation compatible with Pusher client libraries written in GO.
* [JayDiff](https://github.com/yazgazan/jaydiff) (:star: 7) - JSON diff utility written in Go.
* [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Leaps](https://github.com/jeffail/leaps) (:star: 542) - Pair programming service using Operational Transforms.
* [limetext](http://limetext.org/) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide) (:star: 4073) - LiteIDE is a simple, open source, cross-platform Go IDE.
* [mockingjay](https://github.com/quii/mockingjay-server) (:star: 334) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.
* [myLG](https://github.com/mehrdadrad/mylg) (:star: 1907) - Command Line Network Diagnostic tool written in Go.
* [naclpipe](https://github.com/unix4fun/naclpipe) (:star: 11) - Simple NaCL EC25519 based crypto pipe tool written in Go.
* [nes](https://github.com/fogleman/nes) (:star: 2949) - Nintendo Entertainment System (NES) emulator written in Go.
* [orange-cat](https://github.com/noraesae/orange-cat) (:star: 160) - Markdown previewer written in Go.
* [Orbit](https://github.com/gulien/orbit) (:star: 69) - A simple tool for running commands and generating files from templates.
* [peg](https://github.com/pointlander/peg) (:star: 445) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
* [Postman](https://github.com/zachlatta/postman) (:star: 693) - Command-line utility for batch-sending email.
* [restic](https://github.com/restic/restic) (:star: 3047) - De-duplicating backup program.
* [rkt](https://github.com/coreos/rkt) (:star: 7531) - App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM.
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) (:star: 4842) - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
* [shell2http](https://github.com/msoap/shell2http) (:star: 162) - Executing shell commands via http server (for prototyping or remote control).
* [snap](https://github.com/intelsdi-x/snap) (:star: 1625) - Powerful telemetry framework.
* [Snitch](https://github.com/lucasgomide/snitch) (:star: 13) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru.
* [Stack Up](https://github.com/pressly/sup) (:star: 1587) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
* [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
* [Tenyks](https://github.com/kyleterry/tenyks) (:star: 162) - Service oriented IRC bot using Redis and JSON for messaging.
* [toto](https://github.com/blogcin/ToTo) (:star: 19) - Simple proxy server written in Go language, can be used together with browser.
* [toxiproxy](https://github.com/shopify/toxiproxy) (:star: 2277) - Proxy to simulate network and system conditions for automated tests.
* [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow) (:star: 291) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.
* [websysd](https://github.com/ian-kent/websysd) (:star: 39) - Web based process manager (like Marathon or Upstart).
* [wellington](https://github.com/wellington/wellington) (:star: 254) - Sass project management tool, extends the language with sprite functions (like Compass).

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench) (:star: 84) - Framework to compare the performance between different Go versions.
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) (:star: 8) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) (:star: 82) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) (:star: 979) - Go HTTP request router benchmark and comparison.
* [go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) (:star: 4) - Naive performance test of two ways to do type assertion in Go.
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) (:star: 583) - Go web framework benchmark.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) (:star: 571) - Benchmarks of Go serialization methods.
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) (:star: 49) - Benchmarks of common basic operations for the Go language.
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) (:star: 13) - Tiny collection of Go micro benchmarks. The intent is to compare some language features to others.
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) (:star: 32) - Collection of benchmarks for popular Go database/SQL utilities.
* [gospeed](https://github.com/feyeleanor/GoSpeed) (:star: 74) - Go micro-benchmarks for calculating the speed of language constructs.
* [kvbench](https://github.com/jimrobinson/kvbench) (:star: 11) - Key/Value database benchmark.
* [skynet](https://github.com/atemerev/skynet) (:star: 804) - Skynet 1M threads microbenchmark.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) (:star: 113) - Compare various Image resize algorithms for the Go language.

## Conferences

* [Capital Go](http://www.capitalgolang.com) - Washington, D.C., USA
* [dotGo](http://www.dotgo.eu) - Paris, France
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan
* [GolangUK](http://golanguk.com/) - London, UK
* [GopherChina](http://gopherchina.org) - Shanghai, China
* [GopherCon](http://www.gophercon.com/) - Denver, USA
* [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, BR
* [GopherCon Dubai](http://www.gophercon.ae/) - Dubai, UAE
* [GopherCon India](http://www.gophercon.in/) - Pune, India
* [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore
* [GothamGo](http://gothamgo.com/) - New York City, USA

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) (:star: 4198) - A curated list of Go books.
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)

## Meetups

* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*Add the group of your city/country here (send **PR**)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## Websites

* [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) (:star: 10603) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) (:star: 20174) - List of other amazingly awesome lists.
* [Flipboard - Go Magazine](https://flipboard.com/section/the-golang-magazine-bVP7nS) - Collection of Go articles and tutorials.
* [Go Blog](http://blog.golang.org) - The official Go blog.
* [Go Challenge](http://golang-challenge.org/) - Learn Go by solving problems and getting feedback from Go experts.
* [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com) - A report card for your Go package.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) (:star: 25) - Collection of Go projects that needs help. Good place to start your open-source way in Go.
* [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
* [Golang Flow](http://golangflow.io) - Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com) - Links and news about Go programming.
* [golang-graphics](https://github.com/mholt/golang-graphics) (:star: 125) - Collection of Go images, graphics, and art.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [r/Golang](https://www.reddit.com/r/golang) - News about Go.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.

### Tutorials

* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) (:star: 19228) - Golang ebook intro how to build a web app with golang.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) (:star: 2503) - Go's reference card.
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
* [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Working with Go](https://github.com/mkaz/working-with-go) (:star: 831) - Intro to go for experienced programmers.

*The repository's stars were updated last at 11/28/2017 - 12:22 (UTC)*