# Goampl: [Golang](https://golang.org/) interface to [AMPL](http://ampl.com/)

## Installation (Linux)

Requires [Linuxbrew](https://github.com/Homebrew/linuxbrew) for easy installation of ASL:

	brew tap homebrew/science

Must change line 28 of asl.rb from 
	
	libtool_cmd = ["ld", "-shared"]
to
	
	libtool_cmd = ["gcc", "-shared"]

Then:

	brew install asl

In goampl.go package alter CFLAGS and LDFLAGS accordingly. Then it should be able to run.

## Fields

The AMPL struct has the following fields

Field   	| Type               | Description
----------------|--------------------|------------------------------------
`Name`      	| `string          ` | Name of problem
`Nvar`      	| `Int             ` | Number of variables
`Ncon`      	| `Int             ` | Number of constraints
`Nobj`      	| `Int             ` | Number of objectives
`Con_name`      | `[]string        ` | Names of constraints
`Con_type`      | `[]string        ` | Types of constraints
`Con_alg`      	| `[]string        ` | Algebraic shapes of constraints
`RHSup`      	| `[]float64       ` | Upper RHS
`RHSlo`     	| `[]float64       ` | Lower RHS
`Obj_name`      | `[]string        ` | Name of objectives
`Obj_sense`     | `[]int           ` | Objective sense
`Obj_alg`     	| `[]string        ` | Algebraic shapes of objectives
`Var_name`      | `[]string        ` | Names of variables
`Vblo`      	| `[]float64       ` | Lower variable bounds
`Vbup`      	| `[]float64       ` | Upper variable bounds
`Var_type`      | `[]string        ` | Types of variables
`Cons`      	| `[][]string      ` | Names of variables in constraints
`Obj`      	| `[][]string      ` | Names of variables in objectives
`varcons`       | `[][]string      ` | Names of constraints for each variable
`varobj`      	| `[][]string      ` | Names of objectives for each variable
`C_ASL`      	| `*ASL_pfgh       ` | Pointer to the C ASL_pfgh struct

## Functions

Table of functions provided by the Goampl package

Function   | Parameters                              | Description
-----------|-----------------------------------------|--------------------------------------------------
`AMPL_init`| `stub string                          ` | Fills all the fields in the AMPL struct through the ASL
`Conval`   | `model AMPL, ncon int, point []float64` | Calculates LHS of constraint (ncon) at the point
`Objval`   | `model AMPL, nobj int, point []float64` | Calculates LHS of objective (nobj) at the point
`Congrd`   | `model AMPL, ncon int, point []float64` | Returns the gradient vector of constraint (ncon) at the point
`Objgrd`   | `model AMPL, nobj int, point []float64` | Returns the gradient vector of objective (nobj) at the point
