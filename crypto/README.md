AES算法中填充模式的区别

* ZeroPadding，数据长度不对齐时使用0填充，否则不填充
* PKCS7Padding，假设数据长度需要填充n(n>0)个字节才对齐，那么填充n个字节，每个字节都是n;如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小
* PKCS5Padding，PKCS7Padding的子集，块大小固定为8字节。

PKCS5Padding和PKCS7Padding两者的区别在于PKCS5Padding是限制块大小的PKCS7Padding.

参考链接
-------------------------------------
https://blog.csdn.net/weixin_39530838/article/details/111021390