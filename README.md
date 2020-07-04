# Awesome Go

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinorun)

Financial support to Awesome Go

**We have no monthly cost**_, but we have employees **hard working** to maintain the Awesome Go, with money raised we can repay the effort of each person involved!_

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Bot Building](#bot-building)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Dynamic DNS](#dynamic-dns)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Error Handling](#error-handling)
    - [Files](#Files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Functional](#functional)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Geographic](#geographic)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot-internet-of-things)
    - [Job Scheduler](#job-scheduler)
    - [JSON](#json)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Microsoft Office](#microsoft-office)
        - [Microsoft Excel](#microsoft-excel)
    - [Miscellaneous](#miscellaneous)
        - [Dependency Injection](#dependency-injection)
        - [Project Layout](#project-layout)
        - [Strings](#strings)
        - [Uncategorized](#uncategorized)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
        - [HTTP Clients](#http-clients)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Performance](#performance)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Server Applications](#server-applications)
    - [Stream Processing](#stream-processing)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [UUID](#uuid)
    - [Validation](#validation)
    - [Version Control](#version-control)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
        - [Routers](#routers)
    - [WebAssembly](#webassembly)
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

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) - EasyMidi is a simple and reliable library for working with standard midi file (SMF). ![](https://img.shields.io/github/last-commit/algoGuy/EasyMIDI.svg)
* [flac](https://github.com/mewkiz/flac) - Native Go FLAC encoder/decoder with support for FLAC streams. ![](https://img.shields.io/github/last-commit/mewkiz/flac.svg)
* [gaad](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser. ![](https://img.shields.io/github/last-commit/Comcast/gaad.svg)
* [go-sox](https://github.com/krig/go-sox) - libsox bindings for go. ![](https://img.shields.io/github/last-commit/krig/go-sox.svg)
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) - libmediainfo bindings for go. ![](https://img.shields.io/github/last-commit/zhulik/go_mediainfo.svg)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go. ![](https://img.shields.io/github/last-commit/dh1tw/gosamplerate.svg)
* [id3v2](https://github.com/bogem/id3v2) - Fast and powerful ID3 decoding and encodung library written completely in Go. ![](https://img.shields.io/github/last-commit/bogem/id3v2.svg)
* [malgo](https://github.com/gen2brain/malgo) - Mini audio library. ![](https://img.shields.io/github/last-commit/gen2brain/malgo.svg)
* [minimp3](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library. ![](https://img.shields.io/github/last-commit/tosone/minimp3.svg)
* [mix](https://github.com/go-mix/mix) - Sequence-based Go-native audio mixer for music apps. ![](https://img.shields.io/github/last-commit/go-mix/mix.svg)
* [mp3](https://github.com/tcolgate/mp3) - Native Go MP3 decoder. ![](https://img.shields.io/github/last-commit/tcolgate/mp3.svg)
* [music-theory](https://github.com/go-music-theory/music-theory) - Music theory models in Go. ![](https://img.shields.io/github/last-commit/go-music-theory/music-theory.svg)
* [Oto](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms. ![](https://img.shields.io/github/last-commit/hajimehoshi/oto.svg)
* [PortAudio](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library. ![](https://img.shields.io/github/last-commit/gordonklaus/portaudio.svg)
* [portmidi](https://github.com/rakyll/portmidi) - Go bindings for PortMidi. ![](https://img.shields.io/github/last-commit/rakyll/portmidi.svg)
* [taglib](https://github.com/wtolson/go-taglib) - Go bindings for taglib. ![](https://img.shields.io/github/last-commit/wtolson/go-taglib.svg)
* [vorbis](https://github.com/mccoyst/vorbis) - "Native" Go Vorbis decoder (uses CGO, but has no dependencies). ![](https://img.shields.io/github/last-commit/mccoyst/vorbis.svg)
* [waveform](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams. ![](https://img.shields.io/github/last-commit/mdlayher/waveform.svg)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. ![](https://img.shields.io/github/last-commit/volatiletech/authboss.svg)
* [branca](https://github.com/hako/branca) - Golang implementation of Branca Tokens. ![](https://img.shields.io/github/last-commit/hako/branca.svg)
* [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, ABAC. ![](https://img.shields.io/github/last-commit/hsluoyz/casbin.svg)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) - provides parser of cookies.txt file format. ![](https://img.shields.io/github/last-commit/mengzhuo/cookiestxt.svg)
* [go-guardian](https://goreportcard.com/report/github.com/shaj13/go-guardian) - Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token and Certificate based authentication.
* [go-jose](https://github.com/square/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. ![](https://img.shields.io/github/last-commit/square/go-jose.svg)
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) - Standalone, specification-compliant,  OAuth2 server written in Golang. ![](https://img.shields.io/github/last-commit/RichardKnop/go-oauth2-server.svg)
* [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. ![](https://img.shields.io/github/last-commit/dghubble/gologin.svg)
* [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. ![](https://img.shields.io/github/last-commit/mikespook/gorbac.svg)
* [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. ![](https://img.shields.io/github/last-commit/markbates/goth.svg)
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middleware. ![](https://img.shields.io/github/last-commit/goji/httpauth.svg)
* [jeff](https://github.com/abraithwaite/jeff) - Simple, flexible, secure and idiomatic web session management with pluggable backends. ![](https://img.shields.io/github/last-commit/abraithwaite/jeff.svg)
* [jwt](https://github.com/robbert229/jwt) - Clean and easy to use implementation of JSON Web Tokens (JWT). ![](https://img.shields.io/github/last-commit/robbert229/jwt.svg)
* [jwt](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library. ![](https://img.shields.io/github/last-commit/pascaldekloe/jwt.svg)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options. ![](https://img.shields.io/github/last-commit/adam-hanna/jwt-auth.svg)
* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT). ![](https://img.shields.io/github/last-commit/dgrijalva/jwt-go.svg)
* [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. ![](https://img.shields.io/github/last-commit/tarent/loginsrv.svg)
* [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. ![](https://img.shields.io/github/last-commit/golang/oauth2.svg)
* [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. ![](https://img.shields.io/github/last-commit/openshift/osin.svg)
* [paseto](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO). ![](https://img.shields.io/github/last-commit/o1egl/paseto.svg)
* [permissions2](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. ![](https://img.shields.io/github/last-commit/xyproto/permissions2.svg)
* [rbac](https://github.com/zpatrick/rbac) - Minimalistic RBAC package for Go applications. ![](https://img.shields.io/github/last-commit/zpatrick/rbac.svg)
* [scope](https://github.com/SonicRoshan/scope) - Easily Manage OAuth2 Scopes In Go. ![](https://img.shields.io/github/last-commit/SonicRoshan/scope.svg)
* [scs](https://github.com/alexedwards/scs) - Session Manager for HTTP servers. ![](https://img.shields.io/github/last-commit/alexedwards/scs.svg)
* [securecookie](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding. ![](https://img.shields.io/github/last-commit/chmike/securecookie.svg)
* [session](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE). ![](https://img.shields.io/github/last-commit/icza/session.svg)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) - Go session management using the SessionGate Redis module. ![](https://img.shields.io/github/last-commit/f0rmiga/sessiongate-go.svg)
* [sessions](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers. ![](https://img.shields.io/github/last-commit/adam-hanna/sessions.svg)
* [sessionup](https://github.com/swithek/sessionup) - Simple, yet effective HTTP session management and identification package. ![](https://img.shields.io/github/last-commit/swithek/sessionup.svg)
* [signedvalue](https://github.com/sashka/signedvalue) - Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. ![](https://img.shields.io/github/last-commit/sashka/signedvalue.svg)
* [sjwt](https://github.com/brianvoe/sjwt) - Simple jwt generator and parser. ![](https://img.shields.io/github/last-commit/brianvoe/sjwt.svg)

## Bot Building

*Libraries for building and working with bots.*

* [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) - A Discord bot for managing ephemeral roles based upon voice channel member presence. ![](https://img.shields.io/github/last-commit/ewohltman/ephemeral-roles.svg)
* [go-chat-bot](https://github.com/go-chat-bot/bot) - IRC, Slack & Telegram bot written in Go. ![](https://img.shields.io/github/last-commit/go-chat-bot/bot.svg)
* [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.
* [go-sarah](https://github.com/oklahomer/go-sarah) - Framework to build bot for desired chat services including LINE, Slack, Gitter and more. ![](https://img.shields.io/github/last-commit/oklahomer/go-sarah.svg)
* [go-tgbot](https://github.com/olebedev/go-tgbot) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. ![](https://img.shields.io/github/last-commit/olebedev/go-tgbot.svg)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) - Libary to write bots for twitch.tv chat ![](https://img.shields.io/github/last-commit/gempir/go-twitch-irc.svg)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) - A golang implementation of a console-based trading bot for cryptocurrency exchanges. ![](https://img.shields.io/github/last-commit/saniales/golang-crypto-trading-bot.svg)
* [govkbot](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library. ![](https://img.shields.io/github/last-commit/nikepan/govkbot.svg)
* [hanu](https://github.com/sbstjn/hanu) - Framework for writing Slack bots. ![](https://img.shields.io/github/last-commit/sbstjn/hanu.svg)
* [Kelp](https://github.com/stellar/kelp) - official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. ![](https://img.shields.io/github/last-commit/stellar/kelp.svg)
* [margelet](https://github.com/zhulik/margelet) - Framework for building Telegram bots. ![](https://img.shields.io/github/last-commit/zhulik/margelet.svg)
* [micha](https://github.com/onrik/micha) - Go Library for Telegram bot api. ![](https://img.shields.io/github/last-commit/onrik/micha.svg)
* [olivia](https://github.com/olivia-ai/olivia) - A chatbot built with an artificial neural network. ![](https://img.shields.io/github/last-commit/olivia-ai/olivia.svg)
* [slacker](https://github.com/shomali11/slacker) - Easy to use framework to create Slack bots. ![](https://img.shields.io/github/last-commit/shomali11/slacker.svg)
* [slackscot](https://github.com/alexandre-normand/slackscot) - Another framework for building Slack bots. ![](https://img.shields.io/github/last-commit/alexandre-normand/slackscot.svg)
* [tbot](https://github.com/yanzay/tbot) - Telegram bot server with API similar to net/http. ![](https://img.shields.io/github/last-commit/yanzay/tbot.svg)
* [telebot](https://github.com/tucnak/telebot) - Telegram bot framework written in Go. ![](https://img.shields.io/github/last-commit/tucnak/telebot.svg)
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client. ![](https://img.shields.io/github/last-commit/Syfaro/telegram-bot-api.svg)
* [Tenyks](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging. ![](https://img.shields.io/github/last-commit/kyleterry/tenyks.svg)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [1build](https://github.com/gopinath-langote/1build) - Command line tool to frictionlessly manage project-specific commands. ![](https://img.shields.io/github/last-commit/gopinath-langote/1build.svg)
* [argparse](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module. ![](https://img.shields.io/github/last-commit/akamensky/argparse.svg)
* [argv](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax. ![](https://img.shields.io/github/last-commit/cosiner/argv.svg)
* [cli](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags. ![](https://img.shields.io/github/last-commit/mkideal/cli.svg)
* [cli](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go. ![](https://img.shields.io/github/last-commit/teris-io/cli.svg)
* [cli-init](https://github.com/tcnksm/gcli) - The easy way to start building Golang command line applications. ![](https://img.shields.io/github/last-commit/tcnksm/gcli.svg)
* [climax](http://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command.
* [clîr](https://github.com/leaanthony/clir) - A Simple and Clear CLI library. Dependency free. ![](https://img.shields.io/github/last-commit/leaanthony/clir.svg)
* [cmd](https://github.com/posener/cmd) - Extends the standard `flag` package to support sub commands and more in idomatic way. ![](https://img.shields.io/github/last-commit/posener/cmd.svg)
* [cmdr](https://github.com/hedzr/cmdr) - A POSIX/GNU style, getopt-like command-line UI Go library. ![](https://img.shields.io/github/last-commit/hedzr/cmdr.svg)
* [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. ![](https://img.shields.io/github/last-commit/spf13/cobra.svg)
* [commandeer](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. ![](https://img.shields.io/github/last-commit/jaffee/commandeer.svg)
* [complete](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion. ![](https://img.shields.io/github/last-commit/posener/complete.svg)
* [Dnote](https://github.com/dnote/dnote) - A simple command line notebook with multi-device sync. ![](https://img.shields.io/github/last-commit/dnote/dnote.svg)
* [docopt.go](https://github.com/docopt/docopt.go) - Command-line arguments parser that will make you smile. ![](https://img.shields.io/github/last-commit/docopt/docopt.go.svg)
* [env](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs. ![](https://img.shields.io/github/last-commit/codingconcepts/env.svg)
* [flag](https://github.com/cosiner/flag) - Simple but powerful command line option parsing library for Go supporting subcommand. ![](https://img.shields.io/github/last-commit/cosiner/flag.svg)
* [flaggy](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support. ![](https://img.shields.io/github/last-commit/integrii/flaggy.svg)
* [flagvar](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package. ![](https://img.shields.io/github/last-commit/sgreben/flagvar.svg)
* [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. ![](https://img.shields.io/github/last-commit/alexflint/go-arg.svg)
* [go-commander](https://github.com/yitsushi/go-commander) - Go library to simplify CLI workflow. ![](https://img.shields.io/github/last-commit/yitsushi/go-commander.svg)
* [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. ![](https://img.shields.io/github/last-commit/jessevdk/go-flags.svg)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) - Go option parser inspired on the flexibility of Perl’s GetOpt::Long. ![](https://img.shields.io/github/last-commit/DavidGamba/go-getoptions.svg)
* [gocmd](https://github.com/devfacet/gocmd) - Go library for building command line applications. ![](https://img.shields.io/github/last-commit/devfacet/gocmd.svg)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection.
* [job](https://github.com/liujianping/job) - JOB, make your short-term command as a long-term job. ![](https://img.shields.io/github/last-commit/liujianping/job.svg)
* [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands. ![](https://img.shields.io/github/last-commit/alecthomas/kingpin.svg)
* [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. ![](https://img.shields.io/github/last-commit/peterh/liner.svg)
* [mitchellh/cli](https://github.com/mitchellh/cli) - Go library for implementing command-line interfaces. ![](https://img.shields.io/github/last-commit/mitchellh/cli.svg)
* [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation. ![](https://img.shields.io/github/last-commit/jawher/mow.cli.svg)
* [ops](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator. ![](https://img.shields.io/github/last-commit/nanovms/ops.svg)
* [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. ![](https://img.shields.io/github/last-commit/spf13/pflag.svg)
* [readline](https://github.com/chzyer/readline) - Pure golang implementation that provides most features in GNU-Readline under MIT license. ![](https://img.shields.io/github/last-commit/chzyer/readline.svg)
* [sand](https://github.com/Zaba505/sand) - Simple API for creating interpreters and so much more. ![](https://img.shields.io/github/last-commit/Zaba505/sand.svg)
* [sflags](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. ![](https://img.shields.io/github/last-commit/octago/sflags.svg)
* [strumt](https://github.com/antham/strumt) - Library to create prompt chain. ![](https://img.shields.io/github/last-commit/antham/strumt.svg)
* [ts](https://github.com/liujianping/ts) - Timestamp convert & compare tool. ![](https://img.shields.io/github/last-commit/liujianping/ts.svg)
* [ukautz/clif](https://github.com/ukautz/clif) - Small command line interface framework. ![](https://img.shields.io/github/last-commit/ukautz/clif.svg)
* [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). ![](https://img.shields.io/github/last-commit/urfave/cli.svg)
* [wlog](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency. ![](https://img.shields.io/github/last-commit/dixonwille/wlog.svg)
* [wmenu](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompts users to make choices. ![](https://img.shields.io/github/last-commit/dixonwille/wmenu.svg)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. ![](https://img.shields.io/github/last-commit/guptarohit/asciigraph.svg)
* [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that supports fmt.Printf/Sprintf. ![](https://img.shields.io/github/last-commit/logrusorgru/aurora.svg)
* [cfmt](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes. ![](https://img.shields.io/github/last-commit/mingrammer/cfmt.svg)
* [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. ![](https://img.shields.io/github/last-commit/ttacon/chalk.svg)
* [colourize](https://github.com/TreyBastian/colourize) - Go library for ANSI colour text in terminals. ![](https://img.shields.io/github/last-commit/TreyBastian/colourize.svg)
* [ctc](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method. ![](https://img.shields.io/github/last-commit/wzshiming/ctc.svg)
* [go-ataman](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals. ![](https://img.shields.io/github/last-commit/workanator/go-ataman.svg)
* [go-colorable](https://github.com/mattn/go-colorable) - Colorable writer for windows. ![](https://img.shields.io/github/last-commit/mattn/go-colorable.svg)
* [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals. ![](https://img.shields.io/github/last-commit/daviddengcn/go-colortext.svg)
* [go-isatty](https://github.com/mattn/go-isatty) - isatty for golang. ![](https://img.shields.io/github/last-commit/mattn/go-isatty.svg)
* [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). ![](https://img.shields.io/github/last-commit/c-bata/go-prompt.svg)
* [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. ![](https://img.shields.io/github/last-commit/jroimartin/gocui.svg)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
* [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. ![](https://img.shields.io/github/last-commit/gookit/color.svg)
* [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. ![](https://img.shields.io/github/last-commit/vbauerster/mpb.svg)
* [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. ![](https://img.shields.io/github/last-commit/schollz/progressbar.svg)
* [simpletable](https://github.com/alexeyco/simpletable) - Simple tables in terminal with Go. ![](https://img.shields.io/github/last-commit/alexeyco/simpletable.svg)
* [tabby](https://github.com/cheynewallace/tabby) - A tiny library for super simple Golang tables. ![](https://img.shields.io/github/last-commit/cheynewallace/tabby.svg)
* [tabular](https://github.com/InVisionApp/tabular) - Print ASCII tables from command line utilities without the need to pass large sets of data to the API. ![](https://img.shields.io/github/last-commit/InVisionApp/tabular.svg)
* [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. ![](https://img.shields.io/github/last-commit/nsf/termbox-go.svg)
* [termdash](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). ![](https://img.shields.io/github/last-commit/mum4k/termdash.svg)
* [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). ![](https://img.shields.io/github/last-commit/gizak/termui.svg)
* [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in realtime. ![](https://img.shields.io/github/last-commit/gosuri/uilive.svg)
* [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. ![](https://img.shields.io/github/last-commit/gosuri/uiprogress.svg)
* [uitable](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data. ![](https://img.shields.io/github/last-commit/gosuri/uitable.svg)
* [yacspin](https://github.com/theckman/yacspin) - Yet Another CLi Spinner package, for working with terminal spinners. ![](https://img.shields.io/github/last-commit/theckman/yacspin.svg)

## Configuration

*Libraries for configuration parsing.*

* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want). ![](https://img.shields.io/github/last-commit/ilyakaznacheev/cleanenv.svg)
* [config](https://github.com/golobby/config) - A lightweight yet powerful config package for Go projects. ![](https://img.shields.io/github/last-commit/golobby/config.svg)
* [config](https://github.com/JeremyLoy/config) - Cloud native application configuration. Bind ENV to structs in only two lines. ![](https://img.shields.io/github/last-commit/JeremyLoy/config.svg)
* [config](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing. ![](https://img.shields.io/github/last-commit/olebedev/config.svg)
* [configuration](https://github.com/BoRuDar/configuration) - Library for initializing configuration structs from env variables, files, flags and 'default' tag. ![](https://img.shields.io/github/last-commit/BoRuDar/configuration.svg)
* [configure](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables. ![](https://img.shields.io/github/last-commit/paked/configure.svg)
* [configuro](https://github.com/sherifabdlnaby/configuro) - opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. ![](https://img.shields.io/github/last-commit/sherifabdlnaby/configuro.svg)
* [confita](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct. ![](https://img.shields.io/github/last-commit/heetch/confita.svg)
* [conflate](https://github.com/the4thamigo-uk/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. ![](https://img.shields.io/github/last-commit/the4thamigo-uk/conflate.svg)
* [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). ![](https://img.shields.io/github/last-commit/caarlos0/env.svg)
* [envcfg](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs. ![](https://img.shields.io/github/last-commit/tomazk/envcfg.svg)
* [envconf](https://github.com/ian-kent/envconf) - Configuration from environment. ![](https://img.shields.io/github/last-commit/ian-kent/envconf.svg)
* [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables. ![](https://img.shields.io/github/last-commit/vrischmann/envconfig.svg)
* [envh](https://github.com/antham/envh) - Helpers to manage environment variables. ![](https://img.shields.io/github/last-commit/antham/envh.svg)
* [fig](https://github.com/kkyr/fig) - Tiny library for reading configuration from a file and from environment variables (with validation & defaults). ![](https://img.shields.io/github/last-commit/kkyr/fig.svg)
* [gcfg](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections. ![](https://img.shields.io/github/last-commit/go-gcfg/gcfg.svg)
* [genv](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support. ![](https://img.shields.io/github/last-commit/sakirsensoy/genv.svg)
* [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) - Go package that fetches parameters from AWS System Manager - Parameter Store.  ![](https://img.shields.io/github/last-commit/PaddleHQ/go-aws-ssm.svg)
* [go-ini](https://github.com/subpop/go-ini) - A Go package that marshals and unmarshals INI-files. ![](https://img.shields.io/github/last-commit/subpop/go-ini.svg)
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) - Go utility for loading configuration parameters from AWS SSM (Parameter Store). ![](https://img.shields.io/github/last-commit/ianlopshire/go-ssm-config.svg)
* [go-up](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic. ![](https://img.shields.io/github/last-commit/ufoscout/go-up.svg)
* [goConfig](https://github.com/crgimenes/goConfig) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. ![](https://img.shields.io/github/last-commit/crgimenes/goConfig.svg)
* [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). ![](https://img.shields.io/github/last-commit/joho/godotenv.svg)
* [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy. ![](https://img.shields.io/github/last-commit/ian-kent/gofigure.svg)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gookit/config](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. ![](https://img.shields.io/github/last-commit/gookit/config.svg)
* [harvester](https://github.com/beatlabs/harvester) - Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. ![](https://img.shields.io/github/last-commit/beatlabs/harvester.svg)
* [hjson](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. ![](https://img.shields.io/github/last-commit/hjson/hjson-go.svg)
* [hocon](https://github.com/gurkankaymak/hocon) - Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. ![](https://img.shields.io/github/last-commit/gurkankaymak/hocon.svg)
* [ingo](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file. ![](https://img.shields.io/github/last-commit/schachmat/ingo.svg)
* [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. ![](https://img.shields.io/github/last-commit/go-ini/ini.svg)
* [joshbetz/config](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. ![](https://img.shields.io/github/last-commit/joshbetz/config.svg)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. ![](https://img.shields.io/github/last-commit/kelseyhightower/envconfig.svg)
* [koanf](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. ![](https://img.shields.io/github/last-commit/knadh/koanf.svg)
* [konfig](https://github.com/lalamove/konfig) - Composable, observable and performant config handling for Go for the distributed processing era. ![](https://img.shields.io/github/last-commit/lalamove/konfig.svg)
* [mini](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files. ![](https://img.shields.io/github/last-commit/sasbury/mini.svg)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) - Simple useful package for read environment variables. ![](https://img.shields.io/github/last-commit/nasermirzaei89/env.svg)
* [onion](http://github.com/goraz/onion) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.
* [store](https://github.com/tucnak/store) - Lightweight configuration manager for Go. ![](https://img.shields.io/github/last-commit/tucnak/store.svg)
* [swap](https://github.com/oblq/swap) - Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). ![](https://img.shields.io/github/last-commit/oblq/swap.svg)
* [typenv](https://github.com/diegomarangoni/typenv) - Minimalistic, zero dependency, typed environment variables library. ![](https://img.shields.io/github/last-commit/diegomarangoni/typenv.svg)
* [viper](https://github.com/spf13/viper) - Go configuration with fangs. ![](https://img.shields.io/github/last-commit/spf13/viper.svg)
* [xdg](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). ![](https://img.shields.io/github/last-commit/OpenPeeDeeP/xdg.svg)

## Continuous Integration

*Tools for help with continuous integration.*

* [CDS](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. ![](https://img.shields.io/github/last-commit/ovh/cds.svg)
* [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. ![](https://img.shields.io/github/last-commit/drone/drone.svg)
* [duci](https://github.com/duck8823/duci) - A simple ci server no needs domain specific languages. ![](https://img.shields.io/github/last-commit/duck8823/duci.svg)
* [gomason](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace. ![](https://img.shields.io/github/last-commit/nikogura/gomason.svg)
* [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system. ![](https://img.shields.io/github/last-commit/mattn/goveralls.svg)
* [overalls](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls. ![](https://img.shields.io/github/last-commit/go-playground/overalls.svg)
* [roveralls](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool. ![](https://img.shields.io/github/last-commit/LawrenceWoodman/roveralls.svg)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor. ![](https://img.shields.io/github/last-commit/yosssi/gcss.svg)
* [go-libsass](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project. ![](https://img.shields.io/github/last-commit/wellington/go-libsass.svg)

## Data Structures

*Generic datastructures and algorithms in Go.*

* [algorithms](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study. ![](https://img.shields.io/github/last-commit/shady831213/algorithms.svg)
* [binpacker](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream. ![](https://img.shields.io/github/last-commit/zhuangsirui/binpacker.svg)
* [bit](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions. ![](https://img.shields.io/github/last-commit/yourbasic/bit.svg)
* [bitset](https://github.com/willf/bitset) - Go package implementing bitsets. ![](https://img.shields.io/github/last-commit/willf/bitset.svg)
* [bloom](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. ![](https://img.shields.io/github/last-commit/zhenjl/bloom.svg)
* [bloom](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation. ![](https://img.shields.io/github/last-commit/yourbasic/bloom.svg)
* [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. ![](https://img.shields.io/github/last-commit/tylertreat/BoomFilters.svg)
* [concurrent-writer](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`. ![](https://img.shields.io/github/last-commit/free/concurrent-writer.svg)
* [conjungo](https://github.com/InVisionApp/conjungo) - A small, powerful and flexible merge library. ![](https://img.shields.io/github/last-commit/InVisionApp/conjungo.svg)
* [count-min-log](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). ![](https://img.shields.io/github/last-commit/seiflotfy/count-min-log.svg)
* [crunch](https://github.com/superwhiskers/crunch) - Go package implementing buffers for handling various datatypes easily. ![](https://img.shields.io/github/last-commit/superwhiskers/crunch.svg)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. ![](https://img.shields.io/github/last-commit/seiflotfy/cuckoofilter.svg)
* [deque](https://github.com/edwingeng/deque) - A highly optimized double-ended queue. ![](https://img.shields.io/github/last-commit/edwingeng/deque.svg)
* [deque](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue). ![](https://img.shields.io/github/last-commit/gammazero/deque.svg)
* [dict](https://github.com/srfrog/dict) - Python-like dictionaries (dict) for Go. ![](https://img.shields.io/github/last-commit/srfrog/dict.svg)
* [encoding](https://github.com/zhenjl/encoding) - Integer Compression Libraries for Go. ![](https://img.shields.io/github/last-commit/zhenjl/encoding.svg)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree. ![](https://img.shields.io/github/last-commit/plar/go-adaptive-radix-tree.svg)
* [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. ![](https://img.shields.io/github/last-commit/Workiva/go-datastructures.svg)
* [go-ef](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding. ![](https://img.shields.io/github/last-commit/amallia/go-ef.svg)
* [go-geoindex](https://github.com/hailocab/go-geoindex) - In-memory geo index. ![](https://img.shields.io/github/last-commit/hailocab/go-geoindex.svg)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches. ![](https://img.shields.io/github/last-commit/OrlovEvgeny/go-mcache.svg)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding. ![](https://img.shields.io/github/last-commit/aurelien-rainone/go-rquad.svg)
* [gocache](https://github.com/eko/gocache) - A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. ![](https://img.shields.io/github/last-commit/eko/gocache.svg)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) - Concurrent FIFO queue. ![](https://img.shields.io/github/last-commit/enriquebris/goconcurrentqueue.svg)
* [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. ![](https://img.shields.io/github/last-commit/emirpasic/gods.svg)
* [gofal](https://github.com/xxjwxc/gofal) - fractional api for Go. ![](https://img.shields.io/github/last-commit/xxjwxc/gofal.svg)
* [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. ![](https://img.shields.io/github/last-commit/deckarep/golang-set.svg)
* [goset](https://github.com/zoumo/goset) - A useful Set collection implementation for Go. ![](https://img.shields.io/github/last-commit/zoumo/goset.svg)
* [goskiplist](https://github.com/ryszard/goskiplist) - Skip list implementation in Go. ![](https://img.shields.io/github/last-commit/ryszard/goskiplist.svg)
* [gostl](https://github.com/liyue201/gostl) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL. ![](https://img.shields.io/github/last-commit/liyue201/gostl.svg)
* [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. ![](https://img.shields.io/github/last-commit/kniren/gota.svg)
* [hide](https://github.com/emvi/hide) - ID type with marshalling to/from hash to prevent sending IDs to clients. ![](https://img.shields.io/github/last-commit/emvi/hide.svg)
* [hilbert](https://github.com/google/hilbert) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. ![](https://img.shields.io/github/last-commit/google/hilbert.svg)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. ![](https://img.shields.io/github/last-commit/axiomhq/hyperloglog.svg)
* [iter](https://github.com/disksing/iter) - Go implementation of C++ STL iterators and algorithms. ![](https://img.shields.io/github/last-commit/disksing/iter.svg)
* [levenshtein](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. ![](https://img.shields.io/github/last-commit/agext/levenshtein.svg)
* [levenshtein](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go. ![](https://img.shields.io/github/last-commit/agnivade/levenshtein.svg)
* [mafsa](https://github.com/smartystreets/mafsa) - MA-FSA implementation with Minimal Perfect Hashing. ![](https://img.shields.io/github/last-commit/smartystreets/mafsa.svg)
* [merkletree](https://github.com/cbergoon/merkletree) - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. ![](https://img.shields.io/github/last-commit/cbergoon/merkletree.svg)
* [mspm](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval. ![](https://img.shields.io/github/last-commit/BlackRabbitt/mspm.svg)
* [null](https://github.com/emvi/null) - Nullable Go types that can be marshalled/unmarshalled to/from JSON. ![](https://img.shields.io/github/last-commit/emvi/null.svg)
* [parsefields](https://github.com/MonaxGT/parsefields) - Tools for parse JSON-like logs for collecting unique fields and events. ![](https://img.shields.io/github/last-commit/MonaxGT/parsefields.svg)
* [pipeline](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out. ![](https://img.shields.io/github/last-commit/hyfather/pipeline.svg)
* [ptrie](https://github.com/viant/ptrie) - An implementation of prefix tree. ![](https://img.shields.io/github/last-commit/viant/ptrie.svg)
* [remember-go](https://github.com/rocketlaunchr/remember-go) - A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). ![](https://img.shields.io/github/last-commit/rocketlaunchr/remember-go.svg)
* [ring](https://github.com/TheTannerRyan/ring) - Go implementation of a high performance, thread safe bloom filter. ![](https://img.shields.io/github/last-commit/TheTannerRyan/ring.svg)
* [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. ![](https://img.shields.io/github/last-commit/RoaringBitmap/roaring.svg)
* [set](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap. ![](https://img.shields.io/github/last-commit/StudioSol/set.svg)
* [skiplist](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation. ![](https://img.shields.io/github/last-commit/MauriceGit/skiplist.svg)
* [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go. ![](https://img.shields.io/github/last-commit/gansidui/skiplist.svg)
* [timedmap](https://github.com/zekroTJA/timedmap) - Map with expiring key-value pairs. ![](https://img.shields.io/github/last-commit/zekroTJA/timedmap.svg)
* [treap](https://github.com/perdata/treap) - Persistent, fast ordered map using tree heaps. ![](https://img.shields.io/github/last-commit/perdata/treap.svg)
* [trie](https://github.com/derekparker/trie) - Trie implementation in Go. ![](https://img.shields.io/github/last-commit/derekparker/trie.svg)
* [ttlcache](https://github.com/diegobernardes/ttlcache) - In-memory LRU string-interface{} map with expiration for golang. ![](https://img.shields.io/github/last-commit/diegobernardes/ttlcache.svg)
* [typ](https://github.com/gurukami/typ) - Null Types, Safe primitive type conversion and fetching value from complex structures. ![](https://img.shields.io/github/last-commit/gurukami/typ.svg)
* [willf/bloom](https://github.com/willf/bloom) - Go package implementing Bloom filters. ![](https://img.shields.io/github/last-commit/willf/bloom.svg)

## Database

*Databases implemented in Go.*

* [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. ![](https://img.shields.io/github/last-commit/dgraph-io/badger.svg)
* [bbolt](https://github.com/etcd-io/bbolt) - An embedded key/value database for Go. ![](https://img.shields.io/github/last-commit/etcd-io/bbolt.svg)
* [bcache](https://github.com/iwanbk/bcache) - Eventually consistent distributed in-memory  cache Go library. ![](https://img.shields.io/github/last-commit/iwanbk/bcache.svg)
* [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. ![](https://img.shields.io/github/last-commit/allegro/bigcache.svg)
* [Bitcask](https://github.com/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). ![](https://img.shields.io/github/last-commit/prologic/bitcask.svg)
* [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. ![](https://img.shields.io/github/last-commit/tidwall/buntdb.svg)
* [cache](https://github.com/akyoto/cache) - In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. ![](https://img.shields.io/github/last-commit/akyoto/cache.svg)
* [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. ![](https://img.shields.io/github/last-commit/muesli/cache2go.svg)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration. ![](https://img.shields.io/github/last-commit/oaStuff/clusteredBigCache.svg)
* [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. ![](https://img.shields.io/github/last-commit/cockroachdb/cockroach.svg)
* [Coffer](https://github.com/claygod/coffer) - Simple ACID key-value database that supports transactions. ![](https://img.shields.io/github/last-commit/claygod/coffer.svg)
* [couchcache](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server. ![](https://img.shields.io/github/last-commit/codingsince1985/couchcache.svg)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. ![](https://img.shields.io/github/last-commit/CovenantSQL/CovenantSQL.svg)
* [Databunker](https://github.com/paranoidguy/databunker) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. ![](https://img.shields.io/github/last-commit/paranoidguy/databunker.svg)
* [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. ![](https://img.shields.io/github/last-commit/dgraph-io/dgraph.svg)
* [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. ![](https://img.shields.io/github/last-commit/peterbourgon/diskv.svg)
* [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. ![](https://img.shields.io/github/last-commit/krotik/eliasdb.svg)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. ![](https://img.shields.io/github/last-commit/VictoriaMetrics/fastcache.svg)
* [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. ![](https://img.shields.io/github/last-commit/bluele/gcache.svg)
* [go-cache](https://github.com/pmylund/go-cache) - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. ![](https://img.shields.io/github/last-commit/pmylund/go-cache.svg)
* [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. ![](https://img.shields.io/github/last-commit/syndtr/goleveldb.svg)
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) - Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. ![](https://img.shields.io/github/last-commit/kapitan-k/gorocksdb.svg)
* [gostore](https://github.com/twharmon/gostore) - Gostore is a simple, durable, embedded key-value storage engine written in Go. ![](https://img.shields.io/github/last-commit/twharmon/gostore.svg)
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. ![](https://img.shields.io/github/last-commit/golang/groupcache.svg)
* [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. ![](https://img.shields.io/github/last-commit/influxdb/influxdb.svg)
* [Kivik](https://github.com/go-kivik/kivik) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. ![](https://img.shields.io/github/last-commit/go-kivik/kivik.svg)
* [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. ![](https://img.shields.io/github/last-commit/siddontang/ledisdb.svg)
* [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB. ![](https://img.shields.io/github/last-commit/jmhodges/levigo.svg)
* [moss](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go. ![](https://img.shields.io/github/last-commit/couchbase/moss.svg)
* [nutsdb](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. ![](https://img.shields.io/github/last-commit/xujiajun/nutsdb.svg)
* [piladb](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures. ![](https://img.shields.io/github/last-commit/fern4lvarez/piladb.svg)
* [pogreb](https://github.com/akrylysov/pogreb) - Embedded key-value store for read-heavy workloads. ![](https://img.shields.io/github/last-commit/akrylysov/pogreb.svg)
* [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. ![](https://img.shields.io/github/last-commit/prometheus/prometheus.svg)
* [pudge](https://github.com/recoilme/pudge) - Fast and simple  key/value store written using Go's standard library. ![](https://img.shields.io/github/last-commit/recoilme/pudge.svg)
* [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. ![](https://img.shields.io/github/last-commit/rqlite/rqlite.svg)
* [Scribble](https://github.com/nanobox-io/golang-scribble) - Tiny flat file JSON store. ![](https://img.shields.io/github/last-commit/nanobox-io/golang-scribble.svg)
* [slowpoke](https://github.com/recoilme/slowpoke) - Key-value store with persistence. ![](https://img.shields.io/github/last-commit/recoilme/slowpoke.svg)
* [tempdb](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items. ![](https://img.shields.io/github/last-commit/rafaeljesus/tempdb.svg)
* [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. ![](https://img.shields.io/github/last-commit/pingcap/tidb.svg)
* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. ![](https://img.shields.io/github/last-commit/HouzuoGuo/tiedot.svg)
* [unitdb](https://github.com/unit-io/unitdb) - Fast timeseries database for IoT, realtime messaging  applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. ![](https://img.shields.io/github/last-commit/unit-io/unitdb.svg)
* [Vasto](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. ![](https://img.shields.io/github/last-commit/chrislusf/vasto.svg)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. ![](https://img.shields.io/github/last-commit/VictoriaMetrics/VictoriaMetrics.svg)

*Database schema migration.*

* [avro](https://github.com/khezen/avro) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. ![](https://img.shields.io/github/last-commit/khezen/avro.svg)
* [darwin](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go. ![](https://img.shields.io/github/last-commit/GuiaBolso/darwin.svg)
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library. ![](https://img.shields.io/github/last-commit/RichardKnop/go-fixtures.svg)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg. ![](https://img.shields.io/github/last-commit/robinjoseph08/go-pg-migrations.svg)
* [gondolier](https://github.com/emvi/gondolier) - Database migration library using struct decorators. ![](https://img.shields.io/github/last-commit/emvi/gondolier.svg)
* [goose](https://github.com/steinbacher/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. ![](https://img.shields.io/github/last-commit/steinbacher/goose.svg)
* [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. ![](https://img.shields.io/github/last-commit/go-gormigrate/gormigrate.svg)
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. ![](https://img.shields.io/github/last-commit/golang-migrate/migrate.svg)
* [migrator](https://github.com/lopezator/migrator) - Dead simple Go database migration library. ![](https://img.shields.io/github/last-commit/lopezator/migrator.svg)
* [pravasan](https://github.com/pravasan/pravasan) - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. ![](https://img.shields.io/github/last-commit/pravasan/pravasan.svg)
* [schema](https://github.com/adlio/schema) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. ![](https://img.shields.io/github/last-commit/adlio/schema.svg)
* [skeema](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. ![](https://img.shields.io/github/last-commit/skeema/skeema.svg)
* [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. ![](https://img.shields.io/github/last-commit/rubenv/sql-migrate.svg)

*Database tools.*

* [bucket](https://github.com/PumpkinSeed/bucket) - Optimized data structure framework for Couchbase specialized on one bucket usage. ![](https://img.shields.io/github/last-commit/PumpkinSeed/bucket.svg)
* [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. ![](https://img.shields.io/github/last-commit/Vertamedia/chproxy.svg)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - Collects small insterts and sends big requests to ClickHouse servers. ![](https://img.shields.io/github/last-commit/nikepan/clickhouse-bulk.svg)
* [datagen](https://github.com/codingconcepts/datagen) - A fast data generator that's multi-table aware and supports multi-row DML. ![](https://img.shields.io/github/last-commit/codingconcepts/datagen.svg)
* [dbbench](https://github.com/sj14/dbbench) - Database benchmarking tool with support for several databases and scripts. ![](https://img.shields.io/github/last-commit/sj14/dbbench.svg)
* [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. ![](https://img.shields.io/github/last-commit/siddontang/go-mysql.svg)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically. ![](https://img.shields.io/github/last-commit/siddontang/go-mysql-elasticsearch.svg)
* [kingshard](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang. ![](https://img.shields.io/github/last-commit/flike/kingshard.svg)
* [myreplication](https://github.com/2tvenom/myreplication) - MySql binary log replication listener. Supports statement and row based replication. ![](https://img.shields.io/github/last-commit/2tvenom/myreplication.svg)
* [octillery](https://github.com/knocknote/octillery) - Go package for sharding databases ( Supports every ORM or raw SQL ). ![](https://img.shields.io/github/last-commit/knocknote/octillery.svg)
* [orchestrator](https://github.com/github/orchestrator) - MySQL replication topology manager & visualizer. ![](https://img.shields.io/github/last-commit/github/orchestrator.svg)
* [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. ![](https://img.shields.io/github/last-commit/sosedoff/pgweb.svg)
* [prep](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code. ![](https://img.shields.io/github/last-commit/hexdigest/prep.svg)
* [pREST](https://github.com/nuveo/prest) - Serve a RESTful API from any PostgreSQL database. ![](https://img.shields.io/github/last-commit/nuveo/prest.svg)
* [rwdb](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup. ![](https://img.shields.io/github/last-commit/andizzle/rwdb.svg)
* [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. ![](https://img.shields.io/github/last-commit/youtube/vitess.svg)

*SQL query builder, libraries for building and using SQL.*

* [buildsqlx](https://github.com/arthurkushman/buildsqlx) - Go database query builder library for PostgreSQL. ![](https://img.shields.io/github/last-commit/arthurkushman/buildsqlx.svg)
* [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go. ![](https://img.shields.io/github/last-commit/rocketlaunchr/dbq.svg)
* [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease. ![](https://img.shields.io/github/last-commit/gchaincl/dotsql.svg)
* [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. ![](https://img.shields.io/github/last-commit/didi/gendry.svg)
* [godbal](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. ![](https://img.shields.io/github/last-commit/xujiajun/godbal.svg)
* [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. ![](https://img.shields.io/github/last-commit/doug-martin/goqu.svg)
* [gosql](https://github.com/twharmon/gosql) - SQL Query builder with better null values support. ![](https://img.shields.io/github/last-commit/twharmon/gosql.svg)
* [igor](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. ![](https://img.shields.io/github/last-commit/galeone/igor.svg)
* [jet](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. ![](https://img.shields.io/github/last-commit/go-jet/jet.svg)
* [mpath](https://github.com/spacetab-io/mpath-go) - MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation. ![](https://img.shields.io/github/last-commit/spacetab-io/mpath-go.svg)
* [ormlite](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases. ![](https://img.shields.io/github/last-commit/pupizoid/ormlite.svg)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities. ![](https://img.shields.io/github/last-commit/go-ozzo/ozzo-dbx.svg)
* [qry](https://github.com/HnH/qry) - Tool that generates constants from files with raw SQL queries. ![](https://img.shields.io/github/last-commit/HnH/qry.svg)
* [sqlf](https://github.com/leporo/sqlf) - Fast SQL query builder. ![](https://img.shields.io/github/last-commit/leporo/sqlf.svg)
* [sqrl](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance. ![](https://img.shields.io/github/last-commit/elgris/sqrl.svg)
* [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.
* [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. ![](https://img.shields.io/github/last-commit/Masterminds/squirrel.svg)
* [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. ![](https://img.shields.io/github/last-commit/knq/xo.svg)

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [avatica](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql. ![](https://img.shields.io/github/last-commit/apache/calcite-avatica-go.svg)
    * [bgc](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go. ![](https://img.shields.io/github/last-commit/viant/bgc.svg)
    * [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go. ![](https://img.shields.io/github/last-commit/nakagami/firebirdsql.svg)
    * [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql. ![](https://img.shields.io/github/last-commit/mattn/go-adodb.svg)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. ![](https://img.shields.io/github/last-commit/denisenkom/go-mssqldb.svg)
    * [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql. ![](https://img.shields.io/github/last-commit/mattn/go-oci8.svg)
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. ![](https://img.shields.io/github/last-commit/go-sql-driver/mysql.svg)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. ![](https://img.shields.io/github/last-commit/mattn/go-sqlite3.svg)
    * [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). ![](https://img.shields.io/github/last-commit/minus5/gofreetds.svg)
    * [goracle](https://github.com/go-goracle/goracle) - Oracle driver for Go, using the ODPI-C driver. ![](https://img.shields.io/github/last-commit/go-goracle/goracle.svg)
    * [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. ![](https://img.shields.io/github/last-commit/jackc/pgx.svg)
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. ![](https://img.shields.io/github/last-commit/lib/pq.svg)
    * [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) - SQLite with pure Go. ![](https://img.shields.io/github/last-commit/cvilsmeier/sqinn-go.svg)

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language. ![](https://img.shields.io/github/last-commit/aerospike/aerospike-client-go.svg)
    * [arangolite](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB. ![](https://img.shields.io/github/last-commit/solher/arangolite.svg)
    * [asc](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go. ![](https://img.shields.io/github/last-commit/viant/asc.svg)
    * [dynago](https://github.com/underarmour/dynago) - Dynago is a principle of least surprise client for DynamoDB. ![](https://img.shields.io/github/last-commit/underarmour/dynago.svg)
    * [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB. ![](https://img.shields.io/github/last-commit/couchbase/goforestdb.svg)
    * [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go. ![](https://img.shields.io/github/last-commit/couchbase/go-couchbase.svg)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) - Go client library for Pilosa. ![](https://img.shields.io/github/last-commit/pilosa/go-pilosa.svg)
    * [go-rejson](https://github.com/nitishm/go-rejson) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. ![](https://img.shields.io/github/last-commit/nitishm/go-rejson.svg)
    * [gocb](https://github.com/couchbase/gocb) - Official Couchbase Go SDK. ![](https://img.shields.io/github/last-commit/couchbase/gocb.svg)
    * [gocql](http://gocql.github.io) - Go language driver for Apache Cassandra.
    * [godis](https://github.com/piaohao/godis) - redis client implement by golang, inspired by jedis. ![](https://img.shields.io/github/last-commit/piaohao/godis.svg)
    * [godscache](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. ![](https://img.shields.io/github/last-commit/defcronyke/godscache.svg)
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
    * [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. ![](https://img.shields.io/github/last-commit/dancannon/gorethink.svg)
    * [goriak](https://github.com/zegl/goriak) - Go language driver for Riak KV. ![](https://img.shields.io/github/last-commit/zegl/goriak.svg)
    * [mgm](https://github.com/kamva/mgm) - MongoDB model-based ODM for Go (based on official MongoDB driver). ![](https://img.shields.io/github/last-commit/kamva/mgm.svg)
    * [mgo](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. ![](https://img.shields.io/github/last-commit/globalsign/mgo.svg)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. ![](https://img.shields.io/github/last-commit/mongodb/mongo-go-driver.svg)
    * [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang. ![](https://img.shields.io/github/last-commit/cihangir/neo4j.svg)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang. ![](https://img.shields.io/github/last-commit/davemeehan/Neo4j-GO.svg)
    * [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang. ![](https://img.shields.io/github/last-commit/jmcvetta/neoism.svg)
    * [redeo](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services. ![](https://img.shields.io/github/last-commit/bsm/redeo.svg)
    * [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. ![](https://img.shields.io/github/last-commit/gomodule/redigo.svg)
    * [redis](https://github.com/go-redis/redis) - Redis client for Golang. ![](https://img.shields.io/github/last-commit/go-redis/redis.svg)
    * [xredis](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client. ![](https://img.shields.io/github/last-commit/shomali11/xredis.svg)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. ![](https://img.shields.io/github/last-commit/blevesearch/bleve.svg)
    * [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. ![](https://img.shields.io/github/last-commit/olivere/elastic.svg)
    * [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. ![](https://img.shields.io/github/last-commit/cch123/elasticsql.svg)
    * [elastigo](https://github.com/mattbaird/elastigo) - Elasticsearch client library. ![](https://img.shields.io/github/last-commit/mattbaird/elastigo.svg)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go. ![](https://img.shields.io/github/last-commit/elastic/go-elasticsearch.svg)
    * [goes](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch. ![](https://img.shields.io/github/last-commit/OwnLocal/goes.svg)
    * [riot](https://github.com/go-ego/riot) - Go Open Source, Distributed, Simple and efficient Search Engine. ![](https://img.shields.io/github/last-commit/go-ego/riot.svg)
    * [skizze](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage. ![](https://img.shields.io/github/last-commit/seiflotfy/skizze.svg)

* Multiple Backends.
    * [cachego](https://github.com/fabiorphp/cachego) - Golang Cache component for multiple drivers. ![](https://img.shields.io/github/last-commit/fabiorphp/cachego.svg)
    * [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. ![](https://img.shields.io/github/last-commit/google/cayley.svg)
    * [dsc](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files. ![](https://img.shields.io/github/last-commit/viant/dsc.svg)
    * [gokv](https://github.com/philippgille/gokv) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). ![](https://img.shields.io/github/last-commit/philippgille/gokv.svg)

## Date and Time

*Libraries for working with dates and times.*

* [carbon](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library. ![](https://img.shields.io/github/last-commit/uniplaces/carbon.svg)
* [cronrange](https://github.com/1set/cronrange) - Parses Cron-style time range expressions, checks if the given time is within any ranges. ![](https://img.shields.io/github/last-commit/1set/cronrange.svg)
* [date](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. ![](https://img.shields.io/github/last-commit/rickb777/date.svg)
* [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. ![](https://img.shields.io/github/last-commit/araddon/dateparse.svg)
* [durafmt](https://github.com/hako/durafmt) - Time duration formatting library for Go. ![](https://img.shields.io/github/last-commit/hako/durafmt.svg)
* [feiertage](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... ![](https://img.shields.io/github/last-commit/wlbr/feiertage.svg)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang). ![](https://img.shields.io/github/last-commit/yaa110/go-persian-calendar.svg)
* [go-str2duration](https://github.com/xhit/go-str2duration) - Convert string to duration. Support time.Duration returned string and more. ![](https://img.shields.io/github/last-commit/xhit/go-str2duration.svg)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location. ![](https://img.shields.io/github/last-commit/nathan-osman/go-sunrise.svg)
* [go-week](https://github.com/stoewer/go-week) - An efficient package to work with ISO8601 week dates. ![](https://img.shields.io/github/last-commit/stoewer/go-week.svg)
* [iso8601](https://github.com/relvacode/iso8601) - Efficiently parse ISO8601 date-times without regex. ![](https://img.shields.io/github/last-commit/relvacode/iso8601.svg)
* [kair](https://github.com/GuilhermeCaruso/kair) - Date and Time - Golang Formatting Library. ![](https://img.shields.io/github/last-commit/GuilhermeCaruso/kair.svg)
* [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. ![](https://img.shields.io/github/last-commit/jinzhu/now.svg)
* [NullTime](https://github.com/kirillDanshin/nulltime) - Nullable `time.Time`. ![](https://img.shields.io/github/last-commit/kirillDanshin/nulltime.svg)
* [strftime](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter. ![](https://img.shields.io/github/last-commit/awoodbeck/strftime.svg)
* [timespan](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration. ![](https://img.shields.io/github/last-commit/SaidinWoT/timespan.svg)
* [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package. ![](https://img.shields.io/github/last-commit/leekchan/timeutil.svg)
* [tuesday](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function. ![](https://img.shields.io/github/last-commit/osteele/tuesday.svg)

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. ![](https://img.shields.io/github/last-commit/svcavallar/celeriac.v1.svg)
* [consistent](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads. ![](https://img.shields.io/github/last-commit/buraksezer/consistent.svg)
* [consistenthash](https://github.com/mbrostami/consistenthash) - Consistent hashing with configurable replicas. ![](https://img.shields.io/github/last-commit/mbrostami/consistenthash.svg)
* [dht](https://github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation. ![](https://img.shields.io/github/last-commit/anacrolix/dht.svg)
* [digota](https://github.com/digota/digota) - grpc ecommerce microservice. ![](https://img.shields.io/github/last-commit/digota/digota.svg)
* [dot](https://github.com/dotchain/dot/) - distributed sync using operational transformation/OT.
* [doublejump](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash. ![](https://img.shields.io/github/last-commit/edwingeng/doublejump.svg)
* [dragonboat](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go. ![](https://img.shields.io/github/last-commit/lni/dragonboat.svg)
* [drmaa](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard. ![](https://img.shields.io/github/last-commit/dgruber/drmaa.svg)
* [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
* [dynatomic](https://github.com/tylfin/dynatomic) - A library for using DynamoDB as an atomic counter. ![](https://img.shields.io/github/last-commit/tylfin/dynatomic.svg)
* [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. ![](https://img.shields.io/github/last-commit/emitter-io/emitter.svg)
* [flowgraph](https://github.com/vectaport/flowgraph) - flow-based programming package. ![](https://img.shields.io/github/last-commit/vectaport/flowgraph.svg)
* [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. ![](https://img.shields.io/github/last-commit/chrislusf/gleam.svg)
* [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. ![](https://img.shields.io/github/last-commit/chrislusf/glow.svg)
* [gmsec](https://github.com/gmsec/micro) - A Go distributed systems development framework. ![](https://img.shields.io/github/last-commit/gmsec/micro.svg)
* [go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service. ![](https://img.shields.io/github/last-commit/InVisionApp/go-health.svg)
* [go-jump](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function. ![](https://img.shields.io/github/last-commit/dgryski/go-jump.svg)
* [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. ![](https://img.shields.io/github/last-commit/go-kit/kit.svg)
* [go-micro](https://github.com/micro/go-micro) - A distributed systems development framework. ![](https://img.shields.io/github/last-commit/micro/go-micro.svg)
* [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) - MySQL based distributed lock. ![](https://img.shields.io/github/last-commit/sanketplus/go-mysql-lock.svg)
* [go-pdu](https://github.com/pdupub/go-pdu) - A decentralized identity-based social network. ![](https://img.shields.io/github/last-commit/pdupub/go-pdu.svg)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) - A library built to provide support for defining async service health checks for golang services. ![](https://img.shields.io/github/last-commit/AppsFlyer/go-sundheit.svg)
* [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load. ![](https://img.shields.io/github/last-commit/valyala/gorpc.svg)
* [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. ![](https://img.shields.io/github/last-commit/grpc/grpc-go.svg)
* [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. ![](https://img.shields.io/github/last-commit/hprose/hprose-golang.svg)
* [jsonrpc](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0. ![](https://img.shields.io/github/last-commit/osamingo/jsonrpc.svg)
* [jsonrpc](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation. ![](https://img.shields.io/github/last-commit/ybbus/jsonrpc.svg)
* [KrakenD](https://github.com/devopsfaith/krakend) - Ultra performant API Gateway framework with middlewares. ![](https://img.shields.io/github/last-commit/devopsfaith/krakend.svg)
* [liftbridge](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS. ![](https://img.shields.io/github/last-commit/liftbridge-io/liftbridge.svg)
* [micro](https://github.com/micro/micro) - A distributed systems runtime for the cloud and beyond. ![](https://img.shields.io/github/last-commit/micro/micro.svg)
* [NATS](https://github.com/nats-io/gnatsd) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. ![](https://img.shields.io/github/last-commit/nats-io/gnatsd.svg)
* [outboxer](https://github.com/italolelis/outboxer) - Outboxer is a go library that implements the outbox pattern. ![](https://img.shields.io/github/last-commit/italolelis/outboxer.svg)
* [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.
* [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. ![](https://img.shields.io/github/last-commit/hashicorp/raft.svg)
* [raft](https://github.com/coreos/etcd/tree/master/raft) - Go implementation of the Raft consensus protocol, by CoreOS.
* [rain](https://github.com/cenkalti/rain) - BitTorrent client and library. ![](https://img.shields.io/github/last-commit/cenkalti/rain.svg)
* [redis-lock](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis. ![](https://img.shields.io/github/last-commit/bsm/redislock.svg)
* [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [ringpop-go](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications. ![](https://img.shields.io/github/last-commit/uber/ringpop-go.svg)
* [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. ![](https://img.shields.io/github/last-commit/smallnest/rpcx.svg)
* [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). ![](https://img.shields.io/github/last-commit/ursiform/sleuth.svg)
* [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. ![](https://img.shields.io/github/last-commit/tendermint/tendermint.svg)
* [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. ![](https://img.shields.io/github/last-commit/anacrolix/torrent.svg)

## Dynamic DNS

*Tools for updating dynamic DNS records.*

* [DDNS](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend. ![](https://img.shields.io/github/last-commit/skibish/ddns.svg)
* [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.
* [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. ![](https://img.shields.io/github/last-commit/timothyye/godns.svg)

## Email

*Libraries and tools that implement email creation and sending.*

* [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.
* [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails. ![](https://img.shields.io/github/last-commit/aymerick/douceur.svg)
* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. ![](https://img.shields.io/github/last-commit/jordan-wright/email.svg)
* [go-dkim](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email. ![](https://img.shields.io/github/last-commit/toorop/go-dkim.svg)
* [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. ![](https://img.shields.io/github/last-commit/emersion/go-imap.svg)
* [go-message](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages. ![](https://img.shields.io/github/last-commit/emersion/go-message.svg)
* [go-premailer](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in Go. ![](https://img.shields.io/github/last-commit/vanng822/go-premailer.svg)
* [go-simple-mail](https://github.com/xhit/go-simple-mail) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. ![](https://img.shields.io/github/last-commit/xhit/go-simple-mail.svg)
* [Hectane](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API. ![](https://img.shields.io/github/last-commit/hectane/hectane.svg)
* [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. ![](https://img.shields.io/github/last-commit/matcornic/hermes.svg)
* [mailchain](https://github.com/mailchain/mailchain) - Send encrypted emails to blockchain addresses written in Go. ![](https://img.shields.io/github/last-commit/mailchain/mailchain.svg)
* [mailgun-go](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API. ![](https://img.shields.io/github/last-commit/mailgun/mailgun-go.svg)
* [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. ![](https://img.shields.io/github/last-commit/mailhog/MailHog.svg)
* [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email. ![](https://img.shields.io/github/last-commit/sendgrid/sendgrid-go.svg)
* [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine. ![](https://img.shields.io/github/last-commit/mailhog/smtp.svg)

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. ![](https://img.shields.io/github/last-commit/mattn/anko.svg)
* [binder](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). ![](https://img.shields.io/github/last-commit/alexeyco/binder.svg)
* [cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing. ![](https://img.shields.io/github/last-commit/google/cel-go.svg)
* [expr](https://github.com/antonmedv/expr) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. ![](https://img.shields.io/github/last-commit/antonmedv/expr.svg)
* [gentee](https://github.com/gentee/gentee) - Embeddable scripting programming language. ![](https://img.shields.io/github/last-commit/gentee/gentee.svg)
* [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go. ![](https://img.shields.io/github/last-commit/jcla1/gisp.svg)
* [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go. ![](https://img.shields.io/github/last-commit/olebedev/go-duktape.svg)
* [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. ![](https://img.shields.io/github/last-commit/Shopify/go-lua.svg)
* [go-php](https://github.com/deuill/go-php) - PHP bindings for Go. ![](https://img.shields.io/github/last-commit/deuill/go-php.svg)
* [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. ![](https://img.shields.io/github/last-commit/sbinet/go-python.svg)
* [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API. ![](https://img.shields.io/github/last-commit/aarzilli/golua.svg)
* [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. ![](https://img.shields.io/github/last-commit/yuin/gopher-lua.svg)
* [gval](https://github.com/PaesslerAG/gval) - A highly customizable expression language written in Go. ![](https://img.shields.io/github/last-commit/PaesslerAG/gval.svg)
* [ngaro](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro. ![](https://img.shields.io/github/last-commit/db47h/ngaro.svg)
* [otto](https://github.com/robertkrimen/otto) - JavaScript interpreter written in Go. ![](https://img.shields.io/github/last-commit/robertkrimen/otto.svg)
* [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go. ![](https://img.shields.io/github/last-commit/ian-kent/purl.svg)
* [tengo](https://github.com/d5/tengo) - Bytecode compiled script language for Go. ![](https://img.shields.io/github/last-commit/d5/tengo.svg)

## Error Handling

*Libraries for handling errors.*

* [emperror](https://github.com/emperror/emperror) - Error handling tools and best practices for Go libraries and applications. ![](https://img.shields.io/github/last-commit/emperror/emperror.svg)
* [eris](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. ![](https://img.shields.io/github/last-commit/rotisserie/eris.svg)
* [errlog](https://github.com/snwfdhmp/errlog) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. ![](https://img.shields.io/github/last-commit/snwfdhmp/errlog.svg)
* [errors](https://github.com/emperror/errors) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. ![](https://img.shields.io/github/last-commit/emperror/errors.svg)
* [errors](https://github.com/pkg/errors) - Package that provides simple error handling primitives. ![](https://img.shields.io/github/last-commit/pkg/errors.svg)
* [errors](https://github.com/neuronlabs/errors) - Simple golang error handling with classification primitives. ![](https://img.shields.io/github/last-commit/neuronlabs/errors.svg)
* [errors](https://github.com/PumpkinSeed/errors) - The most simple error wrapper with awesome performance and minimal memory overhead. ![](https://img.shields.io/github/last-commit/PumpkinSeed/errors.svg)
* [errorx](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more. ![](https://img.shields.io/github/last-commit/joomcode/errorx.svg)
* [Falcon](https://github.com/SonicRoshan/falcon) - A Simple Yet Highly Powerful Package For Error Handling. ![](https://img.shields.io/github/last-commit/SonicRoshan/falcon.svg)
* [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. ![](https://img.shields.io/github/last-commit/hashicorp/go-multierror.svg)
* [tracerr](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments. ![](https://img.shields.io/github/last-commit/ztrue/tracerr.svg)
* [werr](https://github.com/txgruppi/werr) - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. ![](https://img.shields.io/github/last-commit/txgruppi/werr.svg)

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. ![](https://img.shields.io/github/last-commit/spf13/afero.svg)
* [afs](https://github.com/viant/afs) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. ![](https://img.shields.io/github/last-commit/viant/afs.svg)
* [bigfile](https://github.com/bigfile/bigfile) - A file transfer system, support to manage files with http api, rpc call and ftp client. ![](https://img.shields.io/github/last-commit/bigfile/bigfile.svg)
* [checksum](https://github.com/codingsince1985/checksum) - Compute message digest, like MD5 and SHA256, for large files. ![](https://img.shields.io/github/last-commit/codingsince1985/checksum.svg)
* [copy](https://github.com/otiai10/copy) - Copy directory recursively. ![](https://img.shields.io/github/last-commit/otiai10/copy.svg)
* [flop](https://github.com/homedepot/flop) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). ![](https://img.shields.io/github/last-commit/homedepot/flop.svg)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) - Load csv file using tag. ![](https://img.shields.io/github/last-commit/artonge/go-csv-tag.svg)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans. ![](https://img.shields.io/github/last-commit/hugocarreira/go-decent-copy.svg)
* [go-exiftool](https://github.com/barasher/go-exiftool) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). ![](https://img.shields.io/github/last-commit/barasher/go-exiftool.svg)
* [go-gtfs](https://github.com/artonge/go-gtfs) - Load gtfs files in go. ![](https://img.shields.io/github/last-commit/artonge/go-gtfs.svg)
* [gut/yos](https://github.com/1set/gut) - Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. ![](https://img.shields.io/github/last-commit/1set/gut.svg)
* [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal. ![](https://img.shields.io/github/last-commit/rjeczalik/notify.svg)
* [opc](https://github.com/qmuntal/opc) - Load Open Packaging Conventions (OPC) files for Go. ![](https://img.shields.io/github/last-commit/qmuntal/opc.svg)
* [parquet](https://github.com/parsyl/parquet) - Read and write [parquet](https://parquet.apache.org) files. ![](https://img.shields.io/github/last-commit/parsyl/parquet.svg)
* [pdfcpu](https://github.com/pdfcpu/pdfcpu) - PDF processor. ![](https://img.shields.io/github/last-commit/pdfcpu/pdfcpu.svg)
* [skywalker](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease. ![](https://img.shields.io/github/last-commit/dixonwille/skywalker.svg)
* [stl](https://gitlab.com/russoj88/stl) - Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [tarfs](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. ![](https://img.shields.io/github/last-commit/posener/tarfs.svg)
* [vfs](https://github.com/C2FO/vfs) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. ![](https://img.shields.io/github/last-commit/C2FO/vfs.svg)

## Financial

*Packages for accounting and finance.*

* [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang. ![](https://img.shields.io/github/last-commit/leekchan/accounting.svg)
* [currency](https://github.com/bnkamalesh/currency) - High performant & accurate currency computation package. ![](https://img.shields.io/github/last-commit/bnkamalesh/currency.svg)
* [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. ![](https://img.shields.io/github/last-commit/shopspring/decimal.svg)
* [go-finance](https://github.com/FlashBoys/go-finance) - Comprehensive financial markets data in Go. ![](https://img.shields.io/github/last-commit/FlashBoys/go-finance.svg)
* [go-finance](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. ![](https://img.shields.io/github/last-commit/alpeb/go-finance.svg)
* [go-finance](https://github.com/pieterclaerhout/go-finance) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. ![](https://img.shields.io/github/last-commit/pieterclaerhout/go-finance.svg)
* [go-finnhub](https://github.com/m1/go-finnhub) - Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. ![](https://img.shields.io/github/last-commit/m1/go-finnhub.svg)
* [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. ![](https://img.shields.io/github/last-commit/rhymond/go-money.svg)
* [ofxgo](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client). ![](https://img.shields.io/github/last-commit/aclindsa/ofxgo.svg)
* [orderbook](https://github.com/i25959341/orderbook) - Matching Engine for Limit Order Book in Golang. ![](https://img.shields.io/github/last-commit/i25959341/orderbook.svg)
* [techan](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies. ![](https://img.shields.io/github/last-commit/sdcoffey/techan.svg)
* [transaction](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode. ![](https://img.shields.io/github/last-commit/claygod/transaction.svg)
* [vat](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates. ![](https://img.shields.io/github/last-commit/dannyvankooten/vat.svg)

## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind) - Bind form data to any Go values. ![](https://img.shields.io/github/last-commit/robfig/bind.svg)
* [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct. ![](https://img.shields.io/github/last-commit/mholt/binding.svg)
* [conform](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. ![](https://img.shields.io/github/last-commit/leebenson/conform.svg)
* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. ![](https://img.shields.io/github/last-commit/go-playground/form.svg)
* [formam](https://github.com/monoculum/formam) - decode form's values into a struct. ![](https://img.shields.io/github/last-commit/monoculum/formam.svg)
* [forms](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. ![](https://img.shields.io/github/last-commit/albrow/forms.svg)
* [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. ![](https://img.shields.io/github/last-commit/gorilla/csrf.svg)
* [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. ![](https://img.shields.io/github/last-commit/justinas/nosurf.svg)
* [queryparam](https://github.com/tomwright/queryparam) - Decode `url.Values` into usable struct values of standard or custom types. ![](https://img.shields.io/github/last-commit/tomwright/queryparam.svg)

## Functional

*Packages to support functional programming in Go.*

* [fpGo](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang. ![](https://img.shields.io/github/last-commit/TeaEntityLab/fpGo.svg)
* [fuego](https://github.com/seborama/fuego) - Functional Experiment in Go. ![](https://img.shields.io/github/last-commit/seborama/fuego.svg)
* [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. ![](https://img.shields.io/github/last-commit/tobyhede/go-underscore.svg)

## Game Development

*Awesome game development libraries.*

* [Azul3D](https://github.com/azul3d/engine) - 3D game engine written in Go. ![](https://img.shields.io/github/last-commit/azul3d/engine.svg)
* [Ebiten](https://github.com/hajimehoshi/ebiten) - dead simple 2D game library in Go. ![](https://img.shields.io/github/last-commit/hajimehoshi/ebiten.svg)
* [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. ![](https://img.shields.io/github/last-commit/EngoEngine/engo.svg)
* [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. ![](https://img.shields.io/github/last-commit/g3n/engine.svg)
* [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm. ![](https://img.shields.io/github/last-commit/beefsack/go-astar.svg)
* [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). ![](https://img.shields.io/github/last-commit/veandco/go-sdl2.svg)
* [go3d](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go. ![](https://img.shields.io/github/last-commit/ungerik/go3d.svg)
* [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. ![](https://img.shields.io/github/last-commit/xtaci/gonet.svg)
* [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping. ![](https://img.shields.io/github/last-commit/xiaonanln/goworld.svg)
* [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. ![](https://img.shields.io/github/last-commit/name5566/leaf.svg)
* [nano](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework. ![](https://img.shields.io/github/last-commit/lonng/nano.svg)
* [Oak](https://github.com/oakmound/oak) - Pure Go game engine. ![](https://img.shields.io/github/last-commit/oakmound/oak.svg)
* [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. ![](https://img.shields.io/github/last-commit/topfreegames/pitaya.svg)
* [Pixel](https://github.com/faiface/pixel) - Hand-crafted 2D game library in Go. ![](https://img.shields.io/github/last-commit/faiface/pixel.svg)
* [prototype](https://github.com/gonutz/prototype) - Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. ![](https://img.shields.io/github/last-commit/gonutz/prototype.svg)
* [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. ![](https://img.shields.io/github/last-commit/gen2brain/raylib-go.svg)
* [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. ![](https://img.shields.io/github/last-commit/JoelOtter/termloop.svg)

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [efaceconv](https://github.com/t0pep0/efaceconv) - Code generation tool for high performance conversion from interface{} to immutable type without allocations. ![](https://img.shields.io/github/last-commit/t0pep0/efaceconv.svg)
* [gen](https://github.com/clipperhouse/gen) - Code generation tool for ‘generics’-like functionality. ![](https://img.shields.io/github/last-commit/clipperhouse/gen.svg)
* [generis](https://github.com/senselogic/GENERIS) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. ![](https://img.shields.io/github/last-commit/senselogic/GENERIS.svg)
* [go-enum](https://github.com/abice/go-enum) - Code generation for enums from code comments. ![](https://img.shields.io/github/last-commit/abice/go-enum.svg)
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. ![](https://img.shields.io/github/last-commit/ahmetalpbalkan/go-linq.svg)
* [go-xray](https://github.com/pieterclaerhout/go-xray) - Helpers for making the use of reflection easier. ![](https://img.shields.io/github/last-commit/pieterclaerhout/go-xray.svg)
* [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types. ![](https://img.shields.io/github/last-commit/awalterschulze/goderive.svg)
* [gotype](https://github.com/wzshiming/gotype) - Golang source code parsing, usage like reflect package. ![](https://img.shields.io/github/last-commit/wzshiming/gotype.svg)
* [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. ![](https://img.shields.io/github/last-commit/hexdigest/gowrap.svg)
* [interfaces](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions. ![](https://img.shields.io/github/last-commit/rjeczalik/interfaces.svg)
* [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. ![](https://img.shields.io/github/last-commit/dave/jennifer.svg)
* [pkgreflect](https://github.com/ungerik/pkgreflect) - Go preprocessor for package scoped reflection. ![](https://img.shields.io/github/last-commit/ungerik/pkgreflect.svg)
* [typeregistry](https://github.com/xiaoxin01/typeregistry) - A library to create type dynamically. ![](https://img.shields.io/github/last-commit/xiaoxin01/typeregistry.svg)

## Geographic

*Geographic tools and servers*

* [geocache](https://github.com/melihmucuk/geocache) - In-memory cache that is suitable for geolocation based applications. ![](https://img.shields.io/github/last-commit/melihmucuk/geocache.svg)
* [geoserver](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. ![](https://img.shields.io/github/last-commit/hishamkaram/geoserver.svg)
* [gismanager](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. ![](https://img.shields.io/github/last-commit/hishamkaram/gismanager.svg)
* [mbtileserver](https://github.com/consbio/mbtileserver) - A simple Go-based server for map tiles stored in mbtiles format. ![](https://img.shields.io/github/last-commit/consbio/mbtileserver.svg)
* [osm](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs. ![](https://img.shields.io/github/last-commit/paulmach/osm.svg)
* [pbf](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder. ![](https://img.shields.io/github/last-commit/maguro/pbf.svg)
* [S2 geojson](https://github.com/pantrif/s2-geojson) - Convert geojson to s2 cells & demonstrating some S2 geometry features on map. ![](https://img.shields.io/github/last-commit/pantrif/s2-geojson.svg)
* [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go. ![](https://img.shields.io/github/last-commit/golang/geo.svg)
* [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. ![](https://img.shields.io/github/last-commit/tidwall/tile38.svg)
* [WGS84](https://github.com/wroge/wgs84) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). ![](https://img.shields.io/github/last-commit/wroge/wgs84.svg)

## Go Compilers

*Tools for compiling Go to other languages.*

* [c4go](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code. ![](https://img.shields.io/github/last-commit/Konstantin8105/c4go.svg)
* [f4go](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code. ![](https://img.shields.io/github/last-commit/Konstantin8105/f4go.svg)
* [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. ![](https://img.shields.io/github/last-commit/gopherjs/gopherjs.svg)
* [llgo](https://github.com/go-llvm/llgo) - LLVM-based compiler for Go. ![](https://img.shields.io/github/last-commit/go-llvm/llgo.svg)
* [tardisgo](https://github.com/tardisgo/tardisgo) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. ![](https://img.shields.io/github/last-commit/tardisgo/tardisgo.svg)

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go. ![](https://img.shields.io/github/last-commit/panjf2000/ants.svg)
* [artifex](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching. ![](https://img.shields.io/github/last-commit/borderstech/artifex.svg)
* [async](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic. ![](https://img.shields.io/github/last-commit/studiosol/async.svg)
* [breaker](https://github.com/kamilsk/breaker) - Flexible mechanism to make execution flow interruptible. ![](https://img.shields.io/github/last-commit/kamilsk/breaker.svg)
* [conexec](https://github.com/ITcathyh/conexec) - A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. ![](https://img.shields.io/github/last-commit/ITcathyh/conexec.svg)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang. ![](https://img.shields.io/github/last-commit/marusama/cyclicbarrier.svg)
* [go-floc](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease. ![](https://img.shields.io/github/last-commit/workanator/go-floc.svg)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order. ![](https://img.shields.io/github/last-commit/kamildrazkiewicz/go-flow.svg)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) - Manage a pool of goroutines using this lightweight library with a simple API. ![](https://img.shields.io/github/last-commit/nikhilsaraf/go-tools.svg)
* [go-trylock](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang. ![](https://img.shields.io/github/last-commit/subchen/go-trylock.svg)
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) - Like `sync.WaitGroup` with error handling and concurrency control. ![](https://img.shields.io/github/last-commit/pieterclaerhout/go-waitgroup.svg)
* [goccm](https://github.com/zenthangplus/goccm) - Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. ![](https://img.shields.io/github/last-commit/zenthangplus/goccm.svg)
* [gohive](https://github.com/loveleshsharma/gohive) - A highly performant and easy to use Goroutine pool for Go. ![](https://img.shields.io/github/last-commit/loveleshsharma/gohive.svg)
* [gollback](https://github.com/vardius/gollback) - asynchronous simple function utilities, for managing execution of closures and callbacks. ![](https://img.shields.io/github/last-commit/vardius/gollback.svg)
* [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. ![](https://img.shields.io/github/last-commit/benmanns/goworker.svg)
* [gowp](https://github.com/xxjwxc/gowp) - gowp is concurrency limiting goroutine pool. ![](https://img.shields.io/github/last-commit/xxjwxc/gowp.svg)
* [gpool](https://github.com/Sherifabdlnaby/gpool) - manages a resizeable pool of context-aware goroutines to bound concurrency. ![](https://img.shields.io/github/last-commit/Sherifabdlnaby/gpool.svg)
* [grpool](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool. ![](https://img.shields.io/github/last-commit/ivpusic/grpool.svg)
* [hands](https://github.com/duanckham/hands) - A process controller used to control the execution and return strategies of multiple goroutines. ![](https://img.shields.io/github/last-commit/duanckham/hands.svg)
* [Hunch](https://github.com/AaronJan/Hunch) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. ![](https://img.shields.io/github/last-commit/AaronJan/Hunch.svg)
* [kyoo](https://github.com/dirkaholic/kyoo) - Provides an unlimited job queue and concurrent worker pools. ![](https://img.shields.io/github/last-commit/dirkaholic/kyoo.svg)
* [nursery](https://github.com/arunsworld/nursery) - Structured concurrency in Go. ![](https://img.shields.io/github/last-commit/arunsworld/nursery.svg)
* [oversight](https://cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel. ![](https://img.shields.io/github/last-commit/rafaeljesus/parallel-fn.svg)
* [pond](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go. ![](https://img.shields.io/github/last-commit/alitto/pond.svg)
* [pool](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. ![](https://img.shields.io/github/last-commit/go-playground/pool.svg)
* [queue](https://github.com/AnikHasibul/queue) - Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. ![](https://img.shields.io/github/last-commit/AnikHasibul/queue.svg)
* [routine](https://github.com/x-mod/routine) - go routine control with context, support: Main, Go, Pool and some useful Executors. ![](https://img.shields.io/github/last-commit/x-mod/routine.svg)
* [semaphore](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. ![](https://img.shields.io/github/last-commit/kamilsk/semaphore.svg)
* [semaphore](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). ![](https://img.shields.io/github/last-commit/marusama/semaphore.svg)
* [stl](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. ![](https://img.shields.io/github/last-commit/ssgreg/stl.svg)
* [threadpool](https://github.com/shettyh/threadpool) - Golang threadpool implementation. ![](https://img.shields.io/github/last-commit/shettyh/threadpool.svg)
* [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. ![](https://img.shields.io/github/last-commit/Jeffail/tunny.svg)
* [worker-pool](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool. ![](https://img.shields.io/github/last-commit/vardius/worker-pool.svg)
* [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. ![](https://img.shields.io/github/last-commit/gammazero/workerpool.svg)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. ![](https://img.shields.io/github/last-commit/murlokswarm/app.svg)
* [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. ![](https://img.shields.io/github/last-commit/fyne-io/fyne.svg)
* [go-astilectron](https://github.com/asticode/go-astilectron) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). ![](https://img.shields.io/github/last-commit/asticode/go-astilectron.svg)
* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. ![](https://img.shields.io/github/last-commit/sciter-sdk/go-sciter.svg)
* [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. ![](https://img.shields.io/github/last-commit/gotk3/gotk3.svg)
* [gowd](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. ![](https://img.shields.io/github/last-commit/dtylman/gowd.svg)
* [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). ![](https://img.shields.io/github/last-commit/therecipe/qt.svg)
* [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. ![](https://img.shields.io/github/last-commit/andlabs/ui.svg)
* [Wails](https://wails.app) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [walk](https://github.com/lxn/walk) - Windows application library kit for Go. ![](https://img.shields.io/github/last-commit/lxn/walk.svg)
* [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). ![](https://img.shields.io/github/last-commit/zserge/webview.svg)

*Interaction*

* [go-appindicator](https://github.com/dawidd6/go-appindicator) - Go bindings for libappindicator3 C library. ![](https://img.shields.io/github/last-commit/dawidd6/go-appindicator.svg)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go. ![](https://img.shields.io/github/last-commit/deckarep/gosx-notifier.svg)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) - OSX library to notify about any (pluggable) activity on your machine. ![](https://img.shields.io/github/last-commit/prashantgupta24/activity-tracker.svg)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) - OSX Sleep/Wake notifications in golang. ![](https://img.shields.io/github/last-commit/prashantgupta24/mac-sleep-notifier.svg)
* [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. ![](https://img.shields.io/github/last-commit/go-vgo/robotgo.svg)
* [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. ![](https://img.shields.io/github/last-commit/getlantern/systray.svg)
* [trayhost](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. ![](https://img.shields.io/github/last-commit/shurcooL/trayhost.svg)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. ![](https://img.shields.io/github/last-commit/anthonynsimon/bild.svg)
* [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. ![](https://img.shields.io/github/last-commit/h2non/bimg.svg)
* [cameron](https://github.com/aofei/cameron) - An avatar generator for Go. ![](https://img.shields.io/github/last-commit/aofei/cameron.svg)
* [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image. ![](https://img.shields.io/github/last-commit/tdewolff/canvas.svg)
* [darkroom](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. ![](https://img.shields.io/github/last-commit/gojek/darkroom.svg)
* [draft](https://github.com/lucasepe/draft) - Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. ![](https://img.shields.io/github/last-commit/lucasepe/draft.svg)
* [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. ![](https://img.shields.io/github/last-commit/pravj/geopattern.svg)
* [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. ![](https://img.shields.io/github/last-commit/fogleman/gg.svg)
* [gift](https://github.com/disintegration/gift) - Package of image processing filters. ![](https://img.shields.io/github/last-commit/disintegration/gift.svg)
* [gltf](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator. ![](https://img.shields.io/github/last-commit/qmuntal/gltf.svg)
* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library. ![](https://img.shields.io/github/last-commit/ungerik/go-cairo.svg)
* [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library. ![](https://img.shields.io/github/last-commit/bolknote/go-gd.svg)
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go. ![](https://img.shields.io/github/last-commit/koyachi/go-nude.svg)
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV. ![](https://img.shields.io/github/last-commit/lazywei/go-opencv.svg)
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go. ![](https://img.shields.io/github/last-commit/jyotiska/go-webcolors.svg)
* [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. ![](https://img.shields.io/github/last-commit/hybridgroup/gocv.svg)
* [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package. ![](https://img.shields.io/github/last-commit/corona10/goimagehash.svg)
* [goimghdr](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go. ![](https://img.shields.io/github/last-commit/corona10/goimghdr.svg)
* [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars. ![](https://img.shields.io/github/last-commit/o1egl/govatar.svg)
* [gridder](https://github.com/shomali11/gridder) - A Grid based 2D Graphics library. ![](https://img.shields.io/github/last-commit/shomali11/gridder.svg)
* [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII. ![](https://img.shields.io/github/last-commit/qeesung/image2ascii.svg)
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. ![](https://img.shields.io/github/last-commit/gographics/imagick.svg)
* [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. ![](https://img.shields.io/github/last-commit/h2non/imaginary.svg)
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. ![](https://img.shields.io/github/last-commit/disintegration/imaging.svg)
* [img](https://github.com/hawx/img) - Selection of image manipulation tools. ![](https://img.shields.io/github/last-commit/hawx/img.svg)
* [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. ![](https://img.shields.io/github/last-commit/fogleman/ln.svg)
* [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). ![](https://img.shields.io/github/last-commit/noelyahan/mergi.svg)
* [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go. ![](https://img.shields.io/github/last-commit/aldor007/mort.svg)
* [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos. ![](https://img.shields.io/github/last-commit/donatj/mpo.svg)
* [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. ![](https://img.shields.io/github/last-commit/thoas/picfit.svg)
* [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. ![](https://img.shields.io/github/last-commit/fogleman/pt.svg)
* [resize](https://github.com/nfnt/resize) - Image resizing for Go with common interpolation methods. ![](https://img.shields.io/github/last-commit/nfnt/resize.svg)
* [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD. ![](https://img.shields.io/github/last-commit/bamiaux/rez.svg)
* [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. ![](https://img.shields.io/github/last-commit/muesli/smartcrop.svg)
* [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography. ![](https://img.shields.io/github/last-commit/auyer/steganography.svg)
* [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image. ![](https://img.shields.io/github/last-commit/DimitarPetrov/stegify.svg)
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. ![](https://img.shields.io/github/last-commit/ajstarks/svgo.svg)
* [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder. ![](https://img.shields.io/github/last-commit/ftrvxmtrx/tga.svg)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT. ![](https://img.shields.io/github/last-commit/connectordb/connectordb.svg)
* [devices](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io. ![](https://img.shields.io/github/last-commit/goiot/devices.svg)
* [eywa](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices. ![](https://img.shields.io/github/last-commit/xcodersun/eywa.svg)
* [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. ![](https://img.shields.io/github/last-commit/tibcosoftware/flogo.svg)
* [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. ![](https://img.shields.io/github/last-commit/paypal/gatt.svg)
* [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [huego](https://github.com/amimof/huego) - An extensive Philips Hue client library for Go. ![](https://img.shields.io/github/last-commit/amimof/huego.svg)
* [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device.
* [mainflux](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server. ![](https://img.shields.io/github/last-commit/Mainflux/mainflux.svg)
* [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.
* [sensorbee](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT. ![](https://img.shields.io/github/last-commit/sensorbee/sensorbee.svg)

## Job Scheduler

*Libraries for scheduling jobs.*

* [clockwerk](http://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax.
* [clockwork](https://github.com/whiteShtef/clockwork) - Simple and intuitive job scheduling library in Go. ![](https://img.shields.io/github/last-commit/whiteShtef/clockwork.svg)
* [go-cron](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. ![](https://img.shields.io/github/last-commit/rk/go-cron.svg)
* [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. ![](https://img.shields.io/github/last-commit/roylee0704/gron.svg)
* [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. ![](https://img.shields.io/github/last-commit/bamzi/jobrunner.svg)
* [jobs](https://github.com/albrow/jobs) - Persistent and flexible background jobs library. ![](https://img.shields.io/github/last-commit/albrow/jobs.svg)
* [leprechaun](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling. ![](https://img.shields.io/github/last-commit/kilgaloon/leprechaun.svg)
* [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy. ![](https://img.shields.io/github/last-commit/carlescere/scheduler.svg)

## JSON

*Libraries for working with JSON.*

* [ajson](https://github.com/spyzhov/ajson) - Abstract JSON for golang with JSONPath support. ![](https://img.shields.io/github/last-commit/spyzhov/ajson.svg)
* [ej](https://github.com/lucassscaravelli/ej) - Write and read JSON from different sources succinctly. ![](https://img.shields.io/github/last-commit/lucassscaravelli/ej.svg)
* [gjo](https://github.com/skanehira/gjo) - Small utility to create JSON objects. ![](https://img.shields.io/github/last-commit/skanehira/gjo.svg)
* [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. ![](https://img.shields.io/github/last-commit/tidwall/gjson.svg)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) - Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. ![](https://img.shields.io/github/last-commit/ddymko/go-jsonerror.svg)
* [go-respond](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses. ![](https://img.shields.io/github/last-commit/nicklaw5/go-respond.svg)
* [gojq](https://github.com/elgs/gojq) - JSON query in Golang. ![](https://img.shields.io/github/last-commit/elgs/gojq.svg)
* [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. ![](https://img.shields.io/github/last-commit/ChimeraCoder/gojson.svg)
* [JayDiff](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go. ![](https://img.shields.io/github/last-commit/yazgazan/jaydiff.svg)
* [jettison](https://github.com/wI2L/jettison) - High performance, reflection-less JSON encoder for Go. ![](https://img.shields.io/github/last-commit/wI2L/jettison.svg)
* [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.
* [json2go](https://github.com/m-zajac/json2go) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. ![](https://img.shields.io/github/last-commit/m-zajac/json2go.svg)
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference. ![](https://img.shields.io/github/last-commit/AmuzaTkts/jsonapi-errors.svg)
* [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON. ![](https://img.shields.io/github/last-commit/miolini/jsonf.svg)
* [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. ![](https://img.shields.io/github/last-commit/ricardolonga/jsongo.svg)
* [jsonhal](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses. ![](https://img.shields.io/github/last-commit/RichardKnop/jsonhal.svg)
* [kazaam](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents. ![](https://img.shields.io/github/last-commit/Qntfy/kazaam.svg)
* [mapslice-json](https://github.com/mickep76/mapslice-json) - Go MapSlice for ordered marshal/ unmarshal of maps in JSON. ![](https://img.shields.io/github/last-commit/mickep76/mapslice-json.svg)
* [mp](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON. ![](https://img.shields.io/github/last-commit/sanbornm/mp.svg)

## Logging

*Libraries for generating and working with log files.*

* [distillog](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels). ![](https://img.shields.io/github/last-commit/amoghe/distillog.svg)
* [glg](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go. ![](https://img.shields.io/github/last-commit/kpango/glg.svg)
* [glo](https://github.com/lajosbencz/glo) - PHP Monolog inspired logging facility with identical severity levels. ![](https://img.shields.io/github/last-commit/lajosbencz/glo.svg)
* [glog](https://github.com/golang/glog) - Leveled execution logs for Go. ![](https://img.shields.io/github/last-commit/golang/glog.svg)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog. ![](https://img.shields.io/github/last-commit/utahta/go-cronowriter.svg)
* [go-log](https://github.com/pieterclaerhout/go-log) - A logging library with strack traces, object dumping and optional timestamps. ![](https://img.shields.io/github/last-commit/pieterclaerhout/go-log.svg)
* [go-log](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers. ![](https://img.shields.io/github/last-commit/subchen/go-log.svg)
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers. ![](https://img.shields.io/github/last-commit/siddontang/go-log.svg)
* [go-log](https://github.com/ian-kent/go-log) - Log4j implementation in Go. ![](https://img.shields.io/github/last-commit/ian-kent/go-log.svg)
* [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers. ![](https://img.shields.io/github/last-commit/apsdehal/go-logger.svg)
* [gologger](https://github.com/sadlil/gologger) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. ![](https://img.shields.io/github/last-commit/sadlil/gologger.svg)
* [gomol](https://github.com/aphistic/gomol) - Multiple-output, structured logging for Go with extensible logging outputs. ![](https://img.shields.io/github/last-commit/aphistic/gomol.svg)
* [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.
* [journald](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging. ![](https://img.shields.io/github/last-commit/ssgreg/journald.svg)
* [log](https://github.com/aerogo/log) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). ![](https://img.shields.io/github/last-commit/aerogo/log.svg)
* [log](https://github.com/apex/log) - Structured logging package for Go. ![](https://img.shields.io/github/last-commit/apex/log.svg)
* [log](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go. ![](https://img.shields.io/github/last-commit/go-playground/log.svg)
* [log](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation. ![](https://img.shields.io/github/last-commit/teris-io/log.svg)
* [log-voyage](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang. ![](https://img.shields.io/github/last-commit/firstrow/logvoyage.svg)
* [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. ![](https://img.shields.io/github/last-commit/inconshreveable/log15.svg)
* [logdump](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging. ![](https://img.shields.io/github/last-commit/ewwwwwqm/logdump.svg)
* [logex](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib. ![](https://img.shields.io/github/last-commit/chzyer/logex.svg)
* [logger](https://github.com/azer/logger) - Minimalistic logging library for Go. ![](https://img.shields.io/github/last-commit/azer/logger.svg)
* [logmatic](https://github.com/borderstech/logmatic) - Colorized logger for Golang with dynamic log level configuration. ![](https://img.shields.io/github/last-commit/borderstech/logmatic.svg)
* [logo](https://github.com/mbndr/logo) - Golang logger to different configurable writers. ![](https://img.shields.io/github/last-commit/mbndr/logo.svg)
* [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. ![](https://img.shields.io/github/last-commit/Sirupsen/logrus.svg)
* [logrusiowriter](https://github.com/cabify/logrusiowriter) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. ![](https://img.shields.io/github/last-commit/cabify/logrusiowriter.svg)
* [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). ![](https://img.shields.io/github/last-commit/sebest/logrusly.svg)
* [logur](https://github.com/logur/logur) - An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). ![](https://img.shields.io/github/last-commit/logur/logur.svg)
* [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger. ![](https://img.shields.io/github/last-commit/hashicorp/logutils.svg)
* [logxi](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy. ![](https://img.shields.io/github/last-commit/mgutz/logxi.svg)
* [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. ![](https://img.shields.io/github/last-commit/natefinch/lumberjack.svg)
* [mlog](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. ![](https://img.shields.io/github/last-commit/jbrodriguez/mlog.svg)
* [onelog](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. ![](https://img.shields.io/github/last-commit/francoispqt/onelog.svg)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). ![](https://img.shields.io/github/last-commit/go-ozzo/ozzo-log.svg)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. ![](https://img.shields.io/github/last-commit/arthurkiller/rollingWriter.svg)
* [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. ![](https://img.shields.io/github/last-commit/cihub/seelog.svg)
* [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. ![](https://img.shields.io/github/last-commit/davecgh/go-spew.svg)
* [sqldb-logger](https://github.com/simukti/sqldb-logger) - A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. ![](https://img.shields.io/github/last-commit/simukti/sqldb-logger.svg)
* [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. ![](https://img.shields.io/github/last-commit/alexcesaro/log.svg)
* [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. ![](https://img.shields.io/github/last-commit/hpcloud/tail.svg)
* [xlog](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. ![](https://img.shields.io/github/last-commit/xfxdev/xlog.svg)
* [xlog](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching. ![](https://img.shields.io/github/last-commit/rs/xlog.svg)
* [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. ![](https://img.shields.io/github/last-commit/uber-go/zap.svg)
* [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. ![](https://img.shields.io/github/last-commit/rs/zerolog.svg)

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. ![](https://img.shields.io/github/last-commit/jbrukh/bayesian.svg)
* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. ![](https://img.shields.io/github/last-commit/ryanbressler/CloudForest.svg)
* [eaopt](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library. ![](https://img.shields.io/github/last-commit/MaxHalford/eaopt.svg)
* [evoli](https://github.com/khezen/evoli) - Genetic Algorithm and Particle Swarm Optimization library. ![](https://img.shields.io/github/last-commit/khezen/evoli.svg)
* [fonet](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go. ![](https://img.shields.io/github/last-commit/Fontinalis/fonet.svg)
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms. ![](https://img.shields.io/github/last-commit/e-XpertSolutions/go-cluster.svg)
* [go-deep](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go. ![](https://img.shields.io/github/last-commit/patrikeh/go-deep.svg)
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library. ![](https://img.shields.io/github/last-commit/white-pony/go-fann.svg)
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang. ![](https://img.shields.io/github/last-commit/thoj/go-galib.svg)
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. ![](https://img.shields.io/github/last-commit/daviddengcn/go-pr.svg)
* [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go. ![](https://img.shields.io/github/last-commit/goml/gobrain.svg)
* [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods. ![](https://img.shields.io/github/last-commit/e-dard/godist.svg)
* [goga](https://github.com/tomcraven/goga) - Genetic algorithm library for Go. ![](https://img.shields.io/github/last-commit/tomcraven/goga.svg)
* [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. ![](https://img.shields.io/github/last-commit/sjwhitworth/golearn.svg)
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go. ![](https://img.shields.io/github/last-commit/danieldk/golinear.svg)
* [GoMind](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go. ![](https://img.shields.io/github/last-commit/surenderthakran/gomind.svg)
* [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. ![](https://img.shields.io/github/last-commit/cdipaolo/goml.svg)
* [gonet](https://github.com/dathoangnd/gonet) - Neural Network for Go. ![](https://img.shields.io/github/last-commit/dathoangnd/gonet.svg)
* [Goptuna](https://github.com/c-bata/goptuna) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. ![](https://img.shields.io/github/last-commit/c-bata/goptuna.svg)
* [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go. ![](https://img.shields.io/github/last-commit/timkaye11/goRecommend.svg)
* [gorgonia](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. ![](https://img.shields.io/github/last-commit/gorgonia/gorgonia.svg)
* [gorse](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go. ![](https://img.shields.io/github/last-commit/zhenghaoz/gorse.svg)
* [goscore](https://github.com/asafschers/goscore) - Go Scoring API for PMML. ![](https://img.shields.io/github/last-commit/asafschers/goscore.svg)
* [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. ![](https://img.shields.io/github/last-commit/otiai10/gosseract.svg)
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14. ![](https://img.shields.io/github/last-commit/datastream/libsvm.svg)
* [neat](https://github.com/jinyeom/neat) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). ![](https://img.shields.io/github/last-commit/jinyeom/neat.svg)
* [neural-go](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation. ![](https://img.shields.io/github/last-commit/schuyler/neural-go.svg)
* [ocrserver](https://github.com/otiai10/ocrserver) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku. ![](https://img.shields.io/github/last-commit/otiai10/ocrserver.svg)
* [onnx-go](https://github.com/owulveryck/onnx-go) - Go Interface to Open Neural Network Exchange (ONNX). ![](https://img.shields.io/github/last-commit/owulveryck/onnx-go.svg)
* [probab](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go. ![](https://img.shields.io/github/last-commit/ThePaw/probab.svg)
* [randomforest](https://github.com/malaschitz/randomForest) - Easy to use Random Forest library for Go. ![](https://img.shields.io/github/last-commit/malaschitz/randomForest.svg)
* [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine. ![](https://img.shields.io/github/last-commit/muesli/regommend.svg)
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go. ![](https://img.shields.io/github/last-commit/eaigner/shield.svg)
* [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. ![](https://img.shields.io/github/last-commit/galeone/tfgo.svg)
* [Varis](https://github.com/Xamber/Varis) - Golang Neural Network. ![](https://img.shields.io/github/last-commit/Xamber/Varis.svg)

## Messaging

*Libraries that implement messaging systems.*

* [ami](https://github.com/kak-tus/ami) - Go client to reliable queues based on Redis Cluster Streams. ![](https://img.shields.io/github/last-commit/kak-tus/ami.svg)
* [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. ![](https://img.shields.io/github/last-commit/sideshow/apns2.svg)
* [Asynq](https://github.com/hibiken/asynq) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis. ![](https://img.shields.io/github/last-commit/hibiken/asynq.svg)
* [Beaver](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. ![](https://img.shields.io/github/last-commit/Clivern/Beaver.svg)
* [Benthos](https://github.com/Jeffail/benthos) - A message streaming bridge between a range of protocols. ![](https://img.shields.io/github/last-commit/Jeffail/benthos.svg)
* [Bus](https://github.com/mustafaturan/bus) - Minimalist message bus implementation for internal communication. ![](https://img.shields.io/github/last-commit/mustafaturan/bus.svg)
* [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. ![](https://img.shields.io/github/last-commit/centrifugal/centrifugo.svg)
* [Commander](https://github.com/jeroenrinzema/commander) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. ![](https://img.shields.io/github/last-commit/jeroenrinzema/commander.svg)
* [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. ![](https://img.shields.io/github/last-commit/confluentinc/confluent-kafka-go.svg)
* [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. ![](https://img.shields.io/github/last-commit/godbus/dbus.svg)
* [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. ![](https://img.shields.io/github/last-commit/appleboy/drone-line.svg)
* [emitter](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. ![](https://img.shields.io/github/last-commit/olebedev/emitter.svg)
* [event](https://github.com/agoalofalife/event) - Implementation of the pattern observer. ![](https://img.shields.io/github/last-commit/agoalofalife/event.svg)
* [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. ![](https://img.shields.io/github/last-commit/asaskevich/EventBus.svg)
* [gaurun-client](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go. ![](https://img.shields.io/github/last-commit/osamingo/gaurun-client.svg)
* [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io). ![](https://img.shields.io/github/last-commit/desertbit/glue.svg)
* [go-mq](https://github.com/cheshir/go-mq) - RabbitMQ client with declarative configuration. ![](https://img.shields.io/github/last-commit/cheshir/go-mq.svg)
* [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec. ![](https://img.shields.io/github/last-commit/TheCreeper/go-notify.svg)
* [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. ![](https://img.shields.io/github/last-commit/nsqio/go-nsq.svg)
* [go-res](https://github.com/jirenius/go-res) - Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. ![](https://img.shields.io/github/last-commit/jirenius/go-res.svg)
* [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. ![](https://img.shields.io/github/last-commit/googollee/go-socket.io.svg)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service. ![](https://img.shields.io/github/last-commit/maxatome/go-vitotrol.svg)
* [Gollum](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. ![](https://img.shields.io/github/last-commit/trivago/gollum.svg)
* [golongpoll](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple. ![](https://img.shields.io/github/last-commit/jcuga/golongpoll.svg)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. ![](https://img.shields.io/github/last-commit/Terry-Mao/gopush-cluster.svg)
* [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). ![](https://img.shields.io/github/last-commit/appleboy/gorush.svg)
* [guble](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. ![](https://img.shields.io/github/last-commit/smancke/guble.svg)
* [hub](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. ![](https://img.shields.io/github/last-commit/leandro-lugaresi/hub.svg)
* [jazz](https://github.com/socifi/jazz) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. ![](https://img.shields.io/github/last-commit/socifi/jazz.svg)
* [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. ![](https://img.shields.io/github/last-commit/RichardKnop/machinery.svg)
* [mangos](https://github.com/go-mangos/mangos) - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. ![](https://img.shields.io/github/last-commit/go-mangos/mangos.svg)
* [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. ![](https://img.shields.io/github/last-commit/olahol/melody.svg)
* [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). ![](https://img.shields.io/github/last-commit/dunglas/mercure.svg)
* [messagebus](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. ![](https://img.shields.io/github/last-commit/vardius/message-bus.svg)
* [NATS Go Client](https://github.com/nats-io/nats) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. ![](https://img.shields.io/github/last-commit/nats-io/nats.svg)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel. ![](https://img.shields.io/github/last-commit/rafaeljesus/nsq-event-bus.svg)
* [oplog](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs. ![](https://img.shields.io/github/last-commit/dailymotion/oplog.svg)
* [pubsub](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go. ![](https://img.shields.io/github/last-commit/tuxychandru/pubsub.svg)
* [rabbus](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues. ![](https://img.shields.io/github/last-commit/rafaeljesus/rabbus.svg)
* [rabtap](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app. ![](https://img.shields.io/github/last-commit/jandelgado/rabtap.svg)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue. ![](https://img.shields.io/github/last-commit/sybrexsys/RapidMQ.svg)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) - redisqueue provides a producer and consumer of a queue that uses Redis streams. ![](https://img.shields.io/github/last-commit/robinjoseph08/redisqueue.svg)
* [rmqconn](https://github.com/sbabiv/rmqconn) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. ![](https://img.shields.io/github/last-commit/sbabiv/rmqconn.svg)
* [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. ![](https://img.shields.io/github/last-commit/Shopify/sarama.svg)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. ![](https://img.shields.io/github/last-commit/uniqush/uniqush-push.svg)
* [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). ![](https://img.shields.io/github/last-commit/pebbe/zmq4.svg)

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. ![](https://img.shields.io/github/last-commit/unidoc/unioffice.svg)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - Golang library for reading and writing Microsoft Excel™ (XLSX) files. ![](https://img.shields.io/github/last-commit/360EntSecGroup-Skylar/excelize.svg)
* [go-excel](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table. ![](https://img.shields.io/github/last-commit/szyhf/go-excel.svg)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. ![](https://img.shields.io/github/last-commit/fterrag/goxlsxwriter.svg)
* [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. ![](https://img.shields.io/github/last-commit/tealeg/xlsx.svg)
* [xlsx](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs. ![](https://img.shields.io/github/last-commit/plandem/xlsx.svg)

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [alice](https://github.com/magic003/alice) - Additive dependency injection container for Golang. ![](https://img.shields.io/github/last-commit/magic003/alice.svg)
* [container](https://github.com/golobby/container) - A powerful IoC Container with fluent and easy-to-use interface. ![](https://img.shields.io/github/last-commit/golobby/container.svg)
* [di](https://github.com/goava/di) - A dependency injection container for go programming language. ![](https://img.shields.io/github/last-commit/goava/di.svg)
* [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go. ![](https://img.shields.io/github/last-commit/uber-go/dig.svg)
* [dingo](https://github.com/i-love-flamingo/dingo) - A dependency injection toolkit for Go, based on Guice. ![](https://img.shields.io/github/last-commit/i-love-flamingo/dingo.svg)
* [fx](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig). ![](https://img.shields.io/github/last-commit/uber-go/fx.svg)
* [gocontainer](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container. ![](https://img.shields.io/github/last-commit/vardius/gocontainer.svg)
* [goioc/di](https://github.com/goioc/di) - Spring-inspired Dependency Injection Container. ![](https://img.shields.io/github/last-commit/goioc/di.svg)
* [linker](https://github.com/logrange/linker) - A reflection based dependency injection and inversion of control library with components lifecycle support. ![](https://img.shields.io/github/last-commit/logrange/linker.svg)
* [wire](https://github.com/Fs02/wire) - Strict Runtime Dependency Injection for Golang. ![](https://img.shields.io/github/last-commit/Fs02/wire.svg)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) - A Go application boilerplate template for quick starting projects following production best practices. ![](https://img.shields.io/github/last-commit/lacion/cookiecutter-golang.svg)
* [go-sample](https://github.com/zitryss/go-sample) - A sample layout for Go application projects with the real code. ![](https://img.shields.io/github/last-commit/zitryss/go-sample.svg)
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. ![](https://img.shields.io/github/last-commit/golang-standards/project-layout.svg)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices. ![](https://img.shields.io/github/last-commit/sagikazarmark/modern-go-application.svg)
* [scaffold](https://github.com/catchplay/scaffold) - Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. ![](https://img.shields.io/github/last-commit/catchplay/scaffold.svg)

### Strings

*Libraries for working with strings.*

* [gobeam/Stringy](https://github.com/gobeam/Stringy) - String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. ![](https://img.shields.io/github/last-commit/gobeam/Stringy.svg)
* [strutil](https://github.com/ozgio/strutil) - String utilities. ![](https://img.shields.io/github/last-commit/ozgio/strutil.svg)
* [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. ![](https://img.shields.io/github/last-commit/huandu/xstrings.svg)

### Uncategorized

*These libraries were placed here because none of the other categories seemed to fit.*

* [anagent](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. ![](https://img.shields.io/github/last-commit/mudler/anagent.svg)
* [antch](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework. ![](https://img.shields.io/github/last-commit/antchfx/antch.svg)
* [archiver](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives. ![](https://img.shields.io/github/last-commit/mholt/archiver.svg)
* [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields. ![](https://img.shields.io/github/last-commit/artyom/autoflags.svg)
* [avgRating](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation. ![](https://img.shields.io/github/last-commit/kirillDanshin/avgRating.svg)
* [banner](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications. ![](https://img.shields.io/github/last-commit/dimiro1/banner.svg)
* [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. ![](https://img.shields.io/github/last-commit/mojocn/base64Captcha.svg)
* [battery](https://github.com/distatus/battery) - Cross-platform, normalized battery information library. ![](https://img.shields.io/github/last-commit/distatus/battery.svg)
* [bitio](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go. ![](https://img.shields.io/github/last-commit/icza/bitio.svg)
* [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](http://browscap.org/). ![](https://img.shields.io/github/last-commit/digitalcrab/browscap_go.svg)
* [captcha](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation. ![](https://img.shields.io/github/last-commit/steambap/captcha.svg)
* [conv](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types. ![](https://img.shields.io/github/last-commit/cstockton/go-conv.svg)
* [datacounter](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter. ![](https://img.shields.io/github/last-commit/miolini/datacounter.svg)
* [ffmt](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans. ![](https://img.shields.io/github/last-commit/go-ffmt/ffmt.svg)
* [ghorg](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. ![](https://img.shields.io/github/last-commit/gabrie30/ghorg.svg)
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. ![](https://img.shields.io/github/last-commit/jolestar/go-commons-pool.svg)
* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.
* [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. ![](https://img.shields.io/github/last-commit/eapache/go-resiliency.svg)
* [go-unarr](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives. ![](https://img.shields.io/github/last-commit/gen2brain/go-unarr.svg)
* [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. ![](https://img.shields.io/github/last-commit/brianvoe/gofakeit.svg)
* [gommit](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns. ![](https://img.shields.io/github/last-commit/antham/gommit.svg)
* [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). ![](https://img.shields.io/github/last-commit/shirou/gopsutil.svg)
* [gosh](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method. ![](https://img.shields.io/github/last-commit/osamingo/gosh.svg)
* [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. ![](https://img.shields.io/github/last-commit/haxpax/gosms.svg)
* [gotoprom](https://github.com/cabify/gotoprom) - Type-safe metrics builder wrapper library for the official Prometheus client. ![](https://img.shields.io/github/last-commit/cabify/gotoprom.svg)
* [gountries](https://github.com/pariz/gountries) - Package that exposes country and subdivision data. ![](https://img.shields.io/github/last-commit/pariz/gountries.svg)
* [health](https://github.com/dimiro1/health) - Easy to use, extensible health check library. ![](https://img.shields.io/github/last-commit/dimiro1/health.svg)
* [healthcheck](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services. ![](https://img.shields.io/github/last-commit/etherlabsio/healthcheck.svg)
* [hostutils](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list. ![](https://img.shields.io/github/last-commit/Wing924/hostutils.svg)
* [indigo](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58. ![](https://img.shields.io/github/last-commit/osamingo/indigo.svg)
* [lk](https://github.com/hyperboloide/lk) - A simple licensing library for golang. ![](https://img.shields.io/github/last-commit/hyperboloide/lk.svg)
* [llvm](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go. ![](https://img.shields.io/github/last-commit/llir/llvm.svg)
* [metrics](https://github.com/pascaldekloe/metrics) - Library for metrics instrumentation and Prometheus exposition. ![](https://img.shields.io/github/last-commit/pascaldekloe/metrics.svg)
* [morse](https://github.com/alwindoss/morse) - Library to convert to and from morse code. ![](https://img.shields.io/github/last-commit/alwindoss/morse.svg)
* [numa](https://github.com/lrita/numa) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. ![](https://img.shields.io/github/last-commit/lrita/numa.svg)
* [pdfgen](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests. ![](https://img.shields.io/github/last-commit/hyperboloide/pdfgen.svg)
* [persian](https://github.com/mavihq/persian) - Some utilities for Persian language in go. ![](https://img.shields.io/github/last-commit/mavihq/persian.svg)
* [sandid](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID. ![](https://img.shields.io/github/last-commit/aofei/sandid.svg)
* [shellwords](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. ![](https://img.shields.io/github/last-commit/Wing924/shellwords.svg)
* [shortid](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs. ![](https://img.shields.io/github/last-commit/teris-io/shortid.svg)
* [shoutrrr](https://github.com/containrrr/shoutrrr) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. ![](https://img.shields.io/github/last-commit/containrrr/shoutrrr.svg)
* [stateless](https://github.com/qmuntal/stateless) - A fluent library for creating state machines. ![](https://img.shields.io/github/last-commit/qmuntal/stateless.svg)
* [stats](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... ![](https://img.shields.io/github/last-commit/go-playground/stats.svg)
* [turtle](https://github.com/hackebrot/turtle) - Emojis for Go. ![](https://img.shields.io/github/last-commit/hackebrot/turtle.svg)
* [url-shortener](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support. ![](https://img.shields.io/github/last-commit/pantrif/url-shortener.svg)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling.
* [xdg](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go. ![](https://img.shields.io/github/last-commit/rkoesters/xdg.svg)
* [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber. ![](https://img.shields.io/github/last-commit/go-xkg/xkg.svg)

## Natural Language Processing

*Libraries for working with human languages.*

* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. ![](https://img.shields.io/github/last-commit/detectlanguage/detectlanguage-go.svg)
* [getlang](https://github.com/rylans/getlang) - Fast natural language detection package. ![](https://img.shields.io/github/last-commit/rylans/getlang.svg)
* [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text.
* [go-localize](https://github.com/m1/go-localize) - Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. ![](https://img.shields.io/github/last-commit/m1/go-localize.svg)
* [go-mystem](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer. ![](https://img.shields.io/github/last-commit/dveselov/mystem.svg)
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work. ![](https://img.shields.io/github/last-commit/nuance/go-nlp.svg)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. ![](https://img.shields.io/github/last-commit/mozillazg/go-pinyin.svg)
* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm. ![](https://img.shields.io/github/last-commit/agonopol/go-stem.svg)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text. ![](https://img.shields.io/github/last-commit/mozillazg/go-unidecode.svg)
* [go2vec](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings. ![](https://img.shields.io/github/last-commit/danieldk/go2vec.svg)
* [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. ![](https://img.shields.io/github/last-commit/yanyiwu/gojieba.svg)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2. ![](https://img.shields.io/github/last-commit/rjohnsondev/golibstemmer.svg)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) ![](https://img.shields.io/github/last-commit/xujiajun/gotokenizer.svg)
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go. ![](https://img.shields.io/github/last-commit/fiam/gounidecode.svg)
* [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. ![](https://img.shields.io/github/last-commit/go-ego/gse.svg)
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. ![](https://img.shields.io/github/last-commit/goodsign/icu.svg)
* [iuliia-go](https://github.com/mehanizm/iuliia-go) - Transliterate Cyrillic → Latin in every possible way. ![](https://img.shields.io/github/last-commit/mehanizm/iuliia-go.svg)
* [kagome](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go. ![](https://img.shields.io/github/last-commit/ikawaha/kagome.svg)
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. ![](https://img.shields.io/github/last-commit/goodsign/libtextcat.svg)
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. ![](https://img.shields.io/github/last-commit/awsong/MMSEGO.svg)
* [nlp](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp. ![](https://img.shields.io/github/last-commit/Shixzie/nlp.svg)
* [nlp](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). ![](https://img.shields.io/github/last-commit/james-bowman/nlp.svg)
* [paicehusk](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm. ![](https://img.shields.io/github/last-commit/rookii/paicehusk.svg)
* [petrovich](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case. ![](https://img.shields.io/github/last-commit/striker2000/petrovich.svg)
* [porter](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. ![](https://img.shields.io/github/last-commit/a2800276/porter.svg)
* [porter2](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer. ![](https://img.shields.io/github/last-commit/zhenjl/porter2.svg)
* [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. ![](https://img.shields.io/github/last-commit/jdkato/prose.svg)
* [RAKE.go](https://github.com/Obaied/RAKE.go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). ![](https://img.shields.io/github/last-commit/Obaied/RAKE.go.svg)
* [segment](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/) ![](https://img.shields.io/github/last-commit/blevesearch/segment.svg)
* [sentences](https://github.com/neurosnap/sentences) - Sentence tokenizer:  converts text into a list of sentences. ![](https://img.shields.io/github/last-commit/neurosnap/sentences.svg)
* [shamoji](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go. ![](https://img.shields.io/github/last-commit/osamingo/shamoji.svg)
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). ![](https://img.shields.io/github/last-commit/goodsign/snowball.svg)
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. ![](https://img.shields.io/github/last-commit/dchest/stemmer.svg)
* [textcat](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text. ![](https://img.shields.io/github/last-commit/pebbe/textcat.svg)
* [transliterator](https://github.com/alexsergivan/transliterator) - Provides one-way string transliteration with supporting of language-specific transliteration rules. ![](https://img.shields.io/github/last-commit/alexsergivan/transliterator.svg)
* [whatlanggo](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). ![](https://img.shields.io/github/last-commit/abadojack/whatlanggo.svg)
* [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. ![](https://img.shields.io/github/last-commit/olebedev/when.svg)

## Networking

*Libraries for working with various layers of the network.*

* [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. ![](https://img.shields.io/github/last-commit/mdlayher/arp.svg)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy. ![](https://img.shields.io/github/last-commit/stabbycutyou/buffstreams.svg)
* [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252). ![](https://img.shields.io/github/last-commit/zubairhamed/canopus.svg)
* [cidranger](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go. ![](https://img.shields.io/github/last-commit/yl2chen/cidranger.svg)
* [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. ![](https://img.shields.io/github/last-commit/mdlayher/dhcp6.svg)
* [dns](https://github.com/miekg/dns) - Go library for working with DNS. ![](https://img.shields.io/github/last-commit/miekg/dns.svg)
* [ether](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames. ![](https://img.shields.io/github/last-commit/songgao/ether.svg)
* [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. ![](https://img.shields.io/github/last-commit/mdlayher/ethernet.svg)
* [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. ![](https://img.shields.io/github/last-commit/valyala/fasthttp.svg)
* [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. ![](https://img.shields.io/github/last-commit/fortio/fortio.svg)
* [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). ![](https://img.shields.io/github/last-commit/jlaffaye/ftp.svg)
* [gaio](https://github.com/xtaci/gaio) - High performance async-io networking for Golang in proactor mode. ![](https://img.shields.io/github/last-commit/xtaci/gaio.svg)
* [gev](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. ![](https://img.shields.io/github/last-commit/Allenxuxu/gev.svg)
* [gmqtt](https://github.com/DrmagicE/gmqtt) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. ![](https://img.shields.io/github/last-commit/DrmagicE/gmqtt.svg)
* [gnet](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. ![](https://img.shields.io/github/last-commit/panjf2000/gnet.svg)
* [gNxI](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols. ![](https://img.shields.io/github/last-commit/google/gnxi.svg)
* [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. ![](https://img.shields.io/github/last-commit/hashicorp/go-getter.svg)
* [go-powerdns](https://github.com/joeig/go-powerdns) - PowerDNS API bindings for Golang. ![](https://img.shields.io/github/last-commit/joeig/go-powerdns.svg)
* [go-stun](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389). ![](https://img.shields.io/github/last-commit/ccding/go-stun.svg)
* [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. ![](https://img.shields.io/github/last-commit/osrg/gobgp.svg)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. ![](https://img.shields.io/github/last-commit/sunwxg/golibwireshark.svg)
* [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. ![](https://img.shields.io/github/last-commit/google/gopacket.svg)
* [gopcap](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap. ![](https://img.shields.io/github/last-commit/akrennmair/gopcap.svg)
* [goshark](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet. ![](https://img.shields.io/github/last-commit/sunwxg/goshark.svg)
* [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. ![](https://img.shields.io/github/last-commit/soniah/gosnmp.svg)
* [gosocsvr](https://github.com/rakeki/gosocsvr) - Socket server made simple. ![](https://img.shields.io/github/last-commit/rakeki/gosocsvr.svg)
* [gotcp](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications. ![](https://img.shields.io/github/last-commit/gansidui/gotcp.svg)
* [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. ![](https://img.shields.io/github/last-commit/cavaliercoder/grab.svg)
* [graval](https://github.com/koofr/graval) - Experimental FTP server framework. ![](https://img.shields.io/github/last-commit/koofr/graval.svg)
* [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. ![](https://img.shields.io/github/last-commit/gchaincl/httplab.svg)
* [httpproxy](https://github.com/wzshiming/httpproxy) - HTTP proxy handler and dialer. ![](https://img.shields.io/github/last-commit/wzshiming/httpproxy.svg)
* [iplib](https://github.com/c-robinson/iplib) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html) ![](https://img.shields.io/github/last-commit/c-robinson/iplib.svg)
* [jazigo](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices. ![](https://img.shields.io/github/last-commit/udhos/jazigo.svg)
* [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. ![](https://img.shields.io/github/last-commit/xtaci/kcp-go.svg)
* [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. ![](https://img.shields.io/github/last-commit/xtaci/kcptun.svg)
* [lhttp](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily. ![](https://img.shields.io/github/last-commit/fanux/lhttp.svg)
* [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces. ![](https://img.shields.io/github/last-commit/ian-kent/linkio.svg)
* [llb](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. ![](https://img.shields.io/github/last-commit/kirillDanshin/llb.svg)
* [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. ![](https://img.shields.io/github/last-commit/hashicorp/mdns.svg)
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). ![](https://img.shields.io/github/last-commit/intel-go/nff-go.svg)
* [packet](https://github.com/aerogo/packet) - Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. ![](https://img.shields.io/github/last-commit/aerogo/packet.svg)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast. ![](https://img.shields.io/github/last-commit/schollz/peerdiscovery.svg)
* [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it. ![](https://img.shields.io/github/last-commit/aybabtme/portproxy.svg)
* [publicip](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress). ![](https://img.shields.io/github/last-commit/polera/publicip.svg)
* [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. ![](https://img.shields.io/github/last-commit/lucas-clemente/quic-go.svg)
* [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface. ![](https://img.shields.io/github/last-commit/mdlayher/raw.svg)
* [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. ![](https://img.shields.io/github/last-commit/pkg/sftp.svg)
* [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). ![](https://img.shields.io/github/last-commit/gliderlabs/ssh.svg)
* [sslb](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. ![](https://img.shields.io/github/last-commit/eduardonunesp/sslb.svg)
* [stun](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol. ![](https://img.shields.io/github/last-commit/go-rtc/stun.svg)
* [tcp_server](https://github.com/firstrow/tcp_server) - Go library for building tcp servers faster. ![](https://img.shields.io/github/last-commit/firstrow/tcp_server.svg)
* [tspool](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server. ![](https://img.shields.io/github/last-commit/two/tspool.svg)
* [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation. ![](https://img.shields.io/github/last-commit/anacrolix/utp.svg)
* [vssh](https://github.com/yahoo/vssh) - Go library for building network and server automation over SSH protocol. ![](https://img.shields.io/github/last-commit/yahoo/vssh.svg)
* [water](https://github.com/songgao/water) - Simple TUN/TAP library. ![](https://img.shields.io/github/last-commit/songgao/water.svg)
* [webrtc](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API. ![](https://img.shields.io/github/last-commit/pions/webrtc.svg)
* [winrm](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines. ![](https://img.shields.io/github/last-commit/masterzen/winrm.svg)
* [xtcp](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. ![](https://img.shields.io/github/last-commit/xfxdev/xtcp.svg)

### HTTP Clients

*Libraries for making HTTP requests.*

* [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. ![](https://img.shields.io/github/last-commit/h2non/gentleman.svg)
* [go-http-client](https://github.com/bozd4g/go-http-client) - Make http calls simply and easily. ![](https://img.shields.io/github/last-commit/bozd4g/go-http-client.svg)
* [grequests](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library. ![](https://img.shields.io/github/last-commit/levigross/grequests.svg)
* [heimdall](https://github.com/gojektech/heimdall) - An enchanced http client with retry and hystrix capabilities. ![](https://img.shields.io/github/last-commit/gojektech/heimdall.svg)
* [httpretry](https://github.com/ybbus/httpretry) - Enriches the default go HTTP client with retry functionality. ![](https://img.shields.io/github/last-commit/ybbus/httpretry.svg)
* [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency. ![](https://img.shields.io/github/last-commit/sethgrid/pester.svg)
* [request](https://github.com/monaco-io/request) - HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. ![](https://img.shields.io/github/last-commit/monaco-io/request.svg)
* [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. ![](https://img.shields.io/github/last-commit/go-resty/resty.svg)
* [rq](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client. ![](https://img.shields.io/github/last-commit/ddo/rq.svg)
* [sling](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests. ![](https://img.shields.io/github/last-commit/dghubble/sling.svg)

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). ![](https://img.shields.io/github/last-commit/go-gl/gl.svg)
* [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. ![](https://img.shields.io/github/last-commit/go-gl/glfw.svg)
* [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). ![](https://img.shields.io/github/last-commit/goxjs/gl.svg)
* [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. ![](https://img.shields.io/github/last-commit/goxjs/glfw.svg)
* [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM. ![](https://img.shields.io/github/last-commit/go-gl/mathgl.svg)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [go-firestorm](https://github.com/jschoedt/go-firestorm) - A simple ORM for Google/Firebase Cloud Firestore. ![](https://img.shields.io/github/last-commit/jschoedt/go-firestorm.svg)
* [go-pg](https://github.com/go-pg/pg) - PostgreSQL ORM with focus on PostgreSQL specific features and performance. ![](https://img.shields.io/github/last-commit/go-pg/pg.svg)
* [go-queryset](https://github.com/jirfag/go-queryset) - 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. ![](https://img.shields.io/github/last-commit/jirfag/go-queryset.svg)
* [go-sql](https://github.com/rushteam/gosql) - A easy ORM for mysql. ![](https://img.shields.io/github/last-commit/rushteam/gosql.svg)
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. ![](https://img.shields.io/github/last-commit/huandu/go-sqlbuilder.svg)
* [go-store](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go. ![](https://img.shields.io/github/last-commit/gosuri/go-store.svg)
* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. ![](https://img.shields.io/github/last-commit/jinzhu/gorm.svg)
* [gormt](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct. ![](https://img.shields.io/github/last-commit/xxjwxc/gormt.svg)
* [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. ![](https://img.shields.io/github/last-commit/go-gorp/gorp.svg)
* [grimoire](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). ![](https://img.shields.io/github/last-commit/Fs02/grimoire.svg)
* [lore](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. ![](https://img.shields.io/github/last-commit/abrahambotros/lore.svg)
* [Marlow](https://github.com/dadleyy/marlow) - Generated ORM from project structs for compile time safety assurances. ![](https://img.shields.io/github/last-commit/dadleyy/marlow.svg)
* [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. ![](https://img.shields.io/github/last-commit/gobuffalo/pop.svg)
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM. ![](https://img.shields.io/github/last-commit/coocood/qbs.svg)
* [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. ![](https://img.shields.io/github/last-commit/go-reform/reform.svg)
* [rel](https://github.com/Fs02/rel) - Golang SQL Repository Layer for Clean (Onion) Architecture. ![](https://img.shields.io/github/last-commit/Fs02/rel.svg)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. ![](https://img.shields.io/github/last-commit/volatiletech/sqlboiler.svg)
* [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. ![](https://img.shields.io/github/last-commit/upper/db.svg)
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go. ![](https://img.shields.io/github/last-commit/go-xorm/xorm.svg)
* [Zoom](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis. ![](https://img.shields.io/github/last-commit/albrow/zoom.svg)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) - Go dependency tool. ![](https://img.shields.io/github/last-commit/golang/dep.svg)
* [vgo](https://go.googlesource.com/vgo/) - Versioned Go.

*Unofficial libraries for package and dependency management.*

* [gigo](https://github.com/LyricalSecurity/gigo) - PIP-like dependency tool for golang, with support for private repositories and hashes. ![](https://img.shields.io/github/last-commit/LyricalSecurity/gigo.svg)
* [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. ![](https://img.shields.io/github/last-commit/Masterminds/glide.svg)
* [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. ![](https://img.shields.io/github/last-commit/tools/godep.svg)
* [gom](https://github.com/mattn/gom) - Go Manager - bundle for go. ![](https://img.shields.io/github/last-commit/mattn/gom.svg)
* [goop](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler. ![](https://img.shields.io/github/last-commit/nitrous-io/goop.svg)
* [gop](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH. ![](https://img.shields.io/github/last-commit/lunny/gop.svg)
* [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. ![](https://img.shields.io/github/last-commit/gpmgo/gopm.svg)
* [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. ![](https://img.shields.io/github/last-commit/kardianos/govendor.svg)
* [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. ![](https://img.shields.io/github/last-commit/pote/gpm.svg)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git. ![](https://img.shields.io/github/last-commit/VividCortex/johnny-deps.svg)
* [mvn-golang](https://github.com/raydac/mvn-golang) - plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. ![](https://img.shields.io/github/last-commit/raydac/mvn-golang.svg)
* [nut](https://github.com/jingweno/nut) - Vendor Go dependencies. ![](https://img.shields.io/github/last-commit/jingweno/nut.svg)
* [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments. ![](https://img.shields.io/github/last-commit/DamnWidget/VenGO.svg)

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. ![](https://img.shields.io/github/last-commit/jaegertracing/jaeger.svg)
* [pixie](https://github.com/pixie-labs/pixie) - No instrumentation tracing for Golang applications via eBPF. ![](https://img.shields.io/github/last-commit/pixie-labs/pixie.svg)
* [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. ![](https://img.shields.io/github/last-commit/pkg/profile.svg)
* [tracer](https://github.com/kamilsk/tracer) - Simple, lightweight tracing. ![](https://img.shields.io/github/last-commit/kamilsk/tracer.svg)

## Query Language

* [api-fu](https://github.com/ccbrown/api-fu) - Comprehensive GraphQL implementation. ![](https://img.shields.io/github/last-commit/ccbrown/api-fu.svg)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. ![](https://img.shields.io/github/last-commit/thedevsaddam/gojsonq.svg)
* [graphql](https://github.com/tmc/graphql) - graphql parser + utilities. ![](https://img.shields.io/github/last-commit/tmc/graphql.svg)
* [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. ![](https://img.shields.io/github/last-commit/neelance/graphql-go.svg)
* [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. ![](https://img.shields.io/github/last-commit/graphql-go/graphql.svg)
* [gws](https://github.com/Zaba505/gws) - Apollos' "GraphQL over Websocket" client and server implementation. ![](https://img.shields.io/github/last-commit/Zaba505/gws.svg)
* [jsonql](https://github.com/elgs/jsonql) - JSON query expression library in Golang. ![](https://img.shields.io/github/last-commit/elgs/jsonql.svg)
* [jsonslice](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters. ![](https://img.shields.io/github/last-commit/bhmj/jsonslice.svg)
* [rql](https://github.com/a8m/rql) - Resource Query Language for REST API. ![](https://img.shields.io/github/last-commit/a8m/rql.svg)
* [rqp](https://github.com/timsolov/rest-query-parser) - Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. ![](https://img.shields.io/github/last-commit/timsolov/rest-query-parser.svg)
* [straf](https://github.com/SonicRoshan/straf) - Easily Convert Golang structs to GraphQL objects. ![](https://img.shields.io/github/last-commit/SonicRoshan/straf.svg)

## Resource Embedding

* [esc](https://github.com/mjibson/esc) - Embeds files into Go programs and provides http.FileSystem interfaces to them. ![](https://img.shields.io/github/last-commit/mjibson/esc.svg)
* [fileb0x](https://github.com/UnnoTed/fileb0x) - Simple tool to embed files in go with focus on "customization" and ease to use. ![](https://img.shields.io/github/last-commit/UnnoTed/fileb0x.svg)
* [go-embed](https://github.com/pyros2097/go-embed) - Generates go code to embed resource files into your library or executable. ![](https://img.shields.io/github/last-commit/pyros2097/go-embed.svg)
* [go-resources](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go. ![](https://img.shields.io/github/last-commit/omeid/go-resources.svg)
* [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. ![](https://img.shields.io/github/last-commit/GeertJohan/go.rice.svg)
* [mule](https://github.com/wlbr/mule) - Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focussed on simplicity. ![](https://img.shields.io/github/last-commit/wlbr/mule.svg)
* [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries. ![](https://img.shields.io/github/last-commit/gobuffalo/packr.svg)
* [statics](https://github.com/go-playground/statics) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. ![](https://img.shields.io/github/last-commit/go-playground/statics.svg)
* [statik](https://github.com/rakyll/statik) - Embeds static files into a Go executable. ![](https://img.shields.io/github/last-commit/rakyll/statik.svg)
* [templify](https://github.com/wlbr/templify) - Embed external template files into Go code to create single file binaries. ![](https://img.shields.io/github/last-commit/wlbr/templify.svg)
* [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem. ![](https://img.shields.io/github/last-commit/shurcooL/vfsgen.svg)

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [assocentity](https://github.com/ndabAP/assocentity) - Package assocentity returns the average distance from words to a given entity. ![](https://img.shields.io/github/last-commit/ndabAP/assocentity.svg)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) - Provides a Bradley-Terry Model for pairwise comparisons. ![](https://img.shields.io/github/last-commit/seanhagen/bradleyterry.svg)
* [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) - Calendar heatmap in plain Go inspired by Github contribution activity. ![](https://img.shields.io/github/last-commit/nikolaydubina/calendarheatmap.svg)
* [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types.  ![](https://img.shields.io/github/last-commit/vdobler/chart.svg)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas). ![](https://img.shields.io/github/last-commit/rocketlaunchr/dataframe-go.svg)
* [evaler](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator. ![](https://img.shields.io/github/last-commit/soniah/evaler.svg)
* [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages. ![](https://img.shields.io/github/last-commit/VividCortex/ewma.svg)
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang. ![](https://img.shields.io/github/last-commit/skelterjohn/geom.svg)
* [go-dsp](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go. ![](https://img.shields.io/github/last-commit/mjibson/go-dsp.svg)
* [go-gt](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language. ![](https://img.shields.io/github/last-commit/ThePaw/go-gt.svg)
* [goent](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures. ![](https://img.shields.io/github/last-commit/kzahedi/goent.svg)
* [gohistogram](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams. ![](https://img.shields.io/github/last-commit/VividCortex/gohistogram.svg)
* [gonum](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. ![](https://img.shields.io/github/last-commit/gonum/gonum.svg)
* [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. ![](https://img.shields.io/github/last-commit/gonum/plot.svg)
* [goraph](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorith visualization). ![](https://img.shields.io/github/last-commit/gyuho/goraph.svg)
* [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. ![](https://img.shields.io/github/last-commit/cpmech/gosl.svg)
* [GoStats](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. ![](https://img.shields.io/github/last-commit/OGFris/GoStats.svg)
* [graph](https://github.com/yourbasic/graph) - Library of basic graph algorithms. ![](https://img.shields.io/github/last-commit/yourbasic/graph.svg)
* [ode](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. ![](https://img.shields.io/github/last-commit/ChristopherRabotin/ode.svg)
* [orb](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. ![](https://img.shields.io/github/last-commit/paulmach/orb.svg)
* [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go. ![](https://img.shields.io/github/last-commit/alixaxel/pagerank.svg)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library. ![](https://img.shields.io/github/last-commit/sgreben/piecewiselinear.svg)
* [PiHex](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. ![](https://img.shields.io/github/last-commit/claygod/PiHex.svg)
* [rootfinding](https://github.com/khezen/rootfinding) - root-finding algorithms library for finding roots of quadratic functions. ![](https://img.shields.io/github/last-commit/khezen/rootfinding.svg)
* [sparse](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. ![](https://img.shields.io/github/last-commit/james-bowman/sparse.svg)
* [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. ![](https://img.shields.io/github/last-commit/montanaflynn/stats.svg)
* [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. ![](https://img.shields.io/github/last-commit/nytlabs/streamtools.svg)
* [TextRank](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. ![](https://img.shields.io/github/last-commit/DavidBelicza/TextRank.svg)
* [triangolatte](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. ![](https://img.shields.io/github/last-commit/tchayen/triangolatte.svg)

## Security

*Libraries that are used to help make your application more secure.*

* [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. ![](https://img.shields.io/github/last-commit/hlandau/acme.svg)
* [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. ![](https://img.shields.io/github/last-commit/cossacklabs/acra.svg)
* [argon2-hashing](https://github.com/andskur/argon2-hashing) - light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. ![](https://img.shields.io/github/last-commit/andskur/argon2-hashing.svg)
* [argon2pw](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison. ![](https://img.shields.io/github/last-commit/raja/argon2pw.svg)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.
* [BadActor](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban. ![](https://img.shields.io/github/last-commit/jaredfolkins/badactor.svg)
* [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. ![](https://img.shields.io/github/last-commit/Ullaakut/cameradar.svg)
* [certificates](https://github.com/mvmaasakkers/certificates) - An opinionated tool for generating tls certificates. ![](https://img.shields.io/github/last-commit/mvmaasakkers/certificates.svg)
* [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) - A rest application to dynamically update firewalld rules on a linux server. ![](https://img.shields.io/github/last-commit/prashantgupta24/firewalld-rest.svg)
* [go-generate-password](https://github.com/m1/go-generate-password) - Password generator that can be used on the cli or as a library. ![](https://img.shields.io/github/last-commit/m1/go-generate-password.svg)
* [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". ![](https://img.shields.io/github/last-commit/hillu/go-yara.svg)
* [goArgonPass](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. ![](https://img.shields.io/github/last-commit/dwin/goArgonPass.svg)
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords. ![](https://img.shields.io/github/last-commit/dwin/goSecretBoxPassword.svg)
* [Interpol](https://bitbucket.org/vahidi/interpol) - Rule-based data generator for fuzzing and penetration testing.
* [lego](https://github.com/go-acme/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). ![](https://img.shields.io/github/last-commit/go-acme/lego.svg)
* [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. ![](https://img.shields.io/github/last-commit/awnumar/memguard.svg)
* [nacl](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's. ![](https://img.shields.io/github/last-commit/kevinburke/nacl.svg)
* [optimus-go](https://github.com/pjebs/optimus-go) - ID hashing and Obfuscation using Knuth's Algorithm. ![](https://img.shields.io/github/last-commit/pjebs/optimus-go.svg)
* [passlib](https://github.com/hlandau/passlib) - Futureproof password hashing library. ![](https://img.shields.io/github/last-commit/hlandau/passlib.svg)
* [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. ![](https://img.shields.io/github/last-commit/unrolled/secure.svg)
* [secureio](https://github.com/xaionaro-go/secureio) - An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. ![](https://img.shields.io/github/last-commit/xaionaro-go/secureio.svg)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in. ![](https://img.shields.io/github/last-commit/elithrar/simple-scrypt.svg)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys. ![](https://img.shields.io/github/last-commit/ssh-vault/ssh-vault.svg)
* [sslmgr](https://github.com/adrianosela/sslmgr) - SSL certificates made easy with a high level wrapper around acme/autocert. ![](https://img.shields.io/github/last-commit/adrianosela/sslmgr.svg)

## Serialization

*Libraries and tools for binary serialization.*

* [asn1](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang. ![](https://img.shields.io/github/last-commit/PromonLogicalis/asn1.svg)
* [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go. ![](https://img.shields.io/github/last-commit/glycerine/bambam.svg)
* [bel](https://github.com/32leaves/bel) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. ![](https://img.shields.io/github/last-commit/32leaves/bel.svg)
* [binstruct](https://github.com/ghostiam/binstruct) - Golang binary decoder for mapping data into the structure. ![](https://img.shields.io/github/last-commit/ghostiam/binstruct.svg)
* [cbor](https://github.com/fxamacker/cbor) - Small, safe, and easy CBOR encoding and decoding library. ![](https://img.shields.io/github/last-commit/fxamacker/cbor.svg)
* [colfer](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format. ![](https://img.shields.io/github/last-commit/pascaldekloe/colfer.svg)
* [csvutil](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures. ![](https://img.shields.io/github/last-commit/jszwec/csvutil.svg)
* [elastic](https://github.com/epiclabs-io/elastic) - Convert slices, maps or any other unknown value across different types at run-time, no matter what. ![](https://img.shields.io/github/last-commit/epiclabs-io/elastic.svg)
* [fixedwidth](https://github.com/huydang284/fixedwidth) - Fixed-width text formatting (UTF-8 supported). ![](https://img.shields.io/github/last-commit/huydang284/fixedwidth.svg)
* [fwencoder](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go. ![](https://img.shields.io/github/last-commit/o1egl/fwencoder.svg)
* [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. ![](https://img.shields.io/github/last-commit/glycerine/go-capnproto.svg)
* [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. ![](https://img.shields.io/github/last-commit/ugorji/go.svg)
* [go-lctree](https://github.com/sbourlon/go-lctree) - Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation). ![](https://img.shields.io/github/last-commit/sbourlon/go-lctree.svg)
* [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. ![](https://img.shields.io/github/last-commit/gogo/protobuf.svg)
* [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. ![](https://img.shields.io/github/last-commit/golang/protobuf.svg)
* [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". ![](https://img.shields.io/github/last-commit/json-iterator/go.svg)
* [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. ![](https://img.shields.io/github/last-commit/mitchellh/mapstructure.svg)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. ![](https://img.shields.io/github/last-commit/yvasiyarov/php_session_decoder.svg)
* [pletter](https://github.com/vimeda/pletter) - A standard way to wrap a proto message for message brokers. ![](https://img.shields.io/github/last-commit/vimeda/pletter.svg)
* [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures. ![](https://img.shields.io/github/last-commit/tuvistavie/structomap.svg)

## Server Applications

* [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. ![](https://img.shields.io/github/last-commit/xyproto/algernon.svg)
* [Caddy](https://github.com/mholt/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. ![](https://img.shields.io/github/last-commit/mholt/caddy.svg)
* [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
* [devd](https://github.com/cortesi/devd) - Local webserver for developers. ![](https://img.shields.io/github/last-commit/cortesi/devd.svg)
* [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. ![](https://img.shields.io/github/last-commit/Bilibili/discovery.svg)
* [dudeldu](https://github.com/krotik/dudeldu) - A simple SHOUTcast server. ![](https://img.shields.io/github/last-commit/krotik/dudeldu.svg)
* [etcd](https://github.com/coreos/etcd) - Highly-available key value store for shared configuration and service discovery. ![](https://img.shields.io/github/last-commit/coreos/etcd.svg)
* [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. ![](https://img.shields.io/github/last-commit/getfider/fider.svg)
* [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. ![](https://img.shields.io/github/last-commit/checkr/flagr.svg)
* [flipt](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js ![](https://img.shields.io/github/last-commit/markphelps/flipt.svg)
* [jackal](https://github.com/ortuman/jackal) - An XMPP server written in Go. ![](https://img.shields.io/github/last-commit/ortuman/jackal.svg)
* [lets-proxy2](https://github.com/rekby/lets-proxy2) - Reverse proxy for handle https with issue certificates in fly from lets-encrypt. ![](https://img.shields.io/github/last-commit/rekby/lets-proxy2.svg)
* [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. ![](https://img.shields.io/github/last-commit/minio/minio.svg)
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) - Nginx log parser and exporter to Prometheus. ![](https://img.shields.io/github/last-commit/blind-oracle/nginx-prometheus.svg)
* [nsq](http://nsq.io/) - A realtime distributed messaging platform.
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) - Stream database events from PostgreSQL to Kafka. ![](https://img.shields.io/github/last-commit/blind-oracle/psql-streamer.svg)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay) - Relay to load-balance Riemann events and/or convert them to Carbon. ![](https://img.shields.io/github/last-commit/blind-oracle/riemann-relay.svg)
* [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. ![](https://img.shields.io/github/last-commit/spiral/roadrunner.svg)
* [SFTPGo](https://github.com/drakkan/sftpgo) - Full featured and highly configurable SFTP server software. ![](https://img.shields.io/github/last-commit/drakkan/sftpgo.svg)
* [Trickster](https://github.com/tricksterproxy/trickster) - HTTP reverse proxy cache and time series accelerator. ![](https://img.shields.io/github/last-commit/tricksterproxy/trickster.svg)


## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) - Go stream processing library. ![](https://img.shields.io/github/last-commit/reugn/go-streams.svg)

## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace) - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. ![](https://img.shields.io/github/last-commit/yosssi/ace.svg)
* [amber](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. ![](https://img.shields.io/github/last-commit/eknkc/amber.svg)
* [damsel](https://github.com/dskinner/damsel) - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. ![](https://img.shields.io/github/last-commit/dskinner/damsel.svg)
* [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. ![](https://img.shields.io/github/last-commit/benbjohnson/ego.svg)
* [extemplate](https://github.com/dannyvankooten/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance. ![](https://img.shields.io/github/last-commit/dannyvankooten/extemplate.svg)
* [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). ![](https://img.shields.io/github/last-commit/valyala/fasttemplate.svg)
* [gofpdf](https://github.com/jung-kurt/gofpdf) - PDF document generator with high level support for text, drawing and images. ![](https://img.shields.io/github/last-commit/jung-kurt/gofpdf.svg)
* [gospin](https://github.com/m1/gospin) - Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. ![](https://img.shields.io/github/last-commit/m1/gospin.svg)
* [goview](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. ![](https://img.shields.io/github/last-commit/foolin/goview.svg)
* [hero](https://github.com/shiyanhui/hero) - Hero is a handy, fast and powerful go template engine. ![](https://img.shields.io/github/last-commit/shiyanhui/hero.svg)
* [jet](https://github.com/CloudyKit/jet) - Jet template engine. ![](https://img.shields.io/github/last-commit/CloudyKit/jet.svg)
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation. ![](https://img.shields.io/github/last-commit/ziutek/kasia.go.svg)
* [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates. ![](https://img.shields.io/github/last-commit/osteele/liquid.svg)
* [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. ![](https://img.shields.io/github/last-commit/johnfercher/maroto.svg)
* [mustache](https://github.com/hoisie/mustache) - Go implementation of the Mustache template language. ![](https://img.shields.io/github/last-commit/hoisie/mustache.svg)
* [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. ![](https://img.shields.io/github/last-commit/flosch/pongo2.svg)
* [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. ![](https://img.shields.io/github/last-commit/valyala/quicktemplate.svg)
* [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go. ![](https://img.shields.io/github/last-commit/aymerick/raymond.svg)
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang. ![](https://img.shields.io/github/last-commit/sipin/gorazor.svg)
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). ![](https://img.shields.io/github/last-commit/robfig/soy.svg)
* [velvet](https://github.com/gobuffalo/velvet) - Complete handlebars implementation in Go. ![](https://img.shields.io/github/last-commit/gobuffalo/velvet.svg)

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [assert](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. ![](https://img.shields.io/github/last-commit/go-playground/assert.svg)
    * [badio](https://github.com/cavaliercoder/badio) - Extensions to Go's `testing/iotest` package. ![](https://img.shields.io/github/last-commit/cavaliercoder/badio.svg)
    * [baloo](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy. ![](https://img.shields.io/github/last-commit/h2non/baloo.svg)
    * [biff](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible. ![](https://img.shields.io/github/last-commit/fulldump/biff.svg)
    * [charlatan](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests. ![](https://img.shields.io/github/last-commit/percolate/charlatan.svg)
    * [commander](https://github.com/SimonBaeumer/commander) - Tool for testing cli applications on windows, linux and osx. ![](https://img.shields.io/github/last-commit/SimonBaeumer/commander.svg)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework. ![](https://img.shields.io/github/last-commit/bradleyjkemp/cupaloy.svg)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby. ![](https://img.shields.io/github/last-commit/khaiql/dbcleaner.svg)
    * [dsunit](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files. ![](https://img.shields.io/github/last-commit/viant/dsunit.svg)
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. ![](https://img.shields.io/github/last-commit/fergusstrange/embedded-postgres.svg)
    * [endly](https://github.com/viant/endly) - Declarative end to end functional testing. ![](https://img.shields.io/github/last-commit/viant/endly.svg)
    * [flute](https://github.com/suzuki-shunsuke/flute) - HTTP client testing framework. ![](https://img.shields.io/github/last-commit/suzuki-shunsuke/flute.svg)
    * [frisby](https://github.com/verdverm/frisby) - REST API testing framework. ![](https://img.shields.io/github/last-commit/verdverm/frisby.svg)
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal. ![](https://img.shields.io/github/last-commit/msoap/go-carpet.svg)
    * [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests. ![](https://img.shields.io/github/last-commit/google/go-cmp.svg)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code. ![](https://img.shields.io/github/last-commit/zimmski/go-mutesting.svg)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package. ![](https://img.shields.io/github/last-commit/maxatome/go-testdeep.svg)
    * [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. ![](https://img.shields.io/github/last-commit/dnaeon/go-vcr.svg)
    * [goblin](https://github.com/franela/goblin) - Mocha like testing framework fo Go. ![](https://img.shields.io/github/last-commit/franela/goblin.svg)
    * [gocheck](http://labix.org/gocheck) - More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload.
    * [gocrest](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions. ![](https://img.shields.io/github/last-commit/corbym/gocrest.svg)
    * [godog](https://github.com/DATA-DOG/godog) - Cucumber or Behat like BDD framework for Go. ![](https://img.shields.io/github/last-commit/DATA-DOG/godog.svg)
    * [gofight](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework. ![](https://img.shields.io/github/last-commit/appleboy/gofight.svg)
    * [gogiven](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go. ![](https://img.shields.io/github/last-commit/corbym/gogiven.svg)
    * [gomatch](https://github.com/jfilipczyk/gomatch) - library created for testing JSON against patterns. ![](https://img.shields.io/github/last-commit/jfilipczyk/gomatch.svg)
    * [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
    * [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language. ![](https://img.shields.io/github/last-commit/orfjackal/gospec.svg)
    * [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. ![](https://img.shields.io/github/last-commit/stesla/gospecify.svg)
    * [gosuite](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. ![](https://img.shields.io/github/last-commit/pavlo/gosuite.svg)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns. ![](https://img.shields.io/github/last-commit/gotestyourself/gotest.tools.svg)
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. ![](https://img.shields.io/github/last-commit/rdrdr/hamcrest.svg)
    * [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. ![](https://img.shields.io/github/last-commit/gavv/httpexpect.svg)
    * [jsonassert](https://github.com/kinbiko/jsonassert) - Package for verifying that your JSON payloads are serialized correctly. ![](https://img.shields.io/github/last-commit/kinbiko/jsonassert.svg)
    * [restit](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test. ![](https://img.shields.io/github/last-commit/yookoala/restit.svg)
    * [schema](https://github.com/jgroeneveld/schema) - Quick and easy expression matching for JSON schemas used in requests and responses. ![](https://img.shields.io/github/last-commit/jgroeneveld/schema.svg)
    * [testcase](https://github.com/adamluzsi/testcase) - Idiomatic testing framework for Behavior Driven Development. ![](https://img.shields.io/github/last-commit/adamluzsi/testcase.svg)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. ![](https://img.shields.io/github/last-commit/go-testfixtures/testfixtures.svg)
    * [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. ![](https://img.shields.io/github/last-commit/stretchr/testify.svg)
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd) - Convert markdown snippets into testable go code.
    * [testsql](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished. ![](https://img.shields.io/github/last-commit/zhulongcheng/testsql.svg)
    * [trial](https://github.com/jgroeneveld/trial) - Quick and easy extendable assertions without introducing much boilerplate. ![](https://img.shields.io/github/last-commit/jgroeneveld/trial.svg)
    * [Tt](https://github.com/vcaesar/tt) - Simple and colorful test tools. ![](https://img.shields.io/github/last-commit/vcaesar/tt.svg)
    * [wstest](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler. ![](https://img.shields.io/github/last-commit/posener/wstest.svg)

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. ![](https://img.shields.io/github/last-commit/maxbrunsfeld/counterfeiter.svg)
    * [go-localstack](https://github.com/elgohr/go-localstack) - Tool for using localstack in AWS testing. ![](https://img.shields.io/github/last-commit/elgohr/go-localstack.svg)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. ![](https://img.shields.io/github/last-commit/DATA-DOG/go-sqlmock.svg)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes. ![](https://img.shields.io/github/last-commit/DATA-DOG/go-txdb.svg)
    * [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. ![](https://img.shields.io/github/last-commit/h2non/gock.svg)
    * [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language. ![](https://img.shields.io/github/last-commit/golang/mock.svg)
    * [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing. ![](https://img.shields.io/github/last-commit/seborama/govcr.svg)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. ![](https://img.shields.io/github/last-commit/SpectoLabs/hoverfly.svg)
    * [httpmock](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources. ![](https://img.shields.io/github/last-commit/jarcoal/httpmock.svg)
    * [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces. ![](https://img.shields.io/github/last-commit/gojuno/minimock.svg)
    * [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter. ![](https://img.shields.io/github/last-commit/tv42/mockhttp.svg)
    * [timex](https://github.com/cabify/timex) - A test-friendly replacement for the native `time` package. ![](https://img.shields.io/github/last-commit/cabify/timex.svg)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. ![](https://img.shields.io/github/last-commit/dvyukov/go-fuzz.svg)
    * [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. ![](https://img.shields.io/github/last-commit/google/gofuzz.svg)
    * [Tavor](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework. ![](https://img.shields.io/github/last-commit/zimmski/tavor.svg)

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. ![](https://img.shields.io/github/last-commit/mafredri/cdp.svg)
    * [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. ![](https://img.shields.io/github/last-commit/knq/chromedp.svg)
    * [ggr](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. ![](https://img.shields.io/github/last-commit/aerokube/ggr.svg)
    * [rod](https://github.com/go-rod/rod) - A Devtools driver to make web automation and scraping easy. ![](https://img.shields.io/github/last-commit/go-rod/rod.svg)
    * [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. ![](https://img.shields.io/github/last-commit/aerokube/selenoid.svg)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) - An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. ![](https://img.shields.io/github/last-commit/pingcap/failpoint.svg)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align) - A general purpose application that aligns text. ![](https://img.shields.io/github/last-commit/Guitarbum722/align.svg)
    * [allot](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots. ![](https://img.shields.io/github/last-commit/sbstjn/allot.svg)
    * [bbConvert](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags. ![](https://img.shields.io/github/last-commit/CalebQ42/bbConvert.svg)
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. ![](https://img.shields.io/github/last-commit/russross/blackfriday.svg)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. ![](https://img.shields.io/github/last-commit/microcosm-cc/bluemonday.svg)
    * [codetree](https://github.com/aerogo/codetree) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. ![](https://img.shields.io/github/last-commit/aerogo/codetree.svg)
    * [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers. ![](https://img.shields.io/github/last-commit/asciimoo/colly.svg)
    * [commonregex](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go. ![](https://img.shields.io/github/last-commit/mingrammer/commonregex.svg)
    * [dataflowkit](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data. ![](https://img.shields.io/github/last-commit/slotix/dataflowkit.svg)
    * [did](https://github.com/ockam-network/did) - DID (Decentralized Identifiers) Parser and Stringer in Go. ![](https://img.shields.io/github/last-commit/ockam-network/did.svg)
    * [doi](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go. ![](https://img.shields.io/github/last-commit/hscells/doi.svg)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go. ![](https://img.shields.io/github/last-commit/editorconfig/editorconfig-core-go.svg)
    * [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). ![](https://img.shields.io/github/last-commit/endeveit/enca.svg)
    * [encdec](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decodersa. ![](https://img.shields.io/github/last-commit/mickep76/encdec.svg)
    * [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings. ![](https://img.shields.io/github/last-commit/alixaxel/genex.svg)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection). ![](https://img.shields.io/github/last-commit/ianlopshire/go-fixedwidth.svg)
    * [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. ![](https://img.shields.io/github/last-commit/dustin/go-humanize.svg)
    * [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language. ![](https://img.shields.io/github/last-commit/adrianmo/go-nmea.svg)
    * [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string. ![](https://img.shields.io/github/last-commit/mattn/go-runewidth.svg)
    * [go-slugify](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support. ![](https://img.shields.io/github/last-commit/mozillazg/go-slugify.svg)
    * [go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools. ![](https://img.shields.io/github/last-commit/pelletier/go-toml.svg)
    * [go-vcard](https://github.com/emersion/go-vcard) - Parse and format vCard. ![](https://img.shields.io/github/last-commit/emersion/go-vcard.svg)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) - Zero-width character detection and removal for Go. ![](https://img.shields.io/github/last-commit/trubitsyn/go-zero-width.svg)
    * [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. ![](https://img.shields.io/github/last-commit/mmcdole/gofeed.svg)
    * [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language. ![](https://img.shields.io/github/last-commit/awalterschulze/gographviz.svg)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - Format bytes to string.
    * [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts. ![](https://img.shields.io/github/last-commit/polera/gonameparts.svg)
    * [goq](https://github.com/andrewstuart/goq) - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). ![](https://img.shields.io/github/last-commit/andrewstuart/goq.svg)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. ![](https://img.shields.io/github/last-commit/PuerkitoBio/goquery.svg)
    * [goregen](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions. ![](https://img.shields.io/github/last-commit/zach-klippenstein/goregen.svg)
    * [goribot](https://github.com/zhshch2002/goribot) - A simple golang spider/scraping framework,build a spider in 3 lines. ![](https://img.shields.io/github/last-commit/zhshch2002/goribot.svg)
    * [gotext](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go. ![](https://img.shields.io/github/last-commit/leonelquinteros/gotext.svg)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text. ![](https://img.shields.io/github/last-commit/endeveit/guesslanguage.svg)
    * [htmlquery](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. ![](https://img.shields.io/github/last-commit/antchfx/htmlquery.svg)
    * [inject](https://github.com/facebookgo/inject) - Package inject provides a reflect based injector. ![](https://img.shields.io/github/last-commit/facebookgo/inject.svg)
    * [ltsv](https://github.com/Wing924/ltsv) - High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. ![](https://img.shields.io/github/last-commit/Wing924/ltsv.svg)
    * [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. ![](https://img.shields.io/github/last-commit/clbanning/mxj.svg)
    * [pagser](https://github.com/foolin/pagser) - Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. ![](https://img.shields.io/github/last-commit/foolin/pagser.svg)
    * [podcast](https://github.com/eduncan911/podcast) - iTunes Compliant and RSS 2.0 Podcast Generator in Golang ![](https://img.shields.io/github/last-commit/eduncan911/podcast.svg)
    * [sdp](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. ![](https://img.shields.io/github/last-commit/gortc/sdp.svg)
    * [sh](https://github.com/mvdan/sh) - Shell parser and formatter. ![](https://img.shields.io/github/last-commit/mvdan/sh.svg)
    * [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. ![](https://img.shields.io/github/last-commit/gosimple/slug.svg)
    * [Slugify](https://github.com/avelino/slugify) - Go slugify application that handles string. ![](https://img.shields.io/github/last-commit/avelino/slugify.svg)
    * [syndfeed](https://github.com/zhengchun/syndfeed) - A syndication feed for Atom 1.0 and RSS 2.0. ![](https://img.shields.io/github/last-commit/zhengchun/syndfeed.svg)
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). ![](https://img.shields.io/github/last-commit/BurntSushi/toml.svg)
* Utility
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) - A sanitization-based swear filter for Go. ![](https://img.shields.io/github/last-commit/JoshuaDoes/gofuckyourself.svg)
    * [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go. ![](https://img.shields.io/github/last-commit/bndr/gotabulate.svg)
    * [kace](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. ![](https://img.shields.io/github/last-commit/codemodus/kace.svg)
    * [parseargs-go](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes. ![](https://img.shields.io/github/last-commit/nproc/parseargs-go.svg)
    * [parth](https://github.com/codemodus/parth) - URL path segmentation parsing. ![](https://img.shields.io/github/last-commit/codemodus/parth.svg)
    * [radix](https://github.com/yourbasic/radix) - fast string sorting algorithm. ![](https://img.shields.io/github/last-commit/yourbasic/radix.svg)
    * [Tagify](https://github.com/zoomio/tagify) - Produces a set of tags from given source. ![](https://img.shields.io/github/last-commit/zoomio/tagify.svg)
    * [textwrap](https://github.com/isbm/textwrap) - Implementation of `textwrap` module from Python. ![](https://img.shields.io/github/last-commit/isbm/textwrap.svg)
    * [TySug](https://github.com/Dynom/TySug) - Alternative suggestions with respect to keyboard layouts. ![](https://img.shields.io/github/last-commit/Dynom/TySug.svg)
    * [xj2go](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct. ![](https://img.shields.io/github/last-commit/stackerzzq/xj2go.svg)
    * [xurls](https://github.com/mvdan/xurls) - Extract urls from text. ![](https://img.shields.io/github/last-commit/mvdan/xurls.svg)

## Third-party APIs

*Libraries for accessing third party APIs.*

* [airtable](https://github.com/mehanizm/airtable) - Go client library for the [Airtable API](https://airtable.com/api). ![](https://img.shields.io/github/last-commit/mehanizm/airtable.svg)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). ![](https://img.shields.io/github/last-commit/ngs/go-amazon-product-advertising-api.svg)
* [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. ![](https://img.shields.io/github/last-commit/ChimeraCoder/anaconda.svg)
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) - The official AWS SDK for the Go programming language. ![](https://img.shields.io/github/last-commit/aws/aws-sdk-go.svg)
* [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API. ![](https://img.shields.io/github/last-commit/naegelejd/brewerydb.svg)
* [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/). ![](https://img.shields.io/github/last-commit/andygrunwald/cachet.svg)
* [circleci](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API. ![](https://img.shields.io/github/last-commit/jszwedko/go-circleci.svg)
* [clarifai](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API. ![](https://img.shields.io/github/last-commit/samuelcouch/clarifai.svg)
* [codeship-go](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2. ![](https://img.shields.io/github/last-commit/codeship/codeship-go.svg)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API. ![](https://img.shields.io/github/last-commit/coinpaprika/coinpaprika-api-go-client.svg)
* [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. ![](https://img.shields.io/github/last-commit/bwmarrin/discordgo.svg)
* [ethrpc](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API. ![](https://img.shields.io/github/last-commit/onrik/ethrpc.svg)
* [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. ![](https://img.shields.io/github/last-commit/huandu/facebook.svg)
* [fcm](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging. ![](https://img.shields.io/github/last-commit/maddevsio/fcm.svg)
* [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API. ![](https://img.shields.io/github/last-commit/emiddleton/gads.svg)
* [gami](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface. ![](https://img.shields.io/github/last-commit/bit4bit/gami.svg)
* [gcm](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging. ![](https://img.shields.io/github/last-commit/Aorioli/gcm.svg)
* [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. ![](https://img.shields.io/github/last-commit/codingsince1985/geo-golang.svg)
* [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. ![](https://img.shields.io/github/last-commit/google/go-github.svg)
* [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. ![](https://img.shields.io/github/last-commit/shurcooL/githubql.svg)
* [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler ![](https://img.shields.io/github/last-commit/axelspringer/go-chronos.svg)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API. ![](https://img.shields.io/github/last-commit/PaulRosset/go-hacknews.svg)
* [go-here](https://github.com/abdullahselek/go-here) - Go client library around the HERE location based APIs. ![](https://img.shields.io/github/last-commit/abdullahselek/go-here.svg)
* [go-imgur](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com) ![](https://img.shields.io/github/last-commit/koffeinsource/go-imgur.svg)
* [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) ![](https://img.shields.io/github/last-commit/andygrunwald/go-jira.svg)
* [go-marathon](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS. ![](https://img.shields.io/github/last-commit/gambol99/go-marathon.svg)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). ![](https://img.shields.io/github/last-commit/nstratos/go-myanimelist.svg)
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection) - Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). ![](https://img.shields.io/github/last-commit/rbretecher/go-postman-collection.svg)
* [go-sophos](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. ![](https://img.shields.io/github/last-commit/esurdam/go-sophos.svg)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) - Go client library for the SPTrans Olho Vivo API. ![](https://img.shields.io/github/last-commit/sergioaugrod/go-sptrans.svg)
* [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.
* [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. ![](https://img.shields.io/github/last-commit/andygrunwald/go-trending.svg)
* [go-twitch](https://github.com/knspriggs/go-twitch) - Go client for interacting with the Twitch v3 API. ![](https://img.shields.io/github/last-commit/knspriggs/go-twitch.svg)
* [go-twitter](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs. ![](https://img.shields.io/github/last-commit/dghubble/go-twitter.svg)
* [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API. ![](https://img.shields.io/github/last-commit/hbagdi/go-unsplash.svg)
* [go-xkcd](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API. ![](https://img.shields.io/github/last-commit/nishanths/go-xkcd.svg)
* [gogtrends](https://github.com/groovili/gogtrends) - Google Trends Unofficial API. ![](https://img.shields.io/github/last-commit/groovili/gogtrends.svg)
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb) - Golang wrapper for The Movie Database API v3. ![](https://img.shields.io/github/last-commit/cyruzin/golang-tmdb.svg)
* [golyrics](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website. ![](https://img.shields.io/github/last-commit/mamal72/golyrics.svg)
* [gomalshare](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](http://www.malshare.com/) ![](https://img.shields.io/github/last-commit/MonaxGT/gomalshare.svg)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library. ![](https://img.shields.io/github/last-commit/michiwend/gomusicbrainz.svg)
* [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. ![](https://img.shields.io/github/last-commit/google/google-api-go-client.svg)
* [google-analytics](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting. ![](https://img.shields.io/github/last-commit/chonthu/go-google-analytics.svg)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. ![](https://img.shields.io/github/last-commit/GoogleCloudPlatform/gcloud-golang.svg)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). ![](https://img.shields.io/github/last-commit/ngs/go-google-email-audit-api.svg)
* [google-play-scraper](https://github.com/n0madic/google-play-scraper) - Get data from Google Play Store. ![](https://img.shields.io/github/last-commit/n0madic/google-play-scraper.svg)
* [gopaapi5](https://github.com/utekaravinash/gopaapi5) - Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). ![](https://img.shields.io/github/last-commit/utekaravinash/gopaapi5.svg)
* [gosip](https://github.com/koltyakov/gosip) - Go client library SharePoint API. ![](https://img.shields.io/github/last-commit/koltyakov/gosip.svg)
* [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. ![](https://img.shields.io/github/last-commit/jsgilmore/gostorm.svg)
* [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. ![](https://img.shields.io/github/last-commit/andybons/hipchat.svg)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP. ![](https://img.shields.io/github/last-commit/daneharrigan/hipchat.svg)
* [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/). ![](https://img.shields.io/github/last-commit/Henry-Sarabia/igdb.svg)
* [lastpass-go](https://github.com/ansd/lastpass-go) - Go client library for the [LastPass](https://www.lastpass.com/) API. ![](https://img.shields.io/github/last-commit/ansd/lastpass-go.svg)
* [libgoffi](https://github.com/clevabit/libgoffi) - Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration ![](https://img.shields.io/github/last-commit/clevabit/libgoffi.svg)
* [Medium](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API. ![](https://img.shields.io/github/last-commit/Medium/medium-sdk-go.svg)
* [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. ![](https://img.shields.io/github/last-commit/andygrunwald/megos.svg)
* [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. ![](https://img.shields.io/github/last-commit/minio/minio-go.svg)
* [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. ![](https://img.shields.io/github/last-commit/dukex/mixpanel.svg)
* [patreon-go](https://github.com/mxpv/patreon-go) - Go library for Patreon API. ![](https://img.shields.io/github/last-commit/mxpv/patreon-go.svg)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API. ![](https://img.shields.io/github/last-commit/logpacker/PayPal-Go-SDK.svg)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK. ![](https://img.shields.io/github/last-commit/playlyfe/playlyfe-go-sdk.svg)
* [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API. ![](https://img.shields.io/github/last-commit/gregdel/pushover.svg)
* [rrdaclient](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. ![](https://img.shields.io/github/last-commit/Omie/rrdaclient.svg)
* [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API. ![](https://img.shields.io/github/last-commit/rapito/go-shopify.svg)
* [simples3](https://github.com/rhnvrm/simples3) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go. ![](https://img.shields.io/github/last-commit/rhnvrm/simples3.svg)
* [slack](https://github.com/nlopes/slack) - Slack API in Go. ![](https://img.shields.io/github/last-commit/nlopes/slack.svg)
* [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API. ![](https://img.shields.io/github/last-commit/sergiotapia/smitego.svg)
* [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API. ![](https://img.shields.io/github/last-commit/rapito/go-spotify.svg)
* [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers. ![](https://img.shields.io/github/last-commit/sostronk/go-steam.svg)
* [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. ![](https://img.shields.io/github/last-commit/stripe/stripe-go.svg)
* [textbelt](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API. ![](https://img.shields.io/github/last-commit/dietsche/textbelt.svg)
* [translate](https://github.com/poorny/translate) - Go online translation package. ![](https://img.shields.io/github/last-commit/poorny/translate.svg)
* [Trello](https://github.com/adlio/trello) - Go wrapper for the Trello API. ![](https://img.shields.io/github/last-commit/adlio/trello.svg)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) - Go wrapper for the TripAdvisor API. ![](https://img.shields.io/github/last-commit/mrbenosborne/tripadvisor-golang.svg)
* [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API. ![](https://img.shields.io/github/last-commit/mattcunningham/gumblr.svg)
* [twitter-scraper](https://github.com/n0madic/twitter-scraper) - Scrape the Twitter Frontend API without authentication and limits. ![](https://img.shields.io/github/last-commit/n0madic/twitter-scraper.svg)
* [uptimerobot](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API. ![](https://img.shields.io/github/last-commit/bitfield/uptimerobot.svg)
* [vl-go](https://github.com/verifid/vl-go) - Go client library around the VerifID identity verification layer API. ![](https://img.shields.io/github/last-commit/verifid/vl-go.svg)
* [webhooks](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket. ![](https://img.shields.io/github/last-commit/go-playground/webhooks.svg)
* [wit-go](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API. ![](https://img.shields.io/github/last-commit/wit-ai/wit-go.svg)
* [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API. ![](https://img.shields.io/github/last-commit/brunomvsouza/ynab.go.svg)
* [zooz](https://github.com/gojuno/go-zooz) - Go client for the Zooz API. ![](https://img.shields.io/github/last-commit/gojuno/go-zooz.svg)

## Utilities

*General utilities and tools to make your life easier.*

* [apm](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API. ![](https://img.shields.io/github/last-commit/topfreegames/apm.svg)
* [backscanner](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. ![](https://img.shields.io/github/last-commit/icza/backscanner.svg)
* [beyond](https://github.com/wesovilabs/beyond) - The Go tool that will drive you to the AOP world! ![](https://img.shields.io/github/last-commit/wesovilabs/beyond.svg)
* [blank](https://github.com/Henry-Sarabia/blank) - Verify or remove blanks and whitespace from strings. ![](https://img.shields.io/github/last-commit/Henry-Sarabia/blank.svg)
* [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. ![](https://img.shields.io/github/last-commit/tmrts/boilr.svg)
* [chyle](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities. ![](https://img.shields.io/github/last-commit/antham/chyle.svg)
* [circuit](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. ![](https://img.shields.io/github/last-commit/cep21/circuit.svg)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. ![](https://img.shields.io/github/last-commit/rubyist/circuitbreaker.svg)
* [clockwork](https://github.com/jonboulle/clockwork) - A simple fake clock for golang. ![](https://img.shields.io/github/last-commit/jonboulle/clockwork.svg)
* [cmd](https://github.com/SimonBaeumer/cmd) - Library for executing shell commands on osx, windows and linux. ![](https://img.shields.io/github/last-commit/SimonBaeumer/cmd.svg)
* [command](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher. ![](https://img.shields.io/github/last-commit/txgruppi/command.svg)
* [copy-pasta](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage. ![](https://img.shields.io/github/last-commit/jutkko/copy-pasta.svg)
* [countries](https://github.com/biter777/countries) - Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts. ![](https://img.shields.io/github/last-commit/biter777/countries.svg)
* [create-go-app](https://github.com/create-go-app/cli) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. ![](https://img.shields.io/github/last-commit/create-go-app/cli.svg)
* [ctop](https://github.com/bcicen/ctop) - [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. ![](https://img.shields.io/github/last-commit/bcicen/ctop.svg)
* [ctxutil](https://github.com/posener/ctxutil) - A collection of utility functions for contexts. ![](https://img.shields.io/github/last-commit/posener/ctxutil.svg)
* [dbt](https://github.com/nikogura/dbt) - A framework for running self-updating signed binaries from a central, trusted repository. ![](https://img.shields.io/github/last-commit/nikogura/dbt.svg)
* [Death](https://github.com/vrecan/death) - Managing go application shutdown with signals. ![](https://img.shields.io/github/last-commit/vrecan/death.svg)
* [Deepcopier](https://github.com/ulule/deepcopier) - Simple struct copying for Go. ![](https://img.shields.io/github/last-commit/ulule/deepcopier.svg)
* [delve](https://github.com/derekparker/delve) - Go debugger. ![](https://img.shields.io/github/last-commit/derekparker/delve.svg)
* [dlog](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls. ![](https://img.shields.io/github/last-commit/kirillDanshin/dlog.svg)
* [equalizer](https://github.com/reugn/equalizer) - Quota manager and rate limiter collection for Go. ![](https://img.shields.io/github/last-commit/reugn/equalizer.svg)
* [ergo](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy. ![](https://img.shields.io/github/last-commit/cristianoliveira/ergo.svg)
* [evaluator](https://github.com/nullne/evaluator) - Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. ![](https://img.shields.io/github/last-commit/nullne/evaluator.svg)
* [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. ![](https://img.shields.io/github/last-commit/h2non/filetype.svg)
* [filler](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag. ![](https://img.shields.io/github/last-commit/yaronsumel/filler.svg)
* [filter](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data. ![](https://img.shields.io/github/last-commit/gookit/filter.svg)
* [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. ![](https://img.shields.io/github/last-commit/junegunn/fzf.svg)
* [gaper](https://github.com/maxcnunes/gaper) - Builds and restarts a Go project when it crashes or some watched file changes. ![](https://img.shields.io/github/last-commit/maxcnunes/gaper.svg)
* [generate](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex. ![](https://img.shields.io/github/last-commit/go-playground/generate.svg)
* [ghokin](https://github.com/antham/ghokin) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). ![](https://img.shields.io/github/last-commit/antham/ghokin.svg)
* [git-time-metric](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git. ![](https://img.shields.io/github/last-commit/git-time-metric/gtm.svg)
* [go-astitodo](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code. ![](https://img.shields.io/github/last-commit/asticode/go-astitodo.svg)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only). ![](https://img.shields.io/github/last-commit/wendigo/go-bind-plugin.svg)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) - Pure Go bsdiff and bspatch libraries and CLI tools. ![](https://img.shields.io/github/last-commit/gabstv/go-bsdiff.svg)
* [go-convert](https://github.com/Eun/go-convert) - Package go-convert enbles you to convert a value into another type. ![](https://img.shields.io/github/last-commit/Eun/go-convert.svg)
* [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go. ![](https://img.shields.io/github/last-commit/ungerik/go-dry.svg)
* [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). ![](https://img.shields.io/github/last-commit/thoas/go-funk.svg)
* [go-health](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services. ![](https://img.shields.io/github/last-commit/Talento90/go-health.svg)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields. ![](https://img.shields.io/github/last-commit/mozillazg/go-httpheader.svg)
* [go-lock](https://github.com/viney-shih/go-lock) - go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. ![](https://img.shields.io/github/last-commit/viney-shih/go-lock.svg)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) - Go package for working with Problem Details. ![](https://img.shields.io/github/last-commit/mvmaasakkers/go-problemdetails.svg)
* [go-rate](https://github.com/beefsack/go-rate) - Timed rate limiter for Go. ![](https://img.shields.io/github/last-commit/beefsack/go-rate.svg)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go. ![](https://img.shields.io/github/last-commit/ikeikeikeike/go-sitemap-generator.svg)
* [go-trigger](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. ![](https://img.shields.io/github/last-commit/sadlil/go-trigger.svg)
* [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package. ![](https://img.shields.io/github/last-commit/carlescere/goback.svg)
* [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons. ![](https://img.shields.io/github/last-commit/VividCortex/godaemon.svg)
* [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. ![](https://img.shields.io/github/last-commit/dropbox/godropbox.svg)
* [gohper](https://github.com/cosiner/gohper) - Various tools/modules help for development. ![](https://img.shields.io/github/last-commit/cosiner/gohper.svg)
* [golarm](https://github.com/msempere/golarm) - Fire alarms with system events. ![](https://img.shields.io/github/last-commit/msempere/golarm.svg)
* [golog](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks. ![](https://img.shields.io/github/last-commit/mlimaloureiro/golog.svg)
* [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs. ![](https://img.shields.io/github/last-commit/bndr/gopencils.svg)
* [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images. ![](https://img.shields.io/github/last-commit/michiwend/goplaceholder.svg)
* [goreadability](https://github.com/philipjkim/goreadability) - Webpage summary extractor using Facebook Open Graph and arc90's readability. ![](https://img.shields.io/github/last-commit/philipjkim/goreadability.svg)
* [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. ![](https://img.shields.io/github/last-commit/goreleaser/goreleaser.svg)
* [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. ![](https://img.shields.io/github/last-commit/wgliang/goreporter.svg)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features. ![](https://img.shields.io/github/last-commit/linxGnu/goseaweedfs.svg)
* [gostrutils](https://github.com/ik5/gostrutils) - Collections of string manipulation and conversion functions. ![](https://img.shields.io/github/last-commit/ik5/gostrutils.svg)
* [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go. ![](https://img.shields.io/github/last-commit/subosito/gotenv.svg)
* [gpath](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection. ![](https://img.shields.io/github/last-commit/tenntenn/gpath.svg)
* [gubrak](https://github.com/novalagung/gubrak) - Golang utility library with syntactic sugar. It's like lodash, but for golang. ![](https://img.shields.io/github/last-commit/novalagung/gubrak.svg)
* [handy](https://github.com/miguelpragier/handy) - Many utilities and helpers like string handlers/formatters and validators. ![](https://img.shields.io/github/last-commit/miguelpragier/handy.svg)
* [hostctl](https://github.com/guumaster/hostctl) - A CLI tool to manage /etc/hosts with easy commands. ![](https://img.shields.io/github/last-commit/guumaster/hostctl.svg)
* [htcat](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility. ![](https://img.shields.io/github/last-commit/htcat/htcat.svg)
* [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. ![](https://img.shields.io/github/last-commit/github/hub.svg)
* [hystrix-go](https://github.com/afex/hystrix-go) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. ![](https://img.shields.io/github/last-commit/afex/hystrix-go.svg)
* [immortal](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor. ![](https://img.shields.io/github/last-commit/immortal/immortal.svg)
* [intrinsic](https://github.com/mengzhuo/intrinsic) - Use x86 SIMD without writing any assembly code. ![](https://img.shields.io/github/last-commit/mengzhuo/intrinsic.svg)
* [jsend](https://github.com/clevergo/jsend) - JSend's implementation writen in Go. ![](https://img.shields.io/github/last-commit/clevergo/jsend.svg)
* [jump](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits. ![](https://img.shields.io/github/last-commit/gsamokovarov/jump.svg)
* [koazee](https://github.com/wesovilabs/koazee) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. ![](https://img.shields.io/github/last-commit/wesovilabs/koazee.svg)
* [limiters](https://github.com/mennanov/limiters) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. ![](https://img.shields.io/github/last-commit/mennanov/limiters.svg)
* [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go. ![](https://img.shields.io/github/last-commit/jaschaephraim/lrserver.svg)
* [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. ![](https://img.shields.io/github/last-commit/minio/mc.svg)
* [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. ![](https://img.shields.io/github/last-commit/imdario/mergo.svg)
* [mimemagic](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility. ![](https://img.shields.io/github/last-commit/zRedShift/mimemagic.svg)
* [mimesniffer](https://github.com/aofei/mimesniffer) - A MIME type sniffer for Go. ![](https://img.shields.io/github/last-commit/aofei/mimesniffer.svg)
* [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. ![](https://img.shields.io/github/last-commit/gabriel-vasile/mimetype.svg)
* [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. ![](https://img.shields.io/github/last-commit/tdewolff/minify.svg)
* [minquery](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). ![](https://img.shields.io/github/last-commit/icza/minquery.svg)
* [mmake](https://github.com/tj/mmake) - Modern Make. ![](https://img.shields.io/github/last-commit/tj/mmake.svg)
* [moldova](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template. ![](https://img.shields.io/github/last-commit/StabbyCutyou/moldova.svg)
* [mole](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels. ![](https://img.shields.io/github/last-commit/davrodpin/mole.svg)
* [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) - Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines. ![](https://img.shields.io/github/last-commit/gobeam/mongo-go-pagination.svg)
* [mssqlx](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. ![](https://img.shields.io/github/last-commit/linxGnu/mssqlx.svg)
* [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers. ![](https://img.shields.io/github/last-commit/VividCortex/multitick.svg)
* [myhttp](https://github.com/inancgumus/myhttp) - Simple API to make HTTP GET requests with timeout support. ![](https://img.shields.io/github/last-commit/inancgumus/myhttp.svg)
* [netbug](https://github.com/e-dard/netbug) - Easy remote profiling of your services. ![](https://img.shields.io/github/last-commit/e-dard/netbug.svg)
* [nfdump](https://github.com/chrispassas/nfdump) - Read nfdump netflow files. ![](https://img.shields.io/github/last-commit/chrispassas/nfdump.svg)
* [nostromo](https://github.com/pokanop/nostromo) - CLI for building powerful aliases. ![](https://img.shields.io/github/last-commit/pokanop/nostromo.svg)
* [okrun](https://github.com/xta/okrun) - go run error steamroller. ![](https://img.shields.io/github/last-commit/xta/okrun.svg)
* [olaf](https://github.com/btnguyen2k/olaf) - Twitter Snowflake implemented in Go. ![](https://img.shields.io/github/last-commit/btnguyen2k/olaf.svg)
* [onecache](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). ![](https://img.shields.io/github/last-commit/adelowo/onecache.svg)
* [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. ![](https://img.shields.io/github/last-commit/maruel/panicparse.svg)
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) - Pattern matching libray. ![](https://img.shields.io/github/last-commit/alexpantyukhin/go-pattern-match.svg)
* [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. ![](https://img.shields.io/github/last-commit/peco/peco.svg)
* [pgo](https://github.com/arthurkushman/pgo) - Convenient functions for PHP community. ![](https://img.shields.io/github/last-commit/arthurkushman/pgo.svg)
* [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API. ![](https://img.shields.io/github/last-commit/VividCortex/pm.svg)
* [ptr](https://github.com/gotidy/ptr) - Package that provide functions for simplified creation of pointers from constants of basic types. ![](https://img.shields.io/github/last-commit/gotidy/ptr.svg)
* [r](https://github.com/is5/r) - Python-like `range()` experience for Go. ![](https://img.shields.io/github/last-commit/github.com/is5/r.svg)
* [rclient](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs. ![](https://img.shields.io/github/last-commit/zpatrick/rclient.svg)
* [realize](https://github.com/tockins/realize) - Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. ![](https://img.shields.io/github/last-commit/tockins/realize.svg)
* [repeat](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating. ![](https://img.shields.io/github/last-commit/ssgreg/repeat.svg)
* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. ![](https://img.shields.io/github/last-commit/mozillazg/request.svg)
* [rerate](https://github.com/abo/rerate) - Redis-based rate counter and rate limiter for Go. ![](https://img.shields.io/github/last-commit/abo/rerate.svg)
* [rerun](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes. ![](https://img.shields.io/github/last-commit/ivpusic/rerun.svg)
* [rest-go](https://github.com/edermanoel94/rest-go) - A package that provide many helpful methods for working with rest api. ![](https://img.shields.io/github/last-commit/edermanoel94/rest-go.svg)
* [retry](https://github.com/kamilsk/retry) - The most advanced functional mechanism to perform actions repetitively until successful. ![](https://img.shields.io/github/last-commit/kamilsk/retry.svg)
* [retry](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go. ![](https://img.shields.io/github/last-commit/percolate/retry.svg)
* [retry](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go. ![](https://img.shields.io/github/last-commit/thedevsaddam/retry.svg)
* [retry](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done. ![](https://img.shields.io/github/last-commit/shafreeck/retry.svg)
* [retry-go](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang. ![](https://img.shields.io/github/last-commit/rafaeljesus/retry-go.svg)
* [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics. ![](https://img.shields.io/github/last-commit/VividCortex/robustly.svg)
* [scan](https://github.com/blockloop/scan) - Scan golang `sql.Rows` directly to structs, slices, or primitive types. ![](https://img.shields.io/github/last-commit/blockloop/scan.svg)
* [serve](https://github.com/syntaqx/serve) - A static http server anywhere you need. ![](https://img.shields.io/github/last-commit/syntaqx/serve.svg)
* [shutdown](https://github.com/ztrue/shutdown) - App shutdown hooks for `os.Signal` handling. ![](https://img.shields.io/github/last-commit/ztrue/shutdown.svg)
* [silk](https://github.com/chrispassas/silk) - Read silk netflow files. ![](https://img.shields.io/github/last-commit/chrispassas/silk.svg)
* [slice](https://github.com/psampaz/slice) - Type-safe functions for common Go slice operations. ![](https://img.shields.io/github/last-commit/psampaz/slice.svg)
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) - Slice conversion between primitive types. ![](https://img.shields.io/github/last-commit/Henry-Sarabia/sliceconv.svg)
* [slicer](https://github.com/leaanthony/slicer) - Makes working with slices easier. ![](https://img.shields.io/github/last-commit/leaanthony/slicer.svg)
* [sorty](https://github.com/jfcg/sorty) - Fast Concurrent / Parallel Sorting. ![](https://img.shields.io/github/last-commit/jfcg/sorty.svg)
* [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. ![](https://img.shields.io/github/last-commit/briandowns/spinner.svg)
* [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. ![](https://img.shields.io/github/last-commit/jmoiron/sqlx.svg)
* [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. ![](https://img.shields.io/github/last-commit/asdine/storm.svg)
* [structs](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs. ![](https://img.shields.io/github/last-commit/PumpkinSeed/structs.svg)
* [Task](https://github.com/go-task/task) - simple "Make" alternative. ![](https://img.shields.io/github/last-commit/go-task/task.svg)
* [taskctl](https://github.com/taskctl/taskctl) - Concurrent task runner. ![](https://img.shields.io/github/last-commit/taskctl/taskctl.svg)
* [tome](https://github.com/cyruzin/tome) - Tome was designed to paginate simple RESTful APIs. ![](https://img.shields.io/github/last-commit/cyruzin/tome.svg)
* [toolbox](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. ![](https://img.shields.io/github/last-commit/viant/toolbox.svg)
* [ugo](https://github.com/alxrm/ugo) - ugo is slice toolbox with concise syntax for Go. ![](https://img.shields.io/github/last-commit/alxrm/ugo.svg)
* [UNIS](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go. ![](https://img.shields.io/github/last-commit/esemplastic/unis.svg)
* [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. ![](https://img.shields.io/github/last-commit/knq/usql.svg)
* [util](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...). ![](https://img.shields.io/github/last-commit/shomali11/util.svg)
* [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. ![](https://img.shields.io/github/last-commit/asciimoo/wuzz.svg)
* [xferspdy](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang. ![](https://img.shields.io/github/last-commit/monmohan/xferspdy.svg)

## UUID

*Libraries for working with UUIDs.*

* [goid](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs. ![](https://img.shields.io/github/last-commit/jakehl/goid.svg)
* [nanoid](https://github.com/aidarkhanov/nanoid) - A tiny and efficient Go unique string ID generator. ![](https://img.shields.io/github/last-commit/aidarkhanov/nanoid.svg)
* [sno](https://github.com/muyo/sno) - Compact, sortable and fast unique IDs with embedded metadata. ![](https://img.shields.io/github/last-commit/muyo/sno.svg)
* [ulid](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). ![](https://img.shields.io/github/last-commit/oklog/ulid.svg)
* [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.
* [uuid](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. ![](https://img.shields.io/github/last-commit/agext/uuid.svg)
* [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. ![](https://img.shields.io/github/last-commit/gofrs/uuid.svg)
* [uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. ![](https://img.shields.io/github/last-commit/google/uuid.svg)
* [wuid](https://github.com/edwingeng/wuid) - An extremely fast unique number generator, 10-135 times faster than UUID. ![](https://img.shields.io/github/last-commit/edwingeng/wuid.svg)

## Validation

*Libraries for validation.*

* [checkdigit](https://github.com/osamingo/checkdigit) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). ![](https://img.shields.io/github/last-commit/osamingo/checkdigit.svg)
* [gody](https://github.com/guiferpa/gody) - :balloon: A lightweight struct validator for Go. ![](https://img.shields.io/github/last-commit/guiferpa/gody.svg)
* [govalid](https://github.com/twharmon/govalid) - Fast, tag-based validation for structs. ![](https://img.shields.io/github/last-commit/twharmon/govalid.svg)
* [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. ![](https://img.shields.io/github/last-commit/asaskevich/govalidator.svg)
* [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. ![](https://img.shields.io/github/last-commit/thedevsaddam/govalidator.svg)
* [jio](https://github.com/faceair/jio) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). ![](https://img.shields.io/github/last-commit/faceair/jio.svg)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. ![](https://img.shields.io/github/last-commit/go-ozzo/ozzo-validation.svg)
* [terraform-validator](https://github.com/thazelart/terraform-validator) - A norms and conventions validator for Terraform. ![](https://img.shields.io/github/last-commit/thazelart/terraform-validator.svg)
* [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. ![](https://img.shields.io/github/last-commit/gookit/validate.svg)
* [validate](https://github.com/gobuffalo/validate) - This package provides a framework for writing validations for Go applications. ![](https://img.shields.io/github/last-commit/gobuffalo/validate.svg)
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. ![](https://img.shields.io/github/last-commit/go-playground/validator.svg)

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. ![](https://img.shields.io/github/last-commit/rjeczalik/gh.svg)
* [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. ![](https://img.shields.io/github/last-commit/libgit2/git2go.svg)
* [go-git](https://github.com/src-d/go-git) - highly extensible Git implementation in pure Go. ![](https://img.shields.io/github/last-commit/src-d/go-git.svg)
* [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go. ![](https://img.shields.io/github/last-commit/sourcegraph/go-vcs.svg)
* [hercules](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history. ![](https://img.shields.io/github/last-commit/src-d/hercules.svg)
* [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. ![](https://img.shields.io/github/last-commit/beyang/hgo.svg)

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries. ![](https://img.shields.io/github/last-commit/3d0c/gmf.svg)
* [go-astisub](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). ![](https://img.shields.io/github/last-commit/asticode/go-astisub.svg)
* [go-astits](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO. ![](https://img.shields.io/github/last-commit/asticode/go-astits.svg)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) - Parser and generator library for Apple m3u8 playlists. ![](https://img.shields.io/github/last-commit/quangngotan95/go-m3u8.svg)
* [go-mpd](https://github.com/unki2aut/go-mpd) - Parser and generator library for MPEG-DASH manifest files. ![](https://img.shields.io/github/last-commit/unki2aut/go-mpd.svg)
* [goav](https://github.com/giorgisio/goav) - Comphrensive Go bindings for FFmpeg. ![](https://img.shields.io/github/last-commit/giorgisio/goav.svg)
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer. ![](https://img.shields.io/github/last-commit/ziutek/gst.svg)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) - Subtitle format support for go. Supports .srt, .ttml, and .ass. ![](https://img.shields.io/github/last-commit/wargarblgarbl/libgosubs.svg)
* [libvlc-go](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). ![](https://img.shields.io/github/last-commit/adrg/libvlc-go.svg)
* [m3u8](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS. ![](https://img.shields.io/github/last-commit/grafov/m3u8.svg)
* [v4l](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go. ![](https://img.shields.io/github/last-commit/korandiz/v4l.svg)

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
* [Aero](https://github.com/aerogo/aero) - High-performance web framework for Go, reaches top scores in Lighthouse. ![](https://img.shields.io/github/last-commit/aerogo/aero.svg)
* [Air](https://github.com/aofei/air) - An ideally refined web framework for Go. ![](https://img.shields.io/github/last-commit/aofei/air.svg)
* [appy](https://github.com/appist/appy) - An opinionated productive web framework that helps scaling business easier. ![](https://img.shields.io/github/last-commit/appist/appy.svg)
* [Banjo](https://github.com/nsheremet/banjo) - Very simple and fast web framework for Go. ![](https://img.shields.io/github/last-commit/nsheremet/banjo.svg)
* [Beego](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language. ![](https://img.shields.io/github/last-commit/astaxie/beego.svg)
* [Buffalo](http://gobuffalo.io) - Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. ![](https://img.shields.io/github/last-commit/labstack/echo.svg)
* [Fiber](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp. ![](https://img.shields.io/github/last-commit/gofiber/fiber.svg)
* [Fireball](https://github.com/zpatrick/fireball) - More "natural" feeling web framework. ![](https://img.shields.io/github/last-commit/zpatrick/fireball.svg)
* [Flamingo](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. ![](https://img.shields.io/github/last-commit/i-love-flamingo/flamingo.svg)
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. ![](https://img.shields.io/github/last-commit/i-love-flamingo/flamingo-commerce.svg)
* [Gearbox](https://github.com/abahmed/gearbox) - A web framework written in Go with a focus on high performance and memory optimization. ![](https://img.shields.io/github/last-commit/abahmed/gearbox.svg)
* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. ![](https://img.shields.io/github/last-commit/gin-gonic/gin.svg)
* [Ginrpc](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools. ![](https://img.shields.io/github/last-commit/xxjwxc/ginrpc.svg)
* [Gizmo](https://github.com/NYTimes/gizmo) - Microservice toolkit used by the New York Times. ![](https://img.shields.io/github/last-commit/NYTimes/gizmo.svg)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) - Quick and easy way to setup a RESTful JSON API. ![](https://img.shields.io/github/last-commit/ant0ine/go-json-rest.svg)
* [go-rest](https://github.com/ungerik/go-rest) - Small and evil REST framework for Go. ![](https://img.shields.io/github/last-commit/ungerik/go-rest.svg)
* [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go. ![](https://img.shields.io/github/last-commit/goadesign/goa.svg)
* [goa](https://github.com/goa-go/goa) - goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. ![](https://img.shields.io/github/last-commit/goa-go/goa.svg)
* [Golax](https://github.com/fulldump/golax) - A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. ![](https://img.shields.io/github/last-commit/fulldump/golax.svg)
* [Golf](https://github.com/dinever/golf) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. ![](https://img.shields.io/github/last-commit/dinever/golf.svg)
* [Gondola](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster. ![](https://img.shields.io/github/last-commit/rainycape/gondola.svg)
* [gongular](https://github.com/mustafaakin/gongular) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection. ![](https://img.shields.io/github/last-commit/mustafaakin/gongular.svg)
* [goweb](https://github.com/twharmon/goweb) - Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. ![](https://img.shields.io/github/last-commit/twharmon/goweb.svg)
* [Goyave](https://github.com/System-Glitch/goyave) - Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities. ![](https://img.shields.io/github/last-commit/System-Glitch/goyave.svg)
* [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support. ![](https://img.shields.io/github/last-commit/hidevopsio/hiboot.svg)
* [Macaron](https://github.com/go-macaron/macaron) - Macaron is a high productive and modular design web framework in Go. ![](https://img.shields.io/github/last-commit/go-macaron/macaron.svg)
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. ![](https://img.shields.io/github/last-commit/paulbellamy/mango.svg)
* [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang. ![](https://img.shields.io/github/last-commit/claygod/microservice.svg)
* [neo](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API. ![](https://img.shields.io/github/last-commit/ivpusic/neo.svg)
* [patron](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity. ![](https://img.shields.io/github/last-commit/beatlabs/patron.svg)
* [Resoursea](https://github.com/resoursea/api) - REST framework for quickly writing resource based services. ![](https://img.shields.io/github/last-commit/resoursea/api.svg)
* [REST Layer](http://rest-layer.io) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. ![](https://img.shields.io/github/last-commit/revel/revel.svg)
* [rex](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. ![](https://img.shields.io/github/last-commit/goanywhere/rex.svg)
* [rux](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications. ![](https://img.shields.io/github/last-commit/gookit/rux.svg)
* [tango](https://github.com/lunny/tango) - Micro & pluggable web framework for Go. ![](https://img.shields.io/github/last-commit/lunny/tango.svg)
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - Go framework for building JSON web services inspired by Dropwizard. ![](https://img.shields.io/github/last-commit/rcrowley/go-tigertonic.svg)
* [uAdmin](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django. ![](https://img.shields.io/github/last-commit/uadmin/uadmin.svg)
* [utron](https://github.com/gernest/utron) - Lightweight MVC framework for Go(Golang). ![](https://img.shields.io/github/last-commit/gernest/utron.svg)
* [vox](https://github.com/aisk/vox) - A golang web framework for humans, inspired by Koa heavily. ![](https://img.shields.io/github/last-commit/aisk/vox.svg)
* [WebGo](https://github.com/bnkamalesh/webgo) - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). ![](https://img.shields.io/github/last-commit/bnkamalesh/webgo.svg)
* [YARF](https://github.com/yarf-framework/yarf) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way. ![](https://img.shields.io/github/last-commit/yarf-framework/yarf.svg)

### Middlewares

#### Actual middlewares

* [client-timing](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header. ![](https://img.shields.io/github/last-commit/posener/client-timing.svg)
* [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. ![](https://img.shields.io/github/last-commit/rs/cors.svg)
* [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST. ![](https://img.shields.io/github/last-commit/rs/formjson.svg)
* [go-fault](https://github.com/github/go-fault) - Fault injection middleware for Go. ![](https://img.shields.io/github/last-commit/github/go-fault.svg)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header. ![](https://img.shields.io/github/last-commit/mitchellh/go-server-timing.svg)
* [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. ![](https://img.shields.io/github/last-commit/ulule/limiter.svg)
* [ln-paywall](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). ![](https://img.shields.io/github/last-commit/philippgille/ln-paywall.svg)
* [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. ![](https://img.shields.io/github/last-commit/didip/tollbooth.svg)
* [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends. ![](https://img.shields.io/github/last-commit/sebest/xff.svg)

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. ![](https://img.shields.io/github/last-commit/justinas/alice.svg)
* [catena](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain"). ![](https://img.shields.io/github/last-commit/codemodus/catena.svg)
* [chain](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware"). ![](https://img.shields.io/github/last-commit/codemodus/chain.svg)
* [go-wrap](https://github.com/go-on/wrap) - Small middlewares package for net/http. ![](https://img.shields.io/github/last-commit/go-on/wrap.svg)
* [gores](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. ![](https://img.shields.io/github/last-commit/alioygur/gores.svg)
* [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang. ![](https://img.shields.io/github/last-commit/carbocation/interpose.svg)
* [mediary](https://github.com/HereMobilityDevelopers/mediary) - add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. ![](https://img.shields.io/github/last-commit/HereMobilityDevelopers/mediary.svg)
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http. ![](https://img.shields.io/github/last-commit/stephens2424/muxchain.svg)
* [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. ![](https://img.shields.io/github/last-commit/urfave/negroni.svg)
* [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. ![](https://img.shields.io/github/last-commit/unrolled/render.svg)
* [renderer](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. ![](https://img.shields.io/github/last-commit/thedevsaddam/renderer.svg)
* [rye](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. ![](https://img.shields.io/github/last-commit/InVisionApp/rye.svg)
* [stats](https://github.com/thoas/stats) - Go middleware that stores various information about your web application. ![](https://img.shields.io/github/last-commit/thoas/stats.svg)

### Routers

* [alien](https://github.com/gernest/alien) - Lightweight and fast http router from outer space. ![](https://img.shields.io/github/last-commit/gernest/alien.svg)
* [bellt](https://github.com/GuilhermeCaruso/bellt) - A simple Go HTTP router. ![](https://img.shields.io/github/last-commit/GuilhermeCaruso/bellt.svg)
* [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. ![](https://img.shields.io/github/last-commit/go-zoo/bone.svg)
* [Bxog](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. ![](https://img.shields.io/github/last-commit/claygod/Bxog.svg)
* [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. ![](https://img.shields.io/github/last-commit/go-chi/chi.svg)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`. ![](https://img.shields.io/github/last-commit/buaazp/fasthttprouter.svg)
* [FastRouter](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go. ![](https://img.shields.io/github/last-commit/razonyang/fastrouter.svg)
* [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. ![](https://img.shields.io/github/last-commit/gocraft/web.svg)
* [Goji](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. ![](https://img.shields.io/github/last-commit/goji/goji.svg)
* [goroute](https://github.com/goroute/route) - Simple yet powerful HTTP request multiplexer. ![](https://img.shields.io/github/last-commit/goroute/route.svg)
* [GoRouter](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. ![](https://img.shields.io/github/last-commit/vardius/gorouter.svg)
* [gowww/router](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface. ![](https://img.shields.io/github/last-commit/gowww/router.svg)
* [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. ![](https://img.shields.io/github/last-commit/julienschmidt/httprouter.svg)
* [httptreemux](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. ![](https://img.shields.io/github/last-commit/dimfeld/httptreemux.svg)
* [lars](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. ![](https://img.shields.io/github/last-commit/go-playground/lars.svg)
* [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. ![](https://img.shields.io/github/last-commit/gorilla/mux.svg)
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. ![](https://img.shields.io/github/last-commit/go-ozzo/ozzo-routing.svg)
* [pure](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation. ![](https://img.shields.io/github/last-commit/go-playground/pure.svg)
* [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers. ![](https://img.shields.io/github/last-commit/VividCortex/siesta.svg)
* [vestigo](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications. ![](https://img.shields.io/github/last-commit/husobee/vestigo.svg)
* [violetear](https://github.com/nbari/violetear) - Go HTTP router. ![](https://img.shields.io/github/last-commit/nbari/violetear.svg)
* [xmux](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support. ![](https://img.shields.io/github/last-commit/rs/xmux.svg)
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go. ![](https://img.shields.io/github/last-commit/xujiajun/gorouter.svg)

## WebAssembly

* [dom](https://github.com/dennwc/dom) - DOM library. ![](https://img.shields.io/github/last-commit/dennwc/dom.svg)
* [go-canvas](https://github.com/markfarnan/go-canvas) - Library to use HTML5 Canvas, with all drawing within go code. ![](https://img.shields.io/github/last-commit/markfarnan/go-canvas.svg)
* [tinygo](https://github.com/tinygo-org/tinygo) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. ![](https://img.shields.io/github/last-commit/tinygo-org/tinygo.svg)
* [vert](https://github.com/norunners/vert) - Interop between Go and JS values. ![](https://img.shields.io/github/last-commit/norunners/vert.svg)
* [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) - Run Go WASM tests in your browser. ![](https://img.shields.io/github/last-commit/agnivade/wasmbrowsertest.svg)
* [webapi](https://github.com/gowebapi/webapi) - Bindings for DOM and HTML generated from WebIDL. ![](https://img.shields.io/github/last-commit/gowebapi/webapi.svg)

## Windows

* [d3d9](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9. ![](https://img.shields.io/github/last-commit/gonutz/d3d9.svg)
* [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. ![](https://img.shields.io/github/last-commit/go-ole/go-ole.svg)
* [gosddl](https://github.com/MonaxGT/gosddl) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. ![](https://img.shields.io/github/last-commit/MonaxGT/gosddl.svg)

## XML

*Libraries and tools for manipulating XML.*

* [XML-Comp](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags. ![](https://img.shields.io/github/last-commit/xml-comp/xml-comp.svg)
* [xml2map](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang. ![](https://img.shields.io/github/last-commit/sbabiv/xml2map.svg)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module. ![](https://img.shields.io/github/last-commit/shabbyrobe/xmlwriter.svg)
* [xpath](https://github.com/antchfx/xpath) - XPath package for Go. ![](https://img.shields.io/github/last-commit/antchfx/xpath.svg)
* [xquery](https://github.com/antchfx/xquery) - XQuery lets you extract data from HTML/XML documents using XPath expression. ![](https://img.shields.io/github/last-commit/antchfx/xquery.svg)
* [zek](https://github.com/miku/zek) - Generate a Go struct from XML. ![](https://img.shields.io/github/last-commit/miku/zek.svg)

# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes. ![](https://img.shields.io/github/last-commit/bradleyfalzon/apicompat.svg)
* [dupl](https://github.com/mibk/dupl) - Tool for code clone detection. ![](https://img.shields.io/github/last-commit/mibk/dupl.svg)
* [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. ![](https://img.shields.io/github/last-commit/kisielk/errcheck.svg)
* [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time. ![](https://img.shields.io/github/last-commit/davecheney/gcvis.svg)
* [go-checkstyle](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. ![](https://img.shields.io/github/last-commit/qiniu/checkstyle.svg)
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. ![](https://img.shields.io/github/last-commit/roblaszczak/go-cleanarch.svg)
* [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. ![](https://img.shields.io/github/last-commit/go-critic/go-critic.svg)
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) - An easy way to find outdated dependencies of your Go projects. ![](https://img.shields.io/github/last-commit/psampaz/go-mod-outdated.svg)
* [go-outdated](https://github.com/firstrow/go-outdated) - Console application that displays outdated packages. ![](https://img.shields.io/github/last-commit/firstrow/go-outdated.svg)
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer. ![](https://img.shields.io/github/last-commit/yuroyoro/goast-viewer.svg)
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/) - GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [golines](https://github.com/segmentio/golines) - Formatter that automatically shortens long lines in Go code. ![](https://img.shields.io/github/last-commit/segmentio/golines.svg)
* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code. ![](https://img.shields.io/github/last-commit/golang/lint.svg)
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. ![](https://img.shields.io/github/last-commit/jfeliu007/goplantuml.svg)
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - gosimple is a linter for Go source code that specialises on simplifying code.
* [gostatus](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages. ![](https://img.shields.io/github/last-commit/shurcooL/gostatus.svg)
* [lint](https://github.com/surullabs/lint) - Run linters as part of go test. ![](https://img.shields.io/github/last-commit/surullabs/lint.svg)
* [php-parser](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go. ![](https://img.shields.io/github/last-commit/z7zmey/php-parser.svg)
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) - tarp finds functions and methods without direct unit tests in Go source code. ![](https://img.shields.io/github/last-commit/verygoodsoftwarenotvirus/tarp.svg)
* [tickgit](https://github.com/augmentable-dev/tickgit) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. ![](https://img.shields.io/github/last-commit/augmentable-dev/tickgit.svg)
* [unconvert](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source. ![](https://img.shields.io/github/last-commit/mdempsky/unconvert.svg)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - unused checks Go code for unused constants, variables, functions and types.
* [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags. ![](https://img.shields.io/github/last-commit/mccoyst/validate.svg)

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. ![](https://img.shields.io/github/last-commit/theia-ide/go-language-server.svg)
* [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs. ![](https://img.shields.io/github/last-commit/dominikh/go-mode.el.svg)
* [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. ![](https://img.shields.io/github/last-commit/joefitzgerald/go-plus.svg)
* [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. ![](https://img.shields.io/github/last-commit/nsf/gocode.svg)
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. ![](https://img.shields.io/github/last-commit/DisposaBoy/GoSublime.svg)
* [gounit-vim](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature. ![](https://img.shields.io/github/last-commit/hexdigest/gounit-vim.svg)
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) - Go language support for the Theia IDE. ![](https://img.shields.io/github/last-commit/theia-ide/theia-go-extension.svg)
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save. ![](https://img.shields.io/github/last-commit/rjohnsondev/vim-compiler-go.svg)
* [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. ![](https://img.shields.io/github/last-commit/fatih/vim-go.svg)
* [vscode-go](https://github.com/Microsoft/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. ![](https://img.shields.io/github/last-commit/Microsoft/vscode-go.svg)
* [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes. ![](https://img.shields.io/github/last-commit/eaburns/Watch.svg)

## Go Generate Tools

* [generic](https://github.com/usk81/generic) - flexible data type for Go. ![](https://img.shields.io/github/last-commit/usk81/generic.svg)
* [genny](https://github.com/cheekybits/genny) - Elegant generics for Go. ![](https://img.shields.io/github/last-commit/cheekybits/genny.svg)
* [gocontracts](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation. ![](https://img.shields.io/github/last-commit/Parquery/gocontracts.svg)
* [gonerics](http://github.com/bouk/gonerics) - Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. ![](https://img.shields.io/github/last-commit/cweill/gotests.svg)
* [gounit](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates. ![](https://img.shields.io/github/last-commit/hexdigest/gounit.svg)
* [hasgo](https://github.com/DylanMeeus/hasgo) - Generate Haskell inspired functions for your slices. ![](https://img.shields.io/github/last-commit/DylanMeeus/hasgo.svg)
* [re2dfa](https://github.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code. ![](https://img.shields.io/github/last-commit/opennota/re2dfa.svg)
* [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.
* [xgen](https://github.com/xuri/xgen) - XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. ![](https://img.shields.io/github/last-commit/xuri/xgen.svg)

## Go Tools

* [colorgo](https://github.com/songgao/colorgo) - Wrapper around `go` command for colorized `go build` output. ![](https://img.shields.io/github/last-commit/songgao/colorgo.svg)
* [depth](https://github.com/KyleBanks/depth) - Visualize dependency trees of any package by analyzing imports. ![](https://img.shields.io/github/last-commit/KyleBanks/depth.svg)
* [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) - A [Yeoman](http://yeoman.io) generator to get new Go projects started. ![](https://img.shields.io/github/last-commit/axelspringer/generator-go-lang.svg)
* [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.
* [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. ![](https://img.shields.io/github/last-commit/TrueFurby/go-callvis.svg)
* [go-james](https://github.com/pieterclaerhout/go-james) - Go project skeleton creator, builds and tests your projects without the manual setup. ![](https://img.shields.io/github/last-commit/pieterclaerhout/go-james.svg)
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) - Bash completion for go and wgo. ![](https://img.shields.io/github/last-commit/skelterjohn/go-pkg-complete.svg)
* [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. ![](https://img.shields.io/github/last-commit/go-swagger/go-swagger.svg)
* [godbg](https://github.com/tylerwince/godbg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development. ![](https://img.shields.io/github/last-commit/tylerwince/godbg.svg)
* [gomodrun](https://github.com/dustinblackman/gomodrun/) - Go tool that executes and caches binaries included in go.mod files.
* [gothanks](https://github.com/psampaz/gothanks) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. ![](https://img.shields.io/github/last-commit/psampaz/gothanks.svg)
* [igo](https://github.com/rocketlaunchr/igo) - An igo to go transpiler (new language features for Go language!) ![](https://img.shields.io/github/last-commit/rocketlaunchr/igo.svg)
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. ![](https://img.shields.io/github/last-commit/OctoLinker/browser-extension.svg)
* [richgo](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations. ![](https://img.shields.io/github/last-commit/kyoh86/richgo.svg)
* [rts](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses. ![](https://img.shields.io/github/last-commit/galeone/rts.svg)
* [typex](https://github.com/dtgorski/typex) - Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. ![](https://img.shields.io/github/last-commit/dtgorski/typex.svg)

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool. ![](https://img.shields.io/github/last-commit/smira/aptly.svg)
* [aurora](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console. ![](https://img.shields.io/github/last-commit/xuri/aurora.svg)
* [awsenv](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile. ![](https://img.shields.io/github/last-commit/soniah/awsenv.svg)
* [Blast](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs. ![](https://img.shields.io/github/last-commit/dave/blast.svg)
* [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. ![](https://img.shields.io/github/last-commit/codesenberg/bombardier.svg)
* [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. ![](https://img.shields.io/github/last-commit/bosun-monitor/bosun.svg)
* [cassowary](https://github.com/rogerwelin/cassowary) - Modern cross-platform HTTP load-testing tool written in Go. ![](https://img.shields.io/github/last-commit/rogerwelin/cassowary.svg)
* [DepCharge](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects. ![](https://img.shields.io/github/last-commit/centerorbit/depcharge.svg)
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) - A go library and an executable that produces valid Dockerfiles using various input channels. ![](https://img.shields.io/github/last-commit/ozankasikci/dockerfile-generator.svg)
* [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). ![](https://img.shields.io/github/last-commit/liudng/dogo.svg)
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI. ![](https://img.shields.io/github/last-commit/appleboy/drone-jenkins.svg)
* [drone-scp](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI. ![](https://img.shields.io/github/last-commit/appleboy/drone-scp.svg)
* [Dropship](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn. ![](https://img.shields.io/github/last-commit/chrismckenzie/dropship.svg)
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. ![](https://img.shields.io/github/last-commit/appleboy/easyssh-proxy.svg)
* [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts. ![](https://img.shields.io/github/last-commit/mkchoi212/fac.svg)
* [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. ![](https://img.shields.io/github/last-commit/gaia-pipeline/gaia.svg)
* [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. ![](https://img.shields.io/github/last-commit/go-gitea/gitea.svg)
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [go-furnace](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. ![](https://img.shields.io/github/last-commit/go-furnace/go-furnace.svg)
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. ![](https://img.shields.io/github/last-commit/sanbornm/go-selfupdate.svg)
* [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go. ![](https://img.shields.io/github/last-commit/cryptojuice/gobrew.svg)
* [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application. ![](https://img.shields.io/github/last-commit/sirnewton01/godbg.svg)
* [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. ![](https://img.shields.io/github/last-commit/inconshreveable/gonative.svg)
* [govvv](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries. ![](https://img.shields.io/github/last-commit/ahmetalpbalkan/govvv.svg)
* [gox](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool. ![](https://img.shields.io/github/last-commit/mitchellh/gox.svg)
* [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging. ![](https://img.shields.io/github/last-commit/laher/goxc.svg)
* [grapes](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease. ![](https://img.shields.io/github/last-commit/yaronsumel/grapes.svg)
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. ![](https://img.shields.io/github/last-commit/moovweb/gvm.svg)
* [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. ![](https://img.shields.io/github/last-commit/rakyll/hey.svg)
* [jcli](https://github.com/jenkins-zh/jenkins-cli) - Jenkins CLI allows you manage your Jenkins as an easy way. ![](https://img.shields.io/github/last-commit/jenkins-zh/jenkins-cli.svg)
* [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. ![](https://img.shields.io/github/last-commit/ajvb/kala.svg)
* [kcli](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages. ![](https://img.shields.io/github/last-commit/cswank/kcli.svg)
* [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. ![](https://img.shields.io/github/last-commit/kubernetes/kubernetes.svg)
* [lstags](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries. ![](https://img.shields.io/github/last-commit/ivanilves/lstags.svg)
* [lwc](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command. ![](https://img.shields.io/github/last-commit/timdp/lwc.svg)
* [manssh](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily. ![](https://img.shields.io/github/last-commit/xwjdsh/manssh.svg)
* [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. ![](https://img.shields.io/github/last-commit/moby/moby.svg)
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data. ![](https://img.shields.io/github/last-commit/emicklei/mora.svg)
* [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. ![](https://img.shields.io/github/last-commit/ostrost/ostent.svg)
* [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. ![](https://img.shields.io/github/last-commit/mitchellh/packer.svg)
* [Pewpew](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester. ![](https://img.shields.io/github/last-commit/bengadbois/pewpew.svg)
* [Pomerium](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy. ![](https://img.shields.io/github/last-commit/pomerium/pomerium.svg)
* [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies. ![](https://img.shields.io/github/last-commit/alouche/rodent.svg)
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) - S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). ![](https://img.shields.io/github/last-commit/oxyno-zeta/s3-proxy.svg)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. ![](https://img.shields.io/github/last-commit/rlmcpherson/s3gof3r.svg)
* [s5cmd](https://github.com/peak/s5cmd) - Blazing fast S3 and local filesystem execution tool. ![](https://img.shields.io/github/last-commit/peak/s5cmd.svg)
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker). ![](https://img.shields.io/github/last-commit/scaleway/scaleway-cli.svg)
* [script](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. ![](https://img.shields.io/github/last-commit/bitfield/script.svg)
* [sg](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. ![](https://img.shields.io/github/last-commit/ChristopherRabotin/sg.svg)
* [skm](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! ![](https://img.shields.io/github/last-commit/TimothyYe/skm.svg)
* [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. ![](https://img.shields.io/github/last-commit/sanathp/statusok.svg)
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. ![](https://img.shields.io/github/last-commit/dikhan/terraform-provider-openapi.svg)
* [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. ![](https://img.shields.io/github/last-commit/containous/traefik.svg)
* [uTask](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml. ![](https://img.shields.io/github/last-commit/ovh/utask.svg)
* [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! ![](https://img.shields.io/github/last-commit/tsenart/vegeta.svg)
* [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. ![](https://img.shields.io/github/last-commit/adnanh/webhook.svg)
* [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
* [winrm-cli](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines. ![](https://img.shields.io/github/last-commit/masterzen/winrm-cli.svg)

### Other Software

* [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. ![](https://img.shields.io/github/last-commit/crufter/borg.svg)
* [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine. ![](https://img.shields.io/github/last-commit/tejo/boxed.svg)
* [Cherry](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go. ![](https://img.shields.io/github/last-commit/rafael-santiago/cherry.svg)
* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. ![](https://img.shields.io/github/last-commit/gocircuit/circuit.svg)
* [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. ![](https://img.shields.io/github/last-commit/tylertreat/Comcast.svg)
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. ![](https://img.shields.io/github/last-commit/kelseyhightower/confd.svg)
* [croc](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another. ![](https://img.shields.io/github/last-commit/schollz/croc.svg)
* [Docker](http://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. ![](https://img.shields.io/github/last-commit/documize/community.svg)
* [dp](https://github.com/scryinfo/dp) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. ![](https://img.shields.io/github/last-commit/scryinfo/dp.svg)
* [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. ![](https://img.shields.io/github/last-commit/odeke-em/drive.svg)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. ![](https://img.shields.io/github/last-commit/gilbertchen/duplicacy.svg)
* [gfile](https://github.com/Antonito/gfile) - Securely transfer files between two computers, without any third party, over WebRTC. ![](https://img.shields.io/github/last-commit/Antonito/gfile.svg)
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH. ![](https://img.shields.io/github/last-commit/shurcooL/Go-Package-Store.svg)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client. ![](https://img.shields.io/github/last-commit/Sioro-Neoku/go-peerflix.svg)
* [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. ![](https://img.shields.io/github/last-commit/Humpheh/goboy.svg)
* [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go. ![](https://img.shields.io/github/last-commit/goccmack/gocc.svg)
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list. ![](https://img.shields.io/github/last-commit/diankong/GoDocTooltip.svg)
* [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. ![](https://img.shields.io/github/last-commit/buger/gor.svg)
* [hugo](http://gohugo.io/) - Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go. ![](https://img.shields.io/github/last-commit/thestrukture/ide.svg)
* [ipe](https://github.com/dimiro1/ipe) - Open source Pusher server implementation compatible with Pusher client libraries written in GO. ![](https://img.shields.io/github/last-commit/dimiro1/ipe.svg)
* [joincap](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together. ![](https://img.shields.io/github/last-commit/assafmo/joincap.svg)
* [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Leaps](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms. ![](https://img.shields.io/github/last-commit/jeffail/leaps.svg)
* [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. ![](https://img.shields.io/github/last-commit/yunabe/lgo.svg)
* [limetext](http://limetext.org/) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. ![](https://img.shields.io/github/last-commit/visualfc/liteide.svg)
* [mockingjay](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. ![](https://img.shields.io/github/last-commit/quii/mockingjay-server.svg)
* [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. ![](https://img.shields.io/github/last-commit/mehrdadrad/mylg.svg)
* [naclpipe](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go. ![](https://img.shields.io/github/last-commit/unix4fun/naclpipe.svg)
* [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. ![](https://img.shields.io/github/last-commit/fogleman/nes.svg)
* [orange-cat](https://github.com/noraesae/orange-cat) - Markdown previewer written in Go. ![](https://img.shields.io/github/last-commit/noraesae/orange-cat.svg)
* [Orbit](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates. ![](https://img.shields.io/github/last-commit/gulien/orbit.svg)
* [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. ![](https://img.shields.io/github/last-commit/pointlander/peg.svg)
* [restic](https://github.com/restic/restic) - De-duplicating backup program. ![](https://img.shields.io/github/last-commit/restic/restic.svg)
* [scc](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. ![](https://img.shields.io/github/last-commit/boyter/scc.svg)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek. ![](https://img.shields.io/github/last-commit/chrislusf/seaweedfs.svg)
* [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). ![](https://img.shields.io/github/last-commit/msoap/shell2http.svg)
* [snap](https://github.com/intelsdi-x/snap) - Powerful telemetry framework. ![](https://img.shields.io/github/last-commit/intelsdi-x/snap.svg)
* [Snitch](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru. ![](https://img.shields.io/github/last-commit/lucasgomide/snitch.svg)
* [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. ![](https://img.shields.io/github/last-commit/pressly/sup.svg)
* [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
* [term-quiz](https://github.com/crazcalm/term-quiz) - Quizzes for your terminal. ![](https://img.shields.io/github/last-commit/crazcalm/term-quiz.svg)
* [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. ![](https://img.shields.io/github/last-commit/shopify/toxiproxy.svg)
* [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. ![](https://img.shields.io/github/last-commit/VerizonDigital/vflow.svg)
* [wellington](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass). ![](https://img.shields.io/github/last-commit/wellington/wellington.svg)

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions. ![](https://img.shields.io/github/last-commit/davecheney/autobench.svg)
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. ![](https://img.shields.io/github/last-commit/mrLSD/go-benchmark-app.svg)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. ![](https://img.shields.io/github/last-commit/tylertreat/go-benchmarks.svg)
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. ![](https://img.shields.io/github/last-commit/julienschmidt/go-http-routing-benchmark.svg)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. ![](https://img.shields.io/github/last-commit/smallnest/go-web-framework-benchmark.svg)
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. ![](https://img.shields.io/github/last-commit/alecthomas/go_serialization_benchmarks.svg)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. ![](https://img.shields.io/github/last-commit/PuerkitoBio/gocostmodel.svg)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities. ![](https://img.shields.io/github/last-commit/tyler-smith/golang-sql-benchmark.svg)
* [gospeed](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs. ![](https://img.shields.io/github/last-commit/feyeleanor/GoSpeed.svg)
* [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark. ![](https://img.shields.io/github/last-commit/jimrobinson/kvbench.svg)
* [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. ![](https://img.shields.io/github/last-commit/atemerev/skynet.svg)
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language. ![](https://img.shields.io/github/last-commit/fawick/speedtest-resize.svg)

## Conferences

* [Capital Go](http://www.capitalgolang.com) - Washington, D.C., USA.
* [dotGo](http://www.dotgo.eu) - Paris, France.
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan.
* [GoDays](https://www.godays.io/) - Berlin, Germany.
* [GoLab](http://golab.io/) - Florence, Italy.
* [GolangUK](http://golanguk.com/) - London, UK.
* [GopherChina](http://gopherchina.org) - Shanghai, China.
* [GopherCon](http://www.gophercon.com/) - Denver, USA.
* [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
* [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, BR.
* [GopherCon Europe](https://gophercon.is/) - Berlin, Germany.
* [GopherCon India](https://www.gophercon.in/) - Pune, India.
* [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel.
* [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia.
* [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore.
* [GopherCon Vietnam](https://gophercon.vn/) - Ho Chi Minh City, Vietnam.
* [GothamGo](http://gothamgo.com/) - New York City, USA.
* [GoWayFest](https://goway.io/) - Minsk, Belarus.

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
* [Go Bootcamp](http://golangbootcamp.com)
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) - in Persian. ![](https://img.shields.io/github/last-commit/thedevsir/gosuccinctly.svg)
* [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. ![](https://img.shields.io/github/last-commit/dariubs/GoBooks.svg)
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) ![](https://img.shields.io/github/last-commit/polaris1119/The-Golang-Standard-Library-by-Example.svg)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. ![](https://img.shields.io/github/last-commit/MariaLetta/free-gophers-pack.svg)
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg]. ![](https://img.shields.io/github/last-commit/keygx/Go-gopher-Vector.svg)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) - adorable gopher logos. ![](https://img.shields.io/github/last-commit/GolangUA/gopher-logos.svg)
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers) ![](https://img.shields.io/github/last-commit/tenntenn/gopher-stickers.svg)
* [gopher-vector](https://github.com/golang-samples/gopher-vector) ![](https://img.shields.io/github/last-commit/golang-samples/gopher-vector.svg)
* [gophericons](https://github.com/shalakhin/gophericons) ![](https://img.shields.io/github/last-commit/shalakhin/gophericons.svg)
* [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself. ![](https://img.shields.io/github/last-commit/matryer/gopherize.me.svg)
* [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara. ![](https://img.shields.io/github/last-commit/ashleymcnamara/gophers.svg)
* [gophers](https://github.com/egonelbre/gophers) - Free gophers. ![](https://img.shields.io/github/last-commit/egonelbre/gophers.svg)
* [gophers](https://github.com/rogeralsing/gophers) - random gopher graphics. ![](https://img.shields.io/github/last-commit/rogeralsing/gophers.svg)
* [gophers](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern. ![](https://img.shields.io/github/last-commit/sillecelik/go-gopher.svg)

## Meetups

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
* [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
* [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBandung](https://www.meetup.com/GoBandung/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
* [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
* [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
* [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
* [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
* [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
* [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Казань](https://www.meetup.com/GolangKazan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*Add the group of your city/country here (send **PR**)*

## Social Media
### Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

### Reddit
 * [r/golang](https://www.reddit.com/r/golang/)

## Websites

* [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers. ![](https://img.shields.io/github/last-commit/lukasz-madon/awesome-remote-job.svg)
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. ![](https://img.shields.io/github/last-commit/bayandin/awesome-awesomeness.svg)
* [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org) - The official Go blog.
* [Go Challenge](http://golang-challenge.org/) - Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com) - A report card for your Go package.
* [go.dev](https://go.dev/) - A hub for Go developers.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go. ![](https://img.shields.io/github/last-commit/ninedraft/gocryforhelp.svg)
* [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
* [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusivly for Golang related Roles.
* [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com) - Links and news about Go programming.
* [golang-graphics](https://github.com/mholt/golang-graphics) - Collection of Go images, graphics, and art. ![](https://img.shields.io/github/last-commit/mholt/golang-graphics.svg)
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
* [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang) - News about Go.
* [studygolang](https://studygolang.com) - The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang. ![](https://img.shields.io/github/last-commit/astaxie/build-web-application-with-golang.svg)
* [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) - We’ll write an API with the help of the powerful Gorilla Mux.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) - A little e-book on Ethereum Development with Go. ![](https://img.shields.io/github/last-commit/miguelmota/ethereum-development-with-go-book.svg)
* [Games With Go](http://gameswithgo.org/) - A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card. ![](https://img.shields.io/github/last-commit/a8m/go-lang-cheat-sheet.svg)
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
* [Go Playground for iOS](https://codeplayground.app) - Interactively edit & play Go snippets on your mobile device.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms. ![](https://img.shields.io/github/last-commit/tmrts/go-patterns.svg)
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning. ![](https://img.shields.io/github/last-commit/miguelmota/golang-for-nodejs-developers.svg)
* [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/) - Collection of code snippets and tutorials to help tackle every day issues.
* [GopherSnippets](https://gophersnippets.com/) - Code snippets with tests and testable examples for the Go programming language.
* [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5) - Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
* [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Learn how to use Docker for Go development and how to build production Docker images.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development. ![](https://img.shields.io/github/last-commit/quii/learn-go-with-tests.svg)
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/) - Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers. ![](https://img.shields.io/github/last-commit/mkaz/working-with-go.svg)
* [Your basic Go](http://yourbasic.org/golang) - Huge collection of tutorials and how to's.
