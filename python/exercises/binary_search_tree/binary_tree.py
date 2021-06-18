class BinaryTreeNode:
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.leftChild = left
        self.rightChild = right


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
