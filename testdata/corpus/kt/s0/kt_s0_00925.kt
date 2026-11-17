package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7570 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7570 = "registry-49"
    }
}
data class Pair7571<A, B>(val first: A, val second: B)

fun combine7571(left: Int, right: Int): Pair7571<Int, Int> {
    return Pair7571(left + 56, right + 12)
}
class Counter7572 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
data class Pair7573<A, B>(val first: A, val second: B)

fun combine7573(left: Int, right: Int): Pair7573<Int, Int> {
    return Pair7573(left + 70, right + 34)
}
sealed class Outcome7574 {
    data class Success(val payload: String) : Outcome7574()
    data object Failure : Outcome7574()
}

fun describe7574(outcome: Outcome7574): String = when (outcome) {
    is Outcome7574.Success -> outcome.payload
    Outcome7574.Failure -> "failed"
}
