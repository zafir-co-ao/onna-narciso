package main

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/tests/stubs"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	r := inmem.NewAppointmentRepository()
	cacl := stubs.CustomerAclStub{}
	pacl := stubs.Pacl
	sacl := stubs.Sacl

	s := scheduling.NewAppointmentScheduler(r, cacl, pacl, sacl)
	c := scheduling.NewAppointmentCanceler(r)

	http.Handle("/", web.NewRouter(s, c))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
