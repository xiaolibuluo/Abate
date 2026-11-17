
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2225_0(gate):
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
def traced2225_1(value: int) -> int:
    return value + 31


def decorate2225_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2225_1
def doubled2225_1(value: int) -> int:
    return value + value
def expand2225_2(values):
    for item in values:
        if item > 38:
            yield item * 51


def collect2225_2(values):
    return list(expand2225_2(values))
def guarded2225_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2225_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 52}
def expand2225_5(values):
    for item in values:
        if item > 59:
            yield item * 23


def collect2225_5(values):
    return list(expand2225_5(values))
async def fetch2225_6(gate):
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
def expand2225_7(values):
    for item in values:
        if item > 73:
            yield item * 45


def collect2225_7(values):
    return list(expand2225_7(values))
