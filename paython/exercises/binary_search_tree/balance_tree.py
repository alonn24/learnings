import sys
import math
from binary_tree import BinaryTreeNode

def storeBSTNodes(root, nodes):
  if not root:
    return
  storeBSTNodes(root.leftChild, nodes)
  nodes.append(root)
  storeBSTNodes(root.rightChild, nodes)

def buildTreeUtil(nodes, start, end):
  if start > end:
    return None
  mid=(start+end)//2
  node=nodes[mid]
  node.leftChild = buildTreeUtil(nodes, start, mid-1)
  node.rightChild = buildTreeUtil(nodes, mid+1, end)
  return node

def balanceTree(root):
  nodes=[]
  storeBSTNodes(root, nodes)
  n = len(nodes)
  return buildTreeUtil(nodes, 0, n-1)
