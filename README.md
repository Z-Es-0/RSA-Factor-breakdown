# RSA-Factor breakdown
使用工作池对RSA中的n进行素因子拆分

bin 文件夹里的rsa.exe 是一个有一个命令行参数的可执行文件，参数为 待分解的n值，输出为拆分后的p和q

```shell
➜  .\rsa.exe 10000000000037024930000000092241
n is 10000000000037024930000000092241
Factor found: 10000000000037
p and q are 10000000000037 and 1000000000000002493
```

# 注意
 会大量占用计算机资源 (cpu使用到100%)

 # 效率

 当然这只是普通算法的并行优化版本，对于几百上千的密钥长度依然无可奈何，只会比普通的串行暴力分解快一点

 
