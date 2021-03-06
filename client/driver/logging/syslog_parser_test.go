// +build !windows

package logging

import (
	"log"
	"log/syslog"
	"os"
	"testing"
)

func TestLogParser_Priority(t *testing.T) {
	line := []byte("<30>2016-02-10T10:16:43-08:00 d-thinkpad docker/e2a1e3ebd3a3[22950]: 1:C 10 Feb 18:16:43.391 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf")
	d := NewDockerLogParser(log.New(os.Stdout, "", log.LstdFlags))
	p, _, err := d.parsePriority(line)
	if err != nil {
		t.Fatalf("got an err: %v", err)
	}
	if p.Severity != syslog.LOG_INFO {
		t.Fatalf("expected serverity: %v, got: %v", syslog.LOG_INFO, p.Severity)
	}

	idx := d.logContentIndex(line)
	expected := 68
	if idx != expected {
		t.Fatalf("expected idx: %v, got: %v", expected, idx)
	}
}
