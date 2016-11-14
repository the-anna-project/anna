package boot

import (
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/object/config"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new boot command.
func New() *Command {
	return &Command{}
}

// Command represents the boot command.
type Command struct {
	// Dependencies.

	configCollection  *config.Collection
	serviceCollection servicespec.Collection

	// Settings.

	bootOnce     sync.Once
	gitCommit    string
	shutdownOnce sync.Once
	version      string
}

// Boot makes the neural network boot and run.
func (c *Command) Boot() {
	c.serviceCollection = c.newServiceCollection()
	go c.serviceCollection.Boot()

	go c.ListenToSignal()

	// Block the main goroutine forever. The process is only supposed to be ended
	// by a call to Shutdown or ForceShutdown.
	select {}
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	c.Boot()
}

// ForceShutdown forces the process to stop immediately.
func (c *Command) ForceShutdown() {
	os.Exit(0)
}

// New creates a new cobra command for the boot command.
func (c *Command) New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "boot",
		Short: "Boot and run the anna daemon.",
		Long:  "Boot and run the anna daemon.",
		Run:   c.Execute,
	}

	c.configCollection.Endpoint().Text().SetAddress(newCmd.PersistentFlags().String("endpoint.text.address", "127.0.0.1:9119", "host:port to bind the text endpoint to"))
	c.configCollection.Endpoint().Metric().SetAddress(newCmd.PersistentFlags().String("endpoint.metric.address", "127.0.0.1:9120", "host:port to bind the metric endpoint to"))

	c.configCollection.Space().Connection().SetWeight(newCmd.PersistentFlags().Int("space.connection.weight", 0, "default weight of new connections within the connection space"))

	c.configCollection.Space().Dimension().SetCount(newCmd.PersistentFlags().Int("space.dimension.count", 3, "default number of directional coordinates within the connection space"))
	c.configCollection.Space().Dimension().SetDepth(newCmd.PersistentFlags().Int("space.dimension.depth", 1000000, "default size of each directional coordinate within the connection space"))

	c.configCollection.Space().Peer().SetPosition(newCmd.PersistentFlags().String("space.peer.position", "0,0,0", "default position of new peers within the connection space"))

	c.configCollection.Storage().Connection().SetAddress(newCmd.PersistentFlags().String("storage.connection.address", "127.0.0.1:6379", "host:port to connect to connection storage"))
	c.configCollection.Storage().Connection().SetKind(newCmd.PersistentFlags().String("storage.connection.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().Connection().SetPrefix(newCmd.PersistentFlags().String("storage.connection.prefix", "anna", "prefix used to prepend to connection storage keys"))

	c.configCollection.Storage().Feature().SetAddress(newCmd.PersistentFlags().String("storage.feature.address", "127.0.0.1:6380", "host:port to connect to feature storage"))
	c.configCollection.Storage().Feature().SetKind(newCmd.PersistentFlags().String("storage.feature.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().Feature().SetPrefix(newCmd.PersistentFlags().String("storage.feature.prefix", "anna", "prefix used to prepend to feature storage keys"))

	c.configCollection.Storage().General().SetAddress(newCmd.PersistentFlags().String("storage.general.address", "127.0.0.1:6381", "host:port to connect to general storage"))
	c.configCollection.Storage().General().SetKind(newCmd.PersistentFlags().String("storage.general.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().General().SetPrefix(newCmd.PersistentFlags().String("storage.general.prefix", "anna", "prefix used to prepend to general storage keys"))

	return newCmd
}

// ListenToSignal listens to OS signals to be catched and processed if desired.
func (c *Command) ListenToSignal() {
	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go c.Shutdown()

	<-listener

	c.ForceShutdown()
}

// SetConfigCollection sets the config collection for the boot command to
// configure the neural network.
func (c *Command) SetConfigCollection(configCollection *config.Collection) {
	c.configCollection = configCollection
}

// SetGitCommit sets the git commit for the version command to be displayed.
func (c *Command) SetGitCommit(gitCommit string) {
	c.gitCommit = gitCommit
}

// Shutdown initializes the shutdown of the neural network.
func (c *Command) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.serviceCollection.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		os.Exit(0)
	})
}
