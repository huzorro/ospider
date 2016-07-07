package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gosexy/redis"
	"github.com/huzorro/ospider/util"
	"github.com/huzorro/ospider/web/user"
	"github.com/huzorro/spfactor/sexredis"
	"github.com/huzorro/woplus/tools"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	// "github.com/adjust/redismq"
	"net/http"
	"github.com/adjust/rmq"
	"github.com/huzorro/ospider/cms"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/crontab"
    "github.com/huzorro/ospider/attack"
	"github.com/huzorro/ospider/spider/processor"
	"github.com/huzorro/ospider/web/handler"
)

func main() {
	portPtr := flag.String("port", ":10086", "service port")
	redisIdlePtr := flag.Int("redis", 20, "redis idle connections")
	dbMaxPtr := flag.Int("db", 10, "max db open connections")
	//config path
	cfgPathPtr := flag.String("config", "config.json", "config path name")
	//web
	webPtr := flag.Bool("web", false, "web sm start")
	//crontab
	cronTaskPtr := flag.Bool("cron", false, "crontab task")
    //attack crontab
    ca := flag.Bool("ca", false, "attack crontab")
    //reset crontab
    reset := flag.Bool("reset", false, "reset crontab")
    //attack
    attackPtr := flag.Bool("attack", false, "attack task")
	//spider
	spiderPtr := flag.Bool("spider", false, "spider start")
	//restful
	restfulPtr := flag.Bool("restful", false, "restful start")
	//template path
	templatePath := flag.String("template", "templates", "template path")
	//Static path
	staticPath := flag.String("static", "public", "Static path")

	flag.Parse()
	//json config
	var (
		// cfg       user.Cfg
		// cfg       crontab.Cfg
		cfg       processor.Cfg
		mtn       *martini.ClassicMartini
		redisPool *sexredis.RedisPool
		db        *sql.DB
		err       error
	)
	if err := tools.Json2Struct(*cfgPathPtr, &cfg); err != nil {
		log.Printf("load json config fails %s", err)
		panic(err.Error())
	}

	logger := log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	redisPool = &sexredis.RedisPool{Connections: make(chan *redis.Client, *redisIdlePtr), ConnFn: func() (*redis.Client, error) {
		client := redis.New()
		err := client.Connect("localhost", uint(6379))
		return client, err
	}}
	db, err = sql.Open(cfg.Dbtype, cfg.Dburi)
	db.SetMaxOpenConns(*dbMaxPtr)

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	mtn = martini.Classic()
	mtn.Use(martini.Static(*staticPath))
	mtn.Map(logger)
	mtn.Map(redisPool)
	mtn.Map(db)

	cache := &user.Cache{DB: db, Pool: redisPool}
	mtn.Map(cache)
	//	load rbac node
	if nMap, err := cache.RbacNodeToMap(); err != nil {
		logger.Printf("rbac node to map fails %s", err)
	} else {
		mtn.Map(nMap)
	}
	//load rbac menu
	if ms, err := cache.RbacMenuToSlice(); err != nil {
		logger.Printf("rbac menu to slice fails %s", err)
	} else {
		mtn.Map(ms)
	}
	//session
	store := sessions.NewCookieStore([]byte("secret123"))
	mtn.Use(sessions.Sessions("Qsession", store))
	//render
	rOptions := render.Options{}
	rOptions.Extensions = []string{".tmpl", ".html"}
	rOptions.Directory = *templatePath
	rOptions.Funcs = []template.FuncMap{util.FuncMaps}
	mtn.Use(render.Renderer(rOptions))

	mtn.Map(&cfg.Cfg.Cfg)

	if *cronTaskPtr {
		cronTask := crontab.New()
		cronTask.Cfg = cfg.Cfg

		// taskQueue := redismq.CreateQueue("localhost", "6379", "", 0, cfg.TaskQueueName)
		connection := rmq.OpenConnection("cron_producer", "tcp", "localhost:6379", 0)
		cronTask.Ospider = common.Ospider{Log: logger, P: redisPool, Db: db}
		cronTask.Queue = connection.OpenQueue(cfg.TaskQueueName)
		cronTask.Handler()
		// cronTask := &crontab.Task{common.Ospider{Log:logger, P:redisPool, Db:db}, cfg}
	}
    if *ca {
        attackCron := crontab.NewAttack()
        attackCron.Cfg = cfg.Cfg
		pconnection := rmq.OpenConnection("attack_producer", "tcp", "localhost:6379", 0)
		attackCron.Ospider = common.Ospider{Log: logger, P: redisPool, Db: db}  
		attackCron.Queue = pconnection.OpenQueue(cfg.AttackQueueName)        
        attackCron.Handler()                     
    }
    if *reset {
        restCron := crontab.NewRest()
        restCron.Cfg = cfg.Cfg
        restCron.Ospider = common.Ospider{Log: logger, P: redisPool, Db: db} 
        restCron.Handler()          
    }
    if *attackPtr {        
		cconnection := rmq.OpenConnection("attack_consumer", "tcp", "localhost:6379", 0)
		co := common.Ospider{Log: logger, P: redisPool, Db: db}
        processors := make([]common.Processor, 0)
        result := &attack.Task{cfg, co, cconnection, processors}
        result.AddProcessor(attack.NewAttackSubmit(co)).AddProcessor(attack.NewUpdateHost(co))
        result.Handler()                                     
    }
	if *spiderPtr {
		//spider
		// server := redismq.NewServer("localhost", "6379", "", 0, "9999")
		// server.Start()
		// taskQueue := redismq.CreateQueue("localhost", "6379", "", 0, cfg.TaskQueueName)
		connection := rmq.OpenConnection("task_consumer&link_producer", "tcp", "localhost:6379", 0)
		task := &processor.Task{cfg, common.Ospider{Log: logger, P: redisPool, Db: db}, connection}
		task.Handler()

		// resultQueue := redismq.CreateQueue("localhost", "6379", "", 0, cfg.ResultQueueName)
		connectionLinkResult := rmq.OpenConnection("link_consumer&result_producer", "tcp", "localhost:6379", 0)
		result := &processor.TaskContent{cfg, common.Ospider{Log: logger, P: redisPool, Db: db}, connectionLinkResult}
		result.Handler()
	}
	if *restfulPtr {
		connection := rmq.OpenConnection("result_consumer", "tcp", "localhost:6379", 0)
		co := common.Ospider{Log: logger, P: redisPool, Db: db}
        processors := make([]common.Processor, 0)
        result := &cms.Task{cfg, co, connection, processors}
        result.AddProcessor(cms.NewCommitMysql(co)).AddProcessor(cms.NewCommitRestful(co))
        result.Handler()
	}
	if *webPtr {
		//rbac filter
		rbac := &user.RBAC{}
		mtn.Use(rbac.Filter())
		mtn.Get("/login", func(r render.Render) {
			r.HTML(200, "login", "")
		})
		mtn.Get("/logout", user.Logout)
		mtn.Post("/login/check", user.LoginCheck)
		//user

		mtn.Post("/user/view", user.ViewUserAction)
		mtn.Get("/usersview", user.ViewUsersAction)
		mtn.Post("/user/add", user.AddUserAction)
		mtn.Post("/user/edit", user.EditUserAction)
		//spider manager
		mtn.Get("/", handler.GetSites)
		mtn.Post("/spider/one", handler.GetSpider)
		mtn.Post("/spider/edit", handler.EditSpider)
		mtn.Post("/spider/add", handler.AddSpider)
		mtn.Get("/spiderview", handler.GetSpiders)
		//spider restful
		mtn.Post("/api/spiders", handler.GetSpidersApi)
		//spider crontab
		mtn.Get("/crontabview", handler.GetCrontabs)
		mtn.Post("/crontab/one", handler.GetCrontab)
		mtn.Post("/crontab/edit", handler.EditCrontab)
		mtn.Post("/crontab/add", handler.AddCrontab)
		//crontab restful
		mtn.Post("/api/crontabs", handler.GetCrontabsApi)

		//spider rule
		mtn.Get("/ruleview", handler.GetRules)
		mtn.Post("/rule/one", handler.GetRule)
		mtn.Post("/rule/edit", handler.EditRule)
		mtn.Post("/rule/add", handler.AddRule)
		//rule restful
		mtn.Post("/api/rules", handler.GetRulesApi)

		//spider site
		mtn.Get("/siteview", handler.GetSites)
		mtn.Post("/site/one", handler.GetSite)
		mtn.Post("/site/edit", handler.EditSite)
		mtn.Post("/site/add", handler.AddSite)
        
        //flood api
        mtn.Get("/floodapiview", handler.GetFloods)
        mtn.Post("/floodapi/one", handler.GetFloodApi)
        mtn.Post("/floodapi/edit", handler.EditFloodApi)
        mtn.Post("/floodapi/add", handler.AddFloodApi)
        
        //flood target
        mtn.Get("/floodtargetview", handler.GetfloodTargets)
        mtn.Post("/floodtarget/one", handler.GetFloodTarget)
        mtn.Post("/floodtarget/edit", handler.EditfloodTarget)
        mtn.Post("/floodtarget/add", handler.AddfloodTarget)
        
		go http.ListenAndServe(*portPtr, mtn)
	}
	select {}
}
