# tar包源码文字说明

tar包用于对tar文件的访问，它旨在统一处理tar文件。包括GNU以及BSD对产生的tar包的差异
之所以说会有不同是因为不同的系统生成的tar文件的header并不一样
有的header的定义是这样的。
···c
    //  FreeBSD
    struct header_old_tar {
		   char	name[100];
		   char	mode[8];
		   char	uid[8];
		   char	gid[8];
		   char	size[12];
		   char	mtime[12];
		   char	checksum[8];
		   char	linkflag[1];
		   char	linkname[100];
		   char	pad[255];
	 };

	 // POSIX
	 struct header_posix_ustar {
     		   char	name[100];
     		   char	mode[8];
     		   char	uid[8];
     		   char	gid[8];
     		   char	size[12];
     		   char	mtime[12];
     		   char	checksum[8];
     		   char	typeflag[1];
     		   char	linkname[100];
     		   char	magic[6];
     		   char	version[2];
     		   char	uname[32];
     		   char	gname[32];
     		   char	devmajor[8];
     		   char	devminor[8];
     		   char	prefix[155];
     		   char	pad[12];
     };
     // 我们可以看到这两种头从定义上就有很大的区别
···
