package traversal

import (
	"context"
	"fmt"
	"io"

	ipld "github.com/ipld/go-ipld-prime"
	ipldfree "github.com/ipld/go-ipld-prime/impl/free"
	"github.com/ipld/go-ipld-prime/impl/typed"
)

// init sets all the values in TraveralConfig to reasonable defaults
// if they're currently the zero value.
func (tc *Config) init() {
	if tc.Ctx == nil {
		tc.Ctx = context.Background()
	}
	if tc.LinkLoader == nil {
		tc.LinkLoader = func(ipld.Link, ipld.LinkContext) (io.Reader, error) {
			return nil, fmt.Errorf("no link loader configured")
		}
	}
	if tc.LinkNodeBuilderChooser == nil {
		tc.LinkNodeBuilderChooser = func(lnk ipld.Link, lnkCtx ipld.LinkContext) ipld.NodeBuilder {
			if tlnkNd, ok := lnkCtx.LinkNode.(typed.LinkNode); ok {
				return tlnkNd.ReferencedNodeBuilder()
			}
			return ipldfree.NodeBuilder()
		}
	}
	if tc.LinkStorer == nil {
		tc.LinkStorer = func(ipld.LinkContext) (io.Writer, ipld.StoreCommitter, error) {
			return nil, nil, fmt.Errorf("no link storer configured")
		}
	}
}

func (prog *Progress) init() {
	if prog.Cfg == nil {
		prog.Cfg = &Config{}
	}
	prog.Cfg.init()
}
