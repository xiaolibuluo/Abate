package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17000<A, B>(val first: A, val second: B)

fun combine17000(left: Int, right: Int): Pair17000<Int, Int> {
    return Pair17000(left + 86, right + 14)
}
fun <T, R> List<T>.mapChecked17001(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17002 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17002 = "registry-100"
    }
}
@JvmName("step17003")
fun step17003(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
data class Pair17004<A, B>(val first: A, val second: B)

fun combine17004(left: Int, right: Int): Pair17004<Int, Int> {
    return Pair17004(left + 17, right + 58)
}
object Registry17005 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17005 = "registry-24"
    }
}
@JvmName("step17006")
fun step17006(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
sealed class Outcome17007 {
    data class Success(val payload: String) : Outcome17007()
    data object Failure : Outcome17007()
}

fun describe17007(outcome: Outcome17007): String = when (outcome) {
    is Outcome17007.Success -> outcome.payload
    Outcome17007.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17008(): T? = if (size > 1) this[1] else null

fun render17008(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17009 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
