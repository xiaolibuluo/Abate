package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10440 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
data class Pair10441<A, B>(val first: A, val second: B)

fun combine10441(left: Int, right: Int): Pair10441<Int, Int> {
    return Pair10441(left + 101, right + 69)
}
sealed class Outcome10442 {
    data class Success(val payload: String) : Outcome10442()
    data object Failure : Outcome10442()
}

fun describe10442(outcome: Outcome10442): String = when (outcome) {
    is Outcome10442.Success -> outcome.payload
    Outcome10442.Failure -> "failed"
}
sealed class Outcome10443 {
    data class Success(val payload: String) : Outcome10443()
    data object Failure : Outcome10443()
}

fun describe10443(outcome: Outcome10443): String = when (outcome) {
    is Outcome10443.Success -> outcome.payload
    Outcome10443.Failure -> "failed"
}
