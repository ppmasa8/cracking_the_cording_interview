def main()
    s1 = 'waterbottle'
    s2 = 'erbottlewat'
    p s1
    p s2
    p is_rotation(s1, s2)
end

def is_substring(s1s1, s2)
    for i in 0...s1s1.size - s2.size
        if s1s1[i..i+s2.size-1] == s2
            return true
            break
        end
    end
    return false
end

def is_rotation(s1, s2)
    if s1.size == s2.size && s1.size > 0
        s1s1 = s1 + s1
        return is_substring(s1s1, s2)
    end
    return false
end

main()