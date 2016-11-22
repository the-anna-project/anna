package redis

import (
	"strings"
	"testing"

	"github.com/the-anna-project/log"
	servicespec "github.com/the-anna-project/spec/service"
)

// rootLogger implements spec.RootLogger and is used to capture log messages.
type rootLogger struct {
	Args []interface{}
}

func (rl *rootLogger) ArgsToString() string {
	args := ""
	for _, v := range rl.Args {
		if arg, ok := v.(string); ok {
			args += arg
		}
	}
	return args
}

func (rl *rootLogger) Println(v ...interface{}) {
	rl.Args = v
}

func (rl *rootLogger) ResetArgs() {
	rl.Args = []interface{}{}
}

func testMustNewRootLogger(t *testing.T) servicespec.RootLogger {
	return &rootLogger{Args: []interface{}{}}
}

func Test_RedisStorage_retryErrorLogger(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLog := log.New()
	newLog.SetRootLogger(newRootLogger)

	err := newLog.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err := newLog.Configure()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newStorage.(*storage).retryErrorLogger(invalidConfigError, 0)
	result := newRootLogger.(*rootLogger).ArgsToString()

	if !strings.Contains(result, invalidConfigError.Error()) {
		t.Fatal("expected", invalidConfigError.Error(), "got", result)
	}
}

func Test_RedisStorage_withPrefix(t *testing.T) {
	newConfig := DefaultStorageConfig()
	newConfig.Prefix = "test-prefix"
	newStorage, err := NewStorage(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	expected := "test-prefix:my:test:key"
	newKey := newStorage.(*storage).withPrefix("my", "test", "key")
	if newKey != expected {
		t.Fatal("expected", expected, "got", newKey)
	}
}

func Test_RedisStorage_withPrefix_Empty(t *testing.T) {
	newConfig := DefaultStorageConfig()
	newConfig.Prefix = "test-prefix"
	newStorage, err := NewStorage(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newKey := newStorage.(*storage).withPrefix()
	if newKey != newConfig.Prefix {
		t.Fatal("expected", newConfig.Prefix, "got", newKey)
	}
}
