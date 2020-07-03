## Exec multi command on multi nodes 

#####  Dùng để deploy code lên nhiều node, sử dụng file pem để auth

1. Chỉnh sửa list ip và command trong `config.yaml`

2. Thêm file `pem` vào thư mục gốc

3. `make install`

4. Run

   ```sh
   sudo dplt deploy --config=config.yaml --pem=pem.pem
   ```

   

