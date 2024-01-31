class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        v1 = version1.split(".")
        v2 = version2.split(".")

        for i in range(len(v1)):
            if len(v1[i]) != v1[i].count("0"):
                v1[i] = int(v1[i])
            else:
                v1[i] = 0

        for i in range(len(v2)):
            if len(v2[i]) != v2[i].count("0"):
                v2[i] = int(v2[i])
            else:
                v2[i] = 0
        
        if len(v1) > len(v2):
            for i in range(len(v2)):
                if v1[i] < v2[i]:
                    return -1
                elif v1[i] > v2[i]:
                    return 1
            for i in range(len(v2), len(v1), 1):
                if v1[i] > 0:
                    return 1
        else:
            for i in range(len(v1)):
                if v1[i] < v2[i]:
                    return -1
                elif v1[i] > v2[i]:
                    return 1
            for i in range(len(v1), len(v2), 1):
                if v2[i] > 0:
                    return -1

        return 0