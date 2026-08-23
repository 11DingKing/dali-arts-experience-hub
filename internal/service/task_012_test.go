package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali012CanceledLogoutDoesNotRevoke(t *testing.T){e:=newTestEnvironment(t);r,err:=e.service.Login(context.Background(),"duty","duty-secret");if err!=nil{t.Fatal(err)};ctx,cancel:=context.WithCancel(context.Background());cancel();if err=e.service.Logout(ctx,r.SessionID,e.duty);err==nil{t.Fatal("canceled logout succeeded")};if _,err=e.service.Authenticate(context.Background(),r.Token);err!=nil{t.Fatalf("session was revoked: %v",err)}}