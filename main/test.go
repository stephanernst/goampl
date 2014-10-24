package main

import "goampl"
import "fmt"

func main() {

	model:= goampl.AMPL_init("dietd.nl")
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

	fmt.Println("\nVariables in Objectives")
	for i:=0; i < model.Nobj; i++ {
		fmt.Printf("\nObjective %v: ", i)
		for j:=0; j < len(model.Obj[i]); j++ {
			fmt.Printf("%v, ", model.Obj[i][j])		
		}	
	}

	fmt.Println("\nVariables in Constraints")
	for i:=0; i < model.Ncon; i++ {
		fmt.Printf("\nConstraint %v: ", i)
		for j:=0; j < len(model.Cons[i]); j++ {
			fmt.Printf("%v, ", model.Cons[i][j])		
		}	
	}

	fmt.Println("\n\nObjectives that use the Variable")
	for i:=0; i < model.Nvar; i++ {
		fmt.Printf("\nVariable %v: ", i)
		for j:=0; j < len(model.Varobj[i]); j++ {
			fmt.Printf("%v, ", model.Varobj[i][j])		
		}	
	}

	fmt.Println("\n\nConstraints that use the Variable")
	for i:=0; i < model.Nvar; i++ {
		fmt.Printf("\nVariable %v: ", i)
		for j:=0; j < len(model.Varcons[i]); j++ {
			fmt.Printf("%v, ", model.Varcons[i][j])		
		}	
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
