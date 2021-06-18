from binary_tree import insert
from check_balance import checkBalancedBinaryTree
from tree_to_string import treeToString


def test_balanced():
    root = insert(None, 15)
    insert(root, 10)
    insert(root, 25)
    insert(root, 6)
    insert(root, 14)
    insert(root, 20)
    insert(root, 60)
    print(treeToString(root))
    assert checkBalancedBinaryTree(root)
