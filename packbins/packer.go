package packbins

import (
	"fmt"
	"log"
	"sort"
)

type Bin struct {
	W, H int64
	Name string
}

func (b *Bin) String() string {
	return fmt.Sprintf("Bin{W=%d,H=%d,Name:'%s'}", b.W, b.H, b.Name)
}

type Fit struct {
	X, Y int64
	Name string
}

type Box struct {
	W, H int64
	Fits []*Fit
}

type Packer interface {
	Fit(bins []*Bin) *Box
}

type gNode struct {
	x, y, w, h  int64
	right, down *gNode
	used        bool
}

func (n *gNode) String() string {
	return fmt.Sprintf("Node{x=%d,y=%d,w=%d,h=%d,used=%t}", n.x, n.y, n.w, n.h, n.used)
}

type GrowPacker struct {
	root *gNode
}

type BinWidthCmp []*Bin

func (bwc BinWidthCmp) Len() int           { return len(bwc) }
func (bwc BinWidthCmp) Less(i, j int) bool { return bwc[i].W > bwc[j].W }
func (bwc BinWidthCmp) Swap(i, j int)      { bwc[i], bwc[j] = bwc[j], bwc[i] }

type BinHeightCmp []*Bin

func (bhc BinHeightCmp) Len() int           { return len(bhc) }
func (bhc BinHeightCmp) Less(i, j int) bool { return bhc[i].H > bhc[j].H }
func (bhc BinHeightCmp) Swap(i, j int)      { bhc[i], bhc[j] = bhc[j], bhc[i] }

type BinMaxsizeCmp []*Bin

func (bmc BinMaxsizeCmp) Len() int { return len(bmc) }
func (bmc BinMaxsizeCmp) Less(i, j int) bool {
	return max(bmc[i].W, bmc[i].H) > max(bmc[j].W, bmc[j].H)
}
func (bmc BinMaxsizeCmp) Swap(i, j int) { bmc[i], bmc[j] = bmc[j], bmc[i] }

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

type BinAreaCmp []*Bin

func (bac BinAreaCmp) Len() int           { return len(bac) }
func (bac BinAreaCmp) Less(i, j int) bool { return bac[i].W*bac[i].H > bac[j].W*bac[j].H }
func (bac BinAreaCmp) Swap(i, j int)      { bac[i], bac[j] = bac[j], bac[i] }

type BinWHCmp []*Bin

func (bwhc BinWHCmp) Len() int { return len(bwhc) }
func (bwhc BinWHCmp) Less(i, j int) bool {
	return bwhc[i].W > bwhc[j].W || (bwhc[i].W == bwhc[j].W && bwhc[i].H > bwhc[j].H)
}
func (bwhc BinWHCmp) Swap(i, j int) { bwhc[i], bwhc[j] = bwhc[j], bwhc[i] }

type BinHWCmp []*Bin

func (bhwc BinHWCmp) Len() int { return len(bhwc) }
func (bhwc BinHWCmp) Less(i, j int) bool {
	return bhwc[i].H > bhwc[j].H || (bhwc[i].H == bhwc[j].H && bhwc[i].W > bhwc[j].W)
}
func (bhwc BinHWCmp) Swap(i, j int) { bhwc[i], bhwc[j] = bhwc[j], bhwc[i] }

func (gp *GrowPacker) DirectlyFit(bins []*Bin) *Box {
	return gp.fit(bins)
}

