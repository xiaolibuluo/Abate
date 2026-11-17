
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2397_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2397_1(value: int) -> int:
    return value + 36


def decorate2397_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2397_1
def doubled2397_1(value: int) -> int:
    return value + value
def expand2397_2(values):
    for item in values:
        if item > 43:
            yield item * 35


def collect2397_2(values):
    return list(expand2397_2(values))
async def fetch2397_3(gate):
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
