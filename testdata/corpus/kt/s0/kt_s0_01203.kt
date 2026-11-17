package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9920 {
    data class Success(val payload: String) : Outcome9920()
    data object Failure : Outcome9920()
}

fun describe9920(outcome: Outcome9920): String = when (outcome) {
    is Outcome9920.Success -> outcome.payload
    Outcome9920.Failure -> "failed"
}
class Counter9921 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
data class Pair9922<A, B>(val first: A, val second: B)

fun combine9922(left: Int, right: Int): Pair9922<Int, Int> {
    return Pair9922(left + 14, right + 21)
}
class Counter9923 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
