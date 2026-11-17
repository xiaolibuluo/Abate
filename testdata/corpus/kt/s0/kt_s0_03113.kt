package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25550 {
    data class Success(val payload: String) : Outcome25550()
    data object Failure : Outcome25550()
}

fun describe25550(outcome: Outcome25550): String = when (outcome) {
    is Outcome25550.Success -> outcome.payload
    Outcome25550.Failure -> "failed"
}
sealed class Outcome25551 {
    data class Success(val payload: String) : Outcome25551()
    data object Failure : Outcome25551()
}

fun describe25551(outcome: Outcome25551): String = when (outcome) {
    is Outcome25551.Success -> outcome.payload
    Outcome25551.Failure -> "failed"
}
class Counter25552 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
data class Pair25553<A, B>(val first: A, val second: B)

fun combine25553(left: Int, right: Int): Pair25553<Int, Int> {
    return Pair25553(left + 67, right + 54)
}
