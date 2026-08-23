package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali002DuplicateRegionRejected(t *testing.T){e:=newTestEnvironment(t);w,err:=e.service.PublishWarning(context.Background(),domain.PublishWarningCommand{Warning:domain.Warning{ID:"w002",TyphoonName:"苍山云",Number:"2",Level:domain.WarningBlue,IssuedAt:e.now,EffectiveFrom:e.now,EffectiveUntil:e.now.Add(time.Hour)},Actor:e.duty});if err!=nil{t.Fatal(err)};_,err=e.service.ActivateIncident(context.Background(),domain.ActivateIncidentCommand{IncidentID:"i002",WarningID:w.ID,Name:"凤阳邑手作活动",Regions:[]string{"coast-a","coast-a"},Level:domain.ResponseIV,Actor:e.duty});if err==nil{t.Fatal("duplicate region accepted")}}