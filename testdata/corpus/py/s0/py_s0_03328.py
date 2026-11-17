
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1538_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1538_1(values):
    for item in values:
        if item > 24:
            yield item * 50


def collect1538_1(values):
    return list(expand1538_1(values))
def traced1538_2(value: int) -> int:
    return value + 31


def decorate1538_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1538_2
def doubled1538_2(value: int) -> int:
    return value + value
def project1538_3(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
async def fetch1538_4(gate):
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
async def fetch1538_5(gate):
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
