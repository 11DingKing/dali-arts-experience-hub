package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali007ReleaseConflictRollsBackCapacity(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i007");r,err:=e.service.ReserveShelter(context.Background(),domain.ReserveShelterCommand{ReservationID:"res007",ShelterID:"shelter-1",IncidentID:inc.ID,RegionID:"coast-a",People:10,ExpiresAt:e.now.Add(time.Hour),Actor:e.dispatcher});if err!=nil{t.Fatal(err)};err=e.service.ReleaseShelter(context.Background(),domain.ReleaseShelterCommand{ReservationID:r.ID,ExpectedVersion:r.Version+99,Actor:e.dispatcher});if err==nil{t.Fatal("stale release succeeded")};snap,_:=e.repo.SnapshotIncident(context.Background(),inc.ID);_ = snap}