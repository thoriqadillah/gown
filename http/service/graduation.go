package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"sync"

	"github.com/thoriqadillah/gown/http/api"
	"github.com/thoriqadillah/gown/worker"
)

type Graduation struct {
	jobid int
	api   api.GraduationAPI
	year  string
	wg    *sync.WaitGroup
	mutex *sync.Mutex
}

func NewGraduation(api api.GraduationAPI, year string, wg *sync.WaitGroup) worker.Job {
	return &Graduation{
		jobid: rand.Intn(100),
		api:   api,
		year:  year,
		wg:    wg,
	}
}

func (g *Graduation) Execute() error {
	defer g.wg.Done()

	ext := ".csv"
	grads := g.api.GetGraduees(g.year)
	file, err := os.Create("./storage/" + g.year + ext)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(file)
	forfield := true

	rows := make([][]string, len(grads.Result.Records))
	lencol := reflect.TypeOf(grads.Result.Records[0]).NumField()
	fields := make([]string, lencol)
	for i, record := range grads.Result.Records {
		rows[i] = make([]string, lencol)

		val := reflect.Indirect(reflect.ValueOf(record))
		for j := 0; j < lencol; j++ {
			rows[i][j] = fmt.Sprint(val.Field(j).Interface())
			fields[j] = val.Type().Field(j).Name
		}
		if forfield {
			csvwriter.Write(fields)
			forfield = false
		}

		csvwriter.Write(rows[i])
	}

	csvwriter.Flush()

	return nil
}

func (g *Graduation) HandleError(err error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	log.Printf("Error executing job %d\n", g.jobid)
}
