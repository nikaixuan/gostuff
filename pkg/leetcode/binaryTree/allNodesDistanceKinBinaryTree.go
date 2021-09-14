package binaryTree

// 863
var resk []int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	resk = make([]int, 0)
	// parent or other children of same parent with k distance
	postOrderFindNode(root, target, k)
	// child node with k distance
	findKNode(target, k)
	return resk
}

func postOrderFindNode(root *TreeNode, target *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	if root == target {
		return 1
	}
	// left, right distance
	ldis := postOrderFindNode(root.Left, target, k)
	rdis := postOrderFindNode(root.Right, target, k)
	// k distance
	if ldis == k || rdis == k {
		resk = append(resk, root.Val)
	}
	if ldis > 0 && ldis < k {
		findKNode(root.Right, k-ldis-1)
		return ldis + 1
	}
	if rdis > 0 && rdis < k {
		findKNode(root.Left, k-rdis-1)
		return rdis + 1
	}
	return 0
}

func findKNode(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if k == 0 {
		resk = append(resk, root.Val)
		return
	}
	findKNode(root.Left, k-1)
	findKNode(root.Right, k-1)
}
