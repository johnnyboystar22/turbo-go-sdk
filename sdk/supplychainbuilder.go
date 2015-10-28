package sdk

type SupplyChainBuilder struct {
	SupplyChainNodes map[*EntityDTO_EntityType]*SupplyChainNodeBuilder
	currentNode      *SupplyChainNodeBuilder
}

func NewSupplyChainBuilder() *SupplyChainBuilder {
	return &SupplyChainBuilder{}
}

// TODO: supply chain here is a set(map) or a list
// if a map, key is the node builder, value is a boolean
func (scb *SupplyChainBuilder) Create() []*TemplateDTO {
	allNodeBuilders := make(map[*SupplyChainNodeBuilder]bool)
	for _, nodeBuilder := range scb.SupplyChainNodes {
		allNodeBuilders[nodeBuilder] = true
	}

	// create nodes from all node builders and put it into result
	var allNodes []*TemplateDTO
	for nodeBuilder, _ := range allNodeBuilders {
		allNodes = append(allNodes, nodeBuilder.Create())
	}

	return allNodes
}

// To build a supply chain, must specify the top nodebuilder first.
// Here it initialize SupplyChainNodes and currentNode
func (scb *SupplyChainBuilder) Top(topNode *SupplyChainNodeBuilder) *SupplyChainBuilder {
	supplyChianNodesBuilderMap := make(map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)
	scb.SupplyChainNodes = supplyChianNodesBuilderMap
	topNodeEntityType := topNode.getEntity()
	scb.SupplyChainNodes[&topNodeEntityType] = topNode

	scb.currentNode = topNode

	return scb
}

// Add an entity node to supply chain
func (scb *SupplyChainBuilder) Entity(node *SupplyChainNodeBuilder) *SupplyChainBuilder {
	if hasTop := scb.hasTopNode(); !hasTop {
		//TODO should have error
		return scb
	}
	nodeEntityType := node.getEntity()
	scb.SupplyChainNodes[&nodeEntityType] = node

	scb.currentNode = node

	return scb
}

// check if the top node has been initialized
func (scb *SupplyChainBuilder) hasTopNode() bool {
	if scb.SupplyChainNodes == nil {
		return false
	}
	return true
}
