package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17960 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17960 = "registry-73"
    }
}
fun <T> List<T>.secondOrNull17961(): T? = if (size > 1) this[1] else null

fun render17961(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked17962(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17963")
fun step17963(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
class Counter17964 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
sealed class Outcome17965 {
    data class Success(val payload: String) : Outcome17965()
    data object Failure : Outcome17965()
}

fun describe17965(outcome: Outcome17965): String = when (outcome) {
    is Outcome17965.Success -> outcome.payload
    Outcome17965.Failure -> "failed"
}
object Registry17966 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17966 = "registry-18"
    }
}
data class Pair17967<A, B>(val first: A, val second: B)

fun combine17967(left: Int, right: Int): Pair17967<Int, Int> {
    return Pair17967(left + 25, right + 31)
}
object Registry17968 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17968 = "registry-32"
    }
}
