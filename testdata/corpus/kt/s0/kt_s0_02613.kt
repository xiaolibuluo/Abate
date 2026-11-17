package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21550 {
    data class Success(val payload: String) : Outcome21550()
    data object Failure : Outcome21550()
}

fun describe21550(outcome: Outcome21550): String = when (outcome) {
    is Outcome21550.Success -> outcome.payload
    Outcome21550.Failure -> "failed"
}
class Counter21551 {
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
data class Pair21552<A, B>(val first: A, val second: B)

fun combine21552(left: Int, right: Int): Pair21552<Int, Int> {
    return Pair21552(left + 98, right + 49)
}
