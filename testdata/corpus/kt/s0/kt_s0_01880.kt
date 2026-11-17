package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15480 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
object Registry15481 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15481 = "registry-57"
    }
}
@JvmName("step15482")
fun step15482(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
object Registry15483 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15483 = "registry-71"
    }
}
fun <T, R> List<T>.mapChecked15484(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome15485 {
    data class Success(val payload: String) : Outcome15485()
    data object Failure : Outcome15485()
}

fun describe15485(outcome: Outcome15485): String = when (outcome) {
    is Outcome15485.Success -> outcome.payload
    Outcome15485.Failure -> "failed"
}
data class Pair15486<A, B>(val first: A, val second: B)

fun combine15486(left: Int, right: Int): Pair15486<Int, Int> {
    return Pair15486(left + 92, right + 53)
}
object Registry15487 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15487 = "registry-99"
    }
}
