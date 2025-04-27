package ecs

import (
	"slices"
	"strconv"

	"github.com/cxncxl/gogame/internal/utils"
)

type rawId uint32

type Id struct {
    raw rawId
}
func (self Id) String() string {
    return strconv.Itoa(int(self.raw))
}

type Entity = Id
type ComponentId = Id

type Component interface {
    ComponentId() ComponentId
}

// Each entity can have several components on it
// Those components will define entities type
type Type []ComponentId
// Used for Type as map's key
type typeKey string
var BaseType = Type{}

func NewType(components ...ComponentId) Type {
    return utils.Map(
        utils.QuickSort(
            utils.Map(components, func(v ComponentId, _ int) rawId {
                return v.raw
            }),
        ),
        func (v rawId, _ int) ComponentId {
            return Id{
                raw: v,
            }
        },
    )
}

// Convert Type to typeKey
func (self Type) Key() typeKey {
    res := ""

    utils.ForEach(self, func(v ComponentId, _ int) {
        res += v.String() + ":"
    })

    return typeKey(res)
}

// TODO: implement
func (self Type) EmptyComponents() Column {
    res := Column{}

    for _, id := range self {
        res = append(res, emptyComponents[id])
    }

    return res
}

func (self Type) IncludesAll(comps ...ComponentId) bool {
    for _, c := range comps {
        found := false
        for _, v := range self {
            if v.raw == c.raw {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }

    return true
}

// List of implementations of a Component
// E.g. we have 3 entities with Position component (id = 2)
// Positions []Column = { PositionComponent{}, PositionComponent{}, PositionComponent{} }
type Column = []Component

type ArchetypeId = Id
type Archetype = struct {
    Id ArchetypeId
    Type Type
    Components []Column
    // Graph for archetypes
    // i.e. PositionAndRotationArchetype.Edges[PositionComponentId].Remove = RotationArchetype
    // PositionArchetype.Edges[RotationComponentId].Add = PositionAndRotationArchetype
    // [1, 2].Edges.Add[5] == [1, 2, 5]
    // [1, 5, 8].Edges.Remove[5] == [1, 8]
    Edges map[ComponentId]*ArchetypeEdge
    // Components[i] is related to Entity[i]
    Entities []Entity
}

type ArchetypeEdge struct {
    Add *Archetype
    Remove *Archetype
}

type EntityRecord struct {
    Entity Entity
    Index int
    Archetype *Archetype
}

type World struct {
    initialized bool
    lastUsedId rawId
    Entities map[Entity]EntityRecord
    Archetypes map[typeKey]*Archetype
    allArchetypes []*Archetype
}
var world World

func GetWorld() *World {
    if world.initialized == false {
        world = *newWorld()
    }

    return &world
}

func newWorld() *World {
    return &World{
        initialized: true,
        lastUsedId: 0,
        Entities: make(map[Entity]EntityRecord),
        Archetypes: make(map[typeKey]*Archetype),
    }
}

func (self *World) GetNewId() Id {
    self.lastUsedId += 1

    return Id{
        raw: self.lastUsedId,
    }
}

func (self *World) GetArchetypeByType(t Type) *Archetype {
    for _, v := range self.Archetypes {
        if v.Type.Key() == t.Key() {
            return v
        }
    }

    return nil
}

func (self *World) NewArchetype(t Type) *Archetype {
    arch := Archetype{
        Id: self.GetNewId(),
        Type: t,
        Entities: make([]Entity, 0),
        Components: make([]Column, 0),
        Edges: make(map[ComponentId]*ArchetypeEdge, 1),
    }

    for i, v := range t {
        typeWithout := make(Type, 0, len(t) - 1)
        typeWithout = append(typeWithout, t[:1]...)
        typeWithout = append(typeWithout, t[i+1:]...)

        if without := self.GetArchetypeByType(typeWithout); without != nil {
            edge := ArchetypeEdge{
                Remove: without,
            }

            arch.Edges[v] = &edge

            exstEdge, ok := without.Edges[v]

            if ok == false {
                exstEdge = &ArchetypeEdge{}
                without.Edges[v] = exstEdge
            }

            exstEdge.Add = &arch
        }
    }

    self.Archetypes[t.Key()] = &arch
    self.allArchetypes = append(self.allArchetypes, &arch)

    return &arch
}

func (self *World) NewEntity(t Type) EntityRecord {
    if t == nil {
        t = BaseType
    }

    arch := self.GetArchetypeByType(t)
    if arch == nil {
        arch = self.NewArchetype(t)
    }

    ent := self.GetNewId()

    arch.Entities = append(arch.Entities, ent)

    comps := t.EmptyComponents()
    if len(arch.Components) == 0 {
        arch.Components = make([]Column, len(comps))
    }
    for i, comp := range comps {
        arch.Components[i] = append(arch.Components[i], comp)
    }

    rec := EntityRecord{
        Entity: ent,
        Archetype: arch,
        Index: len(arch.Entities) - 1,
    }

    self.Entities[ent] = rec

    return rec
}

func (self *World) getEntityRecord(ent Entity) (*EntityRecord, error) {
    rec, ok := self.Entities[ent]
    if ok == false {
        return nil, EntityNotFoundError
    }

    return &rec, nil
}

func (self *World) GetEntityComponents(ent Entity) (map[ComponentId]Component, error) {
    rec, err := self.getEntityRecord(ent)
    if err != nil {
        return nil, err
    }

    arch := rec.Archetype
    if arch == nil {
        return nil, EntityNotFoundError
    }

    res := make(map[ComponentId]Component, 1)

    for i, compId := range arch.Type {
        res[compId] = arch.Components[i][rec.Index]
    }

    return res, nil
}

func (self *World) SetEntityComponent(ent Entity, comp Component) error {
    rec, err := self.getEntityRecord(ent)
    if err != nil {
        return err
    }

    arch := rec.Archetype
    compIdx := getComponentIndexInType(arch.Type, comp.ComponentId())
    if compIdx < 0 {
        newType := append(arch.Type, comp.ComponentId())
        newArch, ok := self.Archetypes[newType.Key()]
        if ok == false {
            newArch = self.NewArchetype(newType)
        }

        self.migrateEntity(ent, newArch)
        rec, _ = self.getEntityRecord(ent)
        arch = rec.Archetype
        compIdx = getComponentIndexInType(arch.Type, comp.ComponentId())
    }

    arch.Components[compIdx][rec.Index] = comp
    
    return nil
}

func (self *World) Query(comps ...ComponentId) []Entity {
    totalEnts := 0
    filtered := utils.Filter(
        self.allArchetypes,
        func(v *Archetype, _ int) bool {
            if v.Type.IncludesAll(comps...) {
                totalEnts += len(v.Entities)
                return true
            }

            return false
        },
    )

    res := make([]Entity, 0, totalEnts)

    for _, v := range filtered {
        res = append(res, v.Entities...)
    }

    return res
}

func (self *World) migrateEntity(ent Entity, to *Archetype) error {
    rec, err := self.getEntityRecord(ent)
    if err != nil {
        return err
    }

    from := rec.Archetype
    eidx := rec.Index

    to.Entities = append(to.Entities, ent)
    rec.Archetype = to
    rec.Index = len(to.Entities) - 1

    if len(to.Components) == 0 {
        to.Components = make([]Column, len(to.Type))
    }

    for _, v := range from.Type {
        tidx := getComponentIndexInType(to.Type, v)
        fidx := getComponentIndexInType(from.Type, v)
        to.Components[tidx] = append(to.Components[tidx], from.Components[fidx][eidx])
    }

    from.Entities = slices.Delete(
        from.Entities,
        eidx,
        eidx + 1,
    )
    
    for i := range from.Components {
        from.Components[i] = slices.Delete(
            from.Components[i],
            eidx,
            eidx + 1,
        )
    }

    self.Entities[ent] = *rec

    return nil
}

func getComponentIndexInType(t Type, comp ComponentId) int {
    for i, v := range t {
        if v.raw == comp.raw {
            return i
        }
    }

    return -1
}

const maxEntityCount int = 4096

type EntityNotFound struct {}
func (EntityNotFound) Error() string {
    return "Entity not registered in this world";
}
var EntityNotFoundError = EntityNotFound {}
