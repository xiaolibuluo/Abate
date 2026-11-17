package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12280 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
sealed class Outcome12281 {
    data class Success(val payload: String) : Outcome12281()
    data object Failure : Outcome12281()
}

fun describe12281(outcome: Outcome12281): String = when (outcome) {
    is Outcome12281.Success -> outcome.payload
    Outcome12281.Failure -> "failed"
}
