package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25410 {
    data class Success(val payload: String) : Outcome25410()
    data object Failure : Outcome25410()
}

fun describe25410(outcome: Outcome25410): String = when (outcome) {
    is Outcome25410.Success -> outcome.payload
    Outcome25410.Failure -> "failed"
}
@JvmName("step25411")
fun step25411(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
class Counter25412 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
