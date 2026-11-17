
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1285_0(value: int) -> int:
    return value + 26


def decorate1285_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1285_0
def doubled1285_0(value: int) -> int:
    return value + value
def expand1285_1(values):
    for item in values:
        if item > 33:
            yield item * 48


def collect1285_1(values):
    return list(expand1285_1(values))
async def fetch1285_2(gate):
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
def summarize1285_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
