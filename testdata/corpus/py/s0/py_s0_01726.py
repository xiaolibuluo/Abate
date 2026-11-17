
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch790_0(gate):
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
def traced790_1(value: int) -> int:
    return value + 97


def decorate790_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate790_1
def doubled790_1(value: int) -> int:
    return value + value
def guarded790_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
