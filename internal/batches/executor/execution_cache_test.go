
	"github.com/sourcegraph/sourcegraph/lib/batches"
	"github.com/sourcegraph/sourcegraph/lib/batches/execution"
	"github.com/sourcegraph/sourcegraph/lib/batches/execution/cache"
var cacheRepo1 = batches.Repository{
	ID:          "src-cli",
	Name:        "github.com/sourcegraph/src-cli",
	BaseRef:     "refs/heads/main",
	BaseRev:     "d34db33f",
	FileMatches: []string{"README.md", "main.go"},
}
var cacheRepo2 = batches.Repository{
	ID:          "sourcegraph",
	Name:        "github.com/sourcegraph/sourcegraph",
	BaseRef:     "refs/heads/main-2",
	BaseRev:     "c0ff33",
	FileMatches: []string{"main.go"},
	cacheKey1 := &cache.ExecutionKey{
		Repository: cacheRepo1,
	}
	cacheKey2 := &cache.ExecutionKey{
		Repository: cacheRepo2,
	}
	value := execution.Result{
func assertCacheHit(t *testing.T, c ExecutionDiskCache, k *cache.ExecutionKey, want execution.Result) {
func assertCacheMiss(t *testing.T, c ExecutionDiskCache, k *cache.ExecutionKey) {