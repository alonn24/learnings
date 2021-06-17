from binary_tree import height

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
