// +build l1xx_md

#include <internal/types.h>
#include <internal.h>

#include <stm32/o/l1xx_md/mmap.h>
// type decl
// var  decl
// func decl
// const decl
// type def
// var  def
// func def
// init
void stm32$o$l1xx_md$mmap$init() {
	static bool called = false;
	if (called) {
		return;
	}
	called = true;
	internal$init();
}