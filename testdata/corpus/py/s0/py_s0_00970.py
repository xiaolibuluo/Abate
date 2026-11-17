
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced450_0(value: int) -> int:
    return value + 35


def decorate450_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate450_0
def doubled450_0(value: int) -> int:
    return value + value
async def fetch450_1(gate):
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
def traced450_2(value: int) -> int:
    return value + 49


def decorate450_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate450_2
def doubled450_2(value: int) -> int:
    return value + value
