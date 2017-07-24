package flex

import "testing"

func TestDont_cache_computed_flex_basis_between_layouts(t *testing.T) {
	config := NewConfig()
	ConfigSetExperimentalFeatureEnabled(config, ExperimentalFeatureWebFlexBasis, true)

	root := NewNodeWithConfig(config)
	root.StyleSetHeightPercent(100)
	root.StyleSetWidthPercent(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexBasisPercent(100)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 100, Undefined, DirectionLTR)
	NodeCalculateLayout(root, 100, 100, DirectionLTR)

	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())
}

func TestRecalculate_resolvedDimonsion_onchange(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	rootChild0.StyleSetMinHeight(10)
	rootChild0.StyleSetMaxHeight(10)
	NodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	rootChild0.StyleSetMinHeight(Undefined)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())
}
