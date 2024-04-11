package cli

import (
	"fmt"
	"os/signal"
	"strings"
	"syscall"

	"github.com/xpzouying/go-clean-code-template/internal/config"
	"github.com/xpzouying/go-clean-code-template/internal/constant"
	"github.com/xpzouying/go-clean-code-template/internal/domain"
	"github.com/xpzouying/go-clean-code-template/internal/repo"
	"github.com/xpzouying/go-clean-code-template/internal/service"
	"github.com/xpzouying/go-clean-code-template/internal/usecase"
	"github.com/xpzouying/go-clean-code-template/log"

	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var cliLogger = log.MustNewLogger("info").Sugar().Named("cli-app")

func startAction(c *cli.Context) error {
	cfg, err := InitConfigAndComponents()
	if err != nil {
		cliLogger.Fatalf("Failed to initialize config and components: %v", err)
	}

	// --- begin init service ---

	svc, err := newService(cfg)
	if err != nil {
		cliLogger.Fatalf("Failed to initialize service: %v", err)
	}

	// --- end init service ---

	mainCtx, stop := signal.NotifyContext(c.Context, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg := errgroup.Group{}

	wg.Go(func() error {
		if err := startHTTPServer(mainCtx, cfg, svc); err != nil {
			return errors.Wrap(err, "failed to start HTTP server")
		}
		return nil
	})

	if err := wg.Wait(); err != nil {
		cliLogger.Errorf("start actions error: %v", err)
		return err
	}

	cliLogger.Info("Starting...")
	return nil
}

func CreateCliApp() *cli.App {

	cli.VersionPrinter = func(c *cli.Context) {
		println("Welcome!")
		fmt.Printf("Version: %s\n", constant.Version)
		fmt.Printf("Build Time: %s\n", constant.BuildTime)
		fmt.Printf("Git Commit: %s\n", constant.GitRevision)
	}

	app := cli.NewApp()
	app.Name = "go-template-project"
	app.Flags = RootFlags
	app.Version = constant.Version
	app.Usage = "go-template-project is a template project for go command line application."
	app.Action = startAction

	return app
}

func newService(cfg *config.Config) (*service.UserService, error) {

	db, err := newDBConn(cfg.DBConnStr)
	if err != nil {
		return nil, err
	}

	var userRepo domain.UserRepo
	{
		userRepo = repo.NewUserRepo(db)
	}

	var userUC *usecase.UserUsecase
	{
		userUC = usecase.NewUserUsecase(userRepo)
	}

	return service.NewUserService(userUC), nil
}

func newDBConn(dbConnStr string) (*gorm.DB, error) {

	dbType, dbDsn, err := parseDBConnStr(dbConnStr)
	if err != nil {
		return nil, err
	}

	// simply support sqlite3
	if dbType != "sqlite3" {
		return nil, errors.New("unsupported database type")
	}

	return gorm.Open(sqlite.Open(dbDsn), &gorm.Config{
		PrepareStmt: true,
	})

}

func parseDBConnStr(dbConnStr string) (dbType, dbDsn string, err error) {

	// 使用 SplitN 将字符串分割成两部分，限制为只分割一次，确保只分离协议和后续部分
	parts := strings.SplitN(dbConnStr, "://", 2)
	if len(parts) != 2 {
		// 如果没有找到预期的分割，返回空字符串
		err = errors.New("invalid db connection string")
		return
	}

	dbType = parts[0]
	dbDsn = parts[1]
	return
}
