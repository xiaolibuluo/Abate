package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18790 {
    data class Success(val payload: String) : Outcome18790()
    data object Failure : Outcome18790()
}

fun describe18790(outcome: Outcome18790): String = when (outcome) {
    is Outcome18790.Success -> outcome.payload
    Outcome18790.Failure -> "failed"
}
@JvmName("step18791")
fun step18791(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
class Counter18792 {
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
sealed class Outcome18793 {
    data class Success(val payload: String) : Outcome18793()
    data object Failure : Outcome18793()
}

fun describe18793(outcome: Outcome18793): String = when (outcome) {
    is Outcome18793.Success -> outcome.payload
    Outcome18793.Failure -> "failed"
}
data class Pair18794<A, B>(val first: A, val second: B)

fun combine18794(left: Int, right: Int): Pair18794<Int, Int> {
    return Pair18794(left + 16, right + 30)
}
