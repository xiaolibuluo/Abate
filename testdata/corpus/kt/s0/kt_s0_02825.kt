package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23300 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked23301(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked23302(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair23303<A, B>(val first: A, val second: B)

fun combine23303(left: Int, right: Int): Pair23303<Int, Int> {
    return Pair23303(left + 52, right + 65)
}
sealed class Outcome23304 {
    data class Success(val payload: String) : Outcome23304()
    data object Failure : Outcome23304()
}

fun describe23304(outcome: Outcome23304): String = when (outcome) {
    is Outcome23304.Success -> outcome.payload
    Outcome23304.Failure -> "failed"
}
object Registry23305 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23305 = "registry-66"
    }
}
object Registry23306 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23306 = "registry-73"
    }
}
@JvmName("step23307")
fun step23307(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
