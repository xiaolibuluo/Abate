
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced106_0(value: int) -> int:
    return value + 25


def decorate106_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate106_0
def doubled106_0(value: int) -> int:
    return value + value
def traced106_1(value: int) -> int:
    return value + 32


def decorate106_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate106_1
def doubled106_1(value: int) -> int:
    return value + value
async def fetch106_2(gate):
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
