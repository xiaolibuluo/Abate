package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16780<A, B>(val first: A, val second: B)

fun combine16780(left: Int, right: Int): Pair16780<Int, Int> {
    return Pair16780(left + 91, right + 43)
}
sealed class Outcome16781 {
    data class Success(val payload: String) : Outcome16781()
    data object Failure : Outcome16781()
}

fun describe16781(outcome: Outcome16781): String = when (outcome) {
    is Outcome16781.Success -> outcome.payload
    Outcome16781.Failure -> "failed"
}
object Registry16782 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16782 = "registry-8"
    }
}
sealed class Outcome16783 {
    data class Success(val payload: String) : Outcome16783()
    data object Failure : Outcome16783()
}

fun describe16783(outcome: Outcome16783): String = when (outcome) {
    is Outcome16783.Success -> outcome.payload
    Outcome16783.Failure -> "failed"
}
class Counter16784 {
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
