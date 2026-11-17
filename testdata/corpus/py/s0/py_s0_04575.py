
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2123_0(value: int) -> int:
    return value + 56


def decorate2123_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2123_0
def doubled2123_0(value: int) -> int:
    return value + value
async def fetch2123_1(gate):
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
def pair2123_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
def guarded2123_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2123_4(gate):
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
def traced2123_5(value: int) -> int:
    return value + 91


def decorate2123_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2123_5
def doubled2123_5(value: int) -> int:
    return value + value
