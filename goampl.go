package goampl

/*#cgo CFLAGS: -I/home/stephane/.linuxbrew/Cellar/asl/20140723/include/asl 
#cgo LDFLAGS: -L/home/stephane/.linuxbrew/Cellar/asl/20140723/lib -lasl -lm -ldl
#include <asl_pfgh.h>

//ASL struct

ASL_pfgh *asl_init(char *stub) {
	FILE *nl;
	ASL_pfgh *asl;

	asl = (ASL_pfgh *)ASL_alloc(ASL_read_pfgh);
	nl = jac0dim(stub, (fint)strlen(stub));
	pfgh_read(nl,0);

	return asl;
}

int asl_nobj(ASL_pfgh *asl) {
	return asl->i.n_obj_;
}
int asl_nvar(ASL_pfgh *asl) {
	return asl->i.n_var_;
}
int asl_ncon(ASL_pfgh *asl) {
	return asl->i.n_con_;
}
int asl_nlvc(ASL_pfgh *asl) {
	return asl->i.nlvc_;
}
int asl_nlvo(ASL_pfgh *asl) {
	return asl->i.nlvo_;
}
int asl_nwv(ASL_pfgh *asl) {
	return asl->i.nwv_;
}
int asl_nlvar(ASL_pfgh *asl) {
	if (asl->i.nlvc_ > asl->i.nlvo_)
		return 	asl->i.nlvc_;
	else return asl->i.nlvo_;
}
int asl_nbv(ASL_pfgh *asl) {
	return asl->i.nbv_;
}
int asl_niv(ASL_pfgh *asl) {
	return asl->i.niv_;
}
int asl_otherlinear(ASL_pfgh *asl) {
	return asl_nvar(asl) - (asl_nlvar(asl) + asl_niv(asl) + asl_nbv(asl) + asl_nwv(asl));
}
int asl_nlnc(ASL_pfgh *asl) {
	return asl->i.nlnc_;
}
int asl_lnc(ASL_pfgh *asl) {
	return asl->i.lnc_;
}
int asl_nlc(ASL_pfgh *asl) {
	return asl->i.nlc_;
}
int asl_lgeneral(ASL_pfgh *asl) {
	return asl_ncon(asl) - (asl_nlc(asl) + asl_lnc(asl));
}
int asl_nlgeneral(ASL_pfgh *asl) {
	return asl_nlc(asl) - asl_nlnc(asl);
}
char *asl_objtype(ASL_pfgh *asl) {
	return asl->i.objtype_;
}
real *asl_LUrhs(ASL_pfgh *asl) {
	return asl->i.LUrhs_;
}
real *asl_LUv(ASL_pfgh *asl) {
	return asl->i.LUv_;
}
cgrad *asl_Cgrad(ASL_pfgh *asl, int row) {
	return asl->i.Cgrad_[row];
}
ograd *asl_Ograd(ASL_pfgh *asl, int row) {
	return asl->i.Ograd_[row];
}
real cconival(ASL_pfgh *asl, int ncon, real *X, fint *nerror) {
	return (*((ASL*)asl)->p.Conival)((ASL*)asl,ncon,X,nerror);
}
real cobjval(ASL_pfgh *asl, int np, real *X, fint *nerror) {
	return (*((ASL*)asl)->p.Objval)((ASL*)asl,np,X,nerror);
}
void ccongrad(ASL_pfgh *asl, int i, real *X, real *G, fint *nerror) {
	return (*((ASL*)asl)->p.Congrd)((ASL*)asl,i,X,G,nerror);
}
void cobjgrad(ASL_pfgh *asl, int np, real *X, real *G, fint *nerror) {
	return (*((ASL*)asl)->p.Objgrd)((ASL*)asl,np,X,G,nerror);
}
*/
import "C"
import "unsafe"
import "math"
import "reflect"

//stores needed data from the ASL struct
type AMPL struct {
	Name      string	//Problem name
	Nvar      int		//Number of variables
	Ncon      int		//Number of constraints
	Nobj      int		//Number of objectives
	Con_name  []string	//Constraints names
	Con_type  []string	//Constraint type
	Con_alg   []string	//Constraint algebraic shape
	RHSup     []float64	//Upper Right Hand Side
	RHSlo     []float64	//Lower Right Hand Side
	
	Obj_name  []string	//Objective Name
	Obj_sense []int		//Objective Sense (min = 0, max = 1)
	Obj_alg   []string	//Objective algebraic shape

	Var_name  []string	//Variable names
	Vblo      []float64	//Lower variable bounds
	Vbup      []float64	//Upper variable bounds
	Var_type  []string	//Variable type

	Cons      [][]string	//Names of variables in constraints
	Obj       [][]string	//Names of variables in objectives

	Varcons   [][]string	//Names of constraints for each variable
	Varobj    [][]string	//Names of objectives for each variable

	C_ASL     *C.struct_ASL_pfgh
}

