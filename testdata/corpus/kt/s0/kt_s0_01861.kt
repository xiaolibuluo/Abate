package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15290<A, B>(val first: A, val second: B)

fun combine15290(left: Int, right: Int): Pair15290<Int, Int> {
    return Pair15290(left + 94, right + 37)
}
sealed class Outcome15291 {
    data class Success(val payload: String) : Outcome15291()
    data object Failure : Outcome15291()
}

fun describe15291(outcome: Outcome15291): String = when (outcome) {
    is Outcome15291.Success -> outcome.payload
    Outcome15291.Failure -> "failed"
}
class Counter15292 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
data class Pair15293<A, B>(val first: A, val second: B)

fun combine15293(left: Int, right: Int): Pair15293<Int, Int> {
    return Pair15293(left + 18, right + 9)
}
