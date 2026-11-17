
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2448_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def expand2448_1(values):
    for item in values:
        if item > 20:
            yield item * 15


def collect2448_1(values):
    return list(expand2448_1(values))
async def fetch2448_2(gate):
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
def traced2448_3(value: int) -> int:
    return value + 34


def decorate2448_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2448_3
def doubled2448_3(value: int) -> int:
    return value + value
