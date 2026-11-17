package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3190 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3190 = "registry-78"
    }
}
data class Pair3191<A, B>(val first: A, val second: B)

fun combine3191(left: Int, right: Int): Pair3191<Int, Int> {
    return Pair3191(left + 85, right + 57)
}
sealed class Outcome3192 {
    data class Success(val payload: String) : Outcome3192()
    data object Failure : Outcome3192()
}

fun describe3192(outcome: Outcome3192): String = when (outcome) {
    is Outcome3192.Success -> outcome.payload
    Outcome3192.Failure -> "failed"
}
object Registry3193 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3193 = "registry-99"
    }
}
class Counter3194 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
