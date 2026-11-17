package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2250 {
    data class Success(val payload: String) : Outcome2250()
    data object Failure : Outcome2250()
}

fun describe2250(outcome: Outcome2250): String = when (outcome) {
    is Outcome2250.Success -> outcome.payload
    Outcome2250.Failure -> "failed"
}
class Counter2251 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
data class Pair2252<A, B>(val first: A, val second: B)

fun combine2252(left: Int, right: Int): Pair2252<Int, Int> {
    return Pair2252(left + 34, right + 20)
}
sealed class Outcome2253 {
    data class Success(val payload: String) : Outcome2253()
    data object Failure : Outcome2253()
}

fun describe2253(outcome: Outcome2253): String = when (outcome) {
    is Outcome2253.Success -> outcome.payload
    Outcome2253.Failure -> "failed"
}
class Counter2254 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
