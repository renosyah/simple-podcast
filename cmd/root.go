package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/renosyah/simple-podcast/auth"
	"github.com/renosyah/simple-podcast/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbPool  *sql.DB
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use: "app",
	PreRun: func(cmd *cobra.Command, args []string) {
		router.Init(dbPool)
		auth.Init(dbPool, &auth.Option{
			Enable: viper.GetBool("auth.enable"),
		})
	},
	Run: func(cmd *cobra.Command, args []string) {

		r := mux.NewRouter()

		// users auth
		r.Handle("/users/auth/login", router.HandlerFunc(router.HandlerLoginUser)).Methods(http.MethodPost)
		r.Handle("/users/auth/register", router.HandlerFunc(router.HandlerRegisterUser)).Methods(http.MethodPost)

		apiRouter := r.PathPrefix("/api/v1").Subrouter()
		apiRouter.Use(auth.AuthenticationMiddleware)

		// users api
		apiRouter.Handle("/users", router.HandlerFunc(router.HandlerAddUser)).Methods(http.MethodPost)
		apiRouter.Handle("/users-list", router.HandlerFunc(router.HandlerAllUser)).Methods(http.MethodPost)
		apiRouter.Handle("/users/{id}", router.HandlerFunc(router.HandlerOneUser)).Methods(http.MethodGet)
		apiRouter.Handle("/users/{id}", router.HandlerFunc(router.HandlerUpdateUser)).Methods(http.MethodPut)
		apiRouter.Handle("/users/{id}", router.HandlerFunc(router.HandlerDeleteUser)).Methods(http.MethodDelete)

		// categories api
		apiRouter.Handle("/categories", router.HandlerFunc(router.HandlerAddCategory)).Methods(http.MethodPost)
		apiRouter.Handle("/categories-list", router.HandlerFunc(router.HandlerAllCategory)).Methods(http.MethodPost)
		apiRouter.Handle("/categories/{id}", router.HandlerFunc(router.HandlerOneCategory)).Methods(http.MethodGet)
		apiRouter.Handle("/categories/{id}", router.HandlerFunc(router.HandlerUpdateCategory)).Methods(http.MethodPut)
		apiRouter.Handle("/categories/{id}", router.HandlerFunc(router.HandlerDeleteCategory)).Methods(http.MethodDelete)

		// musics api
		apiRouter.Handle("/musics", router.HandlerFunc(router.HandlerAddMusic)).Methods(http.MethodPost)
		apiRouter.Handle("/musics-list", router.HandlerFunc(router.HandlerAllMusic)).Methods(http.MethodPost)
		apiRouter.Handle("/musics/{id}", router.HandlerFunc(router.HandlerOneMusic)).Methods(http.MethodGet)
		apiRouter.Handle("/musics/{id}", router.HandlerFunc(router.HandlerUpdateMusic)).Methods(http.MethodPut)
		apiRouter.Handle("/musics/{id}", router.HandlerFunc(router.HandlerDeleteMusic)).Methods(http.MethodDelete)

		// serve music file
		apiRouter.Handle("/musics-file/{id}", http.HandlerFunc(router.HandlerFileMusic)).Methods(http.MethodGet)

		// podcasts api
		apiRouter.Handle("/podcasts", router.HandlerFunc(router.HandlerAddPodcast)).Methods(http.MethodPost)
		apiRouter.Handle("/podcasts-list", router.HandlerFunc(router.HandlerAllPodcast)).Methods(http.MethodPost)
		apiRouter.Handle("/podcasts/{id}", router.HandlerFunc(router.HandlerOnePodcast)).Methods(http.MethodGet)
		apiRouter.Handle("/podcasts/{id}", router.HandlerFunc(router.HandlerUpdatePodcast)).Methods(http.MethodPut)
		apiRouter.Handle("/podcasts/{id}", router.HandlerFunc(router.HandlerDeletePodcast)).Methods(http.MethodDelete)

		// upload file
		apiRouter.Handle("/upload", router.HandlerFunc(router.HandlerUploadFile)).Methods(http.MethodPost)

		// r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("files"))))

		port := viper.GetInt("app.port")
		p := os.Getenv("PORT")
		if p != "" {
			port, _ = strconv.Atoi(p)
		}

		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      r,
			ReadTimeout:  time.Duration(viper.GetInt("read_timeout")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt("write_timeout")) * time.Second,
			IdleTimeout:  time.Duration(viper.GetInt("idle_timeout")) * time.Second,
		}

		done := make(chan bool, 1)
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, os.Interrupt)

		go func() {
			<-quit
			log.Println("Server is shutting down...")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			server.SetKeepAlivesEnabled(false)
			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
			}
			close(done)
		}()

		log.Println("Server is ready to handle requests at", fmt.Sprintf(":%d", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", fmt.Sprintf(":%d", port), err)
		}

		<-done
		log.Println("Server stopped")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is github.com/renosyah/simple-podcast/.server.toml)")
	cobra.OnInitialize(initConfig, initDB)
}

func initDB() {

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"))

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error open DB: %v\n", err))
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error ping DB: %v\n", err))
		return
	}

	dbPool = db
}

func initConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/simple-podcast")
		viper.SetConfigName(".server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
