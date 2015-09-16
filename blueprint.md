# Blueprint

The following is a high-level architectural sketch of **xhub**, a Go-based web service developed by the RCC.  

We consider the main [resources](https://github.com/rcc-uchicago/xhub/blob/master/blueprint.md#resources) it deals with, the [buckets](https://github.com/rcc-uchicago/xhub/blob/master/blueprint.md#buckets) in which information about these resources is persisted, and the [routes](https://github.com/rcc-uchicago/xhub/blob/master/blueprint.md#routes) where this information can be sent/retrieved. 

Note that we're utilizing [`buckets`](https://github.com/joyrexus/buckets) for data persistence and [`httprouter`](https://github.com/julienschmidt/httprouter) for URL routing.


## Resources

A resource is ... [BRIEFLY EXPLAIN W/ LINK TO MORE INFO]. 

Xhub deals with the following resources:

* **study**
* **trial**
* **file**
* user
* group

The first three reflect the basic resources comprising an experimental setup:

    studies/

        STUDY_A/
            files/
                FILE_1
                FILE_2
                FILE_3
            trials/
                TRIAL_1/
                    FILE_3
                    FILE_4
                    FILE_5
                TRIAL_2/
                    FILE_6
                    FILE_7
                    FILE_8
                ...
                TRIAL_N/
                    ...

        STUDY_B/
            files/
                FILES
            trials/
                TRIAL/
                    FILES

        ...

        STUDY_N
            ...


## Buckets

A bucket is ... [BRIEFLY EXPLAIN W/ LINK TO MORE INFO]. 

Xhub uses the following buckets for storing and retrieving
information about resources:

* `META` - namespace for incoming resource metadata
* `STUDIES` - list of study names for prefix scans in `META`
* `FILES` - log of file transfers (date, status, location)
* `USERS` - list of users
* `GROUPS` - user groups for specifying group permissions on resources
* `STORES` - list of SQL-based datastores and their SQL statements
* `CONFIG` - JSON-config files for `xpub`


## Routes / Handles

A route is ... [BRIEFLY EXPLAIN W/ LINK TO MORE INFO]. 

A handle is ... [BRIEFLY EXPLAIN W/ LINK TO MORE INFO]. 

Xhub uses the following routes/handles for storing and retrieving
information about resources:


    mux.GET("/", Index)

    h := NewHandles(host, port, db)
    mux := httprouter.New() 
    mux.GET("/", Index)

	// study handles
    mux.GET("/studies", h.studies.List)
    mux.POST("/studies", h.studies.Create)
    mux.GET("/studies/:id", h.studies.Get)
    mux.PUT("/studies/:id", h.studies.Replace)
    mux.DELETE("/studies/:id", h.studies.Delete)

	// file handles (study level)
	mux.GET("/studies/:study/files", h.files.List)
	mux.POST("/studies/:study/files", h.files.Create)
	mux.GET("/studies/:study/files/:id", h.files.Get)
	mux.PUT("/studies/:study/files/:id", h.files.Replace)
	mux.DELETE("/studies/:study/files/:id", h.files.Delete)

	// trial handles (verbose routes)
	mux.GET("/studies/:study/trials", h.trials.List)
	mux.POST("/studies/:study/trials", h.trials.Create)
	mux.GET("/studies/:study/trials/:id", h.trials.Get)
	mux.PUT("/studies/:study/trials/:id", h.trials.Replace)
	mux.DELETE("/studies/:study/trials/:id", h.trials.Delete)

	// trial handles (convenience routes)
	mux.GET("/trials/:study", h.trials.List)
	mux.POST("/trials/:study", h.trials.Create)
	mux.GET("/trials/:study/:id", h.trials.Get)
	mux.PUT("/trials/:study/:id", h.trials.Replace)
	mux.DELETE("/trials/:study/:id", h.trials.Delete)

	// file handles (trial level)
	mux.GET("/files/:study/:trial", h.files.List)
	mux.POST("/files/:study/:trial", h.files.Create)
	mux.GET("/files/:study/:trial/:id", h.files.Get)
	mux.PUT("/files/:study/:trial/:id", h.files.Replace)
	mux.DELETE("/files/:study/:trial/:id", h.files.Delete)
