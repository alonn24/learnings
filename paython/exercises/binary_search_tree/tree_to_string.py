def treeToString(root):
  def innerTreeToString(node, string):
    if node is None:
      return
    string.append(str(node.data))

    if node.leftChild:
      string.append('(')
      innerTreeToString(node.leftChild, string)
      string.append(')')

    if node.rightChild:
      string.append('(')
      innerTreeToString(node.rightChild, string)
      string.append(')')
  res = []
  innerTreeToString(root, res)
  return ''.join(res)
  