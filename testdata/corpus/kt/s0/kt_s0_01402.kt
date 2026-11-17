package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11540<A, B>(val first: A, val second: B)

fun combine11540(left: Int, right: Int): Pair11540<Int, Int> {
    return Pair11540(left + 69, right + 35)
}
class Counter11541 {
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
fun <T, R> List<T>.mapChecked11542(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11543 {
    data class Success(val payload: String) : Outcome11543()
    data object Failure : Outcome11543()
}

fun describe11543(outcome: Outcome11543): String = when (outcome) {
    is Outcome11543.Success -> outcome.payload
    Outcome11543.Failure -> "failed"
}
@JvmName("step11544")
fun step11544(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
sealed class Outcome11545 {
    data class Success(val payload: String) : Outcome11545()
    data object Failure : Outcome11545()
}

fun describe11545(outcome: Outcome11545): String = when (outcome) {
    is Outcome11545.Success -> outcome.payload
    Outcome11545.Failure -> "failed"
}
