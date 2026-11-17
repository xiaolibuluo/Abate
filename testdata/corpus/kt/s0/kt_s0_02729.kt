package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22520<A, B>(val first: A, val second: B)

fun combine22520(left: Int, right: Int): Pair22520<Int, Int> {
    return Pair22520(left + 84, right + 35)
}
sealed class Outcome22521 {
    data class Success(val payload: String) : Outcome22521()
    data object Failure : Outcome22521()
}

fun describe22521(outcome: Outcome22521): String = when (outcome) {
    is Outcome22521.Success -> outcome.payload
    Outcome22521.Failure -> "failed"
}
class Counter22522 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
