package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22860 {
    data class Success(val payload: String) : Outcome22860()
    data object Failure : Outcome22860()
}

fun describe22860(outcome: Outcome22860): String = when (outcome) {
    is Outcome22860.Success -> outcome.payload
    Outcome22860.Failure -> "failed"
}
@JvmName("step22861")
fun step22861(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
data class Pair22862<A, B>(val first: A, val second: B)

fun combine22862(left: Int, right: Int): Pair22862<Int, Int> {
    return Pair22862(left + 55, right + 51)
}
fun <T, R> List<T>.mapChecked22863(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22864 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
class Counter22865 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step22866")
fun step22866(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
sealed class Outcome22867 {
    data class Success(val payload: String) : Outcome22867()
    data object Failure : Outcome22867()
}

fun describe22867(outcome: Outcome22867): String = when (outcome) {
    is Outcome22867.Success -> outcome.payload
    Outcome22867.Failure -> "failed"
}
