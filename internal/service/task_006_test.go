package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali006RejectedReceiptReopensTask(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i006");task:=e.seedTask(t,inc.ID,"task006");task,err:=e.service.AdvanceTask(context.Background(),domain.AdvanceTaskCommand{TaskID:task.ID,ExpectedVersion:task.Version,Target:domain.TaskExecuting,Actor:e.dispatcher});if err!=nil{t.Fatal(err)};r,err:=e.service.SubmitReceipt(context.Background(),domain.SubmitReceiptCommand{Receipt:domain.Receipt{ID:"r006",TaskID:task.ID,IncidentID:inc.ID,RegionID:"coast-a",Summary:"体验完成",EvidenceCount:1},TaskVersion:task.Version,Actor:e.field});if err!=nil{t.Fatal(err)};if _,err=e.service.ReviewReceipt(context.Background(),domain.ReviewReceiptCommand{ReceiptID:r.ID,ExpectedVersion:r.Version,Accept:false,Actor:e.reviewer});err!=nil{t.Fatal(err)};got,err:=e.repo.SnapshotIncident(context.Background(),inc.ID);if err!=nil{t.Fatal(err)};for _,v:=range got.Tasks{if v.ID==task.ID && v.Status!=domain.TaskExecuting{t.Fatalf("task not reopened: %s",v.Status)}}}