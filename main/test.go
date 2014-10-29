package main

import "goampl"
import "fmt"

func main() {

	model:= goampl.AMPL_init("test6.nl")
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
	fmt.Println("\nObjective Names and Sense:")
	for i:=0; i<model.Nobj; i++ {
		fmt.Printf("%v %v\n", model.Obj_name[i], model.Obj_sense[i])
	}
	fmt.Println("\nConstraint Names, Alg and Type:")
	for i:=0; i<model.Ncon; i++ {
		fmt.Println(model.Con_name[i]+"\t"+model.Con_alg[i]+"\t"+model.Con_type[i])
	}
	fmt.Println("\nVariable Names and Type:")
	for i:=0; i<model.Nvar; i++ {
		fmt.Println(model.Var_name[i]+"\t"+model.Var_type[i])
	}
	fmt.Println("\n\nVariables in Objectives")
	for i:=0; i < model.Nobj; i++ {
		fmt.Printf("\nObjective %v: ", i)
		for j:=0; j < len(model.Obj[i]); j++ {
			fmt.Printf("%v ", model.Obj[i][j])		
		}	
	}
	fmt.Println("\n\nVariables in Constraints")
	for i:=0; i < model.Ncon; i++ {
		fmt.Printf("\nConstraint %v: ", i)
		for j:=0; j < len(model.Cons[i]); j++ {
			fmt.Printf("%v ", model.Cons[i][j])		
		}	
	}
	fmt.Println("\n\nObjectives that use the Variable")
	for i:=0; i < model.Nvar; i++ {
		fmt.Printf("\nVariable %v: ", i)
		for j:=0; j < len(model.Varobj[i]); j++ {
			fmt.Printf("%v ", model.Varobj[i][j])		
		}	
	}
	fmt.Println("\n\nConstraints that use the Variable")
	for i:=0; i < model.Nvar; i++ {
		fmt.Printf("\nVariable %v: ", i)
		for j:=0; j < len(model.Varcons[i]); j++ {
			fmt.Printf("%v ", model.Varcons[i][j])		
		}	
	}
	point := make([]float64, model.Nvar)
	fmt.Println("\n\nTest Functions @ 1-point:")
	for i:=0; i<model.Nvar; i++ {
		point[i] = 1
	}
	fmt.Println("Constraint Functions")
	for i:=0; i<model.Ncon; i++ {
		conval, nerr:= goampl.Conval(model, i, point)
		fmt.Printf("Constraint %v: %v (number of errors: %v)\n", i, conval, nerr)
		fmt.Println(goampl.Congrd(model, i, point))
	}
	fmt.Println("Objective Functions")
	for i:=0; i<model.Nobj; i++ {
		objval, nerr:= goampl.Objval(model, i, point)
		fmt.Printf("Objective %v: %v (number of errors: %v)\n", i, objval, nerr)
		fmt.Println(goampl.Objgrd(model, i, point))
	}
	
}
