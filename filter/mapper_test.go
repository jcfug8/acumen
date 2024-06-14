package filter_test

import (
	"testing"

	"go.einride.tech/aip/filtering"
)

type request struct {
	filter string
}

func (r request) GetFilter() string {
	return r.filter
}

func TestMapping(t *testing.T) {
	r := &request{
		filter: `
		200 = 200 AND
		a.b = 200
		`,
	}

	decl, err := filtering.NewDeclarations(
		filtering.DeclareIdent(
			"a.b",
			filtering.TypeInt,
		),
		filtering.DeclareFunction(
			filtering.FunctionEquals,
			filtering.NewFunctionOverload(
				filtering.FunctionOverloadEqualsInt,
				filtering.TypeBool,
				filtering.TypeInt,
				filtering.TypeInt,
			),
		),
		filtering.DeclareFunction(
			filtering.FunctionAnd,
			filtering.NewFunctionOverload(
				filtering.FunctionOverloadAndBool,
				filtering.TypeBool,
				filtering.TypeBool,
				filtering.TypeBool,
			),
		),
		filtering.DeclareFunction(
			filtering.FunctionFuzzyAnd,
			filtering.NewFunctionOverload(
				filtering.FunctionOverloadAndBool,
				filtering.TypeBool,
				filtering.TypeBool,
				filtering.TypeBool,
			),
		),
	)
	if err != nil {
		t.Errorf("Error creating declarations: %v", err)
	}

	parsedFilter, err := filtering.ParseFilter(r, decl)
	if err != nil {
		t.Errorf("Error parsing filter: %v", err)
	}

	if parsedFilter.CheckedExpr == nil {
		t.Errorf("Parsed filter is nil")
	}

	// walkExpr(decl, parsedFilter.CheckedExpr.Expr, nil)
}

// // func walkExpr(decl *filtering.Declarations, e *expr.Expr, db *gorm.DB) {
// // 	filtering.Walk(func(currExpr, parentExpr *expr.Expr) bool {
// // 		fmt.Print(currExpr.GetId(), " ")
// // 		switch currExpr.GetExprKind().(type) {
// // 		case *expr.Expr_ConstExpr:
// // 			switch currExpr.GetConstExpr().GetConstantKind().(type) {
// // 			case *expr.Constant_BoolValue:
// // 				fmt.Println("Bool -", currExpr.GetConstExpr().GetBoolValue())
// // 			case *expr.Constant_DoubleValue:
// // 				fmt.Println("Double -", currExpr.GetConstExpr().GetDoubleValue())
// // 			case *expr.Constant_Int64Value:
// // 				fmt.Println("Int64 -", currExpr.GetConstExpr().GetInt64Value())
// // 			case *expr.Constant_StringValue:
// // 				fmt.Println("String -", currExpr.GetConstExpr().GetStringValue())
// // 			default:
// // 				fmt.Println("Unknown Constant")
// // 			}
// // 		case *expr.Expr_IdentExpr:
// // 			fmt.Println("Ident -", currExpr.GetIdentExpr().Name)
// // 		case *expr.Expr_SelectExpr:
// // 			fmt.Println("Select -", currExpr.GetSelectExpr().GetField())
// // 		case *expr.Expr_CallExpr:
// // 			switch currExpr.GetCallExpr().GetFunction() {
// // 			case filtering.FunctionAnd:
// // 				fmt.Println("And")
// // 			case filtering.FunctionFuzzyAnd:
// // 				fmt.Println("FuzzyAnd")
// // 			case filtering.FunctionEquals:
// // 				fmt.Println("Equals")
// // 			default:
// // 			}
// // 		default:
// // 			fmt.Println("Unknown")
// // 		}
// // 		return true
// // 	}, e)
// // }

// func walkExpr(decl *filtering.Declarations, e *expr.Expr, db *gorm.DB) *gorm.DB {
//     filtering.Walk(func(currExpr, parentExpr *expr.Expr) bool {
//         switch currExpr.GetExprKind().(type) {
//         case *expr.Expr_ConstExpr:
//             // Constants are handled in the parent expression
//         case *expr.Expr_IdentExpr:
//             // Identifiers are handled in the parent expression
//         case *expr.Expr_SelectExpr:
//             // Select expressions are handled in the parent expression
//         case *expr.Expr_CallExpr:
//             call := currExpr.GetCallExpr()
//             switch call.GetFunction() {
//             case filtering.FunctionAnd:
//                 // For AND, add all arguments as conditions
//                 for _, arg := range call.GetArgs() {
//                     db = addCondition(db, arg)
//                 }
//             case filtering.FunctionOr:
//                 // For OR, add all arguments as conditions with OR
//                 for _, arg := range call.GetArgs() {
//                     db = db.Or(addCondition(db.New(), arg))
//                 }
//             case filtering.FunctionEquals:
//                 // For EQUALS, add a condition
//                 db = addCondition(db, currExpr)
//             // Add more cases for other functions as needed
//             default:
//                 fmt.Println("Unknown function")
//             }
//         default:
//             fmt.Println("Unknown expression kind")
//         }
//         return true
//     }, e)
//     return db
// }

// func addCondition(db *gorm.DB, e *expr.Expr) *gorm.DB {
//     switch e.GetExprKind().(type) {
//     case *expr.Expr_CallExpr:
//         call := e.GetCallExpr()
//         if call.GetFunction() == filtering.FunctionEquals && len(call.GetArgs()) == 2 {
//             arg1, arg2 := call.GetArgs()[0], call.GetArgs()[1]
//             // Check if both arguments are constants
//             if _, ok := arg1.GetExprKind().(*expr.Expr_ConstExpr); ok {
//                 if _, ok := arg2.GetExprKind().(*expr.Expr_ConstExpr); ok {
//                     // Both arguments are constants, compare their values
//                     return db.Where(fmt.Sprintf("? = ?", arg1.GetConstExpr().GetValue(), arg2.GetConstExpr().GetValue()))
//                 }
//             }
//             // For EQUALS, add a condition with the identifier and the constant
//             ident := arg1.GetIdentExpr()
//             constant := arg2.GetConstExpr()
//             return db.Where(fmt.Sprintf("%s = ?", ident.GetName()), constant.GetValue())
//         }
//     // Add more cases for other kinds of conditions as needed
//     default:
//         fmt.Println("Unknown condition kind")
//     }
//     return db
// }
