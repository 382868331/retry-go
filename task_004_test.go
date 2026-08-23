package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry004LastErrorFlag(t *testing.T){ got:=New(LastErrorOnly(true)).lastErrorOnly; if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry004LastErrorFlagAdjacent(t *testing.T){
 got:=New(LastErrorOnly(false)).lastErrorOnly
 if got != false{ t.Fatalf("adjacent got %v want %v",got,false) }
}
