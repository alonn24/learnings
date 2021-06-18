from binary_tree import BinaryTreeNode
from tree_to_string import treeToString
from balance_tree import balanceTree


def test_treeToString():
    root = BinaryTreeNode(4, BinaryTreeNode(
        3, BinaryTreeNode(2, BinaryTreeNode(1))))
    assert treeToString(root) == "4(3(2(1)))"
    balanced = balanceTree(root)
    assert treeToString(balanced) == "2(1)(3(4))"
