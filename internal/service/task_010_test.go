package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali010CancelAllocationRestoresLot(t *testing.T){e:=newTestEnvironment(t);inc:=e.seedIncident(t,"i010");task:=e.seedTask(t,inc.ID,"task010");a,err:=e.service.AllocateSupply(context.Background(),domain.AllocateSupplyCommand{AllocationID:"a010",LotID:"lot-water",IncidentID:inc.ID,RegionID:"coast-a",TaskID:task.ID,Quantity:5,LotVersion:1,Actor:e.dispatcher});if err!=nil{t.Fatal(err)};if err=e.service.CancelAllocation(context.Background(),domain.CancelAllocationCommand{AllocationID:a.ID,AllocationVersion:a.Version,Actor:e.dispatcher});err!=nil{t.Fatal(err)};var reserved int;if err=e.repo.WithinTx(context.Background(),func(tx repository.Tx)error{lot,e:=tx.GetSupplyLot(context.Background(),"lot-water");reserved=lot.Reserved;return e});err!=nil{t.Fatal(err)};if reserved!=0{t.Fatalf("reserved remains %d",reserved)}}