package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7640")
fun step7640(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
fun <T, R> List<T>.mapChecked7641(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step7642")
fun step7642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
data class Pair7643<A, B>(val first: A, val second: B)

fun combine7643(left: Int, right: Int): Pair7643<Int, Int> {
    return Pair7643(left + 64, right + 22)
}
sealed class Outcome7644 {
    data class Success(val payload: String) : Outcome7644()
    data object Failure : Outcome7644()
}

fun describe7644(outcome: Outcome7644): String = when (outcome) {
    is Outcome7644.Success -> outcome.payload
    Outcome7644.Failure -> "failed"
}
class Counter7645 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked7646(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
