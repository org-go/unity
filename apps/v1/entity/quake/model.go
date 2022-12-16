package quake

import (
    "sync"
    "unity.service/apps/v1/entity/v_quake/ent"
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
