// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR
//  0x04 32  KEYR
//  0x08 32  OPTKEYR
//  0x0C 32  SR
//  0x10 32  CR
//  0x14 32  AR
//  0x18 32  RESERVED
//  0x1C 32  OBR
//  0x20 32  WRPR
// Import:
//  stm32/o/f10x_hd/mmap
package flash

const (
	LATENCY   ACR_Bits = 0x03 << 0 //+ LATENCY[2:0] bits (Latency).
	LATENCY_0 ACR_Bits = 0x00 << 0 //  Bit 0.
	LATENCY_1 ACR_Bits = 0x01 << 0 //  Bit 0.
	LATENCY_2 ACR_Bits = 0x02 << 0 //  Bit 1.
	HLFCYA    ACR_Bits = 0x01 << 3 //+ Flash Half Cycle Access Enable.
	PRFTBE    ACR_Bits = 0x01 << 4 //+ Prefetch Buffer Enable.
	PRFTBS    ACR_Bits = 0x01 << 5 //+ Prefetch Buffer Status.
)

const (
	FKEYR KEYR_Bits = 0xFFFFFFFF << 0 //+ FPEC Key.
)

const ()

const (
	BSY      SR_Bits = 0x01 << 0 //+ Busy.
	PGERR    SR_Bits = 0x01 << 2 //+ Programming Error.
	WRPRTERR SR_Bits = 0x01 << 4 //+ Write Protection Error.
	EOP      SR_Bits = 0x01 << 5 //+ End of operation.
)

const (
	PG     CR_Bits = 0x01 << 0  //+ Programming.
	PER    CR_Bits = 0x01 << 1  //+ Page Erase.
	MER    CR_Bits = 0x01 << 2  //+ Mass Erase.
	OPTPG  CR_Bits = 0x01 << 4  //+ Option Byte Programming.
	OPTER  CR_Bits = 0x01 << 5  //+ Option Byte Erase.
	STRT   CR_Bits = 0x01 << 6  //+ Start.
	LOCK   CR_Bits = 0x01 << 7  //+ Lock.
	OPTWRE CR_Bits = 0x01 << 9  //+ Option Bytes Write Enable.
	ERRIE  CR_Bits = 0x01 << 10 //+ Error Interrupt Enable.
	EOPIE  CR_Bits = 0x01 << 12 //+ End of operation interrupt enable.
)

const (
	FAR AR_Bits = 0xFFFFFFFF << 0 //+ Flash Address.
)

const (
	OPTERR     OBR_Bits = 0x01 << 0 //+ Option Byte Error.
	RDPRT      OBR_Bits = 0x01 << 1 //+ Read protection.
	USER       OBR_Bits = 0xFF << 2 //+ User Option Bytes.
	WDG_SW     OBR_Bits = 0x01 << 2 //  WDG_SW.
	nRST_STOP  OBR_Bits = 0x02 << 2 //  nRST_STOP.
	nRST_STDBY OBR_Bits = 0x04 << 2 //  nRST_STDBY.
	BFB2       OBR_Bits = 0x08 << 2 //  BFB2.
)

const (
	WRP WRPR_Bits = 0xFFFFFFFF << 0 //+ Write Protect.
)