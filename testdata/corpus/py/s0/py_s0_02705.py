
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1236_0(gate):
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
def expand1236_1(values):
    for item in values:
        if item > 75:
            yield item * 10


def collect1236_1(values):
    return list(expand1236_1(values))
def pair1236_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 82}
def traced1236_3(value: int) -> int:
    return value + 89


def decorate1236_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1236_3
def doubled1236_3(value: int) -> int:
    return value + value
def traced1236_4(value: int) -> int:
    return value + 96


def decorate1236_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1236_4
def doubled1236_4(value: int) -> int:
    return value + value
