package ast

type ASTNode interface {
	// GetType returns the type of the AST node.
	GetType() string

	// GetChildren returns the children of the AST node.
	GetChildren() []ASTNode

	// GetParent returns the parent of the AST node.
	GetParent() *ASTNode

	// GetValue returns the value of the AST node. It is an empty
	// interface, and calling code will need to understand the node
	// type and type assertion to get the concrete value.
	GetValue() interface{}
}

type BaseASTNode struct {
	Parent *ASTNode
}

func (v BaseASTNode) GetParent() *ASTNode {
	return v.Parent
}

func (v BaseASTNode) GetValue() interface{} {
	// Not implemented by default.
	return nil
}

type ProgramNode struct {
	BaseASTNode
	Statements []ASTNode
}

func (v ProgramNode) GetType() string {
	return "ProgramNode"
}

func (v ProgramNode) GetChildren() []ASTNode {
	return v.Statements
}

type StringLiteralNode struct {
	BaseASTNode
	Value string
}

func (v StringLiteralNode) GetType() string {
	return "StringLiteralNode"
}

func (v StringLiteralNode) GetChildren() []ASTNode {
	return nil
}

func (v StringLiteralNode) GetValue() interface{} {
	return v.Value
}

type NumberLiteralNode struct {
	BaseASTNode
	Value float64
}

func (v NumberLiteralNode) GetType() string {
	return "NumberLiteralNode"
}

func (v NumberLiteralNode) GetChildren() []ASTNode {
	return nil
}

func (v NumberLiteralNode) GetValue() interface{} {
	return v.Value
}

type ComputationDefinitionNode struct {
	BaseASTNode
	Name        StringLiteralNode
	Inputs      []InputDefinitionNode
	Assignments []ASTNode
	Calculation ASTNode
}

func (v ComputationDefinitionNode) GetType() string {
	return "ComputationDefinitionNode"
}

func (v ComputationDefinitionNode) GetChildren() []ASTNode {
	children := make([]ASTNode, len(v.Inputs)+1+len(v.Assignments)+1)
	children[0] = v.Name
	for i, input := range v.Inputs {
		children[i+1] = input
	}
	for i, assignment := range v.Assignments {
		children[i+1+len(v.Inputs)] = assignment
	}
	children[len(v.Inputs)+1+len(v.Assignments)] = v.Calculation
	return children
}

// InputDefinitionNode represents an input to a computation.
type InputDefinitionNode struct {
	BaseASTNode
	Name    StringLiteralNode
	Quality StringLiteralNode

	// UnitOfMeasure is optional, and will be the empty string if not set.
	UnitOfMeasure StringLiteralNode
}

func (v InputDefinitionNode) GetType() string {
	return "InputDefinitionNode"
}

func (v InputDefinitionNode) GetChildren() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = v.Name
	children[1] = v.Quality
	if v.UnitOfMeasure.Value != "" {
		children = append(children, v.UnitOfMeasure)
	}
	return children
}

type AssignmentNode struct {
	BaseASTNode
	VariableName StringLiteralNode
	Expression   ASTNode
}

func (v AssignmentNode) GetType() string {
	return "AssignmentNode"

}

func (v AssignmentNode) GetChildren() []ASTNode {
	return []ASTNode{v.VariableName, v.Expression}
}

type GroupNode struct {
	BaseASTNode
	InnerExpression ASTNode
}

func (v GroupNode) GetType() string {
	return "GroupNode"
}

func (v GroupNode) GetChildren() []ASTNode {
	return []ASTNode{v.InnerExpression}
}

const (
	// Operators
	PlusOperator     = "+"
	MinusOperator    = "-"
	MultiplyOperator = "*"
	DivideOperator   = "/"
	PowerOperator    = "^"
)

type BinaryOperatorNode struct {
	BaseASTNode
	Left  ASTNode
	Right ASTNode
	Op    string
}

func (v BinaryOperatorNode) GetType() string {
	return "BinaryOperatorNode"
}

func (v BinaryOperatorNode) GetChildren() []ASTNode {
	var parent ASTNode = v
	var op ASTNode = StringLiteralNode{
		BaseASTNode: BaseASTNode{Parent: &parent},
		Value:       v.Op,
	}
	return []ASTNode{v.Left, op, v.Right}
}

func (v BinaryOperatorNode) GetValue() interface{} {
	return v.Op
}

type FnCallNode struct {
	BaseASTNode
	FnName StringLiteralNode
	Args   []ASTNode
}

func (v FnCallNode) GetType() string {
	return "FnCallNode"
}

func (v FnCallNode) GetChildren() []ASTNode {
	children := make([]ASTNode, len(v.Args)+1)
	children[0] = v.FnName
	for i, arg := range v.Args {
		children[i+1] = arg
	}
	return children
}

type IdentifierNode struct {
	BaseASTNode
	Name StringLiteralNode
}

func (v IdentifierNode) GetType() string {
	return "IdentifierNode"
}

func (v IdentifierNode) GetChildren() []ASTNode {
	return []ASTNode{v.Name}
}

func (v IdentifierNode) GetValue() interface{} {
	return v.Name.Value
}
