# xhub

**xhub** is a web service for the [`xpub`](https://github.com/rcc-uchicago/xpub) client.  It's a backend server for storing and retrieving [xromm]()-related study data.

The `xpub` client is an alternative command-line interface to the [`xromm.rcc.uchicago.edu`](http://xromm.rcc.uchicago.edu/) data portal.

* create new studies
* create new trials
* initiate async file transfers to the Midway storage cluster
* collect metadata attributes of the files to be transferred

**xhub** stores/retrieves the JSON-encoded metadata sent/requested by the `xpub`
client.
