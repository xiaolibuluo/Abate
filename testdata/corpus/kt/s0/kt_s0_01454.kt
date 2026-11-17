package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12010 {
    data class Success(val payload: String) : Outcome12010()
    data object Failure : Outcome12010()
}

fun describe12010(outcome: Outcome12010): String = when (outcome) {
    is Outcome12010.Success -> outcome.payload
    Outcome12010.Failure -> "failed"
}
sealed class Outcome12011 {
    data class Success(val payload: String) : Outcome12011()
    data object Failure : Outcome12011()
}

fun describe12011(outcome: Outcome12011): String = when (outcome) {
    is Outcome12011.Success -> outcome.payload
    Outcome12011.Failure -> "failed"
}
data class Pair12012<A, B>(val first: A, val second: B)

fun combine12012(left: Int, right: Int): Pair12012<Int, Int> {
    return Pair12012(left + 15, right + 20)
}
class Counter12013 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
sealed class Outcome12014 {
    data class Success(val payload: String) : Outcome12014()
    data object Failure : Outcome12014()
}

fun describe12014(outcome: Outcome12014): String = when (outcome) {
    is Outcome12014.Success -> outcome.payload
    Outcome12014.Failure -> "failed"
}
