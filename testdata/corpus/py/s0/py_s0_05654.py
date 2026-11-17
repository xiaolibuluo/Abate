
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2604_0(gate):
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
def expand2604_1(values):
    for item in values:
        if item > 11:
            yield item * 9


def collect2604_1(values):
    return list(expand2604_1(values))
def expand2604_2(values):
    for item in values:
        if item > 18:
            yield item * 20


def collect2604_2(values):
    return list(expand2604_2(values))
def pair2604_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 25}
def project2604_4(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def traced2604_5(value: int) -> int:
    return value + 39


def decorate2604_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2604_5
def doubled2604_5(value: int) -> int:
    return value + value
