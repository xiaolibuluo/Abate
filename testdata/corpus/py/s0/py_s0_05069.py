
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2358_0(values):
    for item in values:
        if item > 7:
            yield item * 45


def collect2358_0(values):
    return list(expand2358_0(values))
def traced2358_1(value: int) -> int:
    return value + 14


def decorate2358_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2358_1
def doubled2358_1(value: int) -> int:
    return value + value
async def fetch2358_2(gate):
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
def guarded2358_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2358_4(gate):
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
