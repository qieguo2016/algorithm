package sort

// BuildDictTree 字典树
func BuildDictTree(arr []string) map[string]interface{} {
	tree := map[string]interface{}{}
	for _, chars := range arr {
		curr := tree
		for _, c := range chars {
			cs := string(c)
			if v, exist := curr[cs]; exist {
				curr = v.(map[string]interface{})
			} else {
				curr[cs] = map[string]interface{}{}
				curr = curr[cs].(map[string]interface{})
			}
		}
		curr["value"] = 1
	}
	return tree
}

// B+树实现
// B+树的索引节点在下层节点中，每一个索引节点对应下层的一个子分页，这与B树不同，B树是左右子分页的结构
// 比如：
// 5     											28        					65
// 5, 		10, 			20 				28				35				65
// 5,6,9 	10,12,15	20,22,27	28,30,33	35,38,50	65,80,120

const (
	BPTREE_PAGE_TYPE_LEAF  = 1 // 叶子页
	BPTREE_PAGE_TYPE_INDEX = 2 // 索引页

	CODE_SUCCESS   = 100000
	CODE_NOT_FOUND = 100001
)

type BPTree struct {
	Order    int
	RootPage *BPTreePage
	LeafPage *BPTreePage
}

type BPTreePage struct {
	Type       int
	Pre, Next  *BPTreePage
	ParentNode *BPTreeNode
	HeadNode   *BPTreeNode
	Len        int
	Cap        int
}

type BPTreeNode struct {
	Index     int
	Data      interface{}
	Next      *BPTreeNode
	Page      *BPTreePage
	ChildPage *BPTreePage
}

type Response struct {
	Code int
	Data interface{}
}

func NewBPTree(order int) BPTree {
	return BPTree{
		Order:    order,
		LeafPage: &BPTreePage{Cap: order - 1, Type: BPTREE_PAGE_TYPE_LEAF},
		RootPage: &BPTreePage{Cap: order - 1, Type: BPTREE_PAGE_TYPE_INDEX},
	}
}

func (bp *BPTree) Insert(index int, data interface{}) {
	page := bp.LeafPage
	if bp.RootPage.Len == 0 {
		parentNode := &BPTreeNode{Index: index, Page: bp.RootPage, ChildPage: page}
		page.ParentNode = parentNode
		bp.RootPage.Insert(parentNode)
	} else {
		page = bp.searchLeafPage(index, bp.RootPage)
	}
	node := &BPTreeNode{Index: index, Data: data}
	bp.addNodeToPage(node, page)
}

func (bp *BPTree) searchLeafPage(index int, page *BPTreePage) *BPTreePage {
	if page.Type == BPTREE_PAGE_TYPE_LEAF {
		return page
	}
	cur := page.HeadNode
	for cur != nil && cur.Next != nil && cur.Next.Index <= index {
		cur = cur.Next
	}
	childPage := cur.ChildPage
	return bp.searchLeafPage(index, childPage)
}

func (bp *BPTree) searchNode(index int) *BPTreeNode {
	page := bp.LeafPage
	if bp.RootPage.Len > 0 {
		page = bp.searchLeafPage(index, bp.RootPage)
	}
	if page == nil {
		return nil
	}
	cur := page.HeadNode
	for cur != nil {
		if cur.Index == index {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (page *BPTreePage) Insert(node *BPTreeNode) {
	cur := page.HeadNode
	if cur == nil || cur.Index > node.Index {
		node.Next = cur
		page.HeadNode = node
	} else {
		for cur.Next != nil && cur.Next.Index < node.Index {
			cur = cur.Next
		}
		node.Next = cur.Next
		cur.Next = node
	}
	node.Page = page
	page.Len++
}

func (page *BPTreePage) Split() (*BPTreePage, *BPTreePage) {
	// 找中心
	cur := page.HeadNode
	n := 1
	for n < (page.Len+1)/2 {
		cur = cur.Next
		n++
	}

	newHeadNode := cur.Next
	cur.Next = nil // 切断原链表

	// 新page
	newPage := &BPTreePage{
		Type:     page.Type,
		Pre:      page,
		Next:     page.Next,
		HeadNode: newHeadNode,
		Cap:      page.Cap,
		Len:      page.Len - n,
	}

	// 调整原page
	page.Next = newPage
	page.Len = n

	// 调整new page节点归属
	for newHeadNode.Next != nil {
		newHeadNode.Page = newPage
		newHeadNode = newHeadNode.Next
	}

	return page, newPage
}

func (page *BPTreePage) Up() {
	if page.ParentNode != nil && page.ParentNode.Index != page.HeadNode.Index {
		page.ParentNode.Index = page.HeadNode.Index
		page.ParentNode.Page.Up()
	}
}

func (page *BPTreePage) DeleteByIndex(index int) *BPTreeNode {
	var pre *BPTreeNode
	cur := page.HeadNode
	for cur != nil {
		if cur.Index == index {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil || cur.Index != index {
		return nil
	}
	if pre == nil { // 头部
		page.HeadNode = cur.Next
	} else {
		pre.Next = cur.Next
	}
	cur.Next = nil
	cur.Page = nil
	page.Len--
	if page.Len == 0 { // 删光了，去掉空页
		n := page.Next
		p := page.Pre
		if p != nil {
			p.Next = n
		}
		if n != nil {
			n.Pre = p
		}
		page.Next = nil
		page.Pre = nil
		if page.ParentNode != nil { // 向上递归删除
			page.ParentNode.Page.DeleteByIndex(index)
		}
		return cur
	}
	page.Up() // 向上递归检查
	return cur
}

func (bp *BPTree) addNodeToPage(node *BPTreeNode, page *BPTreePage) {
	// 加入page
	page.Insert(node)
	page.Up()

	// 页面未满
	if page.Len <= page.Cap {
		return
	}

	// 页满，需要拆分
	leftPage, rightPage := page.Split()
	// 新页上层node
	rightIndexNode := &BPTreeNode{
		Index:     rightPage.HeadNode.Index,
		ChildPage: rightPage,
	}
	rightPage.ParentNode = rightIndexNode

	if leftPage.ParentNode == nil { // 顶层
		bp.RootPage = &BPTreePage{Cap: bp.Order - 1, Type: BPTREE_PAGE_TYPE_INDEX} // 当前可能是root，需要重置
		leftIndexNode := &BPTreeNode{
			Index:     leftPage.HeadNode.Index,
			ChildPage: leftPage,
		}
		leftPage.ParentNode = leftIndexNode
		bp.addNodeToPage(leftIndexNode, bp.RootPage)
	}
	// 递归向上
	bp.addNodeToPage(rightIndexNode, leftPage.ParentNode.Page)
}

func (bp *BPTree) Delete(index int) Response {
	page := bp.LeafPage
	if bp.RootPage.Len > 0 {
		page = bp.searchLeafPage(index, bp.RootPage)
	}
	if page == nil {
		return Response{Code: CODE_NOT_FOUND, Data: nil}
	}
	node := page.DeleteByIndex(index)
	if node == nil {
		return Response{Code: CODE_NOT_FOUND, Data: nil}
	}
	return Response{Code: CODE_SUCCESS, Data: node.Data}
}

func (bp *BPTree) Search(index int) Response {
	node := bp.searchNode(index)
	if node == nil {
		return Response{Code: CODE_NOT_FOUND, Data: nil}
	}
	return Response{Code: CODE_SUCCESS, Data: node.Data}
}
