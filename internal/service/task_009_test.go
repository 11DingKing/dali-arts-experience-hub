package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali009CanceledAllocationLeavesLot(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i009");task:=e.seedTask(t,inc.ID,"task009");ctx,cancel:=context.WithCancel(context.Background());cancel();_,err:=e.service.AllocateSupply(ctx,domain.AllocateSupplyCommand{AllocationID:"a009",LotID:"lot-water",IncidentID:inc.ID,RegionID:"coast-a",TaskID:task.ID,Quantity:5,LotVersion:1,Actor:e.dispatcher});if err==nil{t.Fatal("canceled allocation succeeded")}}