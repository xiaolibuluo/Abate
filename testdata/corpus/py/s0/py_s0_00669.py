
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch309_0(gate):
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
def expand309_1(values):
    for item in values:
        if item > 52:
            yield item * 48


def collect309_1(values):
    return list(expand309_1(values))
def traced309_2(value: int) -> int:
    return value + 59


def decorate309_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate309_2
def doubled309_2(value: int) -> int:
    return value + value
def pair309_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 66}
