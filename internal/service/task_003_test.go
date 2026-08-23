package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali003PendingReceiptBlocksClose(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i003");task:=e.seedTask(t,inc.ID,"task003");task,err:=e.service.AdvanceTask(context.Background(),domain.AdvanceTaskCommand{TaskID:task.ID,ExpectedVersion:task.Version,Target:domain.TaskExecuting,Actor:e.dispatcher});if err!=nil{t.Fatal(err)};_,err=e.service.SubmitReceipt(context.Background(),domain.SubmitReceiptCommand{Receipt:domain.Receipt{ID:"receipt003",TaskID:task.ID,IncidentID:inc.ID,RegionID:"coast-a",Summary:"手作体验凭证",EvidenceCount:1},TaskVersion:task.Version,Actor:e.field});if err!=nil{t.Fatal(err)};if _,err=e.service.CloseIncident(context.Background(),domain.CloseIncidentCommand{IncidentID:inc.ID,ExpectedVersion:inc.Version,Actor:e.duty});err==nil{t.Fatal("closed with pending receipt")}}