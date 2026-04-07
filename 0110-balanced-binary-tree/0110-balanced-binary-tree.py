# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        self.diff = 0
        def maxDepth(curr: Optional[Treenod]) -> int:
            if not curr:
                return 0
            
            left = maxDepth(curr.left)
            right = maxDepth(curr.right)

            diff = abs(left - right)
            self.diff =  diff if diff > 1 else self.diff

            return max(left, right) + 1

        maxDepth(root)
        return False if self.diff > 1 else True