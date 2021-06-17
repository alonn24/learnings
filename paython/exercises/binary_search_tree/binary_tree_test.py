from binary_tree import insert, height

def test_height():
    root= insert(None,15)
    assert height(root) == 1
    insert(root,10)
    assert height(root) == 2
    insert(root,25)
    assert height(root) == 2
    insert(root,6)
    assert height(root) == 3
    insert(root,14)
    insert(root,20)
    insert(root,60)
    assert height(root) == 3
