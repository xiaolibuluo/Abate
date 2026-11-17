
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2364_0(values):
    for item in values:
        if item > 85:
            yield item * 26


def collect2364_0(values):
    return list(expand2364_0(values))
async def fetch2364_1(gate):
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
def traced2364_2(value: int) -> int:
    return value + 99


def decorate2364_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2364_2
def doubled2364_2(value: int) -> int:
    return value + value
def guarded2364_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2364_4(value: int) -> int:
    return value + 16


def decorate2364_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2364_4
def doubled2364_4(value: int) -> int:
    return value + value
def expand2364_5(values):
    for item in values:
        if item > 23:
            yield item * 20


def collect2364_5(values):
    return list(expand2364_5(values))
