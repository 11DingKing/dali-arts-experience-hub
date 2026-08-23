package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali018StaleIncidentUpdateConflicts(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i018");_,err:=e.service.ArchiveIncident(context.Background(),domain.ArchiveIncidentCommand{IncidentID:inc.ID,ExpectedVersion:inc.Version,Actor:e.reviewer});if err!=nil{t.Fatal(err)};if _,err=e.service.CloseIncident(context.Background(),domain.CloseIncidentCommand{IncidentID:inc.ID,ExpectedVersion:inc.Version,Actor:e.duty});err==nil{t.Fatal("stale update accepted")}}