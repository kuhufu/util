openssl ecparam -genkey -name secp521r1 -out private.pem #生成私钥
openssl ec -in private.pem -pubout -out public.pem #生成公钥
