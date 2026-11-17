
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1820_0(values):
    for item in values:
        if item > 94:
            yield item * 61


def collect1820_0(values):
    return list(expand1820_0(values))
def traced1820_1(value: int) -> int:
    return value + 101


def decorate1820_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1820_1
def doubled1820_1(value: int) -> int:
    return value + value
async def fetch1820_2(gate):
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
def summarize1820_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