func (gp *GrowPacker) Fit(bins []*Bin) *Box {
	if len(bins) < 1 {
		return nil
	}
	var sum int64 = 0
	for _, bin := range bins {
		sum += bin.W * bin.H
	}
	sort.Sort(BinWidthCmp(bins))
	wbox := gp.fit(bins)
	log.Printf("sorted by width box{w=%d,h=%d} filled=%.4f%%", wbox.W, wbox.H, float64(100*sum)/float64(wbox.W*wbox.H))
	sort.Sort(BinHeightCmp(bins))
	hbox := gp.fit(bins)
	log.Printf("sorted by height box{w=%d,h=%d} filled=%.4f%%", hbox.W, hbox.H, float64(100*sum)/float64(hbox.W*hbox.H))
	sort.Sort(BinMaxsizeCmp(bins))
	mbox := gp.fit(bins)
	log.Printf("sorted by maxsize box{w=%d,h=%d} filled=%.4f%%", mbox.W, mbox.H, float64(100*sum)/float64(mbox.W*mbox.H))
	sort.Sort(BinAreaCmp(bins))
	abox := gp.fit(bins)
	log.Printf("sorted by area box{w=%d,h=%d} filled=%.4f%%", abox.W, abox.H, float64(100*sum)/float64(abox.W*abox.H))
	sort.Sort(BinWHCmp(bins))
	whbox := gp.fit(bins)
	log.Printf("sorted by width then height box{w=%d,h=%d} filled=%.4f%%", whbox.W, whbox.H, float64(100*sum)/float64(whbox.W*whbox.H))
	sort.Sort(BinHWCmp(bins))
	hwbox := gp.fit(bins)
	log.Printf("sorted by height then width box{w=%d,h=%d} filled=%.4f%%", hwbox.W, hwbox.H, float64(100*sum)/float64(hwbox.W*hwbox.H))

	box := wbox
	if box.W*box.H > hbox.W*hbox.H {
		box = hbox
	}
	if box.W*box.H > mbox.W*mbox.H {
		box = mbox
	}
	if box.W*box.H > abox.W*abox.H {
		box = abox
	}
	if box.W*box.H > whbox.W*whbox.H {
		box = whbox
	}
	if box.W*box.H > hwbox.W*hwbox.H {
		box = hwbox
	}
	return box
}

func (gp *GrowPacker) fit(bins []*Bin) *Box {
	var fits = make([]*Fit, 0, len(bins))
	gp.root = &gNode{x: 0, y: 0, w: bins[0].W, h: bins[0].H}
	for _, bin := range bins {
		if node := gp.findNode(gp.root, bin.W, bin.H); node != nil {
			if n := gp.splitNode(node, bin.W, bin.H); n != nil {
				fits = append(fits, &Fit{X: n.x, Y: n.y, Name: bin.Name})
			} else {
				panic(fmt.Errorf("cannot split node root=%s bin=%s", gp.root, bin))
			}
		} else {
			if n := gp.growNode(gp.root, bin.W, bin.H); n != nil {
				fits = append(fits, &Fit{X: n.x, Y: n.y, Name: bin.Name})
			} else {
				panic(fmt.Errorf("cannot grow node root=%s bin=%s", gp.root, bin))
			}
		}
	}
	return &Box{W: gp.root.w, H: gp.root.h, Fits: fits}
}

func (gp *GrowPacker) findNode(node *gNode, w, h int64) *gNode {
	if node.used == true {
		if r := gp.findNode(node.right, w, h); r != nil {
			return r
		} else if d := gp.findNode(node.down, w, h); d != nil {
			return d
		}
	} else if node.w >= w && node.h >= h {
		return node
	}
	return nil
}

func (gp *GrowPacker) splitNode(node *gNode, w, h int64) *gNode {
	node.used = true
	node.down = &gNode{x: node.x, y: node.y + h, w: node.w, h: node.h - h}
	node.right = &gNode{x: node.x + w, y: node.y, w: node.w - w, h: h}
	return node
}

func (gp *GrowPacker) growNode(node *gNode, w, h int64) *gNode {
	var canGrowRight, canGrowDown = node.h >= h, node.w >= w
	var shouldGrowRight, shouldGrowDown = canGrowRight && node.h >= node.w+w, canGrowDown && node.w >= node.h+h
	if shouldGrowRight == true {
		return gp.growRight(w, h)
	} else if shouldGrowDown == true {
		return gp.growDown(w, h)
	} else if canGrowRight == true {
		return gp.growRight(w, h)
	} else if canGrowDown == true {
		return gp.growDown(w, h)
	}
	return nil
}

func (gp *GrowPacker) growRight(w, h int64) *gNode {
	lastRoot := gp.root
	gp.root = &gNode{
		x:     0,
		y:     0,
		w:     lastRoot.w + w,
		h:     lastRoot.h,
		right: &gNode{x: lastRoot.w, y: 0, w: w, h: lastRoot.h},
		down:  lastRoot,
		used:  true,
	}
	if node := gp.findNode(gp.root, w, h); node != nil {
		return gp.splitNode(node, w, h)
	}
	return nil
}

func (gp *GrowPacker) growDown(w, h int64) *gNode {
	lastRoot := gp.root
	gp.root = &gNode{
		x:     0,
		y:     0,
		w:     lastRoot.w,
		h:     lastRoot.h + h,
		right: lastRoot,
		down:  &gNode{x: 0, y: lastRoot.h, w: lastRoot.w, h: h},
		used:  true,
	}
	if node := gp.findNode(gp.root, w, h); node != nil {
		return gp.splitNode(node, w, h)
	}
	return nil
}
