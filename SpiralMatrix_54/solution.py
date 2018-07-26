class Solution:
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """

        if len(matrix) == 0 or len(matrix[0]) == 0:
            return []

        T = 0
        R = len(matrix[0]) - 1
        B = len(matrix) - 1
        L = 0

        result = []
        direct = 0
        while T <= B and R >= L:
            if direct == 0:
                for i in range(L, R + 1):
                    result.append(matrix[T][i])
                T += 1
            elif direct == 1:
                for i in range(T, B + 1):
                    result.append(matrix[i][R])
                R -= 1
            elif direct == 2:
                for i in range(R, L - 1, -1):
                    result.append(matrix[B][i])
                B -= 1
            elif direct == 3:
                for i in range(B, T - 1, -1):
                    result.append(matrix[i][L])
                L += 1

            direct = (direct + 1) % 4

        return result