// DO NOT EDIT: This file is autogenerated via the builtin command.

package debug

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Comments: nil,
		Errors:   nil,
		Loc:      nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Comments: nil,
			Errors:   nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 13,
					Line:   5,
				},
				File:   "debug.flux",
				Source: "package debug\n\n// pass will pass any incoming tables directly next to the following transformation.\n// It is best used to interrupt any planner rules that rely on a specific ordering.\nbuiltin pass",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: []ast.Comment{ast.Comment{Text: "// pass will pass any incoming tables directly next to the following transformation.\n"}, ast.Comment{Text: "// It is best used to interrupt any planner rules that rely on a specific ordering.\n"}},
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   5,
					},
					File:   "debug.flux",
					Source: "builtin pass",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
			Colon: nil,
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   5,
						},
						File:   "debug.flux",
						Source: "pass",
						Start: ast.Position{
							Column: 9,
							Line:   5,
						},
					},
				},
				Name: "pass",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 54,
							Line:   5,
						},
						File:   "debug.flux",
						Source: "(<-tables: [A]) => [A] where A: Record",
						Start: ast.Position{
							Column: 16,
							Line:   5,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 54,
								Line:   5,
							},
							File:   "debug.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 45,
								Line:   5,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 54,
									Line:   5,
								},
								File:   "debug.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 48,
									Line:   5,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 46,
									Line:   5,
								},
								File:   "debug.flux",
								Source: "A",
								Start: ast.Position{
									Column: 45,
									Line:   5,
								},
							},
						},
						Name: "A",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 38,
								Line:   5,
							},
							File:   "debug.flux",
							Source: "(<-tables: [A]) => [A]",
							Start: ast.Position{
								Column: 16,
								Line:   5,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   5,
								},
								File:   "debug.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 17,
									Line:   5,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   5,
									},
									File:   "debug.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 19,
										Line:   5,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   5,
									},
									File:   "debug.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 27,
										Line:   5,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 29,
											Line:   5,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 28,
											Line:   5,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 29,
												Line:   5,
											},
											File:   "debug.flux",
											Source: "A",
											Start: ast.Position{
												Column: 28,
												Line:   5,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   5,
								},
								File:   "debug.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 35,
									Line:   5,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 37,
										Line:   5,
									},
									File:   "debug.flux",
									Source: "A",
									Start: ast.Position{
										Column: 36,
										Line:   5,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 37,
											Line:   5,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 36,
											Line:   5,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "debug.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   1,
					},
					File:   "debug.flux",
					Source: "package debug",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 14,
							Line:   1,
						},
						File:   "debug.flux",
						Source: "debug",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "debug",
			},
		},
	}},
	Package: "debug",
	Path:    "internal/debug",
}
