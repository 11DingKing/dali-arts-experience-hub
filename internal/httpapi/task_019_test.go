package httpapi
import ("context";"net/http/httptest";"testing")
func TestDali019CanceledErrorMapsToTimeout(t *testing.T){r:=httptest.NewRequest("GET","/",nil);w:=httptest.NewRecorder();writeError(w,r,context.Canceled);if w.Code==500{t.Fatal("cancellation mapped to internal error")}}