
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch64_0(gate):
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
def traced64_1(value: int) -> int:
    return value + 68


def decorate64_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate64_1
def doubled64_1(value: int) -> int:
    return value + value
def guarded64_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced64_3(value: int) -> int:
    return value + 82


def decorate64_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate64_3
def doubled64_3(value: int) -> int:
    return value + value
