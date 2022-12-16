package quake

import (
    _ "github.com/go-sql-driver/mysql"
)


type (
    //  iOption
    iOption interface {
        Link(dns string) error
    }
    //  Option
    Option func(opts ...*options) iOption
    //  options
    options struct {

    }
    //  Repo
    Repo struct {
        M *model
        o *options
    }
)

/**
 * $(NewRepo)
 * @Description:
 * @param opts
 * @return *Repo
 */
func NewRepo(opts ...Option) *Repo {

    opt := &options{}
    for _, o := range opts {
        o(opt)
    }
    return &Repo{
        o: opt,
        M: new(model),
    }
}

/**
 * $(Call)
 * @Description:
 * @receiver this
 * @param schema
 * @return *Repo
 */
func (this *Repo) Call(schema IschemaInterface) *Repo {

    this.M.run(schema.Link())
    return this

}
