package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/utils"
)

const queryEntsCount = 4
// todo: justify
const queryTimeLimit = time.Microsecond * queryEntsCount

func TestCreateEntity(t *testing.T) {
    w := ecs.GetWorld()

    entsBefore := len(w.Entities)

    w.NewEntity(ecs.Type{})

    if len(w.Entities) - entsBefore != 1 {
        t.Error("Entity wasn't created")
    }
}

func TestCreateEntityWithType(t *testing.T) {
    w := ecs.GetWorld()

    rec := w.NewEntity(ecstype)

    arch := rec.Archetype

    compsBefore := len(arch.Components)
    colBefore := len(arch.Components[0])

    w.NewEntity(ecstype)

    if len(arch.Components) != compsBefore {
        t.Error("Adding new entity of the same type added new Column to archetype!")
    }

    if len(arch.Components[0]) - colBefore != 1 {
        t.Error("Components for new entity weren't inserted")
    }
}

func TestSetComponent(t *testing.T) {
    w := ecs.GetWorld()

    rec := w.NewEntity(ecstype)

    w.SetEntityComponent(rec.Entity, dummyComp1{ Val: 69420 })

    comps := rec.Archetype.Components[ecs.GetComponentIndexInType(ecstype, dummyComp1CompId)]
    val := comps[rec.Index]

    dcomp, ok := val.(dummyComp1)

    if !ok {
        t.Error("Got component of wrong type")
    }

    if dcomp.Val != 69420 {
        t.Error("Component's value invalid")
    }
}

func TestQuery(t *testing.T) {
    w := ecs.GetWorld()

    for range queryEntsCount {
        w.NewEntity(ecs.Type{
            queryTestCompCompId,
        })
    }

    ents, dur := utils.MeasureTime(func() []ecs.Entity {
        return w.Query(queryTestCompCompId)
    })

    if dur > queryTimeLimit {
        fmt.Println("Query took too long")
        t.Fail()
    }

    if len(ents) != queryEntsCount {
        t.Errorf(
            "Query returned wrong amount of entities. Expected: %d, got: %d",
            queryEntsCount,
            len(ents),
        )
    }
}

type dummyComp1 struct {
    Val int
}
var dummyComp1CompId = ecs.GetWorld().GetNewId()
func (dummyComp1) ComponentId() ecs.ComponentId {
    return dummyComp1CompId
}

type dummyComp2 struct {
    Val int
}
var dummyComp2CompId = ecs.GetWorld().GetNewId()
func (dummyComp2) ComponentId() ecs.ComponentId {
    return dummyComp1CompId
}

type queryTestComp struct {
    Val string
}
var queryTestCompCompId = ecs.GetWorld().GetNewId()
func (queryTestComp) ComponentId() ecs.ComponentId {
    return queryTestCompCompId
}


var ecstype = ecs.NewType(dummyComp1CompId, dummyComp2CompId);
