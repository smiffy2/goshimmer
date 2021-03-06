# v0.2.2 - 2020-07-27
* Improves message and transaction validation: 
    * Adjust max transaction inputs count;
    * Adds signature validation before issuance; 
    * Enforce max message size in message factory.
* Improves API:
    * Changes granularity of spammer API to accept messages per minute;
    * Adds API middleware and set CORS to allow from every origin;
    * Adds sendTransactionByJSON to allow a client to issue transactions by providing them in a JSON format;
    * Adds tool API endpoint to facilitate debugging of the solidification status;
    * Removes old API documentation;
* Improves synchronization process:
    * Refactors message requester to be more reliable;
    * Improves solidification process;
    * Refactors worker pool management;
    * Replaces bootstrap plugin with the more secure and reliable beacon plugin.
* Improves analysis-server dashboard:
    * Adds the ability to distinguish between different GoShimmer node versions;
    * Refactors the interaction between server side and dashboard;
    * Improves consensus visualization;
    * Improves dashboard tooling.
* Adds a new electron-based wallet.
* Increases max gossip packet size.
* Adds command to the CLI to override database dirty flag.
* Grafana local dashboard
    * Adds messages in database chart (solid, not solid, total)
    * Adds average solidification time chart
    * Adds Message Request Queue Size chart
* **Breaking**: bumps network and database versions

# v0.2.1 - 2020-07-01
* Adds PoW requirement to faucet payloads
* Adds tips broadcaster to ensure that all chains are getting solidified
* Fixes being able to send a double-spend via one node
* **Breaking**: bumps network and database versions

# v0.2.0 - 2020-06-30
* Adds the value transfer dApp:
    * New binary transaction layout
    * UTXO model 
    * Support for transactions with Ed25519 and BLS signatures
    * Parallel reality based ledger state
    * Support for colored tokens
    * Conflict resolution via FPC
    * Applied FCoB rules
* Adds the network delay dApp which is used to gather the avg. network delay occurring in the network
* Adds the faucet dApp giving the ability to request funds for testing via the dashboard or web API
* Adds the DRNG dApp which is used to propagate random numbers produced by a dRand committee (this dApp is inactive)
* Adds the base communication layer
* Adds improved analysis server:
    * Splits it into 3 separate plugins `analysis-client`/`analysis-dashboard`/`analysis-server`
    * Applies heartbeat pattern for robustness on both client and server side
    * Uses TLV denoted messages for communication
    * Complete new dashboard with live visualisations of the network graph and ongoing conflicts
    * Use short node IDs throughout the analysis dashboard
    * Prometheus exporter to expose global network metrics
    * Storage for conflicts inside a MongoDB for further processing
    * Complete rewritten autopeering data retention
* Adds additional HTTP API routes:
    * "healtz" route to query the health of the node (for LBs)
    * Query transaction attachment locations
    * Query transactions by IDs
    * Send transactions
    * Get UTXOs on addresses
    * Query info about the node
    * Issue a faucet funding request
    * Query dRNG committee
    * Query dRNG random number
    * Issue data payloads
* Adds dashboard improvements:
    * Tips chart
    * Communication layer visualizer
    * Address and UTXOs view
    * Message payload view
    * DRNG live feed
    * Faucet page to request funds
    * Support different payload views (data, value, faucet, unknown)
* Adds integration test framework and corresponding tests:
    * Autopeering/Network Split
    * Message propagation 
    * FPC 50/50 network split voting
    * Faucet funding
    * Value transfers
    * Synchronization
