package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25780 {
    data class Success(val payload: String) : Outcome25780()
    data object Failure : Outcome25780()
}

fun describe25780(outcome: Outcome25780): String = when (outcome) {
    is Outcome25780.Success -> outcome.payload
    Outcome25780.Failure -> "failed"
}
sealed class Outcome25781 {
    data class Success(val payload: String) : Outcome25781()
    data object Failure : Outcome25781()
}

fun describe25781(outcome: Outcome25781): String = when (outcome) {
    is Outcome25781.Success -> outcome.payload
    Outcome25781.Failure -> "failed"
}
sealed class Outcome25782 {
    data class Success(val payload: String) : Outcome25782()
    data object Failure : Outcome25782()
}

fun describe25782(outcome: Outcome25782): String = when (outcome) {
    is Outcome25782.Success -> outcome.payload
    Outcome25782.Failure -> "failed"
}
class Counter25783 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
data class Pair25784<A, B>(val first: A, val second: B)

fun combine25784(left: Int, right: Int): Pair25784<Int, Int> {
    return Pair25784(left + 82, right + 43)
}
sealed class Outcome25785 {
    data class Success(val payload: String) : Outcome25785()
    data object Failure : Outcome25785()
}

fun describe25785(outcome: Outcome25785): String = when (outcome) {
    is Outcome25785.Success -> outcome.payload
    Outcome25785.Failure -> "failed"
}
sealed class Outcome25786 {
    data class Success(val payload: String) : Outcome25786()
    data object Failure : Outcome25786()
}

fun describe25786(outcome: Outcome25786): String = when (outcome) {
    is Outcome25786.Success -> outcome.payload
    Outcome25786.Failure -> "failed"
}
