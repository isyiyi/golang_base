package proDataType

import "text/template"
import "os"
import "fmt"

const tmpl = `count = {{.Count}}
	{{range .Data}}
	name: {{.Name}}
	age: {{.Age}}
	---------------------
	{{end}}`

type student2 struct {
	Name string
	Age int
}

type result struct {
	Count int
	Data [] student2
}

func MainProcess2() {
	var stu1 = student2{"kobe", 34}
	var stu2 = student2{"james", 36}
	var stu3 = student2{"jordan", 40}
	var data = make([]student2, 3)
	data[0] = stu1
	data[1] = stu2
	data[2] = stu3
	var res = result{Count: len(data), Data:data}
	
	var report = template.Must(template.New("test").Parse(tmpl))
	if err := report.Execute(os.Stdout, res); err != nil {
		fmt.Println(err)
	}

}
