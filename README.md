### Supported file formats

- cbz

---

```shell
cbmerge λ ./cbmerge -h
Usage of ./cbmerge:
  -dst string
    	output cbz file
  -src string
    	input directory (default ".")
    	
cbmerge λ BLOCKSIZE=1 du -s sample
du: minimum blocksize is 512
3938464	sample

cbmerge λ find sample -type f -exec ls -l {} \; | awk '{sum += $5} END {print sum}'
2016481785

cbmerge λ ./cbmerge -src sample/Parasyte\ \(2015-2016\)\ \(Digital\)\ \(danke-Empire\)/ -dst Parasyte.cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v01 (2015) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v02 (2015) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v03 (2015) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v04 (2015) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v05 (2015) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v06 (2016) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v07 (2016) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:40 opening file sample/Parasyte (2015-2016) (Digital) (danke-Empire)/Parasyte v08 (2016) (Digital) (F) (danke-Empire).cbz
2022/11/12 15:53:43 done merging sample/Parasyte (2015-2016) (Digital) (danke-Empire)/*.cbz to Parasyte.cbz

cbmerge λ ls -s Parasyte.cbz | awk '{size = $1 * 512} END {print size}'
2023804928
```