package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6100 {
    data class Success(val payload: String) : Outcome6100()
    data object Failure : Outcome6100()
}

fun describe6100(outcome: Outcome6100): String = when (outcome) {
    is Outcome6100.Success -> outcome.payload
    Outcome6100.Failure -> "failed"
}
data class Pair6101<A, B>(val first: A, val second: B)

fun combine6101(left: Int, right: Int): Pair6101<Int, Int> {
    return Pair6101(left + 85, right + 20)
}
sealed class Outcome6102 {
    data class Success(val payload: String) : Outcome6102()
    data object Failure : Outcome6102()
}

fun describe6102(outcome: Outcome6102): String = when (outcome) {
    is Outcome6102.Success -> outcome.payload
    Outcome6102.Failure -> "failed"
}
object Registry6103 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6103 = "registry-99"
    }
}
class Counter6104 {
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
