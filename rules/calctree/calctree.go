package calctree

import (
	"fmt"
	"math"
)

type CalcTree interface {
	OutputNode() CalcNode
	NumberValue(string) (float64, error)
	StringValue(string) (string, error)
	SetStringInput(string, string) error
	SetNumberInput(string, float64) error
}

type CalcNode interface {
	Type() string
	Value() interface{}
	GeneratedBy() []CalcNode
	UsedBy() []CalcNode
}

type NamedNode interface {
	CalcNode
	Name() string
}

type CalcTreeBuilder interface {
	CalcGraph() CalcTree
	SetOutputNode(node CalcNode) error
	SetNumberInput(name string, value float64) error
	setStringInput(name string, value float64) error
	AddChild(parent CalcNode, child CalcNode)
}

type NodeBuilder interface {
	AsCalcNode() CalcNode
	AddUsedBy(node NodeBuilder)
	AddGeneratedBy(node NodeBuilder)
}

type BaseCalcNode struct {
	generatedBy []CalcNode
	usedBy      []CalcNode
	name        string
}

func (n *BaseCalcNode) GeneratedBy() []CalcNode {
	return n.generatedBy
}

func (n *BaseCalcNode) UsedBy() []CalcNode {
	return n.usedBy
}

func (n *BaseCalcNode) Type() string {
	return "Undefined"
}

func (n *BaseCalcNode) Name() string {
	return n.name
}

func (n *BaseCalcNode) Value() interface{} {
	return nil
}

func (n *BaseCalcNode) AddUsedBy(node NodeBuilder) {
	n.usedBy = append(n.usedBy, node.AsCalcNode())
	node.AddGeneratedBy(n)
}

func (n *BaseCalcNode) AddGeneratedBy(node NodeBuilder) {
	n.generatedBy = append(n.generatedBy, node.AsCalcNode())
}

func (n *BaseCalcNode) AsCalcNode() CalcNode {
	return n
}

type NumberValueNode struct {
	BaseCalcNode
	value float64
}

func (n *NumberValueNode) Type() string {
	return "Number"
}

func (n *NumberValueNode) Value() interface{} {
	return n.value
}

func (n *NumberValueNode) AddUsedBy(child NodeBuilder) {
	panic("AddUsedBy called on a terminal NumberValueNode")
}

func NewNumberValueNode(value float64) NodeBuilder {
	return &NumberValueNode{
		BaseCalcNode: BaseCalcNode{
			generatedBy: nil,
			usedBy:      nil,
			name:        "",
		},
		value: value,
	}
}

type StringValueNode struct {
	BaseCalcNode
	strVal string
}

func (n *StringValueNode) Type() string {
	return "String"
}

func (n *StringValueNode) Value() interface{} {
	return n.strVal
}

func (n *StringValueNode) AddUsedBy(child NodeBuilder) {
	panic("AddUsedBy called on a terminal NumberValueNode")
}

func NewStringValueNode(str string) NodeBuilder {
	return &StringValueNode{
		BaseCalcNode: BaseCalcNode{
			generatedBy: nil,
			usedBy:      nil,
			name:        "",
		},
		strVal: str,
	}
}

type SummationNode struct {
	BaseCalcNode
}

func (s *SummationNode) Value() interface{} {
	var sum float64
	for _, input := range s.generatedBy {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		sum += val
	}
	return sum
}

type MultiplicationNode struct {
	BaseCalcNode
}

func (m *MultiplicationNode) Value() interface{} {
	var product float64 = 1
	for _, input := range m.generatedBy {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		product *= val
	}
	return product
}

type DivisionNode struct {
	BaseCalcNode
}

func (d *DivisionNode) Value() interface{} {
	var quotient float64 = 1
	for _, input := range d.generatedBy {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		if val == 0 {
			return nil
		}
		quotient /= val
	}
	return quotient
}

type SubtractionNode struct {
	BaseCalcNode
}

func (s *SubtractionNode) Value() interface{} {
	var difference float64
	for _, input := range s.generatedBy {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		difference -= val
	}
	return difference
}

type ExponentiationNode struct {
	BaseCalcNode
}

func (e *ExponentiationNode) Value() interface{} {
	var result float64 = 1
	for _, input := range e.generatedBy {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		result = math.Pow(result, val)
	}
	return result
}

func (g *CTree) SetInputNode(name string, node InputNode) error {
	if g.inputNodes == nil {
		g.inputNodes = make(map[string]InputNode)
	}
	if _, exists := g.inputNodes[name]; exists {
		return fmt.Errorf("input node with name %s already exists", name)
	}
	g.inputNodes[name] = node
	return nil
}

type CTree struct {
	outputNode CalcNode
	inputNodes map[string]InputNode
}

func NewGraph() *CTree {
	return &CTree{}
}

func (g *CTree) OutputNode() CalcNode {
	return g.outputNode
}

type NumberOutputNode struct {
	BaseOperationNode
	cachedValue float64
}

func (n *NumberOutputNode) Type() string {
	return "Number"
}

func (n *NumberOutputNode) Value() interface{} {
	if !n.isCalculated {
		val, ok := n.calculatedBy.Calculate().Value().(float64)
		if !ok {

		}
		n.cachedValue = val
		n.isCalculated = true
	}
	return n.cachedValue
}

type StringOutputNode struct {
	BaseOperationNode
	cachedValue string
}

func (n *StringOutputNode) Type() string {
	return "String"
}

func (n *StringOutputNode) Value() interface{} {
	if !n.isCalculated {
		val, ok := n.calculatedBy.Calculate().Value().(string)
		if !ok {
			n.cachedValue = ""
			n.isCalculated = false
		}
		n.cachedValue = val
		n.isCalculated = true
	}
	return n.cachedValue
}

type SummationOperation struct {
	inputs []CalcNode
}

func (s *SummationNode) Calculate() BaseOperationNode {
	var sum float64
	for _, input := range s.inputs {
		val, ok := input.Value().(float64)
		if !ok {
			return nil
		}
		sum += val
	}
	var opNode OperationNode = s
	var out BaseOperationNode = &NumberOutputNode{
		BaseOperationNode: BaseOperationNode{
			isCalculated: false,
			calculatedBy: opNode,
		},
		cachedValue: sum,
	}
	return out
}

func (s *SummationNode) Inputs() []CalcNode {
	return s.inputs
}

func (g *CTree) SetInputNode(name string, node InputNode) error {
	if g.inputNodes == nil {
		g.inputNodes = make(map[string]InputNode)
	}
	if _, exists := g.inputNodes[name]; exists {
		return fmt.Errorf("input node with name %s already exists", name)
	}
	g.inputNodes[name] = node
	return nil
}
