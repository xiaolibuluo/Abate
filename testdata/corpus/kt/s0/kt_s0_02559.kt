package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21060<A, B>(val first: A, val second: B)

fun combine21060(left: Int, right: Int): Pair21060<Int, Int> {
    return Pair21060(left + 29, right + 50)
}
class Counter21061 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
sealed class Outcome21062 {
    data class Success(val payload: String) : Outcome21062()
    data object Failure : Outcome21062()
}

fun describe21062(outcome: Outcome21062): String = when (outcome) {
    is Outcome21062.Success -> outcome.payload
    Outcome21062.Failure -> "failed"
}
