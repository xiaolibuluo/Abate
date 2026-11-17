
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced935_0(value: int) -> int:
    return value + 35


def decorate935_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate935_0
def doubled935_0(value: int) -> int:
    return value + value
async def fetch935_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
async def fetch935_2(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
