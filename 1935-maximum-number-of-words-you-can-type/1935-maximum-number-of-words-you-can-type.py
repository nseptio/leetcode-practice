class Solution:
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        count = 0
        broken_list = list(brokenLetters)
        text_list = text.split()
        for word in text_list:
            for i in word:
                if i in broken_list:
                    count += 1
                    break

        return len(text_list) - count