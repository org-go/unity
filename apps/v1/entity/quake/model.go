package quake

import (
    "entgo.io/ent/examples/start/ent"
    "sync"

)

var do sync.Once

type (
    IschemaInterface interface {
        Link() (client *ent.Client)
    }

    model struct {
        *ent.Client
    }
)

func (this *model) run(client *ent.Client) *model {

    do.Do(func() {
        this.Client = client
    })
    return this
}
