
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair664_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 101}
async def fetch664_1(gate):
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
def pair664_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def traced664_3(value: int) -> int:
    return value + 25


def decorate664_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate664_3
def doubled664_3(value: int) -> int:
    return value + value
async def fetch664_4(gate):
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
def traced664_5(value: int) -> int:
    return value + 39


def decorate664_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate664_5
def doubled664_5(value: int) -> int:
    return value + value
def traced664_6(value: int) -> int:
    return value + 46


def decorate664_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate664_6
def doubled664_6(value: int) -> int:
    return value + value
