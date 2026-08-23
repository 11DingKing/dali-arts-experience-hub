package service
import (
 "context"; "testing"; "time"
 "github.com/11DingKing/dali-arts-experience-hub/internal/domain"
 "github.com/11DingKing/dali-arts-experience-hub/internal/repository"
)
var _ = time.Time{}
var _ repository.Repository
func TestDali011ExpiredSessionRejected(t *testing.T){e:=newTestEnvironment(t);c:=New(e.repo,func()time.Time{return e.now},-time.Minute);r,err:=c.Login(context.Background(),"duty","duty-secret");if err!=nil{t.Fatal(err)};if _,err=c.Authenticate(context.Background(),r.Token);err==nil{t.Fatal("expired session accepted")}}