# ------------------------------------------- MAKE ALL ----------------------------------------------- #
#                                    official part of LunaScript                                       #
#                                        Authors: LunaTeam                                             #
# ---------------------------------------------------------------------------------------------------- #

import os

size = os.get_terminal_size()
s = size.columns

print(f"{' Building compiler ':-^{s}}")
from make_compiler import erro
print()
print(f"{' Building runtime ':-^{s}}")
from make_runtime import erro
print()
print("Done")
