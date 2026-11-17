
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project962_0(values):
    select = lambda item: item + 95
    return [select(item) for item in values if item != 0]
async def fetch962_1(gate):
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
def project962_2(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def traced962_3(value: int) -> int:
    return value + 19


def decorate962_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate962_3
def doubled962_3(value: int) -> int:
    return value + value
def traced962_4(value: int) -> int:
    return value + 26


def decorate962_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate962_4
def doubled962_4(value: int) -> int:
    return value + value
def project962_5(values):
    select = lambda item: item + 33
    return [select(item) for item in values if item != 0]
def summarize962_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize962_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
