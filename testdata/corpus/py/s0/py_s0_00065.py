
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch33_0(gate):
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
async def fetch33_1(gate):
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
def traced33_2(value: int) -> int:
    return value + 60


def decorate33_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate33_2
def doubled33_2(value: int) -> int:
    return value + value
async def fetch33_3(gate):
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
def project33_4(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def traced33_5(value: int) -> int:
    return value + 81


def decorate33_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate33_5
def doubled33_5(value: int) -> int:
    return value + value
def summarize33_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
