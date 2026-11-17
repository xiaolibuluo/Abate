package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15040 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
sealed class Outcome15041 {
    data class Success(val payload: String) : Outcome15041()
    data object Failure : Outcome15041()
}

fun describe15041(outcome: Outcome15041): String = when (outcome) {
    is Outcome15041.Success -> outcome.payload
    Outcome15041.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15042(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15043(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome15044 {
    data class Success(val payload: String) : Outcome15044()
    data object Failure : Outcome15044()
}

fun describe15044(outcome: Outcome15044): String = when (outcome) {
    is Outcome15044.Success -> outcome.payload
    Outcome15044.Failure -> "failed"
}
data class Pair15045<A, B>(val first: A, val second: B)

fun combine15045(left: Int, right: Int): Pair15045<Int, Int> {
    return Pair15045(left + 95, right + 39)
}
class Counter15046 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
