/*source from Internet whose original author I don't know.
* just to save here for future using.
* and thanks the author*/
#define     HEXLEN      20                /* 十六进制日志字符数 */

struct  HEXBUF
{
 int   point;           /* 数据显示位置 */
 int   chn  ;           /* 汉字标识     */
 char  scop[ 6  + 1 ];  /* 位置偏移数   */
 char  sstr[ 20 + 1 ];  /* ASCII 打印串 */
 char  shex[ 61 + 1 ];  /* 十六进制打印串 */
} ;
 
/*
 * ** 十六进制打印缓冲区初始化
 * */
void  HexBuf_init(struct HEXBUF  *p )
{
 p->point = 0;
 memset(p->scop, 0x0 , sizeof(p->scop));
 memset(p->sstr, 0x20, HEXLEN );
 memset(p->shex, 0x20, HEXLEN*3 );
 return ;
}
/*
 * ** 定长数据包 十六进制打印程序
 * */
int printHexLog( void *PInfo, char *Title, int mLen)
{
 int   cnt;        // 当前数据处理位置
 int   point ;
 char  temp[5];
 unsigned char  *ptr =  (unsigned  char *) PInfo ;
 struct HEXBUF   hex;
 memset(&hex, 0x0, sizeof(hex));
 printf( "%25s------- %s (%d) ------- \n", "", Title, mLen);
 HexBuf_init(&hex);
 for( cnt = 0; cnt < mLen ; cnt ++, ptr ++ )
 {
  if(hex.point && (cnt % HEXLEN) == 0 )
  {
   hex.chn = hex.chn & 0x01;
   if( hex.chn ) hex.sstr[HEXLEN-1]='.';
   	printf("%s: %s | %s |\n", hex.scop, hex.shex, hex.sstr);
   HexBuf_init( &hex );
  }
  if(hex.point == 0 )
   sprintf(hex.scop, "%04d", cnt );
  point = hex.point;
  if( (*ptr) < 0x20 ) hex.sstr[point] = '.';
  else {
   if(point == 0 && hex.chn )
   {
    hex.chn = 0;
    hex.sstr[0] = '.';
   }
   else
   {
    hex.sstr[point] = (*ptr);
    if( (*ptr & 0x80) == 0x80 ) hex.chn ++ ;
   }
  }
  point *= 3;
  sprintf(temp, "%02x ", (*ptr & 0xff ) );
  memcpy(hex.shex + point , temp, 3);
  hex.point ++;
 }
 if( hex.point )
  printf("%s: %s | %s |\n", hex.scop, hex.shex, hex.sstr);
 printf("%25s------- End ------- \n", "");
 return 0 ;
}