/*	Initializes all the fields of the AMPL struct
*	Returns an AMPL struct
*/
func AMPL_init(stub string) AMPL{
	var model AMPL
	var asl *C.struct_ASL_pfgh
	stubc := C.CString(stub)
	defer C.free(unsafe.Pointer(stubc))

	//Setting up AMPL struct
	asl = C.asl_init(stubc)	
	model.Name = stub
	model.Nvar = int(C.asl_nvar(asl))
	model.Ncon = int(C.asl_ncon(asl))
	model.Nobj = int(C.asl_nobj(asl))

	model.Con_name = make([]string, model.Ncon)	
	model.Con_type = make([]string, model.Ncon)
	model.Con_alg = make([]string, model.Ncon)
	model.RHSup = make([]float64, model.Ncon)
	model.RHSlo = make([]float64, model.Ncon)
	
	model.Obj_name = make([]string, model.Nobj)
	model.Obj_sense = make([]int, model.Nobj)
	model.Obj_alg = make([]string, model.Nobj)

	model.Var_name = make([]string, model.Nvar)
	model.Vblo = make([]float64, model.Nvar)
	model.Vbup = make([]float64, model.Nvar)
	model.Var_type = make([]string, model.Nvar)

	model.Cons = make([][]string, model.Ncon)
	model.Obj = make([][]string, model.Nobj)

	model.Varcons = make([][]string, model.Nvar)
	model.Varobj = make([][]string, model.Nvar)

	model.C_ASL = asl

	//Filling in Obj_sense, Vblo, Vbup, RHSlo, and RHSup
	objtype := unsafe.Pointer(C.asl_objtype(asl))
	LUv     := unsafe.Pointer(C.asl_LUv(asl))
	LUrhs   := unsafe.Pointer(C.asl_LUrhs(asl))
	
	obj := *(*[]byte)   (arrayToSlice(objtype, model.Nobj))
	vb  := *(*[]float64)(arrayToSlice(LUv, model.Nvar*2))
	rhs := *(*[]float64)(arrayToSlice(LUrhs, model.Ncon*2))
	
	for i:=0; i < model.Nobj; i++ {
		model.Obj_sense[i] = int(obj[i])	
	}
	for i:=0; i < model.Nvar; i++ {
		model.Vblo[i] = vb[2*i]
		model.Vbup[i] = vb[2*i+1]
	}
	for i:=0; i < model.Ncon; i++ {
		model.RHSlo[i] = rhs[2*i]
		model.RHSup[i] = rhs[2*i+1]
	}
	
	//Filling in constraint type
	for i:=0; i < model.Ncon; i++ {
		if math.IsInf(model.RHSup[i], 1) && !math.IsInf(model.RHSlo[i], 0) {
			model.Con_type[i] = "G"
		} else if !math.IsInf(model.RHSup[i], 1) && math.IsInf(model.RHSlo[i], 0){
			model.Con_type[i] = "L"
		} else if model.RHSup[i] == model.RHSlo[i] {
			model.Con_type[i] = "E"		
		} else if !math.IsInf(model.RHSup[i], 1) && !math.IsInf(model.RHSlo[i], 0){
			model.Con_type[i] = "R"
		} else {
			model.Con_type[i] = "N"		
		}
		
	}

	//Filling in Con_name, Obj_name, and Var_name
	for i:=0; i < model.Ncon; i++ {
		conname := C.con_name_ASL((*C.struct_ASL)(unsafe.Pointer(asl)),C.int(i))
		model.Con_name[i] = C.GoString(conname)
	}
	for i:=0; i < model.Nobj; i++ {
		objname := C.obj_name_ASL((*C.struct_ASL)(unsafe.Pointer(asl)),C.int(i))
		model.Obj_name[i] = C.GoString(objname)
	}
	for i:=0; i < model.Nvar; i++ {
		varname := C.var_name_ASL((*C.struct_ASL)(unsafe.Pointer(asl)),C.int(i))
		model.Var_name[i] = C.GoString(varname)
	}

	//Filling in Variable type (Var_type)
	k:=0
	for i:= 0; i < int(C.asl_nlvar(asl)); i++ {
		model.Var_type[k] = "NL"
		k++	}
	for i:= 0; i < int(C.asl_nwv(asl)); i++ {
		model.Var_type[k] = "LA"	
		k++	}
	for i:= 0; i < int(C.asl_otherlinear(asl)); i++ {
		model.Var_type[k] = "OL"	
		k++	}
	for i:= 0; i < int(C.asl_nbv(asl)); i++ {
		model.Var_type[k] = "B"	
		k++	}
	for i:= 0; i < int(C.asl_niv(asl)); i++ {
		model.Var_type[k] = "I"	
		k++	}

	//Filling in Constraint Algebraic Shape (Con_alg)
	j:=0
	for i:=0; i < int(C.asl_nlgeneral(asl)); i++ {
		model.Con_alg[j] = "NLG"
		j++	}
	for i:=0; i < int(C.asl_nlnc(asl)); i++ {
		model.Con_alg[j] = "NLN"
		j++	}
	for i:=0; i < int(C.asl_lnc(asl)); i++ {
		model.Con_alg[j] = "LN"
		j++	}
	for i:=0; i < int(C.asl_lgeneral(asl)); i++ {
		model.Con_alg[j] = "LG"
		j++	}
	
	//Filling in Cons and Varcons
	for i:=0; i < model.Ncon; i++ {
		cgrad:= C.asl_Cgrad(asl, C.int(i))
		for cgrad != nil {
			varno:= int(cgrad.varno)
			model.Cons[i] = append(model.Cons[i], model.Var_name[varno])
			if !contains(model.Varcons[varno], model.Con_name[i]) {
				model.Varcons[varno] = append(model.Varcons[varno], model.Con_name[i])			
			}	
			cgrad = cgrad.next	
		}
	}

	//Filling in Obj and Varobj
	for i:=0; i < model.Nobj; i++ {
		ograd:= C.asl_Ograd(asl, C.int(i))
		for ograd != nil {
			varno:= int(ograd.varno)
			model.Obj[i] = append(model.Obj[i], model.Var_name[varno])
			if !contains(model.Varobj[varno], model.Obj_name[i]) {
				model.Varobj[varno] = append(model.Varobj[varno], model.Obj_name[i])			
			}
			ograd = ograd.next	
		}
	}

	return model
}

