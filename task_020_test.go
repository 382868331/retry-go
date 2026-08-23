package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry020WrapContextFlag(t *testing.T){ got:=New(WrapContextErrorWithLastError(true)).wrapContextErrorWithLastError; if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry020WrapContextFlagAdjacent(t *testing.T){
 got:=New(WrapContextErrorWithLastError(false)).wrapContextErrorWithLastError
 if got != false{ t.Fatalf("adjacent got %v want %v",got,false) }
}
