class Solution:
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """

        combs = []
        for n in range(1, target + 1):
            if n in candidates:
                if n == target:
                    combs.append([n])
                else:
                    for subComb in self.combinationSum(candidates, target - n):
                        comb = [n] + subComb
                        comb.sort()
                        if comb not in combs:
                            combs.append(comb)

        return combs