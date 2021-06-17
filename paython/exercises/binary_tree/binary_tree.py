class BinaryTreeNode:
    def __init__(self, data):
        self.data = data
        self.leftChild = None
        self.rightChild = None


def insert(root, newValue):
    if root is None:
        root = BinaryTreeNode(newValue)
        return root
    if newValue < root.data:
        root.leftChild = insert(root.leftChild, newValue)
    else:
        root.rightChild = insert(root.rightChild, newValue)
    return root

def height(root):
  if root is None:
    return 0
  leftHeight = height(root.leftChild)
  rightHeight = height(root.rightChild)
  if (rightHeight > leftHeight):
    return rightHeight+1
  else:
    return leftHeight+1

def checkBalancedBinaryTree(root):
  if root is None:
    return True
  leftheight = height(root.leftChild)
  rightHeight = height(root.rightChild)
  if (abs(leftheight - rightHeight) > 1):
    return False
  leftCheck = checkBalancedBinaryTree(root.leftChild)
  rightCheck = checkBalancedBinaryTree(root.rightChild)
  return leftCheck == True and rightCheck == True

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
  