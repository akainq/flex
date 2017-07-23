package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newHadOverflowTests() (*YGConfig, *YGNode) {
	config := YGConfigNew()
	root := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(root, 200)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionColumn)
	YGNodeStyleSetFlexWrap(root, YGWrapNoWrap)
	return config, root
}

func TestChildren_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child0, 80)
	YGNodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(child0, YGEdgeBottom, 15)
	YGNodeInsertChild(root, child0, 0)
	child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1, 80)
	YGNodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, YGEdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_no_wrap_and_no_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child0, 80)
	YGNodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1, 80)
	YGNodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, YGEdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}

func TestNo_overflow_no_wrap_and_flex_children(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child0, 80)
	YGNodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1, 80)
	YGNodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, YGEdgeBottom, 5)
	YGNodeStyleSetFlexShrink(child1, 1)
	YGNodeInsertChild(root, child1, 1)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.False(t, YGNodeLayoutGetHadOverflow(root))
}

func TestHadOverflow_gets_reset_if_not_logger_valid(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child0, 80)
	YGNodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1, 80)
	YGNodeStyleSetHeight(child1, 40)
	YGNodeStyleSetMargin(child1, YGEdgeBottom, 5)
	YGNodeInsertChild(root, child1, 1)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))

	YGNodeStyleSetFlexShrink(child1, 1)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.False(t, YGNodeLayoutGetHadOverflow(root))
}

func TestSpacing_overflow_in_nested_nodes(t *testing.T) {
	config, root := newHadOverflowTests()
	child0 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child0, 80)
	YGNodeStyleSetHeight(child0, 40)
	YGNodeStyleSetMargin(child0, YGEdgeTop, 10)
	YGNodeStyleSetMargin(child0, YGEdgeBottom, 10)
	YGNodeInsertChild(root, child0, 0)
	child1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1, 80)
	YGNodeStyleSetHeight(child1, 40)
	YGNodeInsertChild(root, child1, 1)
	child1_1 := YGNodeNewWithConfig(config)
	YGNodeStyleSetWidth(child1_1, 80)
	YGNodeStyleSetHeight(child1_1, 40)
	YGNodeStyleSetMargin(child1_1, YGEdgeBottom, 5)
	YGNodeInsertChild(child1, child1_1, 0)

	YGNodeCalculateLayout(root, 200, 100, YGDirectionLTR)

	assert.True(t, YGNodeLayoutGetHadOverflow(root))
}