import unittest

def remove_duplicates(i):
    dup = set()
    output = []
    for item in i:
        if item not in dup:
            dup.add(item)
            output.append(item)
    return output

class TestRemoveDuplicates(unittest.TestCase):
    
    def test_remove_duplicates(self):
        self.assertEqual(remove_duplicates([1, 2, 3, 3, 4, 5, 2, 1]), [1, 2, 3, 4, 5])
    def test_remove_duplicates_all_duplicates(self):
        self.assertEqual(remove_duplicates([1, 1, 1, 1]), [1])
    def test_remove_duplicates_empty_list(self):
        self.assertEqual(remove_duplicates([]), [])
    def test_remove_duplicates_no_duplicate(self):
        self.assertEqual(remove_duplicates([5, 2, 3, 1]), [5,2,3,1])

if __name__ == '__main__':
    unittest.main()