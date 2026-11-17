package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10880 {
    data class Success(val payload: String) : Outcome10880()
    data object Failure : Outcome10880()
}

fun describe10880(outcome: Outcome10880): String = when (outcome) {
    is Outcome10880.Success -> outcome.payload
    Outcome10880.Failure -> "failed"
}
class Counter10881 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
data class Pair10882<A, B>(val first: A, val second: B)

fun combine10882(left: Int, right: Int): Pair10882<Int, Int> {
    return Pair10882(left + 98, right + 22)
}
