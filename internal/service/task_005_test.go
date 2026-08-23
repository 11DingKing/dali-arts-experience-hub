package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali005ConcurrentDispatchSingleWinner(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i005");a:=e.seedTask(t,inc.ID,"task005a");b:=e.seedTask(t,inc.ID,"task005b");start:=make(chan struct{});errs:=make(chan error,2);go func(){<-start;_,err:=e.service.DispatchTeam(context.Background(),domain.DispatchTeamCommand{DispatchID:"d005a",TeamID:"team-1",IncidentID:inc.ID,RegionID:"coast-a",TaskID:a.ID,TeamVersion:1,Actor:e.dispatcher});errs<-err}();go func(){<-start;_,err:=e.service.DispatchTeam(context.Background(),domain.DispatchTeamCommand{DispatchID:"d005b",TeamID:"team-1",IncidentID:inc.ID,RegionID:"coast-a",TaskID:b.ID,TeamVersion:1,Actor:e.dispatcher});errs<-err}();close(start);e1,e2:=<-errs,<-errs;if (e1==nil)==(e2==nil){t.Fatalf("expected one winner: %v %v",e1,e2)}}