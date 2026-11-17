
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2642_0(value: int) -> int:
    return value + 13


def decorate2642_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2642_0
def doubled2642_0(value: int) -> int:
    return value + value
async def fetch2642_1(gate):
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
def pair2642_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
def traced2642_3(value: int) -> int:
    return value + 34


def decorate2642_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2642_3
def doubled2642_3(value: int) -> int:
    return value + value
