package fetch

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rusq/dlog"
	"github.com/rusq/slackdump/v2/internal/chunk/processor"
	"github.com/rusq/slackdump/v2/internal/structures"
	"github.com/slack-go/slack"
)

type Parameters struct {
	Oldest    time.Time
	Latest    time.Time
	List      *structures.EntityList
	DumpFiles bool
}

type streamer interface {
	Client() *slack.Client
	Stream(context.Context, processor.Conversations, string, time.Time, time.Time) error
}

var replacer = strings.NewReplacer("/", "-", ":", "-")

// Conversation dumps a single conversation or thread into a directory,
// returning the name of the state file that was created.  State file contains
// the information about the filename of the chunk recording file, as well as
// paths to downloaded files.
func Conversation(ctx context.Context, sess streamer, dir string, link string, p *Parameters) (string, error) {
	fileprefix := replacer.Replace(link)
	var pattern = fmt.Sprintf("%s-*.jsonl.gz", fileprefix)
	f, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gz := gzip.NewWriter(f)
	defer gz.Close()

	pr, err := processor.NewStandard(ctx, gz, sess.Client(), dir, processor.DumpFiles(p.DumpFiles))
	if err != nil {
		return "", err
	}
	defer pr.Close()
	state, err := pr.State()
	if err != nil {
		return "", err
	}
	state.SetFilename(filepath.Base(f.Name()))
	state.SetIsCompressed(true)
	if p.DumpFiles {
		state.SetFilesDir(fileprefix)
	}
	statefile := filepath.Join(dir, fileprefix+".state")
	defer func() {
		// we are deferring this so that it would execute even if the error
		// has occurred to have a consistent state.
		if err := state.Save(statefile); err != nil {
			dlog.Print(err)
			return
		}
	}()
	if err := sess.Stream(ctx, pr, link, p.Oldest, p.Latest); err != nil {
		return statefile, err
	}
	if ctx.Err() != nil {
		return statefile, ctx.Err()
	}
	state.SetIsComplete(true)
	return statefile, nil
}
