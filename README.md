# Awesome Go

<a href="https://awesome-go.com/"><img align="right" src="https://github.com/avelino/awesome-go/raw/main/tmpl/assets/logo.png" alt="awesome-go" title="awesome-go" /></a>

[![Build Status](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml?query=branch%3Amain)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/awesome)
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)
[![Last Commit](https://img.shields.io/github/last-commit/avelino/awesome-go)](https://github.com/avelino/awesome-go/commits/main)

We use the _[Golang Bridge](https://github.com/gobridge/about-us/blob/master/README.md)_ community Slack for instant communication, follow the [form here to join](https://invite.slack.golangbridge.org/).

<a href="https://www.producthunt.com/posts/awesome-go?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-go" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=291535&theme=light" alt="awesome-go - Curated list awesome Go frameworks, libraries and software | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

**Sponsorships:**

_Special thanks to_

<div align="center">
<table cellpadding="5">
<tbody align="center">
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-workos">
<img src="https://avelino.run/sponsors/workos-logo-white-bg.svg" width="200" alt="WorkOS"><br/>
<b>Your app, enterprise-ready.</b><br/>
<sub>Start selling to enterprise customers with just a few lines of code.</sub><br/>
<sup>Add Single Sign-On (and more) in minutes instead of months.</sup>
</a>
</td>
</tr>
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-digitalocean">
<img src="https://avelino.run/sponsors/do_logo_horizontal_blue-210.png" width="200" alt="Digital Ocean">
</a>
</td>
</tr>
</tbody>
</table>
</div>

**Awesome Go has no monthly fee**_, but we have employees who **work hard** to keep it running. With money raised, we can repay the effort of each person involved! You can see how we calculate our billing and distribution as it is open to the entire community. Want to be a supporter of the project click [here](mailto:avelinorun+oss@gmail.com?subject=awesome-go%3A%20project%20support)._

> A curated list of awesome Go frameworks, libraries, and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

**Contributing:**

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

> _If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!_

## Contents

<details>
<summary>Expand contents</summary>

- [Awesome Go](#awesome-go)
  - [Contents](#contents)
  - [Actor Model](#actor-model)
  - [Artificial Intelligence](#artificial-intelligence)
  - [Audio and Music](#audio-and-music)
  - [Authentication and Authorization](#authentication-and-authorization)
  - [Blockchain](#blockchain)
  - [Bot Building](#bot-building)
  - [Build Automation](#build-automation)
  - [Command Line](#command-line)
    - [Advanced Console UIs](#advanced-console-uis)
    - [Standard CLI](#standard-cli)
  - [Configuration](#configuration)
  - [Continuous Integration](#continuous-integration)
  - [CSS Preprocessors](#css-preprocessors)
  - [Data Integration Frameworks](#data-integration-frameworks)
  - [Data Structures and Algorithms](#data-structures-and-algorithms)
    - [Bit-packing and Compression](#bit-packing-and-compression)
    - [Bit Sets](#bit-sets)
    - [Bloom and Cuckoo Filters](#bloom-and-cuckoo-filters)
    - [Data Structure and Algorithm Collections](#data-structure-and-algorithm-collections)
    - [Iterators](#iterators)
    - [Maps](#maps)
    - [Miscellaneous Data Structures and Algorithms](#miscellaneous-data-structures-and-algorithms)
    - [Nullable Types](#nullable-types)
    - [Queues](#queues)
    - [Sets](#sets)
    - [Text Analysis](#text-analysis)
    - [Trees](#trees)
    - [Pipes](#pipes)
  - [Database](#database)
    - [Caches](#caches)
    - [Databases Implemented in Go](#databases-implemented-in-go)
    - [Database Schema Migration](#database-schema-migration)
    - [Database Tools](#database-tools)
    - [SQL Query Builders](#sql-query-builders)
  - [Database Drivers](#database-drivers)
    - [Interfaces to Multiple Backends](#interfaces-to-multiple-backends)
    - [Relational Database Drivers](#relational-database-drivers)
    - [NoSQL Database Drivers](#nosql-database-drivers)
    - [Search and Analytic Databases](#search-and-analytic-databases)
  - [Date and Time](#date-and-time)
  - [Distributed Systems](#distributed-systems)
  - [Dynamic DNS](#dynamic-dns)
  - [Email](#email)
  - [Embeddable Scripting Languages](#embeddable-scripting-languages)
  - [Error Handling](#error-handling)
  - [File Handling](#file-handling)
  - [Financial](#financial)
  - [Forms](#forms)
  - [Functional](#functional)
  - [Game Development](#game-development)
  - [Generators](#generators)
  - [Geographic](#geographic)
  - [Go Compilers](#go-compilers)
  - [Goroutines](#goroutines)
  - [GUI](#gui)
  - [Hardware](#hardware)
  - [Images](#images)
  - [IoT (Internet of Things)](#iot-internet-of-things)
  - [Job Scheduler](#job-scheduler)
  - [JSON](#json)
  - [Logging](#logging)
  - [Machine Learning](#machine-learning)
  - [Messaging](#messaging)
  - [Microsoft Office](#microsoft-office)
    - [Microsoft Excel](#microsoft-excel)
    - [Microsoft Word](#microsoft-word)
  - [Miscellaneous](#miscellaneous)
    - [Dependency Injection](#dependency-injection)
    - [Project Layout](#project-layout)
    - [Strings](#strings)
    - [Uncategorized](#uncategorized)
  - [Natural Language Processing](#natural-language-processing)
    - [Language Detection](#language-detection)
    - [Morphological Analyzers](#morphological-analyzers)
    - [Slugifiers](#slugifiers)
    - [Tokenizers](#tokenizers)
    - [Translation](#translation)
    - [Transliteration](#transliteration)
  - [Networking](#networking)
    - [HTTP Clients](#http-clients)
  - [OpenGL](#opengl)
  - [ORM](#orm)
  - [Package Management](#package-management)
  - [Performance](#performance)
  - [Query Language](#query-language)
  - [Reflection](#reflection)
  - [Resource Embedding](#resource-embedding)
  - [Science and Data Analysis](#science-and-data-analysis)
  - [Security](#security)
  - [Serialization](#serialization)
  - [Server Applications](#server-applications)
  - [Stream Processing](#stream-processing)
  - [Template Engines](#template-engines)
  - [Testing](#testing)
    - [Testing Frameworks](#testing-frameworks)
    - [Mock](#mock)
    - [Fuzzing and delta-debugging/reducing/shrinking](#fuzzing-and-delta-debuggingreducingshrinking)
    - [Selenium and browser control tools](#selenium-and-browser-control-tools)
    - [Fail injection](#fail-injection)
  - [Text Processing](#text-processing)
    - [Formatters](#formatters)
    - [Markup Languages](#markup-languages)
    - [Parsers/Encoders/Decoders](#parsersencodersdecoders)
    - [Regular Expressions](#regular-expressions)
    - [Sanitation](#sanitation)
    - [Scrapers](#scrapers)
    - [RSS](#rss)
    - [Utility/Miscellaneous](#utilitymiscellaneous)
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
  - [Webhooks Server](#webhooks-server)
  - [Windows](#windows)
  - [Workflow Frameworks](#workflow-frameworks)
  - [XML](#xml)
  - [Zero Trust](#zero-trust)
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
    - [E-books for purchase](#e-books-for-purchase)
    - [Free e-books](#free-e-books)
  - [Gophers](#gophers)
  - [Meetups](#meetups)
  - [Style Guides](#style-guides)
  - [Social Media](#social-media)
    - [Twitter](#twitter)
    - [Reddit](#reddit)
  - [Websites](#websites)
    - [Tutorials](#tutorials)
    - [Guided Learning](#guided-learning)

**[⬆ back to top](#contents)**



</details>

## Actor Model

_Libraries for building actor-based programs._

- [asyncmachine-go/pkg/machine](https://github.com/pancsta/asyncmachine-go/tree/main/pkg/machine) ![](https://img.shields.io/github/stars/pancsta/asyncmachine-go.svg) - Graph control flow library (AOP, actor, state-machine).
- [Ergo](https://github.com/ergo-services/ergo) ![](https://img.shields.io/github/stars/ergo-services/ergo.svg) - An actor-based Framework with network transparency for creating event-driven architecture in Golang. Inspired by Erlang.
- [Goakt](https://github.com/Tochemey/goakt) ![](https://img.shields.io/github/stars/Tochemey/goakt.svg) - Fast and Distributed Actor framework using protocol buffers as message for Golang.
- [Hollywood](https://github.com/anthdm/hollywood) ![](https://img.shields.io/github/stars/anthdm/hollywood.svg) - Blazingly fast and light-weight Actor engine written in Golang.
- [ProtoActor](https://github.com/asynkron/protoactor-go) ![](https://img.shields.io/github/stars/asynkron/protoactor-go.svg) - Distributed actors for Go, C#, and Java/Kotlin.

**[⬆ back to top](#contents)**

## Artificial Intelligence

_Libraries for building programs that leverage AI._

- [chromem-go](https://github.com/philippgille/chromem-go) ![](https://img.shields.io/github/stars/philippgille/chromem-go.svg) - Embeddable vector database for Go with Chroma-like interface and zero third-party dependencies. In-memory with optional persistence.
- [fun](https://gitlab.com/tozd/go/fun) - The simplest but powerful way to use large language models (LLMs) in Go.
- [langchaingo](https://github.com/tmc/langchaingo) ![](https://img.shields.io/github/stars/tmc/langchaingo.svg) - LangChainGo is a framework for developing applications powered by language models.
- [LocalAI](https://github.com/mudler/LocalAI) ![](https://img.shields.io/github/stars/mudler/LocalAI.svg) - Open Source OpenAI alternative, self-host AI models.
- [Ollama](https://github.com/jmorganca/ollama) ![](https://img.shields.io/github/stars/jmorganca/ollama.svg) - Run large language models locally.
- [OllamaFarm](https://github.com/presbrey/ollamafarm) ![](https://img.shields.io/github/stars/presbrey/ollamafarm.svg) - Manage, load-balance, and failover packs of Ollamas.

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio._

- [beep](https://github.com/gopxl/beep) ![](https://img.shields.io/github/stars/gopxl/beep.svg) - A simple library for playback and audio manipulation.
- [flac](https://github.com/mewkiz/flac) ![](https://img.shields.io/github/stars/mewkiz/flac.svg) - Native Go FLAC encoder/decoder with support for FLAC streams.
- [gaad](https://github.com/Comcast/gaad) ![](https://img.shields.io/github/stars/Comcast/gaad.svg) - Native Go AAC bitstream parser.
- [go-mpris](https://github.com/leberKleber/go-mpris) ![](https://img.shields.io/github/stars/leberKleber/go-mpris.svg) - Client for mpris dbus interfaces.
- [GoAudio](https://github.com/DylanMeeus/GoAudio) ![](https://img.shields.io/github/stars/DylanMeeus/GoAudio.svg) - Native Go Audio Processing Library.
- [gosamplerate](https://github.com/dh1tw/gosamplerate) ![](https://img.shields.io/github/stars/dh1tw/gosamplerate.svg) - libsamplerate bindings for go.
- [id3v2](https://github.com/bogem/id3v2) ![](https://img.shields.io/github/stars/bogem/id3v2.svg) - ID3 decoding and encoding library for Go.
- [malgo](https://github.com/gen2brain/malgo) ![](https://img.shields.io/github/stars/gen2brain/malgo.svg) - Mini audio library.
- [minimp3](https://github.com/tosone/minimp3) ![](https://img.shields.io/github/stars/tosone/minimp3.svg) - Lightweight MP3 decoder library.
- [Oto](https://github.com/hajimehoshi/oto) ![](https://img.shields.io/github/stars/hajimehoshi/oto.svg) - A low-level library to play sound on multiple platforms.
- [PortAudio](https://github.com/gordonklaus/portaudio) ![](https://img.shields.io/github/stars/gordonklaus/portaudio.svg) - Go bindings for the PortAudio audio I/O library.

**[⬆ back to top](#contents)**

## Authentication and Authorization

_Libraries for implementing authentication and authorization._

- [authboss](https://github.com/volatiletech/authboss) ![](https://img.shields.io/github/stars/volatiletech/authboss.svg) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time.
- [branca](https://github.com/essentialkaos/branca) ![](https://img.shields.io/github/stars/essentialkaos/branca.svg) - branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+.
- [casbin](https://github.com/hsluoyz/casbin) ![](https://img.shields.io/github/stars/hsluoyz/casbin.svg) - Authorization library that supports access control models like ACL, RBAC, and ABAC.
- [cookiestxt](https://github.com/mengzhuo/cookiestxt) ![](https://img.shields.io/github/stars/mengzhuo/cookiestxt.svg) - provides a parser of cookies.txt file format.
- [go-githubauth](https://github.com/jferrl/go-githubauth) ![](https://img.shields.io/github/stars/jferrl/go-githubauth.svg) - Utilities for GitHub authentication: generate and use GitHub application and installation tokens.
- [go-guardian](https://github.com/shaj13/go-guardian) ![](https://img.shields.io/github/stars/shaj13/go-guardian.svg) - Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token, and Certificate based authentication.
- [go-iam](https://github.com/melvinodsa/go-iam) ![](https://img.shields.io/github/stars/melvinodsa/go-iam.svg) - Developer-first Identity and Access Management system with a simple UI.
- [go-jose](https://github.com/go-jose/go-jose) ![](https://img.shields.io/github/stars/go-jose/go-jose.svg) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
- [go-jwt](https://github.com/pardnchiu/go-jwt) ![](https://img.shields.io/github/stars/pardnchiu/go-jwt.svg) - JWT authentication package providing access tokens and refresh tokens with fingerprinting, Redis storage, and automatic refresh capabilities.
- [goiabada](https://github.com/leodip/goiabada) ![](https://img.shields.io/github/stars/leodip/goiabada.svg) - An open-source authentication and authorization server supporting OAuth2 and OpenID Connect.
- [gologin](https://github.com/dghubble/gologin) ![](https://img.shields.io/github/stars/dghubble/gologin.svg) - chainable handlers for login with OAuth1 and OAuth2 authentication providers.
- [gorbac](https://github.com/mikespook/gorbac) ![](https://img.shields.io/github/stars/mikespook/gorbac.svg) - provides a lightweight role-based access control (RBAC) implementation in Golang.
- [gosession](https://github.com/Kwynto/gosession) ![](https://img.shields.io/github/stars/Kwynto/gosession.svg) - This is quick session for net/http in GoLang. This package is perhaps the best implementation of the session mechanism, or at least it tries to become one.
- [goth](https://github.com/markbates/goth) ![](https://img.shields.io/github/stars/markbates/goth.svg) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.
- [jeff](https://github.com/abraithwaite/jeff) ![](https://img.shields.io/github/stars/abraithwaite/jeff.svg) - Simple, flexible, secure, and idiomatic web session management with pluggable backends.
- [jwt](https://github.com/pascaldekloe/jwt) ![](https://img.shields.io/github/stars/pascaldekloe/jwt.svg) - Lightweight JSON Web Token (JWT) library.
- [jwt](https://github.com/cristalhq/jwt) ![](https://img.shields.io/github/stars/cristalhq/jwt.svg) - Safe, simple, and fast JSON Web Tokens for Go.
- [jwt-auth](https://github.com/adam-hanna/jwt-auth) ![](https://img.shields.io/github/stars/adam-hanna/jwt-auth.svg) - JWT middleware for Golang http servers with many configuration options.
- [jwt-go](https://github.com/golang-jwt/jwt) ![](https://img.shields.io/github/stars/golang-jwt/jwt.svg) - A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs.
- [jwx](https://github.com/lestrrat-go/jwx) ![](https://img.shields.io/github/stars/lestrrat-go/jwx.svg) - Go module implementing various JWx (JWA/JWE/JWK/JWS/JWT, otherwise known as JOSE) technologies.
- [keto](https://github.com/ory/keto) ![](https://img.shields.io/github/stars/ory/keto.svg) - Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models.
- [loginsrv](https://github.com/tarent/loginsrv) ![](https://img.shields.io/github/stars/tarent/loginsrv.svg) - JWT login microservice with pluggable backends such as OAuth2 (Github), htpasswd, osiam.
- [oauth2](https://github.com/golang/oauth2) ![](https://img.shields.io/github/stars/golang/oauth2.svg) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support.
- [oidc](https://github.com/zitadel/oidc) ![](https://img.shields.io/github/stars/zitadel/oidc.svg) - Easy to use OpenID Connect client and server library written for Go and certified by the OpenID Foundation.
- [openfga](https://github.com/openfga/openfga) ![](https://img.shields.io/github/stars/openfga/openfga.svg) - Implementation of fine-grained authorization based on the "Zanzibar: Google's Consistent, Global Authorization System" paper. Backed by [CNCF](https://www.cncf.io/).
- [osin](https://github.com/openshift/osin) ![](https://img.shields.io/github/stars/openshift/osin.svg) - Golang OAuth2 server library.
- [otpgen](https://github.com/grijul/otpgen) ![](https://img.shields.io/github/stars/grijul/otpgen.svg) - Library to generate TOTP/HOTP codes.
- [otpgo](https://github.com/jltorresm/otpgo) ![](https://img.shields.io/github/stars/jltorresm/otpgo.svg) - Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go.
- [paseto](https://github.com/o1egl/paseto) ![](https://img.shields.io/github/stars/o1egl/paseto.svg) - Golang implementation of Platform-Agnostic Security Tokens (PASETO).
- [permissions](https://github.com/xyproto/permissions) ![](https://img.shields.io/github/stars/xyproto/permissions.svg) - Library for keeping track of users, login states, and permissions. Uses secure cookies and bcrypt.
- [scope](https://github.com/SonicRoshan/scope) ![](https://img.shields.io/github/stars/SonicRoshan/scope.svg) - Easily Manage OAuth2 Scopes In Go.
- [scs](https://github.com/alexedwards/scs) ![](https://img.shields.io/github/stars/alexedwards/scs.svg) - Session Manager for HTTP servers.
- [securecookie](https://github.com/chmike/securecookie) ![](https://img.shields.io/github/stars/chmike/securecookie.svg) - Efficient secure cookie encoding/decoding.
- [session](https://github.com/icza/session) ![](https://img.shields.io/github/stars/icza/session.svg) - Go session management for web servers (including support for Google App Engine - GAE).
- [sessions](https://github.com/adam-hanna/sessions) ![](https://img.shields.io/github/stars/adam-hanna/sessions.svg) - Dead simple, highly performant, highly customizable sessions service for go http servers.
- [sessionup](https://github.com/swithek/sessionup) ![](https://img.shields.io/github/stars/swithek/sessionup.svg) - Simple, yet effective HTTP session management and identification package.
- [sjwt](https://github.com/brianvoe/sjwt) ![](https://img.shields.io/github/stars/brianvoe/sjwt.svg) - Simple jwt generator and parser.
- [spicedb](https://github.com/authzed/spicedb) ![](https://img.shields.io/github/stars/authzed/spicedb.svg) - A Zanzibar-inspired database that enables fine-grained authorization.
- [x509proxy](https://github.com/vkuznet/x509proxy) ![](https://img.shields.io/github/stars/vkuznet/x509proxy.svg) - Library to handle X509 proxy certificates.

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [cometbft](https://github.com/cometbft/cometbft) ![](https://img.shields.io/github/stars/cometbft/cometbft.svg) - A distributed, Byzantine fault-tolerant, deterministic state machine replication engine. It is a fork of Tendermint Core and implements the Tendermint consensus algorithm.
- [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) ![](https://img.shields.io/github/stars/cosmos/cosmos-sdk.svg) - A Framework for Building Public Blockchains in the Cosmos Ecosystem.
- [gno](https://github.com/gnolang/gno) ![](https://img.shields.io/github/stars/gnolang/gno.svg) - A comprehensive smart contract suite built with Golang and Gnolang, a deterministic, purpose-built Go variant for blockchains.
- [go-ethereum](https://github.com/ethereum/go-ethereum) ![](https://img.shields.io/github/stars/ethereum/go-ethereum.svg) - Official Go implementation of the Ethereum protocol.
- [gosemble](https://github.com/LimeChain/gosemble) ![](https://img.shields.io/github/stars/LimeChain/gosemble.svg) - A Go-based framework for building Polkadot/Substrate-compatible runtimes.
- [gossamer](https://github.com/ChainSafe/gossamer) ![](https://img.shields.io/github/stars/ChainSafe/gossamer.svg) - A Go implementation of the Polkadot Host.
- [kubo](https://github.com/ipfs/kubo) ![](https://img.shields.io/github/stars/ipfs/kubo.svg) - An IPFS implementation in Go. It provides content-addressable storage which can be used for decentralized storage in DApps. It is based on the IPFS protocol.
- [lnd](https://github.com/lightningnetwork/lnd) ![](https://img.shields.io/github/stars/lightningnetwork/lnd.svg) - A complete implementation of a Lightning Network node.
- [nview](https://github.com/blinklabs-io/nview) ![](https://img.shields.io/github/stars/blinklabs-io/nview.svg) - Local monitoring tool for a Cardano Node. It's a TUI (terminal user interface) designed to fit most screens.
- [solana-go](https://github.com/gagliardetto/solana-go) ![](https://img.shields.io/github/stars/gagliardetto/solana-go.svg) - Go library to interface with Solana JSON RPC and WebSocket interfaces.
- [tendermint](https://github.com/tendermint/tendermint) ![](https://img.shields.io/github/stars/tendermint/tendermint.svg) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [arikawa](https://github.com/diamondburned/arikawa) ![](https://img.shields.io/github/stars/diamondburned/arikawa.svg) - A library and framework for the Discord API.
- [bot](https://github.com/go-telegram/bot) ![](https://img.shields.io/github/stars/go-telegram/bot.svg) - Zero-dependencies Telegram Bot library with additional UI components.
- [echotron](https://github.com/NicoNex/echotron) ![](https://img.shields.io/github/stars/NicoNex/echotron.svg) - An elegant and concurrent library for Telegram Bots in Go.
- [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.
- [go-sarah](https://github.com/oklahomer/go-sarah) ![](https://img.shields.io/github/stars/oklahomer/go-sarah.svg) - Framework to build a bot for desired chat services including LINE, Slack, Gitter, and more.
- [go-tg](https://github.com/mr-linch/go-tg) ![](https://img.shields.io/github/stars/mr-linch/go-tg.svg) - Generated from official docs Go client library for accessing Telegram Bot API, with batteries for building complex bots included.
- [go-twitch-irc](https://github.com/gempir/go-twitch-irc) ![](https://img.shields.io/github/stars/gempir/go-twitch-irc.svg) - Library to write bots for twitch.tv chat
- [micha](https://github.com/onrik/micha) ![](https://img.shields.io/github/stars/onrik/micha.svg) - Go Library for Telegram bot api.
- [slack-bot](https://github.com/innogames/slack-bot) ![](https://img.shields.io/github/stars/innogames/slack-bot.svg) - Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github...
- [slacker](https://github.com/slack-io/slacker) ![](https://img.shields.io/github/stars/slack-io/slacker.svg) - Easy to use framework to create Slack bots.
- [telebot](https://github.com/tucnak/telebot) ![](https://img.shields.io/github/stars/tucnak/telebot.svg) - Telegram bot framework is written in Go.
- [telego](https://github.com/mymmrac/telego) ![](https://img.shields.io/github/stars/mymmrac/telego.svg) - Telegram Bot API library for Golang with full one-to-one API implementation.
- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) ![](https://img.shields.io/github/stars/go-telegram-bot-api/telegram-bot-api.svg) - Simple and clean Telegram bot client.
- [wayback](https://github.com/wabarc/wayback) ![](https://img.shields.io/github/stars/wabarc/wayback.svg) - A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages.

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [1build](https://github.com/gopinath-langote/1build) ![](https://img.shields.io/github/stars/gopinath-langote/1build.svg) - Command line tool to frictionlessly manage project-specific commands.
- [air](https://github.com/cosmtrek/air) ![](https://img.shields.io/github/stars/cosmtrek/air.svg) - Air - Live reload for Go apps.
- [anko](https://github.com/GuilhermeCaruso/anko) ![](https://img.shields.io/github/stars/GuilhermeCaruso/anko.svg) - Simple application watcher for multiple programming languages.
- [gaper](https://github.com/maxclaus/gaper) ![](https://img.shields.io/github/stars/maxclaus/gaper.svg) - Builds and restarts a Go project when it crashes or some watched file changes.
- [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.
- [gob](https://github.com/kcmvp/gob) ![](https://img.shields.io/github/stars/kcmvp/gob.svg) - [Gradle](https://docs.gradle.org/)/[Maven](https://maven.apache.org/) like build tool for Go projects.
- [goyek](https://github.com/goyek/goyek) ![](https://img.shields.io/github/stars/goyek/goyek.svg) - Create build pipelines in Go.
- [mage](https://github.com/magefile/mage) ![](https://img.shields.io/github/stars/magefile/mage.svg) - Mage is a make/rake-like build tool using Go.
- [mmake](https://github.com/tj/mmake) ![](https://img.shields.io/github/stars/tj/mmake.svg) - Modern Make.
- [realize](https://github.com/tockins/realize) ![](https://img.shields.io/github/stars/tockins/realize.svg) - Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths.
- [Task](https://github.com/go-task/task) ![](https://img.shields.io/github/stars/go-task/task.svg) - simple "Make" alternative.
- [taskctl](https://github.com/taskctl/taskctl) ![](https://img.shields.io/github/stars/taskctl/taskctl.svg) - Concurrent task runner.
- [xc](https://github.com/joerdav/xc) ![](https://img.shields.io/github/stars/joerdav/xc.svg) - Task runner with README.md defined tasks, executable markdown.

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [asciigraph](https://github.com/guptarohit/asciigraph) ![](https://img.shields.io/github/stars/guptarohit/asciigraph.svg) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.
- [aurora](https://github.com/logrusorgru/aurora) ![](https://img.shields.io/github/stars/logrusorgru/aurora.svg) - ANSI terminal colors that support fmt.Printf/Sprintf.
- [box-cli-maker](https://github.com/Delta456/box-cli-maker) ![](https://img.shields.io/github/stars/Delta456/box-cli-maker.svg) - Make Highly Customized Boxes for your CLI.
- [bubble-table](https://github.com/Evertras/bubble-table) ![](https://img.shields.io/github/stars/Evertras/bubble-table.svg) - An interactive table component for bubbletea.
- [bubbles](https://github.com/charmbracelet/bubbles) ![](https://img.shields.io/github/stars/charmbracelet/bubbles.svg) - TUI components for bubbletea.
- [bubbletea](https://github.com/charmbracelet/bubbletea) ![](https://img.shields.io/github/stars/charmbracelet/bubbletea.svg) - Go framework to build terminal apps, based on The Elm Architecture.
- [crab-config-files-templating](https://github.com/alfiankan/crab-config-files-templating) ![](https://img.shields.io/github/stars/alfiankan/crab-config-files-templating.svg) - Dynamic configuration file templating tool for kubernetes manifest or general configuration files.
- [ctc](https://github.com/wzshiming/ctc) ![](https://img.shields.io/github/stars/wzshiming/ctc.svg) - The non-invasive cross-platform terminal color library does not need to modify the Print method.
- [fx](https://github.com/antonmedv/fx) ![](https://img.shields.io/github/stars/antonmedv/fx.svg) - Terminal JSON viewer & processor.
- [go-ataman](https://github.com/workanator/go-ataman) ![](https://img.shields.io/github/stars/workanator/go-ataman.svg) - Go library for rendering ANSI colored text templates in terminals.
- [go-colorable](https://github.com/mattn/go-colorable) ![](https://img.shields.io/github/stars/mattn/go-colorable.svg) - Colorable writer for windows.
- [go-colortext](https://github.com/daviddengcn/go-colortext) ![](https://img.shields.io/github/stars/daviddengcn/go-colortext.svg) - Go library for color output in terminals.
- [go-isatty](https://github.com/mattn/go-isatty) ![](https://img.shields.io/github/stars/mattn/go-isatty.svg) - isatty for golang.
- [go-palette](https://github.com/abusomani/go-palette) ![](https://img.shields.io/github/stars/abusomani/go-palette.svg) - Go library that provides elegant and convenient style definitions using ANSI colors. Fully compatible & wraps the [fmt library](https://pkg.go.dev/fmt) for nice terminal layouts.
- [go-prompt](https://github.com/c-bata/go-prompt) ![](https://img.shields.io/github/stars/c-bata/go-prompt.svg) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).
- [gocui](https://github.com/jroimartin/gocui) ![](https://img.shields.io/github/stars/jroimartin/gocui.svg) - Minimalist Go library aimed at creating Console User Interfaces.
- [gommon/color](https://github.com/labstack/gommon/tree/master/color) ![](https://img.shields.io/github/stars/labstack/gommon.svg) - Style terminal text.
- [gookit/color](https://github.com/gookit/color) ![](https://img.shields.io/github/stars/gookit/color.svg) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.
- [lipgloss](https://github.com/charmbracelet/lipgloss) ![](https://img.shields.io/github/stars/charmbracelet/lipgloss.svg) - Declaratively define styles for color, format and layout in the terminal.
- [marker](https://github.com/cyucelen/marker) ![](https://img.shields.io/github/stars/cyucelen/marker.svg) - Easiest way to match and mark strings for colorful terminal outputs.
- [mpb](https://github.com/vbauerster/mpb) ![](https://img.shields.io/github/stars/vbauerster/mpb.svg) - Multi progress bar for terminal applications.
- [progressbar](https://github.com/schollz/progressbar) ![](https://img.shields.io/github/stars/schollz/progressbar.svg) - Basic thread-safe progress bar that works in every OS.
- [pterm](https://github.com/pterm/pterm) ![](https://img.shields.io/github/stars/pterm/pterm.svg) - A library to beautify console output on every platform with many combinable components.
- [simpletable](https://github.com/alexeyco/simpletable) ![](https://img.shields.io/github/stars/alexeyco/simpletable.svg) - Simple tables in a terminal with Go.
- [spinner](https://github.com/briandowns/spinner) ![](https://img.shields.io/github/stars/briandowns/spinner.svg) - Go package to easily provide a terminal spinner with options.
- [tabby](https://github.com/cheynewallace/tabby) ![](https://img.shields.io/github/stars/cheynewallace/tabby.svg) - A tiny library for super simple Golang tables.
- [table](https://github.com/tomlazar/table) ![](https://img.shields.io/github/stars/tomlazar/table.svg) - Small library for terminal color based tables.
- [termbox-go](https://github.com/nsf/termbox-go) ![](https://img.shields.io/github/stars/nsf/termbox-go.svg) - Termbox is a library for creating cross-platform text-based interfaces.
- [termdash](https://github.com/mum4k/termdash) ![](https://img.shields.io/github/stars/mum4k/termdash.svg) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).
- [termenv](https://github.com/muesli/termenv) ![](https://img.shields.io/github/stars/muesli/termenv.svg) - Advanced ANSI style & color support for your terminal applications.
- [termui](https://github.com/gizak/termui) ![](https://img.shields.io/github/stars/gizak/termui.svg) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).
- [uilive](https://github.com/gosuri/uilive) ![](https://img.shields.io/github/stars/gosuri/uilive.svg) - Library for updating terminal output in real time.
- [uiprogress](https://github.com/gosuri/uiprogress) ![](https://img.shields.io/github/stars/gosuri/uiprogress.svg) - Flexible library to render progress bars in terminal applications.
- [uitable](https://github.com/gosuri/uitable) ![](https://img.shields.io/github/stars/gosuri/uitable.svg) - Library to improve readability in terminal apps using tabular data.
- [yacspin](https://github.com/theckman/yacspin) ![](https://img.shields.io/github/stars/theckman/yacspin.svg) - Yet Another CLi Spinner package, for working with terminal spinners.

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [acmd](https://github.com/cristalhq/acmd) ![](https://img.shields.io/github/stars/cristalhq/acmd.svg) - Simple, useful, and opinionated CLI package in Go.
- [argparse](https://github.com/akamensky/argparse) ![](https://img.shields.io/github/stars/akamensky/argparse.svg) - Command line argument parser inspired by Python's argparse module.
- [argv](https://github.com/cosiner/argv) ![](https://img.shields.io/github/stars/cosiner/argv.svg) - Go library to split command line string as arguments array using the bash syntax.
- [carapace](https://github.com/rsteube/carapace) ![](https://img.shields.io/github/stars/rsteube/carapace.svg) - Command argument completion generator for spf13/cobra.
- [carapace-bin](https://github.com/rsteube/carapace-bin) ![](https://img.shields.io/github/stars/rsteube/carapace-bin.svg) - Multi-shell multi-command argument completer.
- [carapace-spec](https://github.com/rsteube/carapace-spec) ![](https://img.shields.io/github/stars/rsteube/carapace-spec.svg) - Define simple completions using a spec file.
- [climax](https://github.com/tucnak/climax) ![](https://img.shields.io/github/stars/tucnak/climax.svg) - Alternative CLI with "human face", in spirit of Go command.
- [clîr](https://github.com/leaanthony/clir) ![](https://img.shields.io/github/stars/leaanthony/clir.svg) - A Simple and Clear CLI library. Dependency free.
- [cmd](https://github.com/posener/cmd) ![](https://img.shields.io/github/stars/posener/cmd.svg) - Extends the standard `flag` package to support sub commands and more in idiomatic way.
- [cmdr](https://github.com/hedzr/cmdr) ![](https://img.shields.io/github/stars/hedzr/cmdr.svg) - A POSIX/GNU style, getopt-like command-line UI Go library.
- [cobra](https://github.com/spf13/cobra) ![](https://img.shields.io/github/stars/spf13/cobra.svg) - Commander for modern Go CLI interactions.
- [command-chain](https://github.com/rainu/go-command-chain) ![](https://img.shields.io/github/stars/rainu/go-command-chain.svg) - A go library for configure and run command chains - such as pipelining in unix shells.
- [commandeer](https://github.com/jaffee/commandeer) ![](https://img.shields.io/github/stars/jaffee/commandeer.svg) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.
- [complete](https://github.com/posener/complete) ![](https://img.shields.io/github/stars/posener/complete.svg) - Write bash completions in Go + Go command bash completion.
- [console](https://github.com/reeflective/console) ![](https://img.shields.io/github/stars/reeflective/console.svg) Closed-loop application library for Cobra commands, with oh-my-posh prompts, and more.
- [Dnote](https://github.com/dnote/dnote) ![](https://img.shields.io/github/stars/dnote/dnote.svg) - A simple command line notebook with multi-device sync.
- [elvish](https://github.com/elves/elvish) ![](https://img.shields.io/github/stars/elves/elvish.svg) - An expressive programming language and a versatile interactive shell.
- [env](https://github.com/codingconcepts/env) ![](https://img.shields.io/github/stars/codingconcepts/env.svg) - Tag-based environment configuration for structs.
- [flaggy](https://github.com/integrii/flaggy) ![](https://img.shields.io/github/stars/integrii/flaggy.svg) - A robust and idiomatic flags package with excellent subcommand support.
- [flagvar](https://github.com/sgreben/flagvar) ![](https://img.shields.io/github/stars/sgreben/flagvar.svg) - A collection of flag argument types for Go's standard `flag` package.
- [flash-flags](https://github.com/agilira/flash-flags) ![](https://img.shields.io/github/stars/agilira/flash-flags.svg) - Ultra-fast, zero-dependency, POSIX-compliant flag parsing library that can be used as drop-in stdlib replacement with security hardening.
- [getopt](https://github.com/jon-codes/getopt) ![](https://img.shields.io/github/stars/jon-codes/getopt.svg) - An accurate Go `getopt`, validated against the GNU libc implementation.
- [go-arg](https://github.com/alexflint/go-arg) ![](https://img.shields.io/github/stars/alexflint/go-arg.svg) - Struct-based argument parsing in Go.
- [go-flags](https://github.com/jessevdk/go-flags) ![](https://img.shields.io/github/stars/jessevdk/go-flags.svg) - go command line option parser.
- [go-getoptions](https://github.com/DavidGamba/go-getoptions) ![](https://img.shields.io/github/stars/DavidGamba/go-getoptions.svg) - Go option parser inspired by the flexibility of Perl’s GetOpt::Long.
- [go-readline-ny](https://github.com/nyaosorg/go-readline-ny) ![](https://img.shields.io/github/stars/nyaosorg/go-readline-ny.svg) - A customizable line-editing library with Emacs keybindings, Unicode support, completion, and syntax highlighting. Used in NYAGOS shell.
- [gocmd](https://github.com/devfacet/gocmd) ![](https://img.shields.io/github/stars/devfacet/gocmd.svg) - Go library for building command line applications.
- [goopt](https://github.com/napalu/goopt) ![](https://img.shields.io/github/stars/napalu/goopt.svg) - A declarative, struct-tag based CLI framework for Go, with a broad feature set such as hierarchical commands/flags, i18n, shell completion, and validation.
- [hashicorp/cli](https://github.com/hashicorp/cli) ![](https://img.shields.io/github/stars/hashicorp/cli.svg) - Go library for implementing command-line interfaces.
- [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) ![](https://img.shields.io/github/stars/hidevopsio/hiboot.svg) - cli application framework with auto configuration and dependency injection.
- [job](https://github.com/liujianping/job) ![](https://img.shields.io/github/stars/liujianping/job.svg) - JOB, make your short-term command as a long-term job.
- [kingpin](https://github.com/alecthomas/kingpin) ![](https://img.shields.io/github/stars/alecthomas/kingpin.svg) - Command line and flag parser supporting sub commands (superseded by `kong`; see below).
- [liner](https://github.com/peterh/liner) ![](https://img.shields.io/github/stars/peterh/liner.svg) - Go readline-like library for command-line interfaces.
- [mcli](https://github.com/jxskiss/mcli) ![](https://img.shields.io/github/stars/jxskiss/mcli.svg) - A minimal but very powerful cli library for Go.
- [mkideal/cli](https://github.com/mkideal/cli) ![](https://img.shields.io/github/stars/mkideal/cli.svg) - Feature-rich and easy to use command-line package based on golang struct tags.
- [mow.cli](https://github.com/jawher/mow.cli) ![](https://img.shields.io/github/stars/jawher/mow.cli.svg) - Go library for building CLI applications with sophisticated flag and argument parsing and validation.
- [ops](https://github.com/nanovms/ops) ![](https://img.shields.io/github/stars/nanovms/ops.svg) - Unikernel Builder/Orchestrator.
- [orpheus](https://github.com/agilira/orpheus) ![](https://img.shields.io/github/stars/agilira/orpheus.svg) - CLI framework with security hardening, plugin storage system, and production observability features.
- [pflag](https://github.com/spf13/pflag) ![](https://img.shields.io/github/stars/spf13/pflag.svg) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.
- [readline](https://github.com/reeflective/readline) ![](https://img.shields.io/github/stars/reeflective/readline.svg) - Shell library with modern and easy to use UI features.
- [sflags](https://github.com/octago/sflags) ![](https://img.shields.io/github/stars/octago/sflags.svg) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin, and other libraries.
- [structcli](https://github.com/leodido/structcli) ![](https://img.shields.io/github/stars/leodido/structcli.svg) - Eliminate Cobra boilerplate: build powerful, feature-rich CLIs declaratively from Go structs.
- [strumt](https://github.com/antham/strumt) ![](https://img.shields.io/github/stars/antham/strumt.svg) - Library to create prompt chain.
- [subcmd](https://github.com/bobg/subcmd) ![](https://img.shields.io/github/stars/bobg/subcmd.svg) - Another approach to parsing and running subcommands. Works alongside the standard `flag` package.
- [teris-io/cli](https://github.com/teris-io/cli) ![](https://img.shields.io/github/stars/teris-io/cli.svg) - Simple and complete API for building command line interfaces in Go.
- [urfave/cli](https://github.com/urfave/cli) ![](https://img.shields.io/github/stars/urfave/cli.svg) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).
- [version](https://github.com/mszostok/version) ![](https://img.shields.io/github/stars/mszostok/version.svg) - Collects and displays CLI version information in multiple formats along with upgrade notice.
- [wlog](https://github.com/dixonwille/wlog) ![](https://img.shields.io/github/stars/dixonwille/wlog.svg) - Simple logging interface that supports cross-platform color and concurrency.
- [wmenu](https://github.com/dixonwille/wmenu) ![](https://img.shields.io/github/stars/dixonwille/wmenu.svg) - Easy to use menu structure for cli applications that prompt users to make choices.

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [aconfig](https://github.com/cristalhq/aconfig) ![](https://img.shields.io/github/stars/cristalhq/aconfig.svg) - Simple, useful and opinionated config loader.
- [argus](https://github.com/agilira/argus) ![](https://img.shields.io/github/stars/agilira/argus.svg) - File watching and configuration management with MPSC ring buffer, adaptive batching strategies, and universal format parsing (JSON, YAML, TOML, INI, HCL, Properties).
- [azureappconfiguration](https://github.com/Azure/AppConfiguration-GoProvider) ![](https://img.shields.io/github/stars/Azure/AppConfiguration-GoProvider.svg) - The configuration provider for consuming data in Azure App Configuration from Go applications.
- [bcl](https://github.com/wkhere/bcl) ![](https://img.shields.io/github/stars/wkhere/bcl.svg) - BCL is a configuration language similar to HCL.
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) ![](https://img.shields.io/github/stars/ilyakaznacheev/cleanenv.svg) - Minimalistic configuration reader (from files, ENV, and wherever you want).
- [config](https://github.com/JeremyLoy/config) ![](https://img.shields.io/github/stars/JeremyLoy/config.svg) - Cloud native application configuration. Bind ENV to structs in only two lines.
- [config](https://github.com/num30/config) ![](https://img.shields.io/github/stars/num30/config.svg) - configure your app using file, environment variables, or flags in two lines of code.
- [configuration](https://github.com/BoRuDar/configuration) ![](https://img.shields.io/github/stars/BoRuDar/configuration.svg) - Library for initializing configuration structs from env variables, files, flags and 'default' tag.
- [configuro](https://github.com/sherifabdlnaby/configuro) ![](https://img.shields.io/github/stars/sherifabdlnaby/configuro.svg) - opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications.
- [confiq](https://github.com/greencoda/confiq) ![](https://img.shields.io/github/stars/greencoda/confiq.svg) - Structured data format to config struct decoder library for Go - supporting multiple data formats.
- [confita](https://github.com/heetch/confita) ![](https://img.shields.io/github/stars/heetch/confita.svg) - Load configuration in cascade from multiple backends into a struct.
- [conflate](https://github.com/the4thamigo-uk/conflate) ![](https://img.shields.io/github/stars/the4thamigo-uk/conflate.svg) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
- [enflag](https://github.com/atelpis/enflag) ![](https://img.shields.io/github/stars/atelpis/enflag.svg) - Container-oriented, zero-dependency configuration library that unifies Env variable and Flag parsing. Uses generics for type safety, without reflection or struct tags.
- [env](https://github.com/caarlos0/env) ![](https://img.shields.io/github/stars/caarlos0/env.svg) - Parse environment variables to Go structs (with defaults).
- [env](https://github.com/junk1tm/env) ![](https://img.shields.io/github/stars/junk1tm/env.svg) - A lightweight package for loading environment variables into structs.
- [env](https://github.com/syntaqx/env) ![](https://img.shields.io/github/stars/syntaqx/env.svg) - An environment utility package with support for unmarshaling into structs.
- [envconfig](https://github.com/vrischmann/envconfig) ![](https://img.shields.io/github/stars/vrischmann/envconfig.svg) - Read your configuration from environment variables.
- [envh](https://github.com/antham/envh) ![](https://img.shields.io/github/stars/antham/envh.svg) - Helpers to manage environment variables.
- [envyaml](https://github.com/yuseferi/envyaml) ![](https://img.shields.io/github/stars/yuseferi/envyaml.svg) - Yaml with environment variables reader. it helps to have secrets as environment variable but load them configs as structured Yaml.
- [fig](https://github.com/kkyr/fig) ![](https://img.shields.io/github/stars/kkyr/fig.svg) - Tiny library for reading configuration from a file and from environment variables (with validation & defaults).
- [genv](https://github.com/sakirsensoy/genv) ![](https://img.shields.io/github/stars/sakirsensoy/genv.svg) - Read environment variables easily with dotenv support.
- [go-array](https://github.com/deatil/go-array) ![](https://img.shields.io/github/stars/deatil/go-array.svg) - A Go package that read or set data from map, slice or json.
- [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) ![](https://img.shields.io/github/stars/PaddleHQ/go-aws-ssm.svg) - Go package that fetches parameters from AWS System Manager - Parameter Store.
- [go-cfg](https://github.com/dsbasko/go-cfg) ![](https://img.shields.io/github/stars/dsbasko/go-cfg.svg) - The library provides a unified way to read configuration data into a structure from various sources, such as env, flags, and configuration files (.json, .yaml, .toml, .env).
- [go-conf](https://github.com/ThomasObenaus/go-conf) ![](https://img.shields.io/github/stars/ThomasObenaus/go-conf.svg) - Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters.
- [go-config](https://github.com/MordaTeam/go-config) ![](https://img.shields.io/github/stars/MordaTeam/go-config.svg) - Simple and convenient library for working with app configurations.
- [go-ini](https://github.com/subpop/go-ini) ![](https://img.shields.io/github/stars/subpop/go-ini.svg) - A Go package that marshals and unmarshals INI-files.
- [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) ![](https://img.shields.io/github/stars/ianlopshire/go-ssm-config.svg) - Go utility for loading configuration parameters from AWS SSM (Parameter Store).
- [go-up](https://github.com/ufoscout/go-up) ![](https://img.shields.io/github/stars/ufoscout/go-up.svg) - A simple configuration library with recursive placeholders resolution and no magic.
- [GoCfg](https://github.com/Jagerente/gocfg) ![](https://img.shields.io/github/stars/Jagerente/gocfg.svg) - Config manager with Struct Tags based contracts, custom value providers, parsers, and documentation generation. Customizable yet simple.
- [godotenv](https://github.com/joho/godotenv) ![](https://img.shields.io/github/stars/joho/godotenv.svg) - Go port of Ruby's dotenv library (Loads environment variables from `.env`).
- [GoLobby/Config](https://github.com/golobby/config) ![](https://img.shields.io/github/stars/golobby/config.svg) - GoLobby Config is a lightweight yet powerful configuration manager for the Go programming language.
- [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) ![](https://img.shields.io/github/stars/One-com/gone.svg) - Modular JSON configuration. Keep your config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
- [gonfig](https://github.com/milad-abbasi/gonfig) ![](https://img.shields.io/github/stars/milad-abbasi/gonfig.svg) - Tag-based configuration parser which loads values from different providers into typesafe struct.
- [gookit/config](https://github.com/gookit/config) ![](https://img.shields.io/github/stars/gookit/config.svg) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.
- [harvester](https://github.com/beatlabs/harvester) ![](https://img.shields.io/github/stars/beatlabs/harvester.svg) - Harvester, a easy to use static and dynamic configuration package supporting seeding, env vars and Consul integration.
- [hedzr/store](https://github.com/hedzr/store) ![](https://img.shields.io/github/stars/hedzr/store.svg) - Extensible, high-performance configuration management library, optimized for hierarchical data.
- [hjson](https://github.com/hjson/hjson-go) ![](https://img.shields.io/github/stars/hjson/hjson-go.svg) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.
- [hocon](https://github.com/gurkankaymak/hocon) ![](https://img.shields.io/github/stars/gurkankaymak/hocon.svg) - Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files.
- [ini](https://github.com/go-ini/ini) ![](https://img.shields.io/github/stars/go-ini/ini.svg) - Go package to read and write INI files.
- [ini](https://github.com/wlevene/ini) ![](https://img.shields.io/github/stars/wlevene/ini.svg) - INI Parser & Write Library, Unmarshal to Struct, Marshal to Json, Write File, watch file.
- [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) ![](https://img.shields.io/github/stars/kelseyhightower/envconfig.svg) - Go library for managing configuration data from environment variables.
- [koanf](https://github.com/knadh/koanf) ![](https://img.shields.io/github/stars/knadh/koanf.svg) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.
- [konf](https://github.com/nil-go/konf) ![](https://img.shields.io/github/stars/nil-go/konf.svg) - The simplest API for reading/watching config from file, env, flag and clouds (e.g. AWS, Azure, GCP).
- [konfig](https://github.com/lalamove/konfig) ![](https://img.shields.io/github/stars/lalamove/konfig.svg) - Composable, observable and performant config handling for Go for the distributed processing era.
- [kong](https://github.com/alecthomas/kong) ![](https://img.shields.io/github/stars/alecthomas/kong.svg) - Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`).
- [nasermirzaei89/env](https://github.com/nasermirzaei89/env) ![](https://img.shields.io/github/stars/nasermirzaei89/env.svg) - Simple useful package for read environment variables.
- [nfigure](https://github.com/muir/nfigure) ![](https://img.shields.io/github/stars/muir/nfigure.svg) - Per-library struct-tag based configuration from command lines (Posix & Go-style); environment, JSON, YAML
- [onion](https://github.com/goraz/onion) ![](https://img.shields.io/github/stars/goraz/onion.svg) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.
- [piper](https://github.com/Yiling-J/piper) ![](https://img.shields.io/github/stars/Yiling-J/piper.svg) - Viper wrapper with config inheritance and key generation.
- [sonic](https://github.com/bytedance/sonic) ![](https://img.shields.io/github/stars/bytedance/sonic.svg) - A blazingly fast JSON serializing & deserializing library.
- [swap](https://github.com/oblq/swap) ![](https://img.shields.io/github/stars/oblq/swap.svg) - Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env).
- [typenv](https://github.com/diegomarangoni/typenv) ![](https://img.shields.io/github/stars/diegomarangoni/typenv.svg) - Minimalistic, zero dependency, typed environment variables library.
- [uConfig](https://github.com/omeid/uconfig) ![](https://img.shields.io/github/stars/omeid/uconfig.svg) - Lightweight, zero-dependency, and extendable configuration management.
- [viper](https://github.com/spf13/viper) ![](https://img.shields.io/github/stars/spf13/viper.svg) - Go configuration with fangs.
- [xdg](https://github.com/adrg/xdg) ![](https://img.shields.io/github/stars/adrg/xdg.svg) - Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).
- [yamagiconf](https://github.com/romshark/yamagiconf) ![](https://img.shields.io/github/stars/romshark/yamagiconf.svg) - The "safe subset" of YAML for Go configs.
- [zerocfg](https://github.com/chaindead/zerocfg) ![](https://img.shields.io/github/stars/chaindead/zerocfg.svg) - Zero-effort, concise configuration management that avoids boilerplate and repetitive code, supports multiple sources with priority overrides.

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [abstruse](https://github.com/bleenco/abstruse) ![](https://img.shields.io/github/stars/bleenco/abstruse.svg) - Abstruse is a distributed CI platform.
- [Bencher](https://bencher.dev/) - A suite of continuous benchmarking tools designed to catch performance regressions in CI.
- [CDS](https://github.com/ovh/cds) ![](https://img.shields.io/github/stars/ovh/cds.svg) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.
- [dot](https://github.com/opnlabs/dot) ![](https://img.shields.io/github/stars/opnlabs/dot.svg) - A minimal, local first continuous integration system that uses Docker to run jobs concurrently in stages.
- [drone](https://github.com/drone/drone) ![](https://img.shields.io/github/stars/drone/drone.svg) - Drone is a Continuous Integration platform built on Docker, written in Go.
- [go-beautiful-html-coverage](https://github.com/gha-common/go-beautiful-html-coverage) ![](https://img.shields.io/github/stars/gha-common/go-beautiful-html-coverage.svg) - A GitHub Action to track code coverage in your pull requests, with a beautiful HTML preview, for free.
- [go-fuzz-action](https://github.com/jidicula/go-fuzz-action) ![](https://img.shields.io/github/stars/jidicula/go-fuzz-action.svg) - Use Go 1.18's built-in fuzz testing in GitHub Actions.
- [go-semver-release](https://github.com/s0ders/go-semver-release) ![](https://img.shields.io/github/stars/s0ders/go-semver-release.svg) - Automate the semantic versioning of Git repositories.
- [go-test-coverage](https://github.com/marketplace/actions/go-test-coverage) ![](https://img.shields.io/github/stars/marketplace/actions.svg) - A GitHub Action which reports issues when test coverage is below set threshold.
- [gomason](https://github.com/nikogura/gomason) ![](https://img.shields.io/github/stars/nikogura/gomason.svg) - Test, Build, Sign, and Publish your go binaries from a clean workspace.
- [gotestfmt](https://github.com/GoTestTools/gotestfmt) ![](https://img.shields.io/github/stars/GoTestTools/gotestfmt.svg) - go test output for humans.
- [goveralls](https://github.com/mattn/goveralls) ![](https://img.shields.io/github/stars/mattn/goveralls.svg) - Go integration for Coveralls.io continuous code coverage tracking system.
- [muffet](https://github.com/raviqqe/muffet) ![](https://img.shields.io/github/stars/raviqqe/muffet.svg) - Fast website link checker in Go, see [alternatives](https://github.com/lycheeverse/lychee#features).
- [overalls](https://github.com/go-playground/overalls) ![](https://img.shields.io/github/stars/go-playground/overalls.svg) - Multi-Package go project coverprofile for tools like goveralls.
- [roveralls](https://github.com/LawrenceWoodman/roveralls) ![](https://img.shields.io/github/stars/LawrenceWoodman/roveralls.svg) - Recursive coverage testing tool.
- [woodpecker](https://github.com/woodpecker-ci/woodpecker) ![](https://img.shields.io/github/stars/woodpecker-ci/woodpecker.svg) - Woodpecker is a community fork of the Drone CI system.

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._

- [go-css](https://github.com/napsy/go-css) ![](https://img.shields.io/github/stars/napsy/go-css.svg) - A very simple CSS parser, written in Go.
- [go-libsass](https://github.com/wellington/go-libsass) ![](https://img.shields.io/github/stars/wellington/go-libsass.svg) - Go wrapper to the 100% Sass compatible libsass project.

**[⬆ back to top](#contents)**

## Data Integration Frameworks

_Frameworks for performing ELT / ETL_

- [Benthos](https://github.com/benthosdev/benthos) ![](https://img.shields.io/github/stars/benthosdev/benthos.svg) - A message streaming bridge between a range of protocols.
- [CloudQuery](http://github.com/cloudquery/cloudquery) - A high-performance ELT data integration framework with pluggable architecture.
- [omniparser](https://github.com/jf-tech/omniparser) ![](https://img.shields.io/github/stars/jf-tech/omniparser.svg) - A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema.

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [bingo](https://github.com/iancmcc/bingo) ![](https://img.shields.io/github/stars/iancmcc/bingo.svg) - Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes.
- [binpacker](https://github.com/zhuangsirui/binpacker) ![](https://img.shields.io/github/stars/zhuangsirui/binpacker.svg) - Binary packer and unpacker helps user build custom binary stream.
- [bit](https://github.com/yourbasic/bit) ![](https://img.shields.io/github/stars/yourbasic/bit.svg) - Golang set data structure with bonus bit-twiddling functions.
- [crunch](https://github.com/superwhiskers/crunch) ![](https://img.shields.io/github/stars/superwhiskers/crunch.svg) - Go package implementing buffers for handling various datatypes easily.
- [go-ef](https://github.com/amallia/go-ef) ![](https://img.shields.io/github/stars/amallia/go-ef.svg) - A Go implementation of the Elias-Fano encoding.
- [roaring](https://github.com/RoaringBitmap/roaring) ![](https://img.shields.io/github/stars/RoaringBitmap/roaring.svg) - Go package implementing compressed bitsets.

### Bit Sets

- [bitmap](https://github.com/kelindar/bitmap) ![](https://img.shields.io/github/stars/kelindar/bitmap.svg) - Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go.
- [bitset](https://github.com/bits-and-blooms/bitset) ![](https://img.shields.io/github/stars/bits-and-blooms/bitset.svg) - Go package implementing bitsets.

### Bloom and Cuckoo Filters

- [bloom](https://github.com/bits-and-blooms/bloom) ![](https://img.shields.io/github/stars/bits-and-blooms/bloom.svg) - Go package implementing Bloom filters.
- [bloom](https://github.com/zhenjl/bloom) ![](https://img.shields.io/github/stars/zhenjl/bloom.svg) - Bloom filters implemented in Go.
- [bloom](https://github.com/yourbasic/bloom) ![](https://img.shields.io/github/stars/yourbasic/bloom.svg) - Golang Bloom filter implementation.
- [bloomfilter](https://github.com/OldPanda/bloomfilter) ![](https://img.shields.io/github/stars/OldPanda/bloomfilter.svg) - Yet another Bloomfilter implementation in Go, compatible with Java's Guava library.
- [boomfilters](https://github.com/tylertreat/BoomFilters) ![](https://img.shields.io/github/stars/tylertreat/BoomFilters.svg) - Probabilistic data structures for processing continuous, unbounded streams.
- [cuckoo-filter](https://github.com/linvon/cuckoo-filter) ![](https://img.shields.io/github/stars/linvon/cuckoo-filter.svg) - Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper are available.
- [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) ![](https://img.shields.io/github/stars/seiflotfy/cuckoofilter.svg) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
- [ring](https://github.com/TheTannerRyan/ring) ![](https://img.shields.io/github/stars/TheTannerRyan/ring.svg) - Go implementation of a high performance, thread safe bloom filter.

### Data Structure and Algorithm Collections

- [algorithms](https://github.com/shady831213/algorithms) ![](https://img.shields.io/github/stars/shady831213/algorithms.svg) - Algorithms and data structures.CLRS study.
- [go-datastructures](https://github.com/Workiva/go-datastructures) ![](https://img.shields.io/github/stars/Workiva/go-datastructures.svg) - Collection of useful, performant, and thread-safe data structures.
- [gods](https://github.com/emirpasic/gods) ![](https://img.shields.io/github/stars/emirpasic/gods.svg) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.
- [gostl](https://github.com/liyue201/gostl) ![](https://img.shields.io/github/stars/liyue201/gostl.svg) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL.

### Iterators

- [gloop](https://github.com/alvii147/gloop) ![](https://img.shields.io/github/stars/alvii147/gloop.svg) - Convenient looping using Go's range-over-func feature.
- [goterator](https://github.com/yaa110/goterator) ![](https://img.shields.io/github/stars/yaa110/goterator.svg) - Iterator implementation to provide map and reduce functionalities.
- [iter](https://github.com/disksing/iter) ![](https://img.shields.io/github/stars/disksing/iter.svg) - Go implementation of C++ STL iterators and algorithms.

### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.

- [cmap](https://github.com/lrita/cmap) ![](https://img.shields.io/github/stars/lrita/cmap.svg) - a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards.
- [concurrent-swiss-map](https://github.com/mhmtszr/concurrent-swiss-map) ![](https://img.shields.io/github/stars/mhmtszr/concurrent-swiss-map.svg) - A high-performance, thread-safe generic concurrent hash map implementation with Swiss Map.
- [dict](https://github.com/srfrog/dict) ![](https://img.shields.io/github/stars/srfrog/dict.svg) - Python-like dictionaries (dict) for Go.
- [go-shelve](https://github.com/lucmq/go-shelve) ![](https://img.shields.io/github/stars/lucmq/go-shelve.svg) - A persistent, map-like object for the Go programming language. Supports multiple embedded key-value stores.
- [goradd/maps](https://github.com/goradd/maps) ![](https://img.shields.io/github/stars/goradd/maps.svg) - Go 1.18+ generic map interface for maps; safe maps; ordered maps; ordered, safe maps; etc.
- [hmap](https://github.com/lyonnee/hmap) ![](https://img.shields.io/github/stars/lyonnee/hmap.svg) - HMap is a concurrent and secure, generic support Map implementation designed to provide an easy-to-use API.

### Miscellaneous Data Structures and Algorithms

- [concurrent-writer](https://github.com/free/concurrent-writer) ![](https://img.shields.io/github/stars/free/concurrent-writer.svg) - Highly concurrent drop-in replacement for `bufio.Writer`.
- [count-min-log](https://github.com/seiflotfy/count-min-log) ![](https://img.shields.io/github/stars/seiflotfy/count-min-log.svg) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).
- [fsm](https://github.com/cocoonspace/fsm) ![](https://img.shields.io/github/stars/cocoonspace/fsm.svg) - Finite-State Machine package.
- [genfuncs](https://github.com/nwillc/genfuncs) ![](https://img.shields.io/github/stars/nwillc/genfuncs.svg) - Go 1.18+ generics package inspired by Kotlin's Sequence and Map.
- [go-generics](https://github.com/bobg/go-generics) ![](https://img.shields.io/github/stars/bobg/go-generics.svg) - Generic slice, map, set, iterator, and goroutine utilities.
- [go-geoindex](https://github.com/hailocab/go-geoindex) ![](https://img.shields.io/github/stars/hailocab/go-geoindex.svg) - In-memory geo index.
- [go-rampart](https://github.com/francesconi/go-rampart) ![](https://img.shields.io/github/stars/francesconi/go-rampart.svg) - Determine how intervals relate to each other.
- [go-rquad](https://github.com/aurelien-rainone/go-rquad) ![](https://img.shields.io/github/stars/aurelien-rainone/go-rquad.svg) - Region quadtrees with efficient point location and neighbour finding.
- [go-tuple](https://github.com/barweiss/go-tuple) ![](https://img.shields.io/github/stars/barweiss/go-tuple.svg) - Generic tuple implementation for Go 1.18+.
- [go18ds](https://github.com/daichi-m/go18ds) ![](https://img.shields.io/github/stars/daichi-m/go18ds.svg) - Go Data Structures using Go 1.18 generics.
- [gofal](https://github.com/xxjwxc/gofal) ![](https://img.shields.io/github/stars/xxjwxc/gofal.svg) - fractional api for Go.
- [gogu](https://github.com/esimov/gogu) ![](https://img.shields.io/github/stars/esimov/gogu.svg) - A comprehensive, reusable and efficient concurrent-safe generics utility functions and data structures library.
- [gota](https://github.com/kniren/gota) ![](https://img.shields.io/github/stars/kniren/gota.svg) - Implementation of dataframes, series, and data wrangling methods for Go.
- [hide](https://github.com/emvi/hide) ![](https://img.shields.io/github/stars/emvi/hide.svg) - ID type with marshalling to/from hash to prevent sending IDs to clients.
- [hyperloglog](https://github.com/axiomhq/hyperloglog) ![](https://img.shields.io/github/stars/axiomhq/hyperloglog.svg) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.
- [quadtree](https://github.com/s0rg/quadtree) ![](https://img.shields.io/github/stars/s0rg/quadtree.svg) - Generic, zero-alloc, 100%-test covered quadtree.
- [slices](https://github.com/twharmon/slices) ![](https://img.shields.io/github/stars/twharmon/slices.svg) - Pure, generic functions for slices.

### Nullable Types

- [nan](https://github.com/kak-tus/nan) ![](https://img.shields.io/github/stars/kak-tus/nan.svg) - Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers.
- [null](https://github.com/emvi/null) ![](https://img.shields.io/github/stars/emvi/null.svg) - Nullable Go types that can be marshalled/unmarshalled to/from JSON.
- [typ](https://github.com/gurukami/typ) ![](https://img.shields.io/github/stars/gurukami/typ.svg) - Null Types, Safe primitive type conversion and fetching value from complex structures.

### Queues

- [deque](https://github.com/edwingeng/deque) ![](https://img.shields.io/github/stars/edwingeng/deque.svg) - A highly optimized double-ended queue.
- [deque](https://github.com/gammazero/deque) ![](https://img.shields.io/github/stars/gammazero/deque.svg) - Fast ring-buffer deque (double-ended queue).
- [dqueue](https://github.com/vodolaz095/dqueue) ![](https://img.shields.io/github/stars/vodolaz095/dqueue.svg) - Simple, in memory, zero dependency and battle tested, thread-safe deferred queue.
- [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) ![](https://img.shields.io/github/stars/enriquebris/goconcurrentqueue.svg) - Concurrent FIFO queue.
- [hatchet](https://github.com/hatchet-dev/hatchet) ![](https://img.shields.io/github/stars/hatchet-dev/hatchet.svg) - Distributed, Fault-tolerant task queue.
- [list](https://github.com/koss-null/list) ![](https://img.shields.io/github/stars/koss-null/list.svg) - A generic, thread-safe doubly linked list with full iterator support and an intrusive singly linked list for embedded use; a feature-rich replacement for container/list.
- [memlog](https://github.com/embano1/memlog) ![](https://img.shields.io/github/stars/embano1/memlog.svg) - An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka.
- [queue](https://github.com/adrianbrad/queue) ![](https://img.shields.io/github/stars/adrianbrad/queue.svg) - Multiple thread-safe, generic queue implementations for Go.

### Sets

- [dsu](https://github.com/ihebu/dsu) ![](https://img.shields.io/github/stars/ihebu/dsu.svg) - Disjoint Set data structure implementation in Go.
- [golang-set](https://github.com/deckarep/golang-set) ![](https://img.shields.io/github/stars/deckarep/golang-set.svg) - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
- [goset](https://github.com/zoumo/goset) ![](https://img.shields.io/github/stars/zoumo/goset.svg) - A useful Set collection implementation for Go.
- [set](https://github.com/StudioSol/set) ![](https://img.shields.io/github/stars/StudioSol/set.svg) - Simple set data structure implementation in Go using LinkedHashMap.

### Text Analysis

- [bleve](https://github.com/blevesearch/bleve) ![](https://img.shields.io/github/stars/blevesearch/bleve.svg) - Modern text indexing library for go.
- [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) ![](https://img.shields.io/github/stars/plar/go-adaptive-radix-tree.svg) - Go implementation of Adaptive Radix Tree.
- [go-edlib](https://github.com/hbollon/go-edlib) ![](https://img.shields.io/github/stars/hbollon/go-edlib.svg) - Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode.
- [levenshtein](https://github.com/agext/levenshtein) ![](https://img.shields.io/github/stars/agext/levenshtein.svg) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
- [levenshtein](https://github.com/agnivade/levenshtein) ![](https://img.shields.io/github/stars/agnivade/levenshtein.svg) - Implementation to calculate levenshtein distance in Go.
- [mspm](https://github.com/BlackRabbitt/mspm) ![](https://img.shields.io/github/stars/BlackRabbitt/mspm.svg) - Multi-String Pattern Matching Algorithm for information retrieval.
- [parsefields](https://github.com/MonaxGT/parsefields) ![](https://img.shields.io/github/stars/MonaxGT/parsefields.svg) - Tools for parse JSON-like logs for collecting unique fields and events.
- [ptrie](https://github.com/viant/ptrie) ![](https://img.shields.io/github/stars/viant/ptrie.svg) - An implementation of prefix tree.
- [trie](https://github.com/derekparker/trie) ![](https://img.shields.io/github/stars/derekparker/trie.svg) - Trie implementation in Go.

### Trees

- [graphlib](https://github.com/aio-arch/graphlib) ![](https://img.shields.io/github/stars/aio-arch/graphlib.svg) - Topological sort lib,Sorting and pruning of DAG graphs.
- [hashsplit](http://github.com/bobg/hashsplit) - Split byte streams into chunks, and arrange chunks into trees, with boundaries determined by content, not position.
- [merkle](https://github.com/bobg/merkle) ![](https://img.shields.io/github/stars/bobg/merkle.svg) - Space-efficient computation of Merkle root hashes and inclusion proofs.
- [skiplist](https://github.com/MauriceGit/skiplist) ![](https://img.shields.io/github/stars/MauriceGit/skiplist.svg) - Very fast Go Skiplist implementation.
- [skiplist](https://github.com/gansidui/skiplist) ![](https://img.shields.io/github/stars/gansidui/skiplist.svg) - Skiplist implementation in Go.
- [treemap](https://github.com/igrmk/treemap) ![](https://img.shields.io/github/stars/igrmk/treemap.svg) - Generic key-sorted map using a red-black tree under the hood.

### Pipes

- [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) ![](https://img.shields.io/github/stars/tejzpr/ordered-concurrently.svg) - Go module that processes work concurrently and returns output in a channel in the order of input.
- [parapipe](https://github.com/nazar256/parapipe) ![](https://img.shields.io/github/stars/nazar256/parapipe.svg) - FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results.
- [pipeline](https://github.com/hyfather/pipeline) ![](https://img.shields.io/github/stars/hyfather/pipeline.svg) - An implementation of pipelines with fan-in and fan-out.
- [pipelines](https://github.com/nxdir-s/pipelines) ![](https://img.shields.io/github/stars/nxdir-s/pipelines.svg) - Generic pipeline functions for concurrent processing.

**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [2q](https://github.com/floatdrop/2q) ![](https://img.shields.io/github/stars/floatdrop/2q.svg) - 2Q in-memory cache implementation.
- [bcache](https://github.com/iwanbk/bcache) ![](https://img.shields.io/github/stars/iwanbk/bcache.svg) - Eventually consistent distributed in-memory cache Go library.
- [BigCache](https://github.com/allegro/bigcache) ![](https://img.shields.io/github/stars/allegro/bigcache.svg) - Efficient key/value cache for gigabytes of data.
- [cache2go](https://github.com/muesli/cache2go) ![](https://img.shields.io/github/stars/muesli/cache2go.svg) - In-memory key:value cache which supports automatic invalidation based on timeouts.
- [cachego](https://github.com/faabiosr/cachego) ![](https://img.shields.io/github/stars/faabiosr/cachego.svg) - Golang Cache component for multiple drivers.
- [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) ![](https://img.shields.io/github/stars/oaStuff/clusteredBigCache.svg) - BigCache with clustering support and individual item expiration.
- [coherence-go-client](https://github.com/oracle/coherence-go-client) ![](https://img.shields.io/github/stars/oracle/coherence-go-client.svg) - Full implementation of Oracle Coherence cache API for Go applications using gRPC as network transport.
- [couchcache](https://github.com/codingsince1985/couchcache) ![](https://img.shields.io/github/stars/codingsince1985/couchcache.svg) - RESTful caching micro-service backed by Couchbase server.
- [EchoVault](https://github.com/EchoVault/EchoVault) ![](https://img.shields.io/github/stars/EchoVault/EchoVault.svg) - Embeddable Distributed in-memory data store compatible with Redis clients.
- [fastcache](https://github.com/VictoriaMetrics/fastcache) ![](https://img.shields.io/github/stars/VictoriaMetrics/fastcache.svg) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.
- [GCache](https://github.com/bluele/gcache) ![](https://img.shields.io/github/stars/bluele/gcache.svg) - Cache library with support for expirable Cache, LFU, LRU and ARC.
- [gdcache](https://github.com/ulovecode/gdcache) ![](https://img.shields.io/github/stars/ulovecode/gdcache.svg) - A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache.
- [go-cache](https://github.com/viney-shih/go-cache) ![](https://img.shields.io/github/stars/viney-shih/go-cache.svg) - A flexible multi-layer Go caching library to deal with in-memory and shared cache by adopting Cache-Aside pattern.
- [go-freelru](https://github.com/elastic/go-freelru) ![](https://img.shields.io/github/stars/elastic/go-freelru.svg) A GC-less, fast and generic LRU hashmap library with optional locking, sharding, eviction and expiration.
- [go-gcache](https://github.com/szyhf/go-gcache) ![](https://img.shields.io/github/stars/szyhf/go-gcache.svg) - The generic version of `GCache`, cache support for expirable Cache, LFU, LRU and ARC.
- [go-mcache](https://github.com/OrlovEvgeny/go-mcache) ![](https://img.shields.io/github/stars/OrlovEvgeny/go-mcache.svg) - Fast in-memory key:value store/cache library. Pointer caches.
- [gocache](https://github.com/eko/gocache) ![](https://img.shields.io/github/stars/eko/gocache.svg) - A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.
- [gocache](https://github.com/yuseferi/gocache) ![](https://img.shields.io/github/stars/yuseferi/gocache.svg) - A data race free Go ache library with high performance and auto pruge functionality
- [groupcache](https://github.com/golang/groupcache) ![](https://img.shields.io/github/stars/golang/groupcache.svg) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
- [icache](https://github.com/mdaliyan/icache) ![](https://img.shields.io/github/stars/mdaliyan/icache.svg) - A High Performance, Generic, thread-safe, zero-dependency cache package.
- [imcache](https://github.com/erni27/imcache) ![](https://img.shields.io/github/stars/erni27/imcache.svg) - A generic in-memory cache Go library. It supports expiration, sliding expiration, max entries limit, eviction callbacks and sharding.
- [jetcache-go](https://github.com/mgtv-tech/jetcache-go) ![](https://img.shields.io/github/stars/mgtv-tech/jetcache-go.svg) - Unified Go cache library supporting multi-level caching.
- [nscache](https://github.com/no-src/nscache) ![](https://img.shields.io/github/stars/no-src/nscache.svg) - A Go caching framework that supports multiple data source drivers.
- [otter](https://github.com/maypok86/otter) ![](https://img.shields.io/github/stars/maypok86/otter.svg) - A high performance lockless cache for Go. Many times faster than Ristretto and friends.
- [pocache](https://github.com/naughtygopher/pocache) ![](https://img.shields.io/github/stars/naughtygopher/pocache.svg) - Pocache is a minimal cache package which focuses on a preemptive optimistic caching strategy.
- [ristretto](https://github.com/dgraph-io/ristretto) ![](https://img.shields.io/github/stars/dgraph-io/ristretto.svg) - A high performance memory-bound Go cache.
- [sturdyc](https://github.com/viccon/sturdyc) ![](https://img.shields.io/github/stars/viccon/sturdyc.svg) - A caching library with advanced concurrency features designed to make I/O heavy applications robust and highly performant.
- [theine](https://github.com/Yiling-J/theine-go) ![](https://img.shields.io/github/stars/Yiling-J/theine-go.svg) - High performance, near optimal in-memory cache with proactive TTL expiration and generics.
- [timedmap](https://github.com/zekroTJA/timedmap) ![](https://img.shields.io/github/stars/zekroTJA/timedmap.svg) - Map with expiring key-value pairs.
- [ttlcache](https://github.com/jellydator/ttlcache) ![](https://img.shields.io/github/stars/jellydator/ttlcache.svg) - An in-memory cache with item expiration and generics.
- [ttlcache](https://github.com/cheshir/ttlcache) ![](https://img.shields.io/github/stars/cheshir/ttlcache.svg) - In-memory key value storage with TTL for each record.

### Databases Implemented in Go

- [badger](https://github.com/dgraph-io/badger) ![](https://img.shields.io/github/stars/dgraph-io/badger.svg) - Fast key-value store in Go.
- [bbolt](https://github.com/etcd-io/bbolt) ![](https://img.shields.io/github/stars/etcd-io/bbolt.svg) - An embedded key/value database for Go.
- [Bitcask](https://git.mills.io/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).
- [buntdb](https://github.com/tidwall/buntdb) ![](https://img.shields.io/github/stars/tidwall/buntdb.svg) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.
- [clover](https://github.com/ostafen/clover) ![](https://img.shields.io/github/stars/ostafen/clover.svg) - A lightweight document-oriented NoSQL database written in pure Golang.
- [cockroach](https://github.com/cockroachdb/cockroach) ![](https://img.shields.io/github/stars/cockroachdb/cockroach.svg) - Scalable, Geo-Replicated, Transactional Datastore.
- [Coffer](https://github.com/claygod/coffer) ![](https://img.shields.io/github/stars/claygod/coffer.svg) - Simple ACID key-value database that supports transactions.
- [column](https://github.com/kelindar/column) ![](https://img.shields.io/github/stars/kelindar/column.svg) - High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions.
- [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) ![](https://img.shields.io/github/stars/CovenantSQL/CovenantSQL.svg) - CovenantSQL is a SQL database on blockchain.
- [Databunker](https://github.com/paranoidguy/databunker) ![](https://img.shields.io/github/stars/paranoidguy/databunker.svg) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.
- [dgraph](https://github.com/dgraph-io/dgraph) ![](https://img.shields.io/github/stars/dgraph-io/dgraph.svg) - Scalable, Distributed, Low Latency, High Throughput Graph Database.
- [DiceDB](https://github.com/DiceDB/dice) ![](https://img.shields.io/github/stars/DiceDB/dice.svg) - An open-source, fast, reactive, in-memory database optimized for modern hardware. Higher throughput and lower median latencies, making it ideal for modern workloads.
- [diskv](https://github.com/peterbourgon/diskv) ![](https://img.shields.io/github/stars/peterbourgon/diskv.svg) - Home-grown disk-backed key-value store.
- [dolt](https://github.com/dolthub/dolt) ![](https://img.shields.io/github/stars/dolthub/dolt.svg) - Dolt – It's Git for Data.
- [eliasdb](https://github.com/krotik/eliasdb) ![](https://img.shields.io/github/stars/krotik/eliasdb.svg) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.
- [godis](https://github.com/hdt3213/godis) ![](https://img.shields.io/github/stars/hdt3213/godis.svg) - A Golang implemented high-performance Redis server and cluster.
- [goleveldb](https://github.com/syndtr/goleveldb) ![](https://img.shields.io/github/stars/syndtr/goleveldb.svg) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.
- [hare](https://github.com/jameycribbs/hare) ![](https://img.shields.io/github/stars/jameycribbs/hare.svg) - A simple database management system that stores each table as a text file of line-delimited JSON.
- [immudb](https://github.com/codenotary/immudb) ![](https://img.shields.io/github/stars/codenotary/immudb.svg) - immudb is a lightweight, high-speed immutable database for systems and applications written in Go.
- [influxdb](https://github.com/influxdb/influxdb) ![](https://img.shields.io/github/stars/influxdb/influxdb.svg) - Scalable datastore for metrics, events, and real-time analytics.
- [ledisdb](https://github.com/siddontang/ledisdb) ![](https://img.shields.io/github/stars/siddontang/ledisdb.svg) - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
- [levigo](https://github.com/jmhodges/levigo) ![](https://img.shields.io/github/stars/jmhodges/levigo.svg) - Levigo is a Go wrapper for LevelDB.
- [libradb](https://github.com/amit-davidson/LibraDB) ![](https://img.shields.io/github/stars/amit-davidson/LibraDB.svg) - LibraDB is a simple database with less than 1000 lines of code for learning.
- [LinDB](https://github.com/lindb/lindb) ![](https://img.shields.io/github/stars/lindb/lindb.svg) - LinDB is a scalable, high performance, high availability distributed time series database.
- [lotusdb](https://github.com/flower-corp/lotusdb) ![](https://img.shields.io/github/stars/flower-corp/lotusdb.svg) - Fast k/v database compatible with lsm and b+tree.
- [Milvus](https://github.com/milvus-io/milvus) ![](https://img.shields.io/github/stars/milvus-io/milvus.svg) - Milvus is a vector database for embedding management, analytics and search.
- [moss](https://github.com/couchbase/moss) ![](https://img.shields.io/github/stars/couchbase/moss.svg) - Moss is a simple LSM key-value storage engine written in 100% Go.
- [NoKV](https://github.com/feichai0017/NoKV) ![](https://img.shields.io/github/stars/feichai0017/NoKV.svg) - High-performance distributed KV storage based on LSM Tree.
- [nutsdb](https://github.com/xujiajun/nutsdb) ![](https://img.shields.io/github/stars/xujiajun/nutsdb.svg) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set.
- [objectbox-go](https://github.com/objectbox/objectbox-go) ![](https://img.shields.io/github/stars/objectbox/objectbox-go.svg) - High-performance embedded Object Database (NoSQL) with Go API.
- [pebble](https://github.com/cockroachdb/pebble) ![](https://img.shields.io/github/stars/cockroachdb/pebble.svg) - RocksDB/LevelDB inspired key-value database in Go.
- [piladb](https://github.com/fern4lvarez/piladb) ![](https://img.shields.io/github/stars/fern4lvarez/piladb.svg) - Lightweight RESTful database engine based on stack data structures.
- [pogreb](https://github.com/akrylysov/pogreb) ![](https://img.shields.io/github/stars/akrylysov/pogreb.svg) - Embedded key-value store for read-heavy workloads.
- [prometheus](https://github.com/prometheus/prometheus) ![](https://img.shields.io/github/stars/prometheus/prometheus.svg) - Monitoring system and time series database.
- [pudge](https://github.com/recoilme/pudge) ![](https://img.shields.io/github/stars/recoilme/pudge.svg) - Fast and simple key/value store written using Go's standard library.
- [redka](https://github.com/nalgeon/redka) ![](https://img.shields.io/github/stars/nalgeon/redka.svg) - Redis re-implemented with SQLite.
- [rosedb](https://github.com/roseduan/rosedb) ![](https://img.shields.io/github/stars/roseduan/rosedb.svg) - An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset.
- [rotom](https://github.com/xgzlucario/rotom) ![](https://img.shields.io/github/stars/xgzlucario/rotom.svg) - A tiny Redis server built with Golang, compatible with RESP protocols.
- [rqlite](https://github.com/rqlite/rqlite) ![](https://img.shields.io/github/stars/rqlite/rqlite.svg) - The lightweight, distributed, relational database built on SQLite.
- [tempdb](https://github.com/rafaeljesus/tempdb) ![](https://img.shields.io/github/stars/rafaeljesus/tempdb.svg) - Key-value store for temporary items.
- [tidb](https://github.com/pingcap/tidb) ![](https://img.shields.io/github/stars/pingcap/tidb.svg) - TiDB is a distributed SQL database. Inspired by the design of Google F1.
- [tiedot](https://github.com/HouzuoGuo/tiedot) ![](https://img.shields.io/github/stars/HouzuoGuo/tiedot.svg) - Your NoSQL database powered by Golang.
- [unitdb](https://github.com/unit-io/unitdb) ![](https://img.shields.io/github/stars/unit-io/unitdb.svg) - Fast timeseries database for IoT, realtime messaging applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application.
- [Vasto](https://github.com/chrislusf/vasto) ![](https://img.shields.io/github/stars/chrislusf/vasto.svg) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) ![](https://img.shields.io/github/stars/VictoriaMetrics/VictoriaMetrics.svg) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.

### Database Schema Migration

- [atlas](https://github.com/ariga/atlas) ![](https://img.shields.io/github/stars/ariga/atlas.svg) - A Database Toolkit. A CLI designed to help companies better work with their data.
- [avro](https://github.com/khezen/avro) ![](https://img.shields.io/github/stars/khezen/avro.svg) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.
- [bytebase](https://github.com/bytebase/bytebase) ![](https://img.shields.io/github/stars/bytebase/bytebase.svg) - Safe database schema change and version control for DevOps teams.
- [darwin](https://github.com/GuiaBolso/darwin) ![](https://img.shields.io/github/stars/GuiaBolso/darwin.svg) - Database schema evolution library for Go.
- [dbmate](https://github.com/amacneil/dbmate) ![](https://img.shields.io/github/stars/amacneil/dbmate.svg) - A lightweight, framework-agnostic database migration tool.
- [go-fixtures](https://github.com/RichardKnop/go-fixtures) ![](https://img.shields.io/github/stars/RichardKnop/go-fixtures.svg) - Django style fixtures for Golang's excellent built-in database/sql library.
- [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) ![](https://img.shields.io/github/stars/lawzava/go-pg-migrate.svg) - CLI-friendly package for go-pg migrations management.
- [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) ![](https://img.shields.io/github/stars/robinjoseph08/go-pg-migrations.svg) - A Go package to help write migrations with go-pg/pg.
- [goavro](https://github.com/linkedin/goavro) ![](https://img.shields.io/github/stars/linkedin/goavro.svg) - A Go package that encodes and decodes Avro data.
- [godfish](https://github.com/rafaelespinoza/godfish) ![](https://img.shields.io/github/stars/rafaelespinoza/godfish.svg) - Database migration manager, works with native query language. Support for cassandra, mysql, postgres, sqlite3.
- [goose](https://github.com/pressly/goose) ![](https://img.shields.io/github/stars/pressly/goose.svg) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
- [gorm-seeder](https://github.com/Kachit/gorm-seeder) ![](https://img.shields.io/github/stars/Kachit/gorm-seeder.svg) - Simple database seeder for Gorm ORM.
- [gormigrate](https://github.com/go-gormigrate/gormigrate) ![](https://img.shields.io/github/stars/go-gormigrate/gormigrate.svg) - Database schema migration helper for Gorm ORM.
- [libschema](https://github.com/muir/libschema) ![](https://img.shields.io/github/stars/muir/libschema.svg) - Define your migrations separately in each library. Migrations for open source libraries. MySQL & PostgreSQL.
- [migrate](https://github.com/golang-migrate/migrate) ![](https://img.shields.io/github/stars/golang-migrate/migrate.svg) - Database migrations. CLI and Golang library.
- [migrator](https://github.com/lopezator/migrator) ![](https://img.shields.io/github/stars/lopezator/migrator.svg) - Dead simple Go database migration library.
- [migrator](https://github.com/larapulse/migrator) ![](https://img.shields.io/github/stars/larapulse/migrator.svg) - MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code.
- [schema](https://github.com/adlio/schema) ![](https://img.shields.io/github/stars/adlio/schema.svg) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.
- [skeema](https://github.com/skeema/skeema) ![](https://img.shields.io/github/stars/skeema/skeema.svg) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.
- [soda](https://github.com/gobuffalo/pop/tree/master/soda) ![](https://img.shields.io/github/stars/gobuffalo/pop.svg) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
- [sql-migrate](https://github.com/rubenv/sql-migrate) ![](https://img.shields.io/github/stars/rubenv/sql-migrate.svg) - Database migration tool. Allows embedding migrations into the application using go-bindata.
- [sqlize](https://github.com/sunary/sqlize) ![](https://img.shields.io/github/stars/sunary/sqlize.svg) - Database migration generator. Allows generate sql migration from model and existing sql by differ them.

### Database Tools

- [chproxy](https://github.com/Vertamedia/chproxy) ![](https://img.shields.io/github/stars/Vertamedia/chproxy.svg) - HTTP proxy for ClickHouse database.
- [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) ![](https://img.shields.io/github/stars/nikepan/clickhouse-bulk.svg) - Collects small inserts and sends big requests to ClickHouse servers.
- [database-gateway](https://github.com/kazhuravlev/database-gateway) ![](https://img.shields.io/github/stars/kazhuravlev/database-gateway.svg) - Running SQL in production with ACLs, logs, and shared links.
- [dbbench](https://github.com/sj14/dbbench) ![](https://img.shields.io/github/stars/sj14/dbbench.svg) - Database benchmarking tool with support for several databases and scripts.
- [dg](https://github.com/codingconcepts/dg) ![](https://img.shields.io/github/stars/codingconcepts/dg.svg) - A fast data generator that produces CSV files from generated relational data.
- [gatewayd](https://github.com/gatewayd-io/gatewayd) ![](https://img.shields.io/github/stars/gatewayd-io/gatewayd.svg) - Cloud-native database gateway and framework for building data-driven applications. Like API gateways, for databases.
- [go-mysql](https://github.com/siddontang/go-mysql) ![](https://img.shields.io/github/stars/siddontang/go-mysql.svg) - Go toolset to handle MySQL protocol and replication.
- [gorm-multitenancy](https://github.com/bartventer/gorm-multitenancy) ![](https://img.shields.io/github/stars/bartventer/gorm-multitenancy.svg) - Multi-tenancy support for GORM managed databases.
- [hasql](https://golang.yandex/hasql) - Library for accessing multi-host SQL database installations.
- [octillery](https://github.com/knocknote/octillery) ![](https://img.shields.io/github/stars/knocknote/octillery.svg) - Go package for sharding databases ( Supports every ORM or raw SQL ).
- [onedump](https://github.com/liweiyi88/onedump) ![](https://img.shields.io/github/stars/liweiyi88/onedump.svg) - Database backup from different drivers to different destinations with one command and configuration.
- [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) ![](https://img.shields.io/github/stars/cybertec-postgresql/pg_timetable.svg) - Advanced scheduling for PostgreSQL.
- [pgweb](https://github.com/sosedoff/pgweb) ![](https://img.shields.io/github/stars/sosedoff/pgweb.svg) - Web-based PostgreSQL database browser.
- [prep](https://github.com/hexdigest/prep) ![](https://img.shields.io/github/stars/hexdigest/prep.svg) - Use prepared SQL statements without changing your code.
- [pREST](https://github.com/prest/prest) ![](https://img.shields.io/github/stars/prest/prest.svg) - Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new.
- [rdb](https://github.com/HDT3213/rdb) ![](https://img.shields.io/github/stars/HDT3213/rdb.svg) - Redis RDB file parser for secondary development and memory analysis.
- [rwdb](https://github.com/andizzle/rwdb) ![](https://img.shields.io/github/stars/andizzle/rwdb.svg) - rwdb provides read replica capability for multiple database servers setup.
- [vitess](https://github.com/youtube/vitess) ![](https://img.shields.io/github/stars/youtube/vitess.svg) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.
- [wescale](https://github.com/wesql/wescale) ![](https://img.shields.io/github/stars/wesql/wescale.svg) - WeScale is a database proxy designed to enhance the scalability, performance, security, and resilience of your applications.

### SQL Query Builders

_Libraries for building and using SQL._

- [bqb](https://github.com/nullism/bqb) ![](https://img.shields.io/github/stars/nullism/bqb.svg) - Lightweight and easy to learn query builder.
- [buildsqlx](https://github.com/arthurkushman/buildsqlx) ![](https://img.shields.io/github/stars/arthurkushman/buildsqlx.svg) - Go database query builder library for PostgreSQL.
- [builq](https://github.com/cristalhq/builq) ![](https://img.shields.io/github/stars/cristalhq/builq.svg) - Easily build SQL queries in Go.
- [dbq](https://github.com/rocketlaunchr/dbq) ![](https://img.shields.io/github/stars/rocketlaunchr/dbq.svg) - Zero boilerplate database operations for Go.
- [Dotsql](https://github.com/gchaincl/dotsql) ![](https://img.shields.io/github/stars/gchaincl/dotsql.svg) - Go library that helps you keep sql files in one place and use them with ease.
- [gendry](https://github.com/didi/gendry) ![](https://img.shields.io/github/stars/didi/gendry.svg) - Non-invasive SQL builder and powerful data binder.
- [godbal](https://github.com/xujiajun/godbal) ![](https://img.shields.io/github/stars/xujiajun/godbal.svg) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.
- [goqu](https://github.com/doug-martin/goqu) ![](https://img.shields.io/github/stars/doug-martin/goqu.svg) - Idiomatic SQL builder and query library.
- [gosql](https://github.com/twharmon/gosql) ![](https://img.shields.io/github/stars/twharmon/gosql.svg) - SQL Query builder with better null values support.
- [Hotcoal](https://github.com/motrboat/hotcoal) ![](https://img.shields.io/github/stars/motrboat/hotcoal.svg) - Secure your handcrafted SQL against injection.
- [igor](https://github.com/galeone/igor) ![](https://img.shields.io/github/stars/galeone/igor.svg) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.
- [jet](https://github.com/go-jet/jet) ![](https://img.shields.io/github/stars/go-jet/jet.svg) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.
- [obreron](https://github.com/profe-ajedrez/obreron) ![](https://img.shields.io/github/stars/profe-ajedrez/obreron.svg) - Fast and cheap SQL builder which does only one thing, SQL building.
- [ormlite](https://github.com/pupizoid/ormlite) ![](https://img.shields.io/github/stars/pupizoid/ormlite.svg) - Lightweight package containing some ORM-like features and helpers for sqlite databases.
- [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) ![](https://img.shields.io/github/stars/go-ozzo/ozzo-dbx.svg) - Powerful data retrieval methods as well as DB-agnostic query building capabilities.
- [patcher](https://github.com/Jacobbrewer1/patcher) ![](https://img.shields.io/github/stars/Jacobbrewer1/patcher.svg) - Powerful SQL Query builder that automatically generates SQL queries from structs.
- [qry](https://github.com/HnH/qry) ![](https://img.shields.io/github/stars/HnH/qry.svg) - Tool that generates constants from files with raw SQL queries.
- [sg](https://github.com/go-the-way/sg) ![](https://img.shields.io/github/stars/go-the-way/sg.svg) - A SQL Gen for generating standard SQLs(supports: CRUD) written in Go.
- [sq](https://github.com/bokwoon95/go-structured-query) ![](https://img.shields.io/github/stars/bokwoon95/go-structured-query.svg) - Type-safe SQL builder and struct mapper for Go.
- [sqlc](https://github.com/kyleconroy/sqlc) ![](https://img.shields.io/github/stars/kyleconroy/sqlc.svg) - Generate type-safe code from SQL.
- [sqlf](https://github.com/leporo/sqlf) ![](https://img.shields.io/github/stars/leporo/sqlf.svg) - Fast SQL query builder.
- [sqlingo](https://github.com/lqs/sqlingo) ![](https://img.shields.io/github/stars/lqs/sqlingo.svg) - A lightweight DSL to build SQL in Go.
- [sqrl](https://github.com/elgris/sqrl) ![](https://img.shields.io/github/stars/elgris/sqrl.svg) - SQL query builder, fork of Squirrel with improved performance.
- [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.
- [Squirrel](https://github.com/Masterminds/squirrel) ![](https://img.shields.io/github/stars/Masterminds/squirrel.svg) - Go library that helps you build SQL queries.
- [xo](https://github.com/knq/xo) ![](https://img.shields.io/github/stars/knq/xo.svg) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley](https://github.com/google/cayley) ![](https://img.shields.io/github/stars/google/cayley.svg) - Graph database with support for multiple backends.
- [dsc](https://github.com/viant/dsc) ![](https://img.shields.io/github/stars/viant/dsc.svg) - Datastore connectivity for SQL, NoSQL, structured files.
- [dynamo](https://github.com/fogfish/dynamo) ![](https://img.shields.io/github/stars/fogfish/dynamo.svg) - A simple key-value abstraction to store algebraic and linked-data data types at AWS storage services: AWS DynamoDB and AWS S3.
- [go-transaction-manager](https://github.com/avito-tech/go-transaction-manager) ![](https://img.shields.io/github/stars/avito-tech/go-transaction-manager.svg) - Transaction manager with multiple adapters (sql, sqlx, gorm, mongo, ...) controls transaction boundaries.
- [gokv](https://github.com/philippgille/gokv) ![](https://img.shields.io/github/stars/philippgille/gokv.svg) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).

### Relational Database Drivers

- [avatica](https://github.com/apache/calcite-avatica-go) ![](https://img.shields.io/github/stars/apache/calcite-avatica-go.svg) - Apache Avatica/Phoenix SQL driver for database/sql.
- [bgc](https://github.com/viant/bgc) ![](https://img.shields.io/github/stars/viant/bgc.svg) - Datastore Connectivity for BigQuery for go.
- [firebirdsql](https://github.com/nakagami/firebirdsql) ![](https://img.shields.io/github/stars/nakagami/firebirdsql.svg) - Firebird RDBMS SQL driver for Go.
- [go-adodb](https://github.com/mattn/go-adodb) ![](https://img.shields.io/github/stars/mattn/go-adodb.svg) - Microsoft ActiveX Object DataBase driver for go that uses database/sql.
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) ![](https://img.shields.io/github/stars/denisenkom/go-mssqldb.svg) - Microsoft MSSQL driver for Go.
- [go-oci8](https://github.com/mattn/go-oci8) ![](https://img.shields.io/github/stars/mattn/go-oci8.svg) - Oracle driver for go that uses database/sql.
- [go-rqlite](https://github.com/rqlite/gorqlite) ![](https://img.shields.io/github/stars/rqlite/gorqlite.svg) - A Go client for rqlite, providing easy-to-use abstractions for working with the rqlite API.
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) ![](https://img.shields.io/github/stars/go-sql-driver/mysql.svg) - MySQL driver for Go.
- [go-sqlite3](https://github.com/mattn/go-sqlite3) ![](https://img.shields.io/github/stars/mattn/go-sqlite3.svg) - SQLite3 driver for go that uses database/sql.
- [go-sqlite3](https://github.com/ncruces/go-sqlite3) ![](https://img.shields.io/github/stars/ncruces/go-sqlite3.svg) - This Go module is compatible with the database/sql driver. It allows embedding SQLite into your application, provides direct access to its C API, supports SQLite VFS, and also includes a GORM driver.
- [godror](https://github.com/godror/godror) ![](https://img.shields.io/github/stars/godror/godror.svg) - Oracle driver for Go, using the ODPI-C driver.
- [gofreetds](https://github.com/minus5/gofreetds) ![](https://img.shields.io/github/stars/minus5/gofreetds.svg) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org).
- [KSQL](https://github.com/VinGarcia/ksql) ![](https://img.shields.io/github/stars/VinGarcia/ksql.svg) - A Simple and Powerful Golang SQL Library.
- [pgx](https://github.com/jackc/pgx) ![](https://img.shields.io/github/stars/jackc/pgx.svg) - PostgreSQL driver supporting features beyond those exposed by database/sql.
- [pig](https://github.com/alexeyco/pig) ![](https://img.shields.io/github/stars/alexeyco/pig.svg) - Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily.
- [pq](https://github.com/lib/pq) ![](https://img.shields.io/github/stars/lib/pq.svg) - Pure Go Postgres driver for database/sql.
- [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) ![](https://img.shields.io/github/stars/cvilsmeier/sqinn-go.svg) - SQLite with pure Go.
- [sqlhooks](https://github.com/qustavo/sqlhooks) ![](https://img.shields.io/github/stars/qustavo/sqlhooks.svg) - Attach hooks to any database/sql driver.
- [sqlite](https://pkg.go.dev/modernc.org/sqlite) - Package sqlite is a sql/database driver using a CGo-free port of the C SQLite3 library.
- [surrealdb.go](https://github.com/surrealdb/surrealdb.go) ![](https://img.shields.io/github/stars/surrealdb/surrealdb.go.svg) - SurrealDB Driver for Go.
- [ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk) ![](https://img.shields.io/github/stars/ydb-platform/ydb-go-sdk.svg) - native and database/sql driver YDB (Yandex Database).

### NoSQL Database Drivers

- [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) ![](https://img.shields.io/github/stars/aerospike/aerospike-client-go.svg) - Aerospike client in Go language.
- [arangolite](https://github.com/solher/arangolite) ![](https://img.shields.io/github/stars/solher/arangolite.svg) - Lightweight golang driver for ArangoDB.
- [asc](https://github.com/viant/asc) ![](https://img.shields.io/github/stars/viant/asc.svg) - Datastore Connectivity for Aerospike for go.
- [forestdb](https://github.com/couchbase/goforestdb) ![](https://img.shields.io/github/stars/couchbase/goforestdb.svg) - Go bindings for ForestDB.
- [go-couchbase](https://github.com/couchbase/go-couchbase) ![](https://img.shields.io/github/stars/couchbase/go-couchbase.svg) - Couchbase client in Go.
- [go-mongox](https://github.com/chenmingyong0423/go-mongox) ![](https://img.shields.io/github/stars/chenmingyong0423/go-mongox.svg) - A Go Mongo library based on the official driver, featuring streamlined document operations, generic binding of structs to collections, built-in CRUD, aggregation, automated field updates, struct validation, hooks, and plugin-based programming.
- [go-pilosa](https://github.com/pilosa/go-pilosa) ![](https://img.shields.io/github/stars/pilosa/go-pilosa.svg) - Go client library for Pilosa.
- [go-rejson](https://github.com/nitishm/go-rejson) ![](https://img.shields.io/github/stars/nitishm/go-rejson.svg) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.
- [gocb](https://github.com/couchbase/gocb) ![](https://img.shields.io/github/stars/couchbase/gocb.svg) - Official Couchbase Go SDK.
- [gocosmos](https://github.com/btnguyen2k/gocosmos) ![](https://img.shields.io/github/stars/btnguyen2k/gocosmos.svg) - REST client and standard `database/sql` driver for Azure Cosmos DB.
- [gocql](https://gocql.github.io) - Go language driver for Apache Cassandra.
- [godis](https://github.com/piaohao/godis) ![](https://img.shields.io/github/stars/piaohao/godis.svg) - redis client implement by golang, inspired by jedis.
- [godscache](https://github.com/defcronyke/godscache) ![](https://img.shields.io/github/stars/defcronyke/godscache.svg) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
- [gomemcache](https://github.com/bradfitz/gomemcache/) ![](https://img.shields.io/github/stars/bradfitz/gomemcache.svg) - memcache client library for the Go programming language.
- [gomemcached](https://github.com/aliexpressru/gomemcached) ![](https://img.shields.io/github/stars/aliexpressru/gomemcached.svg) - A binary Memcached client for Go with support for sharding using consistent hashing, along with SASL.
- [gorethink](https://github.com/dancannon/gorethink) ![](https://img.shields.io/github/stars/dancannon/gorethink.svg) - Go language driver for RethinkDB.
- [goriak](https://github.com/zegl/goriak) ![](https://img.shields.io/github/stars/zegl/goriak.svg) - Go language driver for Riak KV.
- [Kivik](https://github.com/go-kivik/kivik) ![](https://img.shields.io/github/stars/go-kivik/kivik.svg) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.
- [mgm](https://github.com/kamva/mgm) ![](https://img.shields.io/github/stars/kamva/mgm.svg) - MongoDB model-based ODM for Go (based on official MongoDB driver).
- [mgo](https://github.com/globalsign/mgo) ![](https://img.shields.io/github/stars/globalsign/mgo.svg) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) ![](https://img.shields.io/github/stars/mongodb/mongo-go-driver.svg) - Official MongoDB driver for the Go language.
- [neo4j](https://github.com/cihangir/neo4j) ![](https://img.shields.io/github/stars/cihangir/neo4j.svg) - Neo4j Rest API Bindings for Golang.
- [neoism](https://github.com/jmcvetta/neoism) ![](https://img.shields.io/github/stars/jmcvetta/neoism.svg) - Neo4j client for Golang.
- [qmgo](https://github.com/qiniu/qmgo) ![](https://img.shields.io/github/stars/qiniu/qmgo.svg) - The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo.
- [redeo](https://github.com/bsm/redeo) ![](https://img.shields.io/github/stars/bsm/redeo.svg) - Redis-protocol compatible TCP servers/services.
- [redigo](https://github.com/gomodule/redigo) ![](https://img.shields.io/github/stars/gomodule/redigo.svg) - Redigo is a Go client for the Redis database.
- [redis](https://github.com/redis/go-redis) ![](https://img.shields.io/github/stars/redis/go-redis.svg) - Redis client for Golang.
- [rueidis](http://github.com/rueian/rueidis) - Fast Redis RESP3 client with auto pipelining and server-assisted client side caching.
- [xredis](https://github.com/shomali11/xredis) ![](https://img.shields.io/github/stars/shomali11/xredis.svg) - Typesafe, customizable, clean & easy to use Redis client.

### Search and Analytic Databases

- [clickhouse-go](https://github.com/ClickHouse/clickhouse-go/) ![](https://img.shields.io/github/stars/ClickHouse/clickhouse-go.svg) - ClickHouse SQL client for Go with a `database/sql` compatibility.
- [effdsl](https://github.com/sdqri/effdsl) ![](https://img.shields.io/github/stars/sdqri/effdsl.svg) - Elasticsearch query builder for Go.
- [elastic](https://github.com/olivere/elastic) ![](https://img.shields.io/github/stars/olivere/elastic.svg) - Elasticsearch client for Go.
- [elasticsql](https://github.com/cch123/elasticsql) ![](https://img.shields.io/github/stars/cch123/elasticsql.svg) - Convert sql to elasticsearch dsl in Go.
- [elastigo](https://github.com/mattbaird/elastigo) ![](https://img.shields.io/github/stars/mattbaird/elastigo.svg) - Elasticsearch client library.
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch) ![](https://img.shields.io/github/stars/elastic/go-elasticsearch.svg) - Official Elasticsearch client for Go.
- [goes](https://github.com/OwnLocal/goes) ![](https://img.shields.io/github/stars/OwnLocal/goes.svg) - Library to interact with Elasticsearch.
- [skizze](https://github.com/seiflotfy/skizze) ![](https://img.shields.io/github/stars/seiflotfy/skizze.svg) - probabilistic data-structures service and storage.
- [zoekt](https://github.com/sourcegraph/zoekt) ![](https://img.shields.io/github/stars/sourcegraph/zoekt.svg) - Fast trigram based code search.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [approx](https://github.com/goschtalt/approx) ![](https://img.shields.io/github/stars/goschtalt/approx.svg) - A Duration extension supporting parsing/printing durations in days, weeks and years.
- [carbon](https://github.com/dromara/carbon) ![](https://img.shields.io/github/stars/dromara/carbon.svg) - A simple, semantic and developer-friendly time package for golang.
- [carbon](https://github.com/uniplaces/carbon) ![](https://img.shields.io/github/stars/uniplaces/carbon.svg) - Simple Time extension with a lot of util methods, ported from PHP Carbon library.
- [cronrange](https://github.com/1set/cronrange) ![](https://img.shields.io/github/stars/1set/cronrange.svg) - Parses Cron-style time range expressions, checks if the given time is within any ranges.
- [date](https://github.com/rickb777/date) ![](https://img.shields.io/github/stars/rickb777/date.svg) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
- [dateparse](https://github.com/araddon/dateparse) ![](https://img.shields.io/github/stars/araddon/dateparse.svg) - Parse date's without knowing format in advance.
- [durafmt](https://github.com/hako/durafmt) ![](https://img.shields.io/github/stars/hako/durafmt.svg) - Time duration formatting library for Go.
- [feiertage](https://github.com/wlbr/feiertage) ![](https://img.shields.io/github/stars/wlbr/feiertage.svg) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
- [go-anytime](https://github.com/ijt/go-anytime) ![](https://img.shields.io/github/stars/ijt/go-anytime.svg) - Parse dates/times like "next dec 22nd at 3pm" and ranges like "from today until next thursday" without knowing the format in advance.
- [go-datebin](https://github.com/deatil/go-datebin) ![](https://img.shields.io/github/stars/deatil/go-datebin.svg) - A simple datetime parse pkg.
- [go-faketime](https://github.com/harkaitz/go-faketime) ![](https://img.shields.io/github/stars/harkaitz/go-faketime.svg) - A simple `time.Now()` that honors the faketime(1) utility.
- [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) ![](https://img.shields.io/github/stars/yaa110/go-persian-calendar.svg) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
- [go-str2duration](https://github.com/xhit/go-str2duration) ![](https://img.shields.io/github/stars/xhit/go-str2duration.svg) - Convert string to duration. Support time.Duration returned string and more.
- [go-sunrise](https://github.com/nathan-osman/go-sunrise) ![](https://img.shields.io/github/stars/nathan-osman/go-sunrise.svg) - Calculate the sunrise and sunset times for a given location.
- [go-week](https://github.com/stoewer/go-week) ![](https://img.shields.io/github/stars/stoewer/go-week.svg) - An efficient package to work with ISO8601 week dates.
- [gostradamus](https://github.com/bykof/gostradamus) ![](https://img.shields.io/github/stars/bykof/gostradamus.svg) - A Go package for working with dates.
- [iso8601](https://github.com/relvacode/iso8601) ![](https://img.shields.io/github/stars/relvacode/iso8601.svg) - Efficiently parse ISO8601 date-times without regex.
- [kair](https://github.com/GuilhermeCaruso/kair) ![](https://img.shields.io/github/stars/GuilhermeCaruso/kair.svg) - Date and Time - Golang Formatting Library.
- [now](https://github.com/jinzhu/now) ![](https://img.shields.io/github/stars/jinzhu/now.svg) - Now is a time toolkit for golang.
- [strftime](https://github.com/awoodbeck/strftime) ![](https://img.shields.io/github/stars/awoodbeck/strftime.svg) - C99-compatible strftime formatter.
- [timespan](https://github.com/SaidinWoT/timespan) ![](https://img.shields.io/github/stars/SaidinWoT/timespan.svg) - For interacting with intervals of time, defined as a start time and a duration.
- [timeutil](https://github.com/leekchan/timeutil) ![](https://img.shields.io/github/stars/leekchan/timeutil.svg) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.
- [tuesday](https://github.com/osteele/tuesday) ![](https://img.shields.io/github/stars/osteele/tuesday.svg) - Ruby-compatible Strftime function.

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [arpc](https://github.com/lesismal/arpc) ![](https://img.shields.io/github/stars/lesismal/arpc.svg) - More effective network communication, support two-way-calling, notify, broadcast.
- [bedrock](https://github.com/z5labs/bedrock) ![](https://img.shields.io/github/stars/z5labs/bedrock.svg) - Provides a minimal, modular and composable foundation for quickly developing services and more use case specific frameworks in Go.
- [capillaries](https://github.com/capillariesio/capillaries) ![](https://img.shields.io/github/stars/capillariesio/capillaries.svg) - distributed batch data processing framework.
- [celeriac](https://github.com/svcavallar/celeriac.v1) ![](https://img.shields.io/github/stars/svcavallar/celeriac.v1.svg) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
- [committer](https://github.com/vadiminshakov/committer) ![](https://img.shields.io/github/stars/vadiminshakov/committer.svg) - A distributed transactions management system (2PC/3PC implementation).
- [consistent](https://github.com/buraksezer/consistent) ![](https://img.shields.io/github/stars/buraksezer/consistent.svg) - Consistent hashing with bounded loads.
- [consistenthash](https://github.com/mbrostami/consistenthash) ![](https://img.shields.io/github/stars/mbrostami/consistenthash.svg) - Consistent hashing with configurable replicas.
- [dht](https://github.com/anacrolix/dht) ![](https://img.shields.io/github/stars/anacrolix/dht.svg) - BitTorrent Kademlia DHT implementation.
- [digota](https://github.com/digota/digota) ![](https://img.shields.io/github/stars/digota/digota.svg) - grpc ecommerce microservice.
- [dot](https://github.com/dotchain/dot/) ![](https://img.shields.io/github/stars/dotchain/dot.svg) - distributed sync using operational transformation/OT.
- [doublejump](https://github.com/edwingeng/doublejump) ![](https://img.shields.io/github/stars/edwingeng/doublejump.svg) - A revamped Google's jump consistent hash.
- [dragonboat](https://github.com/lni/dragonboat) ![](https://img.shields.io/github/stars/lni/dragonboat.svg) - A feature complete and high performance multi-group Raft library in Go.
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) ![](https://img.shields.io/github/stars/dragonflyoss/Dragonfly2.svg) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures.
- [drmaa](https://github.com/dgruber/drmaa) ![](https://img.shields.io/github/stars/dgruber/drmaa.svg) - Job submission library for cluster schedulers based on the DRMAA standard.
- [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
- [dynatomic](https://github.com/tylfin/dynatomic) ![](https://img.shields.io/github/stars/tylfin/dynatomic.svg) - A library for using DynamoDB as an atomic counter.
- [emitter-io](https://github.com/emitter-io/emitter) ![](https://img.shields.io/github/stars/emitter-io/emitter.svg) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.
- [evans](https://github.com/ktr0731/evans) ![](https://img.shields.io/github/stars/ktr0731/evans.svg) - Evans: more expressive universal gRPC client.
- [failured](https://github.com/andy2046/failured) ![](https://img.shields.io/github/stars/andy2046/failured.svg) - adaptive accrual failure detector for distributed systems.
- [flowgraph](https://github.com/vectaport/flowgraph) ![](https://img.shields.io/github/stars/vectaport/flowgraph.svg) - flow-based programming package.
- [gleam](https://github.com/chrislusf/gleam) ![](https://img.shields.io/github/stars/chrislusf/gleam.svg) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.
- [glow](https://github.com/chrislusf/glow) ![](https://img.shields.io/github/stars/chrislusf/glow.svg) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.
- [gmsec](https://github.com/gmsec/micro) ![](https://img.shields.io/github/stars/gmsec/micro.svg) - A Go distributed systems development framework.
- [go-doudou](https://github.com/unionj-cloud/go-doudou) ![](https://img.shields.io/github/stars/unionj-cloud/go-doudou.svg) - A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity.
- [go-eagle](https://github.com/go-eagle/eagle) ![](https://img.shields.io/github/stars/go-eagle/eagle.svg) - A Go framework for the API or Microservice with handy scaffolding tools.
- [go-jump](https://github.com/dgryski/go-jump) ![](https://img.shields.io/github/stars/dgryski/go-jump.svg) - Port of Google's "Jump" Consistent Hash function.
- [go-kit](https://github.com/go-kit/kit) ![](https://img.shields.io/github/stars/go-kit/kit.svg) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.
- [go-micro](https://github.com/micro/go-micro) ![](https://img.shields.io/github/stars/micro/go-micro.svg) - A distributed systems development framework.
- [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) ![](https://img.shields.io/github/stars/sanketplus/go-mysql-lock.svg) - MySQL based distributed lock.
- [go-pdu](https://github.com/pdupub/go-pdu) ![](https://img.shields.io/github/stars/pdupub/go-pdu.svg) - A decentralized identity-based social network.
- [go-sundheit](https://github.com/AppsFlyer/go-sundheit) ![](https://img.shields.io/github/stars/AppsFlyer/go-sundheit.svg) - A library built to provide support for defining async service health checks for golang services.
- [go-zero](https://github.com/tal-tech/go-zero) ![](https://img.shields.io/github/stars/tal-tech/go-zero.svg) - A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity.
- [gorpc](https://github.com/valyala/gorpc) ![](https://img.shields.io/github/stars/valyala/gorpc.svg) - Simple, fast and scalable RPC library for high load.
- [grpc-go](https://github.com/grpc/grpc-go) ![](https://img.shields.io/github/stars/grpc/grpc-go.svg) - The Go language implementation of gRPC. HTTP/2 based RPC.
- [hprose](https://github.com/hprose/hprose-golang) ![](https://img.shields.io/github/stars/hprose/hprose-golang.svg) - Very newbility RPC Library, support 25+ languages now.
- [jsonrpc](https://github.com/osamingo/jsonrpc) ![](https://img.shields.io/github/stars/osamingo/jsonrpc.svg) - The jsonrpc package helps implement of JSON-RPC 2.0.
- [jsonrpc](https://github.com/ybbus/jsonrpc) ![](https://img.shields.io/github/stars/ybbus/jsonrpc.svg) - JSON-RPC 2.0 HTTP client implementation.
- [K8gb](https://github.com/k8gb-io/k8gb) ![](https://img.shields.io/github/stars/k8gb-io/k8gb.svg) - A cloud native Kubernetes Global Balancer.
- [Kitex](https://github.com/cloudwego/kitex) ![](https://img.shields.io/github/stars/cloudwego/kitex.svg) - A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice.
- [Kratos](https://github.com/go-kratos/kratos) ![](https://img.shields.io/github/stars/go-kratos/kratos.svg) - A modular-designed and easy-to-use microservices framework in Go.
- [liftbridge](https://github.com/liftbridge-io/liftbridge) ![](https://img.shields.io/github/stars/liftbridge-io/liftbridge.svg) - Lightweight, fault-tolerant message streams for NATS.
- [lura](https://github.com/luraproject/lura) ![](https://img.shields.io/github/stars/luraproject/lura.svg) - Ultra performant API Gateway framework with middlewares.
- [micro](https://github.com/micro/micro) ![](https://img.shields.io/github/stars/micro/micro.svg) - A distributed systems runtime for the cloud and beyond.
- [mochi mqtt](https://github.com/mochi-co/mqtt) ![](https://img.shields.io/github/stars/mochi-co/mqtt.svg) - Fully spec compliant, embeddable high-performance MQTT v5/v3 broker for IoT, smarthome, and pubsub.
- [NATS](https://github.com/nats-io/nats-server) ![](https://img.shields.io/github/stars/nats-io/nats-server.svg) - NATS is a simple, secure, and performant communications system for digital systems, services, and devices.
- [opentelemetry-go-auto-instrumentation](https://github.com/alibaba/opentelemetry-go-auto-instrumentation) ![](https://img.shields.io/github/stars/alibaba/opentelemetry-go-auto-instrumentation.svg) - OpenTelemetry Compile-Time Instrumentation for Golang.
- [oras](https://github.com/oras-project/oras) ![](https://img.shields.io/github/stars/oras-project/oras.svg) - CLI and library for OCI Artifacts in container registries.
- [outbox](https://github.com/oagudo/outbox) ![](https://img.shields.io/github/stars/oagudo/outbox.svg) - Lightweight library for the transactional outbox pattern in Go, not tied to any specific relational database or broker.
- [outboxer](https://github.com/italolelis/outboxer) ![](https://img.shields.io/github/stars/italolelis/outboxer.svg) - Outboxer is a go library that implements the outbox pattern.
- [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.
- [pjrpc](https://gitlab.com/pjrpc/pjrpc) - Golang JSON-RPC Server-Client with Protobuf spec.
- [raft](https://github.com/hashicorp/raft) ![](https://img.shields.io/github/stars/hashicorp/raft.svg) - Golang implementation of the Raft consensus protocol, by HashiCorp.
- [raft](https://github.com/etcd-io/raft) ![](https://img.shields.io/github/stars/etcd-io/raft.svg) - Go implementation of the Raft consensus protocol, by CoreOS.
- [rain](https://github.com/cenkalti/rain) ![](https://img.shields.io/github/stars/cenkalti/rain.svg) - BitTorrent client and library.
- [redis-lock](https://github.com/bsm/redislock) ![](https://img.shields.io/github/stars/bsm/redislock.svg) - Simplified distributed locking implementation using Redis.
- [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
- [ringpop-go](https://github.com/uber/ringpop-go) ![](https://img.shields.io/github/stars/uber/ringpop-go.svg) - Scalable, fault-tolerant application-layer sharding for Go applications.
- [rpcx](https://github.com/smallnest/rpcx) ![](https://img.shields.io/github/stars/smallnest/rpcx.svg) - Distributed pluggable RPC service framework like alibaba Dubbo.
- [Semaphore](https://github.com/jexia/semaphore) ![](https://img.shields.io/github/stars/jexia/semaphore.svg) - A straightforward (micro) service orchestrator.
- [sleuth](https://github.com/ursiform/sleuth) ![](https://img.shields.io/github/stars/ursiform/sleuth.svg) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).
- [sponge](https://github.com/zhufuyi/sponge) ![](https://img.shields.io/github/stars/zhufuyi/sponge.svg) - A distributed development framework that integrates automatic code generation, gin and grpc frameworks, base development frameworks.
- [Tarmac](https://github.com/tarmac-project/tarmac) ![](https://img.shields.io/github/stars/tarmac-project/tarmac.svg) - Framework for writing functions, microservices, or monoliths with WebAssembly
- [Temporal](https://github.com/temporalio/sdk-go) ![](https://img.shields.io/github/stars/temporalio/sdk-go.svg) - Durable execution system for making code fault-tolerant and simple.
- [torrent](https://github.com/anacrolix/torrent) ![](https://img.shields.io/github/stars/anacrolix/torrent.svg) - BitTorrent client package.
- [trpc-go](https://github.com/trpc-group/trpc-go) ![](https://img.shields.io/github/stars/trpc-group/trpc-go.svg) - The Go language implementation of tRPC, which is a pluggable, high-performance RPC framework.

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [DDNS](https://github.com/skibish/ddns) ![](https://img.shields.io/github/stars/skibish/ddns.svg) - Personal DDNS client with Digital Ocean Networking DNS as backend.
- [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.
- [GoDNS](https://github.com/timothyye/godns) ![](https://img.shields.io/github/stars/timothyye/godns.svg) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.
- [douceur](https://github.com/aymerick/douceur) ![](https://img.shields.io/github/stars/aymerick/douceur.svg) - CSS inliner for your HTML emails.
- [email](https://github.com/jordan-wright/email) ![](https://img.shields.io/github/stars/jordan-wright/email.svg) - A robust and flexible email library for Go.
- [email-verifier](https://github.com/AfterShip/email-verifier) ![](https://img.shields.io/github/stars/AfterShip/email-verifier.svg) - A Go library for email verification without sending any emails.
- [go-dkim](https://github.com/toorop/go-dkim) ![](https://img.shields.io/github/stars/toorop/go-dkim.svg) - DKIM library, to sign & verify email.
- [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) ![](https://img.shields.io/github/stars/dimuska139/go-email-normalizer.svg) - Golang library for providing a canonical representation of email address.
- [go-imap](https://github.com/emersion/go-imap) ![](https://img.shields.io/github/stars/emersion/go-imap.svg) - IMAP library for clients and servers.
- [go-mail](https://github.com/wneessen/go-mail) ![](https://img.shields.io/github/stars/wneessen/go-mail.svg) - A simple Go library for sending mails in Go.
- [go-message](https://github.com/emersion/go-message) ![](https://img.shields.io/github/stars/emersion/go-message.svg) - Streaming library for the Internet Message Format and mail messages.
- [go-premailer](https://github.com/vanng822/go-premailer) ![](https://img.shields.io/github/stars/vanng822/go-premailer.svg) - Inline styling for HTML mail in Go.
- [go-simple-mail](https://github.com/xhit/go-simple-mail) ![](https://img.shields.io/github/stars/xhit/go-simple-mail.svg) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.
- [Hectane](https://github.com/hectane/hectane) ![](https://img.shields.io/github/stars/hectane/hectane.svg) - Lightweight SMTP client providing an HTTP API.
- [hermes](https://github.com/matcornic/hermes) ![](https://img.shields.io/github/stars/matcornic/hermes.svg) - Golang package that generates clean, responsive HTML e-mails.
- [Maddy](https://github.com/foxcpp/maddy) ![](https://img.shields.io/github/stars/foxcpp/maddy.svg) - All-in-one (SMTP, IMAP, DKIM, DMARC, MTA-STS, DANE) email server
- [mailchain](https://github.com/mailchain/mailchain) ![](https://img.shields.io/github/stars/mailchain/mailchain.svg) - Send encrypted emails to blockchain addresses written in Go.
- [mailgun-go](https://github.com/mailgun/mailgun-go) ![](https://img.shields.io/github/stars/mailgun/mailgun-go.svg) - Go library for sending mail with the Mailgun API.
- [MailHog](https://github.com/mailhog/MailHog) ![](https://img.shields.io/github/stars/mailhog/MailHog.svg) - Email and SMTP testing with web and API interface.
- [Mailpit](https://github.com/axllent/mailpit) ![](https://img.shields.io/github/stars/axllent/mailpit.svg) - Email and SMTP testing tool for developers.
- [mailx](https://github.com/valord577/mailx) ![](https://img.shields.io/github/stars/valord577/mailx.svg) - Mailx is a library that makes it easier to send email via SMTP. It is an enhancement of the golang standard library `net/smtp`.
- [mox](https://github.com/mjl-/mox) ![](https://img.shields.io/github/stars/mjl-/mox.svg) - Modern full-featured secure mail server for low-maintenance, self-hosted email.
- [SendGrid](https://github.com/sendgrid/sendgrid-go) ![](https://img.shields.io/github/stars/sendgrid/sendgrid-go.svg) - SendGrid's Go library for sending email.
- [smtp](https://github.com/mailhog/smtp) ![](https://img.shields.io/github/stars/mailhog/smtp.svg) - SMTP server protocol state machine.
- [smtpmock](https://github.com/mocktools/go-smtp-mock) ![](https://img.shields.io/github/stars/mocktools/go-smtp-mock.svg) - Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment.
- [truemail-go](https://github.com/truemail-rb/truemail-go) ![](https://img.shields.io/github/stars/truemail-rb/truemail-go.svg) - Configurable Golang email validator/verifier. Verify email via Regex, DNS, SMTP and even more.

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [anko](https://github.com/mattn/anko) ![](https://img.shields.io/github/stars/mattn/anko.svg) - Scriptable interpreter written in Go.
- [binder](https://github.com/alexeyco/binder) ![](https://img.shields.io/github/stars/alexeyco/binder.svg) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).
- [cel-go](https://github.com/google/cel-go) ![](https://img.shields.io/github/stars/google/cel-go.svg) - Fast, portable, non-Turing complete expression evaluation with gradual typing.
- [ecal](https://github.com/krotik/ecal) ![](https://img.shields.io/github/stars/krotik/ecal.svg) - A simple embeddable scripting language which supports concurrent event processing.
- [expr](https://github.com/antonmedv/expr) ![](https://img.shields.io/github/stars/antonmedv/expr.svg) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing.
- [FrankenPHP](https://github.com/dunglas/frankenphp) ![](https://img.shields.io/github/stars/dunglas/frankenphp.svg) - PHP embedded in Go, with a `net/http` handler.
- [gentee](https://github.com/gentee/gentee) ![](https://img.shields.io/github/stars/gentee/gentee.svg) - Embeddable scripting programming language.
- [gisp](https://github.com/jcla1/gisp) ![](https://img.shields.io/github/stars/jcla1/gisp.svg) - Simple LISP in Go.
- [go-lua](https://github.com/Shopify/go-lua) ![](https://img.shields.io/github/stars/Shopify/go-lua.svg) - Port of the Lua 5.2 VM to pure Go.
- [go-php](https://github.com/deuill/go-php) ![](https://img.shields.io/github/stars/deuill/go-php.svg) - PHP bindings for Go.
- [goal](https://codeberg.org/anaseto/goal) - An embeddable scripting array language.
- [goja](https://github.com/dop251/goja) ![](https://img.shields.io/github/stars/dop251/goja.svg) - ECMAScript 5.1(+) implementation in Go.
- [golua](https://github.com/aarzilli/golua) ![](https://img.shields.io/github/stars/aarzilli/golua.svg) - Go bindings for Lua C API.
- [gopher-lua](https://github.com/yuin/gopher-lua) ![](https://img.shields.io/github/stars/yuin/gopher-lua.svg) - Lua 5.1 VM and compiler written in Go.
- [gval](https://github.com/PaesslerAG/gval) ![](https://img.shields.io/github/stars/PaesslerAG/gval.svg) - A highly customizable expression language written in Go.
- [metacall](https://github.com/metacall/core) ![](https://img.shields.io/github/stars/metacall/core.svg) - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more.
- [ngaro](https://github.com/db47h/ngaro) ![](https://img.shields.io/github/stars/db47h/ngaro.svg) - Embeddable Ngaro VM implementation enabling scripting in Retro.
- [prolog](https://github.com/ichiban/prolog) ![](https://img.shields.io/github/stars/ichiban/prolog.svg) - Embeddable Prolog.
- [purl](https://github.com/ian-kent/purl) ![](https://img.shields.io/github/stars/ian-kent/purl.svg) - Perl 5.18.2 embedded in Go.
- [starlark-go](https://github.com/google/starlark-go) ![](https://img.shields.io/github/stars/google/starlark-go.svg) - Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution.
- [starlet](https://github.com/1set/starlet) ![](https://img.shields.io/github/stars/1set/starlet.svg) - Go wrapper for [starlark-go](https://github.com/google/starlark-go) that simplifies script execution, offers data conversion, and useful Starlark libraries and extensions.
- [tengo](https://github.com/d5/tengo) ![](https://img.shields.io/github/stars/d5/tengo.svg) - Bytecode compiled script language for Go.
- [Wa/凹语言](https://github.com/wa-lang/wa) ![](https://img.shields.io/github/stars/wa-lang/wa.svg) - The Wa Programming Language embedded in Go.

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [emperror](https://github.com/emperror/emperror) ![](https://img.shields.io/github/stars/emperror/emperror.svg) - Error handling tools and best practices for Go libraries and applications.
- [eris](https://github.com/rotisserie/eris) ![](https://img.shields.io/github/stars/rotisserie/eris.svg) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.
- [errlog](https://github.com/snwfdhmp/errlog) ![](https://img.shields.io/github/stars/snwfdhmp/errlog.svg) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.
- [errors](https://github.com/emperror/errors) ![](https://img.shields.io/github/stars/emperror/errors.svg) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.
- [errors](https://github.com/neuronlabs/errors) ![](https://img.shields.io/github/stars/neuronlabs/errors.svg) - Simple golang error handling with classification primitives.
- [errors](https://github.com/PumpkinSeed/errors) ![](https://img.shields.io/github/stars/PumpkinSeed/errors.svg) - The most simple error wrapper with awesome performance and minimal memory overhead.
- [errors](https://gitlab.com/tozd/go/errors) - Providing errors with a stack trace and optional structured details. Compatible with github.com/pkg/errors API but does not use it internally.
- [errors](https://github.com/naughtygopher/errors) ![](https://img.shields.io/github/stars/naughtygopher/errors.svg) - Drop-in replacement for builtin Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions.
- [errors](https://github.com/cockroachdb/errors) ![](https://img.shields.io/github/stars/cockroachdb/errors.svg) - Go error library with error portability over the network.
- [errorx](https://github.com/joomcode/errorx) ![](https://img.shields.io/github/stars/joomcode/errorx.svg) - A feature rich error package with stack traces, composition of errors and more.
- [exception](https://github.com/rbrahul/exception) ![](https://img.shields.io/github/stars/rbrahul/exception.svg) - A simple utility package for exception handling with try-catch in Golang.
- [Falcon](https://github.com/SonicRoshan/falcon) ![](https://img.shields.io/github/stars/SonicRoshan/falcon.svg) - A Simple Yet Highly Powerful Package For Error Handling.
- [Fault](https://github.com/Southclaws/fault) ![](https://img.shields.io/github/stars/Southclaws/fault.svg) - An ergonomic mechanism for wrapping errors in order to facilitate structured metadata and context for error values.
- [go-multierror](https://github.com/hashicorp/go-multierror) ![](https://img.shields.io/github/stars/hashicorp/go-multierror.svg) - Go (golang) package for representing a list of errors as a single error.
- [metaerr](https://github.com/quantumcycle/metaerr) ![](https://img.shields.io/github/stars/quantumcycle/metaerr.svg) - A library to create your custom error builders producing structured errors with metadata from different sources and optional stacktraces.
- [multierr](https://github.com/uber-go/multierr) ![](https://img.shields.io/github/stars/uber-go/multierr.svg) - Package for representing a list of errors as a single error.
- [oops](https://github.com/samber/oops) ![](https://img.shields.io/github/stars/samber/oops.svg) - Error handling with context, stack trace and source fragments.
- [tracerr](https://github.com/ztrue/tracerr) ![](https://img.shields.io/github/stars/ztrue/tracerr.svg) - Golang errors with stack trace and source fragments.

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [afero](https://github.com/spf13/afero) ![](https://img.shields.io/github/stars/spf13/afero.svg) - FileSystem Abstraction System for Go.
- [afs](https://github.com/viant/afs) ![](https://img.shields.io/github/stars/viant/afs.svg) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.
- [baraka](https://github.com/xis/baraka) ![](https://img.shields.io/github/stars/xis/baraka.svg) - A library to process http file uploads easily.
- [checksum](https://github.com/codingsince1985/checksum) ![](https://img.shields.io/github/stars/codingsince1985/checksum.svg) - Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files.
- [copy](https://github.com/otiai10/copy) ![](https://img.shields.io/github/stars/otiai10/copy.svg) - Copy directory recursively.
- [fastwalk](https://github.com/charlievieth/fastwalk) ![](https://img.shields.io/github/stars/charlievieth/fastwalk.svg) - Fast parallel directory traversal library (used by [fzf](https://github.com/junegunn/fzf)).
- [flop](https://github.com/homedepot/flop) ![](https://img.shields.io/github/stars/homedepot/flop.svg) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).
- [gdu](https://github.com/dundee/gdu) ![](https://img.shields.io/github/stars/dundee/gdu.svg) - Disk usage analyzer with console interface.
- [go-csv-tag](https://github.com/artonge/go-csv-tag) ![](https://img.shields.io/github/stars/artonge/go-csv-tag.svg) - Load csv file using tag.
- [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) ![](https://img.shields.io/github/stars/hugocarreira/go-decent-copy.svg) - Copy files for humans.
- [go-exiftool](https://github.com/barasher/go-exiftool) ![](https://img.shields.io/github/stars/barasher/go-exiftool.svg) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).
- [go-gtfs](https://github.com/artonge/go-gtfs) ![](https://img.shields.io/github/stars/artonge/go-gtfs.svg) - Load gtfs files in go.
- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) ![](https://img.shields.io/github/stars/SebastiaanKlippert/go-wkhtmltopdf.svg) - A package to convert an HTML template to a PDF file.
- [gofs](https://github.com/no-src/gofs) ![](https://img.shields.io/github/stars/no-src/gofs.svg) - A cross-platform real-time file synchronization tool out of the box.
- [gulter](https://github.com/adelowo/gulter) ![](https://img.shields.io/github/stars/adelowo/gulter.svg) - A simple HTTP middleware to automatically handle all your file upload needs
- [gut/yos](https://github.com/1set/gut) ![](https://img.shields.io/github/stars/1set/gut.svg) - Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.
- [higgs](https://github.com/dastoori/higgs) ![](https://img.shields.io/github/stars/dastoori/higgs.svg) - A tiny cross-platform Go library to hide/unhide files and directories.
- [iso9660](https://github.com/kdomanski/iso9660) ![](https://img.shields.io/github/stars/kdomanski/iso9660.svg) - A package for reading and creating ISO9660 disk images
- [notify](https://github.com/rjeczalik/notify) ![](https://img.shields.io/github/stars/rjeczalik/notify.svg) - File system event notification library with simple API, similar to os/signal.
- [opc](https://github.com/qmuntal/opc) ![](https://img.shields.io/github/stars/qmuntal/opc.svg) - Load Open Packaging Conventions (OPC) files for Go.
- [parquet](https://github.com/parsyl/parquet) ![](https://img.shields.io/github/stars/parsyl/parquet.svg) - Read and write [parquet](https://parquet.apache.org) files.
- [pathtype](https://github.com/jonchun/pathtype) ![](https://img.shields.io/github/stars/jonchun/pathtype.svg) - Treat paths as their own type instead of using strings.
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) ![](https://img.shields.io/github/stars/pdfcpu/pdfcpu.svg) - PDF processor.
- [skywalker](https://github.com/dixonwille/skywalker) ![](https://img.shields.io/github/stars/dixonwille/skywalker.svg) - Package to allow one to concurrently go through a filesystem with ease.
- [todotxt](https://github.com/1set/todotxt) ![](https://img.shields.io/github/stars/1set/todotxt.svg) - Go library for Gina Trapani's [_todo.txt_](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [_todo.txt_ format](https://github.com/todotxt/todo.txt).
- [vfs](https://github.com/C2FO/vfs) ![](https://img.shields.io/github/stars/C2FO/vfs.svg) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [accounting](https://github.com/leekchan/accounting) ![](https://img.shields.io/github/stars/leekchan/accounting.svg) - money and currency formatting for golang.
- [ach](https://github.com/moov-io/ach) ![](https://img.shields.io/github/stars/moov-io/ach.svg) - A reader, writer, and validator for Automated Clearing House (ACH) files.
- [bbgo](https://github.com/c9s/bbgo) ![](https://img.shields.io/github/stars/c9s/bbgo.svg) - A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies.
- [currency](https://github.com/bojanz/currency) ![](https://img.shields.io/github/stars/bojanz/currency.svg) - Handles currency amounts, provides currency information and formatting.
- [currency](https://github.com/naughtygopher/currency) ![](https://img.shields.io/github/stars/naughtygopher/currency.svg) - High performant & accurate currency computation package.
- [dec128](https://github.com/jokruger/dec128) ![](https://img.shields.io/github/stars/jokruger/dec128.svg) - High performance 128-bit fixed-point decimal numbers.
- [decimal](https://github.com/shopspring/decimal) ![](https://img.shields.io/github/stars/shopspring/decimal.svg) - Arbitrary-precision fixed-point decimal numbers.
- [decimal](https://github.com/govalues/decimal) ![](https://img.shields.io/github/stars/govalues/decimal.svg) - Immutable decimal numbers with panic-free arithmetic.
- [fpdecimal](https://github.com/nikolaydubina/fpdecimal) ![](https://img.shields.io/github/stars/nikolaydubina/fpdecimal.svg) - Fast and precise serialization and arithmetic for small fixed-point decimals
- [fpmoney](https://github.com/nikolaydubina/fpmoney) ![](https://img.shields.io/github/stars/nikolaydubina/fpmoney.svg) - Fast and simple ISO4217 fixed-point decimal money.
- [go-finance](https://github.com/alpeb/go-finance) ![](https://img.shields.io/github/stars/alpeb/go-finance.svg) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.
- [go-finance](https://github.com/pieterclaerhout/go-finance) ![](https://img.shields.io/github/stars/pieterclaerhout/go-finance.svg) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.
- [go-money](https://github.com/rhymond/go-money) ![](https://img.shields.io/github/stars/rhymond/go-money.svg) - Implementation of Fowler's Money pattern.
- [go-nowpayments](https://github.com/matm/go-nowpayments) ![](https://img.shields.io/github/stars/matm/go-nowpayments.svg) - Library for the crypto NOWPayments API.
- [gobl](https://github.com/invopop/gobl) ![](https://img.shields.io/github/stars/invopop/gobl.svg) - Invoice and billing document framework. JSON Schema based. Automates tax calculations and validation, with tooling to convert into global formats.
- [ledger](https://github.com/formancehq/ledger) ![](https://img.shields.io/github/stars/formancehq/ledger.svg) - A programmable financial ledger that provides a foundation for money-moving applications.
- [money](https://github.com/govalues/money) ![](https://img.shields.io/github/stars/govalues/money.svg) - Immutable monetary amounts and exchange rates with panic-free arithmetic.
- [ofxgo](https://github.com/aclindsa/ofxgo) ![](https://img.shields.io/github/stars/aclindsa/ofxgo.svg) - Query OFX servers and/or parse the responses (with example command-line client).
- [orderbook](https://github.com/i25959341/orderbook) ![](https://img.shields.io/github/stars/i25959341/orderbook.svg) - Matching Engine for Limit Order Book in Golang.
- [payme](https://github.com/jovandeginste/payme) ![](https://img.shields.io/github/stars/jovandeginste/payme.svg) - QR code generator (ASCII & PNG) for SEPA payments.
- [swift](https://code.pfad.fr/swift/) - Offline validity check of IBAN (International Bank Account Number) and retrieval of BIC (for some countries).
- [techan](https://github.com/sdcoffey/techan) ![](https://img.shields.io/github/stars/sdcoffey/techan.svg) - Technical analysis library with advanced market analysis and trading strategies.
- [ticker](https://github.com/achannarasappa/ticker) ![](https://img.shields.io/github/stars/achannarasappa/ticker.svg) - Terminal stock watcher and stock position tracker.
- [transaction](https://github.com/claygod/transaction) ![](https://img.shields.io/github/stars/claygod/transaction.svg) - Embedded transactional database of accounts, running in multithreaded mode.
- [udecimal](https://github.com/quagmt/udecimal) ![](https://img.shields.io/github/stars/quagmt/udecimal.svg) - High performance, high precision, zero allocation fixed-point decimal library for financial applications.
- [vat](https://github.com/dannyvankooten/vat) ![](https://img.shields.io/github/stars/dannyvankooten/vat.svg) - VAT number validation & EU VAT rates.

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [bind](https://github.com/robfig/bind) ![](https://img.shields.io/github/stars/robfig/bind.svg) - Bind form data to any Go values.
- [checker](https://github.com/cinar/checker) ![](https://img.shields.io/github/stars/cinar/checker.svg) - Checker helps validating user input through rules defined in struct tags or directly through functions.
- [conform](https://github.com/leebenson/conform) ![](https://img.shields.io/github/stars/leebenson/conform.svg) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
- [form](https://github.com/go-playground/form) ![](https://img.shields.io/github/stars/go-playground/form.svg) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
- [formam](https://github.com/monoculum/formam) ![](https://img.shields.io/github/stars/monoculum/formam.svg) - decode form's values into a struct.
- [forms](https://github.com/albrow/forms) ![](https://img.shields.io/github/stars/albrow/forms.svg) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
- [gbind](https://github.com/bdjimmy/gbind) ![](https://img.shields.io/github/stars/bdjimmy/gbind.svg) - Bind data to any Go value. Can use built-in and custom expression binding capabilities; supports data validation
- [gorilla/csrf](https://github.com/gorilla/csrf) ![](https://img.shields.io/github/stars/gorilla/csrf.svg) - CSRF protection for Go web applications & services.
- [httpin](https://github.com/ggicci/httpin) ![](https://img.shields.io/github/stars/ggicci/httpin.svg) - Decode an HTTP request into a custom struct, including querystring, forms, HTTP headers, etc.
- [nosurf](https://github.com/justinas/nosurf) ![](https://img.shields.io/github/stars/justinas/nosurf.svg) - CSRF protection middleware for Go.
- [qs](https://github.com/sonh/qs) ![](https://img.shields.io/github/stars/sonh/qs.svg) - Go module for encoding structs into URL query parameters.
- [queryparam](https://github.com/tomwright/queryparam) ![](https://img.shields.io/github/stars/tomwright/queryparam.svg) - Decode `url.Values` into usable struct values of standard or custom types.
- [roamer](https://github.com/slipros/roamer) ![](https://img.shields.io/github/stars/slipros/roamer.svg) - Eliminates boilerplate code for parsing HTTP requests by binding cookies, headers, query params, path params, body to structs and more by using simple tags.

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [fp-go](https://github.com/repeale/fp-go) ![](https://img.shields.io/github/stars/repeale/fp-go.svg) - Collection of Functional Programming helpers powered by Golang 1.18+ generics.
- [fpGo](https://github.com/TeaEntityLab/fpGo) ![](https://img.shields.io/github/stars/TeaEntityLab/fpGo.svg) - Monad, Functional Programming features for Golang.
- [fuego](https://github.com/seborama/fuego) ![](https://img.shields.io/github/stars/seborama/fuego.svg) - Functional Experiment in Go.
- [FuncFrog](https://github.com/koss-null/FuncFrog) ![](https://img.shields.io/github/stars/koss-null/FuncFrog.svg) - Functional helpers library providing Map, Filter, Reduce and other stream operations on generic slices Go1.18+ with lazy evaluation and error handling mechanisms.
- [go-functional](https://github.com/BooleanCat/go-functional) ![](https://img.shields.io/github/stars/BooleanCat/go-functional.svg) - Functional programming in Go using generics
- [go-underscore](https://github.com/tobyhede/go-underscore) ![](https://img.shields.io/github/stars/tobyhede/go-underscore.svg) - Useful collection of helpfully functional Go collection utilities.
- [gofp](https://github.com/rbrahul/gofp) ![](https://img.shields.io/github/stars/rbrahul/gofp.svg) - A lodash like powerful utility library for Golang.
- [mo](https://github.com/samber/mo) ![](https://img.shields.io/github/stars/samber/mo.svg) - Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...).
- [underscore](https://github.com/rjNemo/underscore) ![](https://img.shields.io/github/stars/rjNemo/underscore.svg) - Functional programming helpers for Go 1.18 and beyond.
- [valor](https://github.com/phelmkamp/valor) ![](https://img.shields.io/github/stars/phelmkamp/valor.svg) - Generic option and result types that optionally contain a value.

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Ark](https://github.com/mlange-42/ark) ![](https://img.shields.io/github/stars/mlange-42/ark.svg) - Archetype-based Entity Component System (ECS) for Go.
- [Ebitengine](https://github.com/hajimehoshi/ebiten) ![](https://img.shields.io/github/stars/hajimehoshi/ebiten.svg) - dead simple 2D game engine in Go.
- [ecs](https://github.com/andygeiss/ecs) ![](https://img.shields.io/github/stars/andygeiss/ecs.svg) - Build your own Game-Engine based on the Entity Component System concept in Golang.
- [engo](https://github.com/EngoEngine/engo) ![](https://img.shields.io/github/stars/EngoEngine/engo.svg) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.
- [fantasyname](https://github.com/s0rg/fantasyname) ![](https://img.shields.io/github/stars/s0rg/fantasyname.svg) - Fantasy names generator.
- [g3n](https://github.com/g3n/engine) ![](https://img.shields.io/github/stars/g3n/engine.svg) - Go 3D Game Engine.
- [go-astar](https://github.com/beefsack/go-astar) ![](https://img.shields.io/github/stars/beefsack/go-astar.svg) - Go implementation of the A\* path finding algorithm.
- [go-sdl2](https://github.com/veandco/go-sdl2) ![](https://img.shields.io/github/stars/veandco/go-sdl2.svg) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
- [go3d](https://github.com/ungerik/go3d) ![](https://img.shields.io/github/stars/ungerik/go3d.svg) - Performance oriented 2D/3D math package for Go.
- [gonet](https://github.com/xtaci/gonet) ![](https://img.shields.io/github/stars/xtaci/gonet.svg) - Game server skeleton implemented with golang.
- [goworld](https://github.com/xiaonanln/goworld) ![](https://img.shields.io/github/stars/xiaonanln/goworld.svg) - Scalable game server engine, featuring space-entity framework and hot-swapping.
- [grid](https://github.com/s0rg/grid) ![](https://img.shields.io/github/stars/s0rg/grid.svg) - Generic 2D grid with ray-casting, shadow-casting and path finding.
- [Leaf](https://github.com/name5566/leaf) ![](https://img.shields.io/github/stars/name5566/leaf.svg) - Lightweight game server framework.
- [nano](https://github.com/lonng/nano) ![](https://img.shields.io/github/stars/lonng/nano.svg) - Lightweight, facility, high performance golang based game server framework.
- [Oak](https://github.com/oakmound/oak) ![](https://img.shields.io/github/stars/oakmound/oak.svg) - Pure Go game engine.
- [Pi](https://github.com/elgopher/pi) ![](https://img.shields.io/github/stars/elgopher/pi.svg) - Game engine for creating retro games for modern computers. Inspired by Pico-8 and powered by Ebitengine.
- [Pitaya](https://github.com/topfreegames/pitaya) ![](https://img.shields.io/github/stars/topfreegames/pitaya.svg) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.
- [Pixel](https://github.com/gopxl/pixel) ![](https://img.shields.io/github/stars/gopxl/pixel.svg) - Hand-crafted 2D game library in Go.
- [prototype](https://github.com/gonutz/prototype) ![](https://img.shields.io/github/stars/gonutz/prototype.svg) - Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API.
- [raylib-go](https://github.com/gen2brain/raylib-go) ![](https://img.shields.io/github/stars/gen2brain/raylib-go.svg) - Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
- [termloop](https://github.com/JoelOtter/termloop) ![](https://img.shields.io/github/stars/JoelOtter/termloop.svg) - Terminal-based game engine for Go, built on top of Termbox.
- [tile](https://github.com/kelindar/tile) ![](https://img.shields.io/github/stars/kelindar/tile.svg) - Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export.

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [convergen](https://github.com/reedom/convergen) ![](https://img.shields.io/github/stars/reedom/convergen.svg) - Feature rich type-to-type copy code generator.
- [copygen](https://github.com/switchupcb/copygen) ![](https://img.shields.io/github/stars/switchupcb/copygen.svg) - Generate any code based on Go types, including type-to-type converters (copy code) without reflection by default.
- [generis](https://github.com/senselogic/GENERIS) ![](https://img.shields.io/github/stars/senselogic/GENERIS.svg) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
- [go-enum](https://github.com/abice/go-enum) ![](https://img.shields.io/github/stars/abice/go-enum.svg) - Code generation for enums from code comments.
- [go-enum-encoding](https://github.com/nikolaydubina/go-enum-encoding) ![](https://img.shields.io/github/stars/nikolaydubina/go-enum-encoding.svg) - Code generation for enum encoding from code comments.
- [go-linq](https://github.com/ahmetalpbalkan/go-linq) ![](https://img.shields.io/github/stars/ahmetalpbalkan/go-linq.svg) - .NET LINQ-like query methods for Go.
- [goderive](https://github.com/awalterschulze/goderive) ![](https://img.shields.io/github/stars/awalterschulze/goderive.svg) - Derives functions from input types
- [goverter](https://github.com/jmattheis/goverter) ![](https://img.shields.io/github/stars/jmattheis/goverter.svg) - Generate converters by defining an interface.
- [GoWrap](https://github.com/hexdigest/gowrap) ![](https://img.shields.io/github/stars/hexdigest/gowrap.svg) - Generate decorators for Go interfaces using simple templates.
- [interfaces](https://github.com/rjeczalik/interfaces) ![](https://img.shields.io/github/stars/rjeczalik/interfaces.svg) - Command line tool for generating interface definitions.
- [jennifer](https://github.com/dave/jennifer) ![](https://img.shields.io/github/stars/dave/jennifer.svg) - Generate arbitrary Go code without templates.
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) ![](https://img.shields.io/github/stars/deepmap/oapi-codegen.svg) - This package contains a set of utilities for generating Go boilerplate code for services based on OpenAPI 3.0 API definitions.
- [typeregistry](https://github.com/xiaoxin01/typeregistry) ![](https://img.shields.io/github/stars/xiaoxin01/typeregistry.svg) - A library to create type dynamically.

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [borders](https://github.com/kpfaulkner/borders) ![](https://img.shields.io/github/stars/kpfaulkner/borders.svg) - Detects image borders and converts to GeoJSON for GIS operations.
- [geoos](https://github.com/spatial-go/geoos) ![](https://img.shields.io/github/stars/spatial-go/geoos.svg) - A library provides spatial data and geometric algorithms.
- [geoserver](https://github.com/hishamkaram/geoserver) ![](https://img.shields.io/github/stars/hishamkaram/geoserver.svg) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
- [gismanager](https://github.com/hishamkaram/gismanager) ![](https://img.shields.io/github/stars/hishamkaram/gismanager.svg) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.
- [godal](https://github.com/airbusgeo/godal) ![](https://img.shields.io/github/stars/airbusgeo/godal.svg) - Go wrapper for GDAL.
- [H3](https://github.com/uber/h3-go) ![](https://img.shields.io/github/stars/uber/h3-go.svg) - Go bindings for H3, a hierarchical hexagonal geospatial indexing system.
- [H3 GeoJSON](https://github.com/mmadfox/go-geojson2h3) ![](https://img.shields.io/github/stars/mmadfox/go-geojson2h3.svg) - Conversion utilities between H3 indexes and GeoJSON.
- [H3GeoDist](https://github.com/mmadfox/go-h3geo-dist) ![](https://img.shields.io/github/stars/mmadfox/go-h3geo-dist.svg) - Distribution of Uber H3geo cells by virtual nodes.
- [mbtileserver](https://github.com/consbio/mbtileserver) ![](https://img.shields.io/github/stars/consbio/mbtileserver.svg) - A simple Go-based server for map tiles stored in mbtiles format.
- [osm](https://github.com/paulmach/osm) ![](https://img.shields.io/github/stars/paulmach/osm.svg) - Library for reading, writing and working with OpenStreetMap data and APIs.
- [pbf](https://github.com/maguro/pbf) ![](https://img.shields.io/github/stars/maguro/pbf.svg) - OpenStreetMap PBF golang encoder/decoder.
- [S2 geojson](https://github.com/pantrif/s2-geojson) ![](https://img.shields.io/github/stars/pantrif/s2-geojson.svg) - Convert geojson to s2 cells & demonstrating some S2 geometry features on map.
- [S2 geometry](https://github.com/golang/geo) ![](https://img.shields.io/github/stars/golang/geo.svg) - S2 geometry library in Go.
- [simplefeatures](https://github.com/peterstace/simplefeatures) ![](https://img.shields.io/github/stars/peterstace/simplefeatures.svg) - simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them.
- [Tile38](https://github.com/tidwall/tile38) ![](https://img.shields.io/github/stars/tidwall/tile38.svg) - Geolocation DB with spatial index and realtime geofencing.
- [Web-Mercator-Projection](https://github.com/jorelosorio/web-mercator-projection) ![](https://img.shields.io/github/stars/jorelosorio/web-mercator-projection.svg) A project to easily use and convert LonLat, Point and Tile to display info, markers, etc, in a map using the Web Mercator Projection.
- [WGS84](https://github.com/wroge/wgs84) ![](https://img.shields.io/github/stars/wroge/wgs84.svg) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages and vice-versa._

- [bunster](https://github.com/yassinebenaid/bunster) ![](https://img.shields.io/github/stars/yassinebenaid/bunster.svg) - Compile shell scripts to Go.
- [c4go](https://github.com/Konstantin8105/c4go) ![](https://img.shields.io/github/stars/Konstantin8105/c4go.svg) - Transpile C code to Go code.
- [cxgo](https://github.com/gotranspile/cxgo) ![](https://img.shields.io/github/stars/gotranspile/cxgo.svg) - Transpile C code to Go code.
- [esp32](https://github.com/andygeiss/esp32-transpiler) ![](https://img.shields.io/github/stars/andygeiss/esp32-transpiler.svg) - Transpile Go into Arduino code.
- [f4go](https://github.com/Konstantin8105/f4go) ![](https://img.shields.io/github/stars/Konstantin8105/f4go.svg) - Transpile FORTRAN 77 code to Go code.
- [go2hx](https://github.com/go2hx/go2hx) ![](https://img.shields.io/github/stars/go2hx/go2hx.svg) - Compiler from Go to Haxe to Javascript/C++/Java/C#.
- [gopherjs](https://github.com/gopherjs/gopherjs) ![](https://img.shields.io/github/stars/gopherjs/gopherjs.svg) - Compiler from Go to JavaScript.

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [anchor](https://github.com/kyuff/anchor) ![](https://img.shields.io/github/stars/kyuff/anchor.svg) - Library to manage component lifecycle in microservice architectures.
- [ants](https://github.com/panjf2000/ants) ![](https://img.shields.io/github/stars/panjf2000/ants.svg) - A high-performance and low-cost goroutine pool in Go.
- [artifex](https://github.com/borderstech/artifex) ![](https://img.shields.io/github/stars/borderstech/artifex.svg) - Simple in-memory job queue for Golang using worker-based dispatching.
- [async](https://github.com/yaitoo/async) ![](https://img.shields.io/github/stars/yaitoo/async.svg) - An asynchronous task package with async/await style for Go.
- [async](https://github.com/reugn/async) ![](https://img.shields.io/github/stars/reugn/async.svg) - An alternative sync library for Go (Future, Promise, Locks).
- [async](https://github.com/studiosol/async) ![](https://img.shields.io/github/stars/studiosol/async.svg) - A safe way to execute functions asynchronously, recovering them in case of panic.
- [async-job](https://github.com/lab210-dev/async-job) ![](https://img.shields.io/github/stars/lab210-dev/async-job.svg) - AsyncJob is an asynchronous queue job manager with light code, clear and speed.
- [breaker](https://github.com/kamilsk/breaker) ![](https://img.shields.io/github/stars/kamilsk/breaker.svg) - Flexible mechanism to make execution flow interruptible.
- [channelify](https://github.com/ddelizia/channelify) ![](https://img.shields.io/github/stars/ddelizia/channelify.svg) - Transform your function to return channels for easy and powerful parallel processing.
- [conc](https://github.com/sourcegraph/conc) ![](https://img.shields.io/github/stars/sourcegraph/conc.svg) - `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer.
- [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) ![](https://img.shields.io/github/stars/vivek-ng/concurrency-limiter.svg) - Concurrency limiter with support for timeouts, dynamic priority and context cancellation of goroutines.
- [conexec](https://github.com/ITcathyh/conexec) ![](https://img.shields.io/github/stars/ITcathyh/conexec.svg) - A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.
- [cyclicbarrier](https://github.com/marusama/cyclicbarrier) ![](https://img.shields.io/github/stars/marusama/cyclicbarrier.svg) - CyclicBarrier for golang.
- [execpool](https://github.com/hexdigest/execpool) ![](https://img.shields.io/github/stars/hexdigest/execpool.svg) - A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command.
- [flowmatic](https://github.com/carlmjohnson/flowmatic) ![](https://img.shields.io/github/stars/carlmjohnson/flowmatic.svg) - Structured concurrency made easy.
- [go-accumulator](https://github.com/nar10z/go-accumulator) ![](https://img.shields.io/github/stars/nar10z/go-accumulator.svg) - Solution for accumulation of events and their subsequent processing.
- [go-actor](https://github.com/vladopajic/go-actor) ![](https://img.shields.io/github/stars/vladopajic/go-actor.svg) - A tiny library for writing concurrent programs using actor model.
- [go-floc](https://github.com/workanator/go-floc) ![](https://img.shields.io/github/stars/workanator/go-floc.svg) - Orchestrate goroutines with ease.
- [go-flow](https://github.com/kamildrazkiewicz/go-flow) ![](https://img.shields.io/github/stars/kamildrazkiewicz/go-flow.svg) - Control goroutines execution order.
- [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) ![](https://img.shields.io/github/stars/nikhilsaraf/go-tools.svg) - Manage a pool of goroutines using this lightweight library with a simple API.
- [go-trylock](https://github.com/subchen/go-trylock) ![](https://img.shields.io/github/stars/subchen/go-trylock.svg) - TryLock support on read-write lock for Golang.
- [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) ![](https://img.shields.io/github/stars/pieterclaerhout/go-waitgroup.svg) - Like `sync.WaitGroup` with error handling and concurrency control.
- [go-workerpool](https://github.com/zenthangplus/go-workerpool) ![](https://img.shields.io/github/stars/zenthangplus/go-workerpool.svg) - Inspired from Java Thread Pool, Go WorkerPool aims to control heavy Go Routines.
- [goccm](https://github.com/zenthangplus/goccm) ![](https://img.shields.io/github/stars/zenthangplus/goccm.svg) - Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently.
- [gohive](https://github.com/loveleshsharma/gohive) ![](https://img.shields.io/github/stars/loveleshsharma/gohive.svg) - A highly performant and easy to use Goroutine pool for Go.
- [gollback](https://github.com/vardius/gollback) ![](https://img.shields.io/github/stars/vardius/gollback.svg) - asynchronous simple function utilities, for managing execution of closures and callbacks.
- [gowl](https://github.com/hamed-yousefi/gowl) ![](https://img.shields.io/github/stars/hamed-yousefi/gowl.svg) - Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status.
- [goworker](https://github.com/benmanns/goworker) ![](https://img.shields.io/github/stars/benmanns/goworker.svg) - goworker is a Go-based background worker.
- [gowp](https://github.com/xxjwxc/gowp) ![](https://img.shields.io/github/stars/xxjwxc/gowp.svg) - gowp is concurrency limiting goroutine pool.
- [gpool](https://github.com/Sherifabdlnaby/gpool) ![](https://img.shields.io/github/stars/Sherifabdlnaby/gpool.svg) - manages a resizeable pool of context-aware goroutines to bound concurrency.
- [grpool](https://github.com/ivpusic/grpool) ![](https://img.shields.io/github/stars/ivpusic/grpool.svg) - Lightweight Goroutine pool.
- [hands](https://github.com/duanckham/hands) ![](https://img.shields.io/github/stars/duanckham/hands.svg) - A process controller used to control the execution and return strategies of multiple goroutines.
- [Hunch](https://github.com/AaronJan/Hunch) ![](https://img.shields.io/github/stars/AaronJan/Hunch.svg) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
- [kyoo](https://github.com/dirkaholic/kyoo) ![](https://img.shields.io/github/stars/dirkaholic/kyoo.svg) - Provides an unlimited job queue and concurrent worker pools.
- [neilotoole/errgroup](https://github.com/neilotoole/errgroup) ![](https://img.shields.io/github/stars/neilotoole/errgroup.svg) - Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines.
- [nursery](https://github.com/arunsworld/nursery) ![](https://img.shields.io/github/stars/arunsworld/nursery.svg) - Structured concurrency in Go.
- [oversight](https://pkg.go.dev/cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.
- [parallel-fn](https://github.com/rafaeljesus/parallel-fn) ![](https://img.shields.io/github/stars/rafaeljesus/parallel-fn.svg) - Run functions in parallel.
- [pond](https://github.com/alitto/pond) ![](https://img.shields.io/github/stars/alitto/pond.svg) - Minimalistic and High-performance goroutine worker pool written in Go.
- [pool](https://github.com/go-playground/pool) ![](https://img.shields.io/github/stars/go-playground/pool.svg) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.
- [rill](https://github.com/destel/rill) ![](https://img.shields.io/github/stars/destel/rill.svg) - Go toolkit for clean, composable, channel-based concurrency.
- [routine](https://github.com/timandy/routine) ![](https://img.shields.io/github/stars/timandy/routine.svg) - `routine` is a `ThreadLocal` for go library. It encapsulates and provides some easy-to-use, non-competitive, high-performance `goroutine` context access interfaces, which can help you access coroutine context information more gracefully.
- [routine](https://github.com/x-mod/routine) ![](https://img.shields.io/github/stars/x-mod/routine.svg) - go routine control with context, support: Main, Go, Pool and some useful Executors.
- [semaphore](https://github.com/kamilsk/semaphore) ![](https://img.shields.io/github/stars/kamilsk/semaphore.svg) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
- [semaphore](https://github.com/marusama/semaphore) ![](https://img.shields.io/github/stars/marusama/semaphore.svg) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
- [stl](https://github.com/ssgreg/stl) ![](https://img.shields.io/github/stars/ssgreg/stl.svg) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
- [threadpool](https://github.com/shettyh/threadpool) ![](https://img.shields.io/github/stars/shettyh/threadpool.svg) - Golang threadpool implementation.
- [tunny](https://github.com/Jeffail/tunny) ![](https://img.shields.io/github/stars/Jeffail/tunny.svg) - Goroutine pool for golang.
- [worker-pool](https://github.com/vardius/worker-pool) ![](https://img.shields.io/github/stars/vardius/worker-pool.svg) - goworker is a Go simple async worker pool.
- [workerpool](https://github.com/gammazero/workerpool) ![](https://img.shields.io/github/stars/gammazero/workerpool.svg) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [app](https://github.com/murlokswarm/app) ![](https://img.shields.io/github/stars/murlokswarm/app.svg) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.
- [cimgui-go](https://github.com/AllenDang/cimgui-go) ![](https://img.shields.io/github/stars/AllenDang/cimgui-go.svg) - Auto generated Go wrapper for [Dear ImGui](https://github.com/ocornut/imgui) via [cimgui](https://github.com/cimgui/cimgui).
- [Cogent Core](https://github.com/cogentcore/core) ![](https://img.shields.io/github/stars/cogentcore/core.svg) - A framework for building 2D and 3D apps that run on macOS, Windows, Linux, iOS, Android, and the web.
- [DarwinKit](https://github.com/progrium/darwinkit) ![](https://img.shields.io/github/stars/progrium/darwinkit.svg) - Build native macOS applications using Go.
- [energy](https://github.com/energye/energy) ![](https://img.shields.io/github/stars/energye/energy.svg) - Cross-platform based on LCL(Native System UI Control Library) and CEF(Chromium Embedded Framework) (Windows/ macOS / Linux)
- [fyne](https://github.com/fyne-io/fyne) ![](https://img.shields.io/github/stars/fyne-io/fyne.svg) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.
- [gio](https://gioui.org) - Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD, OpenBSD and WebAssembly.
- [go-gtk](https://mattn.github.io/go-gtk/) - Go bindings for GTK.
- [go-sciter](https://github.com/sciter-sdk/go-sciter) ![](https://img.shields.io/github/stars/sciter-sdk/go-sciter.svg) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.
- [Goey](https://bitbucket.org/rj/goey/src/master/) - Cross platform UI toolkit aggregator for Windows / Linux / Mac. GTK, Cocoa, Windows API
- [goradd/html5tag](https://github.com/goradd/html5tag) ![](https://img.shields.io/github/stars/goradd/html5tag.svg) - Library for outputting HTML5 tags.
- [gotk3](https://github.com/gotk3/gotk3) ![](https://img.shields.io/github/stars/gotk3/gotk3.svg) - Go bindings for GTK3.
- [gowd](https://github.com/dtylman/gowd) ![](https://img.shields.io/github/stars/dtylman/gowd.svg) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.
- [qt](https://github.com/therecipe/qt) ![](https://img.shields.io/github/stars/therecipe/qt.svg) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).
- [Spot](https://github.com/roblillack/spot) ![](https://img.shields.io/github/stars/roblillack/spot.svg) - Reactive, cross-platform desktop GUI toolkit.
- [ui](https://github.com/andlabs/ui) ![](https://img.shields.io/github/stars/andlabs/ui.svg) - Platform-native GUI library for Go. Cross platform.
- [unison](https://github.com/richardwilkes/unison) ![](https://img.shields.io/github/stars/richardwilkes/unison.svg) - A unified graphical user experience toolkit for Go desktop applications. macOS, Windows, and Linux are supported.
- [Wails](https://wails.io) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
- [walk](https://github.com/lxn/walk) ![](https://img.shields.io/github/stars/lxn/walk.svg) - Windows application library kit for Go.
- [webview](https://github.com/zserge/webview) ![](https://img.shields.io/github/stars/zserge/webview.svg) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).

_Interaction_

- [AppIndicator Go](https://github.com/gopherlibs/appindicator) ![](https://img.shields.io/github/stars/gopherlibs/appindicator.svg) - Go bindings for libappindicator3 C library.
- [gosx-notifier](https://github.com/deckarep/gosx-notifier) ![](https://img.shields.io/github/stars/deckarep/gosx-notifier.svg) - OSX Desktop Notifications library for Go.
- [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) ![](https://img.shields.io/github/stars/prashantgupta24/activity-tracker.svg) - OSX library to notify about any (pluggable) activity on your machine.
- [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) ![](https://img.shields.io/github/stars/prashantgupta24/mac-sleep-notifier.svg) - OSX Sleep/Wake notifications in golang.
- [robotgo](https://github.com/go-vgo/robotgo) ![](https://img.shields.io/github/stars/go-vgo/robotgo.svg) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.
- [systray](https://github.com/getlantern/systray) ![](https://img.shields.io/github/stars/getlantern/systray.svg) - Cross platform Go library to place an icon and menu in the notification area.
- [trayhost](https://github.com/shurcooL/trayhost) ![](https://img.shields.io/github/stars/shurcooL/trayhost.svg) - Cross-platform Go library to place an icon in the host operating system's taskbar.
- [zenity](https://github.com/ncruces/zenity) ![](https://img.shields.io/github/stars/ncruces/zenity.svg) - Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user.

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli](https://github.com/arduino/arduino-cli) ![](https://img.shields.io/github/stars/arduino/arduino-cli.svg) - Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects.
- [emgo](https://github.com/ziutek/emgo) ![](https://img.shields.io/github/stars/ziutek/emgo.svg) - Go-like language for programming embedded systems (e.g. STM32 MCU).
- [ghw](https://github.com/jaypipes/ghw) ![](https://img.shields.io/github/stars/jaypipes/ghw.svg) - Golang hardware discovery/inspection library.
- [go-osc](https://github.com/hypebeast/go-osc) ![](https://img.shields.io/github/stars/hypebeast/go-osc.svg) - Open Sound Control (OSC) bindings for Go.
- [go-rpio](https://github.com/stianeikeland/go-rpio) ![](https://img.shields.io/github/stars/stianeikeland/go-rpio.svg) - GPIO for Go, doesn't require cgo.
- [goroslib](https://github.com/aler9/goroslib) ![](https://img.shields.io/github/stars/aler9/goroslib.svg) - Robot Operating System (ROS) library for Go.
- [joystick](https://github.com/0xcafed00d/joystick) ![](https://img.shields.io/github/stars/0xcafed00d/joystick.svg) - a polled API to read the state of an attached joystick.
- [sysinfo](https://github.com/zcalusic/sysinfo) ![](https://img.shields.io/github/stars/zcalusic/sysinfo.svg) - A pure Go library providing Linux OS / kernel / hardware system information.

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [bild](https://github.com/anthonynsimon/bild) ![](https://img.shields.io/github/stars/anthonynsimon/bild.svg) - Collection of image processing algorithms in pure Go.
- [bimg](https://github.com/h2non/bimg) ![](https://img.shields.io/github/stars/h2non/bimg.svg) - Small package for fast and efficient image processing using libvips.
- [cameron](https://github.com/aofei/cameron) ![](https://img.shields.io/github/stars/aofei/cameron.svg) - An avatar generator for Go.
- [canvas](https://github.com/tdewolff/canvas) ![](https://img.shields.io/github/stars/tdewolff/canvas.svg) - Vector graphics to PDF, SVG or rasterized image.
- [color-extractor](https://github.com/marekm4/color-extractor) ![](https://img.shields.io/github/stars/marekm4/color-extractor.svg) - Dominant color extractor with no external dependencies.
- [darkroom](https://github.com/gojek/darkroom) ![](https://img.shields.io/github/stars/gojek/darkroom.svg) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.
- [geopattern](https://github.com/pravj/geopattern) ![](https://img.shields.io/github/stars/pravj/geopattern.svg) - Create beautiful generative image patterns from a string.
- [gg](https://github.com/fogleman/gg) ![](https://img.shields.io/github/stars/fogleman/gg.svg) - 2D rendering in pure Go.
- [gift](https://github.com/disintegration/gift) ![](https://img.shields.io/github/stars/disintegration/gift.svg) - Package of image processing filters.
- [gltf](https://github.com/qmuntal/gltf) ![](https://img.shields.io/github/stars/qmuntal/gltf.svg) - Efficient and robust glTF 2.0 reader, writer and validator.
- [go-cairo](https://github.com/ungerik/go-cairo) ![](https://img.shields.io/github/stars/ungerik/go-cairo.svg) - Go binding for the cairo graphics library.
- [go-gd](https://github.com/bolknote/go-gd) ![](https://img.shields.io/github/stars/bolknote/go-gd.svg) - Go binding for GD library.
- [go-nude](https://github.com/koyachi/go-nude) ![](https://img.shields.io/github/stars/koyachi/go-nude.svg) - Nudity detection with Go.
- [go-qrcode](https://github.com/yeqown/go-qrcode) ![](https://img.shields.io/github/stars/yeqown/go-qrcode.svg) - Generate QR codes with personalized styles, allowing adjustments to color, block size, shape, and icons.
- [go-webcolors](https://github.com/jyotiska/go-webcolors) ![](https://img.shields.io/github/stars/jyotiska/go-webcolors.svg) - Port of webcolors library from Python to Go.
- [go-webp](https://github.com/kolesa-team/go-webp) ![](https://img.shields.io/github/stars/kolesa-team/go-webp.svg) - Library for encode and decode webp pictures, using libwebp.
- [gocv](https://github.com/hybridgroup/gocv) ![](https://img.shields.io/github/stars/hybridgroup/gocv.svg) - Go package for computer vision using OpenCV 3.3+.
- [goimagehash](https://github.com/corona10/goimagehash) ![](https://img.shields.io/github/stars/corona10/goimagehash.svg) - Go Perceptual image hashing package.
- [goimghdr](https://github.com/corona10/goimghdr) ![](https://img.shields.io/github/stars/corona10/goimghdr.svg) - The imghdr module determines the type of image contained in a file for Go.
- [govatar](https://github.com/o1egl/govatar) ![](https://img.shields.io/github/stars/o1egl/govatar.svg) - Library and CMD tool for generating funny avatars.
- [govips](https://github.com/davidbyttow/govips) ![](https://img.shields.io/github/stars/davidbyttow/govips.svg) - A lightning fast image processing and resizing library for Go.
- [gowitness](https://github.com/sensepost/gowitness) ![](https://img.shields.io/github/stars/sensepost/gowitness.svg) - Screenshoting webpages using go and headless chrome on command line.
- [gridder](https://github.com/shomali11/gridder) ![](https://img.shields.io/github/stars/shomali11/gridder.svg) - A Grid based 2D Graphics library.
- [image2ascii](https://github.com/qeesung/image2ascii) ![](https://img.shields.io/github/stars/qeesung/image2ascii.svg) - Convert image to ASCII.
- [imagick](https://github.com/gographics/imagick) ![](https://img.shields.io/github/stars/gographics/imagick.svg) - Go binding to ImageMagick's MagickWand C API.
- [imaginary](https://github.com/h2non/imaginary) ![](https://img.shields.io/github/stars/h2non/imaginary.svg) - Fast and simple HTTP microservice for image resizing.
- [imaging](https://github.com/disintegration/imaging) ![](https://img.shields.io/github/stars/disintegration/imaging.svg) - Simple Go image processing package.
- [imagor](https://github.com/cshum/imagor) ![](https://img.shields.io/github/stars/cshum/imagor.svg) - Fast, secure image processing server and Go library, using libvips.
- [img](https://github.com/hawx/img) ![](https://img.shields.io/github/stars/hawx/img.svg) - Selection of image manipulation tools.
- [ln](https://github.com/fogleman/ln) ![](https://img.shields.io/github/stars/fogleman/ln.svg) - 3D line art rendering in Go.
- [mergi](https://github.com/noelyahan/mergi) ![](https://img.shields.io/github/stars/noelyahan/mergi.svg) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
- [mort](https://github.com/aldor007/mort) ![](https://img.shields.io/github/stars/aldor007/mort.svg) - Storage and image processing server written in Go.
- [mpo](https://github.com/donatj/mpo) ![](https://img.shields.io/github/stars/donatj/mpo.svg) - Decoder and conversion tool for MPO 3D Photos.
- [nativewebp](https://github.com/HugoSmits86/nativewebp) ![](https://img.shields.io/github/stars/HugoSmits86/nativewebp.svg) - Go native WebP encoder with zero external dependencies.
- [picfit](https://github.com/thoas/picfit) ![](https://img.shields.io/github/stars/thoas/picfit.svg) - An image resizing server written in Go.
- [pt](https://github.com/fogleman/pt) ![](https://img.shields.io/github/stars/fogleman/pt.svg) - Path tracing engine written in Go.
- [scout](https://github.com/jonoton/scout) ![](https://img.shields.io/github/stars/jonoton/scout.svg) - Scout is a standalone open source software solution for DIY video security.
- [smartcrop](https://github.com/muesli/smartcrop) ![](https://img.shields.io/github/stars/muesli/smartcrop.svg) - Finds good crops for arbitrary images and crop sizes.
- [steganography](https://github.com/auyer/steganography) ![](https://img.shields.io/github/stars/auyer/steganography.svg) - Pure Go Library for LSB steganography.
- [stegify](https://github.com/DimitarPetrov/stegify) ![](https://img.shields.io/github/stars/DimitarPetrov/stegify.svg) - Go tool for LSB steganography, capable of hiding any file within an image.
- [svgo](https://github.com/ajstarks/svgo) ![](https://img.shields.io/github/stars/ajstarks/svgo.svg) - Go Language Library for SVG generation.
- [transformimgs](https://github.com/Pixboost/transformimgs) ![](https://img.shields.io/github/stars/Pixboost/transformimgs.svg) - Transformimgs resizes and optimises images for Web using next-generation formats.
- [webp-server](https://github.com/mehdipourfar/webp-server) ![](https://img.shields.io/github/stars/mehdipourfar/webp-server.svg) - Simple and minimal image server capable of storing, resizing, converting and caching images.

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [connectordb](https://github.com/connectordb/connectordb) ![](https://img.shields.io/github/stars/connectordb/connectordb.svg) - Open-Source Platform for Quantified Self & IoT.
- [devices](https://github.com/goiot/devices) ![](https://img.shields.io/github/stars/goiot/devices.svg) - Suite of libraries for IoT devices, experimental for x/exp/io.
- [ekuiper](https://github.com/lf-edge/ekuiper) ![](https://img.shields.io/github/stars/lf-edge/ekuiper.svg) - Lightweight data stream processing engine for IoT edge.
- [eywa](https://github.com/xcodersun/eywa) ![](https://img.shields.io/github/stars/xcodersun/eywa.svg) - Project Eywa is essentially a connection manager that keeps track of connected devices.
- [flogo](https://github.com/tibcosoftware/flogo) ![](https://img.shields.io/github/stars/tibcosoftware/flogo.svg) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
- [gatt](https://github.com/paypal/gatt) ![](https://img.shields.io/github/stars/paypal/gatt.svg) - Gatt is a Go package for building Bluetooth Low Energy peripherals.
- [gobot](https://github.com/hybridgroup/gobot/) ![](https://img.shields.io/github/stars/hybridgroup/gobot.svg) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
- [huego](https://github.com/amimof/huego) ![](https://img.shields.io/github/stars/amimof/huego.svg) - An extensive Philips Hue client library for Go.
- [iot](https://github.com/vaelen/iot/) ![](https://img.shields.io/github/stars/vaelen/iot.svg) - IoT is a simple framework for implementing a Google IoT Core device.
- [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.
- [rulego](https://github.com/rulego/rulego) ![](https://img.shields.io/github/stars/rulego/rulego.svg) - RuleGo is a lightweight, high-performance, embedded, orchestrable component-based rule engine for IoT edge.
- [sensorbee](https://github.com/sensorbee/sensorbee) ![](https://img.shields.io/github/stars/sensorbee/sensorbee.svg) - Lightweight stream processing engine for IoT.
- [shifu](https://github.com/Edgenesis/shifu) ![](https://img.shields.io/github/stars/Edgenesis/shifu.svg) - Kubernetes native IoT development framework.
- [smart-home](https://github.com/e154/smart-home) ![](https://img.shields.io/github/stars/e154/smart-home.svg) - Software package for IoT automation.

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._

- [cdule](https://github.com/deepaksinghvi/cdule) ![](https://img.shields.io/github/stars/deepaksinghvi/cdule.svg) - Job scheduler library with database support
- [cheek](https://github.com/bart6114/cheek) ![](https://img.shields.io/github/stars/bart6114/cheek.svg) - A simple crontab like scheduler that aims to offer a KISS approach to job scheduling.
- [clockwerk](https://github.com/onatm/clockwerk) ![](https://img.shields.io/github/stars/onatm/clockwerk.svg) - Go package to schedule periodic jobs using a simple, fluent syntax.
- [cronticker](https://github.com/krayzpipes/cronticker) ![](https://img.shields.io/github/stars/krayzpipes/cronticker.svg) - A ticker implementation to support cron schedules.
- [go-cron](https://github.com/rk/go-cron) ![](https://img.shields.io/github/stars/rk/go-cron.svg) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
- [go-job](https://github.com/cybergarage/go-job) ![](https://img.shields.io/github/stars/cybergarage/go-job.svg) - A flexible and extensible job scheduling and execution library for Go.
- [go-quartz](https://github.com/reugn/go-quartz) ![](https://img.shields.io/github/stars/reugn/go-quartz.svg) - Simple, zero-dependency scheduling library for Go.
- [go-scheduler](https://github.com/pardnchiu/go-scheduler) ![](https://img.shields.io/github/stars/pardnchiu/go-scheduler.svg) - Job scheduler supporting standard cron expressions, custom descriptors, intervals, and task dependencies.
- [gocron](https://github.com/go-co-op/gocron) ![](https://img.shields.io/github/stars/go-co-op/gocron.svg) - Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron).
- [goflow](https://github.com/fieldryand/goflow) ![](https://img.shields.io/github/stars/fieldryand/goflow.svg) - A simple but powerful DAG scheduler and dashboard.
- [gron](https://github.com/roylee0704/gron) ![](https://img.shields.io/github/stars/roylee0704/gron.svg) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.
- [gronx](https://github.com/adhocore/gronx) ![](https://img.shields.io/github/stars/adhocore/gronx.svg) - Cron expression parser, task runner and daemon consuming crontab like task list.
- [JobRunner](https://github.com/bamzi/jobrunner) ![](https://img.shields.io/github/stars/bamzi/jobrunner.svg) - Smart and featureful cron job scheduler with job queuing and live monitoring built in.
- [leprechaun](https://github.com/kilgaloon/leprechaun) ![](https://img.shields.io/github/stars/kilgaloon/leprechaun.svg) - Job scheduler that supports webhooks, crons and classic scheduling.
- [sched](https://github.com/romshark/sched) ![](https://img.shields.io/github/stars/romshark/sched.svg) - A job scheduler with the ability to fast-forward time.
- [scheduler](https://github.com/carlescere/scheduler) ![](https://img.shields.io/github/stars/carlescere/scheduler.svg) - Cronjobs scheduling made easy.
- [tasks](https://github.com/madflojo/tasks) ![](https://img.shields.io/github/stars/madflojo/tasks.svg) - An easy to use in-process scheduler for recurring tasks in Go.

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [ajson](https://github.com/spyzhov/ajson) ![](https://img.shields.io/github/stars/spyzhov/ajson.svg) - Abstract JSON for golang with JSONPath support.
- [ask](https://github.com/simonnilsson/ask) ![](https://img.shields.io/github/stars/simonnilsson/ask.svg) - Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types.
- [dynjson](https://github.com/cocoonspace/dynjson) ![](https://img.shields.io/github/stars/cocoonspace/dynjson.svg) - Client-customizable JSON formats for dynamic APIs.
- [ej](https://github.com/lucassscaravelli/ej) ![](https://img.shields.io/github/stars/lucassscaravelli/ej.svg) - Write and read JSON from different sources succinctly.
- [epoch](https://github.com/vtopc/epoch) ![](https://img.shields.io/github/stars/vtopc/epoch.svg) - Contains primitives for marshaling/unmarshalling Unix timestamp/epoch to/from build-in time.Time type in JSON.
- [fastjson](https://github.com/valyala/fastjson) ![](https://img.shields.io/github/stars/valyala/fastjson.svg) - Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection.
- [gabs](https://github.com/Jeffail/gabs) ![](https://img.shields.io/github/stars/Jeffail/gabs.svg) - For parsing, creating and editing unknown or dynamic JSON in Go.
- [gjo](https://github.com/skanehira/gjo) ![](https://img.shields.io/github/stars/skanehira/gjo.svg) - Small utility to create JSON objects.
- [GJSON](https://github.com/tidwall/gjson) ![](https://img.shields.io/github/stars/tidwall/gjson.svg) - Get a JSON value with one line of code.
- [go-jsonerror](https://github.com/ddymko/go-jsonerror) ![](https://img.shields.io/github/stars/ddymko/go-jsonerror.svg) - Go-JsonError is meant to allow us to easily create json response errors that follow the JsonApi spec.
- [go-respond](https://github.com/nicklaw5/go-respond) ![](https://img.shields.io/github/stars/nicklaw5/go-respond.svg) - Go package for handling common HTTP JSON responses.
- [gojmapr](https://github.com/limiu82214/gojmapr) ![](https://img.shields.io/github/stars/limiu82214/gojmapr.svg) - Get simple struct from complex json by json path.
- [gojq](https://github.com/elgs/gojq) ![](https://img.shields.io/github/stars/elgs/gojq.svg) - JSON query in Golang.
- [gojson](https://github.com/ChimeraCoder/gojson) ![](https://img.shields.io/github/stars/ChimeraCoder/gojson.svg) - Automatically generate Go (golang) struct definitions from example JSON.
- [htmljson](https://github.com/nikolaydubina/htmljson) ![](https://img.shields.io/github/stars/nikolaydubina/htmljson.svg) - Rich rendering of JSON as HTML in Go.
- [JayDiff](https://github.com/yazgazan/jaydiff) ![](https://img.shields.io/github/stars/yazgazan/jaydiff.svg) - JSON diff utility written in Go.
- [jettison](https://github.com/wI2L/jettison) ![](https://img.shields.io/github/stars/wI2L/jettison.svg) - Fast and flexible JSON encoder for Go.
- [jscan](https://github.com/romshark/jscan) ![](https://img.shields.io/github/stars/romshark/jscan.svg) - High performance zero-allocation JSON iterator.
- [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.
- [JSON-to-Proto](https://json-to-proto.github.io/) - Convert JSON to Protobuf online.
- [json2go](https://github.com/m-zajac/json2go) ![](https://img.shields.io/github/stars/m-zajac/json2go.svg) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.
- [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) ![](https://img.shields.io/github/stars/AmuzaTkts/jsonapi-errors.svg) - Go bindings based on the JSON API errors reference.
- [jsoncolor](https://github.com/neilotoole/jsoncolor) ![](https://img.shields.io/github/stars/neilotoole/jsoncolor.svg) - Drop-in replacement for `encoding/json` that outputs colorized JSON.
- [jsondiff](https://github.com/wI2L/jsondiff) ![](https://img.shields.io/github/stars/wI2L/jsondiff.svg) - JSON diff library for Go based on RFC6902 (JSON Patch).
- [jsonf](https://github.com/miolini/jsonf) ![](https://img.shields.io/github/stars/miolini/jsonf.svg) - Console tool for highlighted formatting and struct query fetching JSON.
- [jsongo](https://github.com/ricardolonga/jsongo) ![](https://img.shields.io/github/stars/ricardolonga/jsongo.svg) - Fluent API to make it easier to create Json objects.
- [jsonhal](https://github.com/RichardKnop/jsonhal) ![](https://img.shields.io/github/stars/RichardKnop/jsonhal.svg) - Simple Go package to make custom structs marshal into HAL compatible JSON responses.
- [jsonhandlers](https://github.com/abusomani/jsonhandlers) ![](https://img.shields.io/github/stars/abusomani/jsonhandlers.svg) - JSON library to expose simple handlers that lets you easily read and write json from various sources.
- [jsonic](https://github.com/sinhashubham95/jsonic) ![](https://img.shields.io/github/stars/sinhashubham95/jsonic.svg) - Utilities to handle and query JSON without defining structs in a type safe manner.
- [jsonvalue](https://github.com/Andrew-M-C/go.jsonvalue) ![](https://img.shields.io/github/stars/Andrew-M-C/go.jsonvalue.svg) - A fast and convenient library for unstructured JSON data, replacing `encoding/json`.
- [jzon](https://github.com/zerosnake0/jzon) ![](https://img.shields.io/github/stars/zerosnake0/jzon.svg) - JSON library with standard compatible API/behavior.
- [kazaam](https://github.com/Qntfy/kazaam) ![](https://img.shields.io/github/stars/Qntfy/kazaam.svg) - API for arbitrary transformation of JSON documents.
- [mapslice-json](https://github.com/mickep76/mapslice-json) ![](https://img.shields.io/github/stars/mickep76/mapslice-json.svg) - Go MapSlice for ordered marshal/ unmarshal of maps in JSON.
- [marshmallow](https://github.com/PerimeterX/marshmallow) ![](https://img.shields.io/github/stars/PerimeterX/marshmallow.svg) - Performant JSON unmarshalling for flexible use cases.
- [mp](https://github.com/sanbornm/mp) ![](https://img.shields.io/github/stars/sanbornm/mp.svg) - Simple cli email parser. It currently takes stdin and outputs JSON.
- [OjG](https://github.com/ohler55/ojg) ![](https://img.shields.io/github/stars/ohler55/ojg.svg) - Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath.
- [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) ![](https://img.shields.io/github/stars/dedalqq/omg.jsonparser.svg) - Simple JSON parser with validation by condition via golang struct fields tags.
- [SJSON](https://github.com/tidwall/sjson) ![](https://img.shields.io/github/stars/tidwall/sjson.svg) - Set a JSON value with one line of code.  
- [ujson](https://github.com/olvrng/ujson) ![](https://img.shields.io/github/stars/olvrng/ujson.svg) - Fast and minimal JSON parser and transformer that works on unstructured JSON.
- [vjson](https://github.com/miladibra10/vjson) ![](https://img.shields.io/github/stars/miladibra10/vjson.svg) - Go package for validating JSON objects with declaring a JSON schema with fluent API.

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [distillog](https://github.com/amoghe/distillog) ![](https://img.shields.io/github/stars/amoghe/distillog.svg) - distilled levelled logging (think of it as stdlib + log levels).
- [glg](https://github.com/kpango/glg) ![](https://img.shields.io/github/stars/kpango/glg.svg) - glg is simple and fast leveled logging library for Go.
- [glo](https://github.com/lajosbencz/glo) ![](https://img.shields.io/github/stars/lajosbencz/glo.svg) - PHP Monolog inspired logging facility with identical severity levels.
- [glog](https://github.com/golang/glog) ![](https://img.shields.io/github/stars/golang/glog.svg) - Leveled execution logs for Go.
- [go-cronowriter](https://github.com/utahta/go-cronowriter) ![](https://img.shields.io/github/stars/utahta/go-cronowriter.svg) - Simple writer that rotate log files automatically based on current date and time, like cronolog.
- [go-log](https://github.com/pieterclaerhout/go-log) ![](https://img.shields.io/github/stars/pieterclaerhout/go-log.svg) - A logging library with stack traces, object dumping and optional timestamps.
- [go-log](https://github.com/subchen/go-log) ![](https://img.shields.io/github/stars/subchen/go-log.svg) - Simple and configurable Logging in Go, with level, formatters and writers.
- [go-log](https://github.com/siddontang/go-log) ![](https://img.shields.io/github/stars/siddontang/go-log.svg) - Log lib supports level and multi handlers.
- [go-log](https://github.com/ian-kent/go-log) ![](https://img.shields.io/github/stars/ian-kent/go-log.svg) - Log4j implementation in Go.
- [go-logger](https://github.com/apsdehal/go-logger) ![](https://img.shields.io/github/stars/apsdehal/go-logger.svg) - Simple logger of Go Programs, with level handlers.
- [gone/log](https://github.com/One-com/gone/tree/master/log) ![](https://img.shields.io/github/stars/One-com/gone.svg) - Fast, extendable, full-featured, std-lib source compatible log library.
- [httpretty](https://github.com/henvic/httpretty) ![](https://img.shields.io/github/stars/henvic/httpretty.svg) - Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest).
- [journald](https://github.com/ssgreg/journald) ![](https://img.shields.io/github/stars/ssgreg/journald.svg) - Go implementation of systemd Journal's native API for logging.
- [kemba](https://github.com/clok/kemba) ![](https://img.shields.io/github/stars/clok/kemba.svg) - A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications.
- [lazyjournal](https://github.com/Lifailon/lazyjournal) ![](https://img.shields.io/github/stars/Lifailon/lazyjournal.svg) - A TUI for reading and filtering logs from journalctl, file system, Docker and Podman containers, as well Kubernetes pods.
- [log](https://github.com/aerogo/log) ![](https://img.shields.io/github/stars/aerogo/log.svg) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).
- [log](https://github.com/apex/log) ![](https://img.shields.io/github/stars/apex/log.svg) - Structured logging package for Go.
- [log](https://github.com/go-playground/log) ![](https://img.shields.io/github/stars/go-playground/log.svg) - Simple, configurable and scalable Structured Logging for Go.
- [log](https://github.com/teris-io/log) ![](https://img.shields.io/github/stars/teris-io/log.svg) - Structured log interface for Go cleanly separates logging facade from its implementation.
- [log](https://github.com/heartwilltell/log) ![](https://img.shields.io/github/stars/heartwilltell/log.svg) - Simple leveled logging wrapper around standard log package.
- [log](https://github.com/no-src/log) ![](https://img.shields.io/github/stars/no-src/log.svg) - A simple logging framework out of the box.
- [log15](https://github.com/inconshreveable/log15) ![](https://img.shields.io/github/stars/inconshreveable/log15.svg) - Simple, powerful logging for Go.
- [logdump](https://github.com/ewwwwwqm/logdump) ![](https://img.shields.io/github/stars/ewwwwwqm/logdump.svg) - Package for multi-level logging.
- [logex](https://github.com/chzyer/logex) ![](https://img.shields.io/github/stars/chzyer/logex.svg) - Golang log lib, supports tracking and level, wrap by standard log lib.
- [logger](https://github.com/azer/logger) ![](https://img.shields.io/github/stars/azer/logger.svg) - Minimalistic logging library for Go.
- [logo](https://github.com/mbndr/logo) ![](https://img.shields.io/github/stars/mbndr/logo.svg) - Golang logger to different configurable writers.
- [logrus](https://github.com/Sirupsen/logrus) ![](https://img.shields.io/github/stars/Sirupsen/logrus.svg) - Structured logger for Go.
- [logrusiowriter](https://github.com/cabify/logrusiowriter) ![](https://img.shields.io/github/stars/cabify/logrusiowriter.svg) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.
- [logrusly](https://github.com/sebest/logrusly) ![](https://img.shields.io/github/stars/sebest/logrusly.svg) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
- [logutils](https://github.com/hashicorp/logutils) ![](https://img.shields.io/github/stars/hashicorp/logutils.svg) - Utilities for slightly better logging in Go (Golang) extending the standard logger.
- [logxi](https://github.com/mgutz/logxi) ![](https://img.shields.io/github/stars/mgutz/logxi.svg) - 12-factor app logger that is fast and makes you happy.
- [lumberjack](https://github.com/natefinch/lumberjack) ![](https://img.shields.io/github/stars/natefinch/lumberjack.svg) - Simple rolling logger, implements io.WriteCloser.
- [mlog](https://github.com/jbrodriguez/mlog) ![](https://img.shields.io/github/stars/jbrodriguez/mlog.svg) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
- [noodlog](https://github.com/gyozatech/noodlog) ![](https://img.shields.io/github/stars/gyozatech/noodlog.svg) - Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings.
- [onelog](https://github.com/francoispqt/onelog) ![](https://img.shields.io/github/stars/francoispqt/onelog.svg) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation.
- [ozzo-log](https://github.com/go-ozzo/ozzo-log) ![](https://img.shields.io/github/stars/go-ozzo/ozzo-log.svg) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
- [phuslu/log](https://github.com/phuslu/log) ![](https://img.shields.io/github/stars/phuslu/log.svg) - High performance structured logging.
- [pp](https://github.com/k0kubun/pp) ![](https://img.shields.io/github/stars/k0kubun/pp.svg) - Colored pretty printer for Go language.
- [rollingwriter](https://github.com/arthurkiller/rollingWriter) ![](https://img.shields.io/github/stars/arthurkiller/rollingWriter.svg) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.
- [seelog](https://github.com/cihub/seelog) ![](https://img.shields.io/github/stars/cihub/seelog.svg) - Logging functionality with flexible dispatching, filtering, and formatting.
- [sentry-go](https://github.com/getsentry/sentry-go) ![](https://img.shields.io/github/stars/getsentry/sentry-go.svg) - Sentry SDK for Go. Helps monitor and track errors with real-time alerts and performance monitoring.
- [slf4g](https://github.com/echocat/slf4g) ![](https://img.shields.io/github/stars/echocat/slf4g.svg) - Simple Logging Facade for Golang: Simple structured logging; but powerful, extendable and customizable, with huge amount of learnings from decades of past logging frameworks.
- [slog](https://github.com/gookit/slog) ![](https://img.shields.io/github/stars/gookit/slog.svg) - Lightweight, configurable, extensible logger for Go.
- [slog-formatter](https://github.com/samber/slog-formatter) ![](https://img.shields.io/github/stars/samber/slog-formatter.svg) - Common formatters for slog and helpers to build your own.
- [slog-multi](https://github.com/samber/slog-multi) ![](https://img.shields.io/github/stars/samber/slog-multi.svg) - Chain of slog.Handler (pipeline, fanout...).
- [slogor](https://gitlab.com/greyxor/slogor) - A colorful slog handler.
- [spew](https://github.com/davecgh/go-spew) ![](https://img.shields.io/github/stars/davecgh/go-spew.svg) - Implements a deep pretty printer for Go data structures to aid in debugging.
- [sqldb-logger](https://github.com/simukti/sqldb-logger) ![](https://img.shields.io/github/stars/simukti/sqldb-logger.svg) - A logger for Go SQL database driver without modify existing \*sql.DB stdlib usage.
- [stdlog](https://github.com/alexcesaro/log) ![](https://img.shields.io/github/stars/alexcesaro/log.svg) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
- [structy/log](https://github.com/structy/log) ![](https://img.shields.io/github/stars/structy/log.svg) - A simple to use log system, minimalist but with features for debugging and differentiation of messages.
- [tail](https://github.com/hpcloud/tail) ![](https://img.shields.io/github/stars/hpcloud/tail.svg) - Go package striving to emulate the features of the BSD tail program.
- [timberjack](https://github.com/DeRuina/timberjack) ![](https://img.shields.io/github/stars/DeRuina/timberjack.svg) - Rolling logger with size-based, time-based, and scheduled clock-based rotation, supporting compression and cleanup.
- [tint](https://github.com/lmittmann/tint) ![](https://img.shields.io/github/stars/lmittmann/tint.svg) - A slog.Handler that writes tinted logs.
- [xlog](https://github.com/xfxdev/xlog) ![](https://img.shields.io/github/stars/xfxdev/xlog.svg) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
- [xlog](https://github.com/rs/xlog) ![](https://img.shields.io/github/stars/rs/xlog.svg) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching.
- [xylog](https://github.com/xybor-x/xylog) ![](https://img.shields.io/github/stars/xybor-x/xylog.svg) - Leveled and structured logging, dynamic fields, high performance, zone management, simple configuration, and readable syntax.
- [yell](https://github.com/jfcg/yell) ![](https://img.shields.io/github/stars/jfcg/yell.svg) - Yet another minimalistic logging library.
- [zap](https://github.com/uber-go/zap) ![](https://img.shields.io/github/stars/uber-go/zap.svg) - Fast, structured, leveled logging in Go.
- [zax](https://github.com/yuseferi/zax) ![](https://img.shields.io/github/stars/yuseferi/zax.svg) - Integrate Context with Zap logger, which leads to more flexibility in Go logging.
- [zerolog](https://github.com/rs/zerolog) ![](https://img.shields.io/github/stars/rs/zerolog.svg) - Zero-allocation JSON logger.
- [zkits-logger](https://github.com/edoger/zkits-logger) ![](https://img.shields.io/github/stars/edoger/zkits-logger.svg) - A powerful zero-dependency JSON logger.
- [zl](https://github.com/nkmr-jp/zl) ![](https://img.shields.io/github/stars/nkmr-jp/zl.svg) - High Developer Experience, zap based logger. It offers rich functionality but is easy to configure.

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [bayesian](https://github.com/jbrukh/bayesian) ![](https://img.shields.io/github/stars/jbrukh/bayesian.svg) - Naive Bayesian Classification for Golang.
- [catboost-cgo](https://github.com/mirecl/catboost-cgo) ![](https://img.shields.io/github/stars/mirecl/catboost-cgo.svg) - Fast, scalable, high performance Gradient Boosting on Decision Trees library. Golang using Cgo for blazing fast inference CatBoost Model.
- [CloudForest](https://github.com/ryanbressler/CloudForest) ![](https://img.shields.io/github/stars/ryanbressler/CloudForest.svg) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
- [ddt](https://github.com/sgrodriguez/ddt) ![](https://img.shields.io/github/stars/sgrodriguez/ddt.svg) - Dynamic decision tree, create trees defining customizable rules.
- [eaopt](https://github.com/MaxHalford/eaopt) ![](https://img.shields.io/github/stars/MaxHalford/eaopt.svg) - An evolutionary optimization library.
- [evoli](https://github.com/khezen/evoli) ![](https://img.shields.io/github/stars/khezen/evoli.svg) - Genetic Algorithm and Particle Swarm Optimization library.
- [fonet](https://github.com/Fontinalis/fonet) ![](https://img.shields.io/github/stars/Fontinalis/fonet.svg) - A Deep Neural Network library written in Go.
- [go-cluster](https://github.com/e-XpertSolutions/go-cluster) ![](https://img.shields.io/github/stars/e-XpertSolutions/go-cluster.svg) - Go implementation of the k-modes and k-prototypes clustering algorithms.
- [go-deep](https://github.com/patrikeh/go-deep) ![](https://img.shields.io/github/stars/patrikeh/go-deep.svg) - A feature-rich neural network library in Go.
- [go-fann](https://github.com/white-pony/go-fann) ![](https://img.shields.io/github/stars/white-pony/go-fann.svg) - Go bindings for Fast Artificial Neural Networks(FANN) library.
- [go-galib](https://github.com/thoj/go-galib) ![](https://img.shields.io/github/stars/thoj/go-galib.svg) - Genetic Algorithms library written in Go / golang.
- [go-pr](https://github.com/daviddengcn/go-pr) ![](https://img.shields.io/github/stars/daviddengcn/go-pr.svg) - Pattern recognition package in Go lang.
- [gobrain](https://github.com/goml/gobrain) ![](https://img.shields.io/github/stars/goml/gobrain.svg) - Neural Networks written in go.
- [godist](https://github.com/e-dard/godist) ![](https://img.shields.io/github/stars/e-dard/godist.svg) - Various probability distributions, and associated methods.
- [goga](https://github.com/tomcraven/goga) ![](https://img.shields.io/github/stars/tomcraven/goga.svg) - Genetic algorithm library for Go.
- [GoLearn](https://github.com/sjwhitworth/golearn) ![](https://img.shields.io/github/stars/sjwhitworth/golearn.svg) - General Machine Learning library for Go.
- [GoMind](https://github.com/surenderthakran/gomind) ![](https://img.shields.io/github/stars/surenderthakran/gomind.svg) - A simplistic Neural Network Library in Go.
- [goml](https://github.com/cdipaolo/goml) ![](https://img.shields.io/github/stars/cdipaolo/goml.svg) - On-line Machine Learning in Go.
- [GoMLX](https://github.com/gomlx/gomlx) ![](https://img.shields.io/github/stars/gomlx/gomlx.svg) - An accelerated Machine Learning framework for Go.
- [gonet](https://github.com/dathoangnd/gonet) ![](https://img.shields.io/github/stars/dathoangnd/gonet.svg) - Neural Network for Go.
- [Goptuna](https://github.com/c-bata/goptuna) ![](https://img.shields.io/github/stars/c-bata/goptuna.svg) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.
- [goRecommend](https://github.com/timkaye11/goRecommend) ![](https://img.shields.io/github/stars/timkaye11/goRecommend.svg) - Recommendation Algorithms library written in Go.
- [gorgonia](https://github.com/gorgonia/gorgonia) ![](https://img.shields.io/github/stars/gorgonia/gorgonia.svg) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.
- [gorse](https://github.com/zhenghaoz/gorse) ![](https://img.shields.io/github/stars/zhenghaoz/gorse.svg) - An offline recommender system backend based on collaborative filtering written in Go.
- [goscore](https://github.com/asafschers/goscore) ![](https://img.shields.io/github/stars/asafschers/goscore.svg) - Go Scoring API for PMML.
- [gosseract](https://github.com/otiai10/gosseract) ![](https://img.shields.io/github/stars/otiai10/gosseract.svg) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.
- [hugot](https://github.com/knights-analytics/hugot) ![](https://img.shields.io/github/stars/knights-analytics/hugot.svg) - Huggingface transformer pipelines for golang with onnxruntime.
- [libsvm](https://github.com/datastream/libsvm) ![](https://img.shields.io/github/stars/datastream/libsvm.svg) - libsvm golang version derived work based on LIBSVM 3.14.
- [m2cgen](https://github.com/BayesWitnesses/m2cgen) ![](https://img.shields.io/github/stars/BayesWitnesses/m2cgen.svg) - A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support.
- [neural-go](https://github.com/schuyler/neural-go) ![](https://img.shields.io/github/stars/schuyler/neural-go.svg) - Multilayer perceptron network implemented in Go, with training via backpropagation.
- [ocrserver](https://github.com/otiai10/ocrserver) ![](https://img.shields.io/github/stars/otiai10/ocrserver.svg) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku.
- [onnx-go](https://github.com/owulveryck/onnx-go) ![](https://img.shields.io/github/stars/owulveryck/onnx-go.svg) - Go Interface to Open Neural Network Exchange (ONNX).
- [probab](https://github.com/ThePaw/probab) ![](https://img.shields.io/github/stars/ThePaw/probab.svg) - Probability distribution functions. Bayesian inference. Written in pure Go.
- [randomforest](https://github.com/malaschitz/randomForest) ![](https://img.shields.io/github/stars/malaschitz/randomForest.svg) - Easy to use Random Forest library for Go.
- [regommend](https://github.com/muesli/regommend) ![](https://img.shields.io/github/stars/muesli/regommend.svg) - Recommendation & collaborative filtering engine.
- [shield](https://github.com/eaigner/shield) ![](https://img.shields.io/github/stars/eaigner/shield.svg) - Bayesian text classifier with flexible tokenizers and storage backends for Go.
- [tfgo](https://github.com/galeone/tfgo) ![](https://img.shields.io/github/stars/galeone/tfgo.svg) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.
- [Varis](https://github.com/Xamber/Varis) ![](https://img.shields.io/github/stars/Xamber/Varis.svg) - Golang Neural Network.

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [ami](https://github.com/kak-tus/ami) ![](https://img.shields.io/github/stars/kak-tus/ami.svg) - Go client to reliable queues based on Redis Cluster Streams.
- [amqp](https://github.com/rabbitmq/amqp091-go) ![](https://img.shields.io/github/stars/rabbitmq/amqp091-go.svg) - Go RabbitMQ Client Library.
- [APNs2](https://github.com/sideshow/apns2) ![](https://img.shields.io/github/stars/sideshow/apns2.svg) - HTTP/2 Apple Push Notification provider for Go - Send push notifications to iOS, tvOS, Safari and OSX apps.
- [Asynq](https://github.com/hibiken/asynq) ![](https://img.shields.io/github/stars/hibiken/asynq.svg) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis.
- [backlite](https://github.com/mikestefanello/backlite) ![](https://img.shields.io/github/stars/mikestefanello/backlite.svg) - Type-safe, persistent, embedded task queues and background job runner w/ SQLite.
- [Beaver](https://github.com/Clivern/Beaver) ![](https://img.shields.io/github/stars/Clivern/Beaver.svg) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.
- [Bus](https://github.com/mustafaturan/bus) ![](https://img.shields.io/github/stars/mustafaturan/bus.svg) - Minimalist message bus implementation for internal communication.
- [Centrifugo](https://github.com/centrifugal/centrifugo) ![](https://img.shields.io/github/stars/centrifugal/centrifugo.svg) - Real-time messaging (Websockets or SockJS) server in Go.
- [Chanify](https://github.com/chanify/chanify) ![](https://img.shields.io/github/stars/chanify/chanify.svg) - A push notification server send message to your iOS devices.
- [Commander](https://github.com/jeroenrinzema/commander) ![](https://img.shields.io/github/stars/jeroenrinzema/commander.svg) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.
- [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) ![](https://img.shields.io/github/stars/confluentinc/confluent-kafka-go.svg) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform.
- [dbus](https://github.com/godbus/dbus) ![](https://img.shields.io/github/stars/godbus/dbus.svg) - Native Go bindings for D-Bus.
- [drone-line](https://github.com/appleboy/drone-line) ![](https://img.shields.io/github/stars/appleboy/drone-line.svg) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.
- [emitter](https://github.com/olebedev/emitter) ![](https://img.shields.io/github/stars/olebedev/emitter.svg) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
- [event](https://github.com/agoalofalife/event) ![](https://img.shields.io/github/stars/agoalofalife/event.svg) - Implementation of the pattern observer.
- [EventBus](https://github.com/asaskevich/EventBus) ![](https://img.shields.io/github/stars/asaskevich/EventBus.svg) - The lightweight event bus with async compatibility.
- [gaurun-client](https://github.com/osamingo/gaurun-client) ![](https://img.shields.io/github/stars/osamingo/gaurun-client.svg) - Gaurun Client written in Go.
- [Glue](https://github.com/desertbit/glue) ![](https://img.shields.io/github/stars/desertbit/glue.svg) - Robust Go and Javascript Socket Library (Alternative to Socket.io).
- [go-eventbus](https://github.com/stanipetrosyan/go-eventbus) ![](https://img.shields.io/github/stars/stanipetrosyan/go-eventbus.svg) - Simple Event Bus package for Go.
- [Go-MediatR](https://github.com/mehdihadeli/Go-MediatR) ![](https://img.shields.io/github/stars/mehdihadeli/Go-MediatR.svg) - A library for handling mediator patterns and simplified CQRS patterns within an event-driven architecture, inspired by csharp MediatR library.
- [go-mq](https://github.com/cheshir/go-mq) ![](https://img.shields.io/github/stars/cheshir/go-mq.svg) - RabbitMQ client with declarative configuration.
- [go-notify](https://github.com/TheCreeper/go-notify) ![](https://img.shields.io/github/stars/TheCreeper/go-notify.svg) - Native implementation of the freedesktop notification spec.
- [go-nsq](https://github.com/nsqio/go-nsq) ![](https://img.shields.io/github/stars/nsqio/go-nsq.svg) - the official Go package for NSQ.
- [go-res](https://github.com/jirenius/go-res) ![](https://img.shields.io/github/stars/jirenius/go-res.svg) - Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate.
- [go-vitotrol](https://github.com/maxatome/go-vitotrol) ![](https://img.shields.io/github/stars/maxatome/go-vitotrol.svg) - Client library to Viessmann Vitotrol web service.
- [GoEventBus](https://github.com/Raezil/GoEventBus) ![](https://img.shields.io/github/stars/Raezil/GoEventBus.svg) - A blazing‑fast, in‑memory, lock‑free event bus library
- [Gollum](https://github.com/trivago/gollum) ![](https://img.shields.io/github/stars/trivago/gollum.svg) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.
- [golongpoll](https://github.com/jcuga/golongpoll) ![](https://img.shields.io/github/stars/jcuga/golongpoll.svg) - HTTP longpoll server library that makes web pub-sub simple.
- [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) ![](https://img.shields.io/github/stars/Terry-Mao/gopush-cluster.svg) - gopush-cluster is a go push server cluster.
- [gorush](https://github.com/appleboy/gorush) ![](https://img.shields.io/github/stars/appleboy/gorush.svg) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).
- [gosd](https://github.com/alexsniffin/gosd) ![](https://img.shields.io/github/stars/alexsniffin/gosd.svg) - A library for scheduling when to dispatch a message to a channel.
- [guble](https://github.com/smancke/guble) ![](https://img.shields.io/github/stars/smancke/guble.svg) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.
- [hare](https://github.com/leozz37/hare) ![](https://img.shields.io/github/stars/leozz37/hare.svg) - A user friendly library for sending messages and listening to TCP sockets.
- [hub](https://github.com/leandro-lugaresi/hub) ![](https://img.shields.io/github/stars/leandro-lugaresi/hub.svg) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.
- [hypermatch](https://github.com/SchwarzIT/hypermatch) ![](https://img.shields.io/github/stars/SchwarzIT/hypermatch.svg) - A very fast and efficient Go library for matching events to a large set of rules
- [jazz](https://github.com/socifi/jazz) ![](https://img.shields.io/github/stars/socifi/jazz.svg) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.
- [machinery](https://github.com/RichardKnop/machinery) ![](https://img.shields.io/github/stars/RichardKnop/machinery.svg) - Asynchronous task queue/job queue based on distributed message passing.
- [mangos](https://github.com/nanomsg/mangos) ![](https://img.shields.io/github/stars/nanomsg/mangos.svg) - Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability.
- [melody](https://github.com/olahol/melody) ![](https://img.shields.io/github/stars/olahol/melody.svg) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.
- [Mercure](https://github.com/dunglas/mercure) ![](https://img.shields.io/github/stars/dunglas/mercure.svg) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).
- [messagebus](https://github.com/vardius/message-bus) ![](https://img.shields.io/github/stars/vardius/message-bus.svg) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
- [NATS Go Client](https://github.com/nats-io/nats.go) ![](https://img.shields.io/github/stars/nats-io/nats.go.svg) - Go client for the NATS
  messaging system.
- [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) ![](https://img.shields.io/github/stars/rafaeljesus/nsq-event-bus.svg) - A tiny wrapper around NSQ topic and channel.
- [oplog](https://github.com/dailymotion/oplog) ![](https://img.shields.io/github/stars/dailymotion/oplog.svg) - Generic oplog/replication system for REST APIs.
- [pubsub](https://github.com/tuxychandru/pubsub) ![](https://img.shields.io/github/stars/tuxychandru/pubsub.svg) - Simple pubsub package for go.
- [Quamina](https://github.com/timbray/quamina) ![](https://img.shields.io/github/stars/timbray/quamina.svg) - Fast pattern-matching for filtering messages and events.
- [rabbitroutine](https://github.com/furdarius/rabbitroutine) ![](https://img.shields.io/github/stars/furdarius/rabbitroutine.svg) - Lightweight library that handles RabbitMQ auto-reconnect and publishing retries. The library takes into account the need to re-declare entities in RabbitMQ after reconnection.
- [rabbus](https://github.com/rafaeljesus/rabbus) ![](https://img.shields.io/github/stars/rafaeljesus/rabbus.svg) - A tiny wrapper over amqp exchanges and queues.
- [rabtap](https://github.com/jandelgado/rabtap) ![](https://img.shields.io/github/stars/jandelgado/rabtap.svg) - RabbitMQ swiss army knife cli app.
- [RapidMQ](https://github.com/sybrexsys/RapidMQ) ![](https://img.shields.io/github/stars/sybrexsys/RapidMQ.svg) - RapidMQ is a lightweight and reliable library for managing of the local messages queue.
- [Ratus](https://github.com/hyperonym/ratus) ![](https://img.shields.io/github/stars/hyperonym/ratus.svg) - Ratus is a RESTful asynchronous task queue server.
- [redisqueue](https://github.com/robinjoseph08/redisqueue) ![](https://img.shields.io/github/stars/robinjoseph08/redisqueue.svg) - redisqueue provides a producer and consumer of a queue that uses Redis streams.
- [rmqconn](https://github.com/sbabiv/rmqconn) ![](https://img.shields.io/github/stars/sbabiv/rmqconn.svg) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.
- [sarama](https://github.com/Shopify/sarama) ![](https://img.shields.io/github/stars/Shopify/sarama.svg) - Go library for Apache Kafka.
- [Uniqush-Push](https://github.com/uniqush/uniqush-push) ![](https://img.shields.io/github/stars/uniqush/uniqush-push.svg) - Redis backed unified push service for server-side notifications to mobile devices.
- [varmq](https://github.com/goptics/varmq) ![](https://img.shields.io/github/stars/goptics/varmq.svg) - A storage-agnostic message queue and worker pool for concurrent Go programs.
- [Watermill](https://github.com/ThreeDotsLabs/watermill) ![](https://img.shields.io/github/stars/ThreeDotsLabs/watermill.svg) - Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog.
- [zmq4](https://github.com/pebbe/zmq4) ![](https://img.shields.io/github/stars/pebbe/zmq4.svg) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice](https://github.com/unidoc/unioffice) ![](https://img.shields.io/github/stars/unidoc/unioffice.svg) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [cellwalker](https://github.com/chonla/cellwalker) ![](https://img.shields.io/github/stars/chonla/cellwalker.svg) - Virtually traverse Excel cell by cell's name.
- [excelize](https://github.com/xuri/excelize) ![](https://img.shields.io/github/stars/xuri/excelize.svg) - Golang library for reading and writing Microsoft Excel&trade; (XLSX) files.
- [exl](https://github.com/go-the-way/exl) ![](https://img.shields.io/github/stars/go-the-way/exl.svg) - Excel binding to struct written in Go.(Only supports Go1.18+)
- [go-excel](https://github.com/szyhf/go-excel) ![](https://img.shields.io/github/stars/szyhf/go-excel.svg) - A simple and light reader to read a relate-db-like excel as a table.
- [xlsx](https://github.com/tealeg/xlsx) ![](https://img.shields.io/github/stars/tealeg/xlsx.svg) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.
- [xlsx](https://github.com/plandem/xlsx) ![](https://img.shields.io/github/stars/plandem/xlsx.svg) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs.

### Microsoft Word

_Libraries for working with Microsoft Word._

- [godocx](https://github.com/gomutex/godocx) ![](https://img.shields.io/github/stars/gomutex/godocx.svg) - Library for reading and writing Microsoft Word (Docx) files.

**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [alice](https://github.com/magic003/alice) ![](https://img.shields.io/github/stars/magic003/alice.svg) - Additive dependency injection container for Golang.
- [autowire](https://github.com/tiendc/autowire) ![](https://img.shields.io/github/stars/tiendc/autowire.svg) - Dependency injection using Generics and reflection.
- [boot-go](http://github.com/boot-go/boot) - Component-based development with dependency injection using reflections for Go developers.
- [componego](https://github.com/componego/componego) ![](https://img.shields.io/github/stars/componego/componego.svg) - A dependency injection framework based on components, allowing dynamic dependency replacement without duplicating code in tests.
- [cosban/di](https://gitlab.com/cosban/di) - A code generation based dependency injection wiring tool.
- [di](https://github.com/goava/di) ![](https://img.shields.io/github/stars/goava/di.svg) - A dependency injection container for go programming language.
- [dig](https://github.com/uber-go/dig) ![](https://img.shields.io/github/stars/uber-go/dig.svg) - A reflection based dependency injection toolkit for Go.
- [dingo](https://github.com/i-love-flamingo/dingo) ![](https://img.shields.io/github/stars/i-love-flamingo/dingo.svg) - A dependency injection toolkit for Go, based on Guice.
- [do](https://github.com/samber/do) ![](https://img.shields.io/github/stars/samber/do.svg) - A dependency injection framework based on Generics.
- [fx](https://github.com/uber-go/fx) ![](https://img.shields.io/github/stars/uber-go/fx.svg) - A dependency injection based application framework for Go (built on top of dig).
- [Go-Spring](https://github.com/go-spring/spring-core) ![](https://img.shields.io/github/stars/go-spring/spring-core.svg) - A high-performance Go framework inspired by Spring Boot, offering DI, auto-configuration, and lifecycle management while maintaining Go's simplicity and efficiency.
- [gocontainer](https://github.com/vardius/gocontainer) ![](https://img.shields.io/github/stars/vardius/gocontainer.svg) - Simple Dependency Injection Container.
- [godi](https://github.com/junioryono/godi) ![](https://img.shields.io/github/stars/junioryono/godi.svg) - Microsoft-style dependency injection for Go with scoped lifetimes and generics.
- [goioc/di](https://github.com/goioc/di) ![](https://img.shields.io/github/stars/goioc/di.svg) - Spring-inspired Dependency Injection Container.
- [GoLobby/Container](https://github.com/golobby/container) ![](https://img.shields.io/github/stars/golobby/container.svg) - GoLobby Container is a lightweight yet powerful IoC dependency injection container for the Go programming language.
- [gontainer](https://github.com/NVIDIA/gontainer) ![](https://img.shields.io/github/stars/NVIDIA/gontainer.svg) - A dependency injection service container for Go projects.
- [gontainer/gontainer](https://github.com/gontainer/gontainer) ![](https://img.shields.io/github/stars/gontainer/gontainer.svg) - A YAML-based Dependency Injection container for GO. It supports dependencies' scopes, and auto-detection of circular dependencies. Gontainer is concurrent-safe.
- [HnH/di](https://github.com/HnH/di) ![](https://img.shields.io/github/stars/HnH/di.svg) - DI container library that is focused on clean API and flexibility.
- [kinit](https://github.com/go-kata/kinit) ![](https://img.shields.io/github/stars/go-kata/kinit.svg) - Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization.
- [kod](https://github.com/go-kod/kod) ![](https://img.shields.io/github/stars/go-kod/kod.svg) - A generics based dependency injection framework for Go.
- [linker](https://github.com/logrange/linker) ![](https://img.shields.io/github/stars/logrange/linker.svg) - A reflection based dependency injection and inversion of control library with components lifecycle support.
- [nject](https://github.com/muir/nject) ![](https://img.shields.io/github/stars/muir/nject.svg) - A type safe, reflective framework for libraries, tests, http endpoints, and service startup.
- [ore](https://github.com/firasdarwish/ore) ![](https://img.shields.io/github/stars/firasdarwish/ore.svg) - Lightweight, generic & simple dependency injection (DI) container.
- [parsley](https://github.com/matzefriedrich/parsley) ![](https://img.shields.io/github/stars/matzefriedrich/parsley.svg) - A flexible and modular reflection-based DI library with advanced features like scoped contexts and proxy generation, designed for large-scale Go applications.
- [wire](https://github.com/Fs02/wire) ![](https://img.shields.io/github/stars/Fs02/wire.svg) - Strict Runtime Dependency Injection for Golang.

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [ardanlabs/service](https://github.com/ardanlabs/service) ![](https://img.shields.io/github/stars/ardanlabs/service.svg) - A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications.
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) ![](https://img.shields.io/github/stars/lacion/cookiecutter-golang.svg) - A Go application boilerplate template for quick starting projects following production best practices.
- [go-blueprint](https://github.com/Melkeydev/go-blueprint) ![](https://img.shields.io/github/stars/Melkeydev/go-blueprint.svg) - Allows users to spin up a quick Go project using a popular framework.
- [go-module](https://github.com/octomation/go-module) ![](https://img.shields.io/github/stars/octomation/go-module.svg) - Template for a typical module written on Go.
- [go-sample](https://github.com/zitryss/go-sample) ![](https://img.shields.io/github/stars/zitryss/go-sample.svg) - A sample layout for Go application projects with the real code.
- [go-starter](https://github.com/allaboutapps/go-starter) ![](https://img.shields.io/github/stars/allaboutapps/go-starter.svg) - An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers.
- [go-todo-backend](https://github.com/Fs02/go-todo-backend) ![](https://img.shields.io/github/stars/Fs02/go-todo-backend.svg) - Go Todo Backend example using modular project layout for product microservice.
- [goapp](https://github.com/naughtygopher/goapp) ![](https://img.shields.io/github/stars/naughtygopher/goapp.svg) - An opinionated guideline to structure & develop a Go web application/service.
- [gobase](https://github.com/wajox/gobase) ![](https://img.shields.io/github/stars/wajox/gobase.svg) - A simple skeleton for golang application with basic setup for real golang application.
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) ![](https://img.shields.io/github/stars/golang-standards/project-layout.svg) - Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful.
- [golang-templates/seed](https://github.com/golang-templates/seed) ![](https://img.shields.io/github/stars/golang-templates/seed.svg) - Go application GitHub repository template.
- [goxygen](https://github.com/shpota/goxygen) ![](https://img.shields.io/github/stars/shpota/goxygen.svg) - Generate a modern Web project with Go and Angular, React, or Vue in seconds.
- [insidieux/inizio](https://github.com/insidieux/inizio) ![](https://img.shields.io/github/stars/insidieux/inizio.svg) - Golang project layout generator with plugins.
- [kickstart.go](https://github.com/raeperd/kickstart.go) ![](https://img.shields.io/github/stars/raeperd/kickstart.go.svg) - Minimalistic single-file Go HTTP server template without third-party dependencies.
- [modern-go-application](https://github.com/sagikazarmark/modern-go-application) ![](https://img.shields.io/github/stars/sagikazarmark/modern-go-application.svg) - Go application boilerplate and example applying modern practices.
- [nunu](https://github.com/go-nunu/nunu) ![](https://img.shields.io/github/stars/go-nunu/nunu.svg) - Nunu is a scaffolding tool for building Go applications.
- [pagoda](https://github.com/mikestefanello/pagoda) ![](https://img.shields.io/github/stars/mikestefanello/pagoda.svg) - Rapid, easy full-stack web development starter kit built in Go.
- [scaffold](https://github.com/catchplay/scaffold) ![](https://img.shields.io/github/stars/catchplay/scaffold.svg) - Scaffold generates a starter Go project layout. Lets you focus on business logic implemented.
- [wangyoucao577/go-project-layout](https://github.com/wangyoucao577/go-project-layout) ![](https://img.shields.io/github/stars/wangyoucao577/go-project-layout.svg) - Set of practices and discussions on how to structure Go project layout.

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [bexp](https://github.com/happy-sdk/happy/tree/main/pkg/strings/bexp) ![](https://img.shields.io/github/stars/happy-sdk/happy.svg) - Go implementation of Brace Expansion mechanism to generate arbitrary strings.
- [caps](https://github.com/chanced/caps) ![](https://img.shields.io/github/stars/chanced/caps.svg) - A case conversion library.
- [go-formatter](https://gitlab.com/tymonx/go-formatter) - Implements **replacement fields** surrounded by curly braces `{}` format strings.
- [gobeam/Stringy](https://github.com/gobeam/Stringy) ![](https://img.shields.io/github/stars/gobeam/Stringy.svg) - String manipulation library to convert string to camel case, snake case, kebab case / slugify etc.
- [strcase](https://github.com/charlievieth/strcase) ![](https://img.shields.io/github/stars/charlievieth/strcase.svg) - Case-insensitive implementation of the standard library's strings/bytes packages.
- [strutil](https://github.com/ozgio/strutil) ![](https://img.shields.io/github/stars/ozgio/strutil.svg) - String utilities.
- [sttr](https://github.com/abhimanyu003/sttr) ![](https://img.shields.io/github/stars/abhimanyu003/sttr.svg) - cross-platform, cli app to perform various operations on string.
- [xstrings](https://github.com/huandu/xstrings) ![](https://img.shields.io/github/stars/huandu/xstrings.svg) - Collection of useful string functions ported from other languages.

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [anagent](https://github.com/mudler/anagent) ![](https://img.shields.io/github/stars/mudler/anagent.svg) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.
- [antch](https://github.com/antchfx/antch) ![](https://img.shields.io/github/stars/antchfx/antch.svg) - A fast, powerful and extensible web crawling & scraping framework.
- [archives](https://github.com/mholt/archives) ![](https://img.shields.io/github/stars/mholt/archives.svg) - a cross-platform, multi-format Go library for working with archives and compression formats with a unified API and as virtual file systems compatible with io/fs.
- [autoflags](https://github.com/artyom/autoflags) ![](https://img.shields.io/github/stars/artyom/autoflags.svg) - Go package to automatically define command line flags from struct fields.
- [avgRating](https://github.com/kirillDanshin/avgRating) ![](https://img.shields.io/github/stars/kirillDanshin/avgRating.svg) - Calculate average score and rating based on Wilson Score Equation.
- [banner](https://github.com/dimiro1/banner) ![](https://img.shields.io/github/stars/dimiro1/banner.svg) - Add beautiful banners into your Go applications.
- [base64Captcha](https://github.com/mojocn/base64Captcha) ![](https://img.shields.io/github/stars/mojocn/base64Captcha.svg) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.
- [basexx](https://github.com/bobg/basexx) ![](https://img.shields.io/github/stars/bobg/basexx.svg) - Convert to, from, and between digit strings in various number bases.
- [battery](https://github.com/distatus/battery) ![](https://img.shields.io/github/stars/distatus/battery.svg) - Cross-platform, normalized battery information library.
- [bitio](https://github.com/icza/bitio) ![](https://img.shields.io/github/stars/icza/bitio.svg) - Highly optimized bit-level Reader and Writer for Go.
- [browscap_go](https://github.com/digitalcrab/browscap_go) ![](https://img.shields.io/github/stars/digitalcrab/browscap_go.svg) - GoLang Library for [Browser Capabilities Project](https://browscap.org/).
- [captcha](https://github.com/steambap/captcha) ![](https://img.shields.io/github/stars/steambap/captcha.svg) - Package captcha provides an easy to use, unopinionated API for captcha generation.
- [common](https://github.com/kubeservice-stack/common) ![](https://img.shields.io/github/stars/kubeservice-stack/common.svg) - A library for server framework.
- [conv](https://github.com/cstockton/go-conv) ![](https://img.shields.io/github/stars/cstockton/go-conv.svg) - Package conv provides fast and intuitive conversions across Go types.
- [datacounter](https://github.com/miolini/datacounter) ![](https://img.shields.io/github/stars/miolini/datacounter.svg) - Go counters for readers/writer/http.ResponseWriter.
- [fake-useragent](https://github.com/lib4u/fake-useragent) ![](https://img.shields.io/github/stars/lib4u/fake-useragent.svg) - Up-to-date simple useragent faker with real world database in Golang
- [faker](https://github.com/pioz/faker) ![](https://img.shields.io/github/stars/pioz/faker.svg) - Random fake data and struct generator for Go.
- [ffmt](https://github.com/go-ffmt/ffmt) ![](https://img.shields.io/github/stars/go-ffmt/ffmt.svg) - Beautify data display for Humans.
- [gatus](https://github.com/TwinProduction/gatus) ![](https://img.shields.io/github/stars/TwinProduction/gatus.svg) - Automated service health dashboard.
- [go-commandbus](https://github.com/lana/go-commandbus) ![](https://img.shields.io/github/stars/lana/go-commandbus.svg) - A slight and pluggable command-bus for Go.
- [go-commons-pool](https://github.com/jolestar/go-commons-pool) ![](https://img.shields.io/github/stars/jolestar/go-commons-pool.svg) - Generic object pool for Golang.
- [go-openapi](https://github.com/go-openapi) ![](https://img.shields.io/github/stars/go-openapi/go-openapi.svg) - Collection of packages to parse and utilize open-api schemas.
- [go-resiliency](https://github.com/eapache/go-resiliency) ![](https://img.shields.io/github/stars/eapache/go-resiliency.svg) - Resiliency patterns for golang.
- [go-unarr](https://github.com/gen2brain/go-unarr) ![](https://img.shields.io/github/stars/gen2brain/go-unarr.svg) - Decompression library for RAR, TAR, ZIP and 7z archives.
- [gofakeit](https://github.com/brianvoe/gofakeit) ![](https://img.shields.io/github/stars/brianvoe/gofakeit.svg) - Random data generator written in go.
- [gommit](https://github.com/antham/gommit) ![](https://img.shields.io/github/stars/antham/gommit.svg) - Analyze git commit messages to ensure they follow defined patterns.
- [gopsutil](https://github.com/shirou/gopsutil) ![](https://img.shields.io/github/stars/shirou/gopsutil.svg) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).
- [gosh](https://github.com/osamingo/gosh) ![](https://img.shields.io/github/stars/osamingo/gosh.svg) - Provide Go Statistics Handler, Struct, Measure Method.
- [gosms](https://github.com/haxpax/gosms) ![](https://img.shields.io/github/stars/haxpax/gosms.svg) - Your own local SMS gateway in Go that can be used to send SMS.
- [gotoprom](https://github.com/cabify/gotoprom) ![](https://img.shields.io/github/stars/cabify/gotoprom.svg) - Type-safe metrics builder wrapper library for the official Prometheus client.
- [gountries](https://github.com/pariz/gountries) ![](https://img.shields.io/github/stars/pariz/gountries.svg) - Package that exposes country and subdivision data.
- [gtree](https://github.com/ddddddO/gtree) ![](https://img.shields.io/github/stars/ddddddO/gtree.svg) - Provide CLI, Package and Web for tree output and directories creation from Markdown or programmatically.
- [health](https://github.com/alexliesenfeld/health) ![](https://img.shields.io/github/stars/alexliesenfeld/health.svg) - A simple and flexible health check library for Go.
- [health](https://github.com/dimiro1/health) ![](https://img.shields.io/github/stars/dimiro1/health.svg) - Easy to use, extensible health check library.
- [healthcheck](https://github.com/etherlabsio/healthcheck) ![](https://img.shields.io/github/stars/etherlabsio/healthcheck.svg) - An opinionated and concurrent health-check HTTP handler for RESTful services.
- [hostutils](https://github.com/Wing924/hostutils) ![](https://img.shields.io/github/stars/Wing924/hostutils.svg) - A golang library for packing and unpacking FQDNs list.
- [indigo](https://github.com/osamingo/indigo) ![](https://img.shields.io/github/stars/osamingo/indigo.svg) - Distributed unique ID generator of using Sonyflake and encoded by Base58.
- [lk](https://github.com/hyperboloide/lk) ![](https://img.shields.io/github/stars/hyperboloide/lk.svg) - A simple licensing library for golang.
- [llvm](https://github.com/llir/llvm) ![](https://img.shields.io/github/stars/llir/llvm.svg) - Library for interacting with LLVM IR in pure Go.
- [metrics](https://github.com/pascaldekloe/metrics) ![](https://img.shields.io/github/stars/pascaldekloe/metrics.svg) - Library for metrics instrumentation and Prometheus exposition.
- [morse](https://github.com/alwindoss/morse) ![](https://img.shields.io/github/stars/alwindoss/morse.svg) - Library to convert to and from morse code.
- [numa](https://github.com/lrita/numa) ![](https://img.shields.io/github/stars/lrita/numa.svg) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.
- [pdfgen](https://github.com/hyperboloide/pdfgen) ![](https://img.shields.io/github/stars/hyperboloide/pdfgen.svg) - HTTP service to generate PDF from Json requests.
- [persian](https://github.com/mavihq/persian) ![](https://img.shields.io/github/stars/mavihq/persian.svg) - Some utilities for Persian language in go.
- [purego](https://github.com/ebitengine/purego) ![](https://img.shields.io/github/stars/ebitengine/purego.svg) - A library for calling C functions from Go without Cgo.
- [sandid](https://github.com/aofei/sandid) ![](https://img.shields.io/github/stars/aofei/sandid.svg) - Every grain of sand on earth has its own ID.
- [shellwords](https://github.com/Wing924/shellwords) ![](https://img.shields.io/github/stars/Wing924/shellwords.svg) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.
- [shortid](https://github.com/teris-io/shortid) ![](https://img.shields.io/github/stars/teris-io/shortid.svg) - Distributed generation of super short, unique, non-sequential, URL friendly IDs.
- [shoutrrr](https://github.com/containrrr/shoutrrr) ![](https://img.shields.io/github/stars/containrrr/shoutrrr.svg) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others.
- [sitemap-format](https://github.com/mingard/sitemap-format) ![](https://img.shields.io/github/stars/mingard/sitemap-format.svg) - A simple sitemap generator, with a little syntactic sugar.
- [stateless](https://github.com/qmuntal/stateless) ![](https://img.shields.io/github/stars/qmuntal/stateless.svg) - A fluent library for creating state machines.
- [stats](https://github.com/go-playground/stats) ![](https://img.shields.io/github/stars/go-playground/stats.svg) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...
- [turtle](https://github.com/hackebrot/turtle) ![](https://img.shields.io/github/stars/hackebrot/turtle.svg) - Emojis for Go.
- [url-shortener](https://github.com/pantrif/url-shortener) ![](https://img.shields.io/github/stars/pantrif/url-shortener.svg) - A modern, powerful, and robust URL shortener microservice with mysql support.
- [VarHandler](https://github.com/azr/generators/tree/master/varhandler) ![](https://img.shields.io/github/stars/azr/generators.svg) - Generate boilerplate http input and output handling.
- [varint](https://github.com/chmike/varint) ![](https://img.shields.io/github/stars/chmike/varint.svg) - A faster varying length integer encoder/decoder than the one provided in the standard library.
- [xdg](https://github.com/rkoesters/xdg) ![](https://img.shields.io/github/stars/rkoesters/xdg.svg) - FreeDesktop.org (xdg) Specs implemented in Go.
- [xkg](https://github.com/go-xkg/xkg) ![](https://img.shields.io/github/stars/go-xkg/xkg.svg) - X Keyboard Grabber.
- [xz](https://github.com/ulikunitz/xz) ![](https://img.shields.io/github/stars/ulikunitz/xz.svg) - Pure golang package for reading and writing xz-compressed files.
**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) ![](https://img.shields.io/github/stars/detectlanguage/detectlanguage-go.svg) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.
- [getlang](https://github.com/rylans/getlang) ![](https://img.shields.io/github/stars/rylans/getlang.svg) - Fast natural language detection package.
- [guesslanguage](https://github.com/endeveit/guesslanguage) ![](https://img.shields.io/github/stars/endeveit/guesslanguage.svg) - Functions to determine the natural language of a unicode text.
- [lingua-go](https://github.com/pemistahl/lingua-go) ![](https://img.shields.io/github/stars/pemistahl/lingua-go.svg) - An accurate natural language detection library, suitable for long and short text alike. Supports detecting multiple languages in mixed-language text.
- [whatlanggo](https://github.com/abadojack/whatlanggo) ![](https://img.shields.io/github/stars/abadojack/whatlanggo.svg) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).

### Morphological Analyzers

- [go-stem](https://github.com/agonopol/go-stem) ![](https://img.shields.io/github/stars/agonopol/go-stem.svg) - Implementation of the porter stemming algorithm.
- [go2vec](https://github.com/danieldk/go2vec) ![](https://img.shields.io/github/stars/danieldk/go2vec.svg) - Reader and utility functions for word2vec embeddings.
- [golibstemmer](https://github.com/rjohnsondev/golibstemmer) ![](https://img.shields.io/github/stars/rjohnsondev/golibstemmer.svg) - Go bindings for the snowball libstemmer library including porter 2.
- [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) ![](https://img.shields.io/github/stars/dinopuguh/gosentiwordnet.svg) - Sentiment analyzer using sentiwordnet lexicon in Go.
- [govader](https://github.com/jonreiter/govader) ![](https://img.shields.io/github/stars/jonreiter/govader.svg) - Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment).
- [govader-backend](https://github.com/PIMPfiction/govader_backend) ![](https://img.shields.io/github/stars/PIMPfiction/govader_backend.svg) - Microservice implementation of [GoVader](https://github.com/jonreiter/govader).
- [kagome](https://github.com/ikawaha/kagome) ![](https://img.shields.io/github/stars/ikawaha/kagome.svg) - JP morphological analyzer written in pure Go.
- [libtextcat](https://github.com/goodsign/libtextcat) ![](https://img.shields.io/github/stars/goodsign/libtextcat.svg) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
- [nlp](https://github.com/james-bowman/nlp) ![](https://img.shields.io/github/stars/james-bowman/nlp.svg) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).
- [paicehusk](https://github.com/rookii/paicehusk) ![](https://img.shields.io/github/stars/rookii/paicehusk.svg) - Golang implementation of the Paice/Husk Stemming Algorithm.
- [porter](https://github.com/a2800276/porter) ![](https://img.shields.io/github/stars/a2800276/porter.svg) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.
- [porter2](https://github.com/zhenjl/porter2) ![](https://img.shields.io/github/stars/zhenjl/porter2.svg) - Really fast Porter 2 stemmer.
- [RAKE.go](https://github.com/afjoseph/RAKE.Go) ![](https://img.shields.io/github/stars/afjoseph/RAKE.Go.svg) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
- [snowball](https://github.com/goodsign/snowball) ![](https://img.shields.io/github/stars/goodsign/snowball.svg) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
- [spaGO](https://github.com/nlpodyssey/spago) ![](https://img.shields.io/github/stars/nlpodyssey/spago.svg) - Self-contained Machine Learning and Natural Language Processing library in Go.
- [spelling-corrector](https://github.com/jorelosorio/spellingcorrector) ![](https://img.shields.io/github/stars/jorelosorio/spellingcorrector.svg) - A spelling corrector for the Spanish language or create your own.

### Slugifiers

- [go-slugify](https://github.com/mozillazg/go-slugify) ![](https://img.shields.io/github/stars/mozillazg/go-slugify.svg) - Make pretty slug with multiple languages support.
- [slug](https://github.com/gosimple/slug) ![](https://img.shields.io/github/stars/gosimple/slug.svg) - URL-friendly slugify with multiple languages support.
- [Slugify](https://github.com/avelino/slugify) ![](https://img.shields.io/github/stars/avelino/slugify.svg) - Go slugify application that handles string.

### Tokenizers

- [gojieba](https://github.com/yanyiwu/gojieba) ![](https://img.shields.io/github/stars/yanyiwu/gojieba.svg) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.
- [gotokenizer](https://github.com/xujiajun/gotokenizer) ![](https://img.shields.io/github/stars/xujiajun/gotokenizer.svg) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)
- [gse](https://github.com/go-ego/gse) ![](https://img.shields.io/github/stars/go-ego/gse.svg) - Go efficient text segmentation; support english, chinese, japanese and other.
- [MMSEGO](https://github.com/awsong/MMSEGO) ![](https://img.shields.io/github/stars/awsong/MMSEGO.svg) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
- [segment](https://github.com/blevesearch/segment) ![](https://img.shields.io/github/stars/blevesearch/segment.svg) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/)
- [sentences](https://github.com/neurosnap/sentences) ![](https://img.shields.io/github/stars/neurosnap/sentences.svg) - Sentence tokenizer: converts text into a list of sentences.
- [shamoji](https://github.com/osamingo/shamoji) ![](https://img.shields.io/github/stars/osamingo/shamoji.svg) - The shamoji is word filtering package written in Go.
- [stemmer](https://github.com/dchest/stemmer) ![](https://img.shields.io/github/stars/dchest/stemmer.svg) - Stemmer packages for Go programming language. Includes English and German stemmers.
- [textcat](https://github.com/pebbe/textcat) ![](https://img.shields.io/github/stars/pebbe/textcat.svg) - Go package for n-gram based text categorization, with support for utf-8 and raw text.

### Translation

- [ctxi18n](https://github.com/invopop/ctxi18n/) ![](https://img.shields.io/github/stars/invopop/ctxi18n.svg) - Context aware i18n with a short and consise API, pluralization, interpolation, and `fs.FS` support. YAML locale definitions are based on [Rails i18n](https://guides.rubyonrails.org/i18n.html).
- [go-i18n](https://github.com/nicksnyder/go-i18n/) ![](https://img.shields.io/github/stars/nicksnyder/go-i18n.svg) - Package and an accompanying tool to work with localized text.
- [go-mystem](https://github.com/dveselov/mystem) ![](https://img.shields.io/github/stars/dveselov/mystem.svg) - CGo bindings to Yandex.Mystem - russian morphology analyzer.
- [go-pinyin](https://github.com/mozillazg/go-pinyin) ![](https://img.shields.io/github/stars/mozillazg/go-pinyin.svg) - CN Hanzi to Hanyu Pinyin converter.
- [go-words](https://github.com/saleh-rahimzadeh/go-words) ![](https://img.shields.io/github/stars/saleh-rahimzadeh/go-words.svg) - A words table and text resource library for Golang projects.
- [gotext](https://github.com/leonelquinteros/gotext) ![](https://img.shields.io/github/stars/leonelquinteros/gotext.svg) - GNU gettext utilities for Go.
- [iuliia-go](https://github.com/mehanizm/iuliia-go) ![](https://img.shields.io/github/stars/mehanizm/iuliia-go.svg) - Transliterate Cyrillic → Latin in every possible way.
- [spreak](https://github.com/vorlif/spreak) ![](https://img.shields.io/github/stars/vorlif/spreak.svg) - Flexible translation and humanization library for Go, based on the concepts behind gettext.
- [t](https://github.com/youthlin/t) ![](https://img.shields.io/github/stars/youthlin/t.svg) - Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template.

### Transliteration

- [enca](https://github.com/endeveit/enca) ![](https://img.shields.io/github/stars/endeveit/enca.svg) - Minimal cgo bindings for [libenca](https://cihar.com/software/enca/), which detects character encodings.
- [go-unidecode](https://github.com/mozillazg/go-unidecode) ![](https://img.shields.io/github/stars/mozillazg/go-unidecode.svg) - ASCII transliterations of Unicode text.
- [gounidecode](https://github.com/fiam/gounidecode) ![](https://img.shields.io/github/stars/fiam/gounidecode.svg) - Unicode transliterator (also known as unidecode) for Go.
- [transliterator](https://github.com/alexsergivan/transliterator) ![](https://img.shields.io/github/stars/alexsergivan/transliterator.svg) - Provides one-way string transliteration with supporting of language-specific transliteration rules.

**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [arp](https://github.com/mdlayher/arp) ![](https://img.shields.io/github/stars/mdlayher/arp.svg) - Package arp implements the ARP protocol, as described in RFC 826.
- [bart](https://github.com/gaissmai/bart) ![](https://img.shields.io/github/stars/gaissmai/bart.svg) - Package bart provides a Balanced-Routing-Table (BART) for very fast IP to CIDR lookups and more.
- [buffstreams](https://github.com/stabbycutyou/buffstreams) ![](https://img.shields.io/github/stars/stabbycutyou/buffstreams.svg) - Streaming protocolbuffer data over TCP made easy.
- [canopus](https://github.com/zubairhamed/canopus) ![](https://img.shields.io/github/stars/zubairhamed/canopus.svg) - CoAP Client/Server implementation (RFC 7252).
- [cidranger](https://github.com/yl2chen/cidranger) ![](https://img.shields.io/github/stars/yl2chen/cidranger.svg) - Fast IP to CIDR lookup for Go.
- [cloudflared](https://github.com/cloudflare/cloudflared) ![](https://img.shields.io/github/stars/cloudflare/cloudflared.svg) - Cloudflare Tunnel client (formerly Argo Tunnel).
- [dhcp6](https://github.com/mdlayher/dhcp6) ![](https://img.shields.io/github/stars/mdlayher/dhcp6.svg) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
- [dns](https://github.com/miekg/dns) ![](https://img.shields.io/github/stars/miekg/dns.svg) - Go library for working with DNS.
- [dnsmonster](https://github.com/mosajjal/dnsmonster) ![](https://img.shields.io/github/stars/mosajjal/dnsmonster.svg) - Passive DNS Capture/Monitoring Framework.
- [easytcp](https://github.com/DarthPestilane/easytcp) ![](https://img.shields.io/github/stars/DarthPestilane/easytcp.svg) - A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful.
- [ether](https://github.com/songgao/ether) ![](https://img.shields.io/github/stars/songgao/ether.svg) - Cross-platform Go package for sending and receiving ethernet frames.
- [ethernet](https://github.com/mdlayher/ethernet) ![](https://img.shields.io/github/stars/mdlayher/ethernet.svg) - Package ethernet implements marshaling and unmarshalling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.
- [event](https://github.com/cheng-zhongliang/event) ![](https://img.shields.io/github/stars/cheng-zhongliang/event.svg) - Simple I/O event notification library written in Golang.
- [fasthttp](https://github.com/valyala/fasthttp) ![](https://img.shields.io/github/stars/valyala/fasthttp.svg) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.
- [fortio](https://github.com/fortio/fortio) ![](https://img.shields.io/github/stars/fortio/fortio.svg) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.
- [ftp](https://github.com/jlaffaye/ftp) ![](https://img.shields.io/github/stars/jlaffaye/ftp.svg) - Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959).
- [ftpserverlib](https://github.com/fclairamb/ftpserverlib) ![](https://img.shields.io/github/stars/fclairamb/ftpserverlib.svg) - Fully featured FTP server library.
- [fullproxy](https://github.com/shoriwe/fullproxy) ![](https://img.shields.io/github/stars/shoriwe/fullproxy.svg) - A fully featured scriptable and daemon configurable proxy and pivoting toolkit with SOCKS5, HTTP, raw ports and reverse proxy protocols.
- [fwdctl](https://github.com/alegrey91/fwdctl) ![](https://img.shields.io/github/stars/alegrey91/fwdctl.svg) - A simple and intuitive CLI to manage IPTables forwards in your Linux server.
- [gaio](https://github.com/xtaci/gaio) ![](https://img.shields.io/github/stars/xtaci/gaio.svg) - High performance async-io networking for Golang in proactor mode.
- [gev](https://github.com/Allenxuxu/gev) ![](https://img.shields.io/github/stars/Allenxuxu/gev.svg) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.
- [gldap](https://github.com/jimlambrt/gldap) ![](https://img.shields.io/github/stars/jimlambrt/gldap.svg) - gldap provides an ldap server implementation and you provide handlers for its ldap operations.
- [gmqtt](https://github.com/DrmagicE/gmqtt) ![](https://img.shields.io/github/stars/DrmagicE/gmqtt.svg) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.
- [gnet](https://github.com/panjf2000/gnet) ![](https://img.shields.io/github/stars/panjf2000/gnet.svg) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.
- [gnet](https://github.com/fish-tennis/gnet) ![](https://img.shields.io/github/stars/fish-tennis/gnet.svg) - `gnet` is a high-performance networking framework,especially for game servers.
- [gNxI](https://github.com/google/gnxi) ![](https://img.shields.io/github/stars/google/gnxi.svg) - A collection of tools for Network Management that use the gNMI and gNOI protocols.
- [go-getter](https://github.com/hashicorp/go-getter) ![](https://img.shields.io/github/stars/hashicorp/go-getter.svg) - Go library for downloading files or directories from various sources using a URL.
- [go-multiproxy](https://github.com/presbrey/go-multiproxy) ![](https://img.shields.io/github/stars/presbrey/go-multiproxy.svg) - Library for making HTTP requests through a pool of proxies offering fault tolerance, load balancing, automatic retries, cookie management, and more, via http.Get/Post replacement or http.Client RoundTripper drop-in
- [go-pcaplite](https://github.com/alexcfv/go-pcaplite) ![](https://img.shields.io/github/stars/alexcfv/go-pcaplite.svg) - Lightweight live packet capture library with HTTPS SNI extraction.
- [go-powerdns](https://github.com/joeig/go-powerdns) ![](https://img.shields.io/github/stars/joeig/go-powerdns.svg) - PowerDNS API bindings for Golang.
- [go-sse](https://github.com/lampctl/go-sse) ![](https://img.shields.io/github/stars/lampctl/go-sse.svg) - Go client and server implementation of HTML server-sent events.
- [go-stun](https://github.com/ccding/go-stun) ![](https://img.shields.io/github/stars/ccding/go-stun.svg) - Go implementation of the STUN client (RFC 3489 and RFC 5389).
- [gobgp](https://github.com/osrg/gobgp) ![](https://img.shields.io/github/stars/osrg/gobgp.svg) - BGP implemented in the Go Programming Language.
- [gopacket](https://github.com/google/gopacket) ![](https://img.shields.io/github/stars/google/gopacket.svg) - Go library for packet processing with libpcap bindings.
- [gopcap](https://github.com/akrennmair/gopcap) ![](https://img.shields.io/github/stars/akrennmair/gopcap.svg) - Go wrapper for libpcap.
- [GoProxy](https://github.com/elazarl/goproxy) ![](https://img.shields.io/github/stars/elazarl/goproxy.svg) - A library to create a customized HTTP/HTTPS proxy server using Go.
- [goshark](https://github.com/sunwxg/goshark) ![](https://img.shields.io/github/stars/sunwxg/goshark.svg) - Package goshark use tshark to decode IP packet and create data struct to analyse packet.
- [gosnmp](https://github.com/soniah/gosnmp) ![](https://img.shields.io/github/stars/soniah/gosnmp.svg) - Native Go library for performing SNMP actions.
- [gotcp](https://github.com/gansidui/gotcp) ![](https://img.shields.io/github/stars/gansidui/gotcp.svg) - Go package for quickly writing tcp applications.
- [grab](https://github.com/cavaliercoder/grab) ![](https://img.shields.io/github/stars/cavaliercoder/grab.svg) - Go package for managing file downloads.
- [graval](https://github.com/koofr/graval) ![](https://img.shields.io/github/stars/koofr/graval.svg) - Experimental FTP server framework.
- [gws](https://github.com/lxzan/gws) ![](https://img.shields.io/github/stars/lxzan/gws.svg) - High-Performance WebSocket Server & Client With AsyncIO Supporting .
- [HTTPLab](https://github.com/gchaincl/httplab) ![](https://img.shields.io/github/stars/gchaincl/httplab.svg) - HTTPLabs let you inspect HTTP requests and forge responses.
- [httpproxy](https://github.com/wzshiming/httpproxy) ![](https://img.shields.io/github/stars/wzshiming/httpproxy.svg) - HTTP proxy handler and dialer.
- [iplib](https://github.com/c-robinson/iplib) ![](https://img.shields.io/github/stars/c-robinson/iplib.svg) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)
- [jazigo](https://github.com/udhos/jazigo) ![](https://img.shields.io/github/stars/udhos/jazigo.svg) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices.
- [kcp-go](https://github.com/xtaci/kcp-go) ![](https://img.shields.io/github/stars/xtaci/kcp-go.svg) - KCP - Fast and Reliable ARQ Protocol.
- [kcptun](https://github.com/xtaci/kcptun) ![](https://img.shields.io/github/stars/xtaci/kcptun.svg) - Extremely simple & fast udp tunnel based on KCP protocol.
- [lhttp](https://github.com/fanux/lhttp) ![](https://img.shields.io/github/stars/fanux/lhttp.svg) - Powerful websocket framework, build your IM server more easily.
- [linkio](https://github.com/ian-kent/linkio) ![](https://img.shields.io/github/stars/ian-kent/linkio.svg) - Network link speed simulation for Reader/Writer interfaces.
- [llb](https://github.com/kirillDanshin/llb) ![](https://img.shields.io/github/stars/kirillDanshin/llb.svg) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.
- [mdns](https://github.com/hashicorp/mdns) ![](https://img.shields.io/github/stars/hashicorp/mdns.svg) - Simple mDNS (Multicast DNS) client/server library in Golang.
- [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
- [natiu-mqtt](https://github.com/soypat/natiu-mqtt) ![](https://img.shields.io/github/stars/soypat/natiu-mqtt.svg) - A dead-simple, non-allocating, low level implementation of MQTT well suited for embedded systems.
- [nbio](https://github.com/lesismal/nbio) ![](https://img.shields.io/github/stars/lesismal/nbio.svg) - Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use.
- [net](https://golang.org/x/net) - This repository holds supplementary Go networking libraries.
- [netpoll](https://github.com/cloudwego/netpoll) ![](https://img.shields.io/github/stars/cloudwego/netpoll.svg) - A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance.
- [NFF-Go](https://github.com/intel-go/nff-go) ![](https://img.shields.io/github/stars/intel-go/nff-go.svg) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).
- [nodepass](https://github.com/NodePassProject/nodepass) ![](https://img.shields.io/github/stars/NodePassProject/nodepass.svg) - A secure, efficient TCP/UDP tunneling solution that delivers fast, reliable access across network restrictions using pre-established TCP/QUIC/WebSocket or HTTP/2 connections.
- [peerdiscovery](https://github.com/schollz/peerdiscovery) ![](https://img.shields.io/github/stars/schollz/peerdiscovery.svg) - Pure Go library for cross-platform local peer discovery using UDP multicast.
- [portproxy](https://github.com/aybabtme/portproxy) ![](https://img.shields.io/github/stars/aybabtme/portproxy.svg) - Simple TCP proxy which adds CORS support to API's which don't support it.
- [psql-wire](https://github.com/jeroenrinzema/psql-wire) ![](https://img.shields.io/github/stars/jeroenrinzema/psql-wire.svg) - PostgreSQL server wire protocol. Build your own server and start serving connections..
- [publicip](https://github.com/polera/publicip) ![](https://img.shields.io/github/stars/polera/publicip.svg) - Package publicip returns your public facing IPv4 address (internet egress).
- [quic-go](https://github.com/lucas-clemente/quic-go) ![](https://img.shields.io/github/stars/lucas-clemente/quic-go.svg) - An implementation of the QUIC protocol in pure Go.
- [sdns](https://github.com/semihalev/sdns) ![](https://img.shields.io/github/stars/semihalev/sdns.svg) - A high-performance, recursive DNS resolver server with DNSSEC support, focused on preserving privacy.
- [sftp](https://github.com/pkg/sftp) ![](https://img.shields.io/github/stars/pkg/sftp.svg) - Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>.
- [ssh](https://github.com/gliderlabs/ssh) ![](https://img.shields.io/github/stars/gliderlabs/ssh.svg) - Higher-level API for building SSH servers (wraps crypto/ssh).
- [sslb](https://github.com/eduardonunesp/sslb) ![](https://img.shields.io/github/stars/eduardonunesp/sslb.svg) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.
- [stun](https://github.com/go-rtc/stun) ![](https://img.shields.io/github/stars/go-rtc/stun.svg) - Go implementation of RFC 5389 STUN protocol.
- [tcpack](https://github.com/lim-yoona/tcpack) ![](https://img.shields.io/github/stars/lim-yoona/tcpack.svg) - tcpack is an application protocol based on TCP to Pack and Unpack bytes stream in go program.
- [tspool](https://github.com/two/tspool) ![](https://img.shields.io/github/stars/two/tspool.svg) - A TCP Library use worker pool to improve performance and protect your server.
- [tun2socks](https://github.com/xjasonlyu/tun2socks) ![](https://img.shields.io/github/stars/xjasonlyu/tun2socks.svg) - A pure go implementation of tun2socks powered by [gVisor](https://gvisor.dev/) TCP/IP stack.
- [utp](https://github.com/anacrolix/utp) ![](https://img.shields.io/github/stars/anacrolix/utp.svg) - Go uTP micro transport protocol implementation.
- [vssh](https://github.com/yahoo/vssh) ![](https://img.shields.io/github/stars/yahoo/vssh.svg) - Go library for building network and server automation over SSH protocol.
- [water](https://github.com/songgao/water) ![](https://img.shields.io/github/stars/songgao/water.svg) - Simple TUN/TAP library.
- [webrtc](https://github.com/pions/webrtc) ![](https://img.shields.io/github/stars/pions/webrtc.svg) - A pure Go implementation of the WebRTC API.
- [winrm](https://github.com/masterzen/winrm) ![](https://img.shields.io/github/stars/masterzen/winrm.svg) - Go WinRM client to remotely execute commands on Windows machines.
- [xtcp](https://github.com/xfxdev/xtcp) ![](https://img.shields.io/github/stars/xfxdev/xtcp.svg) - TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [axios4go](https://github.com/rezmoss/axios4go) ![](https://img.shields.io/github/stars/rezmoss/axios4go.svg) - A Go HTTP client library inspired by Axios, providing a simple and intuitive API for making HTTP requests.
- [azuretls-client](https://github.com/Noooste/azuretls-client) ![](https://img.shields.io/github/stars/Noooste/azuretls-client.svg) - An easy-to-use HTTP client 100% in Go to spoof TLS/JA3 and HTTP2 fingerprint.
- [fast-shot](https://github.com/opus-domini/fast-shot) ![](https://img.shields.io/github/stars/opus-domini/fast-shot.svg) - Hit your API targets with rapid-fire precision using Go's fastest and simple HTTP Client.
- [gentleman](https://github.com/h2non/gentleman) ![](https://img.shields.io/github/stars/h2non/gentleman.svg) - Full-featured plugin-driven HTTP client library.
- [go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) ![](https://img.shields.io/github/stars/hashicorp/go-cleanhttp.svg) - Get easily stdlib HTTP client, which does not share any state with other clients.
- [go-http-client](https://github.com/bozd4g/go-http-client) ![](https://img.shields.io/github/stars/bozd4g/go-http-client.svg) - Make http calls simply and easily.
- [go-ipmux](https://github.com/optimus-hft/go-ipmux) ![](https://img.shields.io/github/stars/optimus-hft/go-ipmux.svg) - A library for Multiplexing HTTP requests based on multiple Source IPs.
- [go-otelroundtripper](https://github.com/NdoleStudio/go-otelroundtripper) ![](https://img.shields.io/github/stars/NdoleStudio/go-otelroundtripper.svg) - Go http.RoundTripper that emits open telemetry metrics for HTTP requests.
- [go-req](https://github.com/wenerme/go-req) ![](https://img.shields.io/github/stars/wenerme/go-req.svg) - Declarative golang HTTP client.
- [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) ![](https://img.shields.io/github/stars/hashicorp/go-retryablehttp.svg) - Retryable HTTP client in Go.
- [go-zoox/fetch](https://github.com/go-zoox/fetch) ![](https://img.shields.io/github/stars/go-zoox/fetch.svg) - A Powerful, Lightweight, Easy Http Client, inspired by Web Fetch API.
- [Grequest](https://github.com/lib4u/grequest) ![](https://img.shields.io/github/stars/lib4u/grequest.svg)  - Simple and lightweight golang package for http requests. based on powerful net/http
- [grequests](https://github.com/levigross/grequests) ![](https://img.shields.io/github/stars/levigross/grequests.svg) - A Go "clone" of the great and famous Requests library.
- [heimdall](https://github.com/gojektech/heimdall) ![](https://img.shields.io/github/stars/gojektech/heimdall.svg) - An enhanced http client with retry and hystrix capabilities.
- [httpretry](https://github.com/ybbus/httpretry) ![](https://img.shields.io/github/stars/ybbus/httpretry.svg) - Enriches the default go HTTP client with retry functionality.
- [pester](https://github.com/sethgrid/pester) ![](https://img.shields.io/github/stars/sethgrid/pester.svg) - Go HTTP client calls with retries, backoff, and concurrency.
- [req](https://github.com/imroc/req) ![](https://img.shields.io/github/stars/imroc/req.svg) - Simple Go HTTP client with Black Magic (Less code and More efficiency).
- [request](https://github.com/monaco-io/request) ![](https://img.shields.io/github/stars/monaco-io/request.svg) - HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency.
- [requests](https://github.com/carlmjohnson/requests) ![](https://img.shields.io/github/stars/carlmjohnson/requests.svg) - HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools.
- [resty](https://github.com/go-resty/resty) ![](https://img.shields.io/github/stars/go-resty/resty.svg) - Simple HTTP and REST client for Go inspired by Ruby rest-client.
- [rq](https://github.com/ddo/rq) ![](https://img.shields.io/github/stars/ddo/rq.svg) - A nicer interface for golang stdlib HTTP client.
- [sling](https://github.com/dghubble/sling) ![](https://img.shields.io/github/stars/dghubble/sling.svg) - Sling is a Go HTTP client library for creating and sending API requests.
- [surf](https://github.com/enetx/surf) ![](https://img.shields.io/github/stars/enetx/surf.svg) - Advanced HTTP client with HTTP/1.1, HTTP/2, HTTP/3 (QUIC), SOCKS5 proxy support and browser-grade TLS fingerprinting.
- [tls-client](https://github.com/bogdanfinn/tls-client) ![](https://img.shields.io/github/stars/bogdanfinn/tls-client.svg) - net/http.Client like HTTP Client with options to select specific client TLS Fingerprints to use for requests.

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [gl](https://github.com/go-gl/gl) ![](https://img.shields.io/github/stars/go-gl/gl.svg) - Go bindings for OpenGL (generated via glow).
- [glfw](https://github.com/go-gl/glfw) ![](https://img.shields.io/github/stars/go-gl/glfw.svg) - Go bindings for GLFW 3.
- [go-glmatrix](https://github.com/technohippy/go-glmatrix) ![](https://img.shields.io/github/stars/technohippy/go-glmatrix.svg) - Go port of [glMatrix](https://glmatrix.net/) library.
- [goxjs/gl](https://github.com/goxjs/gl) ![](https://img.shields.io/github/stars/goxjs/gl.svg) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
- [goxjs/glfw](https://github.com/goxjs/glfw) ![](https://img.shields.io/github/stars/goxjs/glfw.svg) - Go cross-platform glfw library for creating an OpenGL context and receiving events.
- [mathgl](https://github.com/go-gl/mathgl) ![](https://img.shields.io/github/stars/go-gl/mathgl.svg) - Pure Go math package specialized for 3D math, with inspiration from GLM.

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [bob](https://github.com/stephenafamo/bob) ![](https://img.shields.io/github/stars/stephenafamo/bob.svg) - SQL query builder and ORM/Factory generator for Go. Successor of SQLBoiler.
- [bun](https://github.com/uptrace/bun) ![](https://img.shields.io/github/stars/uptrace/bun.svg) - SQL-first Golang ORM. Successor of go-pg.
- [cacheme](https://github.com/Yiling-J/cacheme-go) ![](https://img.shields.io/github/stars/Yiling-J/cacheme-go.svg) - Schema based, typed Redis caching/memoize framework for Go.
- [CQL](https://github.com/FrancoLiberali/cql) ![](https://img.shields.io/github/stars/FrancoLiberali/cql.svg) - Built on top of GORM, adds compile-time verified queries based on auto-generated code.
- [ent](https://github.com/facebook/ent) ![](https://img.shields.io/github/stars/facebook/ent.svg) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
- [go-dbw](https://github.com/hashicorp/go-dbw) ![](https://img.shields.io/github/stars/hashicorp/go-dbw.svg) - A simple package that encapsulates database operations.
- [go-firestorm](https://github.com/jschoedt/go-firestorm) ![](https://img.shields.io/github/stars/jschoedt/go-firestorm.svg) - A simple ORM for Google/Firebase Cloud Firestore.
- [go-sql](https://github.com/rushteam/gosql) ![](https://img.shields.io/github/stars/rushteam/gosql.svg) - A easy ORM for mysql.
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) ![](https://img.shields.io/github/stars/huandu/go-sqlbuilder.svg) - A flexible and powerful SQL string builder library plus a zero-config ORM.
- [go-store](https://github.com/gosuri/go-store) ![](https://img.shields.io/github/stars/gosuri/go-store.svg) - Simple and fast Redis backed key-value store library for Go.
- [golobby/orm](https://github.com/golobby/orm) ![](https://img.shields.io/github/stars/golobby/orm.svg) - Simple, fast, type-safe, generic orm for developer happiness.
- [GORM](https://github.com/go-gorm/gorm) ![](https://img.shields.io/github/stars/go-gorm/gorm.svg) - The fantastic ORM library for Golang, aims to be developer friendly.
- [gormt](https://github.com/xxjwxc/gormt) ![](https://img.shields.io/github/stars/xxjwxc/gormt.svg) - Mysql database to golang gorm struct.
- [gorp](https://github.com/go-gorp/gorp) ![](https://img.shields.io/github/stars/go-gorp/gorp.svg) - Go Relational Persistence, ORM-ish library for Go.
- [grimoire](https://github.com/Fs02/grimoire) ![](https://img.shields.io/github/stars/Fs02/grimoire.svg) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).
- [lore](https://github.com/abrahambotros/lore) ![](https://img.shields.io/github/stars/abrahambotros/lore.svg) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.
- [marlow](https://github.com/marlow/marlow) ![](https://img.shields.io/github/stars/marlow/marlow.svg) - Generated ORM from project structs for compile time safety assurances.
- [pop/soda](https://github.com/gobuffalo/pop) ![](https://img.shields.io/github/stars/gobuffalo/pop.svg) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
- [Prisma](https://github.com/prisma/prisma-client-go) ![](https://img.shields.io/github/stars/prisma/prisma-client-go.svg) - Prisma Client Go, Typesafe database access for Go.
- [reform](https://github.com/go-reform/reform) ![](https://img.shields.io/github/stars/go-reform/reform.svg) - Better ORM for Go, based on non-empty interfaces and code generation.
- [rel](https://github.com/go-rel/rel) ![](https://img.shields.io/github/stars/go-rel/rel.svg) - Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) ![](https://img.shields.io/github/stars/volatiletech/sqlboiler.svg) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.
- [upper.io/db](https://github.com/upper/db) ![](https://img.shields.io/github/stars/upper/db.svg) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
- [XORM](https://gitea.com/xorm/xorm) - Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).
- [Zoom](https://github.com/albrow/zoom) ![](https://img.shields.io/github/stars/albrow/zoom.svg) - Blazing-fast datastore and querying engine built on Redis.

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_

- [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

_Unofficial libraries for package and dependency management._

- [gup](https://github.com/nao1215/gup) ![](https://img.shields.io/github/stars/nao1215/gup.svg) - Update binaries installed by "go install".
- [modup](https://github.com/chaindead/modup) ![](https://img.shields.io/github/stars/chaindead/modup.svg) - Terminal UI for Go dependency updates with outdated module detection and selective upgrading.
- [syft](https://github.com/anchore/syft) ![](https://img.shields.io/github/stars/anchore/syft.svg) - A CLI tool and Go library for generating a Software Bill of Materials (SBOM) from container images and filesystems.

**[⬆ back to top](#contents)**

## Performance

- [ebpf-go](https://github.com/cilium/ebpf) ![](https://img.shields.io/github/stars/cilium/ebpf.svg) - Provides utilities for loading, compiling, and debugging eBPF programs.
- [go-instrument](https://github.com/nikolaydubina/go-instrument) ![](https://img.shields.io/github/stars/nikolaydubina/go-instrument.svg) - Automatically add spans to all methods and functions.
- [jaeger](https://github.com/jaegertracing/jaeger) ![](https://img.shields.io/github/stars/jaegertracing/jaeger.svg) - A distributed tracing system.
- [mm-go](https://github.com/joetifa2003/mm-go) ![](https://img.shields.io/github/stars/joetifa2003/mm-go.svg) - Generic manual memory management for golang.
- [pixie](https://github.com/pixie-labs/pixie) ![](https://img.shields.io/github/stars/pixie-labs/pixie.svg) - No instrumentation tracing for Golang applications via eBPF.
- [profile](https://github.com/pkg/profile) ![](https://img.shields.io/github/stars/pkg/profile.svg) - Simple profiling support package for Go.
- [statsviz](https://github.com/arl/statsviz) ![](https://img.shields.io/github/stars/arl/statsviz.svg) - Live visualization of your Go application runtime statistics.
- [tracer](https://github.com/kamilsk/tracer) ![](https://img.shields.io/github/stars/kamilsk/tracer.svg) - Simple, lightweight tracing.

**[⬆ back to top](#contents)**

## Query Language

- [api-fu](https://github.com/ccbrown/api-fu) ![](https://img.shields.io/github/stars/ccbrown/api-fu.svg) - Comprehensive GraphQL implementation.
- [dasel](https://github.com/tomwright/dasel) ![](https://img.shields.io/github/stars/tomwright/dasel.svg) - Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies.
- [gojsonq](https://github.com/thedevsaddam/gojsonq) ![](https://img.shields.io/github/stars/thedevsaddam/gojsonq.svg) - A simple Go package to Query over JSON Data.
- [goven](https://github.com/SeldonIO/goven) ![](https://img.shields.io/github/stars/SeldonIO/goven.svg) - A drop-in query language for any database schema.
- [gqlgen](https://github.com/99designs/gqlgen) ![](https://img.shields.io/github/stars/99designs/gqlgen.svg) - go generate based graphql server library.
- [grapher](https://github.com/reaganiwadha/grapher) ![](https://img.shields.io/github/stars/reaganiwadha/grapher.svg) - A GraphQL field builder utilizing Go generics with extra utilities and features.
- [graphql](https://github.com/neelance/graphql-go) ![](https://img.shields.io/github/stars/neelance/graphql-go.svg) - GraphQL server with a focus on ease of use.
- [graphql-go](https://github.com/graphql-go/graphql) ![](https://img.shields.io/github/stars/graphql-go/graphql.svg) - Implementation of GraphQL for Go.
- [gws](https://github.com/Zaba505/gws) ![](https://img.shields.io/github/stars/Zaba505/gws.svg) - Apollos' "GraphQL over Websocket" client and server implementation.
- [jsonpath](https://github.com/AsaiYusuke/jsonpath) ![](https://img.shields.io/github/stars/AsaiYusuke/jsonpath.svg) - A query library for retrieving part of JSON based on JSONPath syntax.
- [jsonql](https://github.com/elgs/jsonql) ![](https://img.shields.io/github/stars/elgs/jsonql.svg) - JSON query expression library in Golang.
- [jsonslice](https://github.com/bhmj/jsonslice) ![](https://img.shields.io/github/stars/bhmj/jsonslice.svg) - Jsonpath queries with advanced filters.
- [mql](https://github.com/hashicorp/mql) ![](https://img.shields.io/github/stars/hashicorp/mql.svg) - Model Query Language (mql) is a query language for your database models.
- [play](https://github.com/paololazzari/play) ![](https://img.shields.io/github/stars/paololazzari/play.svg) - A TUI playground to experiment with your favorite programs, such as grep, sed, awk, jq and yq.
- [rql](https://github.com/a8m/rql) ![](https://img.shields.io/github/stars/a8m/rql.svg) - Resource Query Language for REST API.
- [rqp](https://github.com/timsolov/rest-query-parser) ![](https://img.shields.io/github/stars/timsolov/rest-query-parser.svg) - Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query.
- [straf](https://github.com/SonicRoshan/straf) ![](https://img.shields.io/github/stars/SonicRoshan/straf.svg) - Easily Convert Golang structs to GraphQL objects.

**[⬆ back to top](#contents)**

## Reflection

- [copy](https://github.com/gotidy/copy) ![](https://img.shields.io/github/stars/gotidy/copy.svg) - Package for fast copying structs of different types.
- [Deepcopier](https://github.com/ulule/deepcopier) ![](https://img.shields.io/github/stars/ulule/deepcopier.svg) - Simple struct copying for Go.
- [go-deepcopy](https://github.com/tiendc/go-deepcopy) ![](https://img.shields.io/github/stars/tiendc/go-deepcopy.svg) - Fast deep copy library.
- [goenum](https://github.com/lvyahui8/goenum) ![](https://img.shields.io/github/stars/lvyahui8/goenum.svg) - A common enumeration struct based on generics and reflection that allows you to quickly define enumerations and use a set of useful default methods.
- [gotype](https://github.com/wzshiming/gotype) ![](https://img.shields.io/github/stars/wzshiming/gotype.svg) - Golang source code parsing, usage like reflect package.
- [gpath](https://github.com/tenntenn/gpath) ![](https://img.shields.io/github/stars/tenntenn/gpath.svg) - Library to simplify access struct fields with Go's expression in reflection.
- [objwalker](https://github.com/rekby/objwalker) ![](https://img.shields.io/github/stars/rekby/objwalker.svg) - Walk by go objects with reflection.
- [reflectpro](https://github.com/gontainer/reflectpro) ![](https://img.shields.io/github/stars/gontainer/reflectpro.svg) - Callers, copiers, getters and setters for go.
- [reflectutils](https://github.com/muir/reflectutils) ![](https://img.shields.io/github/stars/muir/reflectutils.svg) - Helpers for working with reflection: struct tag parsing; recursive walking; fill value from string.

**[⬆ back to top](#contents)**

## Resource Embedding

- [debme](https://github.com/leaanthony/debme) ![](https://img.shields.io/github/stars/leaanthony/debme.svg) - Create an `embed.FS` from an existing `embed.FS` subdirectory.
- [embed](https://pkg.go.dev/embed) - Package embed provides access to files embedded in the running Go program.
- [rebed](https://github.com/soypat/rebed) ![](https://img.shields.io/github/stars/soypat/rebed.svg) - Recreate folder structures and files from Go 1.16's `embed.FS` type
- [vfsgen](https://github.com/shurcooL/vfsgen) ![](https://img.shields.io/github/stars/shurcooL/vfsgen.svg) - Generates a vfsdata.go file that statically implements the given virtual filesystem.

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [assocentity](https://github.com/ndabAP/assocentity) ![](https://img.shields.io/github/stars/ndabAP/assocentity.svg) - Package assocentity returns the average distance from words to a given entity.
- [bradleyterry](https://github.com/seanhagen/bradleyterry) ![](https://img.shields.io/github/stars/seanhagen/bradleyterry.svg) - Provides a Bradley-Terry Model for pairwise comparisons.
- [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) ![](https://img.shields.io/github/stars/nikolaydubina/calendarheatmap.svg) - Calendar heatmap in plain Go inspired by Github contribution activity.
- [chart](https://github.com/vdobler/chart) ![](https://img.shields.io/github/stars/vdobler/chart.svg) - Simple Chart Plotting library for Go. Supports many graphs types.
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) ![](https://img.shields.io/github/stars/rocketlaunchr/dataframe-go.svg) - Dataframes for machine-learning and statistics (similar to pandas).
- [decimal](https://github.com/db47h/decimal) ![](https://img.shields.io/github/stars/db47h/decimal.svg) - Package decimal implements arbitrary-precision decimal floating-point arithmetic.
- [evaler](https://github.com/soniah/evaler) ![](https://img.shields.io/github/stars/soniah/evaler.svg) - Simple floating point arithmetic expression evaluator.
- [ewma](https://github.com/VividCortex/ewma) ![](https://img.shields.io/github/stars/VividCortex/ewma.svg) - Exponentially-weighted moving averages.
- [geom](https://github.com/skelterjohn/geom) ![](https://img.shields.io/github/stars/skelterjohn/geom.svg) - 2D geometry for golang.
- [go-dsp](https://github.com/mjibson/go-dsp) ![](https://img.shields.io/github/stars/mjibson/go-dsp.svg) - Digital Signal Processing for Go.
- [go-estimate](https://github.com/milosgajdos/go-estimate) ![](https://img.shields.io/github/stars/milosgajdos/go-estimate.svg) - State estimation and filtering algorithms in Go.
- [go-gt](https://github.com/ThePaw/go-gt) ![](https://img.shields.io/github/stars/ThePaw/go-gt.svg) - Graph theory algorithms written in "Go" language.
- [go-hep](https://github.com/go-hep/hep) ![](https://img.shields.io/github/stars/go-hep/hep.svg) - A set of libraries and tools for performing High Energy Physics analyses with ease.
- [godesim](https://github.com/soypat/godesim) ![](https://img.shields.io/github/stars/soypat/godesim.svg) - Extended/multivariable ODE solver framework for event-based simulations with simple API.
- [goent](https://github.com/kzahedi/goent) ![](https://img.shields.io/github/stars/kzahedi/goent.svg) - GO Implementation of Entropy Measures.
- [gograph](https://github.com/hmdsefi/gograph) ![](https://img.shields.io/github/stars/hmdsefi/gograph.svg) - A golang generic graph library that provides mathematical graph-theory and algorithms.
- [gonum](https://github.com/gonum/gonum) ![](https://img.shields.io/github/stars/gonum/gonum.svg) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.
- [gonum/plot](https://github.com/gonum/plot) ![](https://img.shields.io/github/stars/gonum/plot.svg) - gonum/plot provides an API for building and drawing plots in Go.
- [goraph](https://github.com/gyuho/goraph) ![](https://img.shields.io/github/stars/gyuho/goraph.svg) - Pure Go graph theory library(data structure, algorithm visualization).
- [gosl](https://github.com/cpmech/gosl) ![](https://img.shields.io/github/stars/cpmech/gosl.svg) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.
- [GoStats](https://github.com/OGFris/GoStats) ![](https://img.shields.io/github/stars/OGFris/GoStats.svg) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.
- [graph](https://github.com/yourbasic/graph) ![](https://img.shields.io/github/stars/yourbasic/graph.svg) - Library of basic graph algorithms.
- [jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) ![](https://img.shields.io/github/stars/nikolaydubina/jsonl-graph.svg) - Tool to manipulate JSONL graphs with graphviz support.
- [ode](https://github.com/ChristopherRabotin/ode) ![](https://img.shields.io/github/stars/ChristopherRabotin/ode.svg) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.
- [orb](https://github.com/paulmach/orb) ![](https://img.shields.io/github/stars/paulmach/orb.svg) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.
- [pagerank](https://github.com/alixaxel/pagerank) ![](https://img.shields.io/github/stars/alixaxel/pagerank.svg) - Weighted PageRank algorithm implemented in Go.
- [piecewiselinear](https://github.com/sgreben/piecewiselinear) ![](https://img.shields.io/github/stars/sgreben/piecewiselinear.svg) - Tiny linear interpolation library.
- [PiHex](https://github.com/claygod/PiHex) ![](https://img.shields.io/github/stars/claygod/PiHex.svg) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.
- [Poly](https://github.com/bebop/poly) ![](https://img.shields.io/github/stars/bebop/poly.svg) - A Go package for engineering organisms.
- [rootfinding](https://github.com/khezen/rootfinding) ![](https://img.shields.io/github/stars/khezen/rootfinding.svg) - root-finding algorithms library for finding roots of quadratic functions.
- [sparse](https://github.com/james-bowman/sparse) ![](https://img.shields.io/github/stars/james-bowman/sparse.svg) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.
- [stats](https://github.com/montanaflynn/stats) ![](https://img.shields.io/github/stars/montanaflynn/stats.svg) - Statistics package with common functions missing from the Golang standard library.
- [streamtools](https://github.com/nytlabs/streamtools) ![](https://img.shields.io/github/stars/nytlabs/streamtools.svg) - general purpose, graphical tool for dealing with streams of data.
- [TextRank](https://github.com/DavidBelicza/TextRank) ![](https://img.shields.io/github/stars/DavidBelicza/TextRank.svg) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.
- [topk](https://github.com/keilerkonzept/topk) ![](https://img.shields.io/github/stars/keilerkonzept/topk.svg) - Sliding-window and regular top-K sketches, based on the HeavyKeeper algorithm.
- [triangolatte](https://github.com/tchayen/triangolatte) ![](https://img.shields.io/github/stars/tchayen/triangolatte.svg) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [acmetool](https://github.com/hlandau/acme) ![](https://img.shields.io/github/stars/hlandau/acme.svg) - ACME (Let's Encrypt) client tool with automatic renewal.
- [acopw-go](https://sr.ht/~jamesponddotco/acopw-go/) - Small cryptographically secure password generator package for Go.
- [acra](https://github.com/cossacklabs/acra) ![](https://img.shields.io/github/stars/cossacklabs/acra.svg) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.
- [age](https://github.com/FiloSottile/age) ![](https://img.shields.io/github/stars/FiloSottile/age.svg) - A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability.
- [argon2-hashing](https://github.com/andskur/argon2-hashing) ![](https://img.shields.io/github/stars/andskur/argon2-hashing.svg) - light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package.
- [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.
- [BadActor](https://github.com/jaredfolkins/badactor) ![](https://img.shields.io/github/stars/jaredfolkins/badactor.svg) - In-memory, application-driven jailer built in the spirit of fail2ban.
- [beelzebub](https://github.com/mariocandela/beelzebub) ![](https://img.shields.io/github/stars/mariocandela/beelzebub.svg) - A secure low code honeypot framework, leveraging AI for System Virtualization.
- [booster](https://github.com/anatol/booster) ![](https://img.shields.io/github/stars/anatol/booster.svg) - Fast initramfs generator with full-disk encryption support.
- [Cameradar](https://github.com/Ullaakut/cameradar) ![](https://img.shields.io/github/stars/Ullaakut/cameradar.svg) - Tool and library to remotely hack RTSP streams from surveillance cameras.
- [certificates](https://github.com/mvmaasakkers/certificates) ![](https://img.shields.io/github/stars/mvmaasakkers/certificates.svg) - An opinionated tool for generating tls certificates.
- [CertMagic](https://github.com/caddyserver/certmagic) ![](https://img.shields.io/github/stars/caddyserver/certmagic.svg) - Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal.
- [Coraza](https://github.com/corazawaf/coraza) ![](https://img.shields.io/github/stars/corazawaf/coraza.svg) - Enterprise-ready, modsecurity and OWASP CRS compatible WAF library.
- [dongle](https://github.com/golang-module/dongle) ![](https://img.shields.io/github/stars/golang-module/dongle.svg) - A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption.
- [encid](https://github.com/bobg/encid) ![](https://img.shields.io/github/stars/bobg/encid.svg) - Encode and decode encrypted integer IDs.
- [entpassgen](https://github.com/andreimerlescu/entpassgen) ![](https://img.shields.io/github/stars/andreimerlescu/entpassgen.svg) - Entropy Password Generator with extensive command line arguments to generate random strings securely including digits, passwords, and passwords built using obscure dictionary words mixed with symbols and digits.
- [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) ![](https://img.shields.io/github/stars/prashantgupta24/firewalld-rest.svg) - A rest application to dynamically update firewalld rules on a linux server.
- [go-generate-password](https://github.com/m1/go-generate-password) ![](https://img.shields.io/github/stars/m1/go-generate-password.svg) - Password generator that can be used on the cli or as a library.
- [go-htpasswd](https://github.com/tg123/go-htpasswd) ![](https://img.shields.io/github/stars/tg123/go-htpasswd.svg) - Apache htpasswd Parser for Go.
- [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) ![](https://img.shields.io/github/stars/lane-c-wagner/go-password-validator.svg) - Password validator based on raw cryptographic entropy values.
- [go-peer](https://github.com/number571/go-peer) ![](https://img.shields.io/github/stars/number571/go-peer.svg) - A software library for creating secure and anonymous decentralized systems.
- [go-yara](https://github.com/hillu/go-yara) ![](https://img.shields.io/github/stars/hillu/go-yara.svg) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".
- [goArgonPass](https://github.com/dwin/goArgonPass) ![](https://img.shields.io/github/stars/dwin/goArgonPass.svg) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.
- [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) ![](https://img.shields.io/github/stars/dwin/goSecretBoxPassword.svg) - A probably paranoid package for securely hashing and encrypting passwords.
- [Interpol](https://github.com/avahidi/interpol) ![](https://img.shields.io/github/stars/avahidi/interpol.svg) - Rule-based data generator for fuzzing and penetration testing.
- [lego](https://github.com/go-acme/lego) ![](https://img.shields.io/github/stars/go-acme/lego.svg) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt).
- [luks.go](https://github.com/anatol/luks.go) ![](https://img.shields.io/github/stars/anatol/luks.go.svg) - Pure Golang library to manage LUKS partitions.
- [memguard](https://github.com/awnumar/memguard) ![](https://img.shields.io/github/stars/awnumar/memguard.svg) - A pure Go library for handling sensitive values in memory.
- [multikey](https://github.com/adrianosela/multikey) ![](https://img.shields.io/github/stars/adrianosela/multikey.svg) - An n-out-of-N keys encryption/decryption framework based on Shamir's Secret Sharing algorithm.
- [nacl](https://github.com/kevinburke/nacl) ![](https://img.shields.io/github/stars/kevinburke/nacl.svg) - Go implementation of the NaCL set of API's.
- [optimus-go](https://github.com/pjebs/optimus-go) ![](https://img.shields.io/github/stars/pjebs/optimus-go.svg) - ID hashing and Obfuscation using Knuth's Algorithm.
- [passlib](https://github.com/hlandau/passlib) ![](https://img.shields.io/github/stars/hlandau/passlib.svg) - Futureproof password hashing library.
- [passwap](https://github.com/zitadel/passwap) ![](https://img.shields.io/github/stars/zitadel/passwap.svg) - Provides a unified implementation between different password hashing algorithms
- [qrand](https://github.com/bitfield/qrand) ![](https://img.shields.io/github/stars/bitfield/qrand.svg) - Client for the ANU Quantum Numbers (AQN) API, providing quantum-mechanically secure random data.
- [SafeDep/vet](https://github.com/safedep/vet) ![](https://img.shields.io/github/stars/safedep/vet.svg) - Protect against malicious open source packages.
- [secret](https://github.com/rsjethani/secret) ![](https://img.shields.io/github/stars/rsjethani/secret.svg) - Prevent your secrets from leaking into logs, std\* etc.
- [secure](https://github.com/unrolled/secure) ![](https://img.shields.io/github/stars/unrolled/secure.svg) - HTTP middleware for Go that facilitates some quick security wins.
- [secureio](https://github.com/xaionaro-go/secureio) ![](https://img.shields.io/github/stars/xaionaro-go/secureio.svg) - An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519.
- [simple-scrypt](https://github.com/elithrar/simple-scrypt) ![](https://img.shields.io/github/stars/elithrar/simple-scrypt.svg) - Scrypt package with a simple, obvious API and automatic cost calibration built-in.
- [ssh-vault](https://github.com/ssh-vault/ssh-vault) ![](https://img.shields.io/github/stars/ssh-vault/ssh-vault.svg) - encrypt/decrypt using ssh keys.
- [sslmgr](https://github.com/adrianosela/sslmgr) ![](https://img.shields.io/github/stars/adrianosela/sslmgr.svg) - SSL certificates made easy with a high level wrapper around acme/autocert.
- [teler-waf](https://github.com/kitabisa/teler-waf) ![](https://img.shields.io/github/stars/kitabisa/teler-waf.svg) - teler-waf is a Go HTTP middleware that provide teler IDS functionality to protect against web-based attacks and improve the security of Go-based web applications. It is highly configurable and easy to integrate into existing Go applications.
- [themis](https://github.com/cossacklabs/themis) ![](https://img.shields.io/github/stars/cossacklabs/themis.svg) - high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps.
- [urusai](https://github.com/calpa/urusai) ![](https://img.shields.io/github/stars/calpa/urusai.svg) - Urusai ("noisy" in Japanese) is a Go implementation of a random HTTP/DNS traffic noise generator that helps protect privacy by creating digital smokescreens while browsing.

**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [bambam](https://github.com/glycerine/bambam) ![](https://img.shields.io/github/stars/glycerine/bambam.svg) - generator for Cap'n Proto schemas from go.
- [bel](https://github.com/32leaves/bel) ![](https://img.shields.io/github/stars/32leaves/bel.svg) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.
- [binstruct](https://github.com/ghostiam/binstruct) ![](https://img.shields.io/github/stars/ghostiam/binstruct.svg) - Golang binary decoder for mapping data into the structure.
- [cbor](https://github.com/fxamacker/cbor) ![](https://img.shields.io/github/stars/fxamacker/cbor.svg) - Small, safe, and easy CBOR encoding and decoding library.
- [colfer](https://github.com/pascaldekloe/colfer) ![](https://img.shields.io/github/stars/pascaldekloe/colfer.svg) - Code generation for the Colfer binary format.
- [csvutil](https://github.com/jszwec/csvutil) ![](https://img.shields.io/github/stars/jszwec/csvutil.svg) - High Performance, idiomatic CSV record encoding and decoding to native Go structures.
- [elastic](https://github.com/epiclabs-io/elastic) ![](https://img.shields.io/github/stars/epiclabs-io/elastic.svg) - Convert slices, maps or any other unknown value across different types at run-time, no matter what.
- [fixedwidth](https://github.com/huydang284/fixedwidth) ![](https://img.shields.io/github/stars/huydang284/fixedwidth.svg) - Fixed-width text formatting (UTF-8 supported).
- [fwencoder](https://github.com/o1egl/fwencoder) ![](https://img.shields.io/github/stars/o1egl/fwencoder.svg) - Fixed width file parser (encoding and decoding library) for Go.
- [go-capnproto](https://github.com/glycerine/go-capnproto) ![](https://img.shields.io/github/stars/glycerine/go-capnproto.svg) - Cap'n Proto library and parser for go.
- [go-codec](https://github.com/ugorji/go) ![](https://img.shields.io/github/stars/ugorji/go.svg) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.
- [go-csvlib](https://github.com/tiendc/go-csvlib) ![](https://img.shields.io/github/stars/tiendc/go-csvlib.svg) - High level and rich functionalities CSV serialization/deserialization library.
- [goprotobuf](https://github.com/golang/protobuf) ![](https://img.shields.io/github/stars/golang/protobuf.svg) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.
- [gotiny](https://github.com/raszia/gotiny) ![](https://img.shields.io/github/stars/raszia/gotiny.svg) - Efficient Go serialization library, gotiny is almost as fast as serialization libraries that generate code.
- [jsoniter](https://github.com/json-iterator/go) ![](https://img.shields.io/github/stars/json-iterator/go.svg) - High-performance 100% compatible drop-in replacement of "encoding/json".
- [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) ![](https://img.shields.io/github/stars/yvasiyarov/php_session_decoder.svg) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.
- [pletter](https://github.com/vimeda/pletter) ![](https://img.shields.io/github/stars/vimeda/pletter.svg) - A standard way to wrap a proto message for message brokers.
- [structomap](https://github.com/tuvistavie/structomap) ![](https://img.shields.io/github/stars/tuvistavie/structomap.svg) - Library to easily and dynamically generate maps from static structures.
- [unitpacking](https://github.com/recolude/unitpacking) ![](https://img.shields.io/github/stars/recolude/unitpacking.svg) - Library to pack unit vectors into as fewest bytes as possible.

**[⬆ back to top](#contents)**

## Server Applications

- [algernon](https://github.com/xyproto/algernon) ![](https://img.shields.io/github/stars/xyproto/algernon.svg) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
- [Caddy](https://github.com/caddyserver/caddy) ![](https://img.shields.io/github/stars/caddyserver/caddy.svg) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
- [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
- [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) ![](https://img.shields.io/github/stars/blind-oracle/cortex-tenant.svg) - Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels.
- [devd](https://github.com/cortesi/devd) ![](https://img.shields.io/github/stars/cortesi/devd.svg) - Local webserver for developers.
- [discovery](https://github.com/Bilibili/discovery) ![](https://img.shields.io/github/stars/Bilibili/discovery.svg) - A registry for resilient mid-tier load balancing and failover.
- [dudeldu](https://github.com/krotik/dudeldu) ![](https://img.shields.io/github/stars/krotik/dudeldu.svg) - A simple SHOUTcast server.
- [Easegress](https://github.com/megaease/easegress) ![](https://img.shields.io/github/stars/megaease/easegress.svg) - A cloud native high availability/performance traffic orchestration system with observability and extensibility.
- [Engity's Bifröst](https://bifroest.engity.org/) - Highly customizable SSH server with several ways to authorize a user how to execute its session (local or in containers).
- [etcd](https://github.com/etcd-io/etcd) ![](https://img.shields.io/github/stars/etcd-io/etcd.svg) - Highly-available key value store for shared configuration and service discovery.
- [Euterpe](https://github.com/ironsmile/euterpe) ![](https://img.shields.io/github/stars/ironsmile/euterpe.svg) - Self-hosted music streaming server with built-in web UI and REST API.
- [Fider](https://github.com/getfider/fider) ![](https://img.shields.io/github/stars/getfider/fider.svg) - Fider is an open platform to collect and organize customer feedback.
- [Flagr](https://github.com/checkr/flagr) ![](https://img.shields.io/github/stars/checkr/flagr.svg) - Flagr is an open-source feature flagging and A/B testing service.
- [flipt](https://github.com/markphelps/flipt) ![](https://img.shields.io/github/stars/markphelps/flipt.svg) - A self contained feature flag solution written in Go and Vue.js
- [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) ![](https://img.shields.io/github/stars/thomaspoignant/go-feature-flag.svg) - A simple, complete and lightweight self-hosted feature flag solution 100% Open Source.
- [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) ![](https://img.shields.io/github/stars/fabiocicerchia/go-proxy-cache.svg) - Simple Reverse Proxy with Caching, written in Go, using Redis.
- [gondola](https://github.com/bmf-san/gondola) ![](https://img.shields.io/github/stars/bmf-san/gondola.svg) - A YAML based golang reverse proxy.
- [lets-proxy2](https://github.com/rekby/lets-proxy2) ![](https://img.shields.io/github/stars/rekby/lets-proxy2.svg) - Reverse proxy for handle https with issue certificates in fly from lets-encrypt.
- [minio](https://github.com/minio/minio) ![](https://img.shields.io/github/stars/minio/minio.svg) - Minio is a distributed object storage server.
- [Moxy](https://github.com/sinhashubham95/moxy) ![](https://img.shields.io/github/stars/sinhashubham95/moxy.svg) - Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint.
- [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) ![](https://img.shields.io/github/stars/blind-oracle/nginx-prometheus.svg) - Nginx log parser and exporter to Prometheus.
- [nsq](https://nsq.io/) - A realtime distributed messaging platform.
- [OpenRun](https://github.com/openrundev/openrun) ![](https://img.shields.io/github/stars/openrundev/openrun.svg) - Open-source alternative to Google Cloud Run and AWS App Runner. Easily deploy internal tools across a team.
- [pocketbase](https://github.com/pocketbase/pocketbase) ![](https://img.shields.io/github/stars/pocketbase/pocketbase.svg) - PocketBase is a realtime backend in 1 file consisting of embedded database (SQLite) with realtime subscriptions, built-in auth management and much more.
- [protoxy](https://github.com/camgraff/protoxy) ![](https://img.shields.io/github/stars/camgraff/protoxy.svg) - A proxy server that converts JSON request bodies to Protocol Buffers.
- [psql-streamer](https://github.com/blind-oracle/psql-streamer) ![](https://img.shields.io/github/stars/blind-oracle/psql-streamer.svg) - Stream database events from PostgreSQL to Kafka.
- [riemann-relay](https://github.com/blind-oracle/riemann-relay) ![](https://img.shields.io/github/stars/blind-oracle/riemann-relay.svg) - Relay to load-balance Riemann events and/or convert them to Carbon.
- [RoadRunner](https://github.com/spiral/roadrunner) ![](https://img.shields.io/github/stars/spiral/roadrunner.svg) - High-performance PHP application server, load-balancer and process manager.
- [SFTPGo](https://github.com/drakkan/sftpgo) ![](https://img.shields.io/github/stars/drakkan/sftpgo.svg) - Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage.
- [Trickster](https://github.com/tricksterproxy/trickster) ![](https://img.shields.io/github/stars/tricksterproxy/trickster.svg) - HTTP reverse proxy cache and time series accelerator.
- [wd-41](https://github.com/baalimago/wd-41) ![](https://img.shields.io/github/stars/baalimago/wd-41.svg) - A (w)eb (d)evelopment server with automatic live-reload on file changes.
- [Wish](https://github.com/charmbracelet/wish) ![](https://img.shields.io/github/stars/charmbracelet/wish.svg) - Make SSH apps, just like that!

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-etl](https://github.com/Breeze0806/go-etl) ![](https://img.shields.io/github/stars/Breeze0806/go-etl.svg) - A lightweight toolkit for data source extraction, transformation, and loading (ETL).
- [go-streams](https://github.com/reugn/go-streams) ![](https://img.shields.io/github/stars/reugn/go-streams.svg) - Go stream processing library.
- [goio](https://github.com/primetalk/goio) ![](https://img.shields.io/github/stars/primetalk/goio.svg) - An implementation of IO, Stream, Fiber for Golang, inspired by awesome Scala libraries cats and fs2.
- [gostream](https://github.com/mariomac/gostream) ![](https://img.shields.io/github/stars/mariomac/gostream.svg) - Type-safe stream processing library inspired by the Java Streams API.
- [machine](https://github.com/whitaker-io/machine) ![](https://img.shields.io/github/stars/whitaker-io/machine.svg) - Go library for writing and generating stream workers with built in metrics and traceability.
- [nibbler](https://github.com/naughtygopher/nibbler) ![](https://img.shields.io/github/stars/naughtygopher/nibbler.svg) - A lightweight package for micro batch processing.
- [ro](https://github.com/samber/ro) ![](https://img.shields.io/github/stars/samber/ro.svg) - Reactive Programming: declarative and composable API for event-driven applications.
- [stream](https://github.com/youthlin/stream) ![](https://img.shields.io/github/stars/youthlin/stream.svg) - Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce...
- [StreamSQL](https://github.com/rulego/streamsql) ![](https://img.shields.io/github/stars/rulego/streamsql.svg) - A lightweight streaming SQL engine for real-time data processing.

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [ego](https://github.com/benbjohnson/ego) ![](https://img.shields.io/github/stars/benbjohnson/ego.svg) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
- [extemplate](https://git.sr.ht/~dvko/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance.
- [fasttemplate](https://github.com/valyala/fasttemplate) ![](https://img.shields.io/github/stars/valyala/fasttemplate.svg) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/).
- [gomponents](https://www.gomponents.com) - HTML 5 components in pure Go, that look something like this: `func(name string) g.Node { return Div(Class("headline"), g.Textf("Hi %v!", name)) }`.
- [got](https://github.com/goradd/got) ![](https://img.shields.io/github/stars/goradd/got.svg) - A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more.
- [goview](https://github.com/foolin/goview) ![](https://img.shields.io/github/stars/foolin/goview.svg) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
- [htmgo](https://htmgo.dev) - build simple and scalable systems with go + htmx
- [jet](https://github.com/CloudyKit/jet) ![](https://img.shields.io/github/stars/CloudyKit/jet.svg) - Jet template engine.
- [liquid](https://github.com/osteele/liquid) ![](https://img.shields.io/github/stars/osteele/liquid.svg) - Go implementation of Shopify Liquid templates.
- [maroto](https://github.com/johnfercher/maroto) ![](https://img.shields.io/github/stars/johnfercher/maroto.svg) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.
- [pongo2](https://github.com/flosch/pongo2) ![](https://img.shields.io/github/stars/flosch/pongo2.svg) - Django-like template-engine for Go.
- [quicktemplate](https://github.com/valyala/quicktemplate) ![](https://img.shields.io/github/stars/valyala/quicktemplate.svg) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.
- [Razor](https://github.com/sipin/gorazor) ![](https://img.shields.io/github/stars/sipin/gorazor.svg) - Razor view engine for Golang.
- [Soy](https://github.com/robfig/soy) ![](https://img.shields.io/github/stars/robfig/soy.svg) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).
- [sprout](https://github.com/go-sprout/sprout) ![](https://img.shields.io/github/stars/go-sprout/sprout.svg) - Useful template functions for Go templates.
- [tbd](https://github.com/lucasepe/tbd) ![](https://img.shields.io/github/stars/lucasepe/tbd.svg) - A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata.
- [templ](https://github.com/a-h/templ) ![](https://img.shields.io/github/stars/a-h/templ.svg) - A HTML templating language that has great developer tooling.
- [templator](https://github.com/alesr/templator) ![](https://img.shields.io/github/stars/alesr/templator.svg) - A type-safe HTML template rendering engine for Go.

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

### Testing Frameworks

- [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
- [arch-go](https://github.com/arch-go/arch-go) ![](https://img.shields.io/github/stars/arch-go/arch-go.svg) - Architecture testing tool for Go projects.
- [assert](https://github.com/go-playground/assert) ![](https://img.shields.io/github/stars/go-playground/assert.svg) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions.
- [baloo](https://github.com/h2non/baloo) ![](https://img.shields.io/github/stars/h2non/baloo.svg) - Expressive and versatile end-to-end HTTP API testing made easy.
- [be](https://github.com/carlmjohnson/be) ![](https://img.shields.io/github/stars/carlmjohnson/be.svg) - The minimalist generic test assertion library.
- [biff](https://github.com/fulldump/biff) ![](https://img.shields.io/github/stars/fulldump/biff.svg) - Bifurcation testing framework, BDD compatible.
- [charlatan](https://github.com/percolate/charlatan) ![](https://img.shields.io/github/stars/percolate/charlatan.svg) - Tool to generate fake interface implementations for tests.
- [commander](https://github.com/SimonBaeumer/commander) ![](https://img.shields.io/github/stars/SimonBaeumer/commander.svg) - Tool for testing cli applications on windows, linux and osx.
- [cupaloy](https://github.com/bradleyjkemp/cupaloy) ![](https://img.shields.io/github/stars/bradleyjkemp/cupaloy.svg) - Simple snapshot testing addon for your test framework.
- [dbcleaner](https://github.com/khaiql/dbcleaner) ![](https://img.shields.io/github/stars/khaiql/dbcleaner.svg) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby.
- [dft](https://github.com/abecodes/dft) ![](https://img.shields.io/github/stars/abecodes/dft.svg) - Lightweight, zero dependency docker containers for testing (or more).
- [dsunit](https://github.com/viant/dsunit) ![](https://img.shields.io/github/stars/viant/dsunit.svg) - Datastore testing for SQL, NoSQL, structured files.
- [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) ![](https://img.shields.io/github/stars/fergusstrange/embedded-postgres.svg) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.
- [endly](https://github.com/viant/endly) ![](https://img.shields.io/github/stars/viant/endly.svg) - Declarative end to end functional testing.
- [envite](https://github.com/PerimeterX/envite) ![](https://img.shields.io/github/stars/PerimeterX/envite.svg) - Dev and testing environment management framework.
- [fixenv](https://github.com/rekby/fixenv) ![](https://img.shields.io/github/stars/rekby/fixenv.svg) - Fixture manage engine, inspired by pytest fixtures.
- [flute](https://github.com/suzuki-shunsuke/flute) ![](https://img.shields.io/github/stars/suzuki-shunsuke/flute.svg) - HTTP client testing framework.
- [frisby](https://github.com/verdverm/frisby) ![](https://img.shields.io/github/stars/verdverm/frisby.svg) - REST API testing framework.
- [gherkingen](https://github.com/hedhyw/gherkingen) ![](https://img.shields.io/github/stars/hedhyw/gherkingen.svg) - BDD boilerplate generator and framework.
- [ginkgo](https://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
- [gnomock](https://github.com/orlangure/gnomock) ![](https://img.shields.io/github/stars/orlangure/gnomock.svg) - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks.
- [go-carpet](https://github.com/msoap/go-carpet) ![](https://img.shields.io/github/stars/msoap/go-carpet.svg) - Tool for viewing test coverage in terminal.
- [go-cmp](https://github.com/google/go-cmp) ![](https://img.shields.io/github/stars/google/go-cmp.svg) - Package for comparing Go values in tests.
- [go-hit](https://github.com/Eun/go-hit) ![](https://img.shields.io/github/stars/Eun/go-hit.svg) - Hit is an http integration test framework written in golang.
- [go-httpbin](https://github.com/mccutchen/go-httpbin) ![](https://img.shields.io/github/stars/mccutchen/go-httpbin.svg) - HTTP testing and debugging tool with various endpoints for client testing.
- [go-mutesting](https://github.com/zimmski/go-mutesting) ![](https://img.shields.io/github/stars/zimmski/go-mutesting.svg) - Mutation testing for Go source code.
- [go-mysql-test-container](https://github.com/arikama/go-mysql-test-container) ![](https://img.shields.io/github/stars/arikama/go-mysql-test-container.svg) - Golang MySQL testcontainer to help with MySQL integration testing.
- [go-snaps](http://github.com/gkampitakis/go-snaps) - Jest-like snapshot testing in Golang.
- [go-test-coverage](https://github.com/vladopajic/go-test-coverage) ![](https://img.shields.io/github/stars/vladopajic/go-test-coverage.svg) - Tool that reports coverage of files below set threshold.
- [go-testdeep](https://github.com/maxatome/go-testdeep) ![](https://img.shields.io/github/stars/maxatome/go-testdeep.svg) - Extremely flexible golang deep comparison, extends the go testing package.
- [go-testpredicate](https://github.com/maargenton/go-testpredicate) ![](https://img.shields.io/github/stars/maargenton/go-testpredicate.svg) - Test predicate style assertions library with extensive diagnostics output.
- [go-vcr](https://github.com/dnaeon/go-vcr) ![](https://img.shields.io/github/stars/dnaeon/go-vcr.svg) - Record and replay your HTTP interactions for fast, deterministic and accurate tests.
- [goblin](https://github.com/franela/goblin) ![](https://img.shields.io/github/stars/franela/goblin.svg) - Mocha like testing framework of Go.
- [goc](https://github.com/qiniu/goc) ![](https://img.shields.io/github/stars/qiniu/goc.svg) - Goc is a comprehensive coverage testing system for The Go Programming Language.
- [gocheck](https://labix.org/gocheck) - More advanced testing framework alternative to gotest.
- [GoConvey](https://github.com/smartystreets/goconvey/) ![](https://img.shields.io/github/stars/smartystreets/goconvey.svg) - BDD-style framework with web UI and live reload.
- [gocrest](https://github.com/corbym/gocrest) ![](https://img.shields.io/github/stars/corbym/gocrest.svg) - Composable hamcrest-like matchers for Go assertions.
- [godog](https://github.com/cucumber/godog) ![](https://img.shields.io/github/stars/cucumber/godog.svg) - Cucumber BDD framework for Go.
- [gofight](https://github.com/appleboy/gofight) ![](https://img.shields.io/github/stars/appleboy/gofight.svg) - API Handler Testing for Golang Router framework.
- [gogiven](https://github.com/corbym/gogiven) ![](https://img.shields.io/github/stars/corbym/gogiven.svg) - YATSPEC-like BDD testing framework for Go.
- [gomatch](https://github.com/jfilipczyk/gomatch) ![](https://img.shields.io/github/stars/jfilipczyk/gomatch.svg) - library created for testing JSON against patterns.
- [gomega](https://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
- [Gont](https://github.com/stv0g/gont) ![](https://img.shields.io/github/stars/stv0g/gont.svg) - Go network testing toolkit for testing building complex network topologies using Linux namespaces.
- [gospecify](https://github.com/stesla/gospecify) ![](https://img.shields.io/github/stars/stesla/gospecify.svg) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
- [gosuite](https://github.com/pavlo/gosuite) ![](https://img.shields.io/github/stars/pavlo/gosuite.svg) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.
- [got](https://github.com/ysmood/got) ![](https://img.shields.io/github/stars/ysmood/got.svg) - An enjoyable golang test framework.
- [gotest.tools](https://github.com/gotestyourself/gotest.tools) ![](https://img.shields.io/github/stars/gotestyourself/gotest.tools.svg) - A collection of packages to augment the go testing package and support common patterns.
- [Hamcrest](https://github.com/rdrdr/hamcrest) ![](https://img.shields.io/github/stars/rdrdr/hamcrest.svg) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
- [httpexpect](https://github.com/gavv/httpexpect) ![](https://img.shields.io/github/stars/gavv/httpexpect.svg) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing.
- [is](https://github.com/matryer/is) ![](https://img.shields.io/github/stars/matryer/is.svg) - Professional lightweight testing mini-framework for Go.
- [jsonassert](https://github.com/kinbiko/jsonassert) ![](https://img.shields.io/github/stars/kinbiko/jsonassert.svg) - Package for verifying that your JSON payloads are serialized correctly.
- [keploy](https://github.com/keploy/keploy) ![](https://img.shields.io/github/stars/keploy/keploy.svg) - Generate Testcase and Data Mocks from API calls automatically.
- [omg.testingtools](https://github.com/dedalqq/omg.testingtools) ![](https://img.shields.io/github/stars/dedalqq/omg.testingtools.svg) - The simple library for change a values of private fields for testing.
- [restit](https://github.com/yookoala/restit) ![](https://img.shields.io/github/stars/yookoala/restit.svg) - Go micro framework to help writing RESTful API integration test.
- [schema](https://github.com/jgroeneveld/schema) ![](https://img.shields.io/github/stars/jgroeneveld/schema.svg) - Quick and easy expression matching for JSON schemas used in requests and responses.
- [stop-and-go](https://github.com/elgohr/stop-and-go) ![](https://img.shields.io/github/stars/elgohr/stop-and-go.svg) - Testing helper for concurrency.
- [testcase](https://github.com/adamluzsi/testcase) ![](https://img.shields.io/github/stars/adamluzsi/testcase.svg) - Idiomatic testing framework for Behavior Driven Development.
- [testcerts](https://github.com/madflojo/testcerts) ![](https://img.shields.io/github/stars/madflojo/testcerts.svg) - Dynamically generate self-signed certificates and certificate authorities within your test functions.
- [testcontainers-go](https://github.com/testcontainers/testcontainers-go) ![](https://img.shields.io/github/stars/testcontainers/testcontainers-go.svg) - A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.
- [testfixtures](https://github.com/go-testfixtures/testfixtures) ![](https://img.shields.io/github/stars/go-testfixtures/testfixtures.svg) - A helper for Rails' like test fixtures to test database applications.
- [Testify](https://github.com/stretchr/testify) ![](https://img.shields.io/github/stars/stretchr/testify.svg) - Sacred extension to the standard go testing package.
- [testsql](https://github.com/zhulongcheng/testsql) ![](https://img.shields.io/github/stars/zhulongcheng/testsql.svg) - Generate test data from SQL files before testing and clear it after finished.
- [testza](https://github.com/MarvinJWendt/testza) ![](https://img.shields.io/github/stars/MarvinJWendt/testza.svg) - Full-featured test framework with nice colorized output.
- [trial](https://github.com/jgroeneveld/trial) ![](https://img.shields.io/github/stars/jgroeneveld/trial.svg) - Quick and easy extendable assertions without introducing much boilerplate.
- [Tt](https://github.com/vcaesar/tt) ![](https://img.shields.io/github/stars/vcaesar/tt.svg) - Simple and colorful test tools.
- [wstest](https://github.com/posener/wstest) ![](https://img.shields.io/github/stars/posener/wstest.svg) - Websocket client for unit-testing a websocket http.Handler.

### Mock

- [connexions](https://github.com/cubahno/connexions) ![](https://img.shields.io/github/stars/cubahno/connexions.svg) - Combine multiple APIs with meaningful responses, configurable latency and error codes based on OpenAPI 3.0 specifications and files.
- [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) ![](https://img.shields.io/github/stars/maxbrunsfeld/counterfeiter.svg) - Tool for generating self-contained mock objects.
- [genmock](https://gitlab.com/so_literate/genmock) - Go mocking system with code generator for building calls of the interface methods.
- [go-localstack](https://github.com/elgohr/go-localstack) ![](https://img.shields.io/github/stars/elgohr/go-localstack.svg) - Tool for using localstack in AWS testing.
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) ![](https://img.shields.io/github/stars/DATA-DOG/go-sqlmock.svg) - Mock SQL driver for testing database interactions.
- [go-txdb](https://github.com/DATA-DOG/go-txdb) ![](https://img.shields.io/github/stars/DATA-DOG/go-txdb.svg) - Single transaction based database driver mainly for testing purposes.
- [gock](https://github.com/h2non/gock) ![](https://img.shields.io/github/stars/h2non/gock.svg) - Versatile HTTP mocking made easy.
- [gomock](https://github.com/uber-go/mock) ![](https://img.shields.io/github/stars/uber-go/mock.svg) - Mocking framework for the Go programming language.
- [govcr](https://github.com/seborama/govcr) ![](https://img.shields.io/github/stars/seborama/govcr.svg) - HTTP mock for Golang: record and replay HTTP interactions for offline testing.
- [hoverfly](https://github.com/SpectoLabs/hoverfly) ![](https://img.shields.io/github/stars/SpectoLabs/hoverfly.svg) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.
- [httpmock](https://github.com/jarcoal/httpmock) ![](https://img.shields.io/github/stars/jarcoal/httpmock.svg) - Easy mocking of HTTP responses from external resources.
- [minimock](https://github.com/gojuno/minimock) ![](https://img.shields.io/github/stars/gojuno/minimock.svg) - Mock generator for Go interfaces.
- [mockery](https://github.com/vektra/mockery) ![](https://img.shields.io/github/stars/vektra/mockery.svg) - Tool to generate Go interfaces.
- [mockfs](https://github.com/balinomad/go-mockfs) ![](https://img.shields.io/github/stars/balinomad/go-mockfs.svg) - Mock filesystem for Go testing with error injection and latency simulation, built on `testing/fstest.MapFS`.
- [mockhttp](https://github.com/tv42/mockhttp) ![](https://img.shields.io/github/stars/tv42/mockhttp.svg) - Mock object for Go http.ResponseWriter.
- [mooncake](https://github.com/GuilhermeCaruso/mooncake) ![](https://img.shields.io/github/stars/GuilhermeCaruso/mooncake.svg) - A simple way to generate mocks for multiple purposes.
- [moq](https://github.com/matryer/moq) ![](https://img.shields.io/github/stars/matryer/moq.svg) - Utility that generates a struct from any interface. The struct can be used in test code as a mock of the interface.
- [moxie](https://lesiw.io/moxie) - Generate mock methods on embedded structs.
- [pgxmock](https://github.com/pashagolub/pgxmock) ![](https://img.shields.io/github/stars/pashagolub/pgxmock.svg) - A mock library implementing [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx/).
- [timex](https://github.com/cabify/timex) ![](https://img.shields.io/github/stars/cabify/timex.svg) - A test-friendly replacement for the native `time` package.
- [xgo](https://github.com/xhd2015/xgo) ![](https://img.shields.io/github/stars/xhd2015/xgo.svg) - A general pureposed function mocking library.

### Fuzzing and delta-debugging/reducing/shrinking

- [go-fuzz](https://github.com/dvyukov/go-fuzz) ![](https://img.shields.io/github/stars/dvyukov/go-fuzz.svg) - Randomized testing system.
- [Tavor](https://github.com/zimmski/tavor) ![](https://img.shields.io/github/stars/zimmski/tavor.svg) - Generic fuzzing and delta-debugging framework.

### Selenium and browser control tools

- [cdp](https://github.com/mafredri/cdp) ![](https://img.shields.io/github/stars/mafredri/cdp.svg) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.
- [chromedp](https://github.com/knq/chromedp) ![](https://img.shields.io/github/stars/knq/chromedp.svg) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.
- [playwright-go](https://github.com/mxschmitt/playwright-go) ![](https://img.shields.io/github/stars/mxschmitt/playwright-go.svg) - browser automation library to control Chromium, Firefox and WebKit with a single API.
- [rod](https://github.com/go-rod/rod) ![](https://img.shields.io/github/stars/go-rod/rod.svg) - A Devtools driver to make web automation and scraping easy.

### Fail injection

- [failpoint](https://github.com/pingcap/failpoint) ![](https://img.shields.io/github/stars/pingcap/failpoint.svg) - An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [address](https://github.com/bojanz/address) ![](https://img.shields.io/github/stars/bojanz/address.svg) - Handles address representation, validation and formatting.
- [align](https://github.com/Guitarbum722/align) ![](https://img.shields.io/github/stars/Guitarbum722/align.svg) - A general purpose application that aligns text.
- [bytes](https://github.com/labstack/gommon/tree/master/bytes) ![](https://img.shields.io/github/stars/labstack/gommon.svg) - Formats and parses numeric byte values (10K, 2M, 3G, etc.).
- [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) ![](https://img.shields.io/github/stars/ianlopshire/go-fixedwidth.svg) - Fixed-width text formatting (encoder/decoder with reflection).
- [go-humanize](https://github.com/dustin/go-humanize) ![](https://img.shields.io/github/stars/dustin/go-humanize.svg) - Formatters for time, numbers, and memory size to human readable format.
- [gotabulate](https://github.com/bndr/gotabulate) ![](https://img.shields.io/github/stars/bndr/gotabulate.svg) - Easily pretty-print your tabular data with Go.
- [sq](https://github.com/neilotoole/sq) ![](https://img.shields.io/github/stars/neilotoole/sq.svg) - Convert data from SQL databases or document formats like CSV or Excel into formats such as JSON, Excel, CSV, HTML, Markdown, XML, and YAML.
- [textwrap](https://github.com/isbm/textwrap) ![](https://img.shields.io/github/stars/isbm/textwrap.svg) - Wraps text at end of lines. Implementation of `textwrap` module from Python.

### Markup Languages

- [bafi](https://github.com/mmalcek/bafi) ![](https://img.shields.io/github/stars/mmalcek/bafi.svg) - Universal JSON, BSON, YAML, XML translator to ANY format using templates.
- [bbConvert](https://github.com/CalebQ42/bbConvert) ![](https://img.shields.io/github/stars/CalebQ42/bbConvert.svg) - Converts bbCode to HTML that allows you to add support for custom bbCode tags.
- [blackfriday](https://github.com/russross/blackfriday) ![](https://img.shields.io/github/stars/russross/blackfriday.svg) - Markdown processor in Go.
- [go-output-format](https://github.com/drewstinnett/go-output-format) ![](https://img.shields.io/github/stars/drewstinnett/go-output-format.svg) - Output go structures into multiple formats (YAML/JSON/etc) in your command line app.
- [go-toml](https://github.com/pelletier/go-toml) ![](https://img.shields.io/github/stars/pelletier/go-toml.svg) - Go library for the TOML format with query support and handy cli tools.
- [goldmark](https://github.com/yuin/goldmark) ![](https://img.shields.io/github/stars/yuin/goldmark.svg) - A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured.
- [goq](https://github.com/andrewstuart/goq) ![](https://img.shields.io/github/stars/andrewstuart/goq.svg) - Declarative unmarshalling of HTML using struct tags with jQuery syntax (uses GoQuery).
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) ![](https://img.shields.io/github/stars/JohannesKaufmann/html-to-markdown.svg) - Convert HTML to Markdown. Even works with entire websites and can be extended through rules.
- [htmlquery](https://github.com/antchfx/htmlquery) ![](https://img.shields.io/github/stars/antchfx/htmlquery.svg) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.
- [htmlyaml](https://github.com/nikolaydubina/htmlyaml) ![](https://img.shields.io/github/stars/nikolaydubina/htmlyaml.svg) - Rich rendering of YAML as HTML in Go.
- [htree](https://github.com/bobg/htree) ![](https://img.shields.io/github/stars/bobg/htree.svg) - Traverse, navigate, filter, and otherwise process trees of [html.Node](https://pkg.go.dev/golang.org/x/net/html#Node) objects.
- [mxj](https://github.com/clbanning/mxj) ![](https://img.shields.io/github/stars/clbanning/mxj.svg) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.
- [toml](https://github.com/BurntSushi/toml) ![](https://img.shields.io/github/stars/BurntSushi/toml.svg) - TOML configuration format (encoder/decoder with reflection).

### Parsers/Encoders/Decoders

- [allot](https://github.com/sbstjn/allot) ![](https://img.shields.io/github/stars/sbstjn/allot.svg) - Placeholder and wildcard text parsing for CLI tools and bots.
- [codetree](https://github.com/aerogo/codetree) ![](https://img.shields.io/github/stars/aerogo/codetree.svg) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.
- [commonregex](https://github.com/mingrammer/commonregex) ![](https://img.shields.io/github/stars/mingrammer/commonregex.svg) - A collection of common regular expressions for Go.
- [did](https://github.com/ockam-network/did) ![](https://img.shields.io/github/stars/ockam-network/did.svg) - DID (Decentralized Identifiers) Parser and Stringer in Go.
- [doi](https://github.com/hscells/doi) ![](https://img.shields.io/github/stars/hscells/doi.svg) - Document object identifier (doi) parser in Go.
- [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) ![](https://img.shields.io/github/stars/editorconfig/editorconfig-core-go.svg) - Editorconfig file parser and manipulator for Go.
- [encdec](https://github.com/mickep76/encdec) ![](https://img.shields.io/github/stars/mickep76/encdec.svg) - Package provides a generic interface to encoders and decoders.
- [go-fasttld](https://github.com/elliotwutingfeng/go-fasttld) ![](https://img.shields.io/github/stars/elliotwutingfeng/go-fasttld.svg) - High performance effective top level domains (eTLD) extraction module.
- [go-nmea](https://github.com/adrianmo/go-nmea) ![](https://img.shields.io/github/stars/adrianmo/go-nmea.svg) - NMEA parser library for the Go language.
- [go-querystring](https://github.com/google/go-querystring) ![](https://img.shields.io/github/stars/google/go-querystring.svg) - Go library for encoding structs into URL query parameters.
- [go-vcard](https://github.com/emersion/go-vcard) ![](https://img.shields.io/github/stars/emersion/go-vcard.svg) - Parse and format vCard.
- [godump](https://github.com/yassinebenaid/godump) ![](https://img.shields.io/github/stars/yassinebenaid/godump.svg) - Pretty print any GO variable with ease, an alternative to Go's `fmt.Printf("%#v")`.
- [godump (goforj)](https://github.com/goforj/godump) ![](https://img.shields.io/github/stars/goforj/godump.svg) - Pretty-print Go structs with Laravel/Symfony-style dumps, full type info, colorized CLI output, cycle detection, and private field access.
- [gofeed](https://github.com/mmcdole/gofeed) ![](https://img.shields.io/github/stars/mmcdole/gofeed.svg) - Parse RSS and Atom feeds in Go.
- [gographviz](https://github.com/awalterschulze/gographviz) ![](https://img.shields.io/github/stars/awalterschulze/gographviz.svg) - Parses the Graphviz DOT language.
- [gonameparts](https://github.com/polera/gonameparts) ![](https://img.shields.io/github/stars/polera/gonameparts.svg) - Parses human names into individual name parts.
- [ltsv](https://github.com/Wing924/ltsv) ![](https://img.shields.io/github/stars/Wing924/ltsv.svg) - High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go.
- [normalize](https://github.com/avito-tech/normalize) ![](https://img.shields.io/github/stars/avito-tech/normalize.svg) - Sanitize, normalize and compare fuzzy text.
- [parseargs-go](https://github.com/nproc/parseargs-go) ![](https://img.shields.io/github/stars/nproc/parseargs-go.svg) - string argument parser that understands quotes and backslashes.
- [prattle](https://github.com/askeladdk/prattle) ![](https://img.shields.io/github/stars/askeladdk/prattle.svg) - Scan and parse LL(1) grammars simply and efficiently.
- [sh](https://github.com/mvdan/sh) ![](https://img.shields.io/github/stars/mvdan/sh.svg) - Shell parser and formatter.
- [tokenizer](https://github.com/bzick/tokenizer) ![](https://img.shields.io/github/stars/bzick/tokenizer.svg) - Parse any string, slice or infinite buffer to any tokens.
- [vdf](https://github.com/andygrunwald/vdf) ![](https://img.shields.io/github/stars/andygrunwald/vdf.svg) - A Lexer and Parser for Valves Data Format (known as vdf) written in Go.
- [when](https://github.com/olebedev/when) ![](https://img.shields.io/github/stars/olebedev/when.svg) - Natural EN and RU language date/time parser with pluggable rules.
- [xj2go](https://github.com/stackerzzq/xj2go) ![](https://img.shields.io/github/stars/stackerzzq/xj2go.svg) - Convert xml or json to go struct.

### Regular Expressions

- [genex](https://github.com/alixaxel/genex) ![](https://img.shields.io/github/stars/alixaxel/genex.svg) - Count and expand Regular Expressions into all matching Strings.
- [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) ![](https://img.shields.io/github/stars/IGLOU-EU/go-wildcard.svg) - Simple and lightweight wildcard pattern matching.
- [goregen](https://github.com/zach-klippenstein/goregen) ![](https://img.shields.io/github/stars/zach-klippenstein/goregen.svg) - Library for generating random strings from regular expressions.
- [regroup](https://github.com/oriser/regroup) ![](https://img.shields.io/github/stars/oriser/regroup.svg) - Match regex expression named groups into go struct using struct tags and automatic parsing.
- [rex](https://github.com/hedhyw/rex) ![](https://img.shields.io/github/stars/hedhyw/rex.svg) - Regular expressions builder.

### Sanitation

- [bluemonday](https://github.com/microcosm-cc/bluemonday) ![](https://img.shields.io/github/stars/microcosm-cc/bluemonday.svg) - HTML Sanitizer.
- [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) ![](https://img.shields.io/github/stars/JoshuaDoes/gofuckyourself.svg) - A sanitization-based swear filter for Go.

### Scrapers

- [colly](https://github.com/asciimoo/colly) ![](https://img.shields.io/github/stars/asciimoo/colly.svg) - Fast and Elegant Scraping Framework for Gophers.
- [dataflowkit](https://github.com/slotix/dataflowkit) ![](https://img.shields.io/github/stars/slotix/dataflowkit.svg) - Web scraping Framework to turn websites into structured data.
- [go-recipe](https://github.com/kkyr/go-recipe) ![](https://img.shields.io/github/stars/kkyr/go-recipe.svg) - A package for scraping recipes from websites.
- [go-sitemap-parser](https://github.com/aafeher/go-sitemap-parser) ![](https://img.shields.io/github/stars/aafeher/go-sitemap-parser.svg) - Go language library for parsing Sitemaps.
- [GoQuery](https://github.com/PuerkitoBio/goquery) ![](https://img.shields.io/github/stars/PuerkitoBio/goquery.svg) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.
- [pagser](https://github.com/foolin/pagser) ![](https://img.shields.io/github/stars/foolin/pagser.svg) - Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler.
- [Tagify](https://github.com/zoomio/tagify) ![](https://img.shields.io/github/stars/zoomio/tagify.svg) - Produces a set of tags from given source.
- [walker](https://github.com/cyucelen/walker) ![](https://img.shields.io/github/stars/cyucelen/walker.svg) - Seamlessly fetch paginated data from any source. Simple and high performance API scraping included.
- [xurls](https://github.com/mvdan/xurls) ![](https://img.shields.io/github/stars/mvdan/xurls.svg) - Extract urls from text.

### RSS

- [podcast](https://github.com/eduncan911/podcast) ![](https://img.shields.io/github/stars/eduncan911/podcast.svg) - iTunes Compliant and RSS 2.0 Podcast Generator in Golang

### Utility/Miscellaneous

- [go-runewidth](https://github.com/mattn/go-runewidth) ![](https://img.shields.io/github/stars/mattn/go-runewidth.svg) - Functions to get fixed width of the character or string.
- [kace](https://github.com/codemodus/kace) ![](https://img.shields.io/github/stars/codemodus/kace.svg) - Common case conversions covering common initialisms.
- [lancet](https://github.com/duke-git/lancet) ![](https://img.shields.io/github/stars/duke-git/lancet.svg) - A comprehensive, Lodash-like utility library for Go
- [petrovich](https://github.com/striker2000/petrovich) ![](https://img.shields.io/github/stars/striker2000/petrovich.svg) - Petrovich is the library which inflects Russian names to given grammatical case.
- [radix](https://github.com/yourbasic/radix) ![](https://img.shields.io/github/stars/yourbasic/radix.svg) - Fast string sorting algorithm.
- [TySug](https://github.com/Dynom/TySug) ![](https://img.shields.io/github/stars/Dynom/TySug.svg) - Alternative suggestions with respect to keyboard layouts.
- [w2vgrep](https://github.com/arunsupe/semantic-grep) ![](https://img.shields.io/github/stars/arunsupe/semantic-grep.svg) - A semantic grep tool using word embeddings to find semantically similar matches. For example, searching for "death" will find "dead", "killing", "murder".

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [airtable](https://github.com/mehanizm/airtable) ![](https://img.shields.io/github/stars/mehanizm/airtable.svg) - Go client library for the [Airtable API](https://airtable.com/api).
- [anaconda](https://github.com/ChimeraCoder/anaconda) ![](https://img.shields.io/github/stars/ChimeraCoder/anaconda.svg) - Go client library for the Twitter 1.1 API.
- [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) ![](https://img.shields.io/github/stars/Kachit/appstore-sdk-go.svg) - Unofficial Golang SDK for AppStore Connect API.
- [aws-encryption-sdk-go](https://github.com/chainifynet/aws-encryption-sdk-go) ![](https://img.shields.io/github/stars/chainifynet/aws-encryption-sdk-go.svg) - Unofficial Go SDK implementation of the [AWS Encryption SDK](https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/index.html).
- [aws-sdk-go](https://github.com/aws/aws-sdk-go-v2) ![](https://img.shields.io/github/stars/aws/aws-sdk-go-v2.svg) - The official AWS SDK for the Go programming language.
- [bqwriter](https://github.com/OTA-Insight/bqwriter) ![](https://img.shields.io/github/stars/OTA-Insight/bqwriter.svg) - High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout.
- [brewerydb](https://github.com/naegelejd/brewerydb) ![](https://img.shields.io/github/stars/naegelejd/brewerydb.svg) - Go library for accessing the BreweryDB API.
- [cachet](https://github.com/andygrunwald/cachet) ![](https://img.shields.io/github/stars/andygrunwald/cachet.svg) - Go client library for [Cachet (open source status page system)](https://cachethq.io/).
- [circleci](https://github.com/jszwedko/go-circleci) ![](https://img.shields.io/github/stars/jszwedko/go-circleci.svg) - Go client library for interacting with CircleCI's API.
- [clarifai](https://github.com/samuelcouch/clarifai) ![](https://img.shields.io/github/stars/samuelcouch/clarifai.svg) - Go client library for interfacing with the Clarifai API.
- [codeship-go](https://github.com/codeship/codeship-go) ![](https://img.shields.io/github/stars/codeship/codeship-go.svg) - Go client library for interacting with Codeship's API v2.
- [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) ![](https://img.shields.io/github/stars/coinpaprika/coinpaprika-api-go-client.svg) - Go client library for interacting with Coinpaprika's API.
- [device-check-go](https://github.com/rinchsan/device-check-go) ![](https://img.shields.io/github/stars/rinchsan/device-check-go.svg) - Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1.
- [discordgo](https://github.com/bwmarrin/discordgo) ![](https://img.shields.io/github/stars/bwmarrin/discordgo.svg) - Go bindings for the Discord Chat API.
- [disgo](https://github.com/switchupcb/disgo) ![](https://img.shields.io/github/stars/switchupcb/disgo.svg) - Go API Wrapper for the Discord API.
- [dusupay-sdk-go](https://github.com/Kachit/dusupay-sdk-go) ![](https://img.shields.io/github/stars/Kachit/dusupay-sdk-go.svg) - Unofficial Dusupay payment gateway API Client for Go
- [ethrpc](https://github.com/onrik/ethrpc) ![](https://img.shields.io/github/stars/onrik/ethrpc.svg) - Go bindings for Ethereum JSON RPC API.
- [facebook](https://github.com/huandu/facebook) ![](https://img.shields.io/github/stars/huandu/facebook.svg) - Go Library that supports the Facebook Graph API.
- [fasapay-sdk-go](https://github.com/Kachit/fasapay-sdk-go) ![](https://img.shields.io/github/stars/Kachit/fasapay-sdk-go.svg) - Unofficial Fasapay payment gateway XML API Client for Golang.
- [fcm](https://github.com/maddevsio/fcm) ![](https://img.shields.io/github/stars/maddevsio/fcm.svg) - Go library for Firebase Cloud Messaging.
- [gads](https://github.com/emiddleton/gads) ![](https://img.shields.io/github/stars/emiddleton/gads.svg) - Google Adwords Unofficial API.
- [gcm](https://github.com/Aorioli/gcm) ![](https://img.shields.io/github/stars/Aorioli/gcm.svg) - Go library for Google Cloud Messaging.
- [geo-golang](https://github.com/codingsince1985/geo-golang) ![](https://img.shields.io/github/stars/codingsince1985/geo-golang.svg) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://developer.mapquest.com/documentation/api/geocoding/), [Nominatim](https://nominatim.org/release-docs/latest/api/Overview/), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.
- [github](https://github.com/google/go-github) ![](https://img.shields.io/github/stars/google/go-github.svg) - Go library for accessing the GitHub REST API v3.
- [githubql](https://github.com/shurcooL/githubql) ![](https://img.shields.io/github/stars/shurcooL/githubql.svg) - Go library for accessing the GitHub GraphQL API v4.
- [go-atlassian](https://github.com/ctreminiom/go-atlassian) ![](https://img.shields.io/github/stars/ctreminiom/go-atlassian.svg) - Go library for accessing the [Atlassian Cloud](https://www.atlassian.com/enterprise/cloud) services (Jira, Jira Service Management, Jira Agile, Confluence, Admin Cloud)
- [go-aws-news](https://github.com/circa10a/go-aws-news) ![](https://img.shields.io/github/stars/circa10a/go-aws-news.svg) - Go application and library to fetch what's new from AWS.
- [go-chronos](https://github.com/axelspringer/go-chronos) ![](https://img.shields.io/github/stars/axelspringer/go-chronos.svg) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler
- [go-gerrit](https://github.com/andygrunwald/go-gerrit) ![](https://img.shields.io/github/stars/andygrunwald/go-gerrit.svg) - Go client library for [Gerrit Code Review](https://www.gerritcodereview.com/).
- [go-hacknews](https://github.com/PaulRosset/go-hacknews) ![](https://img.shields.io/github/stars/PaulRosset/go-hacknews.svg) - Tiny Go client for HackerNews API.
- [go-here](https://github.com/abdullahselek/go-here) ![](https://img.shields.io/github/stars/abdullahselek/go-here.svg) - Go client library around the HERE location based APIs.
- [go-hibp](https://github.com/wneessen/go-hibp) ![](https://img.shields.io/github/stars/wneessen/go-hibp.svg) - Simple Go binding to the "Have I Been Pwned" APIs.
- [go-imgur](https://github.com/koffeinsource/go-imgur) ![](https://img.shields.io/github/stars/koffeinsource/go-imgur.svg) - Go client library for [imgur](https://imgur.com)
- [go-jira](https://github.com/andygrunwald/go-jira) ![](https://img.shields.io/github/stars/andygrunwald/go-jira.svg) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)
- [go-lark](https://github.com/go-lark/lark) ![](https://img.shields.io/github/stars/go-lark/lark.svg) - An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform.
- [go-marathon](https://github.com/gambol99/go-marathon) ![](https://img.shields.io/github/stars/gambol99/go-marathon.svg) - Go library for interacting with Mesosphere's Marathon PAAS.
- [go-myanimelist](https://github.com/nstratos/go-myanimelist) ![](https://img.shields.io/github/stars/nstratos/go-myanimelist.svg) - Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2).
- [go-openai](https://github.com/sashabaranov/go-openai) ![](https://img.shields.io/github/stars/sashabaranov/go-openai.svg) - OpenAI ChatGPT, DALL·E, Whisper API library for Go.
- [go-openproject](https://github.com/manuelbcd/go-openproject) ![](https://img.shields.io/github/stars/manuelbcd/go-openproject.svg) - Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API.
- [go-postman-collection](https://github.com/rbretecher/go-postman-collection) ![](https://img.shields.io/github/stars/rbretecher/go-postman-collection.svg) - Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).
- [go-redoc](https://github.com/mvrilo/go-redoc) ![](https://img.shields.io/github/stars/mvrilo/go-redoc.svg) - Embedded OpenAPI/Swagger documentation ui for Go using [ReDoc](https://redocly.com/).
- [go-restcountries](https://github.com/chriscross0/go-restcountries) ![](https://img.shields.io/github/stars/chriscross0/go-restcountries.svg) - Go library for the [REST Countries API](https://countrylayer.com/).
- [go-salesforce](https://github.com/k-capehart/go-salesforce) ![](https://img.shields.io/github/stars/k-capehart/go-salesforce.svg) - Go client library for interacting with the [Salesforce REST API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_list.htm).
- [go-sophos](https://github.com/esurdam/go-sophos) ![](https://img.shields.io/github/stars/esurdam/go-sophos.svg) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.
- [go-swagger-ui](https://github.com/esurdam/go-swagger-ui) ![](https://img.shields.io/github/stars/esurdam/go-swagger-ui.svg) - Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json.
- [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.
- [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
- [go-unsplash](https://github.com/hbagdi/go-unsplash) ![](https://img.shields.io/github/stars/hbagdi/go-unsplash.svg) - Go client library for the [Unsplash.com](https://unsplash.com) API.
- [go-xkcd](https://github.com/nishanths/go-xkcd) ![](https://img.shields.io/github/stars/nishanths/go-xkcd.svg) - Go client for the xkcd API.
- [go-yapla](https://gitlab.com/adrienK/go-yapla) - Go client library for the Yapla v2.0 API.
- [goagi](https://github.com/staskobzar/goagi) ![](https://img.shields.io/github/stars/staskobzar/goagi.svg) - Go library to build Asterisk PBX agi/fastagi applications.
- [goami2](https://github.com/staskobzar/goami2) ![](https://img.shields.io/github/stars/staskobzar/goami2.svg) - AMI v2 library for Asterisk PBX.
- [GoFreeDB](https://github.com/FreeLeh/GoFreeDB) ![](https://img.shields.io/github/stars/FreeLeh/GoFreeDB.svg) - Golang library providing common and simple database abstractions on top of Google Sheets.
- [gogtrends](https://github.com/groovili/gogtrends) ![](https://img.shields.io/github/stars/groovili/gogtrends.svg) - Google Trends Unofficial API.
- [golang-tmdb](https://github.com/cyruzin/golang-tmdb) ![](https://img.shields.io/github/stars/cyruzin/golang-tmdb.svg) - Golang wrapper for The Movie Database API v3.
- [golyrics](https://github.com/mamal72/golyrics) ![](https://img.shields.io/github/stars/mamal72/golyrics.svg) - Golyrics is a Go library to fetch music lyrics data from the Wikia website.
- [gomalshare](https://github.com/MonaxGT/gomalshare) ![](https://img.shields.io/github/stars/MonaxGT/gomalshare.svg) - Go library MalShare API [malshare.com](https://www.malshare.com/)
- [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) ![](https://img.shields.io/github/stars/michiwend/gomusicbrainz.svg) - Go MusicBrainz WS2 client library.
- [google](https://github.com/google/google-api-go-client) ![](https://img.shields.io/github/stars/google/google-api-go-client.svg) - Auto-generated Google APIs for Go.
- [google-analytics](https://github.com/chonthu/go-google-analytics) ![](https://img.shields.io/github/stars/chonthu/go-google-analytics.svg) - Simple wrapper for easy google analytics reporting.
- [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) ![](https://img.shields.io/github/stars/GoogleCloudPlatform/gcloud-golang.svg) - Google Cloud APIs Go Client Library.
- [gopaapi5](https://github.com/utekaravinash/gopaapi5) ![](https://img.shields.io/github/stars/utekaravinash/gopaapi5.svg) - Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).
- [gopensky](https://github.com/navidys/gopensky) ![](https://img.shields.io/github/stars/navidys/gopensky.svg) - Go client implementation for [OpenSKY Network](https://opensky-network.org/) live's API (airspace ADS-B and Mode S data).
- [gosip](https://github.com/koltyakov/gosip) ![](https://img.shields.io/github/stars/koltyakov/gosip.svg) - Client library for SharePoint.
- [gostorm](https://github.com/jsgilmore/gostorm) ![](https://img.shields.io/github/stars/jsgilmore/gostorm.svg) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.
- [hipchat](https://github.com/andybons/hipchat) ![](https://img.shields.io/github/stars/andybons/hipchat.svg) - This project implements a golang client library for the Hipchat API.
- [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) ![](https://img.shields.io/github/stars/daneharrigan/hipchat.svg) - A golang package to communicate with HipChat over XMPP.
- [igdb](https://github.com/Henry-Sarabia/igdb) ![](https://img.shields.io/github/stars/Henry-Sarabia/igdb.svg) - Go client for the [Internet Game Database API](https://api.igdb.com/).
- [ip2location-io-go](https://github.com/ip2location/ip2location-io-go) ![](https://img.shields.io/github/stars/ip2location/ip2location-io-go.svg) - Go wrapper for the IP2Location.io API [IP2Location.io](https://www.ip2location.io/).
- [jokeapi-go](https://github.com/icelain/jokeapi) ![](https://img.shields.io/github/stars/icelain/jokeapi.svg) - Go client for [JokeAPI](https://sv443.net/jokeapi/v2/).
- [lark](https://github.com/chyroc/lark) ![](https://img.shields.io/github/stars/chyroc/lark.svg) - [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback.
- [lastpass-go](https://github.com/ansd/lastpass-go) ![](https://img.shields.io/github/stars/ansd/lastpass-go.svg) - Go client library for the [LastPass](https://www.lastpass.com/) API.
- [libgoffi](https://github.com/clevabit/libgoffi) ![](https://img.shields.io/github/stars/clevabit/libgoffi.svg) - Library adapter toolbox for native [libffi](https://sourceware.org/libffi/) integration
- [Medium](https://github.com/Medium/medium-sdk-go) ![](https://img.shields.io/github/stars/Medium/medium-sdk-go.svg) - Golang SDK for Medium's OAuth2 API.
- [megos](https://github.com/andygrunwald/megos) ![](https://img.shields.io/github/stars/andygrunwald/megos.svg) - Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster.
- [minio-go](https://github.com/minio/minio-go) ![](https://img.shields.io/github/stars/minio/minio-go.svg) - Minio Go Library for Amazon S3 compatible cloud storage.
- [mixpanel](https://github.com/dukex/mixpanel) ![](https://img.shields.io/github/stars/dukex/mixpanel.svg) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
- [newsapi-go](https://github.com/jellydator/newsapi-go) ![](https://img.shields.io/github/stars/jellydator/newsapi-go.svg) - Go client for [NewsAPI](https://newsapi.org/).
- [openaigo](https://github.com/otiai10/openaigo) ![](https://img.shields.io/github/stars/otiai10/openaigo.svg) - OpenAI GPT3/GPT3.5 ChatGPT API client library for Go.
- [patreon-go](https://github.com/mxpv/patreon-go) ![](https://img.shields.io/github/stars/mxpv/patreon-go.svg) - Go library for Patreon API.
- [paypal](https://github.com/logpacker/PayPal-Go-SDK) ![](https://img.shields.io/github/stars/logpacker/PayPal-Go-SDK.svg) - Wrapper for PayPal payment API.
- [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) ![](https://img.shields.io/github/stars/playlyfe/playlyfe-go-sdk.svg) - The Playlyfe Rest API Go SDK.
- [pushover](https://github.com/gregdel/pushover) ![](https://img.shields.io/github/stars/gregdel/pushover.svg) - Go wrapper for the Pushover API.
- [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) ![](https://img.shields.io/github/stars/dimuska139/rawg-sdk-go.svg) - Go library for the [RAWG Video Games Database](https://rawg.io/) API
- [shopify](https://github.com/rapito/go-shopify) ![](https://img.shields.io/github/stars/rapito/go-shopify.svg) - Go Library to make CRUD request to the Shopify API.
- [simples3](https://github.com/rhnvrm/simples3) ![](https://img.shields.io/github/stars/rhnvrm/simples3.svg) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go.
- [slack](https://github.com/slack-go/slack) ![](https://img.shields.io/github/stars/slack-go/slack.svg) - Slack API in Go.
- [smite](https://github.com/sergiotapia/smitego) ![](https://img.shields.io/github/stars/sergiotapia/smitego.svg) - Go package to wraps access to the Smite game API.
- [spotify](https://github.com/rapito/go-spotify) ![](https://img.shields.io/github/stars/rapito/go-spotify.svg) - Go Library to access Spotify WEB API.
- [steam](https://github.com/sostronk/go-steam) ![](https://img.shields.io/github/stars/sostronk/go-steam.svg) - Go Library to interact with Steam game servers.
- [stripe](https://github.com/stripe/stripe-go) ![](https://img.shields.io/github/stars/stripe/stripe-go.svg) - Go client for the Stripe API.
- [swag](https://github.com/zc2638/swag) ![](https://img.shields.io/github/stars/zc2638/swag.svg) - No comments, simple go wrapper to create swagger 2.0 compatible APIs. Support most routing frameworks, such as built-in, gin, chi, mux, echo, httprouter, fasthttp and more.
- [textbelt](https://github.com/dietsche/textbelt) ![](https://img.shields.io/github/stars/dietsche/textbelt.svg) - Go client for the textbelt.com txt messaging API.
- [Trello](https://github.com/adlio/trello) ![](https://img.shields.io/github/stars/adlio/trello.svg) - Go wrapper for the Trello API.
- [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) ![](https://img.shields.io/github/stars/mrbenosborne/tripadvisor-golang.svg) - Go wrapper for the TripAdvisor API.
- [tumblr](https://github.com/mattcunningham/gumblr) ![](https://img.shields.io/github/stars/mattcunningham/gumblr.svg) - Go wrapper for the Tumblr v2 API.
- [uptimerobot](https://github.com/bitfield/uptimerobot) ![](https://img.shields.io/github/stars/bitfield/uptimerobot.svg) - Go wrapper and command-line client for the Uptime Robot v2 API.
- [vl-go](https://github.com/verifid/vl-go) ![](https://img.shields.io/github/stars/verifid/vl-go.svg) - Go client library around the VerifID identity verification layer API.
- [webhooks](https://github.com/go-playground/webhooks) ![](https://img.shields.io/github/stars/go-playground/webhooks.svg) - Webhook receiver for GitHub and Bitbucket.
- [wit-go](https://github.com/wit-ai/wit-go) ![](https://img.shields.io/github/stars/wit-ai/wit-go.svg) - Go client for wit.ai HTTP API.
- [ynab](https://github.com/brunomvsouza/ynab.go) ![](https://img.shields.io/github/stars/brunomvsouza/ynab.go.svg) - Go wrapper for the YNAB API.
- [zooz](https://github.com/gojuno/go-zooz) ![](https://img.shields.io/github/stars/gojuno/go-zooz.svg) - Go client for the Zooz API.

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [abstract](https://github.com/maxbolgarin/abstract) ![](https://img.shields.io/github/stars/maxbolgarin/abstract.svg) - Abstractions and utilities to get rid of boilerplate code in business logic.
- [apm](https://github.com/topfreegames/apm) ![](https://img.shields.io/github/stars/topfreegames/apm.svg) - Process manager for Golang applications with an HTTP API.
- [backscanner](https://github.com/icza/backscanner) ![](https://img.shields.io/github/stars/icza/backscanner.svg) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.
- [bed](https://github.com/itchyny/bed) ![](https://img.shields.io/github/stars/itchyny/bed.svg) - A Vim-like binary editor written in Go.
- [blank](https://github.com/Henry-Sarabia/blank) ![](https://img.shields.io/github/stars/Henry-Sarabia/blank.svg) - Verify or remove blanks and whitespace from strings.
- [bleep](https://github.com/sinhashubham95/bleep) ![](https://img.shields.io/github/stars/sinhashubham95/bleep.svg) - Perform any number of actions on any set of OS signals in Go.
- [boilr](https://github.com/tmrts/boilr) ![](https://img.shields.io/github/stars/tmrts/boilr.svg) - Blazingly fast CLI tool for creating projects from boilerplate templates.
- [boring](https://github.com/alebeck/boring) ![](https://img.shields.io/github/stars/alebeck/boring.svg) - Simple command-line SSH tunnel manager.
- [changie](https://github.com/miniscruff/changie) ![](https://img.shields.io/github/stars/miniscruff/changie.svg) - Automated changelog tool for preparing releases with lots of customization options.
- [chyle](https://github.com/antham/chyle) ![](https://img.shields.io/github/stars/antham/chyle.svg) - Changelog generator using a git repository with multiple configuration possibilities.
- [circuit](https://github.com/cep21/circuit) ![](https://img.shields.io/github/stars/cep21/circuit.svg) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.
- [circuitbreaker](https://github.com/rubyist/circuitbreaker) ![](https://img.shields.io/github/stars/rubyist/circuitbreaker.svg) - Circuit Breakers in Go.
- [clipboard](https://github.com/golang-design/clipboard) ![](https://img.shields.io/github/stars/golang-design/clipboard.svg) - 📋 cross-platform clipboard package in Go.
- [clockwork](https://github.com/jonboulle/clockwork) ![](https://img.shields.io/github/stars/jonboulle/clockwork.svg) - A simple fake clock for golang.
- [cmd](https://github.com/SimonBaeumer/cmd) ![](https://img.shields.io/github/stars/SimonBaeumer/cmd.svg) - Library for executing shell commands on osx, windows and linux.
- [config-file-validator](https://github.com/Boeing/config-file-validator) ![](https://img.shields.io/github/stars/Boeing/config-file-validator.svg) - Cross Platform tool to validate configuration files.
- [contem](https://github.com/maxbolgarin/contem) ![](https://img.shields.io/github/stars/maxbolgarin/contem.svg) - Drop-in context.Context replacement for graceful shutdown Go applications.
- [cookie](https://github.com/syntaqx/cookie) ![](https://img.shields.io/github/stars/syntaqx/cookie.svg) - Cookie struct parsing and helper package.
- [copy-pasta](https://github.com/jutkko/copy-pasta) ![](https://img.shields.io/github/stars/jutkko/copy-pasta.svg) - Universal multi-workstation clipboard that uses S3 like backend for the storage.
- [countries](https://github.com/biter777/countries) ![](https://img.shields.io/github/stars/biter777/countries.svg) - Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standards.
- [countries](https://github.com/pioz/countries) ![](https://img.shields.io/github/stars/pioz/countries.svg) - All you need when you are working with countries in Go.
- [create-go-app](https://github.com/create-go-app/cli) ![](https://img.shields.io/github/stars/create-go-app/cli.svg) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command.
- [cryptgo](https://github.com/Gituser143/cryptgo) ![](https://img.shields.io/github/stars/Gituser143/cryptgo.svg) - Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time!
- [ctop](https://github.com/bcicen/ctop) ![](https://img.shields.io/github/stars/bcicen/ctop.svg) - [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics.
- [ctxutil](https://github.com/posener/ctxutil) ![](https://img.shields.io/github/stars/posener/ctxutil.svg) - A collection of utility functions for contexts.
- [cvt](https://github.com/shockerli/cvt) ![](https://img.shields.io/github/stars/shockerli/cvt.svg) - Easy and safe convert any value to another type.
- [dbt](https://github.com/nikogura/dbt) ![](https://img.shields.io/github/stars/nikogura/dbt.svg) - A framework for running self-updating signed binaries from a central, trusted repository.
- [Death](https://github.com/vrecan/death) ![](https://img.shields.io/github/stars/vrecan/death.svg) - Managing go application shutdown with signals.
- [debounce](https://github.com/floatdrop/debounce) ![](https://img.shields.io/github/stars/floatdrop/debounce.svg) - A zero-allocation debouncer written in Go.
- [delve](https://github.com/derekparker/delve) ![](https://img.shields.io/github/stars/derekparker/delve.svg) - Go debugger.
- [dive](https://github.com/wagoodman/dive) ![](https://img.shields.io/github/stars/wagoodman/dive.svg) - A tool for exploring each layer in a Docker image.
- [dlog](https://github.com/kirillDanshin/dlog) ![](https://img.shields.io/github/stars/kirillDanshin/dlog.svg) - Compile-time controlled logger to make your release smaller without removing debug calls.
- [EaseProbe](https://github.com/megaease/easeprobe) ![](https://img.shields.io/github/stars/megaease/easeprobe.svg) - A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification.
- [equalizer](https://github.com/reugn/equalizer) ![](https://img.shields.io/github/stars/reugn/equalizer.svg) - Quota manager and rate limiter collection for Go.
- [ergo](https://github.com/cristianoliveira/ergo) ![](https://img.shields.io/github/stars/cristianoliveira/ergo.svg) - The management of multiple local services running over different ports made easy.
- [evaluator](https://github.com/nullne/evaluator) ![](https://img.shields.io/github/stars/nullne/evaluator.svg) - Evaluate an expression dynamically based on s-expression. It's simple and easy to extend.
- [Failsafe-go](https://github.com/failsafe-go/failsafe-go) ![](https://img.shields.io/github/stars/failsafe-go/failsafe-go.svg) - Fault tolerance and resilience patterns for Go.
- [filetype](https://github.com/h2non/filetype) ![](https://img.shields.io/github/stars/h2non/filetype.svg) - Small package to infer the file type checking the magic numbers signature.
- [filler](https://github.com/yaronsumel/filler) ![](https://img.shields.io/github/stars/yaronsumel/filler.svg) - small utility to fill structs using "fill" tag.
- [filter](https://github.com/gookit/filter) ![](https://img.shields.io/github/stars/gookit/filter.svg) - provide filtering, sanitizing, and conversion of Go data.
- [fzf](https://github.com/junegunn/fzf) ![](https://img.shields.io/github/stars/junegunn/fzf.svg) - Command-line fuzzy finder written in Go.
- [generate](https://github.com/go-playground/generate) ![](https://img.shields.io/github/stars/go-playground/generate.svg) - runs go generate recursively on a specified path or environment variable and can filter by regex.
- [ghokin](https://github.com/antham/ghokin) ![](https://img.shields.io/github/stars/antham/ghokin.svg) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).
- [git-time-metric](https://github.com/git-time-metric/gtm) ![](https://img.shields.io/github/stars/git-time-metric/gtm.svg) - Simple, seamless, lightweight time tracking for Git.
- [git-tools](https://github.com/kazhuravlev/git-tools) ![](https://img.shields.io/github/stars/kazhuravlev/git-tools.svg) - Tool to help manage git tags.
- [gitbatch](https://github.com/isacikgoz/gitbatch) ![](https://img.shields.io/github/stars/isacikgoz/gitbatch.svg) - manage your git repositories in one place.
- [gitcs](https://github.com/knbr13/gitcs/) ![](https://img.shields.io/github/stars/knbr13/gitcs.svg) - Git Commits Visualizer, CLI tool to visualize your Git commits on your local machine.
- [go-actuator](https://github.com/sinhashubham95/go-actuator) ![](https://img.shields.io/github/stars/sinhashubham95/go-actuator.svg) - Production ready features for Go based web frameworks.
- [go-astitodo](https://github.com/asticode/go-astitodo) ![](https://img.shields.io/github/stars/asticode/go-astitodo.svg) - Parse TODOs in your GO code.
- [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) ![](https://img.shields.io/github/stars/wendigo/go-bind-plugin.svg) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only).
- [go-bsdiff](https://github.com/gabstv/go-bsdiff) ![](https://img.shields.io/github/stars/gabstv/go-bsdiff.svg) - Pure Go bsdiff and bspatch libraries and CLI tools.
- [go-clip](https://github.com/prashantgupta24/go-clip) ![](https://img.shields.io/github/stars/prashantgupta24/go-clip.svg) - A minimalistic clipboard manager for Mac.
- [go-convert](https://github.com/Eun/go-convert) ![](https://img.shields.io/github/stars/Eun/go-convert.svg) - Package go-convert enables you to convert a value into another type.
- [go-countries](https://github.com/mikekonan/go-countries) ![](https://img.shields.io/github/stars/mikekonan/go-countries.svg) - Lightweight lookup over ISO-3166 codes.
- [go-dry](https://github.com/ungerik/go-dry) ![](https://img.shields.io/github/stars/ungerik/go-dry.svg) - DRY (don't repeat yourself) package for Go.
- [go-events](https://github.com/deatil/go-events) ![](https://img.shields.io/github/stars/deatil/go-events.svg) - A go event and event'subscribe package, like wordpress hook functions.
- [go-funk](https://github.com/thoas/go-funk) ![](https://img.shields.io/github/stars/thoas/go-funk.svg) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).
- [go-health](https://github.com/Talento90/go-health) ![](https://img.shields.io/github/stars/Talento90/go-health.svg) - Health package simplifies the way you add health check to your services.
- [go-httpheader](https://github.com/mozillazg/go-httpheader) ![](https://img.shields.io/github/stars/mozillazg/go-httpheader.svg) - Go library for encoding structs into Header fields.
- [go-lambda-cleanup](https://github.com/karl-cardenas-coding/go-lambda-cleanup) ![](https://img.shields.io/github/stars/karl-cardenas-coding/go-lambda-cleanup.svg) - A CLI for removing unused or previous versions of AWS Lambdas.
- [go-lock](https://github.com/viney-shih/go-lock) ![](https://img.shields.io/github/stars/viney-shih/go-lock.svg) - go-lock is a lock library implementing read-write mutex and read-write trylock without starvation.
- [go-pattern-match](https://github.com/PhakornKiong/go-pattern-match) ![](https://img.shields.io/github/stars/PhakornKiong/go-pattern-match.svg) - A Pattern matching library inspired by ts-pattern.
- [go-pkg](https://github.com/chenquan/go-pkg) ![](https://img.shields.io/github/stars/chenquan/go-pkg.svg) - A go toolkit.
- [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) ![](https://img.shields.io/github/stars/mvmaasakkers/go-problemdetails.svg) - Go package for working with Problem Details.
- [go-qr](https://github.com/piglig/go-qr) ![](https://img.shields.io/github/stars/piglig/go-qr.svg) - A native, high-quality and minimalistic QR code generator.
- [go-rate](https://github.com/beefsack/go-rate) ![](https://img.shields.io/github/stars/beefsack/go-rate.svg) - Timed rate limiter for Go.
- [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) ![](https://img.shields.io/github/stars/ikeikeikeike/go-sitemap-generator.svg) - XML Sitemap generator written in Go.
- [go-trigger](https://github.com/sadlil/go-trigger) ![](https://img.shields.io/github/stars/sadlil/go-trigger.svg) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.
- [go-tripper](https://github.com/rajnandan1/go-tripper) ![](https://img.shields.io/github/stars/rajnandan1/go-tripper.svg) - Tripper is a circuit breaker package for Go that allows you to circuit and control the status of circuits.
- [go-type](https://github.com/mikekonan/go-types) ![](https://img.shields.io/github/stars/mikekonan/go-types.svg) - Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types.
- [goback](https://github.com/carlescere/goback) ![](https://img.shields.io/github/stars/carlescere/goback.svg) - Go simple exponential backoff package.
- [goctx](https://github.com/zerosnake0/goctx) ![](https://img.shields.io/github/stars/zerosnake0/goctx.svg) - Get your context value with high performance.
- [godaemon](https://github.com/VividCortex/godaemon) ![](https://img.shields.io/github/stars/VividCortex/godaemon.svg) - Utility to write daemons.
- [godropbox](https://github.com/dropbox/godropbox) ![](https://img.shields.io/github/stars/dropbox/godropbox.svg) - Common libraries for writing Go services/applications from Dropbox.
- [gofn](https://github.com/tiendc/gofn) ![](https://img.shields.io/github/stars/tiendc/gofn.svg) - High performance utility functions written using Generics for Go 1.18+.
- [golarm](https://github.com/msempere/golarm) ![](https://img.shields.io/github/stars/msempere/golarm.svg) - Fire alarms with system events.
- [golog](https://github.com/mlimaloureiro/golog) ![](https://img.shields.io/github/stars/mlimaloureiro/golog.svg) - Easy and lightweight CLI tool to time track your tasks.
- [gopencils](https://github.com/bndr/gopencils) ![](https://img.shields.io/github/stars/bndr/gopencils.svg) - Small and simple package to easily consume REST APIs.
- [goplaceholder](https://github.com/michiwend/goplaceholder) ![](https://img.shields.io/github/stars/michiwend/goplaceholder.svg) - a small golang lib to generate placeholder images.
- [goreadability](https://github.com/philipjkim/goreadability) ![](https://img.shields.io/github/stars/philipjkim/goreadability.svg) - Webpage summary extractor using Facebook Open Graph and arc90's readability.
- [goreleaser](https://github.com/goreleaser/goreleaser) ![](https://img.shields.io/github/stars/goreleaser/goreleaser.svg) - Deliver Go binaries as fast and easily as possible.
- [goreporter](https://github.com/wgliang/goreporter) ![](https://img.shields.io/github/stars/wgliang/goreporter.svg) - Golang tool that does static analysis, unit testing, code review and generate code quality report.
- [goseaweedfs](https://github.com/linxGnu/goseaweedfs) ![](https://img.shields.io/github/stars/linxGnu/goseaweedfs.svg) - SeaweedFS client library with almost full features.
- [gostrutils](https://github.com/ik5/gostrutils) ![](https://img.shields.io/github/stars/ik5/gostrutils.svg) - Collections of string manipulation and conversion functions.
- [gotenv](https://github.com/subosito/gotenv) ![](https://img.shields.io/github/stars/subosito/gotenv.svg) - Load environment variables from `.env` or any `io.Reader` in Go.
- [goval](https://github.com/maja42/goval) ![](https://img.shields.io/github/stars/maja42/goval.svg) - Evaluate arbitrary expressions in Go.
- [graterm](https://github.com/skovtunenko/graterm) ![](https://img.shields.io/github/stars/skovtunenko/graterm.svg) - Provides primitives to perform ordered (sequential/concurrent) GRAceful TERMination (aka shutdown) in Go application.
- [grofer](https://github.com/pesos/grofer) ![](https://img.shields.io/github/stars/pesos/grofer.svg) - A system and resource monitoring tool written in Golang!
- [gubrak](https://github.com/novalagung/gubrak) ![](https://img.shields.io/github/stars/novalagung/gubrak.svg) - Golang utility library with syntactic sugar. It's like lodash, but for golang.
- [handy](https://github.com/miguelpragier/handy) ![](https://img.shields.io/github/stars/miguelpragier/handy.svg) - Many utilities and helpers like string handlers/formatters and validators.
- [healthcheck](https://github.com/kazhuravlev/healthcheck) ![](https://img.shields.io/github/stars/kazhuravlev/healthcheck.svg) - A simple yet powerful readiness test for Kubernetes.
- [hostctl](https://github.com/guumaster/hostctl) ![](https://img.shields.io/github/stars/guumaster/hostctl.svg) - A CLI tool to manage /etc/hosts with easy commands.
- [htcat](https://github.com/htcat/htcat) ![](https://img.shields.io/github/stars/htcat/htcat.svg) - Parallel and Pipelined HTTP GET Utility.
- [hub](https://github.com/github/hub) ![](https://img.shields.io/github/stars/github/hub.svg) - wrap git commands with additional functionality to interact with github from the terminal.
- [immortal](https://github.com/immortal/immortal) ![](https://img.shields.io/github/stars/immortal/immortal.svg) - \*nix cross-platform (OS agnostic) supervisor.
- [jet](https://github.com/NicoNex/jet) ![](https://img.shields.io/github/stars/NicoNex/jet.svg) - Just Edit Text: a fast and powerful tool for finding and replacing file content and names using regular expressions.
- [jsend](https://github.com/clevergo/jsend) ![](https://img.shields.io/github/stars/clevergo/jsend.svg) - JSend's implementation written in Go.
- [json-log-viewer](https://github.com/hedhyw/json-log-viewer) ![](https://img.shields.io/github/stars/hedhyw/json-log-viewer.svg) - Interactive viewer for JSON logs.
- [jump](https://github.com/gsamokovarov/jump) ![](https://img.shields.io/github/stars/gsamokovarov/jump.svg) - Jump helps you navigate faster by learning your habits.
- [just](https://github.com/kazhuravlev/just) ![](https://img.shields.io/github/stars/kazhuravlev/just.svg) - Just a collection of useful functions for working with generic data structures.
- [koazee](https://github.com/wesovilabs/koazee) ![](https://img.shields.io/github/stars/wesovilabs/koazee.svg) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.
- [lang](https://github.com/maxbolgarin/lang) ![](https://img.shields.io/github/stars/maxbolgarin/lang.svg) - Generic one-liners to work with variables, slices and maps without boilerplate code.
- [lets-go](https://github.com/aplescia-chwy/lets-go) ![](https://img.shields.io/github/stars/aplescia-chwy/lets-go.svg) - Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities.
- [limiters](https://github.com/mennanov/limiters) ![](https://img.shields.io/github/stars/mennanov/limiters.svg) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.
- [lo](https://github.com/samber/lo) ![](https://img.shields.io/github/stars/samber/lo.svg) - A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...)
- [loncha](https://github.com/kazu/loncha) ![](https://img.shields.io/github/stars/kazu/loncha.svg) - A high-performance slice Utilities.
- [lrserver](https://github.com/jaschaephraim/lrserver) ![](https://img.shields.io/github/stars/jaschaephraim/lrserver.svg) - LiveReload server for Go.
- [mani](https://github.com/alajmo/mani) ![](https://img.shields.io/github/stars/alajmo/mani.svg) - CLI tool to help you manage multiple repositories.
- [mc](https://github.com/minio/mc) ![](https://img.shields.io/github/stars/minio/mc.svg) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.
- [mergo](https://github.com/imdario/mergo) ![](https://img.shields.io/github/stars/imdario/mergo.svg) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.
- [mimemagic](https://github.com/zRedShift/mimemagic) ![](https://img.shields.io/github/stars/zRedShift/mimemagic.svg) - Pure Go ultra performant MIME sniffing library/utility.
- [mimetype](https://github.com/gabriel-vasile/mimetype) ![](https://img.shields.io/github/stars/gabriel-vasile/mimetype.svg) - Package for MIME type detection based on magic numbers.
- [minify](https://github.com/tdewolff/minify) ![](https://img.shields.io/github/stars/tdewolff/minify.svg) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.
- [minquery](https://github.com/icza/minquery) ![](https://img.shields.io/github/stars/icza/minquery.svg) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).
- [moldova](https://github.com/StabbyCutyou/moldova) ![](https://img.shields.io/github/stars/StabbyCutyou/moldova.svg) - Utility for generating random data based on an input template.
- [mole](https://github.com/davrodpin/mole) ![](https://img.shields.io/github/stars/davrodpin/mole.svg) - cli app to easily create ssh tunnels.
- [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) ![](https://img.shields.io/github/stars/gobeam/mongo-go-pagination.svg) - Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines.
- [mssqlx](https://github.com/linxGnu/mssqlx) ![](https://img.shields.io/github/stars/linxGnu/mssqlx.svg) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
- [multitick](https://github.com/VividCortex/multitick) ![](https://img.shields.io/github/stars/VividCortex/multitick.svg) - Multiplexor for aligned tickers.
- [netbug](https://github.com/e-dard/netbug) ![](https://img.shields.io/github/stars/e-dard/netbug.svg) - Easy remote profiling of your services.
- [nfdump](https://github.com/chrispassas/nfdump) ![](https://img.shields.io/github/stars/chrispassas/nfdump.svg) - Read nfdump netflow files.
- [nostromo](https://github.com/pokanop/nostromo) ![](https://img.shields.io/github/stars/pokanop/nostromo.svg) - CLI for building powerful aliases.
- [okrun](https://github.com/xta/okrun) ![](https://img.shields.io/github/stars/xta/okrun.svg) - go run error steamroller.
- [olaf](https://github.com/btnguyen2k/olaf) ![](https://img.shields.io/github/stars/btnguyen2k/olaf.svg) - Twitter Snowflake implemented in Go.
- [onecache](https://github.com/adelowo/onecache) ![](https://img.shields.io/github/stars/adelowo/onecache.svg) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
- [optional](https://github.com/kazhuravlev/optional) ![](https://img.shields.io/github/stars/kazhuravlev/optional.svg) - Optional struct fields and vars.
- [panicparse](https://github.com/maruel/panicparse) ![](https://img.shields.io/github/stars/maruel/panicparse.svg) - Groups similar goroutines and colorizes stack dump.
- [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) ![](https://img.shields.io/github/stars/alexpantyukhin/go-pattern-match.svg) - Pattern matching library.
- [peco](https://github.com/peco/peco) ![](https://img.shields.io/github/stars/peco/peco.svg) - Simplistic interactive filtering tool.
- [pgo](https://github.com/arthurkushman/pgo) ![](https://img.shields.io/github/stars/arthurkushman/pgo.svg) - Convenient functions for PHP community.
- [pm](https://github.com/VividCortex/pm) ![](https://img.shields.io/github/stars/VividCortex/pm.svg) - Process (i.e. goroutine) manager with an HTTP API.
- [pointer](https://github.com/xorcare/pointer) ![](https://img.shields.io/github/stars/xorcare/pointer.svg) - Package pointer contains helper routines for simplifying the creation of optional fields of basic type.
- [ptr](https://github.com/gotidy/ptr) ![](https://img.shields.io/github/stars/gotidy/ptr.svg) - Package that provide functions for simplified creation of pointers from constants of basic types.
- [rate](https://github.com/webriots/rate) ![](https://img.shields.io/github/stars/webriots/rate.svg) - High-performance rate limiting library with token bucket and AIMD strategies.
- [rclient](https://github.com/zpatrick/rclient) ![](https://img.shields.io/github/stars/zpatrick/rclient.svg) - Readable, flexible, simple-to-use client for REST APIs.
- [release](https://github.com/tomodian/release) ![](https://img.shields.io/github/stars/tomodian/release.svg) - CLI for Keep-a-changelog formatted changelogs.
- [remote-touchpad](https://github.com/Unrud/remote-touchpad) ![](https://img.shields.io/github/stars/Unrud/remote-touchpad.svg) - Control mouse and keyboard from a smartphone.
- [repeat](https://github.com/ssgreg/repeat) ![](https://img.shields.io/github/stars/ssgreg/repeat.svg) - Go implementation of different backoff strategies useful for retrying operations and heartbeating.
- [request](https://github.com/mozillazg/request) ![](https://img.shields.io/github/stars/mozillazg/request.svg) - Go HTTP Requests for Humans™.
- [rerun](https://github.com/ivpusic/rerun) ![](https://img.shields.io/github/stars/ivpusic/rerun.svg) - Recompiling and rerunning go apps when source changes.
- [rest-go](https://github.com/edermanoel94/rest-go) ![](https://img.shields.io/github/stars/edermanoel94/rest-go.svg) - A package that provide many helpful methods for working with rest api.
- [retro](https://github.com/goioc/retro) ![](https://img.shields.io/github/stars/goioc/retro.svg) - Handy retry-on-error library with extensive flexibility (backoff strategies, caps, etc).
- [retry](https://github.com/kamilsk/retry) ![](https://img.shields.io/github/stars/kamilsk/retry.svg) - The most advanced functional mechanism to perform actions repetitively until successful.
- [retry](https://github.com/percolate/retry) ![](https://img.shields.io/github/stars/percolate/retry.svg) - A simple but highly configurable retry package for Go.
- [retry](https://github.com/thedevsaddam/retry) ![](https://img.shields.io/github/stars/thedevsaddam/retry.svg) - Simple and easy retry mechanism package for Go.
- [retry](https://github.com/shafreeck/retry) ![](https://img.shields.io/github/stars/shafreeck/retry.svg) - A pretty simple library to ensure your work to be done.
- [retry-go](https://github.com/avast/retry-go) ![](https://img.shields.io/github/stars/avast/retry-go.svg) - Simple library for retry mechanism.
- [retry-go](https://github.com/rafaeljesus/retry-go) ![](https://img.shields.io/github/stars/rafaeljesus/retry-go.svg) - Retrying made simple and easy for golang.
- [robustly](https://github.com/VividCortex/robustly) ![](https://img.shields.io/github/stars/VividCortex/robustly.svg) - Runs functions resiliently, catching and restarting panics.
- [rospo](https://github.com/ferama/rospo) ![](https://img.shields.io/github/stars/ferama/rospo.svg) - Simple and reliable ssh tunnels with embedded ssh server in Golang.
- [scan](https://github.com/blockloop/scan) ![](https://img.shields.io/github/stars/blockloop/scan.svg) - Scan golang `sql.Rows` directly to structs, slices, or primitive types.
- [scan](https://github.com/wroge/scan) ![](https://img.shields.io/github/stars/wroge/scan.svg) - Scan sql rows into any type powered by generics.
- [scany](https://github.com/georgysavva/scany) ![](https://img.shields.io/github/stars/georgysavva/scany.svg) - Library for scanning data from a database into Go structs and more.
- [serve](https://github.com/syntaqx/serve) ![](https://img.shields.io/github/stars/syntaqx/serve.svg) - A static http server anywhere you need.
- [sesh](https://github.com/joshmedeski/sesh) ![](https://img.shields.io/github/stars/joshmedeski/sesh.svg) - Sesh is a CLI that helps you create and manage tmux sessions quickly and easily using zoxide.
- [set](https://github.com/nofeaturesonlybugs/set) ![](https://img.shields.io/github/stars/nofeaturesonlybugs/set.svg) - Performant and flexible struct mapping and loose type conversion.
- [shutdown](https://github.com/ztrue/shutdown) ![](https://img.shields.io/github/stars/ztrue/shutdown.svg) - App shutdown hooks for `os.Signal` handling.
- [silk](https://github.com/chrispassas/silk) ![](https://img.shields.io/github/stars/chrispassas/silk.svg) - Read silk netflow files.
- [slice](https://github.com/psampaz/slice) ![](https://img.shields.io/github/stars/psampaz/slice.svg) - Type-safe functions for common Go slice operations.
- [sliceconv](https://github.com/Henry-Sarabia/sliceconv) ![](https://img.shields.io/github/stars/Henry-Sarabia/sliceconv.svg) - Slice conversion between primitive types.
- [slicer](https://github.com/leaanthony/slicer) ![](https://img.shields.io/github/stars/leaanthony/slicer.svg) - Makes working with slices easier.
- [sorty](https://github.com/jfcg/sorty) ![](https://img.shields.io/github/stars/jfcg/sorty.svg) - Fast Concurrent / Parallel Sorting.
- [sqlx](https://github.com/jmoiron/sqlx) ![](https://img.shields.io/github/stars/jmoiron/sqlx.svg) - provides a set of extensions on top of the excellent built-in database/sql package.
- [sqlz](https://github.com/rfberaldo/sqlz) ![](https://img.shields.io/github/stars/rfberaldo/sqlz.svg) - Extension for the database/sql package, adding named queries, struct scanning, and batch operations.
- [sshman](https://github.com/shoobyban/sshman) ![](https://img.shields.io/github/stars/shoobyban/sshman.svg) - SSH Manager for authorized_keys files on multiple remote servers.
- [stacktower](https://github.com/matzehuels/stacktower) ![](https://img.shields.io/github/stars/matzehuels/stacktower.svg) - Visualize dependency graphs as physical tower structures, inspired by XKCD #2347.
- [statiks](https://github.com/janiltonmaciel/statiks) ![](https://img.shields.io/github/stars/janiltonmaciel/statiks.svg) - Fast, zero-configuration, static HTTP filer server.
- [Storm](https://github.com/asdine/storm) ![](https://img.shields.io/github/stars/asdine/storm.svg) - Simple and powerful toolkit for BoltDB.
- [structs](https://github.com/PumpkinSeed/structs) ![](https://img.shields.io/github/stars/PumpkinSeed/structs.svg) - Implement simple functions to manipulate structs.
- [throttle](https://github.com/yudppp/throttle) ![](https://img.shields.io/github/stars/yudppp/throttle.svg) - Throttle is an object that will perform exactly one action per duration.
- [tik](https://github.com/andy2046/tik) ![](https://img.shields.io/github/stars/andy2046/tik.svg) - Simple and easy timing wheel package for Go.
- [tome](https://github.com/cyruzin/tome) ![](https://img.shields.io/github/stars/cyruzin/tome.svg) - Tome was designed to paginate simple RESTful APIs.
- [toolbox](https://github.com/viant/toolbox) ![](https://img.shields.io/github/stars/viant/toolbox.svg) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.
- [UNIS](https://github.com/esemplastic/unis) ![](https://img.shields.io/github/stars/esemplastic/unis.svg) - Common Architecture™ for String Utilities in Go.
- [upterm](https://github.com/owenthereal/upterm) ![](https://img.shields.io/github/stars/owenthereal/upterm.svg) - A tool for developers to share terminal/tmux sessions securely over the web. It’s perfect for remote pair programming, accessing computers behind NATs/firewalls, remote debugging, and more.
- [usql](https://github.com/knq/usql) ![](https://img.shields.io/github/stars/knq/usql.svg) - usql is a universal command-line interface for SQL databases.
- [util](https://github.com/shomali11/util) ![](https://img.shields.io/github/stars/shomali11/util.svg) - Collection of useful utility functions. (strings, concurrency, manipulations, ...).
- [watchhttp](https://github.com/nikolaydubina/watchhttp) ![](https://img.shields.io/github/stars/nikolaydubina/watchhttp.svg) - Run command periodically and expose latest STDOUT or its rich delta as HTTP endpoint.
- [wifiqr](https://github.com/reugn/wifiqr) ![](https://img.shields.io/github/stars/reugn/wifiqr.svg) - Wi-Fi QR Code Generator.
- [wuzz](https://github.com/asciimoo/wuzz) ![](https://img.shields.io/github/stars/asciimoo/wuzz.svg) - Interactive cli tool for HTTP inspection.
- [xferspdy](https://github.com/monmohan/xferspdy) ![](https://img.shields.io/github/stars/monmohan/xferspdy.svg) - Xferspdy provides binary diff and patch library in golang.
- [xpool](https://github.com/peczenyj/xpool) ![](https://img.shields.io/github/stars/peczenyj/xpool.svg) - Yet another golang type safe object pool using generics.
- [yogo](https://github.com/antham/yogo) ![](https://img.shields.io/github/stars/antham/yogo.svg) - Check yopmail mails from command line.

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [fastuuid](https://github.com/rekby/fastuuid) ![](https://img.shields.io/github/stars/rekby/fastuuid.svg) - Fast generate UUIDv4 as string or bytes.
- [goid](https://github.com/jakehl/goid) ![](https://img.shields.io/github/stars/jakehl/goid.svg) - Generate and Parse RFC4122 compliant V4 UUIDs.
- [gouid](https://github.com/twharmon/gouid) ![](https://img.shields.io/github/stars/twharmon/gouid.svg) - Generate cryptographically secure random string IDs with just one allocation.
- [guid](https://github.com/sdrapkin/guid) ![](https://img.shields.io/github/stars/sdrapkin/guid.svg) - Fast cryptographically safe Guid generator for Go (~10x faster than `uuid`).
- [nanoid](https://github.com/aidarkhanov/nanoid) ![](https://img.shields.io/github/stars/aidarkhanov/nanoid.svg) - A tiny and efficient Go unique string ID generator.
- [sno](https://github.com/muyo/sno) ![](https://img.shields.io/github/stars/muyo/sno.svg) - Compact, sortable and fast unique IDs with embedded metadata.
- [ulid](https://github.com/oklog/ulid) ![](https://img.shields.io/github/stars/oklog/ulid.svg) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).
- [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.
- [uuid](https://github.com/agext/uuid) ![](https://img.shields.io/github/stars/agext/uuid.svg) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
- [uuid](https://github.com/gofrs/uuid) ![](https://img.shields.io/github/stars/gofrs/uuid.svg) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.
- [uuid](https://github.com/google/uuid) ![](https://img.shields.io/github/stars/google/uuid.svg) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
- [uuidcheck](https://github.com/ashwingopalsamy/uuidcheck) ![](https://img.shields.io/github/stars/ashwingopalsamy/uuidcheck.svg) - A tiny, dependency-free Go library that validates UUIDs against standard RFC 4122 formatting, converts UUIDv7() into UTC timestamps.
- [wuid](https://github.com/edwingeng/wuid) ![](https://img.shields.io/github/stars/edwingeng/wuid.svg) - An extremely fast globally unique number generator.
- [xid](https://github.com/rs/xid) ![](https://img.shields.io/github/stars/rs/xid.svg) - Xid is a globally unique id generator library, ready to be safely used directly in your server code.

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [checkdigit](https://github.com/osamingo/checkdigit) ![](https://img.shields.io/github/stars/osamingo/checkdigit.svg) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
- [go-validator](https://github.com/tiendc/go-validator) ![](https://img.shields.io/github/stars/tiendc/go-validator.svg) - Validation library using Generics.
- [gody](https://github.com/guiferpa/gody) ![](https://img.shields.io/github/stars/guiferpa/gody.svg) - :balloon: A lightweight struct validator for Go.
- [govalid](https://github.com/twharmon/govalid) ![](https://img.shields.io/github/stars/twharmon/govalid.svg) - Fast, tag-based validation for structs.
- [govalidator](https://github.com/asaskevich/govalidator) ![](https://img.shields.io/github/stars/asaskevich/govalidator.svg) - Validators and sanitizers for strings, numerics, slices and structs.
- [govalidator](https://github.com/thedevsaddam/govalidator) ![](https://img.shields.io/github/stars/thedevsaddam/govalidator.svg) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.
- [hvalid](https://github.com/lyonnee/hvalid) ![](https://img.shields.io/github/stars/lyonnee/hvalid.svg) hvalid is a lightweight validation library written in Go language. It provides a custom validator interface and a series of common validation functions to help developers quickly implement data validation.
- [jio](https://github.com/faceair/jio) ![](https://img.shields.io/github/stars/faceair/jio.svg) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) ![](https://img.shields.io/github/stars/go-ozzo/ozzo-validation.svg) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.
- [validate](https://github.com/gookit/validate) ![](https://img.shields.io/github/stars/gookit/validate.svg) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.
- [validate](https://github.com/gobuffalo/validate) ![](https://img.shields.io/github/stars/gobuffalo/validate.svg) - This package provides a framework for writing validations for Go applications.
- [validator](https://github.com/go-playground/validator) ![](https://img.shields.io/github/stars/go-playground/validator.svg) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
- [Validator](https://github.com/go-the-way/validator) ![](https://img.shields.io/github/stars/go-the-way/validator.svg) - A lightweight model validator written in Go.Contains VFs:Min, Max, MinLength, MaxLength, Length, Enum, Regex.
- [valix](https://github.com/marrow16/valix) ![](https://img.shields.io/github/stars/marrow16/valix.svg) Go package for validating requests
- [Zog](https://github.com/Oudwins/zog) ![](https://img.shields.io/github/stars/Oudwins/zog.svg) - A [Zod](https://github.com/colinhacks/zod) inspired schema builder for runtime value parsing and validation.
  **[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [cli](https://gitlab.com/gitlab-org/cli) - An open-source GitLab command line tool bringing GitLab's cool features to your command line.
- [froggit-go](https://github.com/jfrog/froggit-go) ![](https://img.shields.io/github/stars/jfrog/froggit-go.svg) - Froggit-Go is a Go library, allowing to perform actions on VCS providers.
- [git2go](https://github.com/libgit2/git2go) ![](https://img.shields.io/github/stars/libgit2/git2go.svg) - Go bindings for libgit2.
- [githooks](https://github.com/gabyx/githooks) ![](https://img.shields.io/github/stars/gabyx/githooks.svg) - Per-repo and shared Git hooks with version control and auto update.
- [go-git](https://github.com/go-git/go-git) ![](https://img.shields.io/github/stars/go-git/go-git.svg) - highly extensible Git implementation in pure Go.
- [go-vcs](https://github.com/sourcegraph/go-vcs) ![](https://img.shields.io/github/stars/sourcegraph/go-vcs.svg) - manipulate and inspect VCS repositories in Go.
- [hercules](https://github.com/src-d/hercules) ![](https://img.shields.io/github/stars/src-d/hercules.svg) - gaining advanced insights from Git repository history.
- [hgo](https://github.com/beyang/hgo) ![](https://img.shields.io/github/stars/beyang/hgo.svg) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [gmf](https://github.com/3d0c/gmf) ![](https://img.shields.io/github/stars/3d0c/gmf.svg) - Go bindings for FFmpeg av\* libraries.
- [go-astiav](https://github.com/asticode/go-astiav) ![](https://img.shields.io/github/stars/asticode/go-astiav.svg) - Better C bindings for ffmpeg in GO.
- [go-astisub](https://github.com/asticode/go-astisub) ![](https://img.shields.io/github/stars/asticode/go-astisub.svg) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).
- [go-astits](https://github.com/asticode/go-astits) ![](https://img.shields.io/github/stars/asticode/go-astits.svg) - Parse and demux MPEG Transport Streams (.ts) natively in GO.
- [go-mpd](https://github.com/unki2aut/go-mpd) ![](https://img.shields.io/github/stars/unki2aut/go-mpd.svg) - Parser and generator library for MPEG-DASH manifest files.
- [goav](https://github.com/giorgisio/goav) ![](https://img.shields.io/github/stars/giorgisio/goav.svg) - Comprehensive Go bindings for FFmpeg.
- [gortsplib](https://github.com/aler9/gortsplib) ![](https://img.shields.io/github/stars/aler9/gortsplib.svg) - Pure Go RTSP server and client library.
- [hls-m3u8](https://github.com/Eyevinn/hls-m3u8) ![](https://img.shields.io/github/stars/Eyevinn/hls-m3u8.svg) - Parser and generator for HLS (M3U8) playlists; kept up to date with the spec.
- [libvlc-go](https://github.com/adrg/libvlc-go) ![](https://img.shields.io/github/stars/adrg/libvlc-go.svg) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).
- [mp4ff](https://github.com/Eyevinn/mp4ff) ![](https://img.shields.io/github/stars/Eyevinn/mp4ff.svg) - Library and tools for working with MP4 files containing video, audio, subtitles, or metadata.
- [v4l](https://github.com/korandiz/v4l) ![](https://img.shields.io/github/stars/korandiz/v4l.svg) - Video capture library for Linux, written in Go.

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [Atreugo](https://github.com/savsgio/atreugo) ![](https://img.shields.io/github/stars/savsgio/atreugo.svg) - High performance and extensible micro web framework with zero memory allocations in hot paths.
- [Barf](https://github.com/opensaucerer/barf) ![](https://img.shields.io/github/stars/opensaucerer/barf.svg) - Basically, A Remarkable Framework for building JSON-based web APIs. It is entirely unobtrusive and re-invents no wheel. It is crafted such that getting started is easy and quick while being flexible enough for more complex use cases.
- [Beego](https://github.com/beego/beego) ![](https://img.shields.io/github/stars/beego/beego.svg) - beego is an open-source, high-performance web framework for the Go programming language.
- [Confetti Framework](https://confetti-framework.github.io/docs/) - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.
- [Don](https://github.com/abemedia/go-don) ![](https://img.shields.io/github/stars/abemedia/go-don.svg) - A highly performant and simple to use API framework.
- [Echo](https://github.com/labstack/echo) ![](https://img.shields.io/github/stars/labstack/echo.svg) - High performance, minimalist Go web framework.
- [Fastschema](https://github.com/fastschema/fastschema) ![](https://img.shields.io/github/stars/fastschema/fastschema.svg) - A flexible Go web framework and Headless CMS.
- [Fiber](https://github.com/gofiber/fiber) ![](https://img.shields.io/github/stars/gofiber/fiber.svg) - An Express.js inspired web framework build on Fasthttp.
- [Flamingo](https://github.com/i-love-flamingo/flamingo) ![](https://img.shields.io/github/stars/i-love-flamingo/flamingo.svg) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.
- [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) ![](https://img.shields.io/github/stars/i-love-flamingo/flamingo-commerce.svg) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.
- [Fuego](https://github.com/go-fuego/fuego) ![](https://img.shields.io/github/stars/go-fuego/fuego.svg) - The framework for busy Go developers! Web framework generating OpenAPI 3 spec from source code.
- [Gin](https://github.com/gin-gonic/gin) ![](https://img.shields.io/github/stars/gin-gonic/gin.svg) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
- [Ginrpc](https://github.com/xxjwxc/ginrpc) ![](https://img.shields.io/github/stars/xxjwxc/ginrpc.svg) - Gin parameter automatic binding tool,gin rpc tools.
- [go-api-boot](https://github.com/SaiNageswarS/go-api-boot) ![](https://img.shields.io/github/stars/SaiNageswarS/go-api-boot.svg) - A gRpc-first micro-service framework. Features include ODM support for Mongo, cloud resource support (AWS/Azure/Google), and a fluent dependency injection which is customized for gRpc. Additionally, grpc-web is supported directly, enabling browser access to all gRpc APIs without a proxy.
- [Goa](https://github.com/goadesign/goa) ![](https://img.shields.io/github/stars/goadesign/goa.svg) - Goa provides a holistic approach for developing remote APIs and microservices in Go.
- [GoFr](https://github.com/gofr-dev/gofr) ![](https://img.shields.io/github/stars/gofr-dev/gofr.svg) - Gofr is an opinionated microservice development framework.
- [GoFrame](https://github.com/gogf/gf) ![](https://img.shields.io/github/stars/gogf/gf.svg) - GoFrame is a modular, powerful, high-performance and enterprise-class application development framework of Golang.
- [golamb](https://github.com/twharmon/golamb) ![](https://img.shields.io/github/stars/twharmon/golamb.svg) - Golamb makes it easier to write API endpoints for use with AWS Lambda and API Gateway.
- [Gone](https://github.com/gone-io/gone) ![](https://img.shields.io/github/stars/gone-io/gone.svg) - A lightweight dependency injection and web framework inspired by Spring.
- [goravel](https://github.com/goravel/goravel) ![](https://img.shields.io/github/stars/goravel/goravel.svg) - A Laravel-inspired web framework with ORM, authentication, queue, task scheduling, and more built-in features.
- [Goyave](https://github.com/go-goyave/goyave) ![](https://img.shields.io/github/stars/go-goyave/goyave.svg) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities.
- [Hertz](https://github.com/cloudwego/hertz) ![](https://img.shields.io/github/stars/cloudwego/hertz.svg) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices.
- [hiboot](https://github.com/hidevopsio/hiboot) ![](https://img.shields.io/github/stars/hidevopsio/hiboot.svg) - hiboot is a high performance web application framework with auto configuration and dependency injection support.
- [Huma](https://github.com/danielgtaylor/huma/) ![](https://img.shields.io/github/stars/danielgtaylor/huma.svg) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI.
- [iWF](https://github.com/indeedeng/iwf) ![](https://img.shields.io/github/stars/indeedeng/iwf.svg) - iWF is an all-in-one platform for developing long-running business processes. It offers a convenient abstraction for utilizing databases, ElasticSearch, message queues, durable timers, and more, with a clean, simple, and user-friendly interface.
- [Lit](https://github.com/jvcoutinho/lit) ![](https://img.shields.io/github/stars/jvcoutinho/lit.svg) - Highly performant declarative web framework for Golang, aiming for simplicity and quality of life.
- [Microservice](https://github.com/claygod/microservice) ![](https://img.shields.io/github/stars/claygod/microservice.svg) - The framework for the creation of microservices, written in Golang.
- [patron](https://github.com/beatlabs/patron) ![](https://img.shields.io/github/stars/beatlabs/patron.svg) - Patron is a microservice framework following best cloud practices with a focus on productivity.
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux) - Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Revel](https://github.com/revel/revel) ![](https://img.shields.io/github/stars/revel/revel.svg) - High-productivity web framework for the Go language.
- [rk-boot](https://github.com/rookie-ninja/rk-boot) ![](https://img.shields.io/github/stars/rookie-ninja/rk-boot.svg) - A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily.
- [Ronykit](https://github.com/clubpay/ronykit) ![](https://img.shields.io/github/stars/clubpay/ronykit.svg) - Web framework with pluggable architecture and very performant.
- [rux](https://github.com/gookit/rux) ![](https://img.shields.io/github/stars/gookit/rux.svg) - Simple and fast web framework for build golang HTTP applications.
- [templui](https://github.com/axzilla/templui) ![](https://img.shields.io/github/stars/axzilla/templui.svg) - Modern UI Components for Go & Templ.
- [uAdmin](https://github.com/uadmin/uadmin) ![](https://img.shields.io/github/stars/uadmin/uadmin.svg) - Fully featured web framework for Golang, inspired by Django.
- [WebGo](https://github.com/naughtygopher/webgo) ![](https://img.shields.io/github/stars/naughtygopher/webgo.svg) - A micro-framework to build web apps with handler chaining, middleware, and context injection. With standard library-compliant HTTP handlers (i.e., `http.HandlerFunc`)..
- [Xun](https://github.com/yaitoo/xun) ![](https://img.shields.io/github/stars/yaitoo/xun.svg) - Web framework built on Go's built-in html/template and net/http package’s router. It is designed to be lightweight, fast, and easy to use while providing a simple and intuitive API for building web applications with advanced features such as middleware, routing, and template rendering.
- [Yokai](https://github.com/ankorstore/yokai) ![](https://img.shields.io/github/stars/ankorstore/yokai.svg) - Simple, modular, and observable Go framework for backend applications.

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [client-timing](https://github.com/posener/client-timing) ![](https://img.shields.io/github/stars/posener/client-timing.svg) - An HTTP client for Server-Timing header.
- [CORS](https://github.com/rs/cors) ![](https://img.shields.io/github/stars/rs/cors.svg) - Easily add CORS capabilities to your API.
- [echo-middleware](https://github.com/faabiosr/echo-middleware) ![](https://img.shields.io/github/stars/faabiosr/echo-middleware.svg) - Middleware for Echo framework with logging and metrics.
- [formjson](https://github.com/rs/formjson) ![](https://img.shields.io/github/stars/rs/formjson.svg) - Transparently handle JSON input as a standard form POST.
- [go-fault](https://github.com/github/go-fault) ![](https://img.shields.io/github/stars/github/go-fault.svg) - Fault injection middleware for Go.
- [Limiter](https://github.com/ulule/limiter) ![](https://img.shields.io/github/stars/ulule/limiter.svg) - Dead simple rate limit middleware for Go.
- [ln-paywall](https://github.com/philippgille/ln-paywall) ![](https://img.shields.io/github/stars/philippgille/ln-paywall.svg) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).
- [mid](https://github.com/bobg/mid) ![](https://img.shields.io/github/stars/bobg/mid.svg) - Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more.
- [rk-gin](https://github.com/rookie-ninja/rk-gin) ![](https://img.shields.io/github/stars/rookie-ninja/rk-gin.svg) - Middleware for Gin framework with logging, metrics, auth, tracing etc.
- [rk-grpc](https://github.com/rookie-ninja/rk-grpc) ![](https://img.shields.io/github/stars/rookie-ninja/rk-grpc.svg) - Middleware for gRPC with logging, metrics, auth, tracing etc.
- [Tollbooth](https://github.com/didip/tollbooth) ![](https://img.shields.io/github/stars/didip/tollbooth.svg) - Rate limit HTTP request handler.
- [XFF](https://github.com/sebest/xff) ![](https://img.shields.io/github/stars/sebest/xff.svg) - Handle `X-Forwarded-For` header and friends.

#### Libraries for creating HTTP middlewares

- [alice](https://github.com/justinas/alice) ![](https://img.shields.io/github/stars/justinas/alice.svg) - Painless middleware chaining for Go.
- [catena](https://github.com/codemodus/catena) ![](https://img.shields.io/github/stars/codemodus/catena.svg) - http.Handler wrapper catenation (same API as "chain").
- [chain](https://github.com/codemodus/chain) ![](https://img.shields.io/github/stars/codemodus/chain.svg) - Handler wrapper chaining with scoped data (net/context-based "middleware").
- [gores](https://github.com/alioygur/gores) ![](https://img.shields.io/github/stars/alioygur/gores.svg) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
- [interpose](https://github.com/carbocation/interpose) ![](https://img.shields.io/github/stars/carbocation/interpose.svg) - Minimalist net/http middleware for golang.
- [mediary](https://github.com/HereMobilityDevelopers/mediary) ![](https://img.shields.io/github/stars/HereMobilityDevelopers/mediary.svg) - add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses.
- [muxchain](https://github.com/stephens2424/muxchain) ![](https://img.shields.io/github/stars/stephens2424/muxchain.svg) - Lightweight middleware for net/http.
- [negroni](https://github.com/urfave/negroni) ![](https://img.shields.io/github/stars/urfave/negroni.svg) - Idiomatic HTTP middleware for Golang.
- [render](https://github.com/unrolled/render) ![](https://img.shields.io/github/stars/unrolled/render.svg) - Go package for easily rendering JSON, XML, and HTML template responses.
- [renderer](https://github.com/thedevsaddam/renderer) ![](https://img.shields.io/github/stars/thedevsaddam/renderer.svg) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.
- [stats](https://github.com/thoas/stats) ![](https://img.shields.io/github/stars/thoas/stats.svg) - Go middleware that stores various information about your web application.

**[⬆ back to top](#contents)**

### Routers

- [alien](https://github.com/gernest/alien) ![](https://img.shields.io/github/stars/gernest/alien.svg) - Lightweight and fast http router from outer space.
- [bellt](https://github.com/GuilhermeCaruso/bellt) ![](https://img.shields.io/github/stars/GuilhermeCaruso/bellt.svg) - A simple Go HTTP router.
- [Bone](https://github.com/go-zoo/bone) ![](https://img.shields.io/github/stars/go-zoo/bone.svg) - Lightning Fast HTTP Multiplexer.
- [Bxog](https://github.com/claygod/Bxog) ![](https://img.shields.io/github/stars/claygod/Bxog.svg) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
- [chi](https://github.com/go-chi/chi) ![](https://img.shields.io/github/stars/go-chi/chi.svg) - Small, fast and expressive HTTP router built on net/context.
- [fasthttprouter](https://github.com/buaazp/fasthttprouter) ![](https://img.shields.io/github/stars/buaazp/fasthttprouter.svg) - High performance router forked from `httprouter`. The first router fit for `fasthttp`.
- [FastRouter](https://github.com/razonyang/fastrouter) ![](https://img.shields.io/github/stars/razonyang/fastrouter.svg) - a fast, flexible HTTP router written in Go.
- [goblin](https://github.com/bmf-san/goblin) ![](https://img.shields.io/github/stars/bmf-san/goblin.svg) - A golang http router based on trie tree.
- [gocraft/web](https://github.com/gocraft/web) ![](https://img.shields.io/github/stars/gocraft/web.svg) - Mux and middleware package in Go.
- [Goji](https://github.com/goji/goji) ![](https://img.shields.io/github/stars/goji/goji.svg) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.
- [GoLobby/Router](https://github.com/golobby/router) ![](https://img.shields.io/github/stars/golobby/router.svg) - GoLobby Router is a lightweight yet powerful HTTP router for the Go programming language.
- [goroute](https://github.com/goroute/route) ![](https://img.shields.io/github/stars/goroute/route.svg) - Simple yet powerful HTTP request multiplexer.
- [GoRouter](https://github.com/vardius/gorouter) ![](https://img.shields.io/github/stars/vardius/gorouter.svg) - GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
- [gowww/router](https://github.com/gowww/router) ![](https://img.shields.io/github/stars/gowww/router.svg) - Lightning fast HTTP router fully compatible with the net/http.Handler interface.
- [httprouter](https://github.com/julienschmidt/httprouter) ![](https://img.shields.io/github/stars/julienschmidt/httprouter.svg) - High performance router. Use this and the standard http handlers to form a very high performance web framework.
- [httptreemux](https://github.com/dimfeld/httptreemux) ![](https://img.shields.io/github/stars/dimfeld/httptreemux.svg) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.
- [lars](https://github.com/go-playground/lars) ![](https://img.shields.io/github/stars/go-playground/lars.svg) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.
- [mux](https://github.com/gorilla/mux) ![](https://img.shields.io/github/stars/gorilla/mux.svg) - Powerful URL router and dispatcher for golang.
- [nchi](https://github.com/muir/nchi) ![](https://img.shields.io/github/stars/muir/nchi.svg) - chi-like router built on httprouter with dependency injection based middleware wrappers
- [ngamux](https://github.com/ngamux/ngamux) ![](https://img.shields.io/github/stars/ngamux/ngamux.svg) - Simple HTTP router for Go.
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) ![](https://img.shields.io/github/stars/go-ozzo/ozzo-routing.svg) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.
- [pure](https://github.com/go-playground/pure) ![](https://img.shields.io/github/stars/go-playground/pure.svg) - Is a lightweight HTTP router that sticks to the std "net/http" implementation.
- [Siesta](https://github.com/VividCortex/siesta) ![](https://img.shields.io/github/stars/VividCortex/siesta.svg) - Composable framework to write middleware and handlers.
- [vestigo](https://github.com/husobee/vestigo) ![](https://img.shields.io/github/stars/husobee/vestigo.svg) - Performant, stand-alone, HTTP compliant URL Router for go web applications.
- [violetear](https://github.com/nbari/violetear) ![](https://img.shields.io/github/stars/nbari/violetear.svg) - Go HTTP router.
- [xmux](https://github.com/rs/xmux) ![](https://img.shields.io/github/stars/rs/xmux.svg) - High performance muxer based on `httprouter` with `net/context` support.
- [xujiajun/gorouter](https://github.com/xujiajun/gorouter) ![](https://img.shields.io/github/stars/xujiajun/gorouter.svg) - A simple and fast HTTP router for Go.

**[⬆ back to top](#contents)**

## WebAssembly

- [dom](https://github.com/dennwc/dom) ![](https://img.shields.io/github/stars/dennwc/dom.svg) - DOM library.
- [Extism Go SDK](https://github.com/extism/go-sdk) ![](https://img.shields.io/github/stars/extism/go-sdk.svg) - Universal, cross-language WebAssembly framework for building plug-in systems and polyglot apps.
- [go-canvas](https://github.com/markfarnan/go-canvas) ![](https://img.shields.io/github/stars/markfarnan/go-canvas.svg) - Library to use HTML5 Canvas, with all drawing within go code.
- [tinygo](https://github.com/tinygo-org/tinygo) ![](https://img.shields.io/github/stars/tinygo-org/tinygo.svg) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.
- [vert](https://github.com/norunners/vert) ![](https://img.shields.io/github/stars/norunners/vert.svg) - Interop between Go and JS values.
- [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) ![](https://img.shields.io/github/stars/agnivade/wasmbrowsertest.svg) - Run Go WASM tests in your browser.
- [webapi](https://github.com/gowebapi/webapi) ![](https://img.shields.io/github/stars/gowebapi/webapi.svg) - Bindings for DOM and HTML generated from WebIDL.

**[⬆ back to top](#contents)**

## Webhooks Server

- [webhook](https://github.com/adnanh/webhook) ![](https://img.shields.io/github/stars/adnanh/webhook.svg) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.
- [webhooked](https://github.com/42Atomys/webhooked) ![](https://img.shields.io/github/stars/42Atomys/webhooked.svg) - A webhook receiver on steroids: handle, secure, format and store a Webhook payload has never been easier.
- [WebhookX](https://github.com/webhookx-io/webhookx) ![](https://img.shields.io/github/stars/webhookx-io/webhookx.svg) - A webhooks gateway for message receiving, processing, and reliable delivering.

**[⬆ back to top](#contents)**

## Windows

- [d3d9](https://github.com/gonutz/d3d9) ![](https://img.shields.io/github/stars/gonutz/d3d9.svg) - Go bindings for Direct3D9.
- [go-ole](https://github.com/go-ole/go-ole) ![](https://img.shields.io/github/stars/go-ole/go-ole.svg) - Win32 OLE implementation for golang.
- [gosddl](https://github.com/MonaxGT/gosddl) ![](https://img.shields.io/github/stars/MonaxGT/gosddl.svg) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

**[⬆ back to top](#contents)**

## Workflow Frameworks

_Libraries for creating Workflows._

- [Cadence-client](https://github.com/uber-go/cadence-client) ![](https://img.shields.io/github/stars/uber-go/cadence-client.svg) - A framework for authoring workflows and activities running on top of the Cadence orchestration engine made by Uber.
- [Dagu](https://github.com/dagu-go/dagu) ![](https://img.shields.io/github/stars/dagu-go/dagu.svg) - No-code workflow executor. it executes DAGs defined in a simple YAML format.
- [go-dag](https://github.com/rhosocial/go-dag) ![](https://img.shields.io/github/stars/rhosocial/go-dag.svg) - A framework developed in Go that manages the execution of workflows described by directed acyclic graphs.
- [go-taskflow](https://github.com/noneback/go-taskflow) ![](https://img.shields.io/github/stars/noneback/go-taskflow.svg) - A taskflow-like General-purpose Task-parallel Programming Framework with integrated visualizer and profiler.
- [workflow](https://github.com/luno/workflow) ![](https://img.shields.io/github/stars/luno/workflow.svg) - A tech stack agnostic Event Driven Workflow framework.

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._

- [XML-Comp](https://github.com/xml-comp/xml-comp) ![](https://img.shields.io/github/stars/xml-comp/xml-comp.svg) - Simple command line XML comparer that generates diffs of folders, files and tags.
- [xml2map](https://github.com/sbabiv/xml2map) ![](https://img.shields.io/github/stars/sbabiv/xml2map.svg) - XML to MAP converter written Golang.
- [xmlquery](https://github.com/antchfx/xmlquery) ![](https://img.shields.io/github/stars/antchfx/xmlquery.svg) - xmlquery is Golang XPath package for XML query.
- [xmlwriter](https://github.com/shabbyrobe/xmlwriter) ![](https://img.shields.io/github/stars/shabbyrobe/xmlwriter.svg) - Procedural XML generation API based on libxml2's xmlwriter module.
- [xpath](https://github.com/antchfx/xpath) ![](https://img.shields.io/github/stars/antchfx/xpath.svg) - XPath package for Go.
- [zek](https://github.com/miku/zek) ![](https://img.shields.io/github/stars/miku/zek.svg) - Generate a Go struct from XML.

## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign](https://github.com/sigstore/cosign) ![](https://img.shields.io/github/stars/sigstore/cosign.svg) - Container Signing, Verification and Storage in an OCI registry.
- [in-toto](https://github.com/in-toto/in-toto-golang) ![](https://img.shields.io/github/stars/in-toto/in-toto-golang.svg) - Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation.
- [OpenZiti](https://github.com/openziti/ziti) ![](https://img.shields.io/github/stars/openziti/ziti.svg) - A full, open source zero trust overlay network. Including numerous SDKs for numerous languages such as [golang](https://github.com/openziti/sdk-golang) allowing you to embed zero trust principles directly into your applications. The [OpenZiti Test Kitchen](https://github.com/openziti-test-kitchen) has numerous examples to draw inspiration from including a [zero trust ssh client - zssh](https://github.com/openziti-test-kitchen/zssh)
- [Spiffe-Vault](https://github.com/philips-labs/spiffe-vault) ![](https://img.shields.io/github/stars/philips-labs/spiffe-vault.svg) - Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication.
- [Spire](https://github.com/spiffe/spire) ![](https://img.shields.io/github/stars/spiffe/spire.svg) - SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms.

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [apicompat](https://github.com/bradleyfalzon/apicompat) ![](https://img.shields.io/github/stars/bradleyfalzon/apicompat.svg) - Checks recent changes to a Go project for backwards incompatible changes.
- [asty](https://github.com/asty-org/asty) ![](https://img.shields.io/github/stars/asty-org/asty.svg) - Converts golang AST to JSON and JSON to AST.
- [blanket](https://gitlab.com/verygoodsoftwarenotvirus/blanket) - blanket is a tool that helps you catch functions which don't have direct unit tests in your Go packages.
- [ChainJacking](https://github.com/Checkmarx/chainjacking) ![](https://img.shields.io/github/stars/Checkmarx/chainjacking.svg) - Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack.
- [Chronos](https://github.com/amit-davidson/Chronos) ![](https://img.shields.io/github/stars/amit-davidson/Chronos.svg) - Detects race conditions statically
- [dupl](https://github.com/mibk/dupl) ![](https://img.shields.io/github/stars/mibk/dupl.svg) - Tool for code clone detection.
- [errcheck](https://github.com/kisielk/errcheck) ![](https://img.shields.io/github/stars/kisielk/errcheck.svg) - Errcheck is a program for checking for unchecked errors in Go programs.
- [fatcontext](https://github.com/Crocmagnon/fatcontext) ![](https://img.shields.io/github/stars/Crocmagnon/fatcontext.svg) - Fatcontext detects nested contexts in loops or function literals.
- [go-checkstyle](https://github.com/qiniu/checkstyle) ![](https://img.shields.io/github/stars/qiniu/checkstyle.svg) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) ![](https://img.shields.io/github/stars/roblaszczak/go-cleanarch.svg) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.
- [go-critic](https://github.com/go-critic/go-critic) ![](https://img.shields.io/github/stars/go-critic/go-critic.svg) - source code linter that brings checks that are currently not implemented in other linters.
- [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) ![](https://img.shields.io/github/stars/psampaz/go-mod-outdated.svg) - An easy way to find outdated dependencies of your Go projects.
- [goast-viewer](https://github.com/yuroyoro/goast-viewer) ![](https://img.shields.io/github/stars/yuroyoro/goast-viewer.svg) - Web based Golang AST visualizer.
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
- [golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) ![](https://img.shields.io/github/stars/arxdsilva/golang-ifood-sdk.svg) - iFood API SDK.
- [golangci-lint](https://github.com/golangci/golangci-lint) ![](https://img.shields.io/github/stars/golangci/golangci-lint.svg) – A fast Go linters runner. It runs linters in parallel, uses caching, supports `yaml` config, has integrations with all major IDE and has dozens of linters included.
- [golines](https://github.com/segmentio/golines) ![](https://img.shields.io/github/stars/segmentio/golines.svg) - Formatter that automatically shortens long lines in Go code.
- [GoPlantUML](https://github.com/jfeliu007/goplantuml) ![](https://img.shields.io/github/stars/jfeliu007/goplantuml.svg) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.
- [goreturns](https://github.com/sqs/goreturns) ![](https://img.shields.io/github/stars/sqs/goreturns.svg) - Adds zero-value return statements to match the func return types.
- [gostatus](https://github.com/shurcooL/gostatus) ![](https://img.shields.io/github/stars/shurcooL/gostatus.svg) - Command line tool, shows the status of repositories that contain Go packages.
- [lint](https://github.com/surullabs/lint) ![](https://img.shields.io/github/stars/surullabs/lint.svg) - Run linters as part of go test.
- [php-parser](https://github.com/z7zmey/php-parser) ![](https://img.shields.io/github/stars/z7zmey/php-parser.svg) - A Parser for PHP written in Go.
- [revive](https://github.com/mgechev/revive) ![](https://img.shields.io/github/stars/mgechev/revive.svg) – ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for `golint`.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) ![](https://img.shields.io/github/stars/dominikh/go-tools.svg) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
- [testifylint](https://github.com/Antonboom/testifylint) ![](https://img.shields.io/github/stars/Antonboom/testifylint.svg) – A linter that checks usage of [github.com/stretchr/testify](https://github.com/stretchr/testify).
- [tickgit](https://github.com/augmentable-dev/tickgit) ![](https://img.shields.io/github/stars/augmentable-dev/tickgit.svg) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.
- [todocheck](https://github.com/preslavmihaylov/todocheck) ![](https://img.shields.io/github/stars/preslavmihaylov/todocheck.svg) - Static code analyser which links TODO comments in code with issues in your issue tracker.
- [unconvert](https://github.com/mdempsky/unconvert) ![](https://img.shields.io/github/stars/mdempsky/unconvert.svg) - Remove unnecessary type conversions from Go source.
- [usestdlibvars](https://github.com/sashamelentyev/usestdlibvars) ![](https://img.shields.io/github/stars/sashamelentyev/usestdlibvars.svg) - A linter that detect the possibility to use variables/constants from the Go standard library.
- [vacuum](https://github.com/daveshanley/vacuum) ![](https://img.shields.io/github/stars/daveshanley/vacuum.svg) - An ultra-super-fast, lightweight OpenAPI linter and quality checking tool.
- [validate](https://github.com/mccoyst/validate) ![](https://img.shields.io/github/stars/mccoyst/validate.svg) - Automatically validates struct fields with tags.
- [wrapcheck](https://github.com/tomarrell/wrapcheck) ![](https://img.shields.io/github/stars/tomarrell/wrapcheck.svg) - A linter to check that errors from external packages are wrapped.

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [coc-go language server extension for Vim/Neovim](https://github.com/josa42/coc-go) ![](https://img.shields.io/github/stars/josa42/coc-go.svg) - This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim.
- [Go Doc](https://github.com/msyrus/vscode-go-doc) ![](https://img.shields.io/github/stars/msyrus/vscode-go-doc.svg) - A Visual Studio Code extension for showing definition in output and generating go doc.
- [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
- [go-mode](https://github.com/dominikh/go-mode.el) ![](https://img.shields.io/github/stars/dominikh/go-mode.el.svg) - Go mode for GNU/Emacs.
- [gocode](https://github.com/nsf/gocode) ![](https://img.shields.io/github/stars/nsf/gocode.svg) - Autocompletion daemon for the Go programming language.
- [goimports-reviser](https://github.com/incu6us/goimports-reviser) ![](https://img.shields.io/github/stars/incu6us/goimports-reviser.svg) - Formatting tool for imports.
- [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
- [GoSublime](https://github.com/DisposaBoy/GoSublime) ![](https://img.shields.io/github/stars/DisposaBoy/GoSublime.svg) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.
- [gounit-vim](https://github.com/hexdigest/gounit-vim) ![](https://img.shields.io/github/stars/hexdigest/gounit-vim.svg) - Vim plugin for generating Go tests based on the function's or method's signature.
- [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) ![](https://img.shields.io/github/stars/rjohnsondev/vim-compiler-go.svg) - Vim plugin to highlight syntax errors on save.
- [vim-go](https://github.com/fatih/vim-go) ![](https://img.shields.io/github/stars/fatih/vim-go.svg) - Go development plugin for Vim.
- [vscode-go](https://github.com/golang/vscode-go) ![](https://img.shields.io/github/stars/golang/vscode-go.svg) - Extension for Visual Studio Code (VS Code) which provides support for the Go language.
- [Watch](https://github.com/eaburns/Watch) ![](https://img.shields.io/github/stars/eaburns/Watch.svg) - Runs a command in an acme win on file changes.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [envdoc](https://github.com/g4s8/envdoc) ![](https://img.shields.io/github/stars/g4s8/envdoc.svg) - generate documentation for environment variables from Go source files.
- [generic](https://github.com/usk81/generic) ![](https://img.shields.io/github/stars/usk81/generic.svg) - flexible data type for Go.
- [gocontracts](https://github.com/Parquery/gocontracts) ![](https://img.shields.io/github/stars/Parquery/gocontracts.svg) - brings design-by-contract to Go by synchronizing the code with the documentation.
- [godal](https://github.com/mafulong/godal) ![](https://img.shields.io/github/stars/mafulong/godal.svg) - Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm.
- [gonerics](https://github.com/bouk/gonerics) ![](https://img.shields.io/github/stars/bouk/gonerics.svg) - Idiomatic Generics in Go.
- [gotests](https://github.com/cweill/gotests) ![](https://img.shields.io/github/stars/cweill/gotests.svg) - Generate Go tests from your source code.
- [gounit](https://github.com/hexdigest/gounit) ![](https://img.shields.io/github/stars/hexdigest/gounit.svg) - Generate Go tests using your own templates.
- [hasgo](https://github.com/DylanMeeus/hasgo) ![](https://img.shields.io/github/stars/DylanMeeus/hasgo.svg) - Generate Haskell inspired functions for your slices.
- [options-gen](https://github.com/kazhuravlev/options-gen) ![](https://img.shields.io/github/stars/kazhuravlev/options-gen.svg) - Functional options described by Dave Cheney's post "Functional options for friendly APIs".
- [re2dfa](https://gitlab.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code.
- [sqlgen](https://github.com/anqiansong/sqlgen) ![](https://img.shields.io/github/stars/anqiansong/sqlgen.svg) - Generate gorm, xorm, sqlx, bun, sql code from SQL file or DSN.
- [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.
- [xgen](https://github.com/xuri/xgen) ![](https://img.shields.io/github/stars/xuri/xgen.svg) - XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator.

**[⬆ back to top](#contents)**

## Go Tools

- [decouple](https://github.com/bobg/decouple) ![](https://img.shields.io/github/stars/bobg/decouple.svg) - Find “overspecified” function parameters that could be generalized with interface types.
- [docs](https://github.com/go-oas/docs) ![](https://img.shields.io/github/stars/go-oas/docs.svg) - Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard.
- [go-callvis](https://github.com/TrueFurby/go-callvis) ![](https://img.shields.io/github/stars/TrueFurby/go-callvis.svg) - Visualize call graph of your Go program using dot format.
- [go-size-analyzer](https://github.com/Zxilly/go-size-analyzer) ![](https://img.shields.io/github/stars/Zxilly/go-size-analyzer.svg) - Analyze and visualize the size of dependencies in compiled Golang binaries, providing insight into their impact on the final build.
- [go-swagger](https://github.com/go-swagger/go-swagger) ![](https://img.shields.io/github/stars/go-swagger/go-swagger.svg) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.
- [go-template-playground](https://bartventer.github.io/go-template-playground/) - An interactive environment to create and test Go templates.
- [godbg](https://github.com/tylerwince/godbg) ![](https://img.shields.io/github/stars/tylerwince/godbg.svg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development.
- [gomodrun](https://github.com/dustinblackman/gomodrun/) ![](https://img.shields.io/github/stars/dustinblackman/gomodrun.svg) - Go tool that executes and caches binaries included in go.mod files.
- [gotemplate.io](https://gotemplate.io/) - Online tool to preview `text/template` templates live.
- [gotestdox](https://github.com/bitfield/gotestdox) ![](https://img.shields.io/github/stars/bitfield/gotestdox.svg) - Show Go test results as readable sentences.
- [gothanks](https://github.com/psampaz/gothanks) ![](https://img.shields.io/github/stars/psampaz/gothanks.svg) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.
- [gotutor](https://github.com/ahmedakef/gotutor) ![](https://img.shields.io/github/stars/ahmedakef/gotutor.svg) - Online Go Debugger & Visualizer.
- [igo](https://github.com/rocketlaunchr/igo) ![](https://img.shields.io/github/stars/rocketlaunchr/igo.svg) - An igo to go transpiler (new language features for Go language!)
- [modver](https://github.com/bobg/modver) ![](https://img.shields.io/github/stars/bobg/modver.svg) - Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules.
- [MoniGO](https://github.com/iyashjayesh/monigo) ![](https://img.shields.io/github/stars/iyashjayesh/monigo.svg) - A performance monitoring library for Go applications. It provides real-time insights into application performance! 🚀
- [OctoLinker](https://github.com/OctoLinker/browser-extension) ![](https://img.shields.io/github/stars/OctoLinker/browser-extension.svg) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
- [richgo](https://github.com/kyoh86/richgo) ![](https://img.shields.io/github/stars/kyoh86/richgo.svg) - Enrich `go test` outputs with text decorations.
- [roumon](https://github.com/becheran/roumon) ![](https://img.shields.io/github/stars/becheran/roumon.svg) - Monitor current state of all active goroutines via a command line interface.
- [rts](https://github.com/galeone/rts) ![](https://img.shields.io/github/stars/galeone/rts.svg) - RTS: response to struct. Generates Go structs from server responses.
- [textra](https://github.com/ravsii/textra) ![](https://img.shields.io/github/stars/ravsii/textra.svg) - Extract Go struct field names, types and tags for filtering and exporting.
- [typex](https://github.com/dtgorski/typex) ![](https://img.shields.io/github/stars/dtgorski/typex.svg) - Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration.

**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [abbreviate](https://github.com/dnnrly/abbreviate) ![](https://img.shields.io/github/stars/dnnrly/abbreviate.svg) - abbreviate is a tool turning long strings in to shorter ones with configurable separators, for example to embed branch names in to deployment stack IDs.
- [alaz](https://github.com/ddosify/alaz) ![](https://img.shields.io/github/stars/ddosify/alaz.svg) - Effortless, Low-Overhead, eBPF-based Kubernetes Monitoring.
- [aptly](https://github.com/aptly-dev/aptly) ![](https://img.shields.io/github/stars/aptly-dev/aptly.svg) - aptly is a Debian repository management tool.
- [aurora](https://github.com/xuri/aurora) ![](https://img.shields.io/github/stars/xuri/aurora.svg) - Cross-platform web-based Beanstalkd queue server console.
- [awsenv](https://github.com/soniah/awsenv) ![](https://img.shields.io/github/stars/soniah/awsenv.svg) - Small binary that loads Amazon (AWS) environment variables for a profile.
- [Balerter](https://github.com/balerter/balerter) ![](https://img.shields.io/github/stars/balerter/balerter.svg) - A self-hosted script-based alerting manager.
- [Blast](https://github.com/dave/blast) ![](https://img.shields.io/github/stars/dave/blast.svg) - A simple tool for API load testing and batch jobs.
- [bombardier](https://github.com/codesenberg/bombardier) ![](https://img.shields.io/github/stars/codesenberg/bombardier.svg) - Fast cross-platform HTTP benchmarking tool.
- [cassowary](https://github.com/rogerwelin/cassowary) ![](https://img.shields.io/github/stars/rogerwelin/cassowary.svg) - Modern cross-platform HTTP load-testing tool written in Go.
- [chaosmonkey](https://github.com/Netflix/chaosmonkey) ![](https://img.shields.io/github/stars/Netflix/chaosmonkey.svg) - A resiliency tool that helps applications tolerate random instance failures.
- [Ddosify](https://github.com/ddosify/ddosify) ![](https://img.shields.io/github/stars/ddosify/ddosify.svg) - High-performance load testing tool, written in Golang.
- [decompose](https://github.com/s0rg/decompose) ![](https://img.shields.io/github/stars/s0rg/decompose.svg) - tool to generate and process Docker containers connections graphs.
- [DepCharge](https://github.com/centerorbit/depcharge) ![](https://img.shields.io/github/stars/centerorbit/depcharge.svg) - Helps orchestrating the execution of commands across the many dependencies in larger projects.
- [dish](https://github.com/thevxn/dish) ![](https://img.shields.io/github/stars/thevxn/dish.svg) - A lightweight, remotely configurable monitoring service.
- [Docker](https://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
- [docker-go-mingw](https://github.com/x1unix/docker-go-mingw) ![](https://img.shields.io/github/stars/x1unix/docker-go-mingw.svg) - Docker image for building Go binaries for Windows with MinGW toolchain.
- [docker-volume-backup](https://github.com/offen/docker-volume-backup) ![](https://img.shields.io/github/stars/offen/docker-volume-backup.svg) - Backup Docker volumes locally or to any S3, WebDAV, Azure Blob Storage, Dropbox or SSH compatible storage.
- [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) ![](https://img.shields.io/github/stars/ozankasikci/dockerfile-generator.svg) - A go library and an executable that produces valid Dockerfiles using various input channels.
- [dogo](https://github.com/liudng/dogo) ![](https://img.shields.io/github/stars/liudng/dogo.svg) - Monitoring changes in the source file and automatically compile and run (restart).
- [drone-jenkins](https://github.com/appleboy/drone-jenkins) ![](https://img.shields.io/github/stars/appleboy/drone-jenkins.svg) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
- [drone-scp](https://github.com/appleboy/drone-scp) ![](https://img.shields.io/github/stars/appleboy/drone-scp.svg) - Copy files and artifacts via SSH using a binary, docker or Drone CI.
- [Dropship](https://github.com/chrismckenzie/dropship) ![](https://img.shields.io/github/stars/chrismckenzie/dropship.svg) - Tool for deploying code via cdn.
- [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) ![](https://img.shields.io/github/stars/appleboy/easyssh-proxy.svg) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
- [fac](https://github.com/mkchoi212/fac) ![](https://img.shields.io/github/stars/mkchoi212/fac.svg) - Command-line user interface to fix git merge conflicts.
- [Flannel](https://github.com/flannel-io/flannel) ![](https://img.shields.io/github/stars/flannel-io/flannel.svg) - Flannel is a network fabric for containers, designed for Kubernetes.
- [Fleet device management](https://github.com/fleetdm/fleet) ![](https://img.shields.io/github/stars/fleetdm/fleet.svg) - Lightweight, programmable telemetry for servers and workstations.
- [gaia](https://github.com/gaia-pipeline/gaia) ![](https://img.shields.io/github/stars/gaia-pipeline/gaia.svg) - Build powerful pipelines in any programming language.
- [ghorg](https://github.com/gabrie30/ghorg) ![](https://img.shields.io/github/stars/gabrie30/ghorg.svg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket.
- [Gitea](https://github.com/go-gitea/gitea) ![](https://img.shields.io/github/stars/go-gitea/gitea.svg) - Fork of Gogs, entirely community driven.
- [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
- [go-furnace](https://github.com/go-furnace/go-furnace) ![](https://img.shields.io/github/stars/go-furnace/go-furnace.svg) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.
- [go-rocket-update](https://github.com/mouuff/go-rocket-update) ![](https://img.shields.io/github/stars/mouuff/go-rocket-update.svg) - A simple way to make self updating Go applications - Supports Github and Gitlab.
- [go-selfupdate](https://github.com/sanbornm/go-selfupdate) ![](https://img.shields.io/github/stars/sanbornm/go-selfupdate.svg) - Enable your Go applications to self update.
- [gobrew](https://github.com/cryptojuice/gobrew) ![](https://img.shields.io/github/stars/cryptojuice/gobrew.svg) - gobrew lets you easily switch between multiple versions of go.
- [gobrew](https://github.com/kevincobain2000/gobrew) ![](https://img.shields.io/github/stars/kevincobain2000/gobrew.svg) - Go version manager. Super simple tool to install and manage Go versions. Install go without root. Gobrew doesn't require shell rehash.
- [godbg](https://github.com/sirnewton01/godbg) ![](https://img.shields.io/github/stars/sirnewton01/godbg.svg) - Web-based gdb front-end application.
- [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
- [goma-gateway](https://github.com/jkaninda/goma-gateway) ![](https://img.shields.io/github/stars/jkaninda/goma-gateway.svg) - A Lightweight API Gateway and Reverse Proxy with declarative config, robust middleware, and support for REST, GraphQL, TCP, UDP, and gRPC.
- [gonative](https://github.com/inconshreveable/gonative) ![](https://img.shields.io/github/stars/inconshreveable/gonative.svg) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
- [govvv](https://github.com/ahmetalpbalkan/govvv) ![](https://img.shields.io/github/stars/ahmetalpbalkan/govvv.svg) - “go build” wrapper to easily add version information into Go binaries.
- [grapes](https://github.com/yaronsumel/grapes) ![](https://img.shields.io/github/stars/yaronsumel/grapes.svg) - Lightweight tool designed to distribute commands over ssh with ease.
- [GVM](https://github.com/moovweb/gvm) ![](https://img.shields.io/github/stars/moovweb/gvm.svg) - GVM provides an interface to manage Go versions.
- [Hey](https://github.com/rakyll/hey) ![](https://img.shields.io/github/stars/rakyll/hey.svg) - Hey is a tiny program that sends some load to a web application.
- [httpref](https://github.com/dnnrly/httpref) ![](https://img.shields.io/github/stars/dnnrly/httpref.svg) - httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports.
- [jcli](https://github.com/jenkins-zh/jenkins-cli) ![](https://img.shields.io/github/stars/jenkins-zh/jenkins-cli.svg) - Jenkins CLI allows you manage your Jenkins as an easy way.
- [k0s](https://github.com/k0sproject/k0s) ![](https://img.shields.io/github/stars/k0sproject/k0s.svg) - Zero Friction Kubernetes distribution.
- [k3d](https://github.com/k3d-io/k3d) ![](https://img.shields.io/github/stars/k3d-io/k3d.svg) - Little helper to run CNCF's k3s in Docker.
- [k3s](https://github.com/k3s-io/k3s) ![](https://img.shields.io/github/stars/k3s-io/k3s.svg) - Lightweight Kubernetes.
- [k6](https://github.com/grafana/k6) ![](https://img.shields.io/github/stars/grafana/k6.svg) - A modern load testing tool, using Go and JavaScript.
- [k9s](https://github.com/derailed/k9s) ![](https://img.shields.io/github/stars/derailed/k9s.svg) - Kubernetes CLI to manage your clusters in style.
- [kala](https://github.com/ajvb/kala) ![](https://img.shields.io/github/stars/ajvb/kala.svg) - Simplistic, modern, and performant job scheduler.
- [kcli](https://github.com/cswank/kcli) ![](https://img.shields.io/github/stars/cswank/kcli.svg) - Command line tool for inspecting kafka topics/partitions/messages.
- [kind](https://github.com/kubernetes-sigs/kind) ![](https://img.shields.io/github/stars/kubernetes-sigs/kind.svg) - Kubernetes IN Docker - local clusters for testing Kubernetes.
- [ko](https://github.com/google/ko) ![](https://img.shields.io/github/stars/google/ko.svg) - Command line tool for building and deploying Go applications on Kubernetes
- [kool](https://github.com/kool-dev/kool) ![](https://img.shields.io/github/stars/kool-dev/kool.svg) - Command line tool for managing Docker environments as an easy way.
- [kubeblocks](https://github.com/apecloud/kubeblocks) ![](https://img.shields.io/github/stars/apecloud/kubeblocks.svg) - KubeBlocks is an open-source control plane that runs and manages databases, message queues and other data infrastructure on K8s.
- [kubernetes](https://github.com/kubernetes/kubernetes) ![](https://img.shields.io/github/stars/kubernetes/kubernetes.svg) - Container Cluster Manager from Google.
- [kubeshark](https://github.com/kubeshark/kubeshark) ![](https://img.shields.io/github/stars/kubeshark/kubeshark.svg) - API traffic analyzer for Kubernetes, inspired by Wireshark, purposely built for Kubernetes.
- [KubeVela](https://github.com/kubevela/kubevela) ![](https://img.shields.io/github/stars/kubevela/kubevela.svg) - Cloud native application delivery.
- [KubeVPN](https://github.com/kubenetworks/kubevpn) ![](https://img.shields.io/github/stars/kubenetworks/kubevpn.svg) - KubeVPN offers a Cloud-Native Dev Environment that seamlessly connects to your Kubernetes cluster network.
- [KusionStack](https://github.com/KusionStack/kusion) ![](https://img.shields.io/github/stars/KusionStack/kusion.svg) - A unified programmable configuration techstack to deliver modern app in 'platform as code' and 'infra as code' approach.
- [kwatch](https://github.com/abahmed/kwatch) ![](https://img.shields.io/github/stars/abahmed/kwatch.svg) - Monitor & detect crashes in your Kubernetes(K8s) cluster instantly.
- [lstags](https://github.com/ivanilves/lstags) ![](https://img.shields.io/github/stars/ivanilves/lstags.svg) - Tool and API to sync Docker images across different registries.
- [lwc](https://github.com/timdp/lwc) ![](https://img.shields.io/github/stars/timdp/lwc.svg) - A live-updating version of the UNIX wc command.
- [manssh](https://github.com/xwjdsh/manssh) ![](https://img.shields.io/github/stars/xwjdsh/manssh.svg) - manssh is a command line tool for managing your ssh alias config easily.
- [Mantil](https://github.com/mantil-io/mantil) ![](https://img.shields.io/github/stars/mantil-io/mantil.svg) - Go specific framework for building serverless applications on AWS that enables you to focus on pure Go code while Mantil takes care of the infrastructure.
- [minikube](https://github.com/kubernetes/minikube) ![](https://img.shields.io/github/stars/kubernetes/minikube.svg) - Run Kubernetes locally.
- [Moby](https://github.com/moby/moby) ![](https://img.shields.io/github/stars/moby/moby.svg) - Collaborative project for the container ecosystem to assemble container-based systems.
- [Mora](https://github.com/emicklei/mora) ![](https://img.shields.io/github/stars/emicklei/mora.svg) - REST server for accessing MongoDB documents and meta data.
- [ostent](https://github.com/ostrost/ostent) ![](https://img.shields.io/github/stars/ostrost/ostent.svg) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.
- [Packer](https://github.com/mitchellh/packer) ![](https://img.shields.io/github/stars/mitchellh/packer.svg) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
- [Pewpew](https://github.com/bengadbois/pewpew) ![](https://img.shields.io/github/stars/bengadbois/pewpew.svg) - Flexible HTTP command line stress tester.
- [PipeCD](https://github.com/pipe-cd/pipecd) ![](https://img.shields.io/github/stars/pipe-cd/pipecd.svg) - A GitOps-style continuous delivery platform that provides consistent deployment and operations experience for any applications.
- [podinfo](https://github.com/stefanprodan/podinfo) ![](https://img.shields.io/github/stars/stefanprodan/podinfo.svg) - Podinfo is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes. Podinfo is used by CNCF projects like Flux and Flagger for end-to-end testing and workshops.
- [podman-tui](https://github.com/containers/podman-tui) ![](https://img.shields.io/github/stars/containers/podman-tui.svg) - Terminal UI for Podman management.
- [Pomerium](https://github.com/pomerium/pomerium) ![](https://img.shields.io/github/stars/pomerium/pomerium.svg) - Pomerium is an identity-aware access proxy.
- [Rodent](https://github.com/alouche/rodent) ![](https://img.shields.io/github/stars/alouche/rodent.svg) - Rodent helps you manage Go versions, projects and track dependencies.
- [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) ![](https://img.shields.io/github/stars/oxyno-zeta/s3-proxy.svg) - S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r) ![](https://img.shields.io/github/stars/rlmcpherson/s3gof3r.svg) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.
- [s5cmd](https://github.com/peak/s5cmd) ![](https://img.shields.io/github/stars/peak/s5cmd.svg) - Blazing fast S3 and local filesystem execution tool.
- [Scaleway-cli](https://github.com/scaleway/scaleway-cli) ![](https://img.shields.io/github/stars/scaleway/scaleway-cli.svg) - Manage BareMetal Servers from Command Line (as easily as with Docker).
- [script](https://github.com/bitfield/script) ![](https://img.shields.io/github/stars/bitfield/script.svg) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.
- [sg](https://github.com/ChristopherRabotin/sg) ![](https://img.shields.io/github/stars/ChristopherRabotin/sg.svg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.
- [sigma](https://github.com/go-sigma/sigma) ![](https://img.shields.io/github/stars/go-sigma/sigma.svg) - OCI-native container image registry, support OCI-native artifact, scan artifact, image build etc.
- [skm](https://github.com/TimothyYe/skm) ![](https://img.shields.io/github/stars/TimothyYe/skm.svg) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!
- [StatusOK](https://github.com/sanathp/statusok) ![](https://img.shields.io/github/stars/sanathp/statusok.svg) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.
- [tau](https://github.com/taubyte/tau) ![](https://img.shields.io/github/stars/taubyte/tau.svg) - Easily build Cloud Computing Platforms with features like Serverless WebAssembly Functions, Frontend Hosting, CI/CD, Object Storage, K/V Database, and Pub-Sub Messaging.
- [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) ![](https://img.shields.io/github/stars/dikhan/terraform-provider-openapi.svg) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.
- [tf-profile](https://github.com/datarootsio/tf-profile) ![](https://img.shields.io/github/stars/datarootsio/tf-profile.svg) - Profiler for Terraform runs. Generate global stats, resource-level stats or visualizations.
- [tlm](https://github.com/yusufcanb/tlm) ![](https://img.shields.io/github/stars/yusufcanb/tlm.svg) - Local cli copilot, powered by CodeLLaMa
- [traefik](https://github.com/containous/traefik) ![](https://img.shields.io/github/stars/containous/traefik.svg) - Reverse proxy and load balancer with support for multiple backends.
- [trubka](https://github.com/xitonix/trubka) ![](https://img.shields.io/github/stars/xitonix/trubka.svg) - A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka.
- [uTask](https://github.com/ovh/utask) ![](https://img.shields.io/github/stars/ovh/utask.svg) - Automation engine that models and executes business processes declared in yaml.
- [Vegeta](https://github.com/tsenart/vegeta) ![](https://img.shields.io/github/stars/tsenart/vegeta.svg) - HTTP load testing tool and library. It's over 9000!
- [wait-for](https://github.com/dnnrly/wait-for) ![](https://img.shields.io/github/stars/dnnrly/wait-for.svg) - Wait for something to happen (from the command line) before continuing. Easy orchestration of Docker services and other things.
- [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
- [winrm-cli](https://github.com/masterzen/winrm-cli) ![](https://img.shields.io/github/stars/masterzen/winrm-cli.svg) - Cli tool to remotely execute commands on Windows machines.

**[⬆ back to top](#contents)**

### Other Software

- [Better Go Playground](https://goplay.tools) - Go playground with syntax highlight, code completion and other features.
- [blocky](https://github.com/0xERR0R/blocky) ![](https://img.shields.io/github/stars/0xERR0R/blocky.svg) - Fast and lightweight DNS proxy as ad-blocker for local network with many features.
- [bluetuith](https://github.com/bluetuith-org/bluetuith) ![](https://img.shields.io/github/stars/bluetuith-org/bluetuith.svg) - TUI Bluetooth manager for Linux.
- [borg](https://github.com/crufter/borg) ![](https://img.shields.io/github/stars/crufter/borg.svg) - Terminal based search engine for bash snippets.
- [boxed](https://github.com/tejo/boxed) ![](https://img.shields.io/github/stars/tejo/boxed.svg) - Dropbox based blog engine.
- [Chapar](https://github.com/chapar-rest/chapar) ![](https://img.shields.io/github/stars/chapar-rest/chapar.svg) - Chapar is a a cross-platform Postman alternative built with go, aims to help developers to test their api endpoints. it support http and grpc protocols.
- [Cherry](https://github.com/rafael-santiago/cherry) ![](https://img.shields.io/github/stars/rafael-santiago/cherry.svg) - Tiny webchat server in Go.
- [Circuit](https://github.com/gocircuit/circuit) ![](https://img.shields.io/github/stars/gocircuit/circuit.svg) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
- [Comcast](https://github.com/tylertreat/Comcast) ![](https://img.shields.io/github/stars/tylertreat/Comcast.svg) - Simulate bad network connections.
- [confd](https://github.com/kelseyhightower/confd) ![](https://img.shields.io/github/stars/kelseyhightower/confd.svg) - Manage local application configuration files using templates and data from etcd or consul.
- [crawley](https://github.com/s0rg/crawley) ![](https://img.shields.io/github/stars/s0rg/crawley.svg) - Web scraper/crawler for cli.
- [croc](https://github.com/schollz/croc) ![](https://img.shields.io/github/stars/schollz/croc.svg) - Easily and securely send files or folders from one computer to another.
- [Documize](https://github.com/documize/community) ![](https://img.shields.io/github/stars/documize/community.svg) - Modern wiki software that integrates data from SaaS tools.
- [dp](https://github.com/scryinfo/dp) ![](https://img.shields.io/github/stars/scryinfo/dp.svg) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.
- [drive](https://github.com/odeke-em/drive) ![](https://img.shields.io/github/stars/odeke-em/drive.svg) - Google Drive client for the commandline.
- [Duplicacy](https://github.com/gilbertchen/duplicacy) ![](https://img.shields.io/github/stars/gilbertchen/duplicacy.svg) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.
- [fjira](https://github.com/mk-5/fjira) ![](https://img.shields.io/github/stars/mk-5/fjira.svg) - A fuzzy-search based terminal UI application for Attlasian Jira
- [Gebug](https://github.com/moshebe/gebug) ![](https://img.shields.io/github/stars/moshebe/gebug.svg) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
- [gfile](https://github.com/Antonito/gfile) ![](https://img.shields.io/github/stars/Antonito/gfile.svg) - Securely transfer files between two computers, without any third party, over WebRTC.
- [Go Package Store](https://github.com/shurcooL/Go-Package-Store) ![](https://img.shields.io/github/stars/shurcooL/Go-Package-Store.svg) - App that displays updates for the Go packages in your GOPATH.
- [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) ![](https://img.shields.io/github/stars/Sioro-Neoku/go-peerflix.svg) - Video streaming torrent client.
- [goblin](https://goblin.run) - Cloud builder for CLI's written in go lang
- [GoBoy](https://github.com/Humpheh/goboy) ![](https://img.shields.io/github/stars/Humpheh/goboy.svg) - Nintendo Game Boy Color emulator written in Go.
- [gocc](https://github.com/goccmack/gocc) ![](https://img.shields.io/github/stars/goccmack/gocc.svg) - Gocc is a compiler kit for Go written in Go.
- [GoDocTooltip](https://github.com/diankong/GoDocTooltip) ![](https://img.shields.io/github/stars/diankong/GoDocTooltip.svg) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list.
- [Gokapi](https://github.com/Forceu/gokapi) ![](https://img.shields.io/github/stars/Forceu/gokapi.svg) - Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload.
- [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
- [GoNB](https://github.com/janpfeifer/gonb) ![](https://img.shields.io/github/stars/janpfeifer/gonb.svg) - Interactive Go programming with Jupyter Notebooks (also works in VSCode, Binder and Google's Colab).
- [Gor](https://github.com/buger/gor) ![](https://img.shields.io/github/stars/buger/gor.svg) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
- [Guora](https://github.com/meloalright/guora) ![](https://img.shields.io/github/stars/meloalright/guora.svg) - A self-hosted Quora like web application written in Go.
- [hoofli](https://github.com/dnnrly/hoofli) ![](https://img.shields.io/github/stars/dnnrly/hoofli.svg) - Generate PlantUML diagrams from Chrome or Firefox network inspections.
- [hotswap](https://github.com/edwingeng/hotswap) ![](https://img.shields.io/github/stars/edwingeng/hotswap.svg) - A complete solution to reload your go code without restarting your server, interrupting or blocking any ongoing procedure.
- [hugo](https://gohugo.io/) - Fast and Modern Static Website Engine.
- [ide](https://github.com/thestrukture/ide) ![](https://img.shields.io/github/stars/thestrukture/ide.svg) - Browser accessible IDE. Designed for Go with Go.
- [joincap](https://github.com/assafmo/joincap) ![](https://img.shields.io/github/stars/assafmo/joincap.svg) - Command-line utility for merging multiple pcap files together.
- [JuiceFS](https://github.com/juicedata/juicefs) ![](https://img.shields.io/github/stars/juicedata/juicefs.svg) - Distributed POSIX file system built on top of Redis and AWS S3.
- [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
- [Layli](https://layli.app) - Draw pretty layout diagrams as code.
- [Leaps](https://github.com/jeffail/leaps) ![](https://img.shields.io/github/stars/jeffail/leaps.svg) - Pair programming service using Operational Transforms.
- [lgo](https://github.com/yunabe/lgo) ![](https://img.shields.io/github/stars/yunabe/lgo.svg) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.
- [limetext](https://limetext.github.io) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
- [LiteIDE](https://github.com/visualfc/liteide) ![](https://img.shields.io/github/stars/visualfc/liteide.svg) - LiteIDE is a simple, open source, cross-platform Go IDE.
- [mockingjay](https://github.com/quii/mockingjay-server) ![](https://img.shields.io/github/stars/quii/mockingjay-server.svg) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.
- [myLG](https://github.com/mehrdadrad/mylg) ![](https://img.shields.io/github/stars/mehrdadrad/mylg.svg) - Command Line Network Diagnostic tool written in Go.
- [naclpipe](https://github.com/unix4fun/naclpipe) ![](https://img.shields.io/github/stars/unix4fun/naclpipe.svg) - Simple NaCL EC25519 based crypto pipe tool written in Go.
- [Neo-cowsay](https://github.com/Code-Hex/Neo-cowsay) ![](https://img.shields.io/github/stars/Code-Hex/Neo-cowsay.svg) - 🐮 cowsay is reborn. for a New Era.
- [nes](https://github.com/fogleman/nes) ![](https://img.shields.io/github/stars/fogleman/nes.svg) - Nintendo Entertainment System (NES) emulator written in Go.
- [Orbit](https://github.com/gulien/orbit) ![](https://img.shields.io/github/stars/gulien/orbit.svg) - A simple tool for running commands and generating files from templates.
- [peg](https://github.com/pointlander/peg) ![](https://img.shields.io/github/stars/pointlander/peg.svg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
- [Plik](https://github.com/root-gg/plik) ![](https://img.shields.io/github/stars/root-gg/plik.svg) - Plik is a temporary file upload system (Wetransfer like) in Go.
- [portal](https://github.com/SpatiumPortae/portal) ![](https://img.shields.io/github/stars/SpatiumPortae/portal.svg) - Portal is a quick and easy command-line file transfer utility from any computer to another.
- [restic](https://github.com/restic/restic) ![](https://img.shields.io/github/stars/restic/restic.svg) - De-duplicating backup program.
- [sake](https://github.com/alajmo/sake) ![](https://img.shields.io/github/stars/alajmo/sake.svg) - sake is a command runner for local and remote hosts.
- [scc](https://github.com/boyter/scc) ![](https://img.shields.io/github/stars/boyter/scc.svg) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.
- [Seaweed File System](https://github.com/chrislusf/seaweedfs) ![](https://img.shields.io/github/stars/chrislusf/seaweedfs.svg) - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
- [shell2http](https://github.com/msoap/shell2http) ![](https://img.shields.io/github/stars/msoap/shell2http.svg) - Executing shell commands via http server (for prototyping or remote control).
- [Snitch](https://github.com/lucasgomide/snitch) ![](https://img.shields.io/github/stars/lucasgomide/snitch.svg) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru.
- [sonic](https://github.com/go-sonic/sonic) ![](https://img.shields.io/github/stars/go-sonic/sonic.svg) - Sonic is a Go Blogging Platform. Simple and Powerful.
- [Stack Up](https://github.com/pressly/sup) ![](https://img.shields.io/github/stars/pressly/sup.svg) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
- [stew](https://github.com/marwanhawari/stew) ![](https://img.shields.io/github/stars/marwanhawari/stew.svg) - An independent package manager for compiled binaries.
- [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
- [tcpdog](https://github.com/mehrdadrad/tcpdog) ![](https://img.shields.io/github/stars/mehrdadrad/tcpdog.svg) - eBPF based TCP observability.
- [tinycare-tui](https://github.com/DMcP89/tinycare-tui) ![](https://img.shields.io/github/stars/DMcP89/tinycare-tui.svg) - Small terminal app that shows git commits from the last 24 hours and week, current weather, some self care advice, a joke, and you current todo list tasks.
- [toxiproxy](https://github.com/shopify/toxiproxy) ![](https://img.shields.io/github/stars/shopify/toxiproxy.svg) - Proxy to simulate network and system conditions for automated tests.
- [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
- [vaku](https://github.com/lingrino/vaku) ![](https://img.shields.io/github/stars/lingrino/vaku.svg) - CLI & API for folder-based functions in Vault like copy, move, and search.
- [vFlow](https://github.com/VerizonDigital/vflow) ![](https://img.shields.io/github/stars/VerizonDigital/vflow.svg) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.
- [Wave Terminal](https://waveterm.dev) - Wave is an open-source, AI-native terminal built for seamless developer workflows with inline rendering, a modern UI, and persistent sessions.
- [wellington](https://github.com/wellington/wellington) ![](https://img.shields.io/github/stars/wellington/wellington.svg) - Sass project management tool, extends the language with sprite functions (like Compass).
- [woke](https://github.com/get-woke/woke) ![](https://img.shields.io/github/stars/get-woke/woke.svg) - Detect non-inclusive language in your source code.
- [yai](https://github.com/ekkinox/yai) ![](https://img.shields.io/github/stars/ekkinox/yai.svg) - AI powered terminal assistant.
- [zs](https://git.mills.io/prologic/zs) - an extremely minimal static site generator.

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [autobench](https://github.com/davecheney/autobench) ![](https://img.shields.io/github/stars/davecheney/autobench.svg) - Framework to compare the performance between different Go versions.
- [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) ![](https://img.shields.io/github/stars/mrLSD/go-benchmark-app.svg) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.
- [go-benchmarks](https://github.com/tylertreat/go-benchmarks) ![](https://img.shields.io/github/stars/tylertreat/go-benchmarks.svg) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.
- [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) ![](https://img.shields.io/github/stars/julienschmidt/go-http-routing-benchmark.svg) - Go HTTP request router benchmark and comparison.
- [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) ![](https://img.shields.io/github/stars/zerosnake0/go-json-benchmark.svg) - Go JSON benchmark.
- [go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks) ![](https://img.shields.io/github/stars/nikolaydubina/go-ml-benchmarks.svg) - benchmarks for machine learning inference in Go.
- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) ![](https://img.shields.io/github/stars/smallnest/go-web-framework-benchmark.svg) - Go web framework benchmark.
- [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) ![](https://img.shields.io/github/stars/alecthomas/go_serialization_benchmarks.svg) - Benchmarks of Go serialization methods.
- [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) ![](https://img.shields.io/github/stars/PuerkitoBio/gocostmodel.svg) - Benchmarks of common basic operations for the Go language.
- [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) ![](https://img.shields.io/github/stars/SimonWaldherr/golang-benchmarks.svg) - a collection of golang benchmarks.
- [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) ![](https://img.shields.io/github/stars/tyler-smith/golang-sql-benchmark.svg) - Collection of benchmarks for popular Go database/SQL utilities.
- [gospeed](https://github.com/feyeleanor/GoSpeed) ![](https://img.shields.io/github/stars/feyeleanor/GoSpeed.svg) - Go micro-benchmarks for calculating the speed of language constructs.
- [kvbench](https://github.com/jimrobinson/kvbench) ![](https://img.shields.io/github/stars/jimrobinson/kvbench.svg) - Key/Value database benchmark.
- [skynet](https://github.com/atemerev/skynet) ![](https://img.shields.io/github/stars/atemerev/skynet.svg) - Skynet 1M threads microbenchmark.
- [speedtest-resize](https://github.com/fawick/speedtest-resize) ![](https://img.shields.io/github/stars/fawick/speedtest-resize.svg) - Compare various Image resize algorithms for the Go language.

**[⬆ back to top](#contents)**

## Conferences

- [GoCon](https://gocon.connpass.com/) - Tokyo, Japan.
- [GoDays](https://www.godays.io/) - Berlin, Germany.
- [GoLab](https://golab.io/) - Florence, Italy.
- [GopherChina](https://gopherchina.org) - Shanghai, China.
- [GopherCon](https://www.gophercon.com/) - Varied Locations Each Year, USA.
- [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
- [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, Brazil.
- [GopherCon Europe](https://gophercon.eu/) - Berlin, Germany.
- [GopherCon India](https://gopherconindia.org/) - Pune, India.
- [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel.
- [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia.
- [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore.
- [GopherCon UK](https://www.gophercon.co.uk/) - London, UK.
- [GopherCon Vietnam](https://gophercon.vn/) - Ho Chi Minh City, Vietnam.
- [GoWest Conference](https://www.gowestconf.com/) - Lehi, USA.

**[⬆ back to top](#contents)**

## E-Books

### E-books for purchase

- [100 Go Mistakes: How to Avoid Them](https://www.manning.com/books/100-go-mistakes-how-to-avoid-them)
- [Black Hat Go](https://nostarch.com/blackhatgo) - Go programming for hackers and pentesters.
- [Build an Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)
- [Continuous Delivery in Go](https://www.manning.com/books/continuous-delivery-in-go) - This practical guide to continuous delivery shows you how to rapidly establish an automated pipeline that will improve your testing, code quality, and final product.
- [Creative DIY Microcontroller Project With TinyGo and WebAssembly](https://www.packtpub.com/product/creative-diy-microcontroller-projects-with-tinygo-and-webassembly/9781800560208) - An introduction into the TinyGo compiler with projects involving Arduino and WebAssembly.
- [Effective Go: Elegant, efficient, and testable code](https://www.manning.com/books/effective-go) - Unlock Go’s unique perspective on program design, and start writing simple, maintainable, and testable Go code.
- [For the Love of Go](https://bitfieldconsulting.com/books/love) - An introductory book for Go beginners.
- [Go in Practice, Second Edition](https://www.manning.com/books/go-in-practice-second-edition) - Your practical guide on the ins-and-outs of Go development, covering the standard library and the most important tools from Go’s powerful ecosystem.
- [Know Go: Generics](https://bitfieldconsulting.com/books/generics) - A guide to understanding and using generics in Go.
- [Lets-Go](https://lets-go.alexedwards.net) - A step-by-step guide to creating fast, secure and maintanable web applications with Go.
- [Lets-Go-Further](https://lets-go-further.alexedwards.net) - Advanced patterns for building APIs and web applications in Go.
- [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests) - A guide to testing in Go.
- [The Power of Go: Tools](https://bitfieldconsulting.com/books/tools) - A guide to writing command-line tools in Go.
- [Writing A Compiler In Go](https://compilerbook.com)
- [Writing An Interpreter In Go](https://interpreterbook.com) - Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.

### Free e-books

- [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
- [An Introduction to Programming in Go](http://www.golang-book.com/)
- [Build a blockchain from scratch in Go with gRPC](https://github.com/volodymyrprokopyuk/go-blockchain) ![](https://img.shields.io/github/stars/volodymyrprokopyuk/go-blockchain.svg) - The foundational and practical guide for effectively learning and progressively building a blockchain from scratch in Go with gRPC.
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)
- [Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
- [Go AST Book (Chinese)](https://github.com/chai2010/go-ast-book) ![](https://img.shields.io/github/stars/chai2010/go-ast-book.svg) - A book focusing on Go `go/*` packages.
- [Go Faster](https://leanpub.com/gofaster) - This book seeks to shorten your learning curve and help you become a proficient Go programmer, faster.
- [Go Succinctly](https://github.com/thedevsir/gosuccinctly) ![](https://img.shields.io/github/stars/thedevsir/gosuccinctly.svg) - in Persian.
- [Go with the domain](https://threedots.tech/go-with-the-domain/) - A book showing how to apply DDD, Clean Architecture, and CQRS by practical refactoring.
- [GoBooks](https://github.com/dariubs/GoBooks) ![](https://img.shields.io/github/stars/dariubs/GoBooks.svg) - A curated list of Go books.
- [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) - A 600 page introduction to Go aimed at first time developers.
- [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
- [Network Programming With Go](https://jan.newmarch.name/golang/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
- [The Go Programming Language](https://www.gopl.io/)
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) ![](https://img.shields.io/github/stars/polaris1119/The-Golang-Standard-Library-by-Example.svg)
- [The Little Go Book](https://github.com/karlseguin/the-little-go-book) ![](https://img.shields.io/github/stars/karlseguin/the-little-go-book.svg)
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/) ![](https://img.shields.io/github/stars/thewhitetulip/web-dev-golang-anti-textbook.svg)

**[⬆ back to top](#contents)**

## Gophers

- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) ![](https://img.shields.io/github/stars/MariaLetta/free-gophers-pack.svg) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.
- [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) ![](https://img.shields.io/github/stars/keygx/Go-gopher-Vector.svg) - Go gopher Vector Data [.ai, .svg].
- [gopher-logos](https://github.com/GolangUA/gopher-logos) ![](https://img.shields.io/github/stars/GolangUA/gopher-logos.svg) - adorable gopher logos.
- [gopher-stickers](https://github.com/tenntenn/gopher-stickers) ![](https://img.shields.io/github/stars/tenntenn/gopher-stickers.svg)
- [gophericons](https://github.com/shalakhin/gophericons) ![](https://img.shields.io/github/stars/shalakhin/gophericons.svg)
- [gopherize.me](https://github.com/matryer/gopherize.me) ![](https://img.shields.io/github/stars/matryer/gopherize.me.svg) - Gopherize yourself.
- [gophers](https://github.com/ashleymcnamara/gophers) ![](https://img.shields.io/github/stars/ashleymcnamara/gophers.svg) - Gopher artworks by Ashley McNamara.
- [gophers](https://github.com/egonelbre/gophers) ![](https://img.shields.io/github/stars/egonelbre/gophers.svg) - Free gophers.
- [gophers](https://github.com/rogeralsing/gophers) ![](https://img.shields.io/github/stars/rogeralsing/gophers.svg) - random gopher graphics.
- [gophers](https://github.com/sillecelik/go-gopher) ![](https://img.shields.io/github/stars/sillecelik/go-gopher.svg) - Gopher amigurumi toy pattern.
- [gophers](https://github.com/scraly/gophers) ![](https://img.shields.io/github/stars/scraly/gophers.svg) - Gophers by Aurélie Vache.

**[⬆ back to top](#contents)**

## Meetups

- [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
- [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
- [Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)
- [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
- [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
- [Bärner Go Meetup - Berne, Switzerland](https://www.meetup.com/berner-go-meetup/)
- [Go Ireland - Dublin](https://www.meetup.com/goireland/)
- [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
- [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
- [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
- [Go Toronto](https://www.meetup.com/go-toronto/)
- [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
- [GoBandung](https://www.meetup.com/GoBandung/)
- [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
- [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
- [GoJakarta](https://www.meetup.com/GoJakarta/)
- [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
- [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
- [Golang Athens](https://www.meetup.com/Athens-Gophers/)
- [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
- [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
- [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
- [Golang Boston](https://www.meetup.com/bostongo/)
- [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
- [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
- [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
- [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
- [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
- [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
- [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
- [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
- [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
- [Golang Israel](https://www.meetup.com/Go-Israel/)
- [Golang Kathmandu](https://www.meetup.com/Golang-Kathmandu/)
- [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
- [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
- [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
- [Golang Melbourne](https://www.meetup.com/golang-mel/)
- [Golang Milano](https://www.meetup.com/golang-milano/)
- [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
- [Golang Paris](https://www.meetup.com/Golang-Paris/)
- [Golang Poland](https://www.meetup.com/Golang-Poland/)
- [Golang Pune](https://www.meetup.com/Golang-Pune/)
- [Golang Roma](https://www.meetup.com/golangroma/)
- [Golang Rotterdam](https://www.meetup.com/golang-rotterdam/)
- [Golang Singapore](https://www.meetup.com/golangsg/)
- [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
- [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
- [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
- [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
- [Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)
- [Golang Torino](https://www.meetup.com/golang-torino/)
- [Golang Turkey](https://kommunity.com/goturkiye)
- [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
- [Golang Vienna, Austria](https://www.meetup.com/viennago/)
- [Golang Москва](https://www.meetup.com/Golang-Moscow/)
- [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
- [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
- [Lagos Gophers](https://www.meetup.com/GolangNigeria/)
- [Nairobi Gophers](https://www.meetup.com/nairobi-gophers/)
- [Seattle Go Programmers](https://www.meetup.com/golang/)
- [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
- [Utah Go User Group](https://www.meetup.com/utahgophers/)
- [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)
- [Zürich Gophers - Zurich, Switzerland](https://www.meetup.com/zurich-gophers/)

_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide) ![](https://img.shields.io/github/stars/bahlo/go-styleguide.svg)
- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md) ![](https://img.shields.io/github/stars/cockroachdb/cockroach.svg)
- [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
- [Google](https://google.github.io/styleguide/go/)
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst) ![](https://img.shields.io/github/stars/hyperledger/fabric.svg)
- [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)
- [Trybe](https://github.com/betrybe/playbook-go/blob/main/README_EN.md) ![](https://img.shields.io/github/stars/betrybe/playbook-go.svg)
- [Uber](https://github.com/uber-go/guide/blob/master/style.md) ![](https://img.shields.io/github/stars/uber-go/guide.svg)

**[⬆ back to top](#contents)**

## Social Media

### Twitter

- [@GoDiscussions](https://twitter.com/GoDiscussions)
- [@golang](https://twitter.com/golang)
- [@golang_news](https://twitter.com/golang_news)
- [@golangch](https://twitter.com/golangch)
- [@golangweekly](https://twitter.com/golangweekly)

**[⬆ back to top](#contents)**

### Reddit

- [r/golang](https://www.reddit.com/r/golang/)

**[⬆ back to top](#contents)**

## Websites

- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) ![](https://img.shields.io/github/stars/amit-davidson/awesome-golang-workshops.svg) - A curated list of awesome golang workshops.
- [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) ![](https://img.shields.io/github/stars/lukasz-madon/awesome-remote-job.svg) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
- [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) ![](https://img.shields.io/github/stars/bayandin/awesome-awesomeness.svg) - List of other amazingly awesome lists.
- [awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra) ![](https://img.shields.io/github/stars/xwjdsh/awesome-go-extra.svg) - Parse awesome-go README file and generate a new README file with repo info.
- [Code with Mukesh](https://codewithmukesh.com/categories/golang) - Software Engineer and Blogs @ codewithmukesh.com.
- [Coding Mystery](https://codingmystery.com) - Solve exciting escape-room-inspired programming challenges using Go.
- [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
- [Go Blog](https://blog.golang.org) - The official Go blog.
- [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3) - A group of Gophers read and discuss a different Go project every week.
- [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
- [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
- [Go Projects](https://github.com/golang/go/wiki/Projects) ![](https://img.shields.io/github/stars/golang/go.svg) - List of projects on the Go community wiki.
- [Go Proverbs](https://go-proverbs.github.io/) - Go Proverbs by Rob Pike.
- [Go Report Card](https://goreportcard.com) - A report card for your Go package.
- [go.dev](https://go.dev/) - A hub for Go developers.
- [gocryforhelp](https://github.com/ninedraft/gocryforhelp) ![](https://img.shields.io/github/stars/ninedraft/gocryforhelp.svg) - Collection of Go projects that needs help. Good place to start your open-source way in Go.
- [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusively for Golang related Roles.
- [Golang News](https://golangnews.com) - Links and news about Go programming.
- [Golang Nugget](https://golangnugget.com) - A weekly roundup of the best Go content, delivered to your inbox every Monday.
- [Golang Weekly](https://discu.eu/weekly/golang/) - Each monday projects, tutorials and articles about Go.
- [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
- [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
- [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
- [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
- [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
- [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by Francesc Campoy [@francesc](https://twitter.com/francesc).
- [Learn Go Programming](https://blog.learngoprogramming.com) - Learn Go concepts with illustrations.
- [Libs.tech](https://libs.tech/go) – Awesome Go libraries and hidden gems
- [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
- [pkg.go.dev](https://pkg.go.dev/) - Documentation for open source Go packages.
- [studygolang](https://studygolang.com) - The community of studygolang in China.
- [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
- [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

**[⬆ back to top](#contents)**

### Tutorials

- [50 Shades of Go](https://golang50shades.github.io/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
- [A Comprehensive Guide to Structured Logging in Go](https://betterstack.com/community/guides/logging/logging-in-go/) - Delve deep into the world of structured logging in Go with a specific focus on recently accepted slog proposal which aims to bring high performance structured logging with levels to the standard library.
- [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - Building a Golang site for e-commerce (demo included).
- [A Tour of Go](https://tour.golang.org/) - Interactive tour of Go.
- [Build a Database in 1000 lines of code](https://link.medium.com/O9YQlx89Htb) - Build a NoSQL Database From Zero in 1000 Lines of Code.
- [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) ![](https://img.shields.io/github/stars/astaxie/build-web-application-with-golang.svg) - Golang ebook intro how to build a web app with golang.
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) - We’ll write an API with the help of the powerful Gorilla Mux.
- [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
- [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
- [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
- [CodeCrafters Golang Track](https://app.codecrafters.io/tracks/go) - Achieve mastery in advanced Go by building your own Redis, Docker, Git, and SQLite. Featuring goroutines, systems programming, file I/O, and more.
- [Design Patterns in Go](https://github.com/shubhamzanwar/design-patterns) ![](https://img.shields.io/github/stars/shubhamzanwar/design-patterns.svg) - Collection of programming design patterns implemented in Go.
- [Games With Go](https://www.youtube.com/watch?v=9D4yH7e_ea8&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x) - A video series teaching programming and game development.
- [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
- [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) ![](https://img.shields.io/github/stars/a8m/go-lang-cheat-sheet.svg) - Go's reference card.
- [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
- [Go in 7 days](https://github.com/harrytran103/7_days_of_go) ![](https://img.shields.io/github/stars/harrytran103/7_days_of_go.svg) - Learn everything about Go in 7 days (from a Nodejs developer).
- [Go Language Tutorial](https://www.javatpoint.com/go-tutorial) - Learn Go language Tutorial.
- [Go Tutorial](https://www.tutorialspoint.com/go/index.htm) - Learn Go programming.
- [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [go-clean-template](https://github.com/evrone/go-clean-template) ![](https://img.shields.io/github/stars/evrone/go-clean-template.svg) - Clean Architecture template for Golang services.
- [go-patterns](https://github.com/tmrts/go-patterns) ![](https://img.shields.io/github/stars/tmrts/go-patterns.svg) - Curated list of Go design patterns, recipes and idioms.
- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) ![](https://img.shields.io/github/stars/miguelmota/golang-for-nodejs-developers.svg) - Examples of Golang compared to Node.js for learning.
- [Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/) - A List of Free Courses to Learn the Go Programming Language.
- [golang-examples](https://github.com/SimonWaldherr/golang-examples) ![](https://img.shields.io/github/stars/SimonWaldherr/golang-examples.svg) - Many examples to learn Golang.
- [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
- [GopherCoding](https://gophercoding.com/) - Collection of code snippets and tutorials to help tackle every day issues.
- [GopherSnippets](https://gophersnippets.com/) - Code snippets with tests and testable examples for the Go programming language.
- [Gosamples](https://gosamples.dev/) - Collection of code snippets that let you solve everyday code problems.
- [GraphQL with Go](https://hasura.io/learn/graphql/backend-stack/languages/go/) - Learn how to create a Go GraphQL server and client with code generation. Also includes creating REST endpoints.
- [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
- [Hex Monscape](https://github.com/Haraj-backend/hex-monscape) ![](https://img.shields.io/github/stars/Haraj-backend/hex-monscape.svg) - Getting started guidelines in writing maintainable code using Hexagonal Architecture.
- [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5) - Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
- [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Learn how to use Docker for Go development and how to build production Docker images.
- [How to Implement Role-Based Access Control (RBAC) Authorization in Golang](https://www.permit.io/blog/role-based-access-control-rbac-authorization-in-golang) - A guide to implementing Role-Based Access Control (RBAC) in Golang, including code examples, covering various methods to secure app endpoints with role-based authorization.
- [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog - a Behavior-driven development framework for building and testing Go applications.
- [Learn Go with 1000+ Exercises](https://github.com/inancgumus/learngo) ![](https://img.shields.io/github/stars/inancgumus/learngo.svg) - Learn Go with thousands of examples, exercises, and quizzes.
- [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) ![](https://img.shields.io/github/stars/quii/learn-go-with-tests.svg) - Learn Go with test-driven development.
- [Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n) - Series of articles in order to learn Golang language by concrete applications as example.
- [Microservices with Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) - Dive deep into building microservices using Go, including gRPC.
- [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
- [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
- [Scaling Go Applications](https://betterstack.com/community/guides/scaling-go/) - Everything about building, deploying and scaling Go applications in production.
- [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
- [Understanding Go in a visual way](https://dev.to/aurelievache/series/26234) - Learn Go visually
- [W3basic Go Tutorials](https://www.w3basic.com/golang/) - W3Basic provides an in-depth tutorial and well-organized content to learn Golang programming.
- [Your basic Go](https://yourbasic.org/golang) - Huge collection of tutorials and how to's.

**[⬆ back to top](#contents)**

### Guided Learning

- [The Go Developer Roadmap](https://roadmap.sh/golang) - A visual roadmap that new Go developers can follow through to help them learn Go.
- [The Go Interview Practice](https://github.com/RezaSi/go-interview-practice) ![](https://img.shields.io/github/stars/RezaSi/go-interview-practice.svg) - A GitHub repository offering coding challenges for Go technical interview preparation.
- [The Go Learning Path](https://tutorialedge.net/paths/golang/) - A guided learning path containing a mix of free and premium resources.
- [The Go Skill Tree](https://labex.io/skilltrees/go) - A structured learning path that combines both free and premium resources.

**[⬆ back to top](#contents)**

## Contribution

We welcome contributions! Please refer to our [CONTRIBUTING.md](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the [MIT License](https://github.com/avelino/awesome-go/blob/main/LICENSE) - see the LICENSE file for details.
