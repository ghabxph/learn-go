package blog

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/ghabxph/gabriel/blog-server/internal/pkg/test"
	"github.com/ghabxph/gabriel/blog-server/internal/pkg/util"
	"gopkg.in/yaml.v3"
)

type test_struct struct {
	Test []struct {
		Title       string `yaml:"title"`
		Route       string `yaml:"route"`
		Code        int    `yaml:"code"`
		Page        string `yaml:"page"`
		ContentType string `yaml:"contentType"`
	}
}

func TestStaticFileServer(t *testing.T) {
	// Sets the test environment
	os.Setenv("SFS_WEB_DIR", util.Path(util.Workdir()+"/../../../tests/web"))

	// Starting the dummy webserver
	srv := httptest.NewServer(Router())
	defer srv.Close()

	var tt *test_struct
	yaml.Unmarshal(util.ReadFileB(util.Path(util.Workdir()+"/../../../tests/internal/app/blog/static_file_server_test.yaml")), &tt)
	for _, tc := range tt.Test {
		t.Run(tc.Title, func(t *testing.T) {
			res, err := http.Get(srv.URL + "/" + tc.Route)
			if err != nil {
				t.Fatalf("http.Get failed: %v", err)
			}
			defer res.Body.Close()
			test.Assert(t, tc.Title, strconv.Itoa(res.StatusCode), strconv.Itoa(tc.Code))
			test.Assert(t, tc.Title, util.ReadAllS(res.Body), util.ReadFileS(os.Getenv("SFS_WEB_DIR")+"/"+tc.Page))
			test.Assert(t, tc.Title, res.Header.Get("Content-Type"), tc.ContentType)
		})
	}
}
