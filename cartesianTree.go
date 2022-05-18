package algorithms

func NewCartesianTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		//pass
	}
	root := &TreeNode{Val: nums[0]}
	prevNode := root
	for _, val := range nums[1:] {
		curNode := &TreeNode{Val: val}
		p := prevNode
		for ; p.Val > curNode.Val && p.Parent != nil; p = p.Parent {
		}
		if p.Val <= curNode.Val {
			curNode.Left, curNode.Parent = p.Right, p
			p.Right = curNode
		} else {
			p.Parent = curNode
			curNode.Left = p
			root = curNode
		}
		prevNode = curNode
	}
	return root
}
