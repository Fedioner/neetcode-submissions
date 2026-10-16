
class Solution {
public:
    std::string encode(std::vector<std::string>& strs) {
        std::string result;
        for (const std::string& str : strs) {
            result += std::to_string(str.size());
            result += '#';
            result += str;
        }
        return result;
    }

    std::vector<std::string> decode(std::string s) {
        std::vector<std::string> result;
        size_t i = 0;
        while (i < s.size()) {
            size_t j = s.find('#', i);
            size_t length = std::stoul(s.substr(i, j - i));
            size_t start = j + 1;
            result.push_back(s.substr(start, length));
            i = start + length;
        }
        return result;
    }
};