
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2157_0(values):
    for item in values:
        if item > 13:
            yield item * 41


def collect2157_0(values):
    return list(expand2157_0(values))
def guarded2157_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2157_2(value: int) -> int:
    return value + 27


def decorate2157_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2157_2
def doubled2157_2(value: int) -> int:
    return value + value
async def fetch2157_3(gate):
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
