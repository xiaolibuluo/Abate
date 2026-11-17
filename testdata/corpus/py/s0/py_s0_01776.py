
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced811_0(value: int) -> int:
    return value + 72


def decorate811_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate811_0
def doubled811_0(value: int) -> int:
    return value + value
async def fetch811_1(gate):
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
def project811_2(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
async def fetch811_3(gate):
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
def summarize811_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch811_5(gate):
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
