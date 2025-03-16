package evaluator

import (
	"fmt"
	"orangutan/object"
)

var builtins = map[string]*object.Builtin{
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be an array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements
			if len(arr) > 0 {
				return arr[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be an array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements
			length := len(arr)
			if length > 0 {
				return arr[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be an array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements
			length := len(arr)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be an array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements
			origLength := len(arr)

			newElements := make([]object.Object, origLength+1)
			copy(newElements, arr)
			newElements[origLength] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
}
