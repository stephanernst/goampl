package main

import "goampl"
import "fmt"

func main() {

	model:= goampl.AMPL_init("ch3.nl")
	fmt.Println("Problem name:")
	fmt.Println(model.Name)
	fmt.Println("\nNumber of variables:")
	fmt.Println(model.Nvar)
	fmt.Println("\nNumber of constraints:")
	fmt.Println(model.Ncon)
	fmt.Println("\nNumber of objectives:")
	fmt.Println(model.Nobj)
	fmt.Println("\nUpper variable bounds:")
	fmt.Println(model.Vbup)
	fmt.Println("\nLower variable bounds:")
	fmt.Println(model.Vblo)
	fmt.Println("\nUpper RHS:")
	fmt.Println(model.RHSup)
	fmt.Println("\nLower RHS:")
	fmt.Println(model.RHSlo)
	fmt.Println("\nObjective Names:")
	for i:=0; i<model.Nobj; i++ {
		fmt.Println(model.Obj_name[i])
	}
	fmt.Println("\nObjective Sense:")
	for i:=0; i<model.Nobj; i++ {
		fmt.Println(model.Obj_sense[i])
	}
	fmt.Println("\nConstraint Names:")
	for k:=0; k<model.Ncon; k++ {
		fmt.Println(model.Con_name[k])
	}
	fmt.Println("\nVariable Names:")
	for j:=0; j<model.Nvar; j++ {
		fmt.Println(model.Var_name[j])
	}
	point := make([]float64, model.Nvar)
	fmt.Println("\nTest Functions:")
	for i:=0; i<model.Nvar; i++ {
		point[i] = 2
	}
	for i:=0; i<model.Ncon; i++ {
		fmt.Println(goampl.Conval(model, i, point))
	}
	
	fmt.Println(goampl.Objval(model, 0, point))
	fmt.Println(goampl.Objgrd(model, 0, point))
}
