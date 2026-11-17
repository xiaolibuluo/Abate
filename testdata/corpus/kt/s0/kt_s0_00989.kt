package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8110 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8110 = "registry-72"
    }
}
sealed class Outcome8111 {
    data class Success(val payload: String) : Outcome8111()
    data object Failure : Outcome8111()
}

fun describe8111(outcome: Outcome8111): String = when (outcome) {
    is Outcome8111.Success -> outcome.payload
    Outcome8111.Failure -> "failed"
}
@JvmName("step8112")
fun step8112(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
data class Pair8113<A, B>(val first: A, val second: B)

fun combine8113(left: Int, right: Int): Pair8113<Int, Int> {
    return Pair8113(left + 93, right + 46)
}
class Counter8114 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked8115(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
