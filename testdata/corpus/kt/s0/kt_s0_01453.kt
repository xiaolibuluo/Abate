package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12000 {
    data class Success(val payload: String) : Outcome12000()
    data object Failure : Outcome12000()
}

fun describe12000(outcome: Outcome12000): String = when (outcome) {
    is Outcome12000.Success -> outcome.payload
    Outcome12000.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12001(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter12002 {
    var current: Int by Delegates.observable(99) { _, _, new -> println(new) }

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
data class Pair12003<A, B>(val first: A, val second: B)

fun combine12003(left: Int, right: Int): Pair12003<Int, Int> {
    return Pair12003(left + 9, right + 24)
}
sealed class Outcome12004 {
    data class Success(val payload: String) : Outcome12004()
    data object Failure : Outcome12004()
}

fun describe12004(outcome: Outcome12004): String = when (outcome) {
    is Outcome12004.Success -> outcome.payload
    Outcome12004.Failure -> "failed"
}
@JvmName("step12005")
fun step12005(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
