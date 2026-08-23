package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali008DuplicateReservationConflicts(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i008");cmd:=domain.ReserveShelterCommand{ShelterID:"shelter-1",IncidentID:inc.ID,RegionID:"coast-a",People:3,ExpiresAt:e.now.Add(time.Hour),IdempotencyKey:"same",Actor:e.dispatcher};if _,err:=e.service.ReserveShelter(context.Background(),cmd);err!=nil{t.Fatal(err)};cmd.ReservationID="res008b";if _,err:=e.service.ReserveShelter(context.Background(),cmd);err==nil{t.Fatal("duplicate reservation accepted")}}