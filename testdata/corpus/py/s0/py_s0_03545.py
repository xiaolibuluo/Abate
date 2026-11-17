
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1634_0(value: int) -> int:
    return value + 101


def decorate1634_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1634_0
def doubled1634_0(value: int) -> int:
    return value + value
def traced1634_1(value: int) -> int:
    return value + 11


def decorate1634_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1634_1
def doubled1634_1(value: int) -> int:
    return value + value
def expand1634_2(values):
    for item in values:
        if item > 18:
            yield item * 62


def collect1634_2(values):
    return list(expand1634_2(values))
def traced1634_3(value: int) -> int:
    return value + 25


def decorate1634_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1634_3
def doubled1634_3(value: int) -> int:
    return value + value
async def fetch1634_4(gate):
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
def summarize1634_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
