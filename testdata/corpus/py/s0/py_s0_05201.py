
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2421_0(gate):
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
def traced2421_1(value: int) -> int:
    return value + 57


def decorate2421_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2421_1
def doubled2421_1(value: int) -> int:
    return value + value
def traced2421_2(value: int) -> int:
    return value + 64


def decorate2421_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2421_2
def doubled2421_2(value: int) -> int:
    return value + value
def summarize2421_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2421_4(value: int) -> int:
    return value + 78


def decorate2421_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2421_4
def doubled2421_4(value: int) -> int:
    return value + value
def expand2421_5(values):
    for item in values:
        if item > 85:
            yield item * 53


def collect2421_5(values):
    return list(expand2421_5(values))
def pair2421_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 92}
