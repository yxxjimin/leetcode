"""
Runtime : 0 ms
Memory  : 12.39 MB
"""
class Solution(object):
    def toBinary(self, num):
        res = []
        while num > 0:
            res.append(num % 2)
            num /= 2
        return list(reversed(res))
    
    def toDecimal(self, binary):
        ans = 0
        exp = 1
        for i in reversed(binary):
            ans += exp * i
            exp *= 2
        return ans

    def minimizeXor(self, num1, num2):
        """
        :type num1: int
        :type num2: int
        :rtype: int
        """
        num1_binary = self.toBinary(num1)
        num2_binary = self.toBinary(num2)

        fill = [
            0 for _ in range(abs(len(num1_binary) - len(num2_binary)))
        ]
        if len(num1_binary) > len(num2_binary):
            num2_binary = fill + num2_binary
        else:
            num1_binary = fill + num1_binary
        
        binary_len = len(num1_binary)
        answer_binary = [0 for _ in range(binary_len)]
        ones_cnt = sum(num2_binary)
        zeros_cnt = binary_len - ones_cnt

        for i in range(binary_len):
            if num1_binary[i] == 1 and ones_cnt > 0:
                answer_binary[i] = 1
                ones_cnt -= 1
            elif num1_binary[i] == 0 and zeros_cnt > 0:
                zeros_cnt -= 1
            else:
                break
        
        if ones_cnt > 0:
            answer_binary[binary_len - ones_cnt:] = [1] * ones_cnt
            
        return self.toDecimal(answer_binary)


"""
Runtime : 0 ms
Memory  : 12.26 MB
"""
class Solution2(object):
    def _is_set(self, num, bit):
        return (num & (1 << bit)) != 0
    
    def _set_bit(self, num, bit):
        return num | (1 << bit)

    def minimizeXor(self, num1, num2):
        """
        :type num1: int
        :type num2: int
        :rtype: int
        """
        ones_cnt = bin(num2).count("1")
        curr_ones_cnt = 0
        curr_bit = max(len(bin(num1)), len(bin(num2))) - 3
        answer = 0

        while curr_ones_cnt < ones_cnt:
            if (
                self._is_set(num1, curr_bit) 
                or ones_cnt - curr_ones_cnt > curr_bit
            ):
                answer = self._set_bit(answer, curr_bit)
                curr_ones_cnt += 1
            curr_bit -= 1
        
        return answer
    