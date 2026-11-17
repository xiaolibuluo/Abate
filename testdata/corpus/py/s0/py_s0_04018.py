
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1865_0(gate):
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
async def fetch1865_1(gate):
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
def traced1865_2(value: int) -> int:
    return value + 14


def decorate1865_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1865_2
def doubled1865_2(value: int) -> int:
    return value + value
def project1865_3(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
