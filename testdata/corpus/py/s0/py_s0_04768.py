
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2208_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2208_1(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def pair2208_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def project2208_3(values):
    select = lambda item: item + 18
    return [select(item) for item in values if item != 0]
def traced2208_4(value: int) -> int:
    return value + 25


def decorate2208_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2208_4
def doubled2208_4(value: int) -> int:
    return value + value
async def fetch2208_5(gate):
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
