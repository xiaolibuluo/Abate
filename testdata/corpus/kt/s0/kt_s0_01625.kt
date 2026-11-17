package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13420 {
    data class Success(val payload: String) : Outcome13420()
    data object Failure : Outcome13420()
}

fun describe13420(outcome: Outcome13420): String = when (outcome) {
    is Outcome13420.Success -> outcome.payload
    Outcome13420.Failure -> "failed"
}
data class Pair13421<A, B>(val first: A, val second: B)

fun combine13421(left: Int, right: Int): Pair13421<Int, Int> {
    return Pair13421(left + 95, right + 20)
}
object Registry13422 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13422 = "registry-5"
    }
}
sealed class Outcome13423 {
    data class Success(val payload: String) : Outcome13423()
    data object Failure : Outcome13423()
}

fun describe13423(outcome: Outcome13423): String = when (outcome) {
    is Outcome13423.Success -> outcome.payload
    Outcome13423.Failure -> "failed"
}
object Registry13424 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13424 = "registry-19"
    }
}
class Counter13425 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
