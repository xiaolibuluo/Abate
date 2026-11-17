package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter6940 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
data class Pair6941<A, B>(val first: A, val second: B)

fun combine6941(left: Int, right: Int): Pair6941<Int, Int> {
    return Pair6941(left + 13, right + 59)
}
sealed class Outcome6942 {
    data class Success(val payload: String) : Outcome6942()
    data object Failure : Outcome6942()
}

fun describe6942(outcome: Outcome6942): String = when (outcome) {
    is Outcome6942.Success -> outcome.payload
    Outcome6942.Failure -> "failed"
}
