
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2367_0(values):
    for item in values:
        if item > 27:
            yield item * 47


def collect2367_0(values):
    return list(expand2367_0(values))
def traced2367_1(value: int) -> int:
    return value + 34


def decorate2367_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2367_1
def doubled2367_1(value: int) -> int:
    return value + value
def guarded2367_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2367_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 48}
def guarded2367_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2367_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 62}
async def fetch2367_6(gate):
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
