package eval

import (
	"errors"
	"github.com/suzuken/gs/types"
)

// Eval is body of evaluator
func Eval(exp types.Expression, env *types.Env) (types.Expression, error) {
	switch t := exp.(type) {
	case types.Boolean:
		return exp, nil
	case types.Number:
		return exp, nil
	case types.Symbol:
		// it's variable. get value from environment
		e, err := env.Get(t)
		if err != nil {
			return nil, err
		}
		return e, nil
	case []types.Expression:
		// this is multiple expressions pattern
		// at first, get car. car of expression is symbol for each expression
		car, ok := t[0].(types.Symbol)
		if !ok {
			return nil, errors.New("cannot conversion car of expressions. it should be types.Symbol but not.")
		}
		switch car {
		case "define":
		case "if":
		case "cond":
		case "lambda":
		case "begin":
		default:
		}
	default:
		// not found any known operands. failed.
		return nil, errors.New("unkonwn expression type")
	}
	return nil, nil
}

// Apply receives procedure and arguments. if procedure is compounded, evaluate on extended environment.
func Apply(p *types.Expression, args []types.Expression) (*types.Expression, error) {
	if primitiveProcedure(p) {
		ApplyPrimitiveProcedure(p, args)
	} else if compoundProcedure(p) {
		EvalSequence(ProcedureBody(p), ExtendEnvironment(ProcedureParameters(p), args, ProcedureEnvironment(p)))
	} else {
		return nil, errors.New("Unknown procedure type")
	}
	return nil, nil
}

func ProcedureBody(p types.Expression) (exps []types.Expression) {
	return exps
}

func ProcedureParameters(p types.Expression) (exps []types.Expression) {
	return exps
}

func ProcedureEnvironment(p types.Expression) *types.Env {
	return nil
}

func ExtendEnvironment(exps, args []types.Expression, env *types.Env) *types.Env {
	return nil
}

// EvalSeauencd evaluate sequence of expressions in certain environment.
// Return is last expression.
func EvalSequence(exps []types.Expression, env *types.Env) (types.Expression, error) {
	if len(exps) == 1 {
		return Eval(exps[0], env)
	}
	// making environment (Yes, it's pointer)
	if _, err := Eval(exps[0], env); err != nil {
		return nil, err
	}
	return EvalSequence(exps[1:], env)
}

// listOfValues returns arguments for evaluator.
func listOfValues(exps []types.Expression, env *types.Env) (types.Expression, error) {
	if len(exps) <= 0 {
		return nil, nil
	}
	// evaluate exps one by one on each environment
	first, err := Eval(exps[0], env)
	if err != nil {
		return nil, err
	}
	// TODO: should use for in Go way?
	rest, err := listOfValues(exps[1:], env)
	if err != nil {
		return nil, err
	}
	return types.Pair{first, rest}, nil
}

func primitiveProcedure(s *types.Expression) bool {
	return false
}

func compoundProcedure(s *types.Expression) bool {
	return false
}

func ApplyPrimitiveProcedure(p *types.Expression, args []types.Expression) (*types.Expression, error) {
	return nil, nil
}