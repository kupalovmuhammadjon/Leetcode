arr = [1024,512,256,128,64,32,16,8,4,2,1]
res = []
tpls = []
for i in arr:
    x = str(bin(i))
    c = x.count("1")
    tpls.append([i, c])

for i in range(len(tpls)):
    
    for j in range(len(tpls)):
        
        if tpls[i][1] < tpls[j][1]:
            tpls[i], tpls[j] = tpls[j], tpls[i]
        if tpls[i][1] == tpls[j][1] and tpls[i][0] < tpls[j][0]:
            tpls[i], tpls[j] = tpls[j], tpls[i]
            

for i in tpls:
    res.append(i[0])
    
    
print(res)
    

# done