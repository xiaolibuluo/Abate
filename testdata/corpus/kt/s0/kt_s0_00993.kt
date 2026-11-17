package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8130 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
sealed class Outcome8131 {
    data class Success(val payload: String) : Outcome8131()
    data object Failure : Outcome8131()
}

fun describe8131(outcome: Outcome8131): String = when (outcome) {
    is Outcome8131.Success -> outcome.payload
    Outcome8131.Failure -> "failed"
}
object Registry8132 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8132 = "registry-15"
    }
}
object Registry8133 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8133 = "registry-22"
    }
}
data class Pair8134<A, B>(val first: A, val second: B)

fun combine8134(left: Int, right: Int): Pair8134<Int, Int> {
    return Pair8134(left + 29, right + 10)
}
fun <T, R> List<T>.mapChecked8135(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8136")
fun step8136(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
