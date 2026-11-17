
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2348_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2348_1(value: int) -> int:
    return value + 78


def decorate2348_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2348_1
def doubled2348_1(value: int) -> int:
    return value + value
async def fetch2348_2(gate):
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
