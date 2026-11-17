
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand258_0(values):
    for item in values:
        if item > 61:
            yield item * 46


def collect258_0(values):
    return list(expand258_0(values))
async def fetch258_1(gate):
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
def guarded258_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced258_3(value: int) -> int:
    return value + 82


def decorate258_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate258_3
def doubled258_3(value: int) -> int:
    return value + value
