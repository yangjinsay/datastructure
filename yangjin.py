#!/bin/python3
# -*- coding: UTF-8 -*-

def longestCommonPrefix(strs):
    if not strs: return ""
    s1 = min(strs)
    s2 = max(strs)
    for i,x in enumerate(strs):
        if x != s2[i]:
            return s2[:i]
    return s1

s = longestCommonPrefix(["ab", "abc"])
print(s)

print("你好, 世界")