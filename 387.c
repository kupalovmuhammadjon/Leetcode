int firstUniqChar(char *s)
{
    int count = 0;
    for (int i = 0; i < strlen(s); i++)
    {
        for (int j = 0; j < strlen(s); j++)
        {
            if (s[i] == s[j])
            {
                count++;
            }
            if (count >= 2)
            {
                break;
            }
        }
        if (count == 1)
        {
            printf("%d", count);
            return i;
        }
        count = 0;
    }
    return -1;
}