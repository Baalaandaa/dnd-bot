package solver

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const delimiters = "+-dк"

func BuildAST(data string) (*Node, error) {
	exists := strings.ContainsAny(data, delimiters)
	for _, char := range data {
		if char >= '0' && char <= '9' {
			continue
		} else if strings.Contains(delimiters, fmt.Sprintf("%c", char)) {
			continue
		} else {
			return nil, errors.New("invalid string given")
		}
	}
	if !exists {
		number, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil, errors.New("argument is not a number")
		}
		return &Node{
			Data: &DataNode{
				ExpectedScore: number,
				Score:         number,
				MaxScore:      0,
			},
			Operator: nil,
			RawData:  data,
			Left:     nil,
			Right:    nil,
		}, nil
	} else {
		for _, del := range delimiters {
			operator := fmt.Sprintf("%c", del)
			operands := strings.SplitN(data, operator, 2)
			if len(operands) != 2 {
				continue
			}
			left, err := BuildAST(operands[0])
			if err != nil {
				return nil, err
			}
			right, err := BuildAST(operands[1])
			if err != nil {
				return nil, err
			}
			return &Node{
				Data: nil,
				Operator: &OperatorNode{
					Operator: operator,
				},
				RawData: data,
				Left:    left,
				Right:   right,
			}, nil
		}
		panic("that couldn't be possible")
	}
}

func PrintNode(node *Node) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		fmt.Println(node.RawData)
	} else {
		PrintNode(node.Left)
		fmt.Println(node.Operator.Operator)
		PrintNode(node.Right)
	}
}

func Solve(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return node
	} else {
		left := Solve(node.Left)
		right := Solve(node.Right)
		switch node.Operator.Operator {
		case "+":
			node.Data = &DataNode{
				ExpectedScore: left.Data.ExpectedScore + right.Data.ExpectedScore,
				Score:         left.Data.Score + right.Data.Score,
				MaxScore:      left.Data.MaxScore + right.Data.MaxScore,
			}
			node.RawData = node.Left.RawData + "+" + node.Right.RawData
			break
		case "-":
			node.Data = &DataNode{
				ExpectedScore: left.Data.ExpectedScore - right.Data.ExpectedScore,
				Score:         left.Data.Score - right.Data.Score,
				MaxScore:      left.Data.MaxScore - right.Data.MaxScore,
			}
			node.RawData = node.Left.RawData + "-" + node.Right.RawData
			break
		case "d", "к":
			score := int64(0)
			solution := "("
			for i := int64(0); i < left.Data.Score; i++ {
				rng := rand.Int63n(right.Data.Score) + 1
				score += rng
				solution += fmt.Sprint(rng)
				if i+1 != left.Data.Score {
					solution += "+"
				}
			}
			solution += ")"
			node.Data = &DataNode{
				ExpectedScore: left.Data.Score * ((2 + right.Data.Score) / 2),
				Score:         score,
				MaxScore:      left.Data.Score * right.Data.Score,
			}
			node.RawData = solution
			break
		}
	}
	return node
}

func Calc(data string) (*DataNode, error) {
	node, err := BuildAST(data)
	if err != nil {
		return nil, err
	}
	return Solve(node).Data, nil
}
