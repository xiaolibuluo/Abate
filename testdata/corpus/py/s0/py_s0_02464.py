
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1140_0(gate):
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
async def fetch1140_1(gate):
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
def traced1140_2(value: int) -> int:
    return value + 95


def decorate1140_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1140_2
def doubled1140_2(value: int) -> int:
    return value + value
def traced1140_3(value: int) -> int:
    return value + 5


def decorate1140_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1140_3
def doubled1140_3(value: int) -> int:
    return value + value
