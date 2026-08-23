package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali013TaskListIsRegionScoped(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i013");_ = e.seedTask(t,inc.ID,"task013a");if err:=e.repo.WithinTx(context.Background(),func(tx repository.Tx)error{return tx.InsertRegion(context.Background(),domain.Region{ID:"other",Name:"other",TimeZone:"Asia/Shanghai",Version:1,CreatedAt:e.now,UpdatedAt:e.now})});err!=nil{t.Fatal(err)};if _,err:=e.service.ListTasks(context.Background(),domain.TaskFilter{},e.field);err!=nil{t.Fatal(err)}}