* Adds refactored communication layer code
* Adds BLAKE2-based PoW for rate control
* Adds rewritten FPC package
* Adds possibility to change config options with environment variables
* Adds sample Grafana dashboard for local node instances
* Adds snapshot-file import
* Adds "dirty-flag" to the `Database` plugin to check for corrupted databases
* Adds BadgerDB gargbage collection on startup/shutdown and removes periodic GC
* Adds review-dog linter and automatic integration tests execution to continuous integration pipeline
* Adds `Prometheus` exporter plugin with exposure for [following metrics](https://github.com/iotaledger/goshimmer/issues/465)
* Adds `Sync` plugin keeping track of the node's synchronization state
* Adds `Issuer` plugin which takes care of issuing messages (and blocking any issuance when the node is desynced)
* Adds `Profiling` plugin which exposes the `pprof` endpoints and is now configurable
* Adds `Bootstrap` plugin which continuously issues messages to keep the comm. layer graph growing
* Adds proper metrics collection in the `Metrics` plugin
* Adds support for the new HTTP API routes to the Go client library
* Adds `tools/docker-network` to run an isolated GoShimmer network with a chosen amount of nodes. Predefined local and global grafana dashboards are spinned up as well for the network
* Upgrades `hive.go` with various improvements and fixes
* Fixes bind address prints to not be normalized
* Fixes usage of WebSocket message rate limiter
* Fixes disabled/enabled plugins list in info response
* Fixes the `Graceful Shutdown` plugin not showing the actual background workers during shutdown. The node operator is now
able to see pending background workers in the order in which they are supposed to be terminated
* Refactors display of IDs to use Base58 throughout the entire codebase
* Increases network and database versions
* Removes usage of non std `errors` package
* Removes `Graph` plugin
* Renames `SPA` plugin to `Dashboard`
* Makes running GoShimmer with a config file mandatory if not explicitly bypassed

# v0.1.3 - 2020-03-16

* Update SPA plugin's JS dependencies
* Upgrade `github.com/gobuffalo/packr` to v2.8.0
* Resolves a security issue in the packr dependency, which would allow an unauthenticated attacker to read any
 file from the application's filesystem 

# v0.1.2 - 2020-02-24

* Adds `--version` flag to retrieve the GoShimmer version
* Adds the version and commit hash to the remote log logging
* Replaces the autopeering module with the one from hive.go 
* Changed the pprof listen port to `6061` to avoid conflict with Hornet
* Fixes `invalid stored peer` messages
* Fixes masternodes getting removed if they were offline
* Fixes `-c` and `-d` to define config file/dir
* Fixes drop messages about full queues appearing too many times
* Fixes crash due to incopatible transaction size
* Changed the salt lifetime to 2 hours from 30 minutes

# v0.1.1 - 2020-02-07

This release contains a series of fixes:
* Adds logging of the underlying error when a neighbor connection couldn't be established
* Adds the `RemoteLog` plugin (disabled per default) which sends log messages to a centralized logging service
* Removes the status screen plugin in favor of the SPA/dashboard
* Fixes the neighbor send queue being too small causing spam in the log
* Fixes a deadlock which occurred when a neighbor disconnected while at the same time a `request transaction` 
packet was getting processed which was received by the disconnected neighbor
* Fixes memory consumption by disabling BadgerDB's compression
* Fixes analysis server WebSocket replays freezing the backend code
* Fixes sending on the neighbor send queue if the neighbor is disconnected
* Fixes `BufferedConnection`'s read and written bytes to be atomic

# v0.1.0 - 2020-01-31

> Note that this release is a complete breaking change, therefore node operators are instructed to upgrade.

This release mainly integrates the shared codebase between Hornet and GoShimmer called [hive.go](https://github.com/iotaledger/hive.go), 
to foster improvements in both projects simultaneously. Additionally, the autopeering code from [autopeering-sim](https://github.com/iotaledger/autopeering-sim)
has been integrated into the codebase. For developers a Go client library is now available to communicate
with a GoShimmer node in order to issue, get and search zero value transactions and retrieve neighbor information.

A detailed list about the changes in v0.1.0 can be seen under the given [milestone](https://github.com/iotaledger/goshimmer/milestone/1?closed=1).

* Adds config keys to make bind addresses configurable
* Adds a `relay-checker` tool to check transaction propagation
* Adds database versioning
* Adds concrete shutdown order for plugins
* Adds BadgerDB garbage collection routine
* Adds SPA dashboard with TPS and memory chart, neighbors and Tangle explorer (`spa` plugin)
* Adds option to manually define autopeering seed/private keys
* Adds Golang CI and GitHub workflows as part of the continuous integration pipeline
* Adds back off policies when doing network operations
* Adds an open port check which checks whether NATs are configured properly
* Adds `netutil` package
* Adds the glumb visualizer backend as a plugin called `graph` (has to be manually enabled, check the readme)
* Adds rudimentary PoW as a rate control mechanism
* Adds the autopeering code from autopeering-sim
* Adds association between transactions and their address
* Adds the possibility to allow a node to function/start with gossiping disabled
* Adds buffered connections for gossip messages
* Adds an OAS/Swagger specification file for the web API
* Adds web API and refactors endpoints: `broadcastData`, `findTransactionHashes`,
`getNeighbors`, `getTransactionObjectsByHash`, `getTransactionTrytesByHash`, `getTransactionsToApprove`, `spammer`
* Adds a Go client library over the web API
* Adds a complete rewrite of the gossiping layer
* Fixes the autopeering visualizer to conform to the newest autopeering changes.  
  The visualizer can be accessed [here](http://ressims.iota.cafe/).
* Fixes parallel connections by improving the peer selection mechanism
* Fixes that LRU caches are not flushed correctly up on shutdown
* Fixes consistent application naming (removes instances of sole "Shimmer" in favor of "GoShimmer")
* Fixes several analysis server related issues
* Fixes race condition in the PoW code
* Fixes race condition in the `daemon` package
* Fixes several race conditions in the `analysis` package
* Fixes the database not being closed when the node shuts down
* Fixes `webauth` plugin to function properly with the web API
* Fixes solidification related issues
* Fixes several instances of overused goroutine spawning
* Upgrades to [BadgerDB](https://github.com/dgraph-io/badger) v2.0.1
* Upgrades to latest [iota.go](https://github.com/iotaledger/iota.go) library
* Removes sent count from spammed transactions
* Removes usage of `errors.Identifiable` and `github.com/pkg/errors` in favor of standard lib `errors` package
* Use `network`, `parameter`, `events`, `database`, `logger`, `daemon`, `workerpool` and `node` packages from hive.go
* Removes unused plugins (`zmq`, `dashboard`, `ui`)
