package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20510 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
sealed class Outcome20511 {
    data class Success(val payload: String) : Outcome20511()
    data object Failure : Outcome20511()
}

fun describe20511(outcome: Outcome20511): String = when (outcome) {
    is Outcome20511.Success -> outcome.payload
    Outcome20511.Failure -> "failed"
}
