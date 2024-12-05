import subprocess as sp
import platform
import time
import sys as os

sys = platform.system()

def loading(proc):
    d = ""
    while proc.poll() is None:
        print(f"\rBuilding compiler{d}", end="", flush=True)
        d += "."
        if len(d) > 3:
            d = ""
        time.sleep(0.1)
    print(f"\rBuilding compiler... Done!")
    pass
pass

if sys == "Linux" or sys == "Darwin":
    try:
        cmd = sp.Popen([
            "go", "build", "-o", "lunac"
        ], cwd="compiler")
    except Exception as err:
        print(f"\033[1;31mError!\033[0;91m Compilation not successful! make sure you have installed all dependencies\033[3m - ErrorCode: {err}\033[0;0m")
        os.exit(249)
        pass
    loading(cmd)
    pass
pass

if sys == "Windows":
    try:
        cmd = sp.Popen([
            "go", "build", "-o", "lunac"
        ], cwd="compiler")
    except Exception as err:
        print(f"Compilation not successful! make sure you have installed all dependencies - ErrorCode: {err}")
        os.exit(305)
        pass
    loading(cmd)
    pass
pass
