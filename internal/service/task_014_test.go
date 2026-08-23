package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali014TaskPageCountMatchesItems(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i014");task:=e.seedTask(t,inc.ID,"task014");p,err:=e.service.ListTasks(context.Background(),domain.TaskFilter{Status:domain.TaskPending},e.duty);if err!=nil{t.Fatal(err)};if p.Total< len(p.Items){t.Fatal("total smaller than page")};_ = task}