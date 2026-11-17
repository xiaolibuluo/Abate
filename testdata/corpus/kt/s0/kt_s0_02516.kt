package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20680 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
sealed class Outcome20681 {
    data class Success(val payload: String) : Outcome20681()
    data object Failure : Outcome20681()
}

fun describe20681(outcome: Outcome20681): String = when (outcome) {
    is Outcome20681.Success -> outcome.payload
    Outcome20681.Failure -> "failed"
}
data class Pair20682<A, B>(val first: A, val second: B)

fun combine20682(left: Int, right: Int): Pair20682<Int, Int> {
    return Pair20682(left + 34, right + 50)
}
sealed class Outcome20683 {
    data class Success(val payload: String) : Outcome20683()
    data object Failure : Outcome20683()
}

fun describe20683(outcome: Outcome20683): String = when (outcome) {
    is Outcome20683.Success -> outcome.payload
    Outcome20683.Failure -> "failed"
}
