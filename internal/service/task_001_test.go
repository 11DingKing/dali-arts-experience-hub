package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali001CancelledActivationHasNoState(t *testing.T){ e:=newTestEnvironment(t); w,err:=e.service.PublishWarning(context.Background(),domain.PublishWarningCommand{Warning:domain.Warning{ID:"w001",TyphoonName:"大理夏夜",Number:"1",Level:domain.WarningBlue,IssuedAt:e.now,EffectiveFrom:e.now,EffectiveUntil:e.now.Add(time.Hour)},Actor:e.duty});if err!=nil{t.Fatal(err)};ctx,cancel:=context.WithCancel(context.Background());cancel();_,err=e.service.ActivateIncident(ctx,domain.ActivateIncidentCommand{IncidentID:"i001",WarningID:w.ID,Name:"床单厂夜游",Regions:[]string{"coast-a"},Level:domain.ResponseIV,Actor:e.duty});if err==nil{t.Fatal("cancelled activation succeeded")};if _,err=e.repo.SnapshotIncident(context.Background(),"i001");err==nil{t.Fatal("cancelled activation persisted")}}