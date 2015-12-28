// Peripheral: USART_Periph  Universal Synchronous Asynchronous Receiver Transmitter.
// Instances:
//  USART2  mmap.USART2_BASE
//  USART3  mmap.USART3_BASE
//  UART4   mmap.UART4_BASE
//  UART5   mmap.UART5_BASE
//  USART1  mmap.USART1_BASE
// Registers:
//  0x00 16  SR
//  0x04 16  DR
//  0x08 16  BRR
//  0x0C 16  CR1
//  0x10 16  CR2
//  0x14 16  CR3
//  0x18 16  GTPR
// Import:
//  stm32/o/f10x_hd/mmap
package usart

const (
	PE   SR_Bits = 0x01 << 0 //+ Parity Error.
	FE   SR_Bits = 0x01 << 1 //+ Framing Error.
	NE   SR_Bits = 0x01 << 2 //+ Noise Error Flag.
	ORE  SR_Bits = 0x01 << 3 //+ OverRun Error.
	IDLE SR_Bits = 0x01 << 4 //+ IDLE line detected.
	RXNE SR_Bits = 0x01 << 5 //+ Read Data Register Not Empty.
	TC   SR_Bits = 0x01 << 6 //+ Transmission Complete.
	TXE  SR_Bits = 0x01 << 7 //+ Transmit Data Register Empty.
	LBD  SR_Bits = 0x01 << 8 //+ LIN Break Detection Flag.
	CTS  SR_Bits = 0x01 << 9 //+ CTS Flag.
)

const ()

const (
	DIV_Fraction BRR_Bits = 0x0F << 0  //+ Fraction of USARTDIV.
	DIV_Mantissa BRR_Bits = 0xFFF << 4 //+ Mantissa of USARTDIV.
)

const (
	SBK    CR1_Bits = 0x01 << 0  //+ Send Break.
	RWU    CR1_Bits = 0x01 << 1  //+ Receiver wakeup.
	RE     CR1_Bits = 0x01 << 2  //+ Receiver Enable.
	TE     CR1_Bits = 0x01 << 3  //+ Transmitter Enable.
	IDLEIE CR1_Bits = 0x01 << 4  //+ IDLE Interrupt Enable.
	RXNEIE CR1_Bits = 0x01 << 5  //+ RXNE Interrupt Enable.
	TCIE   CR1_Bits = 0x01 << 6  //+ Transmission Complete Interrupt Enable.
	TXEIE  CR1_Bits = 0x01 << 7  //+ PE Interrupt Enable.
	PEIE   CR1_Bits = 0x01 << 8  //+ PE Interrupt Enable.
	PS     CR1_Bits = 0x01 << 9  //+ Parity Selection.
	PCE    CR1_Bits = 0x01 << 10 //+ Parity Control Enable.
	WAKE   CR1_Bits = 0x01 << 11 //+ Wakeup method.
	M      CR1_Bits = 0x01 << 12 //+ Word length.
	UE     CR1_Bits = 0x01 << 13 //+ USART Enable.
	OVER8  CR1_Bits = 0x01 << 15 //+ USART Oversmapling 8-bits.
)

const (
	ADD    CR2_Bits = 0x0F << 0  //+ Address of the USART node.
	LBDL   CR2_Bits = 0x01 << 5  //+ LIN Break Detection Length.
	LBDIE  CR2_Bits = 0x01 << 6  //+ LIN Break Detection Interrupt Enable.
	LBCL   CR2_Bits = 0x01 << 8  //+ Last Bit Clock pulse.
	CPHA   CR2_Bits = 0x01 << 9  //+ Clock Phase.
	CPOL   CR2_Bits = 0x01 << 10 //+ Clock Polarity.
	CLKEN  CR2_Bits = 0x01 << 11 //+ Clock Enable.
	STOP   CR2_Bits = 0x03 << 12 //+ STOP[1:0] bits (STOP bits).
	STOP_0 CR2_Bits = 0x01 << 12 //  Bit 0.
	STOP_1 CR2_Bits = 0x02 << 12 //  Bit 1.
	LINEN  CR2_Bits = 0x01 << 14 //+ LIN mode enable.
)

const (
	EIE    CR3_Bits = 0x01 << 0  //+ Error Interrupt Enable.
	IREN   CR3_Bits = 0x01 << 1  //+ IrDA mode Enable.
	IRLP   CR3_Bits = 0x01 << 2  //+ IrDA Low-Power.
	HDSEL  CR3_Bits = 0x01 << 3  //+ Half-Duplex Selection.
	NACK   CR3_Bits = 0x01 << 4  //+ Smartcard NACK enable.
	SCEN   CR3_Bits = 0x01 << 5  //+ Smartcard mode enable.
	DMAR   CR3_Bits = 0x01 << 6  //+ DMA Enable Receiver.
	DMAT   CR3_Bits = 0x01 << 7  //+ DMA Enable Transmitter.
	RTSE   CR3_Bits = 0x01 << 8  //+ RTS Enable.
	CTSE   CR3_Bits = 0x01 << 9  //+ CTS Enable.
	CTSIE  CR3_Bits = 0x01 << 10 //+ CTS Interrupt Enable.
	ONEBIT CR3_Bits = 0x01 << 11 //+ One Bit method.
)

const (
	PSC   GTPR_Bits = 0xFF << 0 //+ PSC[7:0] bits (Prescaler value).
	PSC_0 GTPR_Bits = 0x01 << 0 //  Bit 0.
	PSC_1 GTPR_Bits = 0x02 << 0 //  Bit 1.
	PSC_2 GTPR_Bits = 0x04 << 0 //  Bit 2.
	PSC_3 GTPR_Bits = 0x08 << 0 //  Bit 3.
	PSC_4 GTPR_Bits = 0x10 << 0 //  Bit 4.
	PSC_5 GTPR_Bits = 0x20 << 0 //  Bit 5.
	PSC_6 GTPR_Bits = 0x40 << 0 //  Bit 6.
	PSC_7 GTPR_Bits = 0x80 << 0 //  Bit 7.
	GT    GTPR_Bits = 0xFF << 8 //+ Guard time value.
)