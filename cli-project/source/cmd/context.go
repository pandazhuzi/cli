package cmd

type Context struct {

}


func NewContext() (*Context,error) {

	ctx := new(Context)

	return ctx,nil
}