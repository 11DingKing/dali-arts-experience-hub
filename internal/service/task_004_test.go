package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali004StaleHandoffIsAtomic(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i004");task:=e.seedTask(t,inc.ID,"task004");d,err:=e.service.DispatchTeam(context.Background(),domain.DispatchTeamCommand{DispatchID:"dispatch004",TeamID:"team-1",IncidentID:inc.ID,RegionID:"coast-a",TaskID:task.ID,TeamVersion:1,Actor:e.dispatcher});if err!=nil{t.Fatal(err)};_,err=e.service.HandoffTeam(context.Background(),domain.HandoffTeamCommand{DispatchID:d.ID,FromTeamID:"team-1",ToTeamID:"team-2",FromVersion:999,ToVersion:1,Actor:e.dispatcher});if err==nil{t.Fatal("stale handoff succeeded")};snap,_:=e.repo.SnapshotIncident(context.Background(),inc.ID);if len(snap.Dispatches)!=0{t.Fatal("partial handoff persisted")}}