/*
*	Calculates the LHS of constraint (ncon) at the given point
*	Returns the value of the LHS and number of errors
*/
func Conval(model AMPL, ncon int, point []float64) (result float64, nerror int){
	var ne C.fint
	cpoint := (*C.double)(unsafe.Pointer(&point[0]))
	return float64(C.cconival(model.C_ASL, C.int(ncon), cpoint, &ne)), int(ne)
}

/*
*	Calculates the LHS of objective (nobj) at the given point
*	Returns the value of the LHS and number of errors
*/
func Objval(model AMPL, nobj int, point []float64) (result float64, nerror int){
	var ne C.fint
	cpoint := (*C.double)(unsafe.Pointer(&point[0]))
	return float64(C.cobjval(model.C_ASL, C.int(nobj), cpoint, &ne)), int(ne)
}

/*
*	Returns the gradient vector of constraint (ncon) at the given point and number of errors
*/
func Congrd(model AMPL, ncon int, point []float64) (gradvec []float64, nerror int){
	var ne C.fint
	grad := make([]float64, model.Nvar)
	ptr := (*C.double)(unsafe.Pointer(&grad[0]))
	cpoint := (*C.double)(unsafe.Pointer(&point[0]))
	C.ccongrad(model.C_ASL, C.int(ncon), cpoint, ptr, &ne)
	
	return grad, int(ne)
}

/*
*	Returns the gradient vector of objective (nobj) at the given point and number of errors
*/
func Objgrd(model AMPL, nobj int, point []float64) (gradvec []float64, nerror int){
	var ne C.fint
	grad := make([]float64, model.Nvar)
	ptr := (*C.double)(unsafe.Pointer(&grad[0]))
	cpoint := (*C.double)(unsafe.Pointer(&point[0]))
	C.cobjgrad(model.C_ASL, C.int(nobj), cpoint, ptr, &ne)

	return grad, int(ne)
}

/*------------------------------------------------------------------------------------------
*	Functions used for AMPL_init()
*/
func arrayToSlice(array unsafe.Pointer, length int) unsafe.Pointer {
	slice := reflect.SliceHeader{
		Data: uintptr(array),
		Len: length,
		Cap: length,
	}
	return unsafe.Pointer(&slice)
}

func contains(list []string, s string) bool{
	for i:=0; i < len(list); i++ {
		if list[i] == s {
			return true
		}
	}
	return false
}
