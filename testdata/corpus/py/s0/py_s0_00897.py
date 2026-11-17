
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project410_0(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def project410_1(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def summarize410_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced410_3(value: int) -> int:
    return value + 21


def decorate410_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate410_3
def doubled410_3(value: int) -> int:
    return value + value
def expand410_4(values):
    for item in values:
        if item > 28:
            yield item * 56


def collect410_4(values):
    return list(expand410_4(values))
async def fetch410_5(gate):
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
